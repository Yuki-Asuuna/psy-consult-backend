package service

import (
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
	"psy-consult-backend/utils/redis"
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
	err := im_message.SendTextMessage(counsellorID, info.VisitorID, "你好，请问有什么可以帮您?")
	if err != nil {
		logrus.Errorf(constant.Service+"AddConversation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	ID, err := database.AddConversation(counsellorID, info.VisitorID)
	if err != nil {
		logrus.Errorf(constant.Service+"AddConversation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", ID))
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

	conversations, err := database.GetConversationList(page, size, startTime, endTime)
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationSearch Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	resp := make([]*api.ConversationSearchResponse, 0)
	for _, conv := range conversations {
		t := &api.ConversationSearchResponse{
			ConversationID: conv.ConversationID,
			StartTime:      conv.StartTime,
			EndTime:        conv.EndTime,
			Status:         conv.Status,
			IsHelped:       conv.IsHelped,
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
				EvaluationID:   evaluation.EvaluationID,
				ConversationID: evaluation.ConversationID,
				Rating:         evaluation.Rating,
				Evaluation:     evaluation.Evaluation,
			}
		}
		resp = append(resp, t)
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
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
	res, err := im_message.SearchAllHistoryMessage(conversationInfo.VisitorID, conversationInfo.CounsellorID, conversationInfo.StartTime.Unix(), conversationInfo.EndTime.Unix())
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationExport Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	// 设置列名 默认Sheet1
	_ = excel.SetCellValue("Sheet1", "A1", "时间")
	_ = excel.SetCellValue("Sheet1", "B1", "发送方")
	_ = excel.SetCellValue("Sheet1", "C1", "接收方")
	_ = excel.SetCellValue("Sheet1", "D1", "消息类型")
	_ = excel.SetCellValue("Sheet1", "E1", "消息内容")
	for index, msg := range res {
		idx := helper.I2S(index + 2)
		_ = excel.SetCellValue("Sheet1", "A"+idx, helper.Timestamp2S(msg.MsgTimeStamp))
		_ = excel.SetCellValue("Sheet1", "B"+idx, msg.FromAccount)
		_ = excel.SetCellValue("Sheet1", "C"+idx, msg.ToAccount)
		_ = excel.SetCellValue("Sheet1", "D"+idx, msg.MsgBody[0].MsgType)
		_ = excel.SetCellValue("Sheet1", "E"+idx, msg.MsgBody[0].MsgContent)
	}
	fileName := "./excel-export/" + conversationID + ".xlsx"
	err = excel.SaveAs(fileName)
	if err != nil {
		logrus.Errorf(constant.Service+"ConversationExport Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.File(fileName)
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}
