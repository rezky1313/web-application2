package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// put render function here
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing Template:", err)
		//panic(err)
	}

}
