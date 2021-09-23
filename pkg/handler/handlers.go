package handler

import (
	"booking/pkg/config"
	"booking/pkg/models"
	"booking/pkg/render"
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

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//perform logic data
	StringMap := make(map[string]string)
	StringMap["test"] = "Ah shit here we go again!"

	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")
	StringMap["remoteIp"] = remoteIp

	//send data to the template
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: StringMap,
	})

}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIp", remoteIp)
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}
