package init

import (
	"test/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQLByConfig(conf config.Mysql) *gorm.DB {
	if db, err := gorm.Open(mysql.Open(conf.Dsn()), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		return db
	}
}
