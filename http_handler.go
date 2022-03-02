package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/constant"
	"psy-consult-backend/exception"
	"psy-consult-backend/service"
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

func httpHandlerInit() {
	logrus.Info(constant.Main + "Init httpHandlerInit")
	// 支持跨域访问
	r.Use(Cors())
	r.GET("/ping", service.Ping)
	r.GET("/home", service.Home)
	imGroup := r.Group("/im")
	imGroup.Use(exception.ErrorHandlingMiddleware)
	{
		imGroup.GET("/sign", service.GetUserSign)
	}

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", service.Login)
		authGroup.POST("/logout", service.Logout)
	}

}
