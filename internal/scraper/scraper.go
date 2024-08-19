package scraper

import (
	"fmt"
	"log"
	"time"

	errorUtils "github.com/GnotAGnoob/kosik-scraper/pkg/utils/errors"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type Scraper struct {
	browser *rod.Browser
}

var scrapper = &Scraper{}

type returnProduct struct {
	Product *Product
	Errors  *[]error
}

func InitScraper() (*Scraper, error) {
	if scrapper.browser != nil {
		return scrapper, nil
	}

	// leakless is a binary that prevents zombie processes
	// but the problem is that windows defender detects it as a virus
	// because according to internet, it is used in many viruses
	launcher := launcher.New().Leakless(false).Set("user-agent", USER_AGENT)
	controlUrl, err := launcher.Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().NoDefaultDevice().ControlURL(controlUrl)
	if error := browser.Connect(); error != nil {
		return nil, error
	}
	scrapper.browser = browser
	return scrapper, nil
}

func (s *Scraper) Cleanup() {
	err := s.browser.Close()
	if err != nil {
		log.Fatalf("Error failed to close browser: %v", err)
	}
}

// todo add location
// todo debug mode with own logging
// todo goroutines for each product and for nutrition page
// todo handle timeout => send what was found and errors for the rest
// todo parse things to floats / ints
// todo caching?
func (s *Scraper) GetKosikProducts(search string) (*[]*returnProduct, error) {
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
	defer func() {
		err := page.Close()
		if err != nil {
			log.Fatalf("Error failed to close page: %v", err)
		}
	}()

	err = page.WaitDOMStable(1*time.Second, 0)
	if err != nil {
		return nil, err
	}

	productSelector := "[data-tid='product-box']:not(:has(.product-amount--vendor-pharmacy))"
	products, err := page.Sleeper(rod.NotFoundSleeper).Elements(productSelector)
	if err != nil {
		return nil, errorUtils.ElementNotFoundError(err, productSelector)
	}

	parsedProducts := make([]*returnProduct, 0, len(products))

	fmt.Printf("Found %d products\n", len(products))

	for _, product := range products {
		parsedProduct, errors := scrapeProduct(product)

		parsedProducts = append(parsedProducts, &returnProduct{
			Product: parsedProduct,
			Errors:  errors,
		})
	}

	return &parsedProducts, nil
}
