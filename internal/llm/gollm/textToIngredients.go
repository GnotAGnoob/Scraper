package gollm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/teilomillet/gollm"
)

const (
	ingredientsPrompt  = "Process the input: if it's a list of ingredients, extract them; if it's a dish name or recipe request (e.g., 'I want a burrito'), generate appropriate ingredients for that dish. the input might be in any language, output should be translates into Czech. and return the ingredients in JSON format."
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
		gollm.WithMaxLength(10000),
	)

	result, err := generateAndParse[shared.LLMIngredientResult](ctx, p.cfg, prompt)
	if err != nil {
		// add the result to the error
		resultStr := "<nil>"
		if result != nil {
			if resultJSON, marshalErr := json.Marshal(result); marshalErr == nil {
				resultStr = string(resultJSON)
			} else {
				resultStr = fmt.Sprintf("<marshal error: %v>", marshalErr)
			}
		}
		logger.Error().Err(err).Str("result", resultStr).Msg("failed to generate LLM response")
		return nil, fmt.Errorf("failed to generate LLM response: %w", err)
	}

	return result.Ingredients, nil
}
