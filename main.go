package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-blog/app/setup"
	"log"
	"net/http"
)

func init() {
	schema := `
	CREATE EXTENSION "uuid-ossp";
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

	setup.Db.MustExec(schema)
	fmt.Println("Scheme Executed")
}

func main() {
	r := setup.Router()

	setup.Auth(r)
	r.Route("/articles", func(r chi.Router) {
		setup.Articles(&r)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
