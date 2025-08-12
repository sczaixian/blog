package models

// 标签

type Tag struct {
	BaseModel
	Name     string    `gorm:"type:varchar(20);not null;comment:标签名" json:"name"`
	Articles []Article `gorm:"many2many:article_tags"` //标签关联的文章
}
