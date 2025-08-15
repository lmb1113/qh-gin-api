package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/lmb1113/qh-gin-api/docs"
	"github.com/lmb1113/qh-gin-api/global"
	"github.com/lmb1113/qh-gin-api/middleware"
	"github.com/lmb1113/qh-gin-api/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func HealthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

func Routers() *gin.Engine {

	Router := gin.Default()
	//gin.SetMode(gin.DebugMode)

	docs.SwaggerInfo.BasePath = global.QGA_CONFIG.System.RouterPrefix
	Router.GET(global.QGA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.QGA_LOG.Info("register swagger handler")
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", HealthCheck)
	}

	Router.Use(middleware.WithRequestId())
	Router.Use(middleware.Recovery())
	Router.Use(gin.Logger())
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求

	userRouter := router.RouterGroupApp.User
	PrivateGroup := Router.Group("")

	PrivateGroup.Use(middleware.JWTAuth())
	userRouter.InitUsersRouter(PrivateGroup)
	global.QGA_ROUTERS = Router.Routes()
	global.QGA_LOG.Info("router register success")
	return Router
}
