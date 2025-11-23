package llm

import (
	"context"
	"errors"
	"strings"

	"github.com/teilomillet/gollm"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
)

const (
	ingredientsPrompt  = "Process the input: if it's a list of ingredients, extract them; if it's a dish name or recipe request (e.g., 'I want a burrito'), generate appropriate ingredients for that dish."
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

func TextToIngredients(ctx context.Context, cfg *config.Config, text string) ([]Ingredient, error) {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return nil, errors.New("text is empty")
	}

	prompt := gollm.NewPrompt(
		ingredientsPrompt,
		gollm.WithContext("Recipe:\n"+trimmed),
		gollm.WithDirectives(promptDirectives),
		gollm.WithOutput(promptOutputSchema),
		gollm.WithMaxLength(1000),
	)

	result, err := generateAndParse[llmIngredientResult](ctx, cfg, prompt)
	if err != nil {
		return nil, err
	}

	return result.Ingredients, nil
}
