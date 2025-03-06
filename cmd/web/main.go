package main

import (
	"fmt"
	"net/http"

	"github.com/rezky1313/web-application2/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.HomeTemplate)
	http.HandleFunc("/about", handlers.AboutTemplate)

	fmt.Println("Starting application on port :8080")
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}
