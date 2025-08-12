package models

import "time"

type ArticleTag struct {
	TagID     uint      `gorm:"primary_key"` // 标签id
	ArticleID uint      `gorm:"primary_key"` // 文章id
	CreatedAt time.Time `json:"created_at"`  // 创建时间
}
