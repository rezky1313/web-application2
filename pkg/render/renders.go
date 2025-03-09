package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// put render function here

// now when user acces the page the application render the file every single time
// those are read from disk, so now we gonna make it render just once
func RenderTemplate1(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing Template:", err)
		//panic(err)
	}

}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have the tempalte
	_, inMap := templateCache[t]
	if !inMap {
		//need create the template to read on disk and parsing the tempalte
		log.Println("creating template and adding to cache")
		err = createTempalteCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//we have the tempalte in cache
		log.Println("using cached template")
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTempalteCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add template to cache
	templateCache[t] = tmpl
	return nil
}

//https://chatgpt.com/g/g-FZNaQRBDg-golang-mentor/c/67cd4df9-839c-8005-b648-4fec1d3f06d4