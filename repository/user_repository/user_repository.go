package user_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-blog/model/domain"
	"go-blog/model/web"
)

type UserRepository interface {
	Login(ctx context.Context, tx *sqlx.Tx, request web.UserRequest)
	FindByEmail(ctx context.Context, tx *sqlx.Tx, email string) domain.Users
	Register(ctx context.Context, tx *sqlx.Tx, request web.UserRequest) bool
}
