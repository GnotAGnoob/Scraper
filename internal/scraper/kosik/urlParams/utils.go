package urlParams

import (
	"net/url"
	"strconv"
)

const kosikSearchEndpoint = "/api/front/page/products/flexible"
const kosikHost = "www.kosik.cz"
const kosikVendor = 1 // idk what it does
const KosikLimit = 30 // the apis max limit. otherwise it returns error.
const kosikSearchMoreEndpoint = "/api/front/products/more"
const kosikProductDetailEndpoint = "/api/front/product/slug"
const KosikSearchMoreLimit = 30 // our own limit. I dont want to get too many products.
const KosikProductLimit = KosikLimit + KosikSearchMoreLimit

const searchParam = "search"
const searchTermParam = "search_term"
const slugParam = "slug"
const orderByParam = "order_by"

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

func GetKosikUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   kosikHost,
	}
}

func GetKosikSearchUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   kosikHost,
		Path:   kosikSearchEndpoint,
	}
}

func GetKosikSearchMoreUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   kosikHost,
		Path:   kosikSearchMoreEndpoint,
	}
}

func GetKosikProductDetailUrl() url.URL {
	return url.URL{
		Scheme: "https",
		Host:   kosikHost,
		Path:   kosikProductDetailEndpoint,
	}
}

func GetDefaultKosikSearchParams() url.Values {
	params := url.Values{}
	params.Add("vendor", strconv.Itoa(kosikVendor))
	params.Add("limit", strconv.Itoa(KosikLimit))

	return params
}
