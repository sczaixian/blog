package models

// 文章

type Article struct {
	BaseModel
	Title        string `json:"title" gorm:"not null; type:varchar(255); comment:标题"`
	Content      string `json:"content" gorm:"type:text;comment:内容"`
	Excerpt      string `json:"excerpt" gorm:"type:text;comment:摘要"`
	CoverImage   string `json:"cover_image" gorm:"default:xxx;comment:封面"`
	Status       int    `gorm:"default:0;comment:状态（0:草稿;1:已发布;2:私密）" json:"status"`
	ViewCount    int    `json:"view_count" gorm:"default:0;comment:浏览数"`
	LikeCount    int    `json:"like_count" gorm:"default:0;comment:点赞数"`
	CommentCount int    `json:"comment_count" gorm:"default:0;comment:评论数"`

	UserID     uint `json:"user_id" gorm:"index"`     // 作者id
	CategoryID uint `json:"category_id" gorm:"index"` // 分类id

	User     User      `json:"user" gorm:"foreignkey:UserId"`         // 作者
	Category Category  `json:"category" gorm:"foreignkey:CategoryID"` // 分类
	Tags     []Tag     `gorm:"many2many:article_tags;"`               // 标签
	Comments []Comment `gorm:"foreignKey:ArticleID"`                  // 评论
	Likes    []Like    `gorm:"foreignKey:ArticleID"`                  // 点赞
}
