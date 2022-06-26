package controller

import (
	"account_service/http/buz_code"
	"account_service/providers"
	"account_service/repository"
	"account_service/service"
	"account_service/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindJSON(ctx *gin.Context, req interface{}) (err error) {
	err = ctx.BindJSON(req)
	return
}
func BindQuery(ctx *gin.Context, form interface{}) (err error) {
	err = ctx.BindQuery(form)
	return
}
func BindMultiForm(ctx *gin.Context, form interface{}) (err error) {
	err = ctx.BindWith(form, binding.FormMultipart)
	return
}

//BindPhone 绑定手机号
func BindPhone(ctx *gin.Context) {
	uid, _ := ctx.Params.Get("uid")
	ctx.Writer.Write([]byte(uid))
}

//BindUserName 绑定用户名
func BindUsername(ctx *gin.Context) {

}
func LockUnlock(ctx *gin.Context) {

}

type RequestGetAssets struct {
	UID     string `form:"uid"`
	AppType int    `form:"app_type"`
	Token   string `form:"b_access_token"`
}

func InitApp(ctx *gin.Context) (res STDResponse, err error) {
	req := &RequestGetAssets{}
	if err = BindQuery(ctx, &req); err != nil {
		res.Code = buz_code.CODE_INVALID_ARGS
		res.Msg = err.Error()
		return
	}
	tx := providers.DBAccount.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	appID, err := service.GetUserAppID(tx, req.UID)
	if err != nil {
		return
	}
	// TODO:还未初始化，可以初始化
	if appID != "" {
		res.Code = buz_code.CODE_AUTH_FAILED
		err = errors.New("err:已初始化")
		return
	}
	if req.AppType == 1 {
		appID = utils.GenerateGroupID()
	} else {
		appID = utils.GenerateAppID()
	}
	_, err = service.UpdateUser(tx, req.UID, appID)
	res.Data = appID
	//TODO: 更新redis缓存appid
	token := ctx.Request.URL.Query().Get("b_access_token")
	err = repository.SetLoginStatus(req.UID, appID, token)

	return
}
