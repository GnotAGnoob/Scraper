package main

import (
	"context"
	"net/http"
	"os"

	"github.com/GnotAGnoob/kosik-scraper/internal/api"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/service"
	"github.com/rs/zerolog/log"
)

// @title Recipe to Kosik Cart API
// @version 1.0
// @description API for converting recipes to Kosik cart items.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	logger.Init(os.Getenv("LOG_LEVEL"))

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal().Msg("GEMINI_API_KEY environment variable is required")
	}

	ctx := context.Background()
	llmClient, err := llm.NewClient(ctx, apiKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize LLM client")
	}
	defer llmClient.Close()

	matcher := service.NewIngredientMatcher()
	handler := api.NewHandler(llmClient, matcher)

	r := api.NewRouter(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info().Msgf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
