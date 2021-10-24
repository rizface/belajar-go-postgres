package setup

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"go-blog/app"
	middleware2 "go-blog/app/middleware"
	"go-blog/controller/articles_controller"
	"go-blog/repository/article_repository"
	"go-blog/service/article_service"
)

func Articles(r *chi.Router) *chi.Router {
	repo := article_repository.NewRepository()
	service := article_service.NewService(Db,validator.New(),repo)
	controller := articles_controller.NewController(service)
	router := *r

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware2.ErrorHandler)
	router.Use(middleware2.Authentication)

	router.Get(app.ARTICLES, controller.Get)
	router.Post(app.ARTICLES,controller.Post)
	router.Delete(app.ARTICLES_DETAIL,controller.Delete)
	router.Put(app.ARTICLES_DETAIL,controller.Update)
	return &router
}

