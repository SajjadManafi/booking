package handler

import (
	"booking/internal/config"
	"booking/internal/models"
	"booking/internal/render"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// About renders the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform logic data
	StringMap := make(map[string]string)
	StringMap["test"] = "Ah shit here we go again!"

	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")
	StringMap["remoteIp"] = remoteIp

	//send data to the template
	render.RenderTemplate(w, r, "about.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})

}

// Home renders the Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIp", remoteIp)
	render.RenderTemplate(w, r, "home.page.gohtml", &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.gohtml", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.gohtml", &models.TemplateData{})
}

// Book renders the Book page
func (m *Repository) Book(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "book.page.gohtml", &models.TemplateData{})
}

// Contact renders the Contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.gohtml", &models.TemplateData{})
}

// Availability renders the Availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.gohtml", &models.TemplateData{})
}

// Reservation renders the Reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.gohtml", &models.TemplateData{})
}

// PostAvailability handles request for Availability
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start-date")
	end := r.Form.Get("end-date")
	w.Write([]byte(fmt.Sprintf("Start date is %s & end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for Availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	Resp := jsonResponse{
		OK: true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(Resp, "" , "     ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(out)
}
