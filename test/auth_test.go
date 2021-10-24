package test

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"go-blog/app"
	"go-blog/app/setup"
	"go-blog/model/web"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Run("success", func(t *testing.T) {
		data := web.UserRequest{
			Email:    "malfarizzi13@gmail.com",
			Username: "fariz",
			Password: "rahasia",
		}
		jsonData,_ := json.Marshal(data)
		reader := strings.NewReader(string(jsonData))
		request := httptest.NewRequest(http.MethodPost,"http://localhost:80820/register",reader)
		recorder := httptest.NewRecorder()

		r := setup.Router()
		router := setup.Auth(r)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("duplicate", func(t *testing.T) {
		data := web.UserRequest{
			Email:    "malfarizzi13@gmail.com",
			Username: "fariz",
			Password: "rahasia",
		}
		jsonData,_ := json.Marshal(data)
		reader := strings.NewReader(string(jsonData))
		request := httptest.NewRequest(http.MethodPost,app.REGISTER,reader)
		recorder := httptest.NewRecorder()

		r := setup.Router()
		router := setup.Auth(r)
		router.ServeHTTP(recorder,request)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusConflict,recorder.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		data := web.UserRequest{
			Email:    "malfarizzi13@gmail.com",
			Password: "rahasia",
		}
		jsonData,_ := json.Marshal(data)
		reader := strings.NewReader(string(jsonData))
		request := httptest.NewRequest(http.MethodPost,app.LOGIN,reader)
		recorder := httptest.NewRecorder()

		r := setup.Router()
		router := setup.Auth(r)
		router.ServeHTTP(recorder,request)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("failed", func(t *testing.T) {
		data := web.UserRequest{
			Email:    "malfarizzi13@gmail.com",
			Password: "rahasia1",
		}
		jsonData,_ := json.Marshal(data)
		reader := strings.NewReader(string(jsonData))
		request := httptest.NewRequest(http.MethodPost,app.LOGIN,reader)
		recorder := httptest.NewRecorder()

		r := setup.Router()
		router := setup.Auth(r)
		router.ServeHTTP(recorder,request)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusForbidden,recorder.Code)
	})
	t.Run("not found", func(t *testing.T) {
		data := web.UserRequest{
			Email:    "malfarizzi123@gmail.com",
			Password: "rahasia1",
		}
		jsonData,_ := json.Marshal(data)
		reader := strings.NewReader(string(jsonData))
		request := httptest.NewRequest(http.MethodPost,app.LOGIN,reader)
		recorder := httptest.NewRecorder()

		r := setup.Router()
		router := setup.Auth(r)
		router.ServeHTTP(recorder,request)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusNotFound,recorder.Code)
	})
}