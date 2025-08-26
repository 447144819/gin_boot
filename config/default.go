package config

// ServerConfig 服务器配置
type ServerConfig struct {
	Name    string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
	Mode    string `mapstructure:"mode"`
	Host    string `mapstructure:"host"`
	Env     string `mapstructure:"env"`
	Version string `mapstructure:"version"`
}
