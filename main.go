package main

import (
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", HomeTemplate)
	http.HandleFunc("/about", AboutTemplate)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}
