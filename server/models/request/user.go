package request

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	// todo：验证码
}

type Register struct {
	Username string `json:"username" example:"用户名"`
	Password string `json:"password" example:"密码"`
}

type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type ResetPassword struct {
	ID       uint   `json:"ID" form:"ID"`
	Password string `json:"password" form:"password" gorm:"comment:用户登录密码"` // 用户登录密码
}
