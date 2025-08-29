package router

import (
	"gin_boot/internal/controller"
	"gin_boot/internal/router/routers"
)

func NewAllHandlers(
	captcha *controller.Captcha,
	userController *controller.UserController,
) []routers.RouteRegistrar {
	return []routers.RouteRegistrar{
		captcha,
		userController,
	}
}
