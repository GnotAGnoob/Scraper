package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes sets up all HTTP routes for the API.
func RegisterRoutes(r chi.Router) {
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/recipe", handleRecipe)
	})
}

