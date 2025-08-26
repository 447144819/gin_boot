package middleware

import (
	"gin_boot/pkg/jwts"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 白名单
var paths = []string{
	"/api/v1/captcha",
	"/api/v1/users/login",
}

type JWTMiddlewareBuilder struct{}

func NewJWTAuthMiddleware() *JWTMiddlewareBuilder {
	return &JWTMiddlewareBuilder{}
}
func (l *JWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 白名单
		for _, path := range paths {
			if path == c.Request.URL.Path {
				return
			}
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证token"})
			return
		}

		// 按照 Bearer <token> 格式解析
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token格式错误"})
			return
		}

		tokenString := parts[1]
		claims, err := jwts.NewJWTHandler().ParseToken(tokenString)
		if err != nil || claims.UserID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的Token", "details": err.Error()})
			return
		}
		// 将用户信息注入到 Gin 的上下文中，后续 handler 可以获取
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
