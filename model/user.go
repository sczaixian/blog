package model

type User struct {
	BaseModel
	Name     string `gorm:"size:50; not null; unique; index; comment:用户名" json:"name"  validate:"required,min=3,max=50,alphanum"`
	Email    string `gorm:"size:100; comment:邮箱" json:"email" validate:"required,email"`
	Password string `gorm:"size:255; not null comment:密码" json:"-" validate:"required,min=8"`
	Role     string `gorm:"size:20;default:'user'" json:"role" validate:"oneof=user editor admin"`

	Posts    []Post    `gorm:"foreignkey:AuthorID" json:"-"`
	Comments []Comment `gorm:"foreignkey:UserID" json:"-"`
}

func (u *User) TableName() string {
	return "users"
}
