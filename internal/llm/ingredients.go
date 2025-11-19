package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Ingredient struct {
	Name        string  `json:"name"`
	AmountGrams float64 `json:"amount_grams"`
}

const ingredientPrompt = `
You are an expert chef and nutritionist.

Extract all ingredients from the following recipe text.
For each ingredient, estimate the total amount in GRAMS.

Return ONLY a JSON array (no markdown, no comments, no explanations) with objects:
[
  {
    "name": "ingredient name in Czech or English, singular form",
    "amount_grams": 123.4
  }
]

Rules:
- Use grams for everything (convert from cups, spoons, pieces, etc.).
- Merge duplicates (e.g. "cibule" and "cibuli" -> "cibule") and sum their grams.
- Ignore water and salt unless their quantities are clearly specified.
- If you are very unsure about the quantity, make a reasonable guess instead of 0.
`

func ExtractIngredients(ctx context.Context, recipeText string) ([]Ingredient, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("GEMINI_API_KEY environment variable is not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("creating gemini client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")

	prompt := ingredientPrompt + "\n\nRECIPE TEXT:\n" + recipeText

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("calling gemini: %w", err)
	}

	raw := responseToString(resp)
	jsonStr, err := extractJSON(raw)
	if err != nil {
		return nil, fmt.Errorf("extracting JSON from model output: %w; raw=%q", err, raw)
	}

	var ingredients []Ingredient
	if err := json.Unmarshal([]byte(jsonStr), &ingredients); err != nil {
		return nil, fmt.Errorf("unmarshalling ingredients JSON: %w; json=%s", err, jsonStr)
	}

	return ingredients, nil
}

func responseToString(resp *genai.GenerateContentResponse) string {
	if resp == nil {
		return ""
	}

	var buf bytes.Buffer
	for _, c := range resp.Candidates {
		if c.Content == nil {
			continue
		}
		for _, part := range c.Content.Parts {
			buf.WriteString(fmt.Sprint(part))
		}
	}
	return buf.String()
}

func extractJSON(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("empty response")
	}

	start := strings.IndexAny(s, "[{")
	if start == -1 {
		return "", errors.New("no JSON start ([ or {) found")
	}

	var end int
	switch s[start] {
	case '[':
		end = strings.LastIndex(s, "]")
	case '{':
		end = strings.LastIndex(s, "}")
	default:
		return "", errors.New("unexpected JSON start character")
	}

	if end == -1 || end <= start {
		return "", errors.New("no JSON end (] or }) found")
	}

	return s[start : end+1], nil
}
