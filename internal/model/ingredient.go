package model

type Ingredient struct {
	Name          string   `json:"name"`
	QuantityGrams *float64 `json:"quantity_grams,omitempty"`
	Notes         string   `json:"notes,omitempty"`
}
