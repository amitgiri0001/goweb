package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amitgiri0001/goweb/pkg/config"
	"github.com/amitgiri0001/goweb/pkg/handlers"
)

func main() {
	appConfig := config.Appconfig{}

	appConfig.SetPort(8080)
	appConfig.SetTemplateDirectory("./templates/")

	repo := handlers.Repository{
		Logger:    log.Default(),
		AppConfig: &appConfig,
	}
	handlers.InitHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("%+v", appConfig.Port)

	http.ListenAndServe(fmt.Sprint(":", appConfig.Port), nil)
}
