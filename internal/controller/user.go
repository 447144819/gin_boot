package controller

import (
	"gin_boot/internal/service"
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
)

type userController struct {
	svc *service.UserService
}

func NewUserController(svc *service.UserService) *userController {
	return &userController{
		svc: svc,
	}
}

func (uc *userController) Login(ctx *gin.Context) {
	response.Success(ctx, "登录功能")
}
