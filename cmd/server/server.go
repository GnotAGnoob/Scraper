package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
)

type Server struct {
	cfg    *config.Config
	router chi.Router
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{cfg: cfg}
	s.router = s.buildRouter()
	return s
}

func (s *Server) Run(ctx context.Context) error {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.cfg.HTTPPort),
		Handler: s.router,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Info().Int("port", s.cfg.HTTPPort).Msg("server starting")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
			return
		}
		serverErr <- nil
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()
		log.Info().Msg("server shutting down")
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			return err
		}
		return nil
	case err := <-serverErr:
		return err
	}
}

func (s *Server) buildRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   s.cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   s.cfg.AllowedHeaders,
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/healthz", s.handleHealth())
	r.Route("/api/v1", func(api chi.Router) {
		api.Post("/recipe", s.handleRecipe())
	})

	r.Get("/swagger/*", s.handleSwaggerPlaceholder())

	return r
}

func (s *Server) handleHealth() http.HandlerFunc {
	type healthResponse struct {
		Status string `json:"status"`
		Env    string `json:"env"`
		Time   string `json:"time"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, healthResponse{
			Status: "ok",
			Env:    s.cfg.Env,
			Time:   time.Now().UTC().Format(time.RFC3339),
		})
	}
}

func (s *Server) handleRecipe() http.HandlerFunc {
	type recipeRequest struct {
		Text string `json:"text"`
	}

	type recipeResponse struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req recipeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
			return
		}

		if strings.TrimSpace(req.Text) == "" {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "text is required"})
			return
		}

		respondJSON(w, http.StatusNotImplemented, recipeResponse{
			Message: "recipe endpoint wiring in progress",
		})
	}
}

func (s *Server) handleSwaggerPlaceholder() http.HandlerFunc {
	type placeholder struct {
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusNotImplemented, placeholder{
			Message: "swagger docs coming soon. generate via `swag init` once handlers annotated",
		})
	}
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Error().Err(err).Msg("write JSON response")
	}
}
