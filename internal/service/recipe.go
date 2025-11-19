package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	scraperLibShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type RecipeService struct {
	llmClient *llm.Client
}

func NewRecipeService() *RecipeService {
	llmClient, err := llm.NewClient()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize LLM client")
	}

	return &RecipeService{
		llmClient: llmClient,
	}
}

func (s *RecipeService) ProcessRecipe(ctx context.Context, recipeText string) ([]MatchedProduct, error) {
	ingredients, err := s.llmClient.ExtractIngredients(ctx, recipeText)
	if err != nil {
		return nil, fmt.Errorf("failed to extract ingredients: %w", err)
	}

	log.Info().Int("count", len(ingredients)).Msg("Extracted ingredients")

	var products []MatchedProduct
	var mu sync.Mutex

	g, gctx := errgroup.WithContext(ctx)

	for _, ingredient := range ingredients {
		ingredient := ingredient
		g.Go(func() error {
			product, err := s.findProductForIngredient(gctx, ingredient.Name)
			if err != nil {
				log.Warn().Err(err).Str("ingredient", ingredient.Name).Msg("Failed to find product")
				return nil
			}

			if product != nil {
				mu.Lock()
				products = append(products, *product)
				mu.Unlock()
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *RecipeService) findProductForIngredient(ctx context.Context, ingredientName string) (*MatchedProduct, error) {
	totalChan := make(chan int, 1)
	productsChan := make(chan *scraperLibShared.ProductResult, 100)

	errChan := make(chan error, 1)
	go func() {
		errChan <- scraper.GetProducts(ingredientName, totalChan, productsChan)
	}()

	total, ok := <-totalChan
	if !ok {
		return nil, fmt.Errorf("failed to get total for %s", ingredientName)
	}

	if err := <-errChan; err != nil {
		return nil, fmt.Errorf("error getting products for %s: %w", ingredientName, err)
	}

	if total == 0 {
		return nil, fmt.Errorf("no products found for %s", ingredientName)
	}

	var firstValidProduct *scraperLibShared.ReturnProduct
	for i := 0; i < total; i++ {
		productResult, ok := <-productsChan
		if !ok {
			break
		}
		if productResult.Result.ScrapeErr == nil && productResult.Result.Value != nil {
			product := productResult.Result.Value
			if product.Name.ScrapeErr == nil && product.Price.ScrapeErr == nil {
				firstValidProduct = productResult.Result
				break
			}
		}
	}

	if firstValidProduct == nil || firstValidProduct.Value == nil {
		return nil, fmt.Errorf("no valid product found for %s", ingredientName)
	}

	p := firstValidProduct.Value
	imageURL := ""
	if p.ImageUrl.ScrapeErr == nil && p.ImageUrl.Value != nil {
		imageURL = p.ImageUrl.Value.String()
	}

	link := ""
	if p.Link.ScrapeErr == nil && p.Link.Value != nil {
		link = p.Link.Value.String()
	}

	return &MatchedProduct{
		IngredientName: ingredientName,
		Name:           p.Name.Value,
		Price:          p.Price.Value,
		Unit:           p.Unit.Value,
		ImageURL:       imageURL,
		Link:           link,
	}, nil
}

