package kosik

import (
	"net/url"
	"strconv"

	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/kosik/urlParams"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
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
	urlResult := structs.ScrapeResult[*url.URL]{Value: linkUrl, ScrapeErr: err}

	imageUrl, err := url.Parse(productData.Image)
	imageResult := structs.ScrapeResult[*url.URL]{Value: imageUrl, ScrapeErr: err}

	pricePerUnit := shared.PricePerUnit{
		Value: productData.PricePerUnit.Price,
		Unit:  productData.PricePerUnit.Unit,
	}

	return &shared.ProductResult{
		Index: index,
		Result: &shared.ReturnProduct{ScrapeResult: structs.ScrapeResult[*shared.Product]{
			Value: &shared.Product{
				Name:         structs.ScrapeResult[string]{Value: productData.Name},
				Price:        structs.ScrapeResult[float64]{Value: productData.Price},
				PricePerUnit: structs.ScrapeResult[shared.PricePerUnit]{Value: pricePerUnit},
				Unit:         structs.ScrapeResult[string]{Value: productData.Unit},
				Link:         urlResult,
				ImageUrl:     imageResult,
				IsSoldOut:    structs.ScrapeResult[bool]{Value: productData.MaxInCart == 0},
			},
			ScrapeErr: nil,
		}},
	}
}

func transformKosikSearchProductDetailToNutrition(detailData *ProductDetail) *shared.Nutrition {
	var calories, protein, fat, saturatedFat, carbs, sugar, fiber structs.ScrapeResult[*float64]
	var ingredients structs.ScrapeResult[*string]

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
