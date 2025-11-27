package orchestrator

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/GnotAGnoob/kosik-scraper/internal/config"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	scraperLibShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
)

const (
	maxWorkers         = 10
	maxProducts        = 25
	productSourceKosik = "kosik"
)

var (
	ErrEmptyRecipe         = errors.New("recipe text is empty")
	ErrEmptyIngredientName = errors.New("empty ingredient name")
	ErrNilConfig           = errors.New("config is nil")
)

type Service struct {
	cfg *config.Config
}

func NewService(ctx context.Context, cfg *config.Config) (*Service, error) {
	logger := applogger.LoggerFromContext(ctx)
	if cfg == nil {
		logger.Error().Msg("config is nil")
		return nil, ErrNilConfig
	}
	return &Service{cfg: cfg}, nil
}

func (s *Service) ProcessTextToProducts(ctx context.Context, text string) ([]IngredientProducts, error) {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return nil, ErrEmptyRecipe
	}

	logger := applogger.LoggerFromContext(ctx)
	logger.Info().Int("text_length", len(trimmed)).Msg("parsing recipe with LLM")

	ingredients, err := llm.TextToIngredients(ctx, s.cfg, trimmed)
	if err != nil {
		logger.Error().Err(err).Msg("LLM recipe parsing failed")
		return nil, err
	}

	ingredientCount := len(ingredients)
	logger.Info().Int("ingredients_count", ingredientCount).Msg("recipe parsed successfully")

	if ingredientCount == 0 {
		logger.Warn().Msg("no ingredients recognized")
		return []IngredientProducts{}, nil
	}

	return s.scrapeIngredients(ctx, ingredients)
}

func (s *Service) scrapeIngredients(ctx context.Context, ingredients []shared.Ingredient) ([]IngredientProducts, error) {
	logger := applogger.LoggerFromContext(ctx)
	if len(ingredients) == 0 {
		logger.Warn().Msg("no ingredients provided")
		return []IngredientProducts{}, nil
	}

	logger.Info().Int("ingredients_count", len(ingredients)).Int("max_workers", maxWorkers).Msg("scraping ingredients")
	ingredientsProducts := make([]IngredientProducts, len(ingredients))
	g, gctx := errgroup.WithContext(ctx)
	semaphore := make(chan struct{}, maxWorkers)

	for idx, ing := range ingredients {
		id := idx
		ingredient := ing
		g.Go(func() error {
			// limits the number of concurrent requests by using a semaphore pattern
			select {
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()
			case <-gctx.Done():
				return gctx.Err()
			}

			ingredientProducts, err := s.scrapeIngredient(gctx, ingredient)
			if err != nil {
				logger.Warn().Err(err).Str("ingredient", ingredient.Name).Msg("scrape failed")
			}
			ingredientsProducts[id] = ingredientProducts
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		logger.Warn().Err(err).Msg("scraping ingredients products cancelled")
		return ingredientsProducts, err
	}

	logger.Info().Int("ingredients_products_count", len(ingredientsProducts)).Msg("scraping ingredients products completed")
	return ingredientsProducts, nil
}

func (s *Service) scrapeIngredient(ctx context.Context, ingredient shared.Ingredient) (IngredientProducts, error) {
	logger := applogger.LoggerFromContext(ctx)
	name := strings.TrimSpace(ingredient.Name)

	ingredientProducts := IngredientProducts{
		Ingredient: ingredient,
		Source:     productSourceKosik,
	}

	if name == "" {
		logger.Error().Str("ingredient", ingredient.Name).Msg("empty ingredient name")
		return ingredientProducts, ErrEmptyIngredientName
	}

	totalChan := make(chan int)
	productsChan := make(chan *scraperLibShared.ProductResult)
	errChan := make(chan error, 1)

	go func() {
		errChan <- scraper.GetProducts(ctx, name, maxProducts, totalChan, productsChan)
	}()

	var total int
	select {
	case t, ok := <-totalChan:
		if !ok {
			// Channel closed, wait for error from goroutine
			select {
			case err := <-errChan:
				logger.Error().Err(err).Msg("failed to read total products")
				return ingredientProducts, err
			case <-ctx.Done():
				return ingredientProducts, ctx.Err()
			}
		}
		total = t
	case <-ctx.Done():
		logger.Error().Err(ctx.Err()).Msg("context done")
		return ingredientProducts, ctx.Err()
	}

	if total == 0 {
		msg := "no products found"
		logger.Warn().Str("ingredient", name).Msg(msg)
		return ingredientProducts, nil
	}

	// Collect up to maxProducts, but total is kept for paging purposes
	expectedCount := maxProducts
	if total < maxProducts {
		expectedCount = total
	}

	// Allocate array of size total to accommodate all possible indices
	products := make([]*scraperLibShared.ReturnProduct, expectedCount)
	collectedCount := 0
	for i := 0; i < expectedCount; i++ {
		select {
		case productResult, ok := <-productsChan:
			if !ok {
				// Channel closed early, collect what we have
				logger.Warn().Int("collected", collectedCount).Int("expected", expectedCount).Msg("channel closed early, collecting partial results")
				// Wait for error from goroutine
				var scrapeErr error
				select {
				case err := <-errChan:
					scrapeErr = err
					logger.Error().Err(err).Msg("scraping failed")
				case <-ctx.Done():
					return ingredientProducts, ctx.Err()
				}
				// Return partial results with error
				transformedProducts := transformProducts(products)
				ingredientProducts.Products = transformedProducts
				return ingredientProducts, scrapeErr
			}
			products[productResult.Index] = productResult.Result
			collectedCount++
		case <-ctx.Done():
			logger.Error().Err(ctx.Err()).Int("collected", collectedCount).Int("expected", expectedCount).Msg("context cancelled, returning partial results")
			// Return what we collected so far
			transformedProducts := transformProducts(products)
			ingredientProducts.Products = transformedProducts
			return ingredientProducts, ctx.Err()
		}
	}

	scrapeErr := <-errChan
	transformedProducts := transformProducts(products)
	ingredientProducts.Products = transformedProducts

	return ingredientProducts, scrapeErr
}
