package scraper

// import (
// 	"fmt"
// 	"sync"
// 	"time"

// 	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
// 	"github.com/GnotAGnoob/kosik-scraper/internal/utils/urlParams"
// 	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/scraping"
// 	"github.com/go-rod/rod"
// 	"github.com/go-rod/rod/lib/launcher"
// 	"github.com/go-rod/rod/lib/proto"
// 	"github.com/rs/zerolog/log"
// )

// type Scraper struct {
// 	browser *rod.Browser
// }

// var scrapper = &Scraper{}

// type ReturnProduct struct {
// 	structs.ScrapeResult[*Product]
// }

// type ProductResult struct {
// 	Index  int
// 	Result *ReturnProduct
// }

// func InitScraper() (*Scraper, error) {
// 	if scrapper.browser != nil {
// 		return scrapper, nil
// 	}

// 	// leakless is a binary that prevents zombie processes
// 	// but the problem is that windows defender detects it as a virus
// 	// because according to internet, it is used in many viruses
// 	launcher := launcher.New().Leakless(false).Set("user-agent", userAgent)
// 	controlUrl, err := launcher.Launch()
// 	if err != nil {
// 		return nil, err
// 	}

// 	browser := rod.New().NoDefaultDevice().ControlURL(controlUrl)
// 	if err := browser.Connect(); err != nil {
// 		return nil, err
// 	}
// 	scrapper.browser = browser
// 	return scrapper, nil
// }

// func (s *Scraper) Cleanup() error {
// 	err := s.browser.Close()
// 	if err != nil {
// 		return fmt.Errorf("error while closing browser: %w", err)
// 	}

// 	s.browser = nil

// 	return err
// }

// // todo add the setting of the address for kosik site
// // todo debug logs
// // todo handle timeout => send what was found and errors for the rest
// // todo parse things to floats / ints
// // todo sometimes the calories are not found but there are joules instead
// // todo caching?
// func (s *Scraper) GetKosikProducts(search string, totalChan chan<- int, productsChan chan<- *ProductResult) error {
// 	defer func() {
// 		close(totalChan)
// 		close(productsChan)
// 	}()

// 	searchUrl, err := urlParams.CreateSearchUrl(search)
// 	if err != nil {
// 		return err
// 	}

// 	page, err := s.browser.Page(proto.TargetCreateTarget{
// 		URL: searchUrl.String(),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		err = page.Close()
// 		if err != nil {
// 			err = fmt.Errorf("error failed to close page: %w", err)
// 		}
// 	}()

// 	isNotFoundPage := false
// 	notFoundHandler := func(el *rod.Element) error {
// 		isNotFoundPage = true
// 		return nil
// 	}

// 	_, err = scraping.RaceSelectors(
// 		page,
// 		productsPageTimeout,
// 		scraping.RaceSelector{Selector: productPageNotFoundWaitSelector, Handler: &notFoundHandler},
// 		scraping.RaceSelector{Selector: productPageWaitSelector},
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if isNotFoundPage {
// 		totalChan <- 0
// 		return err
// 	}

// 	productSelector := "[data-tid='product-box']:not(:has(.product-amount--vendor-pharmacy))"
// 	products, err := page.Sleeper(rod.NotFoundSleeper).Elements(productSelector)
// 	if err != nil {
// 		return err
// 	}

// 	log.Debug().Msgf("Found %d products", len(products))

// 	totalChan <- len(products)

// 	if len(products) == 0 {
// 		return err
// 	}

// 	browser := s.browser.Timeout(time.Duration(len(products)) * perProductTimeout)

// 	wg := sync.WaitGroup{}

// 	for index, product := range products {
// 		wg.Add(1)
// 		go scrapeProductAsync(product, index, browser, productsChan, &wg)
// 	}

// 	wg.Wait()

// 	return err
// }
