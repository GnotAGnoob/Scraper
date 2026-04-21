package orchestrator

import (
	"github.com/GnotAGnoob/kosik-scraper/internal/llm/shared"
)

type fieldResult[T any] struct {
	Value T       `json:"value" validate:"required"`
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

type pricePerUnit struct {
	Value float64 `json:"value" validate:"required"`
	Unit  string  `json:"unit" validate:"required"`
}

type productQuantity struct {
	Prefix string  `json:"prefix" validate:"required"`
	Value  float64 `json:"value" validate:"required"`
	Unit   string  `json:"unit" validate:"required"`
}

type saleLabel struct {
	Name       string `json:"name" validate:"required"`
	Background string `json:"background"`
	StyleKey   string `json:"styleKey"`
	Priority   int    `json:"priority"`
}

type cumulativePriceFields struct {
	Quantity     int          `json:"quantity" validate:"required"`
	Price        float64      `json:"price" validate:"required"`
	PricePerUnit pricePerUnit `json:"pricePerUnit"`
}

type salesFields struct {
	RecommendedPrice   *float64                `json:"recommendedPrice,omitempty"`
	PercentageDiscount *int                    `json:"percentageDiscount,omitempty"`
	ActionLabel        *string                 `json:"actionLabel,omitempty"`
	Labels             []saleLabel             `json:"labels,omitempty"`
	MembershipPrice    *float64                `json:"membershipPrice,omitempty"`
	CumulativePrices   []cumulativePriceFields `json:"cumulativePrices,omitempty"`
}

type productFields struct {
	Name            fieldResult[string]           `json:"name"`
	Price           fieldResult[float64]          `json:"price"`
	Currency        fieldResult[string]           `json:"currency"`
	Unit            fieldResult[string]           `json:"unit"`
	PricePerUnit    fieldResult[pricePerUnit]     `json:"pricePerUnit"`
	ProductQuantity fieldResult[*productQuantity] `json:"productQuantity"`
	Link            fieldResult[*string]          `json:"link"`
	ImageURL        fieldResult[*string]          `json:"imageUrl"`
	IsSoldOut       fieldResult[bool]             `json:"isSoldOut"`
	Sales           fieldResult[*salesFields]     `json:"sales"`
	Nutrition       fieldResult[*nutritionFields] `json:"nutrition"`
}

type IngredientProducts struct {
	Ingredient shared.Ingredient `json:"ingredient" validate:"required"`
	Source     string            `json:"source" validate:"required"`
	SearchURL  string            `json:"searchUrl" validate:"required"`
	Products   []*productFields  `json:"products" validate:"required"`
}

type RecipeResult struct {
	Recipe              *string              `json:"recipe"`
	IngredientsProducts []IngredientProducts `json:"ingredientsProducts"`
}
