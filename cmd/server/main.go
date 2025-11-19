package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/GnotAGnoob/kosik-scraper/docs"
	"github.com/GnotAGnoob/kosik-scraper/internal/api"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Recipe to Kosik API
// @version         0.1
// @description     Convert free-form recipes into structured Kosik product suggestions.
// @BasePath        /api/v1
func main() {
	logLevel := getEnv("LOG_LEVEL", "info")
	logger.Init(logLevel)

	apiKey := os.Getenv("GOOGLE_GEMINI_API_KEY")
	geminiClient, err := llm.NewClient(apiKey)
	if err != nil {
		panic(err)
	}

	workerLimit := parseIntEnv("KOSIK_WORKERS", 4)

	recipeService := service.NewRecipeService(
		geminiClient,
		service.NewKosikFetcher(),
		service.WithMaxConcurrency(workerLimit),
	)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	recipeHandler := api.NewRecipeHandler(recipeService)
	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/recipe", recipeHandler.HandleRecipe)
	})

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	port := getEnv("PORT", "8080")
	addr := fmt.Sprintf(":%s", port)

	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Printf("Server listening on %s\n", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func parseIntEnv(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
