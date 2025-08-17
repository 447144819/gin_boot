package dao

import (
	"gin_boot/internal/model"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	// 自动创建表
	db.AutoMigrate(&model.User{})
	return &UserDao{
		db: db,
	}
}
