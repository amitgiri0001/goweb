package handlers

import (
	"net/http"

	"github.com/amitgiri0001/goweb/pkg/render"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.TemplateParser("./templates/home.page.tmpl")
	parsedTemplate.Execute(rw, nil)
}

func About(rw http.ResponseWriter, r *http.Request) {
	parsedTemplate := render.TemplateParser("./templates/about.page.tmpl")
	parsedTemplate.Execute(rw, nil)
}
