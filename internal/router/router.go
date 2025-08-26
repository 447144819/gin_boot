package router

import (
	"gin_boot/internal/controller"
	"gin_boot/internal/controller/common"
	"github.com/gin-gonic/gin"
)

//type AllHandlers struct {
//	Registrars []common.RouteRegistrar // 存储所有实现了 RouteRegistrar 的 Handler
//}

func NewAllHandlers(
	userHandler *controller.UserController,
	captchaHandler *controller.Captcha,
) []common.RouteRegistrar {
	return []common.RouteRegistrar{
		userHandler,
		captchaHandler,
		// 新增的 Handler 直接加入这个切片
	}
}
func RegisterRoutes(server *gin.Engine, handlers []common.RouteRegistrar) {
	// 定义路由组
	ctx := &common.RouteContext{
		APIV1: server.Group("/api/v1"),
		APIV2: server.Group("/api/v2"),
		Root:  server.Group("/"),
		Admin: server.Group("/admin"),
	}

	// 自动遍历所有 Handler 并注册路由
	for _, registrar := range handlers {
		registrar.RegisterRoutes(ctx)
	}
}
