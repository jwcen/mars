package service

import (
	"context"

	"github.com/jwcen/mars/internal/apiserver/domain"
	"github.com/jwcen/mars/internal/apiserver/repository"
)

type ArticleService interface {
	Save(ctx context.Context, article domain.Article) (int64, error)
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{
		repo: repo,
	}
}

func (svc *articleService) Save(ctx context.Context, article domain.Article) (int64, error) {
	return svc.repo.Create(ctx, article)
}
