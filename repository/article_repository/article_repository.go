package article_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-blog/model/domain"
	"go-blog/model/web"
)

type ArticleRepository interface{
	Get(ctx context.Context, tx *sqlx.Tx) []domain.Articles
	Post(ctx context.Context, tx *sqlx.Tx, request web.ArticleRequest) bool
	GetById(ctx context.Context, tx *sqlx.Tx, id string) domain.Articles
	Delete(ctx context.Context, tx *sqlx.Tx, id string) bool
	Update(ctx context.Context, tx *sqlx.Tx, id string, request web.ArticleRequest) bool
}