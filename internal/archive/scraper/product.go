package scraper

// import (
// 	"errors"
// 	"fmt"
// 	"net/url"
// 	"strings"
// 	"sync"

// 	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
// 	"github.com/GnotAGnoob/kosik-scraper/internal/utils/urlParams"
// 	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/scraping"
// 	"github.com/go-rod/rod"
// 	"github.com/go-rod/rod/lib/proto"
// )

// type nutrition struct {
// 	Calories     structs.ScrapeResult[string]
// 	Protein      structs.ScrapeResult[string]
// 	Fat          structs.ScrapeResult[string]
// 	SaturatedFat structs.ScrapeResult[string]
// 	Carbs        structs.ScrapeResult[string]
// 	Sugar        structs.ScrapeResult[string]
// 	Fiber        structs.ScrapeResult[string]
// 	Ingredients  structs.ScrapeResult[string]
// }

// type Product struct {
// 	Name       structs.ScrapeResult[string]
// 	Price      structs.ScrapeResult[string]
// 	PricePerKg structs.ScrapeResult[string]
// 	Unit       structs.ScrapeResult[string]
// 	Link       structs.ScrapeResult[*url.URL]
// 	// Image      *[]byte
// 	IsSoldOut structs.ScrapeResult[bool]
// 	Nutrition structs.ScrapeResult[*nutrition]
// 	AddButton structs.ScrapeResult[*rod.Element]
// }

// func (product *Product) scrapeNutritions(browser *rod.Browser) error {
// 	if product == nil {
// 		return errors.New("product is not set")
// 	}

// 	if browser == nil {
// 		return errors.New("browser is not set")
// 	}

// 	nutrition := &nutrition{}
// 	product.Nutrition.Value = nutrition

// 	if product.Link.ScrapeErr != nil || product.Link.Value == nil || len(product.Link.Value.String()) == 0 {
// 		return errors.New("product link is not set")
// 	}

// 	ingredientsPage, err := browser.Page(proto.TargetCreateTarget{
// 		URL: product.Link.Value.String(),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		err = ingredientsPage.Close()
// 		if err != nil {
// 			err = fmt.Errorf("error failed to close ingredients page: %w", err)
// 		}
// 	}()

// 	isNotFoundPage := false
// 	notFoundHandler := func(el *rod.Element) error {
// 		isNotFoundPage = true
// 		return nil
// 	}

// 	_, err = scraping.RaceSelectors(
// 		ingredientsPage,
// 		nutritionPageTimeout,
// 		scraping.RaceSelector{Selector: nutritionPageNotFoundWaitSelector, Handler: &notFoundHandler},
// 		scraping.RaceSelector{Selector: nutritionPageWaitSelector},
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	if isNotFoundPage {
// 		return nil
// 	}

// 	_, err = ingredientsPage.Sleeper(rod.NotFoundSleeper).ElementR("button", "/vyprodáno/i")
// 	if err == nil {
// 		product.IsSoldOut.Value = true
// 		return nil
// 	}

// 	ingredients, err := scraping.GetText(ingredientsPage.Sleeper(rod.NotFoundSleeper), ingredientsSelector)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.Ingredients.ScrapeErr = err
// 	}
// 	nutrition.Ingredients.Value = ingredients

// 	nutritionElement, err := ingredientsPage.Sleeper(rod.NotFoundSleeper).Element(nutritionSelector)
// 	if scraping.IsElementNotFound(err) { // some random error
// 		return err
// 	}
// 	if err != nil { // no nutrition table
// 		return nil
// 	}

// 	calories, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "\\d* kcal", true)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.Calories.ScrapeErr = err
// 	}
// 	nutrition.Calories.Value = calories

// 	fat, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Tuky", false)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.Fat.ScrapeErr = err
// 	}
// 	nutrition.Fat.Value = fat

// 	saturatedFat, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Z toho nasycené mastné kyseliny", false)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.SaturatedFat.ScrapeErr = err
// 	}
// 	nutrition.SaturatedFat.Value = saturatedFat

// 	carbs, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Sacharidy", false)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.Carbs.ScrapeErr = err
// 	}
// 	nutrition.Carbs.Value = carbs

// 	sugar, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Z toho cukry", false)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.Sugar.ScrapeErr = err
// 	}
// 	nutrition.Sugar.Value = sugar

// 	fiber, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Vláknina", false)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.Fiber.ScrapeErr = err
// 	}
// 	nutrition.Fiber.Value = fiber

// 	protein, err := scraping.GetTextFromTable(nutritionElement.Sleeper(rod.NotFoundSleeper), "Bílkoviny", false)
// 	if scraping.IsElementNotFound(err) {
// 		nutrition.Protein.ScrapeErr = err
// 	}
// 	nutrition.Protein.Value = protein

// 	return err
// }

// func scrapeProduct(element *rod.Element, browser *rod.Browser) (*Product, error) {
// 	if element == nil {
// 		return nil, errors.New("product is not set")
// 	}

// 	if browser == nil {
// 		return nil, errors.New("browser is not set")
// 	}

// 	product := &Product{}

// 	linkElement, err := element.Sleeper(rod.NotFoundSleeper).Element(linkSelector)
// 	if err != nil {
// 		product.Link.ScrapeErr = err
// 	} else {
// 		url := &url.URL{}
// 		hrefAttribute := "href"
// 		href, err := linkElement.Sleeper(rod.NotFoundSleeper).Attribute(hrefAttribute)
// 		if err != nil {
// 			product.Link.ScrapeErr = err
// 		} else if href != nil {
// 			url, err = urlParams.CreateUrlFromPath(*href)
// 			if err != nil {
// 				product.Link.ScrapeErr = err
// 			}
// 			url.Fragment = "ingredients"
// 		}
// 		product.Link.Value = url
// 	}

// 	name, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), nameSelector)
// 	if err != nil {
// 		product.Name.ScrapeErr = err
// 	}
// 	product.Name.Value = name

// 	_, err = element.Sleeper(rod.NotFoundSleeper).ElementR("span", "/vyprodáno/i")
// 	if err == nil {
// 		product.IsSoldOut.Value = true
// 		return product, nil
// 	}

// 	ingredientsError := product.scrapeNutritions(browser)
// 	product.Nutrition.ScrapeErr = ingredientsError

// 	// todo remove whitespace
// 	unit, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), unitSelector)
// 	if err != nil {
// 		product.Unit.ScrapeErr = err
// 	}
// 	product.Unit.Value = unit

// 	// imgSelector := "a img"
// 	// image, err := scraping.GetImageResource(element.Sleeper(rod.NotFoundSleeper), imgSelector)
// 	// if err != nil {
// 	// 	errors = append(errors, errorUtils.ElementNotFoundError(err, imgSelector))
// 	// }
// 	// product.Image = image

// 	pricePrefix, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), pricePrefixSelector)
// 	_, ok := err.(*rod.ElementNotFoundError)
// 	if err != nil && !ok {
// 		product.Price.ScrapeErr = err
// 	}

// 	price, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), priceSelector)
// 	if err != nil {
// 		product.Price.ScrapeErr = err
// 	}
// 	product.Price.Value = strings.TrimSpace(pricePrefix + " " + price)

// 	pricePerKg, err := scraping.GetText(element.Sleeper(rod.NotFoundSleeper), pricePerKgSelector)
// 	if err != nil {
// 		product.PricePerKg.ScrapeErr = err
// 	}
// 	product.PricePerKg.Value = pricePerKg

// 	buttonElement, err := element.Sleeper(rod.NotFoundSleeper).Element(buttonSelector)
// 	if scraping.IsElementNotFound(err) {
// 		product.AddButton.ScrapeErr = err
// 	}
// 	product.AddButton.Value = buttonElement

// 	return product, nil
// }

// func scrapeProductAsync(product *rod.Element, index int, browser *rod.Browser, ch chan<- *ProductResult, wg *sync.WaitGroup) {
// 	defer func() {
// 		wg.Done()
// 	}()

// 	if product == nil {
// 		ch <- &ProductResult{
// 			Index: index,
// 			Result: &ReturnProduct{ScrapeResult: structs.ScrapeResult[*Product]{
// 				Value:     nil,
// 				ScrapeErr: errors.New("product is not set"),
// 			}},
// 		}
// 		return
// 	}

// 	if browser == nil {
// 		ch <- &ProductResult{
// 			Index: index,
// 			Result: &ReturnProduct{ScrapeResult: structs.ScrapeResult[*Product]{
// 				Value:     nil,
// 				ScrapeErr: errors.New("browser is not set"),
// 			}},
// 		}
// 		return
// 	}

// 	parsedProduct, err := scrapeProduct(product, browser)

// 	ch <- &ProductResult{
// 		Index: index,
// 		Result: &ReturnProduct{ScrapeResult: structs.ScrapeResult[*Product]{
// 			Value:     parsedProduct,
// 			ScrapeErr: err,
// 		}},
// 	}
// }
