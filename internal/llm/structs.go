package llm

type Ingredient struct {
	Name          string `json:"name"`
	QuantityGrams int    `json:"quantity_grams"`
	OriginalUnit  string `json:"original_unit,omitempty"`
}

type RecipeExtraction struct {
	Ingredients []Ingredient `json:"ingredients"`
}

