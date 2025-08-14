package models

type Jwt struct {
	BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
