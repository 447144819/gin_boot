// 简化的日志调用
package log

import (
	"gin_boot/internal/initializa"
	"go.uber.org/zap"
)

func Info(msg string, fields ...zap.Field) {
	initializa.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	initializa.Error(msg, fields...)
}
func Debug(msg string, fields ...zap.Field) {
	initializa.Debug("error", fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	initializa.Fatal(msg, fields...)
}
func Panic(msg string, fields ...zap.Field) {
	initializa.Panic(msg, fields...)
}
