package model

type Comment struct {
	BaseModel
	Content  string     `gorm:"type:text; not null; comment:评论内容" json:"content" validate:"required,min=2,max=1000"`
	UserID   uint       `gorm:"index; comment:用户ID" json:"user_id"`
	PostID   uint       `gorm:"index; comment:文章ID" json:"post_id"`
	User     User       `gorm:"foreignkey:UserID" json:"user"`
	Post     Post       `gorm:"foreignkey:PostID" json:"-"`
	ParentID *uint      `json:"parent_id"`
	Replies  []*Comment `gorm:"foreignkey:ParentID" json:"replies"`
}
