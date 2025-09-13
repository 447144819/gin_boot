package dao

import (
	"context"
	"gin_boot/internal/dao/basedao"
	"gin_boot/internal/model"
	"gorm.io/gorm"
)

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

func (d *UserDao) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := d.DB.WithContext(ctx).Where("username = ?", username).Find(&user).Error
	return user, err
}
