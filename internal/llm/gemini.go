package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type RecipeIngredients struct {
	Ingredients []Ingredient `json:"ingredients"`
}

type GeminiClient struct {
	apiKey string
}

func NewGeminiClient(apiKey string) *GeminiClient {
	return &GeminiClient{apiKey: apiKey}
}

func (c *GeminiClient) ExtractIngredients(ctx context.Context, recipeText string) (*RecipeIngredients, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(c.apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")
	model.SetTemperature(0.2)
	model.ResponseMIMEType = "application/json"

	prompt := `You are a cooking assistant. Extract all ingredients from the following recipe text and return them in a structured JSON format.
For each ingredient, extract:
- name: The ingredient name (in Czech if the recipe is in Czech, otherwise in the original language)
- quantity: The amount as a number (convert to grams/ml where possible, use 0 if not specified)
- unit: The unit of measurement (g, ml, kg, l, ks/pieces, etc.)

Return ONLY valid JSON in this format:
{
  "ingredients": [
    {"name": "mouka", "quantity": 500, "unit": "g"},
    {"name": "mléko", "quantity": 250, "unit": "ml"}
  ]
}

Recipe text:
` + recipeText

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no response from gemini")
	}

	responseText := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	var result RecipeIngredients
	if err := json.Unmarshal([]byte(responseText), &result); err != nil {
		return nil, fmt.Errorf("failed to parse gemini response: %w", err)
	}

	return &result, nil
}
