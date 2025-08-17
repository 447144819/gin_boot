package service

import "gin_boot/internal/dao"

type UserService struct {
	dao *dao.UserDao
}

func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{
		dao: dao,
	}
}
