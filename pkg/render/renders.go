package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/rezky1313/web-application2/pkg/config"
	"github.com/rezky1313/web-application2/pkg/models"
)

// setting aplication wide configuration
// in section before we stil load all the template set when request has been made
// now once i have tempalte set i never want to load it again until application restart
// using app config

var function = template.FuncMap{}

var app *config.AppConfig

// newtemplates sets the config for the tempalte package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//for development mode use the usecache for rebuild
	// if not just tc := app.TemplateCache
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config

		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested tempalte from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get tempalte from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	log.Println("get template from cache")

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
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
