package router

import "github.com/gin-gonic/gin"

type ArticleRouter struct{}

func (a *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article")
	{
		articleRouter.POST("create", ArticleApi.CreateArticle)
		articleRouter.POST("list", ArticleApi.ListArticle)
		articleRouter.GET(":id", ArticleApi.GetArticle)
		articleRouter.DELETE(":id", ArticleApi.DeleteArticle)
		articleRouter.POST(":id", ArticleApi.EditArticle)
	}
}
