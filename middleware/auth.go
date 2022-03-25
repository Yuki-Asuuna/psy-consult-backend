package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/constant"
	"psy-consult-backend/exception"
	"psy-consult-backend/utils/sessions"
)

// 允许跨域访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// 鉴权中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := sessions.GetSessionClient()
		session, err := store.Get(c.Request, "dotcomUser")
		if err != nil {
			c.Error(exception.AuthError())
			logrus.Errorf(constant.Service+"AuthMiddleWare Store Get Session Failed, err= %v", err)
			c.Abort()
		}
		if session.IsNew {
			c.Error(exception.AuthError())
			logrus.Errorf(constant.Service+"AuthMiddleWare Store Session Is New, err= %v", err)
			c.Abort()
		}
		if isauth, ok := session.Values["authenticated"].(bool); !ok || !isauth {
			c.Error(exception.AuthError())
			logrus.Infof(constant.Service + "AuthMiddleWare Store Values Is False")
			c.Abort()
		}
		c.Next()
	}
}
