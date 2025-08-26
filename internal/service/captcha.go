package service

import (
	"gin_boot/config"
	"gin_boot/internal/utils/captcha"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaService struct {
	store *captcha.RedisStore
	cfg   *config.Config
}

func NewCaptchaService(store *captcha.RedisStore, cfg *config.Config) *CaptchaService {
	return &CaptchaService{
		store: store,
		cfg:   cfg,
	}
}

// 生成验证码
func (c *CaptchaService) CaptchaMake() (id, b64s, code string) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          c.cfg.Captcha.Height,
		Width:           c.cfg.Captcha.Width,
		NoiseCount:      c.cfg.Captcha.NoiseCount,
		ShowLineOptions: c.cfg.Captcha.ShowLineOptions,
		Length:          c.cfg.Captcha.Length,
		Source:          c.cfg.Captcha.Source,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, c.store)
	lid, lb64s, code, _ := captcha.Generate()
	//fmt.Println("id:", id, "code:", code)
	return lid, lb64s, code
}

// 验证captcha是否正确
func (c *CaptchaService) CaptVerify(id string, capt string) bool {
	if c.store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
