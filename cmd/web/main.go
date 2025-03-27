package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rezky1313/web-application2/pkg/config"
	"github.com/rezky1313/web-application2/pkg/handlers"
	"github.com/rezky1313/web-application2/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Panicln("Cannot create tempalte cache")
	}

	log.Println("creating new template from main file")
	
	app.TemplateCache = tc
	app.UseCache = false

	log.Println("put the new template that just created into the template cache")

	repo:= handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.HomeTemplate)
	// http.HandleFunc("/about", handlers.Repo.AboutTemplate)

	fmt.Println("Starting application on port :8080")
	// _ = http.ListenAndServe(portNumber, nil)

	server := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}

}
