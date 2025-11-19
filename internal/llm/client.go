package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	genaiClient *genai.Client
	model       *genai.GenerativeModel
}

func NewClient(ctx context.Context, apiKey string) (*Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client: %w", err)
	}

	model := client.GenerativeModel("gemini-2.5-flash")
	model.ResponseMIMEType = "application/json"

	return &Client{
		genaiClient: client,
		model:       model,
	}, nil
}

func (c *Client) Close() {
	c.genaiClient.Close()
}

func (c *Client) ExtractIngredients(ctx context.Context, text string) ([]Ingredient, error) {
	prompt := fmt.Sprintf(`
You are a cooking assistant. Extract a list of ingredients from the following recipe text.
For each ingredient, estimate the quantity in grams. If the unit is "pcs" or count, estimate the weight of that item in grams (e.g., 1 onion ~ 150g).
Return ONLY a valid JSON object with the following structure:
{
  "ingredients": [
    { "name": "ingredient name", "quantity_grams": 100, "original_unit": "cup" }
  ]
}

Recipe Text:
%s
`, text)

	resp, err := c.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no content generated")
	}

	part, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	jsonStr := string(part)
	// Clean up any potential markdown code blocks if the model ignores the mime type instruction (safeguard)
	jsonStr = strings.TrimPrefix(jsonStr, "```json")
	jsonStr = strings.TrimPrefix(jsonStr, "```")
	jsonStr = strings.TrimSuffix(jsonStr, "```")

	var extraction RecipeExtraction
	if err := json.Unmarshal([]byte(jsonStr), &extraction); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w, response: %s", err, jsonStr)
	}

	return extraction.Ingredients, nil
}
