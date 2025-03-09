package handlers

import (
	"net/http"

	"github.com/rezky1313/web-application2/pkg/render"
)

//put all the handler function here

// Home page Handler
func HomeTemplate(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About Page handler
func AboutTemplate(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
