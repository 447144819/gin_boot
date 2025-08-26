package service

import "github.com/google/wire"

// ServiceSet 是所有 Service 构造函数的集合
var ServiceSet = wire.NewSet(
	NewUserService,
	NewCaptchaService,

	// 未来新增 Service，只需在这里追加即可
)
