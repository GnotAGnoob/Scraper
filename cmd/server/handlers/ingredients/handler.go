package ingredients

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/cmd/server/handlers/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/orchestrator"
)

type request struct {
	Text string `json:"text"`
}

const maxRequestBodySize = 10 * 1024 * 1024 // 10MB

// TextProductsHandler godoc
// @Summary Convert text into scraped products shop suggestions
// @Description The LLM extracts ingredients which are matched against shop products.
// @Tags Text
// @Accept json
// @Produce json
// @Param request body request true "Text payload"
// @Success 200 {object} orchestrator.RecipeResult
// @Failure 400 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /api/v1/text/products [post]
func TextProductsHandler(orch *orchestrator.Service) http.HandlerFunc {
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

		if strings.TrimSpace(req.Text) == "" {
			logger.Warn().Msg("empty text provided")
			shared.Error(w, r, http.StatusBadRequest, "text is required")
			return
		}

		textLen := len(req.Text)
		logger.Info().Int("text_length", textLen).Msg("processing text to products")
		result, err := orch.ProcessTextToProducts(r.Context(), req.Text)
		if err != nil {
			logger.Error().Err(err).Msg("processing failed")
			shared.OrchestratorError(w, r, err)
			return
		}

		logger.Info().Int("products_count", len(result.IngredientsProducts)).Msg("processing successful")
		shared.JSON(w, r, http.StatusOK, result)
	}
}
