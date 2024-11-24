package request

type EmailLogin struct {
	UserName string `json:"user_name" binding:"required" msg:"用户名为空"`
	Password string `json:"password" binding:"required" msg:"密码为空"`
}
