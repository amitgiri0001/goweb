package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/amitgiri0001/goweb/internal/config"
)

func main() {
	appConfig := config.Appconfig{}

	appConfig.SetPort(8080)
	wd, _ := os.Getwd()
	appConfig.SetTemplateDirectory(wd + "/templates/")

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
