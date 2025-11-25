package google

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"google.golang.org/genai"
)

type Provider struct {
	cfg *config.Config
}

func NewProvider(ctx context.Context, cfg *config.Config) *Provider {
	return &Provider{cfg: cfg}
}

func generateAndParse[T any](ctx context.Context, cfg *config.Config, systemInstruction, userContent string, responseJsonSchema map[string]any) (*T, error) {
	logger := applogger.LoggerFromContext(ctx)
	result := new(T)

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.LLMAPIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to initialize GenAI client")
		return result, fmt.Errorf("init genai client: %w", err)
	}

	systemInstructionContent := genai.NewContentFromText(systemInstruction, "")
	userContentContent := genai.NewContentFromText(userContent, genai.RoleUser)
	contents := []*genai.Content{userContentContent}

	config := &genai.GenerateContentConfig{
		SystemInstruction:  systemInstructionContent,
		ResponseMIMEType:   "application/json",
		ResponseJsonSchema: responseJsonSchema,
		MaxOutputTokens:    10000,
	}

	response, err := client.Models.GenerateContent(ctx, cfg.LLMModel, contents, config)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate GenAI response")
		return result, fmt.Errorf("genai generate: %w", err)
	}

	// Collect all text from all candidates and parts
	var raw strings.Builder
	for _, candidate := range response.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				if part.Text != "" {
					raw.WriteString(part.Text)
				}
			}
		}
	}

	rawStr := raw.String()
	if rawStr == "" {
		// Fallback to response.Text() if no candidates found
		rawStr = response.Text()
	}
	if rawStr == "" {
		logger.Error().Msg("empty response from GenAI")
		return result, fmt.Errorf("empty response from genai")
	}

	logger.Info().Str("raw", rawStr).Msg("raw response from GenAI")
	cleaned := shared.SanitizeJSON(rawStr)

	if err := json.Unmarshal([]byte(cleaned), result); err != nil {
		logger.Error().Err(err).Str("cleaned", cleaned).Msg("failed to parse GenAI response")
		return result, fmt.Errorf("parse genai response: %w", err)
	}

	return result, nil
}
