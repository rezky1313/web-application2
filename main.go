package main

import (
	"net/http"
	"testing"
)

const portNumber = ":8080"

func TestFunctionHandlerTemplate(t *testing.T) {
	http.HandleFunc("/", HomeTemplate)
	http.HandleFunc("/about", AboutTemplate)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		panic(err)
	}
}
