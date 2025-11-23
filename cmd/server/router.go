package main

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	healthhandler "github.com/GnotAGnoob/kosik-scraper/cmd/server/handlers/health"
	ingredientshandler "github.com/GnotAGnoob/kosik-scraper/cmd/server/handlers/ingredients"
	swaggerhandler "github.com/GnotAGnoob/kosik-scraper/cmd/server/handlers/swagger"
	servermiddleware "github.com/GnotAGnoob/kosik-scraper/cmd/server/middleware"
)

func (s *server) buildRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(servermiddleware.RequestLogger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   s.cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/health", healthhandler.HealthHandler(s.cfg))
	r.Route("/api/v1", func(api chi.Router) {
		api.Post("/text/products", ingredientshandler.TextProductsHandler(s.orchestrator))
	})

	r.Get("/swagger/*", swaggerhandler.SwaggerHandler())

	return r
}
