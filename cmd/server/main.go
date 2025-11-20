package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("load config")
	}

	logger.Init(cfg.LogLevel)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := NewServer(cfg)
	if err := srv.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("run server")
	}
}
