package config

// 验证码
type CaptchaConfig struct {
	Height          int    `yaml:"height"`
	Width           int    `yaml:"width"`
	NoiseCount      int    `yaml:"noise_count"`
	ShowLineOptions int    `yaml:"show_line_options"`
	Length          int    `yaml:"length"`
	Source          string `yaml:"source"`
}
