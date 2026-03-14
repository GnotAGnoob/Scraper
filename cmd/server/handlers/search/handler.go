package search

import (
	"encoding/json"
	"net/http"

	"github.com/GnotAGnoob/kosik-scraper/cmd/server/handlers/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/orchestrator"
)

type request struct {
	Queries []string `json:"queries" validate:"required"`
}

const maxRequestBodySize = 10 * 1024 * 1024 // 10MB

// SearchProductsHandler godoc
// @Summary Search products directly by query or URL
// @Description Scrapes products directly by search terms or Kosik.cz URLs, bypassing LLM ingredient extraction.
// @Tags Search
// @Accept json
// @Produce json
// @Param request body request true "Search queries payload"
// @Success 200 {object} orchestrator.RecipeResult
// @Failure 400 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /api/v1/search/products [post]
func SearchProductsHandler(orch *orchestrator.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := applogger.LoggerFromContext(r.Context())

		r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)
		defer func() {
			r.Body.Close()
		}()

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Warn().Err(err).Msg("invalid JSON body")
			shared.Error(w, r, http.StatusBadRequest, "invalid JSON body")
			return
		}

		if len(req.Queries) == 0 {
			logger.Warn().Msg("empty queries provided")
			shared.Error(w, r, http.StatusBadRequest, "queries are required")
			return
		}

		logger.Info().Int("queries_count", len(req.Queries)).Msg("processing direct search")
		result, err := orch.ProcessDirectSearch(r.Context(), req.Queries)
		if err != nil {
			logger.Error().Err(err).Msg("direct search failed")
			shared.OrchestratorError(w, r, err)
			return
		}

		logger.Info().Int("products_count", len(result.IngredientsProducts)).Msg("direct search successful")
		shared.JSON(w, r, http.StatusOK, result)
	}
}
