package initializa

import (
	"fmt"
	"gin_boot/config"
	"gin_boot/internal/controller/tests"
	"gin_boot/internal/initializa/log"
	"gin_boot/internal/initializa/validator"
	"gin_boot/internal/middleware"
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
	InitDB()

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

	// 测试专用
	test := server.Group("/api/v1/test/")
	{
		test.POST("addUser", tests.NewExampleController().CreateUser)
		test.Use(middleware.JWTAuthMiddleware())
		{
			test.GET("list", func(c *gin.Context) {
				userID, _ := c.Get("userID")
				username, _ := c.Get("username")
				c.JSON(200, gin.H{
					"user":     userID,
					"username": username,
				})
			})
		}
	}

	return server
}
