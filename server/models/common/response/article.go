package response

import (
	"blog/server/models"
	"blog/server/models/common/request"
)

type ArticleQuery struct {
	models.Article
	request.PageInfo
	OrderKey string `json:"order_key"` // 排序字段
	Desc     bool   `json:"desc"`      // 排序方式
}
