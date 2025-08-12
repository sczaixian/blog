package models

type Comment struct {
	BaseModel
	IP     string `gorm:"type:varchar(50);not null;comment:评论IP" json:"ip"`
	Status string `gorm:"default:1;comment:状态(1正常;0待审核;-1屏蔽)" json:"status"`
	//Comments []Comment `gorm:"foreignkey:CommentID"`
	Content string `gorm:"type:text;not null;comment:评论内容" json:"content"`

	UserID    uint `gorm:"index"` // 评论用户
	ArticleID uint `gorm:"index"` // 文章id
	ParentID  uint `gorm:"index"` // 父评论id

	User    User      `gorm:"foreignkey:UserID"`    // 评论用户
	Article Article   `gorm:"foreignkey:ArticleID"` // 评论文章
	Parent  *Comment  `gorm:"foreignkey:ParentID"`  // 父评论
	Replies []Comment `gorm:"foreignkey:ParentID"`  // 子评论
}
