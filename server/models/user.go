package models

// 用户

type Login interface {
	GetUserName() string
	GetUserId() uint
	GetUserInfo() any
}

var _ Login = new(User)

type User struct {
	BaseModel
	UserName      string         `gorm:"type:varchar(50);not null;unique;comment:用户名" json:"user_name"`
	Email         string         `gorm:"size:100;unique;comment:邮箱" json:"email"`
	Password      string         `gorm:"type:varchar(255);not null;comment:密码" json:"-"`
	Avatar        string         `gorm:"default:xxxx;comment:头像" json:"avatar"`
	Bio           string         `gorm:"type:text;comment:个人简介" json:"bio"`
	Status        int            `gorm:"default:1;comment:状态（1：正常，0：禁用）" json:"status"`
	LastLoginIP   string         `gorm:"size:50;comment:最后登录ip" json:"last_login_ip"`
	Articles      []Article      `gorm:"foreignkey:UserID;comment:文章" json:"articles"`
	Followers     []UserFollow   `gorm:"foreignkey:FollowingID" json:"followers"`  // 粉丝
	Followings    []UserFollow   `gorm:"foreignkey:FollowingID" json:"followings"` // 关注的人
	Notifications []Notification `gorm:"foreignKey:UserID"`                        // 用户通知
}

type UserFollow struct {
	BaseModel
	FollowerID  uint `gorm:"index"` // 关注者ID
	FollowingID uint `gorm:"index"` // 被关注者ID

	// Relationships
	Follower  User `gorm:"foreignKey:FollowerID"`  // 关注者
	Following User `gorm:"foreignKey:FollowingID"` // 被关注者
}

func (User) TableName() string {
	return "users"
}

func (UserFollow) TableName() string {
	return "user_follows"
}

func (s *User) GetUserName() string {
	return s.UserName
}
func (s *User) GetUserId() uint {
	return s.ID
}
func (s *User) GetUserInfo() any {
	return *s
}
