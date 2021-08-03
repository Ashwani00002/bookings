package handlers

import (
	"net/http"

	"github.com/Ashwani00002/bookings/pkg/config"
	"github.com/Ashwani00002/bookings/pkg/models"
	"github.com/Ashwani00002/bookings/pkg/render"
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

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Test"] = "Hello, Again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
