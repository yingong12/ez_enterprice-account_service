package controller

import (
	"account_service/http/buz_code"
	"account_service/http/request"
	"account_service/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//发送验证码
func SendVerifyCode(ctx *gin.Context) {
	req := request.SendVerifyCodeRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_INVALID_ARGS,
			"msg":  fmt.Sprintf("invalid params %s\n", err.Error()),
		})
		return
	}
	ok, err := service.SendVerifyCode(req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_SERVER_ERROR,
			"msg":  err,
		})
	}
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": buz_code.CODE_TOO_MUCH_TRY_SMS,
			"msg":  "验证码未过期，勿频繁操作",
		})
		return
	}
	//
	ctx.JSON(http.StatusOK, gin.H{
		"code": buz_code.CODE_OK,
		"msg":  "ok"})
}
