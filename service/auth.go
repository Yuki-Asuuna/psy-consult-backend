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

func Login(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	username := params["username"].(string)
	password := params["password"].(string)
	user_md5 := helper.S2MD5(username)

	// 通过md5生成counsellorID
	user, err := database.GetCounsellorUserByCounsellorID(user_md5)

	if err != nil {
		c.Error(exception.ServerError())
		logrus.Error(constant.Service+"Login Failed, err= %v", err)
		return
	}
	if user == nil {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-1, "Username not found", nil))
		return
	}
	// password = utils.S2MD5(password)
	if password != user.Password {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-1, "Incorrect Password", nil))
		return
	}
	session, _ := sessions.GetSessionClient().Get(c.Request, "dotcomUser")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	err = sessions.GetSessionClient().Save(c.Request, c.Writer, session)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Login Failed, err= %v", err)
		return
	}
	if err := database.UpdateLoginTimeByCounsellorID(helper.S2MD5(username)); err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Login Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func Logout(c *gin.Context) {
	session, _ := sessions.GetSessionClient().Get(c.Request, "dotcomUser")
	session.Values["authenticated"] = false
	err := sessions.GetSessionClient().Save(c.Request, c.Writer, session)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"Logout Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func Me(c *gin.Context) {
	user := sessions.GetCounsellorInfoBySession(c)
	if user == nil {
		c.Error(exception.ServerError())
		logrus.Error(constant.Service + "Me Get Personal Info Failed, user is nil")
		return
	}
	result := &api.MeResponse{
		Username:       user.Username,
		Name:           user.Name,
		Role:           user.Role,
		Status:         user.Status,
		Gender:         user.Gender,
		Age:            user.Age,
		IdentityNumber: user.IdentityNumber,
		PhoneNumber:    user.PhoneNumber,
		LastLogin:      user.LastLogin,
		Avatar:         user.Avatar,
		Email:          user.Email,
		Title:          user.Title,
		Department:     user.Department,
		Qualification:  user.Qualification,
		Introduction:   user.Introduction,
		MaxConsults:    user.MaxConsults,
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", result))
}

func ChangePassword(c *gin.Context) {
	user := sessions.GetCounsellorInfoBySession(c)
	if user == nil {
		c.Error(exception.ServerError())
		logrus.Error(constant.Service + "ChangePassword Get Personal Info Failed, user is nil")
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	oldPassword := params["oldPassword"].(string)
	newPassword := params["newPassword"].(string)
	if user.Password != oldPassword {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-1, "旧密码不正确", nil))
		return
	}
	err := database.UpdatePasswordByCounsellorID(user.CounsellorID, newPassword)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"ChangePassword Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}
