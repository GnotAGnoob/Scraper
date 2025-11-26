package shared

import "context"

type Ingredient struct {
	Name          string  `json:"name"`
	QuantityGrams float64 `json:"quantityGrams"`
}

type LLMIngredientResult struct {
	Ingredients []Ingredient `json:"ingredients"`
}

type Provider interface {
	TextToIngredients(ctx context.Context, text string) ([]Ingredient, error)
}
