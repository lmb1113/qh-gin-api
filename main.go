package main

import (
	"github.com/lmb1113/qh-gin-api/core"
	"github.com/lmb1113/qh-gin-api/global"
	"github.com/lmb1113/qh-gin-api/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       qh-gin-api Swagger API接口文档
// @version                     v1.0.0
// @description                 清欢-开箱即用简易后端框架
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /

func main() {
	// 初始化配置
	global.QGA_VP = core.Viper()      // 初始化Viper
	global.QGA_LOG = core.Zap()       // 初始化zap日志库
	global.QGA_DB = initialize.Gorm() // gorm连接数据库
	zap.ReplaceGlobals(global.QGA_LOG)
	initialize.OtherInit()
	if global.QGA_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
	core.RunServer()
}
