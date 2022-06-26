package main

import (
	"net/http"

	"github.com/ismael3s/go-tests/01/handler"
)

func main() {
	paths := map[string]string{
		"/google": "https://www.google.com",
	}

	http.HandleFunc("/health-check", handler.HealthCheckHandler)
	http.HandleFunc("/", handler.RedirectHandler(paths))
	http.ListenAndServe(":8080", nil)
}
