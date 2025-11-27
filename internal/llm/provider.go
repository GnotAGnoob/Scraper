package llm

import (
	"context"
	"fmt"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/gollm"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/google"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
)

func NewProvider(ctx context.Context, cfg *config.Config) (shared.Provider, error) {
	switch cfg.LLMProvider {
	case "google":
		return google.NewProvider(ctx, cfg), nil
	default:
		return gollm.NewProvider(ctx, cfg), nil
	}
}

func TextToIngredients(ctx context.Context, cfg *config.Config, text string) (*shared.LLMIngredientResult, error) {
	logger := applogger.LoggerFromContext(ctx)
	provider, err := NewProvider(ctx, cfg)
	if err != nil || provider == nil {
		logger.Error().Err(err).Msg("failed to create LLM provider")
		return nil, fmt.Errorf("failed to create LLM provider: %w", err)
	}

	return provider.TextToIngredients(ctx, text)
}
