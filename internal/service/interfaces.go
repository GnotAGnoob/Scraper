package service

import (
	"context"

	"github.com/GnotAGnoob/kosik-scraper/internal/model"
)

type IngredientExtractor interface {
	ExtractIngredients(ctx context.Context, recipeText string) ([]model.Ingredient, error)
}

type ProductFetcher interface {
	FetchBestProduct(ctx context.Context, ingredient model.Ingredient) (*model.ProductSummary, error)
}
