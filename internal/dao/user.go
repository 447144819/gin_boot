package dao

import (
	"gin_boot/internal/dto"
	"gin_boot/internal/model"
	"github.com/gin-gonic/gin"
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

func (d *UserDao) FindByUsername(ctx *gin.Context, username string) (model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).Where("username = ?", username).Find(&user).Error
	return user, err
}

func (d *UserDao) Create(ctx *gin.Context, req dto.UserCreateDTO) error {
	res := d.db.WithContext(ctx).Create(&model.User{
		Username: req.Username,
		Password: req.Password,
		RoleId:   req.RoleId,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
	})
	return res.Error
}

func (d *UserDao) FindById(ctx *gin.Context, id int64) (model.User, error) {
	var user model.User
	res := d.db.WithContext(ctx).Where("id=?", id).First(&user)
	return user, res.Error
}

func (d *UserDao) Delete(ctx *gin.Context, id int64) error {
	var user model.User
	return d.db.WithContext(ctx).Delete(&user, id).Error
}
