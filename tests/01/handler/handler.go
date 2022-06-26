package handler

import (
	"io"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func RedirectHandler(paths map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if destination, ok := paths[path]; ok {
			http.Redirect(w, r, destination, http.StatusFound)
			return
		}

		HealthCheckHandler(w, r)
	}
}
