package urlParams

import (
	"errors"
	"net/url"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/constants"
)

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

func GetUrl(search string) (*url.URL, error) {
	if len(search) == 0 {
		return nil, errors.New("search term is empty")
	}

	searchUrl, err := url.Parse(search)
	if err != nil {
		return nil, err
	}

	if searchUrl.IsAbs() {
		kosikUrl := constants.GetKosikSearchUrl()

		if searchUrl.Hostname() != (&kosikUrl).Hostname() {
			return nil, errors.New("invalid URL: hostname does not match")
		}

		params := searchUrl.Query()
		isCategory := strings.HasPrefix(searchUrl.Path, "/c")

		if _, ok := params[SEARCH]; !ok && !isCategory {
			return nil, errors.New("no search term in URL or category in Path")
		}
	} else {
		kosikUrl := constants.GetKosikSearchUrl()
		searchUrl = &kosikUrl

		params := url.Values{}
		params.Add(ORDER_BY, orderByDefinitions.UnitPriceAsc)
		params.Add(SEARCH, search)

		searchUrl.RawQuery = params.Encode()
	}
	return searchUrl, nil
}
