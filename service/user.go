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

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", "hello,"+sessions.GetUserNameBySession(c)))
}
