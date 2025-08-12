package models

type Notification struct {
	BaseModel
	UserID  uint   `gorm:"index"`                                      //接收用户ID
	Type    string `gorm:"size:50;not null; comment:通知类型" json:"type"` //通知类型(comment, like, follow等)
	Message string `gorm:"type:text;not null" json:"message"`          //通知内容
	IsRead  bool   `gorm:"default:false" json:"is_read"`               //是否已读

	User User `gorm:"foreignkey:UserID" json:"user"` //接收用户
}
