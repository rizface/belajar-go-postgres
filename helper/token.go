package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-blog/app/exception"
	"go-blog/model/domain"
	"time"
)

var secret = []byte("ini rahasia")

type UserClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user domain.Users) string{
	claim := UserClaims{
		Id:               user.Id,
		Username:         user.Username,
		Email:            user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "Blog Postgres",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	token,err := jwt.SignedString(secret)
	PanicIfError(err)
	return token
}

func VerifyToken(plain string)(interface{},error) {
	token,err := jwt.ParseWithClaims(plain,&UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_,ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok {
			return secret,nil
		}
		return nil,errors.New("token is invalid")
	})

	if err != nil {
		exception.PanicForbidden(err)
	}

	claim,claimOK := token.Claims.(*UserClaims)
	if claimOK && token.Valid {
		return claim,nil
	}
	return nil,errors.New("token is invalid")
}