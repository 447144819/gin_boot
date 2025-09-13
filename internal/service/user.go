package service

import (
	"context"
	"errors"
	"gin_boot/internal/dao"
	"gin_boot/internal/dto"
	"gin_boot/internal/model"
	"gin_boot/internal/utils/captcha"
	"gin_boot/internal/utils/hash"
	"gin_boot/internal/vo"
	"gin_boot/pkg/jwts"
)

// UserService 定义服务行为（接口）
type UserService interface {
	ModelToVo(user model.User) vo.UserInfoVO
	Create(ctx context.Context, req dto.UserCreateDTO) error
	Delete(ctx context.Context, id uint64) error
	Edit(ctx context.Context, req dto.UserEditDTO) error
	Detail(ctx context.Context, id uint64) (vo.UserInfoVO, error)
	List(ctx context.Context, req *dto.UserListDTO) ([]vo.UserInfoVO, int64, error)
	Login(ctx context.Context, req dto.UserLoginDTO) (string, error)
}

// userServiceImpl 是接口的实际实现（包内实现，不对外暴露）
type userServiceImpl struct {
	dao        *dao.UserDao
	redisStore *captcha.RedisStore
}

// NewUserService 是构造函数，返回接口类型
func NewUserService(dao *dao.UserDao, redisStore *captcha.RedisStore) UserService {
	return &userServiceImpl{
		dao:        dao,
		redisStore: redisStore,
	}
}

// modelToVo model 转vo
func (s *userServiceImpl) ModelToVo(user model.User) vo.UserInfoVO {
	return vo.UserInfoVO{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Email:    user.Email,
		RoleId:   user.RoleId,
		Ctime:    user.Ctime,
	}
}

func (s *userServiceImpl) Create(ctx context.Context, req dto.UserCreateDTO) error {
	// 判断是否存在
	user, err := s.dao.FindByUsername(ctx, req.Username)
	if err != nil {
		return err
	}
	if user.Id > 0 {
		return errors.New("用户已存在")
	}

	// 设置密码
	req.Password = hash.BcryptMake(req.Password)
	return s.dao.Create(ctx, &model.User{
		Username: req.Username,
		Password: req.Password,
		RoleId:   req.RoleId,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Email:    req.Email,
	})
}

func (s *userServiceImpl) Delete(ctx context.Context, id uint64) error {
	user, err := s.dao.FindById(ctx, id)
	if user.Id < 1 {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}

	return s.dao.Delete(ctx, id)
}

func (s *userServiceImpl) Edit(ctx context.Context, req dto.UserEditDTO) error {
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
	return s.dao.Update(ctx, &user)
}

func (s *userServiceImpl) Detail(ctx context.Context, id uint64) (vo.UserInfoVO, error) {
	user, err := s.dao.FindById(ctx, id)
	if user.Id < 1 {
		return vo.UserInfoVO{}, errors.New("用户不存在")
	}
	if err != nil {
		return vo.UserInfoVO{}, err
	}
	return s.ModelToVo(user), nil
}

func (s *userServiceImpl) List(ctx context.Context, req *dto.UserListDTO) ([]vo.UserInfoVO, int64, error) {
	var userInfo []vo.UserInfoVO
	where := map[string]interface{}{
		"phone": req.Phone,
		"email": req.Email,
	}
	if req.Username != "" {
		where["username like ?"] = "%" + req.Username + "%"
	}
	if req.Nickname != "" {
		where["nickname like ?"] = "%" + req.Nickname + "%"
	}
	users, total, err := s.dao.PageQuery(ctx, req.Page, req.Limit, where, "id desc", []string{})
	if err != nil {
		return userInfo, total, err
	}

	for _, user := range users {
		userInfo = append(userInfo, s.ModelToVo(user))
	}
	return userInfo, total, nil
}

func (s *userServiceImpl) Login(ctx context.Context, req dto.UserLoginDTO) (string, error) {
	// 验证码校验
	isValid := s.redisStore.Verify(req.CodeId, req.Code, true)
	if !isValid {
		//return "",errors.New("验证码错误或已过期")
	}

	// 查询用户
	user, err := s.dao.FindByUsername(ctx, req.Username)
	if user.Id < 1 {
		return "", errors.New("用户不存在")
	}
	if err != nil {
		return "", err
	}

	// 对比密码
	isValid = hash.BcryptCheck(req.Password, user.Password)
	if !isValid {
		return "", errors.New("密码错误")
	}

	// 生成token
	return jwts.NewJWTHandler().SetJWTToken(int64(user.Id), user.Username)
}
