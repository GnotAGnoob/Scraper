package service

import (
	"context"
	"fmt"
	"math"

	"github.com/GnotAGnoob/kosik-scraper/internal/model"
	scraperlib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
)

type KosikFetcher struct{}

func NewKosikFetcher() *KosikFetcher {
	return &KosikFetcher{}
}

func (f *KosikFetcher) FetchBestProduct(ctx context.Context, ingredient model.Ingredient) (*model.ProductSummary, error) {
	query := buildSearchQuery(ingredient)
	if query == "" {
		return nil, ErrNoProductFound
	}

	totalChan := make(chan int)
	productsChan := make(chan *shared.ProductResult)
	errChan := make(chan error, 1)

	go func() {
		errChan <- scraperlib.GetProducts(query, totalChan, productsChan)
	}()

	total, ok := <-totalChan
	if !ok {
		return nil, fmt.Errorf("failed to read totals for %s", query)
	}

	if total == 0 {
		if err := <-errChan; err != nil {
			return nil, err
		}
		return nil, ErrNoProductFound
	}

	var best *model.ProductSummary
	for {
		select {
		case <-ctx.Done():
			go drainProducts(productsChan)
			<-errChan
			return nil, ctx.Err()
		case productResult, ok := <-productsChan:
			if !ok {
				goto exit
			}
			if best == nil {
				best = convertProductResult(productResult)
			}
		}
	}

exit:
	if err := <-errChan; err != nil {
		return nil, err
	}
	if best == nil {
		return nil, ErrNoProductFound
	}
	return best, nil
}

func drainProducts(productsChan <-chan *shared.ProductResult) {
	for range productsChan {
	}
}

func buildSearchQuery(ingredient model.Ingredient) string {
	if ingredient.Name == "" {
		return ""
	}
	if ingredient.QuantityGrams == nil || *ingredient.QuantityGrams == 0 {
		return ingredient.Name
	}
	rounded := math.Round(*ingredient.QuantityGrams)
	if rounded <= 0 {
		return ingredient.Name
	}
	return fmt.Sprintf("%s %dg", ingredient.Name, int(rounded))
}

func convertProductResult(result *shared.ProductResult) *model.ProductSummary {
	if result == nil || result.Result == nil || result.Result.ScrapeErr != nil || result.Result.Value == nil {
		return nil
	}

	product := result.Result.Value
	if product.Name.ScrapeErr != nil || product.Name.Value == "" {
		return nil
	}

	var unit string
	if product.Unit.ScrapeErr == nil {
		unit = product.Unit.Value
	}

	summary := &model.ProductSummary{
		Name: product.Name.Value,
		Unit: unit,
	}

	if product.Price.ScrapeErr == nil {
		price := product.Price.Value
		summary.Price = &price
	}

	if product.PricePerUnit.ScrapeErr == nil {
		ppu := product.PricePerUnit.Value
		summary.PricePerUnit = &model.PricePerUnit{
			Value: ppu.Value,
			Unit:  ppu.Unit,
		}
	}

	if product.Link.ScrapeErr == nil && product.Link.Value != nil {
		summary.Link = product.Link.Value.String()
	}

	if product.ImageUrl.ScrapeErr == nil && product.ImageUrl.Value != nil {
		summary.ImageURL = product.ImageUrl.Value.String()
	}

	if product.IsSoldOut.ScrapeErr == nil {
		val := product.IsSoldOut.Value
		summary.IsSoldOut = &val
	}

	return summary
}
