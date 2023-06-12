package handlers

import (
	"net/http"

	"github.com/waterpigeon420/bookings/pkg/config"
	"github.com/waterpigeon420/bookings/pkg/models"
	"github.com/waterpigeon420/bookings/pkg/render"
)

//TenplateData holds dataset

type Repository struct {
	App *config.AppConfig
} //type repo

var Repo *Repository //pointer to repo

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a, //creates a new repo of app config
	}
}
func NewHandlers(r *Repository) {
	Repo = r //sets the repository for the handlers
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, bro!"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
