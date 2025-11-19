package api

import (
	"encoding/json"
	"net/http"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	"github.com/GnotAGnoob/kosik-scraper/internal/service"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	llmClient *llm.Client
	matcher   *service.IngredientMatcher
}

func NewHandler(llm *llm.Client, matcher *service.IngredientMatcher) *Handler {
	return &Handler{
		llmClient: llm,
		matcher:   matcher,
	}
}

type RecipeRequest struct {
	Text string `json:"text"`
}

type RecipeResponse struct {
	Ingredients []service.MatchedProduct `json:"ingredients"`
}

// HandleRecipeToCart godoc
// @Summary Convert recipe text to Kosik cart products
// @Description Extracts ingredients from text and finds matching products on Kosik
// @Tags recipe
// @Accept json
// @Produce json
// @Param request body RecipeRequest true "Recipe text"
// @Success 200 {object} RecipeResponse
// @Router /recipe [post]
func (h *Handler) HandleRecipeToCart(w http.ResponseWriter, r *http.Request) {
	var req RecipeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Info().Msg("Extracting ingredients from recipe")
	ingredients, err := h.llmClient.ExtractIngredients(r.Context(), req.Text)
	if err != nil {
		log.Error().Err(err).Msg("Failed to extract ingredients")
		http.Error(w, "Failed to process recipe", http.StatusInternalServerError)
		return
	}

	log.Info().Int("count", len(ingredients)).Msg("Ingredients extracted. Matching products...")
	matched := h.matcher.MatchIngredients(r.Context(), ingredients)

	response := RecipeResponse{
		Ingredients: matched,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error().Err(err).Msg("Failed to encode response")
	}
}
