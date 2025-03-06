package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHelloWorldApp(t *testing.T) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world!")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Number of bytes written: %d", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
