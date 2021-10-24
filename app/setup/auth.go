package setup

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"go-blog/app"
	"go-blog/controller/user_controller"
	"go-blog/helper"
	"go-blog/repository/user_repository"
	"go-blog/service/user_service"
)

func Auth(r *chi.Mux) *chi.Mux{
	if Err != nil {
		helper.PanicIfError(Err)
	}
	repo := user_repository.NewRepository()
	service := user_service.NewService(Db, validator.New(),repo)
	controller := user_controller.NewController(service)

	r.Post(app.REGISTER,controller.Register)
	r.Post(app.LOGIN, controller.Login)

	return r
}
