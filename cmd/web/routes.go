package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rezky1313/web-application2/pkg/config"
	"github.com/rezky1313/web-application2/pkg/handlers"
)

/*
when start getting more complex web application its not good to have all
your routes right in your main folder/routine, so we gonna make separate files
wich i will call routes.go
*/



func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.HomeTemplate))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutTemplate))

	//Router Using Chi

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.HomeTemplate)
	mux.Get("/about", handlers.Repo.AboutTemplate)

	return mux

}
