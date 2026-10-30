package kosik

import (
	"net/url"
	"strconv"

	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/kosik/urlParams"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/convertUtils"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/htmlUtils"
)

const caloriesText = "Energetická hodnota"
const caloriesUnitText = "kcal"
const proteinText = "Bílkoviny"
const fatText = "Tuky"
const saturatedFatText = "Z toho nasycené mastné kyseliny"
const carbsText = "Sacharidy"
const sugarText = "Z toho cukry"
const fiberText = "Vláknina"
const ingredientsText = "Složení"

func transformKosikSearchProductToProduct(index int, productData *Product) *shared.ProductResult {
	if productData == nil {
		return nil
	}

	linkUrl, err := urlParams.CreateProductUrl(productData.URL)
	urlResult := shared.ScrapeResult[*url.URL]{Value: linkUrl, ScrapeErr: err}

	storefrontURL, err := urlParams.CreateStorefrontURL(productData.URL)
	storefrontResult := shared.ScrapeResult[*url.URL]{Value: storefrontURL, ScrapeErr: err}

	imageUrl, err := url.Parse(productData.Image)
	imageResult := shared.ScrapeResult[*url.URL]{Value: imageUrl, ScrapeErr: err}

	pricePerUnit := shared.PricePerUnit{
		Value: productData.PricePerUnit.Price,
		Unit:  productData.PricePerUnit.Unit,
	}

	return &shared.ProductResult{
		Index: index,
		Result: &shared.ReturnProduct{ScrapeResult: shared.ScrapeResult[*shared.Product]{
			Value: &shared.Product{
				Name:          shared.ScrapeResult[string]{Value: productData.Name},
				Price:         shared.ScrapeResult[float64]{Value: productData.Price},
				Currency:      shared.ScrapeResult[string]{Value: "CZK"},
				PricePerUnit:  shared.ScrapeResult[shared.PricePerUnit]{Value: pricePerUnit},
				Unit:          shared.ScrapeResult[string]{Value: productData.Unit},
				Link:          urlResult,
				StorefrontURL: storefrontResult,
				ImageUrl:      imageResult,
				IsSoldOut:     shared.ScrapeResult[bool]{Value: productData.MaxInCart == 0},
			},
			ScrapeErr: nil,
		}},
	}
}

func transformKosikSearchProductDetailToNutrition(detailData *ProductDetail) *shared.Nutrition {
	var calories, protein, fat, saturatedFat, carbs, sugar, fiber shared.ScrapeResult[*float64]
	var ingredients shared.ScrapeResult[*string]

	if detailData == nil {
		return nil
	}

	if detailData.NutritionalValues != nil {
		for _, nutrition := range detailData.NutritionalValues.Values {
			parsedValue, err := strconv.ParseFloat(nutrition.Value, 64)

			switch nutrition.Title {
			case caloriesText:
				if nutrition.Unit == caloriesUnitText {
					calories.Value = &parsedValue
				} else {
					kcalValue := convertUtils.KjToKcal(parsedValue)
					calories.Value = &kcalValue
				}
				calories.ScrapeErr = err
			case proteinText:
				protein.Value = &parsedValue
				protein.ScrapeErr = err
			case fatText:
				fat.Value = &parsedValue
				fat.ScrapeErr = err
			case saturatedFatText:
				saturatedFat.Value = &parsedValue
				saturatedFat.ScrapeErr = err
			case carbsText:
				carbs.Value = &parsedValue
				carbs.ScrapeErr = err
			case sugarText:
				sugar.Value = &parsedValue
				sugar.ScrapeErr = err
			case fiberText:
				fiber.Value = &parsedValue
				fiber.ScrapeErr = err
			}
		}
	}

	for _, ingredient := range detailData.Ingredients {
		if ingredient.Title == ingredientsText {
			if ingredient.Type == "html" {
				text, err := htmlUtils.ExtractTextFromHtml(ingredient.Value)
				ingredients.Value = &text
				ingredients.ScrapeErr = err

			} else {
				ingredients.Value = &ingredient.Value
				ingredients.ScrapeErr = nil
			}
			break
		}
	}

	return &shared.Nutrition{
		Calories:     calories,
		Protein:      protein,
		Fat:          fat,
		SaturatedFat: saturatedFat,
		Carbs:        carbs,
		Sugar:        sugar,
		Fiber:        fiber,
		Ingredients:  ingredients,
	}
}
