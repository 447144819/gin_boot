# gin_boot
gin框架脚手架



### 常用返回
```
response.Success(ctx)
response.Success(ctx, "用户创建成功")
response.SuccessData(ctx, "用户详情data", "用户创建成功")

response.Error(ctx)
response.Error(ctx, "用户创建失败")
response.ErrorWithCode(ctx, 201)
response.ErrorWithCode(ctx, 203, "用户创建失败")
```
    