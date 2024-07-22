package handlers

import (
	"github.com/mziyadabek/bookings/pkg/config"
	"github.com/mziyadabek/bookings/pkg/models"
	"github.com/mziyadabek/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplatest(w, "Home.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	render.RenderTemplatest(w, "About.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
