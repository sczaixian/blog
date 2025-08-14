package system

import (
	"blog/server/service"
)

// 搜集包内对外接口
type ApiGroup struct {
	UserApi
	ArticleApi
	CategoryApi
}

// 注册外部服务，在包内可以用，统一管理
var (
	userService     = service.ServiceGroupApp.ServiceGroup.UserService
	articleService  = service.ServiceGroupApp.ServiceGroup.ArticleService
	categoryService = service.ServiceGroupApp.ServiceGroup.CategoryService
)
