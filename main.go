package main

import (
	"fmt"
	"gin_boot/config"
	"gin_boot/internal/initializa"
	"gin_boot/internal/initializa/log"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化
	server := initializa.InitServer()

	server.GET("/", func(ctx *gin.Context) {
	})

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", config.GetServer().Host, config.GetServer().Port)
	log.Info("🚀 服务器启动成功，监听地址: " + addr)
	log.Info("📝 当前运行模式: " + config.GetServer().Mode)
	err := server.Run(addr)
	if err != nil {
		log.Error("启动服务失败")
	}
}
