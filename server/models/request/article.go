package request

import (
	"blog/server/models"
	"blog/server/models/common/request"
)

type CreateArticle struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Excerpt    string `json:"excerpt"`
	CategoryID uint   `json:"category_id"`
}

type UpdateArticle struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Excerpt    string `json:"excerpt"`
	CategoryID uint   `json:"category_id"`
}

type ListArticle struct {
	models.Article
	request.PageInfo
	OrderKey string `json:"order_key"` // 排序
	Desc     string `json:"desc"`      // 排序方式:升序false(默认)|降序true
}
