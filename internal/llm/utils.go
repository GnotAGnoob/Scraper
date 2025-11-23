package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
	"github.com/teilomillet/gollm"
)

func generateAndParse[T any](ctx context.Context, cfg *config.Config, prompt *gollm.Prompt) (*T, error) {
	result := new(T)

	llmClient, err := gollm.NewLLM(
		gollm.SetProvider(cfg.LLMProvider),
		gollm.SetModel(cfg.LLMModel),
		gollm.SetAPIKey(cfg.LLMAPIKey),
	)
	if err != nil {
		return result, fmt.Errorf("init llm client: %w", err)
	}

	raw, err := llmClient.Generate(ctx, prompt)
	if err != nil {
		return result, fmt.Errorf("llm generate: %w", err)
	}

	cleaned := sanitizeJSON(raw)

	if err := json.Unmarshal([]byte(cleaned), result); err != nil {
		return result, fmt.Errorf("parse llm response: %w", err)
	}

	return result, nil
}

func sanitizeJSON(output string) string {
	cleaned := strings.TrimSpace(output)
	cleaned = strings.TrimPrefix(cleaned, "```json")
	cleaned = strings.TrimPrefix(cleaned, "```JSON")
	cleaned = strings.TrimPrefix(cleaned, "```")
	cleaned = strings.TrimSuffix(cleaned, "```")
	return strings.TrimSpace(cleaned)
}
