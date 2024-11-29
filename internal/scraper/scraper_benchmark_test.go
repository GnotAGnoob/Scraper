package scraper_test

import (
	"os"
	"testing"

	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper"
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

func openSearchAndCloseBrowser(b *testing.B, searches ...string) {
	browser, err := scraper.InitScraper()
	if err != nil {
		b.Fatalf("error while initializing scraper: %v", err)
	}
	defer func() {
		err = browser.Cleanup()
		if err != nil {
			b.Errorf("error while cleaning up scraper: %v", err)
		}
	}()

	for _, search := range searches {
		totalChan := make(chan int)
		productsChan := make(chan *scraper.ProductResult)

		drainChan(totalChan)
		drainChan(productsChan)

		browser.GetKosikProducts(search, totalChan, productsChan)
	}
}

func BenchmarkScrape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		openSearchAndCloseBrowser(b, "susenky", "https://www.kosik.cz/c1319-slane?orderBy=unit-price-asc")
	}
}

func BenchmarkNotFoundScrape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		openSearchAndCloseBrowser(b, "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")
	}
}
