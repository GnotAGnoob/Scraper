package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type IngredientExtraction struct {
	Ingredients []Ingredient `json:"ingredients"`
}

type Client struct {
	client   *genai.Client
	model    *genai.GenerativeModel
	modelName string
}

func NewClient() (*Client, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable is not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	// Allow model to be configured via env var, default to gemini-pro (most stable)
	modelName := os.Getenv("GEMINI_MODEL")
	if modelName == "" {
		modelName = "gemini-pro"
	}

	model := client.GenerativeModel(modelName)
	log.Info().Str("model", modelName).Msg("Using Gemini model")

	return &Client{
		client:    client,
		model:     model,
		modelName: modelName,
	}, nil
}

func (c *Client) ExtractIngredients(ctx context.Context, recipeText string) ([]Ingredient, error) {
	prompt := `Extract ingredients and their quantities from the following recipe text. 
Return ONLY a valid JSON object with this exact structure:
{
  "ingredients": [
    {
      "name": "ingredient name",
      "quantity": 100.0,
      "unit": "g"
    }
  ]
}

Convert all quantities to grams (g). If no quantity is specified, use 100g as default.
If the unit is not weight-based (e.g., "pieces", "cups"), convert to approximate grams.
Return ONLY the JSON, no additional text or markdown formatting.

Recipe text:
` + recipeText

	resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		// If model not found, suggest alternatives
		if err.Error() != "" {
			return nil, fmt.Errorf("failed to generate content with model %s: %w. Try setting GEMINI_MODEL env var to: gemini-pro, gemini-1.5-flash-001, or gemini-1.5-pro-001", c.modelName, err)
		}
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no response from Gemini")
	}

	text := ""
	for _, part := range resp.Candidates[0].Content.Parts {
		if str, ok := part.(genai.Text); ok {
			text += string(str)
		}
	}

	log.Debug().Str("response", text).Msg("Gemini raw response")

	var extraction IngredientExtraction
	if err := json.Unmarshal([]byte(text), &extraction); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return extraction.Ingredients, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

