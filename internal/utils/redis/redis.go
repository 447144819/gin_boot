package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisService 提供 Redis 操作的封装方法
type RedisService struct {
	client *redis.Client
}

// NewRedisService 创建 RedisService 实例
func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		client: client,
	}
}

// Set 设置缓存，默认单位：秒
func (s *RedisService) Set(ctx context.Context, key string, value string, expiration int) error {
	expirationTime := time.Duration(expiration) * time.Second
	return s.client.Set(ctx, key, value, expirationTime).Err()
}

// Get 获取缓存
func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}

// Delete 删除指定的键
func (s *RedisService) Delete(ctx context.Context, key string) error {
	return s.client.Del(ctx, key).Err()
}

// Exists 检查键是否存在
func (s *RedisService) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
