package orchestrator

import (
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
	scraperLibShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
)

type fieldResult[T any] struct {
	Value T       `json:"value"`
	Error *string `json:"error,omitempty"`
}

type nutritionFields struct {
	Calories     fieldResult[*float64] `json:"calories"`
	Protein      fieldResult[*float64] `json:"protein"`
	Fat          fieldResult[*float64] `json:"fat"`
	SaturatedFat fieldResult[*float64] `json:"saturatedFat"`
	Carbs        fieldResult[*float64] `json:"carbs"`
	Sugar        fieldResult[*float64] `json:"sugar"`
	Fiber        fieldResult[*float64] `json:"fiber"`
	Ingredients  fieldResult[*string]  `json:"ingredients"`
}

type productFields struct {
	Name         fieldResult[string]                        `json:"name"`
	Price        fieldResult[float64]                       `json:"price"`
	Currency     fieldResult[string]                        `json:"currency"`
	Unit         fieldResult[string]                        `json:"unit"`
	PricePerUnit fieldResult[scraperLibShared.PricePerUnit] `json:"pricePerUnit"`
	Link         fieldResult[*string]                       `json:"link"`
	ImageURL     fieldResult[*string]                       `json:"imageUrl"`
	IsSoldOut    fieldResult[bool]                          `json:"isSoldOut"`
	Nutrition    fieldResult[*nutritionFields]              `json:"nutrition"`
}

type IngredientProducts struct {
	Ingredient shared.Ingredient `json:"ingredient"`
	Source     string            `json:"source"`
	Products   []*productFields  `json:"products"`
}
