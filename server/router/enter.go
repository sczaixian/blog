package router

import (
	"blog/server/api"
)

var (
	UserApi = api.ApiGroupApp.SystemApiGroup.UserApi // 接口是对外的引入的时候也必须是对外的（UserApi必须大写）
)

type RouterGroup struct {
	UserRouter UserRouter
}

var RouterGroupApp = new(RouterGroup)
