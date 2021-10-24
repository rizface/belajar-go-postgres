package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-blog/app/exception"
	"go-blog/helper"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		header := request.Header.Get("Authorization")
		if(!strings.Contains(header,"Bearer")) {
			fmt.Println("error disini")
			exception.PanicForbidden(errors.New("token is invalid"))
		} else {
			items := strings.Split(header, " ")
			if len(items) != 2 {
				fmt.Println("disini")
				exception.PanicForbidden(errors.New("token is invalid"))
			} else {
				claim,err := helper.VerifyToken(items[1])
				exception.PanicForbidden(err)
				request := request.WithContext(context.WithValue(request.Context(),"user-data",claim))
				next.ServeHTTP(writer,request)
			}
		}
	})
}
