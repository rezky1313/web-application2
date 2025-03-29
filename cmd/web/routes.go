package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rezky1313/web-application2/pkg/config"
	"github.com/rezky1313/web-application2/pkg/handlers"
)

/*
when start getting more complex web application its not good to have all
your routes right in your main folder/routine, so we gonna make separate files
wich i will call routes.go
*/

/*
Install a middleware
Middleware actualy allows you to process a request as it comes into your web
and perform an action on it

every time somebody hits the page just write something to the console
*/

func routes(app *config.AppConfig) http.Handler {
	//Router Using Chi

	mux := chi.NewRouter()

	//middleware taht write when panic
	mux.Use(middleware.Recoverer)

	//middleware using nosurf
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.HomeTemplate)
	mux.Get("/about", handlers.Repo.AboutTemplate)

	return mux

}
