package gollm

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/teilomillet/gollm"
)

const (
	ingredientsPrompt  = "Process the input: if it's a list of ingredients, extract them; if it's a dish name or recipe request (e.g., 'I want a burrito'), generate appropriate ingredients for that dish. the input might be in any language, translate it to Czech."
	promptDirectives   = "Return ONLY valid JSON. No prose, no code fences. Convert units to grams. If you must guess, be conservative. Include the original ingredient wording in `originalText`."
	promptOutputSchema = `JSON schema:
{
  "ingredients": [
    {
      "name": "string",
      "quantityGrams": number,
      "originalText": "string"
    }
  ]
}`
)

func (p *Provider) TextToIngredients(ctx context.Context, text string) ([]shared.Ingredient, error) {
	logger := applogger.LoggerFromContext(ctx)
	trimmedText := strings.TrimSpace(text)
	if trimmedText == "" {
		logger.Error().Msg("text is empty")
		return nil, errors.New("text is empty")
	}

	prompt := gollm.NewPrompt(
		ingredientsPrompt,
		gollm.WithContext(trimmedText),
		gollm.WithDirectives(promptDirectives),
		gollm.WithOutput(promptOutputSchema),
		gollm.WithMaxLength(1000),
	)

	result, err := generateAndParse[shared.LLMIngredientResult](ctx, p.cfg, prompt)
	if err != nil {
		logger.Error().Err(err).Msg("failed to generate LLM response")
		return nil, fmt.Errorf("failed to generate LLM response: %w", err)
	}

	return result.Ingredients, nil
}
