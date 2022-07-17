package providers

import (
	"account_service/library"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
)

var RedisClient *library.RedisClient

var DBAccount *library.GormDB

var SMSClient *dysmsapi20170525.Client
