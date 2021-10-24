package user_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-blog/app/exception"
	"go-blog/model/domain"
	"go-blog/model/web"
)

type userRepository struct{}

func NewRepository() UserRepository {
	return userRepository{}
}

func (u userRepository) Login(ctx context.Context, tx *sqlx.Tx, request web.UserRequest) {
	panic("implement me")
}

func (u userRepository) Register(ctx context.Context, tx *sqlx.Tx, request web.UserRequest) bool{
	result,err := tx.NamedExecContext(ctx,"INSERT INTO users(email,username,password) VALUES(:email,:username,:password)",&request)
	exception.PanicDuplicate(err)
	affected,_ :=result.RowsAffected()
	return affected > 0
}

func (u userRepository) FindByEmail(ctx context.Context, tx *sqlx.Tx,email string) domain.Users {
	user := domain.Users{}
	err := tx.Get(&user, "SELECT id,email,username,password FROM users WHERE email = $1", email)
	exception.PanicNotFound(err)
	return user
}

