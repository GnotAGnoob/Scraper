package service

import (
	"context"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/llm"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/rs/zerolog/log"
)

type IngredientMatcher struct {
}

func NewIngredientMatcher() *IngredientMatcher {
	return &IngredientMatcher{}
}

type MatchedProduct struct {
	IngredientName string          `json:"ingredient_name"`
	QuantityGrams  int             `json:"quantity_grams"`
	Product        *shared.Product `json:"product,omitempty"`
	Found          bool            `json:"found"`
}

func (s *IngredientMatcher) MatchIngredients(ctx context.Context, ingredients []llm.Ingredient) []MatchedProduct {
	var wg sync.WaitGroup
	results := make([]MatchedProduct, len(ingredients))

	for i, ing := range ingredients {
		wg.Add(1)
		go func(index int, ingredient llm.Ingredient) {
			defer wg.Done()

			results[index] = MatchedProduct{
				IngredientName: ingredient.Name,
				QuantityGrams:  ingredient.QuantityGrams,
				Found:          false,
			}

			// Use buffered channels to prevent blocking if we were to stop reading (though we will drain)
			// But scraper writes to them, so buffer size helps performance.
			totalChan := make(chan int, 1)
			productsChan := make(chan *shared.ProductResult, 20)

			go func() {
				// TODO: Pass context to scraper if modified to support it
				err := scraper.GetProducts(ingredient.Name, totalChan, productsChan)
				if err != nil {
					log.Error().Err(err).Str("ingredient", ingredient.Name).Msg("Failed to scrape products")
				}
			}()

			// Read total count
			total, ok := <-totalChan
			if !ok || total == 0 {
				// Channel closed or no products
				// We need to ensure we drain productsChan if it was opened but total was 0 (unlikely based on code)
				// Based on code: if total == 0, it returns nil, and deferred close happens.
				// So we just drain productsChan to be sure.
				for range productsChan {
				}
				return
			}

			// Read products and pick the first valid one
			var firstProduct *shared.Product
			
			for pResult := range productsChan {
				if firstProduct == nil && pResult.Result.Value != nil && pResult.Result.ScrapeErr == nil {
					// Make a copy or just point to it? 
					// The product is a pointer, so we can just point to it.
					// We need to dereference *ProductResult -> Result -> Value (*Product)
					// Wait, pResult.Result.Value is *Product.
					firstProduct = pResult.Result.Value
				}
			}

			if firstProduct != nil {
				results[index].Product = firstProduct
				results[index].Found = true
			}

		}(i, ing)
	}

	wg.Wait()
	return results
}

