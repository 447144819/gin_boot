package controller

import (
	"gin_boot/internal/dto"
	"gin_boot/internal/initializa/validator"
	"gin_boot/internal/service"
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (u *UserController) Create(ctx *gin.Context) {
	var req dto.UserCreateDTO
	// 使用封装的验证器
	if errors := validator.GinBind(ctx, &req); errors != nil {
		//response.SuccessData(ctx, errors)
		response.Error(ctx, errors.Error())
		return
	}

	response.Custom(ctx, 200, "添加成功", req)
}

func (u *UserController) Edit(ctx *gin.Context) {

}

func (u *UserController) Delete(ctx *gin.Context) {

}

func (u *UserController) Detail(ctx *gin.Context) {

}

func (u *UserController) List(ctx *gin.Context) {

}
func (u *UserController) Login(ctx *gin.Context) {
	response.Success(ctx, "登录功能")
}

func (u *UserController) Logout(ctx *gin.Context) {

}
