package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"qh-gin-api/config"
	"qh-gin-api/utils/timer"
	"sync"
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
	lock                    sync.RWMutex
)
