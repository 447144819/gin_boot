package dto

type Pagination struct {
	PageNum  int `form:"page_num"`  // 页数
	PageSize int `form:"page_size"` // 每页几条
}
