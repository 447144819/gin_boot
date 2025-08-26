package common

import "github.com/gin-gonic/gin"

// RouteContext 包含所有可能的路由组
type RouteContext struct {
	APIV1 *gin.RouterGroup
	APIV2 *gin.RouterGroup
	Root  *gin.RouterGroup
	Admin *gin.RouterGroup // 可以继续扩展
}

// RouteRegistrar 每个 Handler 实现这个接口来自行注册路由
type RouteRegistrar interface {
	RegisterRoutes(ctx *RouteContext)
}
