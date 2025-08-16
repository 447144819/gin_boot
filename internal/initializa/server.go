package initializa

import (
	"fmt"
	"gin_boot/config"
	"gin_boot/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	// 初始化配置
	if err := config.Init("./config/config.yaml"); err != nil {
		fmt.Printf("配置初始化失败: %v", err)
	}

	// 初始化日志
	InitLogger()

	// 初始化数据接
	InitDB()

	defer Logger.Sync() // 刷新缓冲区

	server := gin.Default()

	// 使用配置中间件
	server.Use(middleware.NewCorsMiddleware().Build(), middleware.RecoveryMiddleware())

	// 设置Gin模式
	if mode := config.GetServer().Mode; mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	return server
}
