package articles_controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-blog/helper"
	"go-blog/model/web"
	"go-blog/service/article_service"
	"net/http"
)

type articleController struct {
	service article_service.ArticleService
}

func NewController(service article_service.ArticleService) ArticlesController {
	return articleController{
		service: service,
	}
}

func (a articleController) Get(w http.ResponseWriter, r *http.Request) {
	articles := a.service.Get(r.Context())
	helper.JsonWriter(w,http.StatusOK, "success", map[string]interface{}{
		"articles":articles,
	})
}

func (a articleController) Post(w http.ResponseWriter, r *http.Request) {
	request := web.ArticleRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)
	claim := r.Context().Value("user-data")
	request.User_Id = claim.(*helper.UserClaims).Id
	success := a.service.Post(r.Context(),request)
	if success {
		helper.JsonWriter(w,http.StatusOK, "contest posted", nil)
	}
}

func (a articleController) Delete(w http.ResponseWriter, r *http.Request) {
	articleId :=  chi.URLParam(r,"articleId")
	success := a.service.Delete(r.Context(),articleId)
	if success {
		helper.JsonWriter(w,http.StatusOK, "content deleted", nil)
	}
}

func (a articleController) Update(w http.ResponseWriter, r *http.Request) {
	articleId :=  chi.URLParam(r,"articleId")
	request := web.ArticleRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)
	success := a.service.Update(r.Context(),articleId,request)
	if success {
		helper.JsonWriter(w,http.StatusOK,"content updated", nil)
	}
}
