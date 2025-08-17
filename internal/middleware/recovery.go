package middleware

import (
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获panic并返回统一错误格式
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.ErrorWithCode(c, response.ServerErrorCode, "服务器错误"+err.(string))
				c.Abort()
			}
		}()
		c.Next()
	}
}
