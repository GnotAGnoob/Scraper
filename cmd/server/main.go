package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	serverconfig "github.com/GnotAGnoob/kosik-scraper/cmd/server/config"
	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := serverconfig.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("load config")
	}

	logger.Init(cfg.LogLevel)
	log.Info().Str("level", cfg.LogLevel).Msg("logger initialized")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv, err := NewServer(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("init server")
	}
	log.Info().Msg("server initialized")

	if err := srv.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("run server")
	}
}
