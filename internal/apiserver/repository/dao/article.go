package dao

import (
	"context"
	"time"

	"github.com/jwcen/mars/internal/apiserver/repository/model"
	"gorm.io/gorm"
)

type ArticleDao interface {
	Insert(ctx context.Context, article model.ArticleM) (int64, error)
	UpdateById(ctx context.Context, article model.ArticleM) error
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

func (dao *articleDao) UpdateById(ctx context.Context, article model.ArticleM) error {
	err := dao.db.WithContext(ctx).Model(&model.ArticleM{}).Where("id = ?", article.Id).
		Updates(map[string]any{
			"title":      article.Title,
			"content":    article.Content,
			"updated_at": time.Now(),
		}).Error
	return err
}
