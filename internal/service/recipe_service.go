package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/kosik"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type MatchedProduct struct {
	IngredientName string  `json:"ingredientName"`
	ProductName    string  `json:"productName"`
	Price          float64 `json:"price"`
	Unit           string  `json:"unit"`
	ImageURL       string  `json:"imageUrl"`
	ProductURL     string  `json:"productUrl"`
	PricePerUnit   float64 `json:"pricePerUnit"`
	PricePerUnitUnit string  `json:"pricePerUnitUnit"`
}

type RecipeService struct {
	geminiClient *llm.GeminiClient
}

func NewRecipeService(geminiApiKey string) *RecipeService {
	return &RecipeService{
		geminiClient: llm.NewGeminiClient(geminiApiKey),
	}
}

func (s *RecipeService) ProcessRecipe(ctx context.Context, recipeText string) ([]MatchedProduct, error) {
	// Extract ingredients using Gemini
	ingredients, err := s.geminiClient.ExtractIngredients(ctx, recipeText)
	if err != nil {
		return nil, fmt.Errorf("failed to extract ingredients: %w", err)
	}

	log.Info().Msgf("Extracted %d ingredients", len(ingredients.Ingredients))

	// Search for products concurrently
	results := make([]MatchedProduct, 0, len(ingredients.Ingredients))
	resultsMu := sync.Mutex{}

	g, gCtx := errgroup.WithContext(ctx)

	for _, ingredient := range ingredients.Ingredients {
		ingredient := ingredient // capture loop variable
		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return gCtx.Err()
			default:
			}

			product, err := s.searchFirstProduct(ingredient.Name)
			if err != nil {
				log.Warn().Err(err).Msgf("Failed to find product for ingredient: %s", ingredient.Name)
				return nil // Continue with other ingredients
			}

			if product != nil {
				resultsMu.Lock()
				results = append(results, *product)
				resultsMu.Unlock()
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *RecipeService) searchFirstProduct(ingredientName string) (*MatchedProduct, error) {
	totalChan := make(chan int, 1)
	productsChan := make(chan *shared.ProductResult, 100)

	go func() {
		if err := kosik.GetProducts(ingredientName, totalChan, productsChan); err != nil {
			log.Error().Err(err).Msgf("Error searching for %s", ingredientName)
		}
	}()

	// Wait for total
	total := <-totalChan
	if total == 0 {
		return nil, fmt.Errorf("no products found for: %s", ingredientName)
	}

	// Get first valid product
	for productResult := range productsChan {
		if productResult.Result.ScrapeErr != nil {
			continue
		}

		product := productResult.Result.Value
		if product == nil {
			continue
		}

		// Found valid product
		matched := &MatchedProduct{
			IngredientName: ingredientName,
		}

		if product.Name.ScrapeErr == nil {
			matched.ProductName = product.Name.Value
		}

		if product.Price.ScrapeErr == nil {
			matched.Price = product.Price.Value
		}

		if product.Unit.ScrapeErr == nil {
			matched.Unit = product.Unit.Value
		}

		if product.ImageUrl.ScrapeErr == nil && product.ImageUrl.Value != nil {
			matched.ImageURL = product.ImageUrl.Value.String()
		}

		if product.Link.ScrapeErr == nil && product.Link.Value != nil {
			matched.ProductURL = product.Link.Value.String()
		}

		if product.PricePerUnit.ScrapeErr == nil {
			matched.PricePerUnit = product.PricePerUnit.Value.Value
			matched.PricePerUnitUnit = product.PricePerUnit.Value.Unit
		}

		// Drain remaining products
		go func() {
			for range productsChan {
			}
		}()

		return matched, nil
	}

	return nil, fmt.Errorf("no valid products found for: %s", ingredientName)
}

