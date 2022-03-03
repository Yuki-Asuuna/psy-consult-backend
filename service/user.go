package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/constant"
	"psy-consult-backend/database"
	"psy-consult-backend/exception"
	"psy-consult-backend/utils"
	"psy-consult-backend/utils/helper"
	"psy-consult-backend/utils/sessions"
)

// 管理员添加后台用户
func AdminPutMs(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}

	params := make(map[string]interface{})
	c.BindJSON(&params)
	username := params["username"].(string)
	password := params["password"].(string)
	role := int(params["role"].(float64))

	user, err := database.GetCounsellorUserByCounsellorID(helper.S2MD5(username))
	if err != nil {
		logrus.Error(constant.Service+"AdminPutMs failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	// 用户已存在
	if user != nil {
		logrus.Infof(constant.Service + "AdminPutMs failed, user exists")
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "用户已存在", nil))
		return
	}

	err = database.AddCounsellorUser(username, password, role)
	if err != nil {
		logrus.Error(constant.Service+"AdminPutMs Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

// 管理员 更新后台用户信息
func AdminPostMs(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}

	params := make(map[string]interface{})
	c.BindJSON(&params)
	username := params["username"].(string)
	name := params["name"].(string)
	password := params["password"].(string)
	status := int(params["status"].(float64))
	gender := int(params["gender"].(float64))
	age := int(params["age"].(float64))
	identityNumber := params["identityNumber"].(string)
	phoneNumber := params["phoneNumber"].(string)
	avatar := params["avatar"].(string)
	email := params["email"].(string)
	title := params["title"].(string)
	department := params["department"].(string)
	qualification := params["qualification"].(string)
	introduction := params["introduction"].(string)
	maxConsults := int(params["maxConsults"].(float64))

	counsellorID := helper.S2MD5(username)
	err := database.UpdateCounsellorUserByCounsellorID(counsellorID, name, password, status, gender, age, identityNumber, phoneNumber, avatar, email, title, department, qualification, introduction, maxConsults)
	if err != nil {
		logrus.Error(constant.Service+"AdminPostMs Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}
