package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
