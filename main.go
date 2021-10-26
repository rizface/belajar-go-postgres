package main

import (
	"github.com/go-chi/chi/v5"
	"go-blog/app/database"
	"go-blog/app/setup"
	"log"
	"net/http"
)

func main() {
	r := setup.Router()
	database.Start()

	setup.Auth(r)
	r.Route("/articles", func(r chi.Router) {
		setup.Articles(&r)
	})

	r.Route("/authors", func(r chi.Router) {
		setup.Author(&r)
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
