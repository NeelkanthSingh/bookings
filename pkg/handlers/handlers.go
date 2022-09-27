package handlers

import (
	"github.com/neelkanthsingh/bookings/pkg/config"
	"github.com/neelkanthsingh/bookings/pkg/models"
	"github.com/neelkanthsingh/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home function return response for Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: map[string]string{
			"Name": "Neelkanth",
		},
	})
}

// About function return response for About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	ipAddress := m.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: map[string]string{
			"remote_ip": ipAddress,
		},
	})
}
