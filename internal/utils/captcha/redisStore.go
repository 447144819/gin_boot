package captcha

import (
	"context"
	"gin_boot/internal/utils/redis"
)

var (
	captchaPrefix  = "captcha:"
	ctx            = context.Background()
	expirationTime = 60 // 过期时间，单位秒
)

// RedisStore 实现 base64Captcha.Store 的三个方法
type RedisStore struct {
	redisSvc *redis.RedisService
}

func NewRedisStore(redisSvc *redis.RedisService) *RedisStore {
	return &RedisStore{redisSvc: redisSvc}
}

func (r *RedisStore) Set(id string, value string) error {
	key := captchaPrefix + id
	return r.redisSvc.Set(ctx, key, value, expirationTime)
}

func (r *RedisStore) Get(id string, clear bool) string {
	key := captchaPrefix + id
	val, err := r.redisSvc.Get(ctx, key)
	if err != nil {
		return ""
	}
	if clear {
		_ = r.redisSvc.Delete(ctx, key)
	}
	return val
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
