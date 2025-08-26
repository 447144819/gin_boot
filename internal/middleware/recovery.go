package middleware

import (
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RecoveryMiddleware 捕获panic并返回统一错误格式
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.ErrorWithCode(c, http.StatusInternalServerError, err.(error), "服务器错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
