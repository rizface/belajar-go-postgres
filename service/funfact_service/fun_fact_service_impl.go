package funfact_service

import (
	"context"
	"fmt"
	"go-blog/app/database"
	"go-blog/helper"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type funFactService struct{}

func NewService() FunFactService {
	return funFactService{}
}

func (f funFactService) PostFunFact(ctx context.Context, request map[string]interface{}) interface{} {
	cx,cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()
	db := database.Mongo.Database("go_blog")
	collection := db.Collection("funFact")
	result,err := collection.InsertOne(cx, request)
	helper.PanicIfError(err)
	fmt.Println(result.InsertedID, "ini dari service")
	return result.InsertedID
}

func (f funFactService) PutFunFact(ctx context.Context, request map[string]interface{}) bool {
	cx,cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()
	fmt.Println(request["user_id"])
	db := database.Mongo.Database("go_blog")
	collection,err := db.Collection("funFact").UpdateOne(cx, bson.M{"user_id": request["user_id"]},bson.M{"$set": request})
	helper.PanicIfError(err)
	return collection.ModifiedCount > 0
}

func (f funFactService) DeleteFunFact(ctx context.Context, userId int) bool {
	cx,cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()
	db := database.Mongo.Database("go_blog")
	collection := db.Collection("funFact")
	result,err := collection.DeleteOne(cx,bson.M{"user_id": userId})
	helper.PanicIfError(err)
	return result.DeletedCount > 0
}
