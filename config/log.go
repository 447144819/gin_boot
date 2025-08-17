package config

// LogConfig 日志配置
type LogConfig struct {
	Level         string `mapstructure:"level"`
	Encoding      string `mapstructure:"encoding"`
	Development   bool   `mapstructure:"development"`
	EnableFile    bool   `mapstructure:"enable_file"`
	EnableConsole bool   `mapstructure:"enable_console"`
}
