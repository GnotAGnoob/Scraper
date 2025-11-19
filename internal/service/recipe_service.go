package service

import (
	"context"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/model"
	"golang.org/x/sync/errgroup"
)

const defaultWorkerLimit = 4

type RecipeService struct {
	extractor     IngredientExtractor
	fetcher       ProductFetcher
	maxConcurrent int
}

type Option func(*RecipeService)

func WithMaxConcurrency(limit int) Option {
	return func(s *RecipeService) {
		if limit > 0 {
			s.maxConcurrent = limit
		}
	}
}

func NewRecipeService(extractor IngredientExtractor, fetcher ProductFetcher, opts ...Option) *RecipeService {
	service := &RecipeService{
		extractor:     extractor,
		fetcher:       fetcher,
		maxConcurrent: defaultWorkerLimit,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

func (s *RecipeService) ProcessRecipe(ctx context.Context, recipeText string) (*model.RecipeResponse, error) {
	if strings.TrimSpace(recipeText) == "" {
		return nil, ErrEmptyRecipeText
	}

	ingredients, err := s.extractor.ExtractIngredients(ctx, recipeText)
	if err != nil {
		return nil, err
	}

	if len(ingredients) == 0 {
		return &model.RecipeResponse{Matches: []model.IngredientMatch{}}, nil
	}

	workerLimit := s.maxConcurrent
	if workerLimit <= 0 {
		workerLimit = defaultWorkerLimit
	}

	matches := make([]model.IngredientMatch, len(ingredients))
	sem := make(chan struct{}, workerLimit)

	eg, egCtx := errgroup.WithContext(ctx)
	for idx, ingredient := range ingredients {
		i := idx
		ing := ingredient
		eg.Go(func() error {
			select {
			case sem <- struct{}{}:
				defer func() { <-sem }()
			case <-egCtx.Done():
				return egCtx.Err()
			}

			product, fetchErr := s.fetcher.FetchBestProduct(egCtx, ing)
			if fetchErr != nil {
				matches[i] = model.IngredientMatch{
					Ingredient: ing,
					Error:      fetchErr.Error(),
				}
				return nil
			}

			matches[i] = model.IngredientMatch{
				Ingredient: ing,
				Product:    product,
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return &model.RecipeResponse{Matches: matches}, nil
}
