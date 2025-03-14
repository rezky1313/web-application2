package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// put render function here

// now when user acces the page the application render the file every single time
// those are read from disk, so now we gonna make it render just once

// creating more complex tempalte cache

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a tempalte cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested tempalte from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the file named *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all file ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, err

}
