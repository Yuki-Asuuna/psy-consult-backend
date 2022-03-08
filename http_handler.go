package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psy-consult-backend/constant"
	"psy-consult-backend/exception"
	"psy-consult-backend/service"
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

func httpHandlerInit() {
	logrus.Info(constant.Main + "Init httpHandlerInit")
	// 支持跨域访问
	r.Use(Cors())
	// 错误处理
	r.Use(exception.ErrorHandlingMiddleware)

	r.GET("/ping", service.Ping)
	r.PUT("/image_upload", service.ImageUpload)
	imGroup := r.Group("/im")
	{
		// imGroup.POST("/login")
		// imGroup.GET("me")
		imGroup.GET("/sign", service.GetUserSign)
	}

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", service.Login)
		authGroup.POST("/logout", AuthMiddleWare(), service.Logout)
		authGroup.GET("/me", AuthMiddleWare(), service.Me)
		authGroup.POST("/password", AuthMiddleWare(), service.ChangePassword)
	}

	userGroup := r.Group("/user")
	{
		userGroup.PUT("/ms", service.AdminPutMs)
		userGroup.POST("/ms", service.AdminPostMs)
		userGroup.DELETE("/ms", service.AdminDeleteMs)
		userGroup.GET("/superuser_get", service.SuperuserGet)
		userGroup.GET("/list", service.GetCounsellorList)
		userGroup.PUT("/bind", service.AddBinding)
		userGroup.DELETE("/bind", service.DeleteBinding)
		userGroup.GET("/bind", service.GetBinding)

	}

	visitorGroup := r.Group("/visitor")
	{
		visitorGroup.GET("/list", service.GetVisitorList)
		visitorGroup.POST("/ban", service.BanVisitor)
		visitorGroup.POST("/activate", service.ActivateVisitor)
	}

	arrangeGroup := r.Group("/arrange")
	{
		arrangeGroup.PUT("/add", service.PutArrange)
		arrangeGroup.GET("/get", service.GetArrange)
		arrangeGroup.DELETE("/delete", service.DeleteArrange)
	}

}
