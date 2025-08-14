package router

import "github.com/gin-gonic/gin"

type CategoryRouter struct{}

func (r *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category")
	{
		categoryRouter.POST("add", CategoryApi.CreateCategory)
	}

}
