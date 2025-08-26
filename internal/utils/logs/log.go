package logs

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

// Init 初始化日志实例
func Init(l *zap.Logger) {
	logger = l
}

// Sync 同步日志缓冲区
func Sync() {
	if logger != nil {
		logger.Sync()
	}
}

// 便捷方法
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Dump(args ...interface{}) {
	fmt.Println(args...)
}
