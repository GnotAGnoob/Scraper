package config

import (
	"flag"
	"strings"

	internalconfig "github.com/GnotAGnoob/kosik-scraper/internal/config"
)

type Config struct {
	*internalconfig.Config
	HTTPPort       int
	AllowedOrigins []string
}

func Load() (*Config, error) {
	if err := internalconfig.EnsureEnv(); err != nil {
		return nil, err
	}

	portDefault := internalconfig.GetEnvAsInt("SERVER_PORT", 8080)
	originsDefault := internalconfig.GetEnv("SERVER_ALLOWED_ORIGINS", "http://localhost:5173")

	portFlag := flag.Int("port", portDefault, "HTTP server port")
	originsFlag := flag.String("allowed-origins", originsDefault, "Comma-separated allowed origins")

	sharedCfg, err := internalconfig.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Config:         sharedCfg,
		HTTPPort:       *portFlag,
		AllowedOrigins: splitAndTrim(*originsFlag),
	}, nil
}

func splitAndTrim(input string) []string {
	parts := strings.Split(input, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
