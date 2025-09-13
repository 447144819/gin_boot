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
// UserDao 
type UserDao struct {
    *basedao.BaseDao[model.User, uint64]
}

// NewUserDao 是构造函数，返回接口类型
func NewUserDao(db *gorm.DB) *UserDao {
   // 自动创建表
   db.AutoMigrate(&model.User{})
   return &UserDao{
      basedao.NewBaseDao[model.User, uint64](db),
   }
}
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
	dao *dao.UserDao
}

// NewUserService 是构造函数，返回接口类型
func NewUserService(dao *dao.UserDao) UserService {
	return &userServiceImpl{
		dao: dao,
	}
}
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

5. 生成wire

在更目录执行命令，会自动生成路由和wire文件
```go
// 安装 wire
go install github.com/google/wire/cmd/wire@latest


// 执行命令，生成
go run .\cmd\runwire.go

```

### 自动创建模块
```angular2html
go run .\cmd\auto\main.go test1
go run cmd/auto/main.go test1
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
