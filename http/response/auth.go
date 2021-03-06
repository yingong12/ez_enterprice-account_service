package response

type SignInUsernameRsp struct {
	AccessToken string `json:"b_access_token" example:"b_u_uasdasd"` //b端用户token
	UID         string `json:"uid" exmample:"zhuyan911"`             //用户ID
	AppID       string `json:"app_id" example:"app_asd"`             //appID
	AppType     uint8  `json:"app_type" example:"1"`                 //类型 0-企业 1-机构
}

type SignUpRsp struct {
	AccessToken string `json:"b_access_token" example:"b_u_uasdasd"` //b端用户token
	UID         string `json:"uid" exmample:"zhuyan911"`             //用户ID
}
