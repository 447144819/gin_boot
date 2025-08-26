# gin_boot
gin框架脚手架
基于wire管理依赖注入，配置文件管理，日志打印，响应返回封装，错误处理，中间件，jwt验证，mysql数据库操作，redis数据库操作，验证码生成，

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
