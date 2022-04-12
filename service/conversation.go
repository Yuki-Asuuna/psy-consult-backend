package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"net/http"
	"psy-consult-backend/api"
	"psy-consult-backend/constant"
	"psy-consult-backend/database"
	"psy-consult-backend/exception"
	"psy-consult-backend/tencent-im/im_message"
	"psy-consult-backend/utils"
	"psy-consult-backend/utils/helper"
	"psy-consult-backend/utils/rabbitmq"
	"psy-consult-backend/utils/redis"
	"psy-consult-backend/utils/sessions"
	"psy-consult-backend/utils/snowflake"
	"strings"
	"time"
)

func AddConversation(c *gin.Context) {
	sessionKey := c.Query("sessionKey")
	var info *database.VisitorUser
	if info = redis.GetVisitorInfoBySessionKey(sessionKey); info == nil {
		c.Error(exception.AuthError())
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	counsellorID := params["counsellorID"].(string)
	conversationID := snowflake.GenID()
	groupID, err := im_message.CreateNewGroup(info.VisitorID, counsellorID, helper.I642S(conversationID))
	if err != nil {
		logrus.Errorf(constant.Service+"AddConversation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	err = im_message.SendGroupMessage(groupID, counsellorID, "你好，请问有什么可以帮您？")
	if err != nil {
		logrus.Errorf(constant.Service+"AddConversation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	err = database.AddConversation(counsellorID, info.VisitorID, conversationID, groupID)
	if err != nil {
		logrus.Errorf(constant.Service+"AddConversation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	resp := api.AddConversationResponse{
		ConversationID: helper.I642S(conversationID),
		GroupID:        groupID,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func EndConversation(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	conversationID := int64(params["conversationID"].(float64))
	conversation, err := database.GetConversationByConversationID(conversationID)
	if err != nil || conversation == nil {
		logrus.Errorf(constant.Service+"EndConversation Failed, err= %v", err)
		c.Error(exception.ParameterError())
		return
	}
	err = database.UpdateConversationByConversationID(conversationID, map[string]interface{}{
		"end_time": time.Now(),
		"status":   1,
	})
	if err != nil {
		logrus.Errorf(constant.Service+"EndConversation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func Supervise(c *gin.Context) {
	// counsellor := sessions.GetCounsellorInfoBySession(c)
	params := make(map[string]interface{})
	c.BindJSON(&params)
	conversationID := helper.S2I64(params["conversationID"].(string))
	supervisorID := params["supervisorID"].(string)
	// 判断supervisor是否存在
	supervisor, err := database.GetCounsellorUserByCounsellorID(supervisorID)
	if err != nil {
		logrus.Errorf(constant.Service+"Supervise Failed, err= %v", err)
		c.Error(exception.ParameterError())
		return
	}
	if supervisor == nil {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "非法督导ID", nil))
		return
	}
	// 获取会话信息
	conversation, err := database.GetConversationByConversationID(conversationID)
	if err != nil {
		logrus.Errorf(constant.Service+"Supervise Failed, err= %v", err)
		c.Error(exception.ParameterError())
		return
	}
	if conversation == nil {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "非法会话ID", nil))
		return
	}
	err = im_message.AddGroupMember(conversation.GroupID, supervisorID)
	if err != nil {
		logrus.Errorf(constant.Service+"Supervise Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}

	// 发起会话（是否后端要参与？）
	//err = im_message.SendTextMessage(counsellor.CounsellorID, supervisorID, "我在咨询中遇到了一个问题，想向您求助")
	//if err != nil {
	//	logrus.Errorf(constant.Service+"Supervise Failed, err= %v", err)
	//	c.Error(exception.ServerError())
	//	return
	//}
	err = database.UpdateConversationByConversationID(conversation.ConversationID, map[string]interface{}{
		"is_helped":            1,
		"helped_supervisor_id": supervisorID,
	})
	if err != nil {
		logrus.Errorf(constant.Service+"Supervise Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func ConversationSearch(c *gin.Context) {
	var page, size, startTime, endTime int64
	var counsellorName, visitorName string
	qc := c.Query("page")
	if qc == "" {
		page = 0
	} else {
		page = helper.S2I64(qc)
	}

	qc = c.Query("size")
	if qc == "" {
		size = 10
	} else {
		size = helper.S2I64(qc)
	}

	qc = c.Query("startTime")
	if qc == "" {
		startTime = 0
	} else {
		startTime = helper.S2I64(qc)
	}

	qc = c.Query("endTime")
	if qc == "" {
		endTime = time.Now().Unix()
	} else {
		endTime = helper.S2I64(qc)
	}

	counsellorName = c.DefaultQuery("counsellorName", "")
	visitorName = c.DefaultQuery("visitorName", "")

	conversations, err := database.GetConversationList(page, size, startTime, endTime)
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationSearch Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	resp := make([]*api.ConversationSearchResponse, 0)
	for _, conv := range conversations {
		t := &api.ConversationSearchResponse{
			ConversationID: helper.I642S(conv.ConversationID),
			StartTime:      conv.StartTime,
			EndTime:        conv.EndTime,
			Status:         conv.Status,
			IsHelped:       conv.IsHelped,
			GroupID:        conv.GroupID,
		}
		counsellor, err := database.GetCounsellorUserByCounsellorID(conv.CounsellorID)
		if err != nil {
			logrus.Errorf(constant.Service+"ConversationSearch Failed, err= %v", err)
			c.Error(exception.ServerError())
			return
		}
		t.Counsellor = &api.CounsellorInfoResponse{
			CounsellorID:   counsellor.CounsellorID,
			Username:       counsellor.Username,
			Name:           counsellor.Name,
			Role:           counsellor.Role,
			Status:         counsellor.Status,
			Gender:         counsellor.Gender,
			Age:            counsellor.Age,
			IdentityNumber: counsellor.IdentityNumber,
			PhoneNumber:    counsellor.PhoneNumber,
			LastLogin:      counsellor.LastLogin,
			Avatar:         counsellor.Avatar,
			Email:          counsellor.Email,
			Title:          counsellor.Title,
			Department:     counsellor.Department,
			Qualification:  counsellor.Qualification,
			Introduction:   counsellor.Introduction,
			MaxConsults:    counsellor.MaxConsults,
		}

		visitor, err := database.GetVisitorUserByVisitorID(conv.VisitorID)
		if err != nil {
			logrus.Errorf(constant.Service+"ConversationSearch Failed, err= %v", err)
			c.Error(exception.ServerError())
			return
		}
		t.Visitor = &api.VisitorInfoResponse{
			VisitorID:        conv.VisitorID,
			Username:         visitor.Username,
			Name:             visitor.Name,
			Status:           visitor.Status,
			Gender:           visitor.Gender,
			PhoneNumber:      visitor.PhoneNumber,
			LastLogin:        visitor.LastLogin,
			Email:            visitor.Email,
			EmergencyContact: visitor.EmergencyContact,
			EmergencyPhone:   visitor.EmergencyPhone,
			HasAgreed:        visitor.HasAgreed,
		}

		if conv.IsHelped == 1 {
			supervisor, err := database.GetCounsellorUserByCounsellorID(conv.HelpedSupervisorID)
			if err != nil {
				logrus.Errorf(constant.Service+"ConversationSearch Failed, err= %v", err)
				c.Error(exception.ServerError())
				return
			}
			t.Supervisor = &api.CounsellorInfoResponse{
				CounsellorID:   supervisor.CounsellorID,
				Username:       supervisor.Username,
				Name:           supervisor.Name,
				Role:           supervisor.Role,
				Status:         supervisor.Status,
				Gender:         supervisor.Gender,
				Age:            supervisor.Age,
				IdentityNumber: supervisor.IdentityNumber,
				PhoneNumber:    supervisor.PhoneNumber,
				LastLogin:      supervisor.LastLogin,
				Avatar:         supervisor.Avatar,
				Email:          supervisor.Email,
				Title:          supervisor.Title,
				Department:     supervisor.Department,
				Qualification:  supervisor.Qualification,
				Introduction:   supervisor.Introduction,
				MaxConsults:    supervisor.MaxConsults,
				IsOnline:       redis.CheckOnline(supervisor.CounsellorID),
			}
		}
		evaluation, err := database.GetEvaluationByConversationID(conv.ConversationID)
		if err != nil {
			logrus.Errorf(constant.Service+"ConversationSearch Failed, err= %v", err)
			c.Error(exception.ServerError())
			return
		}
		if evaluation != nil {
			t.Evaluation = &api.EvaluationInfoResponse{
				EvaluationID:   helper.I642S(evaluation.EvaluationID),
				ConversationID: helper.I642S(evaluation.ConversationID),
				Rating:         evaluation.Rating,
				Evaluation:     evaluation.Evaluation,
			}
		}
		if visitorName != "" {
			if strings.Contains(t.Visitor.Name, visitorName) == false {
				continue
			}
		}
		if counsellorName != "" {
			if strings.Contains(t.Counsellor.Name, counsellorName) == false {
				continue
			}
		}
		resp = append(resp, t)
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func ConversationDetail(c *gin.Context) {
	conversationID := c.Query("conversationID")
	conversationInfo, err := database.GetConversationByConversationID(helper.S2I64(conversationID))
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationExport Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	if conversationInfo == nil {
		logrus.Warnf(constant.Service+"ConversationExport Failed, err= %v", "conversation does not exist")
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "conversation does not exist", nil))
		return
	}
	res, err := im_message.GetAllGroupMessage(conversationInfo.GroupID)
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationExport Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(0, utils.GenSuccessResponse(0, "OK", res))
}

func ConversationExport(c *gin.Context) {
	conversationID := c.Query("conversationID")
	excel := excelize.NewFile()
	conversationInfo, err := database.GetConversationByConversationID(helper.S2I64(conversationID))
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationExport Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	if conversationInfo == nil {
		logrus.Warnf(constant.Service+"ConversationExport Failed, err= %v", "conversation does not exist")
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "conversation does not exist", nil))
		return
	}
	res, err := im_message.GetAllGroupMessage(conversationInfo.GroupID)
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationExport Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	// 设置列名 默认Sheet1
	_ = excel.SetCellValue("Sheet1", "A1", "时间")
	_ = excel.SetCellValue("Sheet1", "B1", "发送方")
	_ = excel.SetCellValue("Sheet1", "C1", "消息类型")
	_ = excel.SetCellValue("Sheet1", "D1", "消息内容")
	for index, msg := range res {
		idx := helper.I2S(index + 2)
		_ = excel.SetCellValue("Sheet1", "A"+idx, helper.Timestamp2S(msg.MsgTimeStamp))
		_ = excel.SetCellValue("Sheet1", "B"+idx, msg.FromAccount)
		_ = excel.SetCellValue("Sheet1", "C"+idx, msg.MsgBody[0].MsgType)
		_ = excel.SetCellValue("Sheet1", "D"+idx, msg.MsgBody[0].MsgContent)
	}
	fileName := "./excel-export/" + conversationID + ".xlsx"
	err = excel.SaveAs(fileName)
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationExport Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.FileAttachment(fileName, conversationID+".xlsx")
	// c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func TodayStat(c *gin.Context) {
	counsellor := sessions.GetCounsellorInfoBySession(c)
	if counsellor == nil {
		c.Error(exception.ServerError())
		return
	}
	lst, err := database.GetConversationListByCounsellorIDAndTimeInterval(counsellor.CounsellorID, helper.GetTodayStartTimeStamp(), helper.GetTodayEndTimeStamp())
	if err != nil {
		logrus.Errorf(constant.Service+"TodayStat Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	if lst == nil {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", api.TodayStatResponse{}))
		return
	}
	var d time.Duration
	var processingCnt int
	for _, c := range lst {
		if c.Status == 0 {
			d += time.Now().Sub(c.StartTime)
			processingCnt += 1
			continue
		}
		delta := c.EndTime.Sub(c.StartTime)
		d += delta
	}
	seconds := int(d.Seconds())
	hour := int(seconds / 3600)
	minute := int((seconds - hour*3600) / 60)
	second := seconds - hour*3600 - minute*60
	resp := api.TodayStatResponse{
		TotalCount:        len(lst),
		Hour:              hour,
		Minute:            minute,
		Second:            second,
		InConversationCnt: processingCnt,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func TodayStatAll(c *gin.Context) {
	counsellor := sessions.GetCounsellorInfoBySession(c)
	if counsellor == nil || counsellor.Role != 0 {
		c.Error(exception.ServerError())
		return
	}
	lst, err := database.GetConversationList(0, 99999, helper.GetTodayStartTimeStamp(), helper.GetTodayEndTimeStamp())
	if err != nil {
		logrus.Errorf(constant.Service+"TodayStatAll Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	if lst == nil {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", api.TodayStatResponse{}))
		return
	}
	var d time.Duration
	for _, c := range lst {
		if c.Status == 0 {
			d += time.Now().Sub(c.StartTime)
			continue
		}
		delta := c.EndTime.Sub(c.StartTime)
		d += delta
	}
	seconds := int(d.Seconds())
	hour := int(seconds / 3600)
	minute := int((seconds - hour*3600) / 60)
	second := seconds - hour*3600 - minute*60
	resp := api.TodayStatResponse{
		TotalCount: len(lst),
		Hour:       hour,
		Minute:     minute,
		Second:     second,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func searchConversation(fromAcc, toAcc string) *database.Conversation {
	// status:0 进行中 status:1 已结束
	c, err := database.GetConversationBySender(fromAcc, toAcc)
	if err != nil {
		return nil
	}
	if c != nil {
		return c
	}
	c, err = database.GetConversationBySender(toAcc, fromAcc)
	if err != nil {
		return nil
	}
	return c
}

func Callback(c *gin.Context) {
	// appid := c.Query("SdkAppid")
	callbackCmd := c.Query("CallbackCommand")
	// 单聊回调
	if callbackCmd == "C2C.CallbackAfterSendMsg" {
		params := make(map[string]interface{})
		c.BindJSON(&params)
		fromAcc := params["From_Account"].(string)
		toAcc := params["To_Account"].(string)
		conversation := searchConversation(fromAcc, toAcc)
		if conversation == nil {
			logrus.Warnf(constant.Service+"Callback Ignored, msg does not belong to any conversation, fromAccount= %v, toAccount= %v", fromAcc, toAcc)
			c.JSON(http.StatusOK, utils.GenSuccessResponse(-4, "msg does not belong to any conversation", nil))
		}
		body, _ := json.Marshal(params)
		err := rabbitmq.PushMessage(helper.I642S(conversation.ConversationID), body)
		if err != nil {
			logrus.Errorf(constant.Service+"Callback Failed, err= %v", err)
			c.Error(exception.ServerError())
			return
		}
	} else {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "cmd does not exist", nil))
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func Pool(c *gin.Context) {
	conversationID := c.Query("conversationID")

	// 从mq中拉取一条消息
	msg, err := rabbitmq.GetMessage(conversationID)
	if err != nil {
		logrus.Errorf(constant.Service+"Pool Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}

	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", string(msg)))
}

func NStat(c *gin.Context) {
	n := helper.S2I(c.DefaultQuery("days", "7"))
	EndTime := helper.GetTodayEndTimeStamp()
	StartTime := helper.GetNDaysBefore(n)
	conversations, err := database.GetConversationList(0, 999999, StartTime, EndTime)
	if err != nil {
		logrus.Errorf(constant.Service+"Nstat Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	stat := make(map[string]int)
	tm, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	for _, c := range conversations {
		tmstr := c.StartTime.Format("2006-01-02")
		stat[tmstr] += 1
	}
	dateList := make([]string, 0)
	countList := make([]int, 0)
	for i := 0; i < n; i++ {
		tmstr := tm.Format("2006-01-02")
		val, ok := stat[tmstr]
		if !ok {
			countList = append(countList, 0)
		} else {
			countList = append(countList, val)
		}
		dateList = append(dateList, tmstr)
		tm = tm.AddDate(0, 0, -1)
		// countList = append(countList,stat)
	}
	helper.ReverseSlice(dateList)
	helper.ReverseSlice(countList)
	resp := api.NStatResponse{
		DateList:  dateList,
		CountList: countList,
	}
	c.JSON(0, utils.GenSuccessResponse(0, "OK", resp))
}
