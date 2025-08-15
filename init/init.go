package init

import "test/config"

func Init() {
	InitMySQLByConfig(config.Mysql{})
}
