package scraper

import (
	"log"

	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/errorsUtils"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func getNutrition(element *rod.Element, searchText string, isDirect bool) (string, error) {
	saturatedFatElement, err := element.ElementR("td", searchText)
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok { // some random error
		return "", err
	}

	if err != nil { // element not found
		return "", nil
	}

	if !isDirect {
		saturatedFatElement, err = saturatedFatElement.Next()
		if err != nil {
			return "", err
		}
	}

	saturatedFat, err := saturatedFatElement.Text()
	if err != nil {
		return "", err
	}

	return saturatedFat, nil
}

func getIngredients(product *ParsedProduct) *[]error {
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

	imgSelector := "img"
	err = indgredientsPage.WaitElementsMoreThan(imgSelector, 0) // there is always an image. Wait until javascript loads it
	if err != nil {
		errors = append(errors, errorsUtils.ElementNotFoundError(err, imgSelector))
		return &errors
	}

	_, err = indgredientsPage.ElementR("button", "/vyprodáno/i")
	if err == nil {
		product.IsSoldOut = true
		return &errors
	}

	ingredientsElement, err := indgredientsPage.Sleeper(rod.NotFoundSleeper).Element("[data-tid='product-detail__ingredients'] dd")
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
		errors = append(errors, err)
	} else if err == nil {
		ingredients, err := ingredientsElement.Text()
		if err != nil {
			errors = append(errors, err)
		}
		product.Nutrition.Ingredients = ingredients
	}

	nutritionElement, err := indgredientsPage.Sleeper(rod.NotFoundSleeper).Element("[data-tid='product-detail__nutrition_table'][aria-describedby='Výživové hodnoty (na 100 g)']")
	if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok { // some error
		errors = append(errors, err)
		return &errors
	}
	if err != nil { // no nutrition table
		return &errors
	}

	calories, err := getNutrition(nutritionElement, "\\d* kcal", true)
	if err != nil {
		errors = append(errors, err)
	}
	product.Nutrition.Calories = calories

	fat, err := getNutrition(nutritionElement, "Tuky", false)
	if err != nil {
		errors = append(errors, err)
	}
	product.Nutrition.Fat = fat

	saturatedFat, err := getNutrition(nutritionElement, "Z toho nasycené mastné kyseliny", false)
	if err != nil {
		errors = append(errors, err)
	}
	product.Nutrition.SaturatedFat = saturatedFat

	carbs, err := getNutrition(nutritionElement, "Sacharidy", false)
	if err != nil {
		errors = append(errors, err)
	}
	product.Nutrition.Carbs = carbs

	sugar, err := getNutrition(nutritionElement, "Z toho cukry", false)
	if err != nil {
		errors = append(errors, err)
	}
	product.Nutrition.Sugar = sugar

	fiber, err := getNutrition(nutritionElement, "Vláknina", false)
	if err != nil {
		errors = append(errors, err)
	}
	product.Nutrition.Fiber = fiber

	protein, err := getNutrition(nutritionElement, "Bílkoviny", false)
	if err != nil {
		errors = append(errors, err)
	}
	product.Nutrition.Protein = protein

	return &errors
}
