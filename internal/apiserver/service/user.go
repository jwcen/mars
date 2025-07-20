package service

import (
	"context"

	"github.com/jwcen/mars/internal/apiserver/domain"
	"github.com/jwcen/mars/internal/apiserver/repository"
)

type UserAndService interface {
	Signup(ctx context.Context, u *domain.User) error
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserAndService {
	return &UserService{
		repo: r,
	}
}

func (svc *UserService) Signup(ctx context.Context, u *domain.User) error {
	//hashPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return err
	//}
	//u.Password = string(hashPwd)
	return svc.repo.Create(ctx, u)
}
