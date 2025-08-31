package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jwcen/mars/internal/apiserver/repository/model"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicate = errors.New("邮箱冲突")
	ErrUserNotFound  = gorm.ErrRecordNotFound
)

const uniqueIndexErr uint16 = 1062

type UserDao interface {
	Insert(ctx context.Context, u model.UserM) error
	FindByEmail(ctx context.Context, email string) (model.UserM, error)
	FindByPhone(ctx context.Context, phone string) (model.UserM, error)

	UserExpansion
}

// UserExpansion 定义了用户操作的附加方法.
type UserExpansion interface{}

type userDao struct {
	db *gorm.DB
}

var _ UserDao = (*userDao)(nil)

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{
		db: db,
	}
}

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&model.UserM{})
}

func (dao *userDao) Insert(ctx context.Context, u model.UserM) error {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now

	err := dao.db.WithContext(ctx).Create(&u).Error
	if e, ok := err.(*mysql.MySQLError); ok {
		if e.Number == uniqueIndexErr {
			return errors.New("用户已存在")
		}
	}

	return err
}

func (dao *userDao) FindByEmail(ctx context.Context, email string) (model.UserM, error) {
	var u model.UserM
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error

	return u, err
}

func (dao *userDao) FindByPhone(ctx context.Context, phone string) (model.UserM, error) {
	var u model.UserM
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&u).Error

	return u, err
}
