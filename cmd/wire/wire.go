//go:build wireinject

package wire

import (
	"gin_boot/config"
	"gin_boot/internal/ioc"
	"gin_boot/internal/router"
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
		ioc.DaoSet,

		// 引入 Service 层的 ProviderSet
		ioc.ServiceSet,

		// 引入 Controller 层的 ProviderSet
		ioc.ControllerSet,

		// 初始化
		ioc.InitWebServer,

		// 自动注册
		router.NewAllHandlers,
	)
	return &ioc.Server{}, nil
}
