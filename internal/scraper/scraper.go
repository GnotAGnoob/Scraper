package scraper

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/GnotAGnoob/kosik-scraper/pkg/constants"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
)

func getUrl(search string) (*url.URL, error) {
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

		if _, ok := params[urlParams.SEARCH]; !ok && !isCategory {
			return nil, errors.New("no search term in URL or category in Path")
		}
	} else {
		kosikUrl := constants.GetKosikSearchUrl()
		searchUrl = &kosikUrl

		params := url.Values{}
		params.Add(urlParams.ORDER_BY, urlParams.GetOrderBy().UnitPriceAsc)
		params.Add(urlParams.SEARCH, search)

		searchUrl.RawQuery = params.Encode()
	}
	return searchUrl, nil
}

func GetKosikItems(search string) ([]*string, error) {
	searchUrl, err := getUrl(search)
	if err != nil {
		return nil, err
	}

	page := rod.New().NoDefaultDevice().MustConnect().MustPage(searchUrl.String())
	defer page.MustClose()

	page.MustWindowMaximize()
	page.MustWaitStable()
	time.Sleep(time.Hour)

	// items := page.MustElements(".product-list .product-item")

	// var names []string
	// for _, item := range items {
	// 	names = append(names, item.MustElement(".product-name").MustText())
	// }

	// return names
	return nil, nil
}
