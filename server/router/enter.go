package router

import (
	"blog/server/api"
)

var (
	UserApi    = api.ApiGroupApp.SystemApiGroup.UserApi // 接口是对外的引入的时候也必须是对外的（UserApi必须大写）
	ArticleApi = api.ApiGroupApp.SystemApiGroup.ArticleApi
)

type RouterGroup struct {
	UserRouter    UserRouter
	ArticleRouter ArticleRouter
}

var RouterGroupApp = new(RouterGroup)
