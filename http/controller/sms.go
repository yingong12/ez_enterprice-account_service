package controller

import (
	"account_service/http/buz_code"
	"account_service/http/request"
	"account_service/service"

	"github.com/gin-gonic/gin"
)

//发送验证码
func SendVerifyCode(ctx *gin.Context) (res *STDResponse, err error) {
	req := request.SendVerifyCodeRequest{}
	if err = ctx.BindJSON(&req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	ok, err := service.SendVerifyCode(req.Phone)
	if err != nil {
		res.Code = buz_code.CODE_SERVER_ERROR
		res.Msg = "server error"
		return
	}
	if !ok {
		res.Code = buz_code.CODE_TOO_MUCH_TRY_SMS
		res.Msg = "验证码未过期，勿频繁操作"
		return
	}
	//
	return
}
