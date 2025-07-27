package repository

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jwcen/mars/internal/apiserver/domain"
	"github.com/jwcen/mars/internal/apiserver/repository/dao"
	daomocks "github.com/jwcen/mars/internal/apiserver/repository/dao/mocks"
	"github.com/jwcen/mars/internal/apiserver/repository/model"
	"github.com/stretchr/testify/assert"
)

func TestUserInfoRepository_Create(t *testing.T) {
	testCases := []struct {
		name    string
		mock    func(*gomock.Controller) dao.UserDao
		ctx     context.Context
		user    *domain.User
		wantErr error
	}{
		{
			name: "创建成功！",
			ctx:  context.Background(),
			user: &domain.User{
				Email:    "123@qq.com",
				Password: "admin123",
			},
			mock: func(c *gomock.Controller) dao.UserDao {
				daoDao := daomocks.NewMockUserDao(c)
				daoDao.EXPECT().Insert(gomock.Any(), &model.UserM{
					Email:    "123@qq.com",
					Password: "admin123",
				}).Return(nil)
				return daoDao
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			userDao := NewUserInfoRepository(tc.mock(ctrl))
			err := userDao.Create(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
