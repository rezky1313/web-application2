package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//put all the handler function here

// Home page Handler
func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "homepage.tmpl")
}

// About Page handler
func AboutTemplate(w http.ResponseWriter, r *http.Request) {

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing Template:", err)
		//panic(err)
	}

}
