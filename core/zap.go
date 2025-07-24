package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"qh-gin-api/core/internal"
	"qh-gin-api/global"
	"qh-gin-api/utils"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.QGA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.QGA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.QGA_CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.QGA_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.QGA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
