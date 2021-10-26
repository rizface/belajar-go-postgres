package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
	_"github.com/lib/pq"
)

var Db *sqlx.DB
var Err error
var Mongo *mongo.Client

func Start() {
	log.Println("init database started")
}

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}
	log.Print("env file loaded")
}

func initPostgres() {
	var dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",os.Getenv("DATABASE_HOST"),os.Getenv("DATABASE_PORT"),os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), "go_blog")
	Db, Err = sqlx.Open("postgres", dsn)
	if Err != nil {
		panic(Err)
	}
	initPostgresTable()
}

func initPostgresTable() {
	schema := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	CREATE TABLE IF NOT EXISTS users(
		id BIGSERIAL NOT NULL PRIMARY KEY,
		username varchar(100) NOT NULL,
		email varchar(100) NOT NULL UNIQUE,
		password varchar(250) NOT NULL UNIQUE
	);
	
	CREATE TABLE IF NOT EXISTS contents(
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
		user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		title varchar(100) NOT NULL,
		content text,
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	Db.MustExec(schema)
	log.Println("init database finished")
}

func initMongo() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",os.Getenv("MONGO_USERNAME"),os.Getenv("MONGO_PASSWORD"),os.Getenv("MONGO_HOST"),os.Getenv("MONGO_PORT"))
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	Mongo, Err = mongo.Connect(ctx,options.Client().ApplyURI(uri))
	if Err != nil {
		panic(Err)
	}
}

func init() {
	initEnv()
	go initPostgres()
	go initMongo()
}


