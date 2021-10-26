package funfact_service

import "context"

type FunFactService interface{
	PostFunFact(ctx context.Context, request map[string]interface{}) interface{}
	PutFunFact(ctx context.Context, reqesut map[string]interface{}) bool
	DeleteFunFact(ctx context.Context, userId int) bool
}
