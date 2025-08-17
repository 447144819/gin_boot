package config

// ServerConfig 服务器配置
type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
}

//// GinLogConfig Gin日志中间件配置
//type GinLogConfig struct {
//	Enable       bool     `yaml:"enable"`        // 是否启用Gin日志中间件
//	SkipPaths    []string `yaml:"skip_paths"`    // 跳过记录的路径
//	TimeFormat   string   `yaml:"time_format"`   // 时间格式
//	CustomFormat bool     `yaml:"custom_format"` // 是否使用自定义格式
//}
