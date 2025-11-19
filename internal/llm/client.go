package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/GnotAGnoob/kosik-scraper/internal/model"
)

const (
	defaultBaseURL        = "https://generativelanguage.googleapis.com/v1beta"
	defaultModel          = "models/gemini-2.5-flash"
	defaultTemperature    = 0.1
	defaultRequestTimeout = 30 * time.Second
)

var (
	ErrEmptyAPIKey      = errors.New("gemini api key is required")
	ErrEmptyModelResult = errors.New("gemini responded with no candidates")
)

type Client struct {
	apiKey     string
	model      string
	baseURL    string
	httpClient *http.Client
}

type Option func(*Client)

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

func WithModel(model string) Option {
	return func(c *Client) {
		if model != "" {
			c.model = model
		}
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		if baseURL != "" {
			c.baseURL = baseURL
		}
	}
}

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	client := &Client{
		apiKey:  apiKey,
		model:   defaultModel,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultRequestTimeout,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: defaultRequestTimeout,
		}
	}

	return client, nil
}

type generateContentRequest struct {
	Contents         []content         `json:"contents"`
	GenerationConfig *generationConfig `json:"generationConfig,omitempty"`
}

type generationConfig struct {
	Temperature      float32 `json:"temperature,omitempty"`
	ResponseMimeType string  `json:"responseMimeType,omitempty"`
}

type content struct {
	Role  string `json:"role,omitempty"`
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text,omitempty"`
}

type apiError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

type generateContentResponse struct {
	Candidates []struct {
		Content content `json:"content"`
	} `json:"candidates"`
}

type structuredIngredients struct {
	Ingredients []struct {
		Name          string   `json:"name"`
		QuantityGrams *float64 `json:"quantity_grams"`
		Notes         string   `json:"notes"`
	} `json:"ingredients"`
}

func (c *Client) ExtractIngredients(ctx context.Context, recipeText string) ([]model.Ingredient, error) {
	payload := generateContentRequest{
		Contents: []content{
			{
				Role: "user",
				Parts: []part{
					{Text: buildPrompt(recipeText)},
				},
			},
		},
		GenerationConfig: &generationConfig{
			Temperature:      defaultTemperature,
			ResponseMimeType: "application/json",
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.generateURL(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Set("key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr apiError
		if decodeErr := json.NewDecoder(resp.Body).Decode(&apiErr); decodeErr == nil && apiErr.Error.Message != "" {
			return nil, fmt.Errorf("gemini error: %s", apiErr.Error.Message)
		}
		return nil, fmt.Errorf("gemini request failed with status %s", resp.Status)
	}

	var llmResp generateContentResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResp); err != nil {
		return nil, err
	}

	if len(llmResp.Candidates) == 0 || len(llmResp.Candidates[0].Content.Parts) == 0 {
		return nil, ErrEmptyModelResult
	}

	text := llmResp.Candidates[0].Content.Parts[0].Text

	var parsed structuredIngredients
	if err := json.Unmarshal([]byte(text), &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse json response: %w", err)
	}

	results := make([]model.Ingredient, 0, len(parsed.Ingredients))
	for _, ing := range parsed.Ingredients {
		if ing.Name == "" {
			continue
		}
		results = append(results, model.Ingredient{
			Name:          ing.Name,
			QuantityGrams: ing.QuantityGrams,
			Notes:         ing.Notes,
		})
	}

	return results, nil
}

func (c *Client) generateURL() string {
	return fmt.Sprintf("%s/%s:generateContent", c.baseURL, c.model)
}

func buildPrompt(recipeText string) string {
	return fmt.Sprintf(`Extract a deduplicated ingredient list from the text below.
Return ONLY valid JSON with this shape:
{"ingredients":[{"name":"string","quantity_grams":number,"notes":"string"}]}
- quantity_grams must be grams if possible, otherwise omit the field
- use concise ingredient names suitable for grocery search

Recipe:
%s`, recipeText)
}
