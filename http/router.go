package http

import (
	"account_service/http/controller"

	"github.com/gin-gonic/gin"
)

func loadRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.New()
	//routes
	router.POST("healthy", controller.Healthy)
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
		auth.POST("/check.post", controller.Check)                    //校验登录态
		auth.POST("/signin/username.post", controller.SignInUsername) //用户名登录
		auth.POST("/signin/sms.post", controller.SignInSMS)           //手机登录
		auth.POST("/signup/username.post", controller.SignUpUsername) //用户名注册
		auth.POST("/signup/sms.post", controller.SignUpSMS)           //手机注册
		// flow.Use(providers.ServiceOrder())
	}
	//账号模块
	account := router.Group("/account")
	{
		//绑定手机号和用户名
		account.POST("/bind/username.post", controller.BindPhone)
		account.POST("/bind/phone.post", controller.BindUsername)
	}
	//sms验证码模块
	sms := router.Group("sms")
	{
		sms.POST("/ask_code.post", controller.AskCode) //向sms服务申请验证码
	}
	return
}
