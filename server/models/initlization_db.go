package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "sc:123@tcp(192.168.3.52:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:123@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
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
		&Jwt{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------init ok!------------")
	return db
}
