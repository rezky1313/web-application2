package handlers

import (
	"net/http"

	"github.com/rezky1313/web-application2/pkg/config"
	"github.com/rezky1313/web-application2/pkg/render"
)

//put all the handler function here

// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo Creates a new repository
func NewRepo(a *config.AppConfig) *Repository  {
	return &Repository {
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository)  {
	Repo = r
}

// Home page Handler
func (m *Repository) HomeTemplate(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About Page handler
func (m *Repository) AboutTemplate(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
