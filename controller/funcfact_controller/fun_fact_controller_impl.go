package funcfact_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-blog/app/exception"
	"go-blog/helper"
	"go-blog/service/funfact_service"
	"net/http"
)

type funFactController struct {
	service funfact_service.FunFactService
}

func NewController(service funfact_service.FunFactService) FunFactController {
	return funFactController{
		service: service,
	}
}

func (f funFactController) PostFuncFact(w http.ResponseWriter, r *http.Request) {
	var request = make(map[string]interface{})
	claim := r.Context().Value("user-data")

	userData,userDataOK := claim.(*helper.UserClaims)

	if userDataOK == false {
		helper.PanicIfError(errors.New("token is invalid"))
	}
	request["user_id"] = userData.Id
	fmt.Println(request)

	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)
	result := f.service.PostFunFact(r.Context(),request)
	fmt.Println(result,"ini dari controller")
	if result != nil {
		helper.JsonWriter(w,http.StatusOK, "fun fact posted",nil)
	}
}

func (f funFactController) PutFunFact(w http.ResponseWriter, r *http.Request) {
	var request = make(map[string]interface{})
	claim := r.Context().Value("user-data")
	dataUser,dataUserok := claim.(*helper.UserClaims)
	if dataUserok == false {
		exception.PanicForbidden(errors.New("token is invalid"))
	}
	request["user_id"] = dataUser.Id
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)
	success := f.service.PutFunFact(r.Context(),request)
	fmt.Println(success)
	if success {
		helper.JsonWriter(w,http.StatusOK,"fun fact updated",nil)
	}
}

func (f funFactController) DeleteFunFact(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value("user-data")
	dataUser,dataUserok := claim.(*helper.UserClaims)
	if dataUserok == false {
		exception.PanicForbidden(errors.New("token is invalid"))
	}
	success := f.service.DeleteFunFact(r.Context(),dataUser.Id)
	if success {
		helper.JsonWriter(w,http.StatusOK, "fun fact deleted", nil)
	}
}


