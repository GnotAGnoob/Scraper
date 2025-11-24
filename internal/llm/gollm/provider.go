package gollm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/teilomillet/gollm"
)

type Provider struct {
	cfg *config.Config
}

func NewProvider(_ context.Context, cfg *config.Config) *Provider {
	return &Provider{cfg: cfg}
}

func generateAndParse[T any](ctx context.Context, cfg *config.Config, prompt *gollm.Prompt) (*T, error) {
	logger := applogger.LoggerFromContext(ctx)
	result := new(T)

	llmClient, err := gollm.NewLLM(
		gollm.SetProvider(cfg.LLMProvider),
		gollm.SetModel(cfg.LLMModel),
		gollm.SetAPIKey(cfg.LLMAPIKey),
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to initialize LLM client")
		return result, fmt.Errorf("init llm client: %w", err)
	}

	raw, err := llmClient.Generate(ctx, prompt)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate LLM response")
		return result, fmt.Errorf("llm generate: %w", err)
	}

	cleaned := shared.SanitizeJSON(raw)

	if err := json.Unmarshal([]byte(cleaned), result); err != nil {
		logger.Error().Err(err).Msg("failed to parse LLM response")
		return result, fmt.Errorf("parse llm response: %w", err)
	}

	return result, nil
}
