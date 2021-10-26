package setup

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-blog/app"
	middleware2 "go-blog/app/middleware"
	"go-blog/controller/funcfact_controller"
	"go-blog/service/funfact_service"
)

func Author(r *chi.Router) *chi.Router {
	router := *r

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware2.ErrorHandler)
	router.Use(middleware2.Authentication)



	service := funfact_service.NewService()
	controller := funcfact_controller.NewController(service)

	router.Post(app.AUTHOR_FUNFACT, controller.PostFuncFact)
	router.Put(app.AUTHOR_FUNFACT,controller.PutFunFact)
	router.Delete(app.AUTHOR_FUNFACT,controller.DeleteFunFact)

	return &router
}
