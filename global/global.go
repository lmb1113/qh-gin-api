package global

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lmb1113/qh-gin-api/config"
	"github.com/lmb1113/qh-gin-api/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	QGA_DB                  *gorm.DB
	QGA_REDIS               *redis.Client
	QGA_CONFIG              config.Server
	QGA_VP                  *viper.Viper
	QGA_LOG                 *zap.Logger
	QGA_ROUTERS             gin.RoutesInfo
	QGA_ACTIVE_DBNAME       *string
	QGA_Timer               = timer.NewTimerTask()
	QGA_Concurrency_Control = &singleflight.Group{}
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex
)
