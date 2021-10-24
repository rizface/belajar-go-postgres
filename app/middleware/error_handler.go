package middleware

import (
	"fmt"
	"go-blog/app/exception"
	"go-blog/helper"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				if duplicate,duplicateOK := err.(exception.Duplicate);duplicateOK {
					field := helper.ExtractDuplicateField(duplicate.Err.Error())
					helper.JsonWriter(writer,duplicate.Code,field + " sudah digunakan", nil)
				} else if notFound,notFoundOK := err.(exception.NotFound); notFoundOK {
					helper.JsonWriter(writer,notFound.Code,notFound.Err.Error(),nil)
				} else if forbidden,forbiddenok := err.(exception.Forbidden); forbiddenok {
					helper.JsonWriter(writer,forbidden.Code,forbidden.Err.Error(),nil)
				}else {
					fmt.Println(err.(error).Error())
					helper.JsonWriter(writer,http.StatusInternalServerError, err.(error).Error(), nil)
				}
			}
		}()
		next.ServeHTTP(writer,request)
	})
}