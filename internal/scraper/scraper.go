package scraper

import (
	"context"

	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/kosik"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
)

// todo add the setting of the address for kosik site
// todo debug logs
// todo caching?
// todo add the ability to scrape other sites
func GetProducts(ctx context.Context, search string, totalChan chan<- int, productsChan chan<- *shared.ProductResult) error {
	return kosik.GetProducts(ctx, search, totalChan, productsChan)
}
