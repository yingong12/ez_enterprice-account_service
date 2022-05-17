package service

import (
	"account_service/model"
	"account_service/providers"
	"account_service/repository"
	"account_service/utils"
	"log"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func Check(accessToken string) (authInfo *model.AuthStatus, err error) {
	authInfo, err = repository.GetAuthStatus(accessToken)
	//空key 未登录
	if err == redis.Nil {
		log.Println(accessToken, err)
		err = nil
	}
	return
}

func SignUpUsername(username, pswd string) (accessToken, uid string, err error) {
	tx := providers.DBAccount.Begin()
	defer func() {
		//网络错误或者已被注册
		if err != nil || uid == "" {
			tx.Rollback()
			return
		}
		//没问题后注册用户并登录
		accessToken, err = setLoginStatus(uid)
		tx.Commit()
	}()
	//没被注册才继续
	if err = repository.GetUserByKey("username", username); err != gorm.ErrRecordNotFound {
		return
	}
	uid, err = repository.InsertUser(username, pswd)
	return
}

//TODO:这里要做同时valid token数量的限制
func SignInUsername(username, pswd string) (accessToken, uid string, err error) {
	m := map[string]interface{}{
		"username": username,
		"pswd":     pswd,
	}
	usr, err := repository.GetUserByKeys(m)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	uid = usr.UID
	accessToken, err = setLoginStatus(uid)
	return

}

func setLoginStatus(uid string) (accessToken string, err error) {
	accessToken = utils.GenerateAccessToken()
	err = repository.SetLoginStatus(uid, accessToken)
	return
}
