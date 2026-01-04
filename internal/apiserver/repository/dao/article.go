package dao

import (
	"context"
	"errors"
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
	res := dao.db.WithContext(ctx).
		Model(&model.ArticleM{}).
		Where("id = ? AND author_id = ?", article.Id, article.AuthorId).
		Updates(map[string]any{
			"title":      article.Title,
			"content":    article.Content,
			"updated_at": time.Now(),
		})
	if res.RowsAffected == 0 {
		return errors.New("更新失败，文章不存在或不是作者")
	}
	return res.Error
}
