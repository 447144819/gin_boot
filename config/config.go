package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config 全局配置结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
	File     FileConfig     `mapstructure:"file"` // 文件配置
	Captcha  CaptchaConfig  `mapstructure:"captcha"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")    // 配置文件名称 (不带扩展名)
	viper.SetConfigType("yaml")      // 或者 json
	viper.AddConfigPath("./config/") // 当前目录

	// 你可以根据需要添加更多路径，如 viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
		return nil
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("配置文件解析失败: %v", err)
		return nil
	}

	log.Printf("🔧 加载配置: %+v", cfg)
	return &cfg
}
