package ioc

import (
	"gin_boot/config"
	"gin_boot/internal/utils/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

func InitLogger(cfg *config.Config) (*zap.Logger, error) {
	// 设置日志级别
	var level zapcore.Level
	switch cfg.Log.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// 创建编码器配置
	encoderConfig := getEncoderConfig(cfg.Log.Development)

	// 创建编码器
	var encoder zapcore.Encoder
	if cfg.Log.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 创建输出核心
	var cores []zapcore.Core

	// 文件输出
	if cfg.Log.EnableFile {
		fileWriter := createFileWriter(cfg)
		fileCore := zapcore.NewCore(encoder, zapcore.AddSync(fileWriter), level)
		cores = append(cores, fileCore)
	}

	// 控制台输出
	if cfg.Log.EnableConsole {
		consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
		cores = append(cores, consoleCore)
	}

	// 创建logger
	core := zapcore.NewTee(cores...)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return zapLogger, nil
}

// getEncoderConfig 获取编码器配置
func getEncoderConfig(development bool) zapcore.EncoderConfig {
	var config zapcore.EncoderConfig

	// 问题：如果开启后，输入到文件的日志错误等级会乱码
	//if development {
	//	config = zap.NewDevelopmentEncoderConfig()
	//	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//} else {
	config = zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	//}

	// 自定义时间格式
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	// 自定义字段名
	config.TimeKey = "timestamp"
	config.LevelKey = "level"
	config.NameKey = "logger"
	config.CallerKey = "caller"
	config.MessageKey = "message"
	config.StacktraceKey = "stacktrace"

	return config
}

// createFileWriter 创建文件写入器
func createFileWriter(cfg *config.Config) io.WriteCloser {
	if cfg.File.DailyRotate {
		// 使用自定义按天轮转
		return logs.NewDailyRotateWriter(cfg.File)
	} else {
		// 使用lumberjack按大小轮转
		return logs.NewLumberjackWriter(cfg.File)
	}
}
