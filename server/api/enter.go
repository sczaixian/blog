package api

import (
	sy "blog/server/api/system"
)

// 收集本 包里面的接口

// 对外提供统一调用
var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup sy.ApiGroup
}
