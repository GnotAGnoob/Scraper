package scraper

import (
	"errors"
	"log"
	"net/url"
	"time"

	errorUtils "github.com/GnotAGnoob/kosik-scraper/pkg/utils/errors"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/scraping"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/structs"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type nutrition struct {
	Calories     structs.ErrorValue[string]
	Protein      structs.ErrorValue[string]
	Fat          structs.ErrorValue[string]
	SaturatedFat structs.ErrorValue[string]
	Carbs        structs.ErrorValue[string]
	Sugar        structs.ErrorValue[string]
	Fiber        structs.ErrorValue[string]
	Ingredients  structs.ErrorValue[string]
}

type Product struct {
	Name       structs.ErrorValue[string]
	Price      structs.ErrorValue[string]
	PricePerKg structs.ErrorValue[string]
	Unit       structs.ErrorValue[string]
	Link       structs.ErrorValue[*url.URL]
	// Image      *[]byte
	IsSoldOut bool
	Nutrition structs.ErrorValue[nutrition]
	AddButton structs.ErrorValue[*rod.Element]
}

func (product *Product) scrapeNutritions() error {
	if product.Link.Err != nil || product.Link.Value == nil || len(product.Link.Value.String()) == 0 {
		return errors.New("Product link is not set")
	}

	indgredientsPage, err := scrapper.browser.Page(proto.TargetCreateTarget{
		URL: product.Link.Value.String(),
	})
	if err != nil {
		return err
	}
	defer func() {
		err = indgredientsPage.Close()
		if err != nil {
			log.Fatalf("Error failed to close ingredients page: %v", err)
		}
	}()

	err = indgredientsPage.WaitDOMStable(1*time.Second, 0)
	if err != nil {
		return err
	}

	_, err = indgredientsPage.Sleeper(rod.NotFoundSleeper).ElementR("button", "/vyprodáno/i")
	if err == nil {
		product.IsSoldOut = true
		return nil
	}

	nutrition := nutrition{}
	ingredients, err := scraping.GetText(indgredientsPage.Sleeper(rod.NotFoundSleeper), "[data-tid='product-detail__ingredients'] dd")
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.Ingredients.Err = err
	}
	nutrition.Ingredients.Value = ingredients

	nutritionElement, err := indgredientsPage.Sleeper(rod.NotFoundSleeper).Element("[data-tid='product-detail__nutrition_table'][aria-describedby='Výživové hodnoty (na 100 g)']")
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok { // some random error
		return err
	}
	if err != nil { // no nutrition table
		return nil
	}

	calories, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "\\d* kcal", true)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.Calories.Err = err
	}
	nutrition.Calories.Value = calories

	fat, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Tuky", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.Fat.Err = err
	}
	nutrition.Fat.Value = fat

	saturatedFat, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Z toho nasycené mastné kyseliny", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.SaturatedFat.Err = err
	}
	nutrition.SaturatedFat.Value = saturatedFat

	carbs, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Sacharidy", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.Carbs.Err = err
	}
	nutrition.Carbs.Value = carbs

	sugar, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Z toho cukry", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.Sugar.Err = err
	}
	nutrition.Sugar.Value = sugar

	fiber, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Vláknina", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.Fiber.Err = err
	}
	nutrition.Fiber.Value = fiber

	protein, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Bílkoviny", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		nutrition.Protein.Err = err
	}
	nutrition.Protein.Value = protein

	return nil
}

func scrapeProduct(element *rod.Element) (*Product, error) {
	product := &Product{}

	_, err := element.Sleeper(rod.NotFoundSleeper).ElementR("span", "/vyprodáno/i")
	if err == nil {
		product.IsSoldOut = true
		return product, nil
	}

	nameSelector := "a[data-tid='product-box__name']"
	nameElement, err := element.Sleeper(rod.NotFoundSleeper).Element(nameSelector)
	if err != nil {
		product.Name.Err = errorUtils.ElementNotFoundError(err, nameSelector)
	} else {
		name, err := nameElement.Text()
		if err != nil {
			product.Name.Err = err
		}
		product.Name.Value = name

		url := &url.URL{}
		hrefAttribute := "href"
		href, err := nameElement.Sleeper(rod.NotFoundSleeper).Attribute(hrefAttribute)
		if err != nil {
			product.Link.Err = errorUtils.ElementNotFoundError(err, hrefAttribute)
		} else {
			url, err = urlParams.CreateUrlFromPath(*href)
			if err != nil {
				product.Link.Err = err
			}
			url.Fragment = "ingredients"
		}
		product.Link.Value = url
	}

	unitSelector := ".attributes"
	unit, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), unitSelector)
	if err != nil {
		product.Unit.Err = errorUtils.ElementNotFoundError(err, unitSelector)
	}
	product.Unit.Value = unit

	// imgSelector := "a img"
	// image, err := scraping.GetImageResource(element.Sleeper(rod.NotFoundSleeper), imgSelector)
	// if err != nil {
	// 	errors = append(errors, errorUtils.ElementNotFoundError(err, imgSelector))
	// }
	// product.Image = image

	pricePrefix, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), ".price__prefix")
	_, ok := err.(*rod.ElementNotFoundError)
	if err != nil && !ok {
		product.Price.Err = err
	} else if err == nil {
		pricePrefix += " "
	}

	if err == nil || err != nil && ok {
		priceSelector := "[data-tid='product-price']"
		price, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), priceSelector)
		if err != nil {
			product.Price.Err = errorUtils.ElementNotFoundError(err, priceSelector)
		}
		product.Price.Value = pricePrefix + price
	}

	pricePerKgSelector := "[aria-label='Cena'] > *:last-child"
	pricePerKg, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), pricePerKgSelector)
	if err != nil {
		product.PricePerKg.Err = errorUtils.ElementNotFoundError(err, pricePerKgSelector)
	}
	product.PricePerKg.Value = pricePerKg

	buttonSelector := "[data-tid='product-to-cart__to-cart']"
	buttonElement, err := element.Sleeper(rod.NotFoundSleeper).Element(buttonSelector)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		product.AddButton.Err = err
	}
	product.AddButton.Value = buttonElement

	ingredientsError := product.scrapeNutritions()
	product.Nutrition.Err = ingredientsError

	return product, nil
}
