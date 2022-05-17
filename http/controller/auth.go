package controller

import (
	"account_service/http/buz_code"
	"account_service/http/request"
	"account_service/http/response"
	"account_service/library/env"
	"account_service/logger"
	"account_service/service"
	"fmt"
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
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
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
	/*
		username
		password
	*/

	req := request.SignInUsernameReques{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
		return
	}
	accessToken, uid, err := service.SignInUsername(req.Username, req.Password)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	buzCode := buz_code.CODE_OK
	msg := "ok"
	if accessToken == "" {
		buzCode = buz_code.CODE_USERNAME_PSWD_NOT_MATCH
		msg = "用户名密码不匹配"
		ctx.JSON(http.StatusOK, gin.H{"code": buzCode, "msg": msg})
		return
	}
	//
	ctx.JSON(http.StatusOK, gin.H{"code": buzCode, "msg": msg, "data": response.SignInUsernameRsp{
		UID:         uid,
		AccessToken: accessToken,
	}})
}

func SignInSMS(ctx *gin.Context) {
	ctx.Writer.Write([]byte("signin/SMS.post"))
}

func SignUpUsername(ctx *gin.Context) {
	/*
		username+pswd
	*/
	req := request.SignUpUsernameRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
		return
	}
	accessToken, uid, err := service.SignUpUsername(req.Username, req.Password)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  "server error",
		})
		return
	}
	rsp := gin.H{}
	if accessToken == "" {
		rsp["code"] = buz_code.CODE_USER_ALREADY_EXISTS
		rsp["msg"] = "用户名已存在"
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	//
	rsp["code"] = buz_code.CODE_OK
	rsp["msg"] = "ok"
	rsp["data"] = response.SignUpRsp{
		UID:         uid,
		AccessToken: accessToken,
	}
	ctx.JSON(http.StatusOK, rsp)
}
func SignUpSMS(ctx *gin.Context) {
}
