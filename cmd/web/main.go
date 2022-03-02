package main

import (
	"net/http"

	"github.com/amitgiri0001/goweb/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":8080", nil)
}
