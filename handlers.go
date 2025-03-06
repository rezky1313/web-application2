package main

import (
	"net/http"
)

//put all the handler function here

// Home page Handler
func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "homepage.tmpl")
}

// About Page handler
func AboutTemplate(w http.ResponseWriter, r *http.Request) {

}
