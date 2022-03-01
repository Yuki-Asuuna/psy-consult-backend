package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"psy-consult-backend/exception"
	"psy-consult-backend/tencent-im"
	"psy-consult-backend/tencent-im/usersig"
)

const (
	expire_time = 3600 // 默认颁发用户令牌有效时长为3600秒
)

func GetUserSign(c *gin.Context) {
	username := c.Query("username")
	result, err := usersig.GenUserSig(tencent_im.SDKAppID, tencent_im.SDKSecretKey, username, expire_time)
	if err != nil {
		c.Error(exception.ServerError())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"result":  result,
	})
}
