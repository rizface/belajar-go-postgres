package test

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go-blog/app"
	"go-blog/app/setup"
	"go-blog/model/web"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := setup.Router()
		r.Route("/articles", func(r chi.Router) {
			token,_ := ioutil.ReadFile("token.txt")
			re := *setup.Articles(&r)
			request := httptest.NewRequest(http.MethodGet,app.ARTICLES,nil)
			request.Header.Add("Authorization", "Bearer " + string(token))
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder,request)
			assert.Equal(t,http.StatusOK,recorder.Code)
		})
	})
	t.Run("failed", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/articles", func(r chi.Router) {
			re := *setup.Articles(&r)
			request := httptest.NewRequest(http.MethodGet,app.ARTICLES,nil)
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder,request)
			assert.Equal(t,http.StatusForbidden,recorder.Code)
		})
	})
}

func TestPostArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/articles", func(r chi.Router) {
			token,_ := ioutil.ReadFile("token.txt")
			re := *setup.Articles(&r)

			data := web.ArticleRequest{
				User_Id: 0,
				Title:   "kedua",
				Content: "kedua",
			}
			dataJson,_ := json.Marshal(data)
			reader := bytes.NewReader(dataJson)

			request := httptest.NewRequest(http.MethodPost,app.ARTICLES,reader)
			request.Header.Add("Authorization", "Bearer " + string(token))
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder,request)
			assert.Equal(t,http.StatusOK,recorder.Code)
		})
	})
	t.Run("invalid token", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/articles", func(r chi.Router) {
			re := *setup.Articles(&r)

			data := web.ArticleRequest{
				User_Id: 0,
				Title:   "kedua",
				Content: "kedua",
			}
			dataJson,_ := json.Marshal(data)
			reader := bytes.NewReader(dataJson)

			request := httptest.NewRequest(http.MethodPost,app.ARTICLES,reader)
			request.Header.Add("Authorization", "Bearer token")
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder,request)
			assert.Equal(t,http.StatusForbidden,recorder.Code)
		})
	})
}

func TestDeleteArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/articles", func(r chi.Router) {
			token, _ := ioutil.ReadFile("token.txt")
			re := *setup.Articles(&r)
			request := httptest.NewRequest(http.MethodDelete, "/articles/4c982d65-6f17-47df-b40e-8cf4170a9440", nil)
			request.Header.Add("Authorization", "Bearer "+string(token))
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder, request)
			assert.Equal(t, http.StatusOK, recorder.Code)
		})
	})
	t.Run("notfound", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/articles", func(r chi.Router) {
			token, _ := ioutil.ReadFile("token.txt")
			re := *setup.Articles(&r)
			request := httptest.NewRequest(http.MethodDelete, "/articles/4c982d65-6f17-47df-b40e-8cf41asd70a9440", nil)
			request.Header.Add("Authorization", "Bearer "+string(token))
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder, request)
			assert.Equal(t, http.StatusNotFound, recorder.Code)
		})
	})
}

func TestPutArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/articles", func(r chi.Router) {
			token,_ := ioutil.ReadFile("token.txt")
			re := *setup.Articles(&r)

			data := web.ArticleRequest{
				User_Id: 0,
				Title:   "samsul",
				Content: "samsul",
			}
			dataJson,_ := json.Marshal(data)
			reader := bytes.NewReader(dataJson)

			request := httptest.NewRequest(http.MethodPut,"/articles/6477da1b-29cc-4e6a-812b-bfe4d13167b4",reader)
			request.Header.Add("Authorization", "Bearer " + string(token))
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder,request)
			assert.Equal(t,http.StatusOK,recorder.Code)
		})
	})
	t.Run("invalid token", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/articles", func(r chi.Router) {
			re := *setup.Articles(&r)

			data := web.ArticleRequest{
				User_Id: 0,
				Title:   "kedua",
				Content: "kedua",
			}
			dataJson,_ := json.Marshal(data)
			reader := bytes.NewReader(dataJson)

			request := httptest.NewRequest(http.MethodPut,"/articles/6477da1b-29cc-4e6a-812b-bfe4d13167b4",reader)
			request.Header.Add("Authorization", "Bearer token")
			recorder := httptest.NewRecorder()
			re.ServeHTTP(recorder,request)
			assert.Equal(t,http.StatusOK,recorder.Code)
		})
	})
}