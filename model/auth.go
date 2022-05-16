package model

type AuthStatus struct {
	UID       string `json:"uid" example:"u_12345678901"`             //b端用户id
	ExpiresAt string `json:"expire_at" example:"2022-05-16 23:00:00"` //过期时间
}
