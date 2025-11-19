package api

import (
	"net/http"

	"github.com/GnotAGnoob/kosik-scraper/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	recipeService *service.RecipeService
}

func NewHandler() *Handler {
	return &Handler{
		recipeService: service.NewRecipeService(),
	}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/recipe", h.handleRecipe)
	return r
}

// handleRecipe extracts ingredients from recipe text and finds matching products
// @Summary      Extract ingredients from recipe text and find products
// @Description  Takes recipe text, extracts ingredients using LLM, and finds matching products from Kosik
// @Tags         recipe
// @Accept       json
// @Produce      json
// @Param        request  body      RecipeRequest  true  "Recipe text"
// @Success      200     {object}  RecipeResponse
// @Failure      400     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /recipe [post]
func (h *Handler) handleRecipe(w http.ResponseWriter, r *http.Request) {
	var req RecipeRequest
	if err := decodeJSONBody(w, r, &req); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.Text == "" {
		respondError(w, http.StatusBadRequest, "text field is required")
		return
	}

	products, err := h.recipeService.ProcessRecipe(r.Context(), req.Text)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, RecipeResponse{
		Products: products,
	})
}

