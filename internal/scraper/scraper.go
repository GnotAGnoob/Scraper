package scraper

import (
	"fmt"
	"time"

	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
	"github.com/GnotAGnoob/kosik-scraper/internal/utils/urlParams"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/rs/zerolog/log"
)

type Scraper struct {
	browser *rod.Browser
}

var scrapper = &Scraper{}

type returnProduct struct {
	structs.ScrapeResult[*Product]
}

func InitScraper() (*Scraper, error) {
	if scrapper.browser != nil {
		return scrapper, nil
	}

	// leakless is a binary that prevents zombie processes
	// but the problem is that windows defender detects it as a virus
	// because according to internet, it is used in many viruses
	launcher := launcher.New().Leakless(false).Set("user-agent", userAgent)
	controlUrl, err := launcher.Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().NoDefaultDevice().ControlURL(controlUrl)
	if err := browser.Connect(); err != nil {
		return nil, err
	}
	scrapper.browser = browser
	return scrapper, nil
}

func (s *Scraper) Cleanup() error {
	err := s.browser.Close()
	if err != nil {
		return fmt.Errorf("error while closing browser: %w", err)
	}

	return err
}

// todo add location
// todo debug logs
// todo goroutines for each product and for nutrition page
// todo instead of returning array return channel
// todo handle timeout => send what was found and errors for the rest
// todo parse things to floats / ints
// todo sometimes the calories are not found but there are joules instead
// todo caching?
func (s *Scraper) GetKosikProducts(search string) ([]*returnProduct, error) {
	searchUrl, err := urlParams.CreateSearchUrl(search)
	if err != nil {
		return nil, err
	}

	page, err := s.browser.Page(proto.TargetCreateTarget{
		URL: searchUrl.String(),
	})
	if err != nil {
		return nil, err
	}
	var deferErr error
	defer func() {
		err := page.Close()
		if err != nil {
			deferErr = fmt.Errorf("error failed to close page: %v", err)
		}
	}()

	err = page.WaitDOMStable(1*time.Second, 0)
	if err != nil {
		return nil, err
	}

	productSelector := "[data-tid='product-box']:not(:has(.product-amount--vendor-pharmacy))"
	products, err := page.Sleeper(rod.NotFoundSleeper).Elements(productSelector)
	if err != nil {
		return nil, err
	}

	parsedProducts := make([]*returnProduct, 0, len(products))

	log.Info().Msgf("Found %d products", len(products))

	for _, product := range products {
		parsedProduct, err := scrapeProduct(product)

		parsedProducts = append(parsedProducts, &returnProduct{ScrapeResult: structs.ScrapeResult[*Product]{
			Value:     parsedProduct,
			ScrapeErr: err,
		},
		})
	}

	return parsedProducts, deferErr
}
