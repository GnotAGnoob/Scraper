package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env             string
	LogLevel        string
	HTTPPort        int
	AllowedOrigins  []string
	AllowedHeaders  []string
	ShutdownTimeout time.Duration
}

func Load() (*Config, error) {
	if err := ensureEnvFile(".env"); err != nil {
		return nil, err
	}

	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	defaultEnv := getEnv("APP_ENV", "development")
	defaultPort := getEnvAsInt("SERVER_PORT", 8080)
	defaultLogLevel := getEnv("LOG_LEVEL", "info")
	defaultOrigins := getEnv("SERVER_ALLOWED_ORIGINS", "http://localhost:5173")
	defaultHeaders := getEnv("SERVER_ALLOWED_HEADERS", "Accept,Content-Type")
	defaultShutdown := getEnvAsDuration("SERVER_SHUTDOWN_TIMEOUT", 5*time.Second)

	envFlag := flag.String("env", defaultEnv, "Runtime environment")
	portFlag := flag.Int("port", defaultPort, "HTTP server port")
	logLevel := flag.String("log-level", defaultLogLevel, "Log level (debug|info|warn|error)")
	origins := flag.String("allowed-origins", defaultOrigins, "Comma-separated allowed origins")
	headers := flag.String("allowed-headers", defaultHeaders, "Comma-separated allowed headers")
	flag.Parse()

	cfg := &Config{
		Env:             *envFlag,
		LogLevel:        *logLevel,
		HTTPPort:        *portFlag,
		AllowedOrigins:  splitAndTrim(*origins),
		AllowedHeaders:  splitAndTrim(*headers),
		ShutdownTimeout: defaultShutdown,
	}

	return cfg, nil
}

func ensureEnvFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf(".env file missing. create %s (see README for template)", filepath.Base(path))
		}

		return err
	}

	return nil
}

func getEnv(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return parsed
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}

	return parsed
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
