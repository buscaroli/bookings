package handlers

import (
	"net/http"

	"github.com/buscaroli/secondGoWebApp/pkg/config"
	"github.com/buscaroli/secondGoWebApp/pkg/models"
	"github.com/buscaroli/secondGoWebApp/pkg/render"
)

// repository used by the handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// handler method on the Repository struct
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// handler method on the Repository struct
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}
