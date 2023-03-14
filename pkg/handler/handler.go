package handler

import (
	"github.com/amir-moshfegh/web-game/pkg/config"
	"github.com/amir-moshfegh/web-game/pkg/models"
	"github.com/amir-moshfegh/web-game/pkg/render"
	"net/http"
)

type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_addr", remoteAddr)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["name"] = "Amir"
	stringMap["remote_addr"] = m.app.Session.GetString(r.Context(), "remote_addr")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
