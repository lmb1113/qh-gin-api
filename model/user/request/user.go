package request

type LoginRequest struct {
	LoginType uint   `json:"login_type" binding:"required,oneof=1 2 3"` // 1 手机号登录 2邮箱登录 3密码登录
	Phone     string `json:"phone" binding:"required_if=LoginType 1,len=11"`
	Email     string `json:"email" binding:"required_if=LoginType 2,email"`
	Code      string `json:"code" binding:"required_if=LoginType 1 2,len=6"`
	Password  string `json:"password" binding:"required,min=8,max=20"`
}
