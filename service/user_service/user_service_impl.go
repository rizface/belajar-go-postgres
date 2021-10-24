package user_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go-blog/helper"
	"go-blog/model/web"
	"go-blog/repository/user_repository"
)

type userService struct {
	db *sqlx.DB
	validate *validator.Validate
	repo user_repository.UserRepository
}

func NewService(db *sqlx.DB, validate *validator.Validate, repo user_repository.UserRepository) UserService {
	return userService{
		db:       db,
		validate: validate,
		repo:     repo,
	}
}

func (u userService) Login(ctx context.Context, request web.UserRequest) string{
	tx := u.db.MustBegin()
	defer helper.CommitOrRollback(tx)
	user := u.repo.FindByEmail(ctx,tx,request.Email)
	helper.Compare(request.Password,user.Password)
	token := helper.GenerateToken(user)
	return token
}

func (u userService) Register(ctx context.Context, request web.UserRequest) bool {
	err := u.validate.Struct(request)
	helper.PanicIfError(err)
	pass,err := helper.Generate(request.Password)
	helper.PanicIfError(err)
	request.Password = string(pass)

	tx := u.db.MustBegin()
	defer helper.CommitOrRollback(tx)
	return u.repo.Register(ctx,tx,request)
}

