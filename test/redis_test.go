package main

import (
	"context"
	"fmt"
	"gin_boot/config"
	"gin_boot/internal/initializa"
	"gin_boot/internal/utils/redis"
	"testing"
)

func init() {
	// 初始化配置
	if err := config.Init("../config/config.yaml"); err != nil {
		fmt.Printf("配置初始化失败: %v", err)
	}

	// 初始化redis
	initializa.InitRedis()
}

var (
	ctx     = context.Background()
	key     = "name"
	val     = "lzw123"
	expTime = 60 // 过期时间，单位秒
)

func TestSet(t *testing.T) {
	err := redis.NewRedisService().Set(ctx, key, val, expTime)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("设置缓存成功")
}

func TestGet(t *testing.T) {
	name, _ := redis.NewRedisService().Get(ctx, key)
	t.Log(name)
}

func TestExists(t *testing.T) {
	ok, _ := redis.NewRedisService().Exists(ctx, key)
	t.Log(ok)
}
