package repository

import (
	"context"

	"github.com/jwcen/mars/internal/apiserver/domain"
	"github.com/jwcen/mars/internal/apiserver/repository/dao"
	"github.com/jwcen/mars/internal/apiserver/repository/model"
)

type ArticleRepository interface {
	Create(ctx context.Context, article domain.Article) (int64, error)
	UpdateById(ctx context.Context, article domain.Article) error
}

type articleRepository struct {
	dao dao.ArticleDao
}

func NewArticleRepository(dao dao.ArticleDao) ArticleRepository {
	return &articleRepository{
		dao: dao,
	}
}

func (repo *articleRepository) Create(ctx context.Context, article domain.Article) (int64, error) {
	return repo.dao.Insert(ctx, model.ArticleM{
		Title:    article.Title,
		Content:  article.Content,
		AuthorId: article.Author.Id,
	})
}

func (repo *articleRepository) UpdateById(ctx context.Context, article domain.Article) error {
	return repo.dao.UpdateById(ctx, model.ArticleM{
		Id:       article.Id,
		Title:    article.Title,
		Content:  article.Content,
		AuthorId: article.Author.Id,
	})
}
