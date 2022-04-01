package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/api"
	"psy-consult-backend/constant"
	"psy-consult-backend/database"
	"psy-consult-backend/exception"
	"psy-consult-backend/utils"
	"psy-consult-backend/utils/helper"
	"psy-consult-backend/utils/sessions"
)

func GetVisitorList(c *gin.Context) {
	page := helper.S2I(c.DefaultQuery("page", "0"))
	size := helper.S2I(c.DefaultQuery("size", "10"))
	name := c.DefaultQuery("name", "")
	list, err := database.GetVisitorUserList(page, size, name)
	if err != nil {
		logrus.Error(constant.Service+"GetVisitorList Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	vlist := make([]*api.VisitorInfoResponse, 0)
	for _, c := range list {
		vlist = append(vlist, &api.VisitorInfoResponse{
			VisitorID:        c.VisitorID,
			Username:         c.Username,
			Name:             c.Name,
			Status:           c.Status,
			Gender:           c.Gender,
			PhoneNumber:      c.PhoneNumber,
			LastLogin:        c.LastLogin,
			Email:            c.Email,
			EmergencyContact: c.EmergencyContact,
			EmergencyPhone:   c.EmergencyPhone,
			HasAgreed:        c.HasAgreed,
		})
	}
	cnt, err := database.GetVisitorUserListCount(name)
	if err != nil {
		logrus.Error(constant.Service+"GetVisitorList Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	resp := &api.GetVisitorListResponse{
		VisitorList: vlist,
		TotalCount:  cnt,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func BanVisitor(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	visitorID := params["visitorID"].(string)
	err := database.UpdateVisitorUserByVisitorID(visitorID, gin.H{"status": 1})
	if err != nil {
		logrus.Errorf(constant.Service+"BanVisitor Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", visitorID))
}

func ActivateVisitor(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	visitorID := params["visitorID"].(string)
	err := database.UpdateVisitorUserByVisitorID(visitorID, gin.H{"status": 0})
	if err != nil {
		logrus.Errorf(constant.Service+"BanVisitor Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", visitorID))
}
