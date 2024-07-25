package scraper

import (
	"fmt"
	"log"

	errorsUtil "github.com/GnotAGnoob/kosik-scraper/pkg/utils/errors"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type ParsedProduct struct {
	Name         string
	Price        string
	PricePerKg   string
	Unit         string
	Link         string
	Image        *[]byte
	Description  string
	Calories     string
	Protein      string
	Fat          string
	SaturatedFat string
	Carbs        string
	Sugar        string
	Fiber        string
	AddButton    *rod.Element
}

type Scraper struct {
	launcher *launcher.Launcher
}

func NewScraper() *Scraper {
	return &Scraper{
		// leakless is a binary that prevents zombie processes
		// but the problem is that windows defender detects it as a virus
		// because according to internet, it is used in many viruses
		launcher: launcher.New().Leakless(false),
	}
}

// todo better error messages -> now it just says: "Error: cannot find element"
// todo separate the code into smaller functions
func (s *Scraper) GetKosikItems(search string) ([]*ParsedProduct, error) {
	searchUrl, err := urlParams.GetUrl(search)
	if err != nil {
		return nil, err
	}

	controlUrl, err := s.launcher.Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().NoDefaultDevice().ControlURL(controlUrl)

	if error := browser.Connect(); error != nil {
		return nil, error
	}
	defer func() {
		err := browser.Close()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

	}()

	page, err := browser.Page(proto.TargetCreateTarget{
		URL: searchUrl.String(),
	})
	if err != nil {
		return nil, err
	}

	productSelector := "[data-tid='product-box']"
	err = page.WaitElementsMoreThan(productSelector, 1)
	if err != nil {
		return nil, errorsUtil.ElementNotFoundError(err, productSelector)
	}

	products, err := page.Elements(productSelector)
	if err != nil {
		return nil, errorsUtil.ElementNotFoundError(err, productSelector)
	}

	parsedProducts := make([]*ParsedProduct, 0, len(products))

	for _, product := range products {
		nameSelector := "a[data-tid='product-box__name']"
		nameElement, err := product.Element(nameSelector)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, nameSelector)
		}

		name, err := nameElement.Text()
		if err != nil {
			return nil, err
		}

		hrefAttribute := "href"
		href, err := nameElement.Attribute(hrefAttribute)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, hrefAttribute)
		}

		unitSelector := ".attributes"
		unitElement, err := product.Element(unitSelector)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, unitSelector)
		}

		unit, err := unitElement.Text()
		if err != nil {
			return nil, err
		}

		imageSelector := "img"
		imageElement, err := product.Element(imageSelector)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, imageSelector)
		}

		image, err := imageElement.Resource()
		if err != nil {
			return nil, err
		}

		pricePrefix := ""
		pricePrefixElement, err := product.Element(".price__prefix")
		if _, ok := err.(*rod.ElementNotFoundError); !ok {
			return nil, err
		} else if err == nil {
			pricePrefix, err = pricePrefixElement.Text()
			if err != nil {
				return nil, err
			}

			pricePrefix += " "
		}

		priceSelector := "[data-tid='product-price']"
		priceElement, err := product.Element(priceSelector)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, priceSelector)
		}

		price, err := priceElement.Text()
		if err != nil {
			return nil, err
		}

		pricePerKgSelector := "[aria-label='Cena'] > *:last-child"
		pricePerKgElement, err := product.Element(pricePerKgSelector)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, pricePerKgSelector)
		}

		pricePerKg, err := pricePerKgElement.Text()
		if err != nil {
			return nil, err
		}

		buttonSelector := "[data-tid='product-to-cart__to-cart']"
		buttonElement, err := product.Element(buttonSelector)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, buttonSelector)
		}

		fmt.Println(*href)

		indgredientsPage, err := browser.Page(proto.TargetCreateTarget{
			URL: *href + "#ingredients",
		})
		if err != nil {
			return nil, err
		}
		caloriesRegex := "\\d* kcal"
		caloriesElement, err := indgredientsPage.ElementR("td", caloriesRegex)
		if err != nil {
			return nil, errorsUtil.ElementNotFoundError(err, caloriesRegex)
		}

		fmt.Println("calories")

		calories, err := caloriesElement.Text()
		if err != nil {
			return nil, err
		}

		fmt.Println(calories)

		parsedProducts = append(parsedProducts, &ParsedProduct{
			Name:       name,
			Price:      fmt.Sprintf("%s%s", pricePrefix, price),
			PricePerKg: pricePerKg,
			Unit:       unit,
			Link:       *href,
			Image:      &image,
			Calories:   calories,
			// Protein      :
			// Fat          :
			// SaturatedFat :
			// Carbs        :
			// Sugar        :
			// Fiber        :
			AddButton: buttonElement,
		})
	}

	return parsedProducts, nil
}
