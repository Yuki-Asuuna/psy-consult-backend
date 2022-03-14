package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/constant"
	"psy-consult-backend/database"
	"psy-consult-backend/exception"
	"psy-consult-backend/utils"
)

func AddEvaluation(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	conversationID := int64(params["conversationID"].(float64))
	rating := int64(params["rating"].(float64))
	evaluation := params["evaluation"].(string)
	conversation, err := database.GetConversationByConversationID(conversationID)
	if err != nil {
		logrus.Error(constant.Service+"AddEvaluation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	if conversation == nil {
		c.Error(exception.ParameterError())
		return
	}
	e, err := database.AddEvaluation(rating, evaluation, conversationID)
	if err != nil {
		logrus.Error(constant.Service+"AddEvaluation Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	if e == nil {
		logrus.Error(constant.Service+"AddEvaluation Failed, err= %v", "evaluation is nil")
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", e.EvaluationID))
}
