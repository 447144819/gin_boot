### gin_boot
gin框架脚手架
基于wire管理依赖注入，配置文件管理，日志打印，响应返回封装，错误处理，中间件，jwt验证，mysql数据库操作，redis数据库操作，验证码生成，

### 使用流程
1. 创建 model
   /gin_boot/internal/model/user.go
```go
package model

import "gorm.io/plugin/soft_delete"

type User struct {
	Id       uint64                `gorm:"primary_key;auto_increment"`
	Username string                `gorm:"unique;type:varchar(30);not null;comment:用户名"`
	Password string                `gorm:"type:varchar(200);not null;comment:密码"`
	...
}
```
2. 创建 dao
   /gin_boot/internal/dao/user.go
```go
// UserDao 定义服务行为（接口）
type UserDao interface {
	Create(ctx context.Context, req dto.UserCreateDTO) error
}

// userDaoImpl 是接口的实际实现（包内实现，不对外暴露）
type userDaoImpl struct {
	db *gorm.DB
}

// NewUserDao 是构造函数，返回接口类型
func NewUserDao(db *gorm.DB) UserDao {
	// 自动创建表
	db.AutoMigrate(&model.User{})
	return &userDaoImpl{
		db: db,
	}
}

```
#### 设置wire
/gin_boot/internal/dao/wire_set.go
```go
package dao

import "github.com/google/wire"

// DaoSet 是所有 DAO 构造函数的集合（ProviderSet）
var DaoSet = wire.NewSet(
	NewUserDao,

	// 未来新增 DAO，只需在这里追加即可
)
```
3. 创建 service
   /gin_boot/internal/service/router.go
```go

// UserService 定义服务行为（接口）
type UserService interface {
	Create(ctx context.Context, req dto.UserCreateDTO) error
}

// userServiceImpl 是接口的实际实现（包内实现，不对外暴露）
type userServiceImpl struct {
	dao dao.UserDao
}

// NewUserService 是构造函数，返回接口类型
func NewUserService(dao dao.UserDao) UserService {
	return &userServiceImpl{
		dao: dao,
	}
}
```

#### 设置wire
/gin_boot/internal/service/wire_set.go
```go
package service

import "github.com/google/wire"

// ServiceSet 是所有 Service 构造函数的集合
var ServiceSet = wire.NewSet(
	NewUserService,

	// 未来新增 Service，只需在这里追加即可
)

```
4. 创建 controller
   /gin_boot/internal/controller/router.go
```go

type UserController struct {
	svc service.UserService
}

func NewUserController(svc service.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (h *UserController) RegisterRoutes(server *common.RouteContext) {
	user := server.APIV1.Group("/user")
	user.POST("/add", h.Create)
	user.PUT("/edit", h.Edit)
	user.DELETE("/:id", h.Delete)
	user.POST("/login", h.Login)
}

```
#### 设置wire
/gin_boot/internal/controller/wire_set.go
```go
package controller

import "github.com/google/wire"

// ControllerSet 是所有 Controller 构造函数的集合
var ControllerSet = wire.NewSet(
	NewUserController,

	// 未来新增 Controller，只需在这里追加即可
)

```
5. 设置路由
   /gin_boot/internal/router/router.go
```go

func NewAllHandlers(
	userHandler *controller.UserController,
) []common.RouteRegistrar {
	return []common.RouteRegistrar{
		userHandler,
		
		// 新增的 Handler 直接加入这个切片
	}
}
```


### 日志打印
```
logs.Info("🚀 hello lzw" )
logs.Error("🚀 系统错误" )
```

### 响应返回
```
response.Success(ctx)
response.Success(ctx, "用户创建成功")
response.SuccessData(ctx, "用户详情data", "用户创建成功")

response.Error(ctx)
response.Error(ctx, "用户创建失败")
response.ErrorWithCode(ctx, 201)
response.ErrorWithCode(ctx, 203, "用户创建失败")
```

### // 注册路由
在控制器中实现RegisterRoutes方法
```angular2html
func (c *Captcha) RegisterRoutes(server *common.RouteContext) {
server.APIV1.GET("/captcha", c.GetCaptcha)
}
```
