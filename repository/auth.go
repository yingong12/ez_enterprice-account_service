package repository

import (
	"account_service/model"
	"account_service/providers"
)

func GetAuthStatus(token string) (as *model.AuthStatus, err error) {
	val, err := providers.RedisClient.Get(token).Result()
	if err != nil {
		return
	}
	as = &model.AuthStatus{}
	as.UID = val
	return
}
