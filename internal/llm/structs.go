package llm

type Ingredient struct {
	Name          string  `json:"name"`
	QuantityGrams float64 `json:"quantityGrams"`
	OriginalText  string  `json:"originalText"`
}

type llmIngredientResult struct {
	Ingredients []Ingredient `json:"ingredients"`
}
