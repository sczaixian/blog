package model

type Category struct {
	BaseModel
	Name        string `gorm:"type:varchar(255); not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Posts       []Post `gorm:"foreignkey:CategoryID" json:"posts"`
}
