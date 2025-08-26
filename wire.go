//go:build wireinject

package main

import (
	"gin_boot/config"
	"gin_boot/internal/controller"
	"gin_boot/internal/dao"
	"gin_boot/internal/ioc"
	"gin_boot/internal/router"
	"gin_boot/internal/service"
	"gin_boot/internal/utils/captcha"
	"gin_boot/internal/utils/redis"
	"github.com/google/wire"
)

func InitWebServer() (*ioc.Server, error) {
	wire.Build(
		config.LoadConfig,
		ioc.InitLogger,
		ioc.InitDB, ioc.InitRedis,

		redis.NewRedisService,

		captcha.NewRedisStore,

		// 引入 DAO 层的 ProviderSet
		dao.DaoSet,

		// 引入 Service 层的 ProviderSet
		service.ServiceSet,

		// 引入 Controller 层的 ProviderSet
		controller.ControllerSet,

		// 初始化
		ioc.InitWebServer,

		// 自动注册
		router.NewAllHandlers,
	)
	return &ioc.Server{}, nil
}
