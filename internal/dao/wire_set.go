package dao

import "github.com/google/wire"

// DaoSet 是所有 DAO 构造函数的集合（ProviderSet）
var DaoSet = wire.NewSet(
	NewUserDao,

	// 未来新增 DAO，只需在这里追加即可
)
