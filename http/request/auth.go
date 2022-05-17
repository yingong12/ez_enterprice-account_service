package request

type SignUpUsernameRequest struct {
	BaseRequest
	Username string `json:"username" example:"zhuyan"`         //用户名
	Password string `json:"password" exmample:"123456@Zhuyan"` //密码，需要包含大小写数字和特殊字符
}

type SignInUsernameReques struct {
	BaseRequest
	Username string `json:"username" example:"zhuyan"`         //用户名
	Password string `json:"password" exmample:"123456@Zhuyan"` //密码，需要包含大小写数字和特殊字符
}
