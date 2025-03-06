package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//put render function here
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing Template:", err)
		//panic(err)
	}

}
