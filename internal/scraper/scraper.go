package scraper

import (
	"fmt"
	"log"

	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type ParsedProduct struct {
	Name       string
	Price      string
	PricePerKg string
	Unit       string
	Link       string
	Image      *[]byte
	AddButton  *rod.Element
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

	err = page.WaitElementsMoreThan("[data-product-id]", 1)
	if err != nil {
		return nil, err
	}

	products, err := page.Elements("[data-product-id]")
	if err != nil {
		return nil, err
	}

	parsedProducts := make([]*ParsedProduct, 0, len(products))

	for _, product := range products {
		nameElement, err := product.Element("a[data-tid='product-box__name']")
		if err != nil {
			return nil, err
		}

		name, err := nameElement.Text()
		if err != nil {
			return nil, err
		}

		href, err := nameElement.Attribute("href")
		if err != nil {
			return nil, err
		}

		unitElement, err := product.Element(".attributes")
		if err != nil {
			return nil, err
		}

		unit, err := unitElement.Text()
		if err != nil {
			return nil, err
		}

		imageElement, err := product.Element("img")
		if err != nil {
			return nil, err
		}

		image, err := imageElement.Resource()
		if err != nil {
			return nil, err
		}

		pricePrefixElement, err := product.Element(".price__prefix")
		if err != nil {
			return nil, err
		}

		pricePrefix, err := pricePrefixElement.Text()
		err, ok := err.(*rod.ObjectNotFoundError)
		if !ok {
			return nil, err
		}

		priceElement, err := product.Element("[data-tid='product-price']")
		if err != nil {
			return nil, err
		}

		price, err := priceElement.Text()
		if err != nil {
			return nil, err
		}

		pricePerKgElement, err := product.Element("[aria-label='Cena'] > *:last-child")
		if err != nil {
			return nil, err
		}

		pricePerKg, err := pricePerKgElement.Text()
		if err != nil {
			return nil, err
		}

		buttonElement, err := product.Element("[data-tid='product-to-cart__to-cart']")
		if err != nil {
			return nil, err
		}

		parsedProducts = append(parsedProducts, &ParsedProduct{
			Name:       name,
			Price:      fmt.Sprintf("%s %s", pricePrefix, price),
			PricePerKg: pricePerKg,
			Unit:       unit,
			Link:       *href,
			Image:      &image,
			AddButton:  buttonElement,
		})
	}

	return parsedProducts, nil
}
