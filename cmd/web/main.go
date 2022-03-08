package main

import (
	"fmt"
	"net/http"

	"github.com/amitgiri0001/goweb/pkg/config"
)

func main() {
	appConfig := config.Appconfig{}

	appConfig.SetPort(8080)
	appConfig.SetTemplateDirectory("./templates/")

	fmt.Printf("%+v", appConfig.Port)

	router := Router{
		AppConfig: &appConfig,
	}

	srv := &http.Server{
		Addr:    fmt.Sprint(":", appConfig.Port),
		Handler: router.Routes(),
	}

	srv.ListenAndServe()
}
