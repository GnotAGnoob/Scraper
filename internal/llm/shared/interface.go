package shared

import "context"

type Ingredient struct {
	Name          string  `json:"name"`
	OriginalName  string  `json:"originalName"`
	QuantityGrams float64 `json:"quantityGrams"`
}

type LLMIngredientResult struct {
	Ingredients []Ingredient `json:"ingredients"`
	Recipe      *string      `json:"recipe"`
}

type Provider interface {
	TextToIngredients(ctx context.Context, text string) (*LLMIngredientResult, error)
}
