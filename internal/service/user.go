package service

import (
	"context"
	"errors"
	"gin_boot/internal/dao"
	"gin_boot/internal/dto"
	"gin_boot/internal/model"
	"gin_boot/internal/utils/hash"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{
		dao: dao,
	}
}

func (s UserService) Create(ctx context.Context, req dto.UserCreateDTO) error {
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

func (s *UserService) Delete(ctx context.Context, id int64) error {
	user, err := s.dao.FindById(ctx, id)
	if user.Id < 1 {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}

	return s.dao.Delete(ctx, id)
}

func (s UserService) Edit(ctx context.Context, req dto.UserEditDTO) error {
	user, err := s.dao.FindById(ctx, req.Id)
	if user.Id < 1 {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	user.Email = req.Email
	user.Phone = req.Phone
	user.Nickname = req.Nickname
	user.RoleId = req.RoleId
	return s.dao.Update(ctx, user)
}

func (s *UserService) Detial(ctx context.Context, id int64) (model.User, error) {
	user, err := s.dao.FindById(ctx, id)
	if user.Id < 1 {
		return model.User{}, errors.New("用户不存在")
	}
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
