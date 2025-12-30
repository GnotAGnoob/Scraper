package shared

import "context"

type Ingredient struct {
	Name          string  `json:"name" validate:"required"`
	OriginalName  string  `json:"originalName" validate:"required"`
	QuantityGrams float64 `json:"quantityGrams" validate:"required"`
}

type LLMIngredientResult struct {
	Ingredients []Ingredient `json:"ingredients"`
	Recipe      *string      `json:"recipe"`
}

type Provider interface {
	TextToIngredients(ctx context.Context, text string) (*LLMIngredientResult, error)
}
