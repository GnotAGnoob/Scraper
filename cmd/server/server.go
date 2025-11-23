package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	serverconfig "github.com/GnotAGnoob/kosik-scraper/cmd/server/config"
	"github.com/GnotAGnoob/kosik-scraper/internal/orchestrator"
	"github.com/rs/zerolog/log"
)

type server struct {
	cfg          *serverconfig.Config
	orchestrator *orchestrator.Service
	router       chi.Router
}

func NewServer(ctx context.Context, cfg *serverconfig.Config) (*server, error) {
	orch, err := orchestrator.NewService(ctx, cfg.Config)
	if err != nil {
		log.Error().Err(err).Msg("orchestrator initialization failed")
		return nil, err
	}
	log.Info().Msg("orchestrator initialized")

	s := &server{
		cfg:          cfg,
		orchestrator: orch,
	}
	s.router = s.buildRouter()
	log.Info().Msg("router built")
	return s, nil
}

func (s *server) Run(ctx context.Context) error {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.cfg.HTTPPort),
		Handler: s.router,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Info().Str("address", httpServer.Addr).Msg("server starting")

		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
			return
		}
		serverErr <- nil
	}()

	select {
	case <-ctx.Done():
		const shutdownTimeout = 10 * time.Second
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		log.Info().Dur("timeout", shutdownTimeout).Msg("server shutting down")

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("server shutdown failed")
			return err
		}

		log.Info().Msg("server shutdown complete")
		return nil
	case err := <-serverErr:
		log.Error().Err(err).Msg("server error")
		return err
	}
}
