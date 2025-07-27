package repository

import (
	"context"

	"github.com/jwcen/mars/internal/apiserver/domain"
	"github.com/jwcen/mars/internal/apiserver/repository/dao"
	"github.com/jwcen/mars/internal/apiserver/repository/model"
)

type UserRepository interface {
	Create(ctx context.Context, u *domain.User) error
}

type UserInfoRepository struct {
	dao dao.UserDao
}

func NewUserInfoRepository(dao dao.UserDao) UserRepository {
	return &UserInfoRepository{
		dao: dao,
	}
}

func (ur *UserInfoRepository) Create(ctx context.Context, u *domain.User) error {
	return ur.dao.Insert(ctx, &model.UserM{
		Id:        u.Id,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	})
}
