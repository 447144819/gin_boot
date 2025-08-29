package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
	"unicode"
)

// GetClientIP 获取客户端IP
func GetClientIP(c *gin.Context) string {
	// 优先尝试从 X-Forwarded-For 获取
	if forwarded := c.GetHeader("X-Forwarded-For"); forwarded != "" {
		// 可能有多个 IP，取第一个
		ips := strings.Split(forwarded, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// 其次尝试 X-Real-IP
	if realIP := c.GetHeader("X-Real-IP"); realIP != "" {
		return realIP
	}

	// 最后使用 Gin 提供的 ClientIP() 方法（内部已处理部分代理情况）
	return c.ClientIP()
}

// 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// ToSnake 转蛇形命名,如 GetRole -> get_role
func ToSnake(s string) string {
	if s == "" {
		return ""
	}

	var result []rune
	first := true
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if !first {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(ch))
			first = false
		} else {
			result = append(result, ch)
			first = false
		}
	}

	return string(result)
}
