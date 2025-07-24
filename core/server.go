package core

import (
	"fmt"
	"qh-gin-api/global"
	"qh-gin-api/initialize"
	"time"
)

func RunServer() {
	if global.QGA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.QGA_CONFIG.System.Addr)

	fmt.Printf(`
	欢迎使用 qh-gin-api
	接口地址 http://127.0.0.1%s
`, address)
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
