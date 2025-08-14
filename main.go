package main

import (
	"blog/server/config"
	"blog/server/core"
	"blog/server/global"
	"blog/server/middleware"
	"blog/server/models"
	rter "blog/server/router"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func initRedisClient(rediscfg config.Redis) (redis.UniversalClient, error) {
	var client redis.UniversalClient
	client = redis.NewClient(&redis.Options{
		Addr:     rediscfg.Addr,
		Password: rediscfg.Password,
		DB:       rediscfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis ping error", zap.String("name", rediscfg.Name), zap.Error(err))
		return nil, err
	}
	global.GVA_LOG.Info("redis ping response", zap.String("name", rediscfg.Name), zap.String("pong", pong))
	return client, nil
}

func Redis() {
	redisClient, err := initRedisClient(global.GVA_CONFIG.Redis)
	if err != nil {
		panic(err)
	}
	global.GVA_REDIS = redisClient
}

func initRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}
	router.Use(middleware.JwtMiddleware())
	sysRouter := rter.RouterGroupApp.UserRouter
	articleRouter := rter.RouterGroupApp.ArticleRouter
	categoryRouter := rter.RouterGroupApp.CategoryRouter
	sysRouter.InitUserRouter(router.Group("/"))
	articleRouter.InitArticleRouter(router.Group("/"))
	categoryRouter.InitCategoryRouter(router.Group("/"))
	return router
}

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = models.InitDB()
	Redis()
	router := initRouter()
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
