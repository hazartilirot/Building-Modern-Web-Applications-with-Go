package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/internal/config"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/internal/forms"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/internal/models"
	"github.com/hazartilirot/Building-Modern-Web-Applications-with-Go/bookings/internal/render"
	"log"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

/*Reservations responsible for rendering room reservation*/
func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["reservations"] = models.Reservations{}
	render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}
func (m *Repository) PostReservations(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	reservations := models.Reservations{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	//form.HasField("first_name", r)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3, r)

	if !form.IsValid() {
		data := make(map[string]interface{})
		data["reservations"] = reservations

		render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
}
func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals-quarters.page.tmpl", &models.TemplateData{})
}
func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors-suite.page.tmpl", &models.TemplateData{})
}

/*RoomAvailability shows available rooms in selected dates*/
func (m *Repository) RoomAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "room-availability.page.tmpl", &models.TemplateData{})
}
func (m *Repository) FormRoomAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	w.Write([]byte(fmt.Sprintf("arrival day is %s, departure day is %s", start, end)))
}

type roomAvailabilityJson struct {
	OK  bool   `json:"ok"`
	Msg string `json:"msg"`
}

func (m *Repository) FormRoomAvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := roomAvailabilityJson{
		OK:  true,
		Msg: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
