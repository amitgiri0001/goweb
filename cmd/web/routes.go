package main

import (
	"log"

	"github.com/amitgiri0001/goweb/pkg/config"
	"github.com/amitgiri0001/goweb/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r.Use(middleware.Logger)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)

	return r
}
