package ingredient

import (
	"context"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
)

type IngredientInput struct {
	Name        string
	AmountGrams float64
}

type Match struct {
	Ingredient    string
	QuantityGrams float64
	Name          string
	URL           string
	ImageURL      string
	Price         float64
	Currency      string
}

func FindMatches(ctx context.Context, ingredients []IngredientInput) ([]Match, error) {
	if len(ingredients) == 0 {
		return nil, nil
	}

	matches := make([]Match, len(ingredients))

	var wg sync.WaitGroup
	errCh := make(chan error, len(ingredients))

	for i, ing := range ingredients {
		i, ing := i, ing
		wg.Add(1)
		go func() {
			defer wg.Done()

			select {
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			default:
			}

			product, err := searchBestProduct(ing.Name)
			if err != nil {
				errCh <- err
				return
			}
			if product == nil {
				return
			}

			match := Match{
				Ingredient:    ing.Name,
				QuantityGrams: ing.AmountGrams,
				Name:          product.Name.Value,
				URL:           product.Link.Value.String(),
				Price:         product.Price.Value,
				Currency:      "CZK",
			}

			if product.ImageUrl.Value != nil {
				match.ImageURL = product.ImageUrl.Value.String()
			}

			matches[i] = match
		}()
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	// Compact out empty matches (where no product was found).
	out := make([]Match, 0, len(matches))
	for _, m := range matches {
		if m.Name == "" {
			continue
		}
		out = append(out, m)
	}

	return out, nil
}

func searchBestProduct(search string) (*shared.Product, error) {
	totalChan := make(chan int)
	productsChan := make(chan *shared.ProductResult)
	errCh := make(chan error, 1)

	go func() {
		errCh <- scraper.GetProducts(search, totalChan, productsChan)
	}()

	// We don't really need the total, but we must drain the channel.
	go func() {
		for range totalChan {
		}
	}()

	var best *shared.Product

	for result := range productsChan {
		if result == nil || result.Result == nil {
			continue
		}
		if result.Result.ScrapeErr != nil {
			continue
		}

		product := result.Result.Value
		if product == nil {
			continue
		}

		// Basic validity checks.
		if product.IsSoldOut.ScrapeErr != nil || product.IsSoldOut.Value {
			continue
		}
		if product.Name.ScrapeErr != nil ||
			product.Price.ScrapeErr != nil ||
			product.Link.ScrapeErr != nil {
			continue
		}

		if best == nil {
			best = product
		}
	}

	err := <-errCh
	if err != nil {
		return nil, err
	}

	return best, nil
}


