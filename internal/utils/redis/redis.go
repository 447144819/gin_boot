package redis

import (
	"context"
	"errors"
	"gin_boot/internal/initializa"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

// ExampleController 定义示例控制器
type ExampleController struct{}

// NewExampleController 创建 ExampleController 实例
func NewExampleController() *ExampleController {
	return &ExampleController{}
}

// Set 设置
func (ec *ExampleController) Set(ctx *gin.Context, key string, value string, expire time.Duration) {
	if key == "" || value == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Key 和 Value 参数必填"})
		return
	}

	rdb := initializa.GetRedisClient()
	ctxRedis := context.Background()

	err := rdb.Set(ctxRedis, key, value, expire).Err() // 设置键值对，过期时间为 10 分钟
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法设置 Redis 键值对: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "成功设置键值对", "key": key, "value": value})
}

// Get 获取
func (ec *ExampleController) Get(ctx *gin.Context, key string) {
	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Key 参数必填"})
		return
	}

	rdb := initializa.GetRedisClient()
	ctxRedis := context.Background()

	val, err := rdb.Get(ctxRedis, key).Result()
	if errors.Is(err, redis.Nil) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "键不存在"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取 Redis 键值: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"key": key, "value": val})
}
