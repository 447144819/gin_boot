package config

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
