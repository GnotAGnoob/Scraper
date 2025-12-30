package scraper_test

import (
	"context"
	"os"
	"testing"

	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	scraperShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
)

func TestMain(m *testing.M) {
	logger.Init("disabled")

	code := m.Run()
	os.Exit(code)
}

func drainChan[T any](ch chan T) {
	go func() {
		for range ch {
		}
	}()
}

func openSearchAndCloseBrowser(searches ...string) {
	for _, search := range searches {
		infoChan := make(chan scraperShared.SearchInfo)
		productsChan := make(chan *scraperShared.ProductResult)
		ctx := context.Background()

		drainChan(infoChan)
		drainChan(productsChan)

		_ = scraper.GetProducts(ctx, search, 10, infoChan, productsChan)
	}
}

func BenchmarkScrape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		openSearchAndCloseBrowser("susenky", "https://www.kosik.cz/c1319-slane?orderBy=unit-price-asc")
	}
}

func BenchmarkNotFoundScrape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		openSearchAndCloseBrowser("wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")
	}
}
