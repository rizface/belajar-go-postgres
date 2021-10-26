package funcfact_controller

import "net/http"

type FunFactController interface {
	PostFuncFact(w http.ResponseWriter, r *http.Request)
	PutFunFact(w http.ResponseWriter, r *http.Request)
	DeleteFunFact(w http.ResponseWriter, r *http.Request)
}
