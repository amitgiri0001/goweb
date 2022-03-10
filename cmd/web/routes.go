package main

import (
	"log"

	"github.com/amitgiri0001/goweb/pkg/config"
	"github.com/amitgiri0001/goweb/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	AppConfig *config.Appconfig
}

func (rt *Router) Routes() *chi.Mux {
	r := chi.NewRouter()

	repo := handlers.Repository{
		Logger:    log.Default(),
		AppConfig: rt.AppConfig,
	}
	handlers.InitHandlers(repo)

	r.Route("/users", func(r chi.Router) {
		r.With(RequestLogger).Get("/{user}", handlers.Repo.Home)
	})
	r.Get("/about", handlers.Repo.About)

	return r
}
