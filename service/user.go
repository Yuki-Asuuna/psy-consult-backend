package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/api"
	"psy-consult-backend/constant"
	"psy-consult-backend/database"
	"psy-consult-backend/exception"
	"psy-consult-backend/tencent-im/account_manage"
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
	err = account_manage.AddIMSDKAccount(helper.S2MD5(username), username, "")
	if err != nil {
		logrus.Error(constant.Service+"AdminPutMs Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", username))
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

func SuperuserGet(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}
	counsellorID := c.Query("counsellorID")
	user, err := database.GetCounsellorUserByCounsellorID(counsellorID)
	if err != nil {
		logrus.Error(constant.Service+"SuperuserGet, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	resp := api.SuperuserGetResponse{
		CounsellorID:   user.CounsellorID,
		Username:       user.Username,
		Name:           user.Name,
		Password:       user.Password,
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
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func AdminDeleteMs(c *gin.Context) {
	// 判断有无管理员权限
	if info := sessions.GetCounsellorInfoBySession(c); info == nil || info.Role != 0 {
		c.Error(exception.AuthError())
		return
	}
	params := make(map[string]interface{})
	c.BindJSON(&params)
	counsellorID := params["counsellorID"].(string)
	err := database.DeleteCounsellorUserByCounsellorID(counsellorID)
	if err != nil {
		logrus.Error(constant.Service+"AdminDeleteMs, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	err = account_manage.DeleteIMSDKAccount(counsellorID)
	if err != nil {
		logrus.Error(constant.Service+"AdminDeleteMs Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", counsellorID))
}

func GetCounsellorList(c *gin.Context) {
	page := helper.S2I(c.DefaultQuery("page", "0"))
	size := helper.S2I(c.DefaultQuery("size", "10"))
	role := helper.S2I(c.DefaultQuery("role", "0"))
	list, err := database.GetCounsellorUserList(page, size, role)
	if err != nil {
		logrus.Error(constant.Service+"GetCounsellorList Failed, err= %v", err)
		c.Error(exception.ServerError())
		return
	}
	resp := make([]*api.CounsellorInfoResponse, 0)
	for _, c := range list {
		resp = append(resp, &api.CounsellorInfoResponse{
			CounsellorID:   c.CounsellorID,
			Username:       c.Username,
			Name:           c.Name,
			Role:           c.Role,
			Status:         c.Status,
			Gender:         c.Gender,
			Age:            c.Age,
			IdentityNumber: c.IdentityNumber,
			PhoneNumber:    c.PhoneNumber,
			LastLogin:      c.LastLogin,
			Avatar:         c.Avatar,
			Email:          c.Email,
			Title:          c.Title,
			Department:     c.Department,
			Qualification:  c.Qualification,
			Introduction:   c.Introduction,
			MaxConsults:    c.MaxConsults,
		})
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}

func AddBinding(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	supervisorID := params["supervisorID"].(string)
	counsellorID := params["counsellorID"].(string)
	info := sessions.GetCounsellorInfoBySession(c)

	// 管理员 或者 当前后台用户为咨询师且咨询师的ID与json body一致
	if info.Role == 0 || (info.CounsellorID == counsellorID && info.Role == 1) {
		supervisorInfo, err := database.GetCounsellorUserByCounsellorID(supervisorID)
		if err != nil {
			logrus.Error(constant.Service+"AddBinding Failed, err= %v", err)
			c.Error(exception.ServerError())
			return
		}
		if supervisorInfo == nil {
			c.JSON(http.StatusOK, utils.GenSuccessResponse(-2, "非法ID", nil))
			return
		}
		err = database.AddBinding(counsellorID, supervisorID)
		if err != nil {
			logrus.Error(constant.Service+"AddBinding Failed, err= %v", err)
			c.Error(exception.ServerError())
			return
		}
		c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
	} else {
		c.JSON(http.StatusOK, utils.GenSuccessResponse(-3, "权限错误", nil))
		return
	}
}

func DeleteBinding(c *gin.Context) {
	params := make(map[string]interface{})
	c.BindJSON(&params)
	bindingID := int64(params["bindingID"].(float64))
	err := database.DeleteBinding(bindingID)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"DeleteBinding Failed, err= %v", err)
		return
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", nil))
}

func GetBinding(c *gin.Context) {
	counsellorID := c.Query("counsellorID")
	bindings, err := database.GetBindingByCounsellorID(counsellorID)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"GetBinding Failed, err= %v", err)
		return
	}
	supervisorIDList := make([]string, 0)
	for _, binding := range bindings {
		supervisorIDList = append(supervisorIDList, binding.SupervisorID)
	}
	counsellorID2InfoMap, err := database.GetCounsellorUsersByCounsellorIDs(supervisorIDList)
	if err != nil {
		c.Error(exception.ServerError())
		logrus.Errorf(constant.Service+"GetBinding Failed, err= %v", err)
		return
	}
	resp := make([]*api.BindingInfoResponse, 0)
	for _, binding := range bindings {
		supervisor, _ := counsellorID2InfoMap[binding.SupervisorID]
		if supervisor == nil {
			continue
		}
		resp = append(resp, &api.BindingInfoResponse{
			BindingID:    binding.BindingID,
			SupervisorID: binding.SupervisorID,
			Name:         supervisor.Name,
		})
	}
	c.JSON(http.StatusOK, utils.GenSuccessResponse(0, "OK", resp))
}
