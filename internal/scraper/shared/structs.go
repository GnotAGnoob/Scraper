package shared

import (
	"net/url"

	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
)

type Nutrition struct {
	Calories     structs.ScrapeResult[*float64]
	Protein      structs.ScrapeResult[*float64]
	Fat          structs.ScrapeResult[*float64]
	SaturatedFat structs.ScrapeResult[*float64]
	Carbs        structs.ScrapeResult[*float64]
	Sugar        structs.ScrapeResult[*float64]
	Fiber        structs.ScrapeResult[*float64]
	Ingredients  structs.ScrapeResult[*string]
}

type PricePerUnit struct {
	Value float64
	Unit  string
}

type Product struct {
	Name         structs.ScrapeResult[string]
	Price        structs.ScrapeResult[float64]
	Unit         structs.ScrapeResult[string]
	PricePerUnit structs.ScrapeResult[PricePerUnit]
	Link         structs.ScrapeResult[*url.URL]
	ImageUrl     structs.ScrapeResult[*url.URL]
	IsSoldOut    structs.ScrapeResult[bool]
	Nutrition    structs.ScrapeResult[*Nutrition]
}

type ReturnProduct struct {
	structs.ScrapeResult[*Product]
}

type ProductResult struct {
	Index  int
	Result *ReturnProduct
}
