package main

import (
	"gin_boot/cmd/wire"
	"gin_boot/internal/utils/logs"
)

func main() {
	// 初始化
	server, err := wire.InitWebServer()

	// 在程序退出前，调用 log.Sync()，确保日志全部刷新
	logs.Sync()

	// 启动服务器
	err = server.Run()
	if err != nil {
		logs.Error("启动服务失败")
	}
}
