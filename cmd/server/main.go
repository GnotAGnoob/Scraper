package main

// @title           ScrapeMonkey API
// @version         1.0
// @description     API that converts recipe text into Kosik.cz product matches.
// @BasePath        /

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	// _ "github.com/GnotAGnoob/kosik-scraper/docs" // TODO: generate with swag init
	"github.com/GnotAGnoob/kosik-scraper/internal/api"
	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP listen address")
	logLevel := flag.String("log-level", "info", "log level: debug|info|disabled")
	flag.Parse()

	logger.Init(*logLevel)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	api.RegisterRoutes(r)

	if err := http.ListenAndServe(*addr, r); err != nil {
		// fall back to std lib logging to avoid circular deps
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
