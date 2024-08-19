package scraper

import (
	"errors"
	"log"
	"net/url"
	"time"

	errorUtils "github.com/GnotAGnoob/kosik-scraper/pkg/utils/errors"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/scraping"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
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

type Product struct {
	Name       string
	Price      string
	PricePerKg string
	Unit       string
	Link       *url.URL
	// Image      *[]byte
	IsSoldOut bool
	Nutrition nutrition
	AddButton *rod.Element
}

func (product *Product) getNutritions() *[]error {
	if product.Link == nil {
		return &[]error{errors.New("Product link is not set")}
	}

	errors := []error{}
	indgredientsPage, err := scrapper.browser.Page(proto.TargetCreateTarget{
		URL: product.Link.String(),
	})
	if err != nil {
		errors = append(errors, err)
		return &errors
	}
	defer func() {
		err = indgredientsPage.Close()
		if err != nil {
			log.Fatalf("Error failed to close ingredients page: %v", err)
		}
	}()

	err = indgredientsPage.WaitDOMStable(1*time.Second, 0)
	if err != nil {
		errors = append(errors, err)
		return &errors
	}

	_, err = indgredientsPage.Sleeper(rod.NotFoundSleeper).ElementR("button", "/vyprodáno/i")
	if err == nil {
		product.IsSoldOut = true
		return &errors
	}

	ingredients, err := scraping.GetText(indgredientsPage.Sleeper(rod.NotFoundSleeper), "[data-tid='product-detail__ingredients'] dd")
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.Ingredients = ingredients

	nutritionElement, err := indgredientsPage.Sleeper(rod.NotFoundSleeper).Element("[data-tid='product-detail__nutrition_table'][aria-describedby='Výživové hodnoty (na 100 g)']")
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok { // some random error
		errors = append(errors, err)
		return &errors
	}
	if err != nil { // no nutrition table
		return &errors
	}

	calories, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "\\d* kcal", true)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.Calories = calories

	fat, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Tuky", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.Fat = fat

	saturatedFat, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Z toho nasycené mastné kyseliny", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.SaturatedFat = saturatedFat

	carbs, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Sacharidy", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.Carbs = carbs

	sugar, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Z toho cukry", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.Sugar = sugar

	fiber, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Vláknina", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.Fiber = fiber

	protein, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Bílkoviny", false)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.Nutrition.Protein = protein

	return &errors
}

func scrapeProduct(element *rod.Element) (*Product, *[]error) {
	errors := []error{}
	product := &Product{}

	_, err := element.Sleeper(rod.NotFoundSleeper).ElementR("span", "/vyprodáno/i")
	if err == nil {
		product.IsSoldOut = true
		return product, &errors
	}

	nameSelector := "a[data-tid='product-box__name']"
	nameElement, err := element.Sleeper(rod.NotFoundSleeper).Element(nameSelector)
	if err != nil {
		errors = append(errors, errorUtils.ElementNotFoundError(err, nameSelector))
	} else {
		name, err := nameElement.Text()
		if err != nil {
			errors = append(errors, err)
		}
		product.Name = name

		url := &url.URL{}
		hrefAttribute := "href"
		href, err := nameElement.Sleeper(rod.NotFoundSleeper).Attribute(hrefAttribute)
		if err != nil {
			errors = append(errors, errorUtils.ElementNotFoundError(err, hrefAttribute))
		} else {
			url, err = urlParams.CreateUrlFromPath(*href)
			if err != nil {
				errors = append(errors, err)
			}
			url.Fragment = "ingredients"
		}
		product.Link = url
	}

	unitSelector := ".attributes"
	unit, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), unitSelector)
	if err != nil {
		errors = append(errors, errorUtils.ElementNotFoundError(err, unitSelector))
	}
	product.Unit = unit

	// imgSelector := "a img"
	// image, err := scraping.GetImageResource(element.Sleeper(rod.NotFoundSleeper), imgSelector)
	// if err != nil {
	// 	errors = append(errors, errorUtils.ElementNotFoundError(err, imgSelector))
	// }
	// product.Image = image

	pricePrefix, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), ".price__prefix")
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	} else if err == nil {
		pricePrefix += " "
	}

	priceSelector := "[data-tid='product-price']"
	price, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), priceSelector)
	if err != nil {
		errors = append(errors, errorUtils.ElementNotFoundError(err, priceSelector))
	}
	product.Price = pricePrefix + price

	pricePerKgSelector := "[aria-label='Cena'] > *:last-child"
	pricePerKg, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), pricePerKgSelector)
	if err != nil {
		errors = append(errors, errorUtils.ElementNotFoundError(err, pricePerKgSelector))
	}
	product.PricePerKg = pricePerKg

	buttonSelector := "[data-tid='product-to-cart__to-cart']"
	buttonElement, err := element.Sleeper(rod.NotFoundSleeper).Element(buttonSelector)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	}
	product.AddButton = buttonElement

	ingredientsErrors := product.getNutritions()
	errors = append(errors, *ingredientsErrors...)

	return product, &errors
}
