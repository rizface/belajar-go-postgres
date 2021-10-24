package helper

import (
	"errors"
	"go-blog/app/exception"
	"golang.org/x/crypto/bcrypt"
)

func Generate(plain string) ([]byte,error) {
	result,err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return result,err
}

func Compare(plain string, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(plain))
	if err != nil {
		exception.PanicForbidden(errors.New("username / password salah"))
	}
}