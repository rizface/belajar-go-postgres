package test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {

	//var container []map[string]interface{}

	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	client,err := mongo.Connect(ctx,options.Client().ApplyURI("mongodb://root:root@172.26.0.2:27017"))
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println(err.Error())
		}
	}()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db := client.Database("go_blog")
	err = client.Ping(ctx,readpref.Primary())
	collection := db.Collection("student")
	res,err := collection.InsertOne(ctx, bson.D{{"nama","ip"},{"fariz",4}})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.InsertedID)

	data,err := collection.Find(ctx,bson.D{{"name","Fariz"}})
	defer data.Close(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	var container []map[string]interface{}
	data.All(ctx,&container)
	//for data.Next(ctx) {
	//	each := make(map[string]interface{})
	//	err := data.Decode(&each)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	container = append(container ,each)
	//}
	for _, v := range container {
		for strKey, element := range v {
			fmt.Println(strKey," -> ",element)
		}
	}
	data2,err2 := db.Collection("student").Find(ctx,bson.D{{"user_id",100}})
	if err2 != nil {
		fmt.Println(err2.Error(), "wkwk")
	} else {
		fmt.Println(data2.Next(ctx), "err")
	}
	fmt.Println(data2.Next(ctx), "ini data 2")


	res,_ = db.Collection("student").InsertOne(ctx,map[string]interface{}{
		"name":"Fariz",
	})
	fmt.Println(res.InsertedID, "wkwkw")

	update,err := db.Collection("student").UpdateMany(ctx,bson.D{{"nama", "ip"}},bson.M{"$inc": bson.M{"ip": 1000}})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(update.ModifiedCount)
	}
}

