package service

type MatchedProduct struct {
	IngredientName string  `json:"ingredient_name"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Unit           string  `json:"unit"`
	ImageURL       string  `json:"image_url"`
	Link           string  `json:"link"`
}

