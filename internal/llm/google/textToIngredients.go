package google

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
)

const (
	ingredientsPrompt = "Process the input: if it's a list of ingredients, extract them; if it's a dish name or recipe request (e.g., 'I want a burrito'), generate appropriate ingredients for that dish. the input might be in any language, translate it to Czech."
	promptDirectives  = "Return ONLY valid JSON. No prose, no code fences. Convert units to grams. If you must guess, be conservative. Include the original ingredient wording in `originalText`."
)

var responseJsonSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"ingredients": map[string]any{
			"type": "array",
			"items": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name": map[string]any{
						"type": "string",
					},
					"quantityGrams": map[string]any{
						"type": "number",
					},
					"originalText": map[string]any{
						"type": "string",
					},
				},
				"required": []string{"name", "quantityGrams", "originalText"},
			},
		},
	},
	"required": []string{"ingredients"},
}

func (p *Provider) TextToIngredients(ctx context.Context, text string) ([]shared.Ingredient, error) {
	logger := applogger.LoggerFromContext(ctx)
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		logger.Error().Msg("text is empty")
		return nil, errors.New("text is empty")
	}

	systemInstruction := fmt.Sprintf("%s\n\n%s", ingredientsPrompt, promptDirectives)
	userContent := fmt.Sprintf("Recipe:\n%s", trimmed)

	result, err := generateAndParse[shared.LLMIngredientResult](ctx, p.cfg, systemInstruction, userContent, responseJsonSchema)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate GenAI response")
		return nil, fmt.Errorf("failed to generate GenAI response: %w", err)
	}

	return result.Ingredients, nil
}
