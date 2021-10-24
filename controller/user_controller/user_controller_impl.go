package user_controller

import (
	"encoding/json"
	"go-blog/helper"
	"go-blog/model/web"
	"go-blog/service/user_service"
	"net/http"
)

type userController struct {
	service user_service.UserService
}

func NewController(service user_service.UserService) UserController {
	return userController{
		service: service,
	}
}

func (u userController) Login(w http.ResponseWriter, r *http.Request) {
	request := web.UserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)
	token := u.service.Login(r.Context(),request)
	helper.JsonWriter(w,http.StatusOK,"login success", map[string]interface{}{
		"token":token,
	})
}

func (u userController) Register(w http.ResponseWriter, r *http.Request) {
	request := web.UserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)
	success := u.service.Register(r.Context(),request)
	if success {
		helper.JsonWriter(w,http.StatusOK, "register success", nil)
	}
}

