package config

// Config 全局配置结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
	File     FileConfig     `yaml:"file"` // 文件配置
	//Gin      GinLogConfig   `yaml:"gin"`  // Gin中间件配置
}

// 全局便捷函数
var globalManager = GetConfigManager()

// Init 全局初始化函数
func Init(configPath string) error {
	return globalManager.InitConfig(configPath)
}

// Get 获取完整配置
func Get() Config {
	return globalManager.GetConfig()
}

// GetServer 获取服务器配置
func GetServer() ServerConfig {
	return globalManager.GetServerConfig()
}

// GetDatabase 获取数据库配置
func GetDatabase() DatabaseConfig {
	return globalManager.GetDatabaseConfig()
}

// GetRedis 获取Redis配置
func GetRedis() RedisConfig {
	return globalManager.GetRedisConfig()
}

// GetLog 获取日志配置
func GetLog() LogConfig {
	return globalManager.GetLogConfig()
}

// GetLog 获取日志配置
func GetFile() FileConfig {
	return globalManager.GetFileConfig()
}
