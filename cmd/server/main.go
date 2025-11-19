package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GnotAGnoob/kosik-scraper/internal/api"
	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// @title           Recipe to Kosik Cart API
// @version         1.0
// @description     API for extracting ingredients from recipes and finding matching products on Kosik.cz
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	_ = godotenv.Load()

	port := flag.String("port", "8080", "server port")
	logLevel := flag.String("log-level", "info", "sets log level")
	flag.Parse()

	logger.Init(*logLevel)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	apiHandler := api.NewHandler()
	r.Mount("/api/v1", apiHandler.Routes())

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+*port+"/swagger/doc.json"),
	))

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: r,
	}

	go func() {
		log.Info().Msgf("Server starting on port %s", *port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited")
}

