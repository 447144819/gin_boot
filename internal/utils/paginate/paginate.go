package paginate

import "gorm.io/gorm"

// Paginate 分页
func Paginate(PageNumber int, PageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if PageSize > 200 {
			PageSize = 200
		}
		offset := (PageNumber - 1) * PageSize
		return db.Offset(offset).Limit(PageSize)
	}
}
