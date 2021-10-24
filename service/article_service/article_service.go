package article_service

import (
	"context"
	"go-blog/model/domain"
	"go-blog/model/web"
)

type ArticleService interface{
	Get(ctx context.Context) []domain.Articles
	Post(ctx context.Context, request web.ArticleRequest) bool
	Delete(ctx context.Context, articleId string) bool
	Update(ctx context.Context, articleId string, request web.ArticleRequest) bool
}
