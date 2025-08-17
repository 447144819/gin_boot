package service

import (
	"errors"
	"gin_boot/internal/dao"
	"gin_boot/internal/dto"
	"gin_boot/internal/utils/hash"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{
		dao: dao,
	}
}

func (s UserService) Create(ctx *gin.Context, req dto.UserCreateDTO) error {
	// 判断用户是否存在
	user, err := s.dao.FindByUsername(ctx, req.Username)
	if user.Id > 0 {
		return errors.New("用户已存在")
	}
	if err != nil {
		return err
	}
	// 设置密码
	req.Password = hash.BcryptMake(req.Password)
	err = s.dao.Create(ctx, req)
	return err
}

func (s *UserService) Delete(ctx *gin.Context, id int64) error {
	user, err := s.dao.FindById(ctx, id)
	if user.Id < 1 {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}

	return s.dao.Delete(ctx, id)
}
