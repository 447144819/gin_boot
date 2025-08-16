// 简化的日志调用
package log

import (
	"gin_boot/internal/initializa/log"
	"go.uber.org/zap"
)

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
func Debug(msg string, fields ...zap.Field) {
	log.Debug("error", fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	log.Panic(msg, fields...)
}
