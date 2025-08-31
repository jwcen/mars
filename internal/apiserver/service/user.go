package service

import (
	"context"
	"errors"

	"github.com/jwcen/mars/internal/apiserver/domain"
	"github.com/jwcen/mars/internal/apiserver/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

var _ UserAndService = &UserService{}

type UserAndService interface {
	Signup(ctx context.Context, u domain.User) error
	Login(ctx context.Context, email, password string) (domain.User, error)
	FindOrCreate(ctx context.Context, phone string) (domain.User, error)
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserAndService {
	return &UserService{
		repo: r,
	}
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPwd)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	// 查找用户是否存在
	u, err := svc.repo.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}

	// 返回用户
	return u, nil
}

func (svc *UserService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	// 这时候，这个地方要怎么办？
	// 这个叫做快路径
	user, err := svc.repo.FindByPhone(ctx, phone)
	if err != repository.ErrUserNotFound {
		// 绝大部分请求进来这里
		// nil 会进来这里
		// 不为 ErrUserNotFound 的也会进来这里
		return user, err
	}

	// TODO: logger
	// 这个叫做慢路径
	// 你明确知道，没有这个用户
	user = domain.User{
		Phone: phone,
	}
	err = svc.repo.Create(ctx, user)
	if err != nil && err != repository.ErrUserDuplicate {
		return user, err
	}

	// 因为这里会遇到主从延迟的问题
	return svc.repo.FindByPhone(ctx, phone)
}
