package models

type Comment struct {
	BaseModel
	Article  Article   `gorm:"foreignkey:ArticleId"`
	Comments []Comment `gorm:"foreignkey:CommentId"`
	User     User      `gorm:"foreignkey:UserId"`
	Content  string    `gorm:"type:text"`
}
