package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, fmt.Sprint("Hello world"))
	})

	http.ListenAndServe(":8080", nil)
}
