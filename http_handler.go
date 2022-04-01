package main

import (
	"github.com/sirupsen/logrus"
	"psy-consult-backend/constant"
	"psy-consult-backend/exception"
	"psy-consult-backend/middleware"
	"psy-consult-backend/service"
)

func httpHandlerInit() {
	logrus.Info(constant.Main + "Init httpHandlerInit")
	// 支持跨域访问
	r.Use(middleware.Cors())
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
		authGroup.POST("/wx_login", service.WxLogin)
		authGroup.POST("/wx_update", service.WxUpdate)
		authGroup.GET("/wx_me", service.WxMe)
		authGroup.POST("/logout", middleware.AuthMiddleWare(), service.Logout)
		authGroup.GET("/me", middleware.AuthMiddleWare(), service.Me)
		authGroup.POST("/me", middleware.AuthMiddleWare(), service.PostMe)
		authGroup.POST("/password", middleware.AuthMiddleWare(), service.ChangePassword)
		authGroup.POST("/avatar_upload", middleware.AuthMiddleWare(), service.AvatarUpload)
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

	conversationGrop := r.Group("/conversation")
	{
		conversationGrop.GET("/today_stat", service.TodayStat)
		conversationGrop.PUT("/start", service.AddConversation)
		conversationGrop.POST("/end", service.EndConversation)
		conversationGrop.POST("/supervise", service.Supervise)
		conversationGrop.GET("/search", service.ConversationSearch)
		conversationGrop.GET("/export", service.ConversationExport)
		conversationGrop.GET("/detail", service.ConversationDetail)
		conversationGrop.POST("/callback", service.Callback)
		conversationGrop.GET("/poll", service.Pool)
	}

	evaluateGroup := r.Group("/evaluate")
	{
		evaluateGroup.PUT("/add", service.AddEvaluation)
	}

}
