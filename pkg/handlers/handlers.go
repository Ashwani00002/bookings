package handlers

import (
	"net/http"

	"github.com/Ashwani00002/go-course/pkg/config"
	"github.com/Ashwani00002/go-course/pkg/render"
)

// Repo is repository used by handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	App *config.Appconfig
}

// NewRepo create a new Repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the Repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
