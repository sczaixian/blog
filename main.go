package main

import (
	"blog/server/core"
	"blog/server/global"
	"blog/server/models"
	"fmt"

	rter "blog/server/router"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}
	sysRouter := rter.RouterGroupApp.UserRouter
	articRouter := rter.RouterGroupApp.ArticleRouter
	sysRouter.InitUserRouter(router.Group("/"))
	articRouter.InitArticleRouter(router.Group("/"))
	return router
}

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = models.InitDB()
	user := &models.User{}
	global.GVA_DB.Model(models.User{}).First(user)
	fmt.Println(user)
	router := initRouter()
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
