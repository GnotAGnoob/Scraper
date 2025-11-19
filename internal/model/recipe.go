package model

type IngredientMatch struct {
	Ingredient Ingredient      `json:"ingredient"`
	Product    *ProductSummary `json:"product,omitempty"`
	Error      string          `json:"error,omitempty"`
}

type RecipeResponse struct {
	Matches []IngredientMatch `json:"matches"`
}
