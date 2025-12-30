package shared

import (
	"net/url"
)

type Nutrition struct {
	Calories     ScrapeResult[*float64]
	Protein      ScrapeResult[*float64]
	Fat          ScrapeResult[*float64]
	SaturatedFat ScrapeResult[*float64]
	Carbs        ScrapeResult[*float64]
	Sugar        ScrapeResult[*float64]
	Fiber        ScrapeResult[*float64]
	Ingredients  ScrapeResult[*string]
}

type PricePerUnit struct {
	Value float64
	Unit  string
}

type ProductQuantity struct {
	Prefix string
	Value  float64
	Unit   string
}

type Product struct {
	Name            ScrapeResult[string]
	Price           ScrapeResult[float64]
	Currency        ScrapeResult[string]
	Unit            ScrapeResult[string]
	PricePerUnit    ScrapeResult[PricePerUnit]
	ProductQuantity ScrapeResult[*ProductQuantity]
	Link            ScrapeResult[*url.URL]
	ImageUrl        ScrapeResult[*url.URL]
	IsSoldOut       ScrapeResult[bool]
	Nutrition       ScrapeResult[*Nutrition]
}

type ReturnProduct struct {
	ScrapeResult[*Product]
}

type ProductResult struct {
	Index  int
	Result *ReturnProduct
}

type ScrapeResult[T any] struct {
	Value     T
	ScrapeErr error
}

type SearchInfo struct {
	Total     int
	SearchURL string
}
