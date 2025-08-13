package request

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	// todo：验证码
}

type Register struct {
	Username string `json:"username" example:"用户名"`
	Password string `json:"password" example:"密码"`
	Email    string `json:"email" example:"邮箱"`
}

type ChangePasswordReq struct {
	ID          uint   `json:"id"`          // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type ResetPassword struct {
	ID       uint   `json:"id" form:"id"`
	Password string `json:"password" form:"password" gorm:"comment:用户登录密码"` // 用户登录密码
}
