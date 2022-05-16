package controller

import (
	"account_service/library/env"
	"account_service/logger"
	"account_service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Create	登录态系统
//@Summary	登录态校验
//@Description	登录态校验
//@Tags	登录态校验
//@Produce	json
//@Param  b_access_token header string true "b端用户token"
//@Success 200 {object} model.AuthStatus
//@Router	/auth/check [get]
func Check(ctx *gin.Context) {
	token := ctx.GetHeader(env.GetStringVal("TOKEN_KEY"))
	//参数校验
	if token == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "缺少token",
		})
		return
	}
	//buz逻辑
	authInfo, err := service.Check(token)
	if err != nil {
		logger.Error(authInfo)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "server error",
		})
		return
	}
	code := 0
	//为nil时没有登录
	if authInfo == nil {
		code = 1
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "ok",
		"data": authInfo,
	})
}

func SignInUsername(ctx *gin.Context) {
	ctx.Writer.Write([]byte("signin/username.post"))
}

func SignInSMS(ctx *gin.Context) {
	ctx.Writer.Write([]byte("signin/SMS.post"))
}

func SignUpUsername(ctx *gin.Context) {

}
func SignUpSMS(ctx *gin.Context) {
}
