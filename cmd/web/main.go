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

	app.TemplateCache = tc
	
	http.HandleFunc("/", handlers.HomeTemplate)
	http.HandleFunc("/about", handlers.AboutTemplate)

	fmt.Println("Starting application on port :8080")
	_ = http.ListenAndServe(portNumber, nil)

}
