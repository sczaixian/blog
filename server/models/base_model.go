package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT;comment:主键id"`
	CreatedAt time.Time `json:"created_at;comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at;comment:更新时间"`
}
