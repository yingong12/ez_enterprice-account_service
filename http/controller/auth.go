package controller

import (
	"account_service/http/buz_code"
	"account_service/http/request"
	"account_service/http/response"
	"account_service/library/env"
	"account_service/service"

	"github.com/gin-gonic/gin"
)

//Create	登录态
//@Summary	登录态校验
//@Description	登录态校验
//@Tags	登录态校验
//@Produce	json
//@Param  b_access_token header string true "b端用户token"
//@Success 200 {object} model.AuthStatus
//@Router	/auth/check [get]
func Check(ctx *gin.Context) (res STDResponse, err error) {
	token := ctx.GetHeader(env.GetStringVal("TOKEN_KEY"))
	//参数校验
	if token == "" {
		res.Code = buz_code.CODE_NO_TOKEN
		res.Msg = "no token"
		return
	}
	//buz逻辑
	authInfo, err := service.Check(token)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	//为nil时没有登录
	if authInfo.AppID == "" && authInfo.UID == "" {
		res.Code = buz_code.CODE_AUTH_FAILED
		res.Msg = "未登录"
	}
	res.Data = authInfo
	return
}

//Create	注册登录
//@Summary	用户名登录
//@Description	用户名登录
//@Tags
//@Produce	json
//@Param xxx body request.SignInUsernameRequest  false "注释"
//@Success 200 {object} response.SignInUsernameRsp
//@Router	/signin/username [post]
func SignInUsername(ctx *gin.Context) (res STDResponse, err error) {
	req := request.SignInUsernameRequest{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	accessToken, uid, appID, err := service.SignInUsername(req.Username, req.Password)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	if accessToken == "" {
		res.Code = buz_code.CODE_USERNAME_PSWD_NOT_MATCH
		res.Msg = "用户名密码不匹配"
		return
	}
	rsp := response.SignInUsernameRsp{
		UID:         uid,
		AccessToken: accessToken,
		AppID:       appID,
	}
	if len(appID) > 0 && appID[0] == 'g' {
		rsp.AppType = 1
	}
	res.Data = rsp
	return
}

//Create	注册登录
//@Summary	用户名注册
//@Description	用户名注册
//@Tags
//@Produce	json
//@Param xxx body request.SignUpUsernameRequest false "注释"
//@Success 200 {object} response.SignUpRsp
//@Router	/signup/username [post]
func SignUpUsername(ctx *gin.Context) (res STDResponse, err error) {
	/*
		username+pswd
	*/
	req := request.SignUpUsernameRequest{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	accessToken, uid, err := service.SignUpUsername(req.Username, req.Phone, req.VerifyCode, req.Password)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	if accessToken == "" {
		res.Code = buz_code.CODE_USER_ALREADY_EXISTS
		res.Msg = "用户名已存在"
		return
	}
	res.Data = response.SignUpRsp{
		UID:         uid,
		AccessToken: accessToken,
	}
	return
}

func SignUpSMS(ctx *gin.Context) (res STDResponse, err error) {
	//sms注册
	/**
	checkCode(code) -> 注册流程
	*/
	// req := request.SignUpSMSRequest{}
	// if err = ctx.BindJSON(&req); err != nil {
	// 	res.Code = buz_code.CODE_INVALID_ARGS
	// 	res.Msg = err.Error()
	// 	return
	// }
	// accessToken, uid, err, buzCode := service.SinUpSMS(req.Phone, req.VerifyCode)
	// if err != nil {
	// 	res.Code = buz_code.CODE_SERVER_ERROR
	// 	res.Msg = "server error"
	// 	return
	// }
	// rsp := gin.H{}
	// if accessToken == "" {
	// 	rsp["code"] = buzCode
	// 	rsp["msg"] = ""
	// 	ctx.JSON(http.StatusOK, rsp)
	// 	return
	// }
	// res.Data = response.SignUpRsp{
	// 	UID:         uid,
	// 	AccessToken: accessToken,
	// }
	return
}

func SignInSMS(ctx *gin.Context) (res STDResponse, err error) {
	req := request.SignInSMSRequest{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	accessToken, uid, err, buzCode := service.SignInSMS(req.Phone, req.VerifyCode)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	res.Code = buzCode
	if accessToken == "" {
		res.Msg = ""
		return
	}
	//
	res.Data = response.SignInUsernameRsp{
		UID:         uid,
		AccessToken: accessToken,
	}
	return
}
