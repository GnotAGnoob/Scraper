package urlParams

import (
	"errors"
	"net/url"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/constants"
)

const searchParam = "search"
const orderByParam = "orderBy"

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

type SearchOptions struct {
	Search  string
	OrderBy string
	Path    string
}

func CreateSearchUrl(search string) (*url.URL, error) {
	if len(search) == 0 {
		return nil, errors.New("search term is empty")
	}

	searchUrl, err := url.Parse(search)
	if err != nil {
		return nil, err
	}

	if searchUrl.IsAbs() {
		kosikUrl := constants.GetKosikSearchUrl()

		if searchUrl.Hostname() != kosikUrl.Hostname() {
			return nil, errors.New("invalid URL: hostname does not match")
		}

		params := searchUrl.Query()
		isCategory := strings.HasPrefix(searchUrl.Path, "/c")

		if _, ok := params[searchParam]; !ok && !isCategory {
			return nil, errors.New("no search term in URL or category in Path")
		}
	} else {
		kosikUrl := constants.GetKosikSearchUrl()
		searchUrl = &kosikUrl

		params := url.Values{}
		params.Add(orderByParam, orderByDefinitions.UnitPriceAsc)
		params.Add(searchParam, search)

		searchUrl.RawQuery = params.Encode()
	}
	return searchUrl, nil
}

func CreateUrlFromPath(path string) (*url.URL, error) {
	if len(path) == 0 {
		return nil, errors.New("path term is empty")
	}

	pathUrl, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	if pathUrl.IsAbs() {
		kosikUrl := constants.GetKosikUrl()

		if pathUrl.Hostname() != kosikUrl.Hostname() {
			return nil, errors.New("invalid URL: hostname does not match")
		}

		params := pathUrl.Query()
		isProduct := strings.HasPrefix(pathUrl.Path, "/p")

		if _, ok := params[searchParam]; !ok && !isProduct {
			return nil, errors.New("no search term in URL or category in Path")
		}
	} else {
		kosikUrl := constants.GetKosikUrl()
		pathUrl = &kosikUrl

		pathUrl.Path = path
	}

	return pathUrl, nil
}
