package model

type Post struct {
	BaseModel
	Title   string `gorm:"size:255; not null; comment:标题" json:"title" validate:"required,min=5,max=255"`
	Content string `gorm:"type:longtext; not null; comment:内容" json:"content" validate:"required,min=20"`
	Status  int    `gorm:"default:0 comment:状态[0草稿;1发布;2私密]" json:"status"`

	AuthorID   uint `gorm:"not null; index comment:作者ID" json:"author_id"`
	CategoryID uint `gorm:"not null; index comment:分类ID" json:"category_id"`

	Category Category `gorm:"foreignkey:CategoryID" json:"category"`
	User     User     `gorm:"foreignkey:UserID" json:"user"`

	Comments []Comment `gorm:"foreignkey:PostID" json:"comments"`
	Tags     []Tag     `gorm:"many2many:post_tag" json:"tags"`
}
