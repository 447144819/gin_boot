package dto

type Pagination struct {
	Page  int `form:"page"`  // 页数
	Limit int `form:"limit"` // 每页几条
}
