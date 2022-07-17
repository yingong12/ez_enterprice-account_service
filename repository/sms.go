package repository

import (
	"account_service/library/env"
	"account_service/providers"
	"time"
)

func SetSMSEntry(phone, verifyCode string) (ok bool, err error) {
	smsTimeMinute := env.GetIntVal("SMS_TIME_DURARTION_MINUTES")
	cmd := providers.RedisClient.SetNX(phone, verifyCode, (time.Duration(smsTimeMinute) * time.Minute))
	ok, err = cmd.Result()
	return
}

func GetSMSEntry(phone string) (res string, err error) {
	cmd := providers.RedisClient.Get(phone)
	res, err = cmd.Result()
	return
}
