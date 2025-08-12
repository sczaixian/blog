package models

// 分类

type Category struct {
	BaseModel
	Name        string    `json:"name" gorm:"index; comment:分类"`
	Description string    `json:"description" gorm:"size:255;comment:描述"`
	Sort        int       `json:"sort" gorm:"default:0;comment:排序"`
	Articles    []Article `json:"articles" gorm:"foreignkey:CategoryID"` // 分类下文章
}
