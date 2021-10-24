package article_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go-blog/helper"
	"go-blog/model/domain"
	"go-blog/model/web"
	"go-blog/repository/article_repository"
)

type articleService struct {
	db *sqlx.DB
	validate *validator.Validate
	repo article_repository.ArticleRepository
}

func NewService(db *sqlx.DB, validate *validator.Validate, repo article_repository.ArticleRepository) ArticleService {
	return articleService{
		db:       db,
		validate: validate,
		repo:     repo,
	}
}

func (a articleService) Get(ctx context.Context) []domain.Articles {
	tx := a.db.MustBegin()
	defer helper.CommitOrRollback(tx)
	articles := a.repo.Get(ctx,tx)
	return articles
}

func (a articleService) Post(ctx context.Context, request web.ArticleRequest) bool {
	err := a.validate.Struct(request)
	helper.PanicIfError(err)
	tx := a.db.MustBegin()
	defer helper.CommitOrRollback(tx)
	return a.repo.Post(ctx,tx,request)
}

func (a articleService) Delete(ctx context.Context, articleId string) bool {
	tx := a.db.MustBegin()
	defer helper.CommitOrRollback(tx)
	find := a.repo.GetById(ctx,tx,articleId)
	return a.repo.Delete(ctx,tx,find.Id)
}

func (a articleService) Update(ctx context.Context, articleId string, request web.ArticleRequest) bool {
	err := a.validate.Struct(request)
	helper.PanicIfError(err)
	tx := a.db.MustBegin()
	defer helper.CommitOrRollback(tx)
	find := a.repo.GetById(ctx,tx,articleId)
	return a.repo.Update(ctx,tx,find.Id,request)
}
