package urlParams

const SEARCH = "search"
const ORDER_BY = "orderBy"

type orderBy struct {
	PriceAsc      string
	PriceDesc     string
	UnitPriceAsc  string
	UnitPriceDesc string
}

var orderByDefinitions = orderBy{
	PriceAsc:      "price-asc",
	PriceDesc:     "price-desc",
	UnitPriceAsc:  "unit-price-asc",
	UnitPriceDesc: "unit-price-desc",
}

func GetOrderBy() orderBy {
	return orderByDefinitions
}
