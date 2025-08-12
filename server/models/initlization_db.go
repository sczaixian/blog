package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := "sc:123@tcp(192.168.3.52:3306)/djangoblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(
		&User{},
		&Article{},
		&Category{},
		&Tag{},
		&Comment{},
		&ArticleTag{},
		&Like{},
		&UserFollow{},
		&Notification{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------init ok!------------")
}
