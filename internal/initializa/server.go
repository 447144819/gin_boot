package initializa

import (
	"fmt"
	"gin_boot/config"
	"gin_boot/internal/initializa/log"
	"gin_boot/internal/initializa/validator"
	"gin_boot/internal/middleware"
	"gin_boot/internal/router"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	// 初始化配置
	if err := config.Init("./config/config.yaml"); err != nil {
		fmt.Printf("配置初始化失败: %v", err)
	}

	// 初始化日志
	log.InitLogger()

	// 初始化数据接
	db := InitDB()

	// 初始化redis
	InitRedis()

	// 初始化 Validator（包含中文翻译和自定义规则）
	validator.Init()

	defer log.Sync() // 刷新缓冲区

	server := gin.Default()

	// 使用配置中间件
	server.Use(
		middleware.NewCorsMiddleware().Build(),
		middleware.RecoveryMiddleware(),
	)

	// 设置Gin模式
	if mode := config.GetServer().Mode; mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 注册路由
	router.InitRouter(server, db)

	return server
}
