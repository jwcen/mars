package dao

import (
	"context"

	"github.com/jwcen/mars/internal/apiserver/repository/model"
	"gorm.io/gorm"
)

type ArticleDao interface {
	Insert(ctx context.Context, article model.ArticleM) (int64, error)
}

type articleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) ArticleDao {
	return &articleDao{
		db: db,
	}
}

func (dao *articleDao) Insert(ctx context.Context, article model.ArticleM) (int64, error) {
	err := dao.db.WithContext(ctx).Create(&article).Error
	if err != nil {
		return 0, err
	}
	return article.Id, nil
}
