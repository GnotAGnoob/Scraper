package swagger

import (
	"net/http"

	"github.com/GnotAGnoob/kosik-scraper/cmd/server/handlers/shared"
)

type response struct {
	Message string `json:"message"`
}

func SwaggerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shared.JSON(w, r, http.StatusNotImplemented, response{
			Message: "swagger docs coming soon. generate via `swag init` once handlers annotated",
		})
	}
}
