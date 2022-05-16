package service

import (
	"account_service/model"
	"account_service/repository"
	"log"

	"github.com/go-redis/redis"
)

func Check(token string) (authInfo *model.AuthStatus, err error) {
	authInfo, err = repository.GetAuthStatus(token)
	//空key 未登录
	if err == redis.Nil {
		log.Println(token, err)
		err = nil
	}
	return
}
