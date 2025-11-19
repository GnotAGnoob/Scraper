package api

import (
	"encoding/json"
	"net/http"

	"github.com/GnotAGnoob/kosik-scraper/internal/service"
	"github.com/rs/zerolog/log"
)

type RecipeRequest struct {
	Text string `json:"text"`
}

type RecipeResponse struct {
	Products []service.MatchedProduct `json:"products"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	recipeService *service.RecipeService
}

func NewHandler(geminiApiKey string) *Handler {
	return &Handler{
		recipeService: service.NewRecipeService(geminiApiKey),
	}
}

// ProcessRecipe godoc
// @Summary Process recipe text and find matching products
// @Description Extracts ingredients from recipe text using AI and searches for matching products on Kosik
// @Tags recipe
// @Accept json
// @Produce json
// @Param request body RecipeRequest true "Recipe text"
// @Success 200 {object} RecipeResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/recipe [post]
func (h *Handler) ProcessRecipe(w http.ResponseWriter, r *http.Request) {
	var req RecipeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Err(err).Msg("Failed to decode request")
		respondError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		respondError(w, "Recipe text is required", http.StatusBadRequest)
		return
	}

	products, err := h.recipeService.ProcessRecipe(r.Context(), req.Text)
	if err != nil {
		log.Error().Err(err).Msg("Failed to process recipe")
		respondError(w, "Failed to process recipe", http.StatusInternalServerError)
		return
	}

	respondJSON(w, RecipeResponse{Products: products}, http.StatusOK)
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, map[string]string{"status": "ok"}, http.StatusOK)
}

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, message string, status int) {
	respondJSON(w, ErrorResponse{Error: message}, status)
}

