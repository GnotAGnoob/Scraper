package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GnotAGnoob/kosik-scraper/internal/api"
	_ "github.com/GnotAGnoob/kosik-scraper/docs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Recipe to Kosik Cart API
// @version 1.0
// @description API for converting recipe text to Kosik shopping cart
// @host localhost:8080
// @BasePath /
func main() {
	// Setup logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Get Gemini API key from environment
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	if geminiApiKey == "" {
		log.Fatal().Msg("GEMINI_API_KEY environment variable is required")
	}

	// Initialize handler
	handler := api.NewHandler(geminiApiKey)

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "http://127.0.0.1:*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Routes
	r.Get("/health", handler.HealthCheck)
	r.Post("/api/v1/recipe", handler.ProcessRecipe)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Info().Msgf("Starting server on %s", addr)
	log.Info().Msg("Swagger UI available at http://localhost:" + port + "/swagger/index.html")

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

