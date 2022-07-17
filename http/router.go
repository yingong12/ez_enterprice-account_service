package http

import (
	_ "account_service/docs" //引入swagger
	"account_service/http/controller"
	"account_service/http/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // register swagger
	//日志中间件
	router.Use(middleware.RequestLogger())
	router.Use(middleware.ControllerErrorLogger())
	//routes
	router.POST("health", controller.Health)
	//swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // register swagger
	//登录模块
	auth := router.Group("/auth")
	{
		/**
		1. 登录态校验  token-> code 0:uid,app_id  code 1: 过期   code 2:wrong token code 3:token missing
		2. 注册： username,pswd,phone(带验证码) -> token
		3. 手机登录： phone#, veriCode -> token
		*/
		auth.GET("/check", controller.STDwrapperJSON(controller.Check))                     //校验登录态
		auth.POST("/signin/username", controller.STDwrapperJSON(controller.SignInUsername)) //用户名登录
		auth.POST("/signin/sms", controller.STDwrapperJSON(controller.SignInSMS))           //手机登录
		auth.POST("/signup/username", controller.STDwrapperJSON(controller.SignUpUsername)) //用户名注册
		auth.POST("/signup/sms", controller.STDwrapperJSON(controller.SignUpSMS))           //手机注册
		auth.POST("/send_sms_code", controller.STDwrapperJSON(controller.SendVerifyCode))   //向sms服务申请验证码
	}
	//账号模块
	account := router.Group("/account")
	{
		account.GET("/init_app", controller.STDwrapperJSON(controller.InitApp)) //初始化企业，绑定空企业或者店铺
	}
	return
}
