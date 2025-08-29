package controller

import (
	"gin_boot/internal/router/routers"
	"gin_boot/internal/service"
	"gin_boot/internal/utils/captcha"
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
)

type Captcha struct {
	redisSvc *captcha.RedisStore
	svc      *service.CaptchaService
}

// 注册路由
func (c *Captcha) RegisterRoutes(server *gin.Engine) {
	apiv1 := server.Group(routers.RouterBase.APIV1)
	apiv1.GET("/captcha", c.GetCaptcha)
}

func NewCaptchaController(redisSvc *captcha.RedisStore, svc *service.CaptchaService) *Captcha {
	return &Captcha{
		redisSvc: redisSvc,
		svc:      svc,
	}
}

func (c *Captcha) GetCaptcha(ctx *gin.Context) {
	id, b64s, _ := c.svc.CaptchaMake()
	response.SuccessData(ctx, gin.H{
		"idKey": id,
		"image": b64s,
	})
}
