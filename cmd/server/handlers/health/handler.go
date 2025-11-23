package health

import (
	"net/http"
	"time"

	serverconfig "github.com/GnotAGnoob/kosik-scraper/cmd/server/config"
	"github.com/GnotAGnoob/kosik-scraper/cmd/server/handlers/shared"
)

type response struct {
	Status string `json:"status"`
	Env    string `json:"env"`
	Time   string `json:"time"`
}

// HealthHandler godoc
// @Summary Health probe
// @Description Returns the server environment and current UTC timestamp.
// @Tags Health
// @Produce json
// @Success 200 {object} response
// @Router /health [get]
func HealthHandler(cfg *serverconfig.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shared.JSON(w, r, http.StatusOK, response{
			Status: "ok",
			Env:    cfg.Env,
			Time:   time.Now().UTC().Format(time.RFC3339),
		})
	}
}
