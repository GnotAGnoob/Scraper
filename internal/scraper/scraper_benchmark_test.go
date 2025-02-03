package scraper_test

import (
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
		totalChan := make(chan int)
		productsChan := make(chan *scraperShared.ProductResult)

		drainChan(totalChan)
		drainChan(productsChan)

		scraper.GetProducts(search, totalChan, productsChan)
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
