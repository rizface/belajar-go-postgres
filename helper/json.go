package helper

import (
	"encoding/json"
	"go-blog/model/response"
	"net/http"
)

func JsonWriter(writer http.ResponseWriter,code int, status string ,data interface{}){
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(response.Response{
		Code:   code,
		Status: status,
		Data:   data,
	})
}
