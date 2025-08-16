package tests

import (
	"gin_boot/internal/initializa/validator"
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
)

// UserRegisterRequest 用户注册请求
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" binding:"required,strong_password" label:"密码"`
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Phone    string `json:"phone" binding:"required,phone" label:"手机号"`
	Age      int    `json:"age" binding:"required,min=1,max=120" label:"年龄"`
}

// ExampleController 定义示例控制器
type ExampleController struct{}

// NewExampleController 创建 ExampleController 实例
func NewExampleController() *ExampleController {
	return &ExampleController{}
}

// CreateUser 处理创建用户的请求
func (ec *ExampleController) CreateUser(c *gin.Context) {
	var req UserRegisterRequest
	// 使用封装的验证器
	if errors := validator.GinBind(c, &req); errors != nil {
		response.Error(c, errors.Error())
		return
	}

	// 验证通过，处理业务逻辑
	response.Success(c, "用户创建成功")

}
