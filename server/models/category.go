package models

import (
	"runtime/metrics"

	"gorm.io/gorm"
)

// 分类

type Category struct {
	BaseModel
	Name        string `json:"name" gorm:"index; comment:分类"`
	Description string `json:"description" gorm:"size:255;comment:描述"`
	Sort        int    `json:"sort" gorm:"default:-1"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"size:50;unique;not null"` // 分类名称
	Description string `gorm:"size:255"`                // 分类描述
	Slug        string `gorm:"size:50;unique"`          // SEO友好URL
	Sort        int    `gorm:"default:0"`               // 排序

	// Relationships
	Articles []Article `gorm:"foreignKey:CategoryID"` // 分类下的文章
}
