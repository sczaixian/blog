package model

type Tag struct {
	BaseModel
	Name string  `gorm:"size:50;not null; comment:标签" json:"name"`
	Post []*Post `gorm:"many2many:post_tag" json:"-"`
}
