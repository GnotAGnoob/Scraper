package shared

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/orchestrator"
)

type errorResponse struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, r *http.Request, status int, payload any) {
	render.Status(r, status)
	render.JSON(w, r, payload)
}

func Error(w http.ResponseWriter, r *http.Request, status int, message string) {
	JSON(w, r, status, errorResponse{Error: message})
}

func OrchestratorError(w http.ResponseWriter, r *http.Request, err error) {
	logger := applogger.LoggerFromContext(r.Context())

	switch {
	case errors.Is(err, orchestrator.ErrNilConfig):
		logger.Warn().Msg("config is nil")
		Error(w, r, http.StatusInternalServerError, "internal server error")
	case errors.Is(err, orchestrator.ErrEmptyRecipe):
		logger.Warn().Msg("empty recipe")
		Error(w, r, http.StatusBadRequest, "bad request")
	default:
		logger.Error().Err(err).Msg("orchestrator error")
		Error(w, r, http.StatusInternalServerError, "internal server error")
	}
}
