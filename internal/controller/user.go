package controller

import (
	"gin_boot/internal/dto"
	"gin_boot/internal/router/routers"
	"gin_boot/internal/service"
	"gin_boot/internal/utils/validator"
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	svc service.UserService
}

func NewUserController(svc service.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (h *UserController) RegisterRoutes(server *gin.Engine) {
	apiv1 := server.Group(routers.RouterBase.APIV1)
	user := apiv1.Group("/user")
	user.POST("/add", h.Create)
	user.PUT("/edit", h.Edit)
	user.DELETE("/delete/:id", h.Delete)
	user.GET("/:id", h.Detail)
	user.GET("/list", h.List)
	user.POST("/login", h.Login)
	user.GET("/logout", h.Logout)
}

func (u *UserController) Create(ctx *gin.Context) {
	var req dto.UserCreateDTO
	// 使用封装的验证器
	if err := validator.GinBind(ctx, &req); err != nil {
		response.Error(ctx, err, response.ParamsError)
		return
	}
	// 在 Gin 项目中，*gin.Context是 HTTP 请求的上下文，只应在 Controller 层使用；
	// 而 context.Context是 Go 标准的、跨层级的请求上下文，应该在 Service、Repository 等所有业务逻辑层使用。
	// 应该从 *gin.Context中提取 ctx.Request.Context()并传给 Service，而不是直接传 *gin.Context。
	// 这样可以保持分层清晰、代码解耦、便于扩展与维护。
	gctx := ctx.Request.Context() // 转为 context.Context
	err := u.svc.Create(gctx, req)
	if err != nil {
		response.Error(ctx, err, response.AddError)
		return
	}
	response.Success(ctx, response.AddSuccess)
}

func (u *UserController) Edit(ctx *gin.Context) {
	var req dto.UserEditDTO
	// 使用封装的验证器
	if errors := validator.GinBind(ctx, &req); errors != nil {
		response.Error(ctx, errors, response.ParamsError)
		return
	}

	gctx := ctx.Request.Context() // 转为 context.Context
	err := u.svc.Edit(gctx, req)
	if err != nil {
		response.Error(ctx, err, response.EditError)
		return
	}
	response.Success(ctx, response.EditSuccess)
}

func (u *UserController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if id < 1 {
		response.Error(ctx, nil, "参数错误")
	}
	err := u.svc.Delete(ctx.Request.Context(), id)
	if err != nil {
		response.Error(ctx, err, response.DeleteError)
		return
	}
	response.Success(ctx, response.DeleteSuccess)
}

func (u *UserController) Detail(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if id < 1 {
		response.Error(ctx, nil, response.ParamsError)
	}
	user, err := u.svc.Detail(ctx.Request.Context(), id)
	if err != nil {
		response.Error(ctx, err, response.ParamsError)
		return
	}
	response.SuccessData(ctx, user)
}

func (u *UserController) List(ctx *gin.Context) {
	var req dto.UserListDTO
	// 使用封装的验证器
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, err, response.ParamsError)
		return
	}

	data, total, err := u.svc.List(ctx.Request.Context(), &req)
	if err != nil {
		response.Error(ctx, err, response.DoError)
	}
	response.PageSuccess(ctx, data, total, req.Page, req.Limit)
}

func (u *UserController) Login(ctx *gin.Context) {
	var req dto.UserLoginDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, err, response.ParamsError)
		return
	}

	// 实现登录功能
	token, err := u.svc.Login(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err, "登录失败")
		return
	}

	response.SuccessData(ctx, gin.H{
		"token": token,
	}, "登录成功")
}

func (u *UserController) Logout(ctx *gin.Context) {
	response.Success(ctx, "退出")
}
