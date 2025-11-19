package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GnotAGnoob/kosik-scraper/internal/service"
)

type RecipeHandler struct {
	service *service.RecipeService
}

type RecipeRequest struct {
	Text string `json:"text"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewRecipeHandler(service *service.RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

// Recipe godoc
// @Summary     Extracts ingredients and matches Kosik products
// @Description Accepts raw recipe text, extracts structured ingredients via Gemini, and fetches the first Kosik product per ingredient.
// @Tags        recipe
// @Accept      json
// @Produce     json
// @Param       request body RecipeRequest true "Recipe input"
// @Success     200 {object} model.RecipeResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/v1/recipe [post]
func (h *RecipeHandler) HandleRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req RecipeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request body"})
		return
	}

	result, err := h.service.ProcessRecipe(r.Context(), req.Text)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, service.ErrEmptyRecipeText) {
			status = http.StatusBadRequest
		}
		writeJSON(w, status, ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
