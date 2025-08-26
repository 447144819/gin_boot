package dao

import (
	"context"
	"gin_boot/internal/dto"
	"gin_boot/internal/model"
	"gin_boot/internal/utils/page"
	"gorm.io/gorm"
)

// UserDao 定义服务行为（接口）
type UserDao interface {
	FindByUsername(ctx context.Context, username string) (model.User, error)
	Create(ctx context.Context, req dto.UserCreateDTO) error
	FindById(ctx context.Context, id int64) (model.User, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user model.User) error
	List(ctx context.Context, req *dto.UserListDTO) ([]model.User, int64, error)
}

// userDaoImpl 是接口的实际实现（包内实现，不对外暴露）
type userDaoImpl struct {
	db *gorm.DB
}

// NewUserDao 是构造函数，返回接口类型
func NewUserDao(db *gorm.DB) UserDao {
	// 自动创建表
	db.AutoMigrate(&model.User{})
	return &userDaoImpl{
		db: db,
	}
}

func (d *userDaoImpl) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).Where("username = ?", username).Find(&user).Error
	return user, err
}

func (d *userDaoImpl) Create(ctx context.Context, req dto.UserCreateDTO) error {
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

func (d *userDaoImpl) FindById(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	res := d.db.WithContext(ctx).Where("id=?", id).First(&user)
	return user, res.Error
}

func (d *userDaoImpl) Delete(ctx context.Context, id int64) error {
	var user model.User
	return d.db.WithContext(ctx).Delete(&user, id).Error
}

func (d *userDaoImpl) Update(ctx context.Context, user model.User) error {
	return d.db.WithContext(ctx).Updates(&user).Error
}

func (d *userDaoImpl) List(ctx context.Context, req *dto.UserListDTO) ([]model.User, int64, error) {
	var total int64
	var users []model.User
	// 1. 构造基础查询（带上下文和模型）
	db := d.db.WithContext(ctx).Model(&model.User{}).Order("id desc")

	// 2. 添加筛选条件
	if req.Username != "" {
		db.Where("username = ?", "%"+req.Username+"%")
	}
	if req.Phone != "" {
		db.Where("phone = ?", req.Phone)
	}
	if req.Email != "" {
		db.Where("email = ?", req.Email)
	}
	if req.Nickname != "" {
		db.Where("nickname = ?", req.Nickname)
	}

	// 3. 分页查询
	err := db.Count(&total).Scopes(page.Paginate(req.Page, req.Page)).Find(&users).Error
	return users, total, err
}
