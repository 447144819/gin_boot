package initializa

import (
	"fmt"
	"gin_boot/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

// InitLogger 初始化Logger
func InitLogger() {
	cfg := config.GetLog()

	fmt.Println(config.GetLog())
	// 设置日志级别
	logLevel := zapcore.InfoLevel
	if cfg.Debug {
		logLevel = zapcore.DebugLevel
	}

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志输出
	var core zapcore.Core
	if cfg.EnableFile {
		// 文件输出
		fileWriter := getLogWriter(cfg.FilePath)
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			fileWriter,
			logLevel,
		)

		if cfg.EnableConsole {
			// 控制台和文件都输出
			consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
			consoleCore := zapcore.NewCore(
				consoleEncoder,
				zapcore.Lock(os.Stdout),
				logLevel,
			)
			core = zapcore.NewTee(consoleCore, fileCore)
		} else {
			// 仅文件输出
			core = fileCore
		}
	} else {
		// 仅控制台输出
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			logLevel,
		)
	}

	// 构造日志
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// getLogWriter 获取日志写入器
func getLogWriter(filePath string) zapcore.WriteSyncer {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

// 封装常用方法
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}
