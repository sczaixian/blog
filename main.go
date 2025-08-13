package main

import (
	"blog/server/global"
	"blog/server/models"
	"fmt"

	rter "blog/server/router"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	//models.InitDB()
	//dsn := "sc:123@tcp(192.168.3.52:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db)
	fmt.Println("hello world")
	global.GVA_DB = db
	user := &models.User{}
	global.GVA_DB.Model(models.User{}).First(user)
	fmt.Println(user)
	router := initRouter()
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
