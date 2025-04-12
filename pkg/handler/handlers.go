package handler

import (
	"net/http"

	"github.com/Leopoldo013/bookings/pkg/config"
	"github.com/Leopoldo013/bookings/pkg/models"

	"github.com/Leopoldo013/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{app}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Valeu Leopold da silva"

	render.RenderTemplates(w, "home.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Leopoldão da silva"

	render.RenderTemplates(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
