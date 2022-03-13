package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/constant"
	"psy-consult-backend/database"
	"psy-consult-backend/exception"
	"psy-consult-backend/tencent-im/im_message"
	"psy-consult-backend/utils"
	"psy-consult-backend/utils/helper"
	"psy-consult-backend/utils/sessions"
	"time"
)

func AddConversation(c *gin.Context) {
	var info *database.VisitorUser
	if info = sessions.GetWeChatUserInfoBySession(c); info == nil {
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
	conversationID := helper.S2I64(params["conversationID"].(string))
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
