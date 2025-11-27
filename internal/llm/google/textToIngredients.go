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
	ingredientsPrompt = "Process the input: if it's a list of ingredients, extract them; if it's a dish name or recipe request (e.g., 'I want a burrito'), generate appropriate ingredients for that dish. The input might be in any language, translate it to Czech and put it in the 'name' field, the name in the original language should be put in the 'originalName' field. if the input is a recipe request, return the recipe in the 'recipe' field in the same language as the input. DO NOT INCLUDE THE INGREDIENTS LIST IN THE RECIPE TEXT FIELD. and had to be in markdown format."
	promptDirectives  = "Return ONLY valid JSON. No prose, no code fences. Convert units to grams. If you must guess, be conservative."
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
					"originalName": map[string]any{
						"type": "string",
					},
					"quantityGrams": map[string]any{
						"type": "number",
					},
				},
				"required": []string{"name", "originalName", "quantityGrams"},
			},
		},
		"recipe": map[string]any{
			"type": "string",
		},
	},
	"required": []string{"ingredients"},
}

func (p *Provider) TextToIngredients(ctx context.Context, text string) (*shared.LLMIngredientResult, error) {
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

	return result, nil
}
