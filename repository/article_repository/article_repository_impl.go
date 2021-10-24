package article_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-blog/app/exception"
	"go-blog/helper"
	"go-blog/model/domain"
	"go-blog/model/web"
)

type articleRepository struct{}

func NewRepository() ArticleRepository {
	return articleRepository{}
}

func (a articleRepository) Get(ctx context.Context, tx *sqlx.Tx) []domain.Articles {
	articles := []domain.Articles{}
	err := tx.SelectContext(ctx,&articles,"SELECT id,title,content,created_at FROM contents ORDER BY id DESC LIMIT 30")
	helper.PanicIfError(err)
	return articles
}

func (a articleRepository) Post(ctx context.Context, tx *sqlx.Tx, request web.ArticleRequest) bool {
	result,err := tx.NamedExecContext(ctx,"INSERT INTO contents(title,content,user_id) VALUES(:title,:content, :user_id)",request)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (a articleRepository) GetById(ctx context.Context, tx *sqlx.Tx, id string) domain.Articles {
	article := domain.Articles{}
	err := tx.GetContext(ctx,&article,"SELECT id,title,content,created_at FROM contents ORDER BY id DESC LIMIT 30")
	exception.PanicNotFound(err)
	return article
}

func (a articleRepository) Delete(ctx context.Context, tx *sqlx.Tx, id string) bool {
	result,err := tx.ExecContext(ctx,"DELETE FROM contents WHERE id = $1",id)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}

func (a articleRepository) Update(ctx context.Context, tx *sqlx.Tx, id string, request web.ArticleRequest) bool {
	result,err := tx.ExecContext(ctx,"UPDATE contents SET title = $1, content = $2 WHERE id = $3", request.Title,request.Content,id)
	helper.PanicIfError(err)
	affected,err := result.RowsAffected()
	helper.PanicIfError(err)
	return affected > 0
}
