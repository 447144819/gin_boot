package redis

import (
	"context"
	"gin_boot/internal/initializa"
	"time"
)

// RedisService 提供 Redis 操作的封装方法
type RedisService struct{}

// NewRedisService 创建 RedisService 实例
func NewRedisService() *RedisService {
	return &RedisService{}
}

// Set 设置缓存，默认单位：秒
func (s *RedisService) Set(ctx context.Context, key string, value string, expiration int) error {
	rdb := initializa.GetRedisClient()
	expirationTime := time.Duration(expiration) * time.Second
	return rdb.Set(ctx, key, value, expirationTime).Err()
}

// Get 获取缓存
func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
	rdb := initializa.GetRedisClient()
	return rdb.Get(ctx, key).Result()
}

// Delete 删除指定的键
func (s *RedisService) Delete(ctx context.Context, key string) error {
	rdb := initializa.GetRedisClient()
	return rdb.Del(ctx, key).Err()
}

// Exists 检查键是否存在
func (s *RedisService) Exists(ctx context.Context, key string) (bool, error) {
	rdb := initializa.GetRedisClient()
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
