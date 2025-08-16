package initializa

import (
	"context"
	"fmt"
	"gin_boot/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// RedisClient 是一个全局的 Redis 客户端实例
var RedisClient *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis() {
	cfg := config.GetRedis()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), // Redis 服务器地址，例如 "localhost:6379"
		Password: cfg.Password,                             // Redis 密码，如果没有则为空字符串 ""
		DB:       cfg.DB,                                   // 使用的数据库编号，默认 0
	})

	// 使用 Ping 命令测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Panic("无法连接到 Redis: %v", err)
	}

	log.Println("成功连接到 Redis")
}

// GetRedisClient 返回 Redis 客户端实例
func GetRedisClient() *redis.Client {
	if RedisClient == nil {
		log.Fatal("Redis 客户端未初始化，请先调用 InitRedis")
	}
	return RedisClient
}
