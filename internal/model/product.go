package model

type PricePerUnit struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type ProductSummary struct {
	Name         string        `json:"name"`
	Price        *float64      `json:"price,omitempty"`
	Unit         string        `json:"unit,omitempty"`
	PricePerUnit *PricePerUnit `json:"price_per_unit,omitempty"`
	Link         string        `json:"link,omitempty"`
	ImageURL     string        `json:"image_url,omitempty"`
	IsSoldOut    *bool         `json:"is_sold_out,omitempty"`
}
