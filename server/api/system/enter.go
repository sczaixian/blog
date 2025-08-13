package system

import (
	"blog/server/service"
)

// 搜集包内对外接口
type ApiGroup struct {
	UserApi
}

// 注册外部服务，在包内可以用，统一管理
var (
	userService = service.ServiceGroupApp.ServiceGroup.UserService
)
