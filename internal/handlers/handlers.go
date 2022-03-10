package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amitgiri0001/goweb/internal/config"
	"github.com/amitgiri0001/goweb/internal/render"
	"github.com/go-chi/chi/v5"
)

type Repository struct {
	AppConfig      *config.Appconfig
	databaseConfig map[string]string
	Logger         *log.Logger
}

var Repo *Repository

func InitHandlers(r Repository) {
	Repo = &Repository{
		Logger:    r.Logger,
		AppConfig: r.AppConfig,
		// Other components.
	}

	render.InitRenderers(Repo.AppConfig.TemplateDirectory)
}

func (r *Repository) Home(rw http.ResponseWriter, hr *http.Request) {
	parsedTemplate := render.Render.TemplateParser("home.page.tmpl")
	data := []string{"Hello"}
	fmt.Println(chi.URLParam(hr, "user"))
	parsedTemplate.Execute(rw, data[0])
	r.Logger.Println("Home executes")
}

func (r *Repository) About(rw http.ResponseWriter, _ *http.Request) {
	parsedTemplate := render.Render.TemplateParser("about.page.tmpl")
	parsedTemplate.Execute(rw, nil)
	r.Logger.Println("about executes")
}
