package models

type Like struct {
	BaseModel
	UserID    uint    `gorm:"index"`                // 用户id
	ArticleID uint    `gorm:"index"`                //文章id
	User      User    `gorm:"foreignkey:UserId"`    // 点赞用户
	Article   Article `gorm:"foreignkey:ArticleId"` // 点赞文章
}
