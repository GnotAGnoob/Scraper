package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	ingredientsvc "github.com/GnotAGnoob/kosik-scraper/internal/service/ingredient"
)

type RecipeRequest struct {
	// Free-form recipe text (CZ/EN) from which ingredients will be extracted.
	Text string `json:"text"`
}

type ProductMatch struct {
	Ingredient    string  `json:"ingredient"`
	QuantityGrams float64 `json:"quantityGrams"`
	Name          string  `json:"name"`
	URL           string  `json:"url"`
	ImageURL      string  `json:"imageUrl"`
	Price         float64 `json:"price"`
	Currency      string  `json:"currency"`
}

type RecipeResponse struct {
	Matches []ProductMatch `json:"matches"`
}

// handleRecipe godoc
// @Summary      Convert recipe text to Kosik products
// @Description  Extract ingredients using Gemini and match them to Kosik.cz products.
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        request  body      RecipeRequest  true  "Recipe text payload"
// @Success      200      {object}  RecipeResponse
// @Failure      400      {string}  string  "invalid json or missing text"
// @Failure      500      {string}  string  "internal error"
// @Router       /api/v1/recipe [post]
func handleRecipe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RecipeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Text) == "" {
		http.Error(w, "`text` is required", http.StatusBadRequest)
		return
	}

	ingredients, err := llm.ExtractIngredients(ctx, req.Text)
	if err != nil {
		http.Error(w, "failed to extract ingredients", http.StatusInternalServerError)
		return
	}

	inputs := make([]ingredientsvc.IngredientInput, 0, len(ingredients))
	for _, ing := range ingredients {
		if strings.TrimSpace(ing.Name) == "" {
			continue
		}
		inputs = append(inputs, ingredientsvc.IngredientInput{
			Name:        ing.Name,
			AmountGrams: ing.AmountGrams,
		})
	}

	matches, err := ingredientsvc.FindMatches(ctx, inputs)
	if err != nil {
		http.Error(w, "failed to find products", http.StatusInternalServerError)
		return
	}

	respMatches := make([]ProductMatch, 0, len(matches))
	for _, m := range matches {
		respMatches = append(respMatches, ProductMatch{
			Ingredient:    m.Ingredient,
			QuantityGrams: m.QuantityGrams,
			Name:          m.Name,
			URL:           m.URL,
			ImageURL:      m.ImageURL,
			Price:         m.Price,
			Currency:      m.Currency,
		})
	}

	resp := RecipeResponse{
		Matches: respMatches,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

