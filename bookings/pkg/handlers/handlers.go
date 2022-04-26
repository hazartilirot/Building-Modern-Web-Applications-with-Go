package handlers

import (
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/pkg/config"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/pkg/models"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

/*Reservations responsible for rendering room reservation*/
func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservations.page.tmpl", &models.TemplateData{})
}
func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals-quarters.page.tmpl", &models.TemplateData{})
}
func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors-suite.page.tmpl", &models.TemplateData{})
}

/*RoomAvailability shows available rooms in selected dates*/
func (m *Repository) RoomAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "room-availability.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
