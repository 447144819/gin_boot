package controller

import (
	"gin_boot/internal/dto"
	"gin_boot/internal/initializa/validator"
	"gin_boot/internal/service"
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
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
		response.Error(ctx, errors.Error())
		return
	}

	err := u.svc.Create(ctx, req)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, "添加成功")
}

func (u *UserController) Edit(ctx *gin.Context) {
	var req dto.UserEditDTO
	// 使用封装的验证器
	if errors := validator.GinBind(ctx, &req); errors != nil {
		response.Error(ctx, errors.Error())
		return
	}

	err := u.svc.Edit(ctx, req)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, "修改成功")
}

func (u *UserController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id < 1 {
		response.Error(ctx, "参数错误")
	}
	err := u.svc.Delete(ctx, id)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.Success(ctx, "删除成功")
}

func (u *UserController) Detail(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if id < 1 {
		response.Error(ctx, "参数错误")
	}
	user, err := u.svc.Detial(ctx, id)
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.SuccessData(ctx, user)
}

func (u *UserController) List(ctx *gin.Context) {
	var req dto.UserListDTO
	// 使用封装的验证器
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, err.Error())
		return
	}

	data, err := u.svc.List(ctx, &req)
	if err != nil {
		response.Error(ctx, err.Error())
	}
	response.SuccessData(ctx, data)
}

func (u *UserController) Login(ctx *gin.Context) {
	response.Success(ctx, "登录功能")
}

func (u *UserController) Logout(ctx *gin.Context) {
	response.Success(ctx, "退出")
}
