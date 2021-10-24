package user_service

import (
	"context"
	"go-blog/model/web"
)

type UserService interface {
	Login(ctx context.Context, request web.UserRequest) string
	Register(ctx context.Context, request web.UserRequest) bool
}