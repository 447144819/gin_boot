package controller

import "github.com/google/wire"

// ControllerSet 是所有 Controller 构造函数的集合
var ControllerSet = wire.NewSet(
	NewUserController,
	NewCaptchaController,

	// 未来新增 Controller，只需在这里追加即可
)
