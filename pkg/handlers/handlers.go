package handlers

import (
	"net/http"

	"github.com/amitgiri0001/goweb/pkg/render"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := render.TemplateParser("./templates/home.page.tmpl")
	if err != nil {
		rw.Write([]byte("Something went wrong"))
		return
	}
	parsedTemplate.Execute(rw, nil)
}

func About(rw http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := render.TemplateParser("./templates/about.page.tmpl")
	if err != nil {
		rw.Write([]byte("Something went wrong"))
		return
	}
	parsedTemplate.Execute(rw, nil)
}
