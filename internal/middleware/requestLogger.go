package middleware

import (
	"gin_boot/internal/utils/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type RequestLogger struct {
}

func NewRequestLogger() *RequestLogger {
	return &RequestLogger{}
}

func (l *RequestLogger) Build(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		if query != "" {
			path = path + "?" + query
		}

		// 处理请求
		c.Next()

		// 计算耗时
		duration := time.Since(start)

		//fmt.Println("记录日志...............")
		// 记录日志
		logs.Info("request handled",
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.Int("status", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Duration("duration", duration),
			zap.String("referer", c.Request.Referer()),
		)
	}
}
