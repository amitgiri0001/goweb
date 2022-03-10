package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := chi.URLParam(r, "user")
		log.Println("from the middleware", params)
		next.ServeHTTP(w, r)
	})
}
