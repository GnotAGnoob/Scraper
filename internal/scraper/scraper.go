package scraper

import (
	"fmt"
	"log"
	"net/url"

	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/errorsUtils"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type nutrition struct {
	Calories     string
	Protein      string
	Fat          string
	SaturatedFat string
	Carbs        string
	Sugar        string
	Fiber        string
	Ingredients  string
}

type ParsedProduct struct {
	Name       string
	Price      string
	PricePerKg string
	Unit       string
	Link       *url.URL
	Image      *[]byte
	IsSoldOut  bool
	Nutrition  *nutrition
	AddButton  *rod.Element
}

type returnProduct struct {
	Product *ParsedProduct
	Errors  *[]error
}

type Scraper struct {
	browser *rod.Browser
}

var scrapper = &Scraper{}

func InitScraper() (*Scraper, error) {
	if scrapper.browser != nil {
		return scrapper, nil
	}

	// leakless is a binary that prevents zombie processes
	// but the problem is that windows defender detects it as a virus
	// because according to internet, it is used in many viruses
	launcher := launcher.New().Leakless(false)
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

// todo separate the code into smaller reusable functions
// todo handle timeout => send what was found and errors for the rest
// todo goroutines for each product and for nutrition page
// todo debug mode with own logging
// todo parse things to floats / ints
func (s *Scraper) GetKosikProducts(search string) (*[]*returnProduct, error) {
	searchUrl, err := urlParams.CreateSearchUrl(search)
	if err != nil {
		return nil, err
	}

	fmt.Println("searchUrl: ", searchUrl)

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

	widgetSelector := ".page-products-widgets"
	err = page.WaitElementsMoreThan(widgetSelector, 0)
	if err != nil {
		return nil, errorsUtils.ElementNotFoundError(err, widgetSelector)
	}

	productSelector := "[data-tid='product-box']:not(:has(.product-amount--vendor-pharmacy))"
	products, err := page.Elements(productSelector)
	if err != nil {
		return nil, errorsUtils.ElementNotFoundError(err, productSelector)
	}

	parsedProducts := make([]*returnProduct, 0, len(products))

	for _, product := range products[:2] {
		// todo put into separate function and defer when soldout
		errors := []error{}
		parsedProduct := &ParsedProduct{}

		_, err := product.ElementR("span", "/vyprodáno/i")
		if err == nil {
			parsedProduct.IsSoldOut = true
		}

		nameSelector := "a[data-tid='product-box__name']"
		nameElement, err := product.Sleeper(rod.NotFoundSleeper).Element(nameSelector)
		if err != nil {
			errors = append(errors, errorsUtils.ElementNotFoundError(err, nameSelector))
		} else {
			name, err := nameElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Name = name

			url := &url.URL{}
			hrefAttribute := "href"
			href, err := nameElement.Attribute(hrefAttribute)
			if err != nil {
				errors = append(errors, errorsUtils.ElementNotFoundError(err, hrefAttribute))
			} else {
				url, err = urlParams.CreateUrlFromPath(*href)
				if err != nil {
					errors = append(errors, err)
				}
				url.Fragment = "ingredients"
			}
			parsedProduct.Link = url
		}

		unitSelector := ".attributes"
		unitElement, err := product.Sleeper(rod.NotFoundSleeper).Element(unitSelector)
		if err != nil {
			errors = append(errors, errorsUtils.ElementNotFoundError(err, unitSelector))
		} else {
			unit, err := unitElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Unit = unit
		}

		imageSelector := "img"
		imageElement, err := product.Sleeper(rod.NotFoundSleeper).Element(imageSelector)
		if err != nil {
			errors = append(errors, errorsUtils.ElementNotFoundError(err, imageSelector))
		} else {
			image, err := imageElement.Resource()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Image = &image
		}

		pricePrefix := ""
		pricePrefixElement, err := product.Sleeper(rod.NotFoundSleeper).Element(".price__prefix")
		if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
			errors = append(errors, err)
		} else if err == nil {
			pricePrefix, err = pricePrefixElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			pricePrefix += " "
		}

		priceSelector := "[data-tid='product-price']"
		priceElement, err := product.Sleeper(rod.NotFoundSleeper).Element(priceSelector)
		if err != nil {
			errors = append(errors, errorsUtils.ElementNotFoundError(err, priceSelector))
		} else {
			price, err := priceElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Price = pricePrefix + price
		}

		pricePerKgSelector := "[aria-label='Cena'] > *:last-child"
		pricePerKgElement, err := product.Sleeper(rod.NotFoundSleeper).Element(pricePerKgSelector)
		if err != nil {
			errors = append(errors, errorsUtils.ElementNotFoundError(err, pricePerKgSelector))
		} else {
			pricePerKg, err := pricePerKgElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.PricePerKg = pricePerKg
		}

		buttonSelector := "[data-tid='product-to-cart__to-cart']"
		buttonElement, err := product.Sleeper(rod.NotFoundSleeper).Element(buttonSelector)
		if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
			errors = append(errors, err)
		}
		parsedProduct.AddButton = buttonElement

		ingredientsErrors := getIngredients(parsedProduct)
		errors = append(errors, *ingredientsErrors...)

		parsedProducts = append(parsedProducts, &returnProduct{
			Product: parsedProduct,
			Errors:  &errors,
		})
	}

	return &parsedProducts, nil
}
