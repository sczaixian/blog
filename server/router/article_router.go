package router

import "github.com/gin-gonic/gin"

type ArticleRouter struct{}

func (a *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article")
	{
		articleRouter.POST("articles", ArticleApi.CreateArticle)
		articleRouter.GET("articles", ArticleApi.ListArticle)
		articleRouter.GET("article/:id", ArticleApi.GetArticle)
		articleRouter.DELETE("article/:id", ArticleApi.DeleteArticle)
		articleRouter.POST("article/:id", ArticleApi.EditArticle)
		articleRouter.PUT("article/:id", ArticleApi.SaveArticle)
	}
}
