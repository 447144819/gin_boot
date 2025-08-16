package log

import (
	"gin_boot/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

// InitLogger 初始化Logger
func InitLogger() {
	cfg := config.GetLog()

	// 设置日志级别
	level := getZapLevel(cfg.Level)

	// 创建编码器配置
	encoderConfig := getEncoderConfig(cfg.Development)

	// 创建编码器
	var encoder zapcore.Encoder
	if cfg.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 创建输出核心
	var cores []zapcore.Core

	// 控制台输出
	if cfg.EnableConsole {
		consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
		cores = append(cores, consoleCore)
	}

	// 文件输出
	if cfg.EnableFile {
		fileWriter := createFileWriter()
		fileCore := zapcore.NewCore(encoder, zapcore.AddSync(fileWriter), level)
		cores = append(cores, fileCore)
	}

	// 创建logger
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// 全局赋值
	Logger = logger
	Sugar = logger.Sugar()

	// 替换全局logger
	zap.ReplaceGlobals(logger)

	Logger.Info("日志系统初始化成功",
		zap.String("level", cfg.Level),
		zap.String("encoding", cfg.Encoding),
		zap.Bool("console", cfg.EnableConsole),
		zap.Bool("file", cfg.EnableFile),
	)
}

// createFileWriter 创建文件写入器
func createFileWriter() io.WriteCloser {
	cfg := config.GetFile()
	if cfg.DailyRotate {
		// 使用自定义按天轮转
		return NewDailyRotateWriter(cfg)
	} else {
		// 使用lumberjack按大小轮转
		return NewLumberjackWriter(cfg)
	}
}

// getZapLevel 获取zap日志级别
func getZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// getEncoderConfig 获取编码器配置
func getEncoderConfig(development bool) zapcore.EncoderConfig {
	var config zapcore.EncoderConfig

	if development {
		config = zap.NewDevelopmentEncoderConfig()
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionEncoderConfig()
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	}

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

// Sync 同步日志缓冲区
func Sync() {
	if Logger != nil {
		Logger.Sync()
	}
}

// 便捷方法
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

// Sugar便捷方法
func Debugf(template string, args ...interface{}) {
	Sugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	Sugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	Sugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	Sugar.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Sugar.Fatalf(template, args...)
}

// WithFields 添加字段
func WithFields(fields ...zap.Field) *zap.Logger {
	return Logger.With(fields...)
}
