package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		// * 和 ctx.GetHeader("Origin") 都是通配 但是 设置为*时存在一个问题是不允许XMLHttpRequest携带Cookie
		//ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
		// 允许前端的headers中携带得到值
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Token, x-token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, x-token")
		// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	}
}
