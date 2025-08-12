package global

import (
	"sync"

	"blog/server/config"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB      *gorm.DB
	GVA_LIST    map[string]*gorm.DB
	GVA_REDIS   redis.UniversalClient
	GVA_MONGO   *qmgo.QmgoClient
	GVA_LOG     *zap.Logger
	GVA_ROUTERS gin.RoutesInfo
	lock        sync.RWMutex
	BlackCache  local_cache.Cache
	GVA_VP      *viper.Viper
	GVA_CONFIG  config.Server
)
