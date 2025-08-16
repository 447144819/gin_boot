package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"sync"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level         string `mapstructure:"level"`
	Encoding      string `mapstructure:"encoding"`
	Development   bool   `mapstructure:"development"`
	EnableFile    bool   `mapstructure:"enable_file"`
	EnableConsole bool   `mapstructure:"enable_console"`
}

// FileConfig 文件日志配置
type FileConfig struct {
	// 基础文件配置
	Dir      string `mapstructure:"dir"`      // 日志目录
	Filename string `mapstructure:"filename"` // 日志文件名（不含扩展名）

	// 轮转配置
	MaxSize    int  `mapstructure:"max_size"`    // 单个文件最大MB
	MaxAge     int  `mapstructure:"max_age"`     // 保留天数
	MaxBackups int  `mapstructure:"max_backups"` // 最大备份文件数
	Compress   bool `mapstructure:"compress"`    // 是否压缩
	LocalTime  bool `mapstructure:"local_time"`  // 是否使用本地时间

	// 按天分割配置
	DailyRotate bool   `mapstructure:"daily_rotate"` // 启用按天轮转
	TimeFormat  string `mapstructure:"time_format"`  // 时间格式
}

//// GinLogConfig Gin日志中间件配置
//type GinLogConfig struct {
//	Enable       bool     `yaml:"enable"`        // 是否启用Gin日志中间件
//	SkipPaths    []string `yaml:"skip_paths"`    // 跳过记录的路径
//	TimeFormat   string   `yaml:"time_format"`   // 时间格式
//	CustomFormat bool     `yaml:"custom_format"` // 是否使用自定义格式
//}

// Config 全局配置结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
	File     FileConfig     `yaml:"file"` // 文件配置
	//Gin      GinLogConfig   `yaml:"gin"`  // Gin中间件配置
}

// ConfigManager 配置管理器
type ConfigManager struct {
	config *Config
	mutex  sync.RWMutex
	viper  *viper.Viper
}

var (
	manager *ConfigManager
	once    sync.Once
)

// GetConfigManager 获取配置管理器实例（单例模式）
func GetConfigManager() *ConfigManager {
	once.Do(func() {
		manager = &ConfigManager{
			config: &Config{},
			viper:  viper.New(),
		}
	})
	return manager
}

// InitConfig 初始化配置
func (cm *ConfigManager) InitConfig(configPath string) error {
	// 设置配置文件路径和名称
	cm.viper.SetConfigFile(configPath)
	cm.viper.SetConfigType("yaml")

	// 设置环境变量前缀
	cm.viper.SetEnvPrefix("APP")
	cm.viper.AutomaticEnv()

	// 读取配置文件
	if err := cm.viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	if err := cm.loadConfig(); err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}

	// 启动热加载监听
	cm.watchConfig()

	log.Printf("✅ 配置初始化完成，监听文件: %s", cm.viper.ConfigFileUsed())
	return nil
}

// loadConfig 加载配置到结构体
func (cm *ConfigManager) loadConfig() error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	return cm.viper.Unmarshal(cm.config)
}

// watchConfig 监听配置文件变化
func (cm *ConfigManager) watchConfig() {
	cm.viper.WatchConfig()
	cm.viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("🔄 检测到配置文件变化: %s", e.Name)

		if err := cm.loadConfig(); err != nil {
			log.Printf("❌ 配置热加载失败: %v", err)
		} else {
			log.Printf("✅ 配置热加载成功")
			cm.printCurrentConfig()
		}
	})
}

// GetConfig 获取完整配置（线程安全）
func (cm *ConfigManager) GetConfig() Config {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return *cm.config
}

// GetServerConfig 获取服务器配置
func (cm *ConfigManager) GetServerConfig() ServerConfig {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.config.Server
}

// GetDatabaseConfig 获取数据库配置
func (cm *ConfigManager) GetDatabaseConfig() DatabaseConfig {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.config.Database
}

// GetRedisConfig 获取Redis配置
func (cm *ConfigManager) GetRedisConfig() RedisConfig {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.config.Redis
}

// GetLogConfig 获取日志配置
func (cm *ConfigManager) GetLogConfig() LogConfig {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.config.Log
}

// GetFileConfig 获取文件配置
func (cm *ConfigManager) GetFileConfig() FileConfig {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.config.File
}

// GetString 获取字符串配置值
func (cm *ConfigManager) GetString(key string) string {
	return cm.viper.GetString(key)
}

// GetInt 获取整数配置值
func (cm *ConfigManager) GetInt(key string) int {
	return cm.viper.GetInt(key)
}

// GetBool 获取布尔配置值
func (cm *ConfigManager) GetBool(key string) bool {
	return cm.viper.GetBool(key)
}

// printCurrentConfig 打印当前配置（调试用）
func (cm *ConfigManager) printCurrentConfig() {
	config := cm.GetConfig()
	log.Printf("当前配置: Server=%+v, Database=%+v", config.Server, config.Database)
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
