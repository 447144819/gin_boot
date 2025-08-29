package routers

import "github.com/gin-gonic/gin"

var RouterBase = &RouteContexts{
	APIV1: "/api/v1",
	APIV2: "/api/v2",
	Root:  "/",
	Admin: "/admin",
}

// RouteContexts RouteContext 包含所有可能的路由组
type RouteContexts struct {
	APIV1 string
	APIV2 string
	Root  string
	Admin string
}

func RegisterRoutes(server *gin.Engine, handlers []RouteRegistrar) {
	// 自动遍历所有 Handler 并注册路由
	for _, registrar := range handlers {
		registrar.RegisterRoutes(server)
	}
}

// RouteRegistrar 每个 Handler 实现这个接口来自行注册路由
type RouteRegistrar interface {
	RegisterRoutes(ctx *gin.Engine)
}
