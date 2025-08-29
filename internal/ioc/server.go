package ioc

import (
	"fmt"
	"gin_boot/config"
	"gin_boot/internal/middleware"
	"gin_boot/internal/router/routers"
	"gin_boot/internal/utils/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	cfg    *config.Config
	log    *zap.Logger
	engine *gin.Engine
	//handlers []common.RouteRegistrar // 只依赖一个 Handler 集合
}

func InitWebServer(cfg *config.Config, handlers []routers.RouteRegistrar, logger *zap.Logger) (*Server, error) {
	s := &Server{
		cfg: cfg,
		log: logger,
		//db: db,
		//handlers: handlers,
	}
	server := gin.Default()

	// 初始化日志
	logs.Init(logger)

	// 中间件
	s.InitGinMiddlewares(server)

	// 注册路由
	routers.RegisterRoutes(server, handlers)

	//设置Gin模式
	if mode := cfg.Server.Mode; mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	s.engine = server
	return s, nil
}

func (s *Server) Run() error {
	// 启动服务器
	addr := fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port)
	//log.Info("🚀 服务器启动成功，监听地址: " + addr)
	//log.Info("📝 当前运行模式: " + s.cfg.Server.Mode)
	return s.engine.Run(addr)
}

func (s *Server) InitGinMiddlewares(server *gin.Engine) {
	server.Use(
		middleware.NewCorsMiddleware().Build(),
		middleware.RecoveryMiddleware(),
		middleware.NewJWTAuthMiddleware().Build(),
		middleware.NewRequestLogger().Build(s.log),
	)
}
