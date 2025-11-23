package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"slices"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	LogLevel    string
	LLMProvider string
	LLMAPIKey   string
	LLMModel    string
}

var allowedLLMProviders = map[string][]string{
	"google": {
		"gemini-2.5-flash",
	},
}

const (
	defaultProvider = "google"
	defaultModel    = "gemini-2.5-flash"
)

func Load() (*Config, error) {
	if err := EnsureEnv(); err != nil {
		return nil, err
	}

	defaultEnv := GetEnv("APP_ENV", "development")
	defaultLogLevel := GetEnv("LOG_LEVEL", "info")
	defaultLLMProvider := GetEnv("LLM_PROVIDER", defaultProvider)
	defaultLLMModel := GetEnv("LLM_MODEL", defaultModel)
	defaultLLMAPIKey := GetEnv("LLM_API_KEY", "")

	envFlag := flag.String("env", defaultEnv, "Runtime environment")
	logLevel := flag.String("log-level", defaultLogLevel, "Log level (debug|info|warn|error)")
	llmProvider := flag.String("llm-provider", defaultLLMProvider, "LLM provider")
	llmModel := flag.String("llm-model", defaultLLMModel, "LLM model identifier")
	llmAPIKey := flag.String("llm-api-key", defaultLLMAPIKey, "LLM API key")

	if !flag.Parsed() {
		flag.Parse()
	}

	if *llmAPIKey == "" {
		return nil, fmt.Errorf("LLM API key is required")
	}

	provider, model, err := normalizeLLM(*llmProvider, *llmModel)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:         strings.TrimSpace(*envFlag),
		LogLevel:    strings.TrimSpace(*logLevel),
		LLMProvider: provider,
		LLMAPIKey:   *llmAPIKey,
		LLMModel:    model,
	}

	return cfg, nil
}

func EnsureEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("load .env: %w", err)
	}

	return nil
}

func GetEnv(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}

	return fallback
}

func GetEnvAsInt(key string, fallback int) int {
	value := GetEnv(key, "")
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}

func normalizeLLM(provider, model string) (normalizedProvider, normalizedModel string, err error) {
	normalizedProvider = strings.ToLower(strings.TrimSpace(provider))
	allowedModels, ok := allowedLLMProviders[normalizedProvider]
	if !ok {
		return "", "", fmt.Errorf("unsupported LLM provider %q", provider)
	}

	normalizedModel = strings.ToLower(strings.TrimSpace(model))
	if !slices.Contains(allowedModels, normalizedModel) {
		return "", "", fmt.Errorf("unsupported model %q for provider %q", model, provider)
	}

	return
}
