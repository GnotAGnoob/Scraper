package urlParams

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// todo zhezcit
// todo vzit orderby z url / pridat
// todo test prevzeti url parametru
// test categories
// test jen textu
// test vyhledavani url
// test nesmyslne url -> neexistujici produkt | kategorie | parametr
func CreateSearchUrl(search string) (*url.URL, error) {
	if len(search) == 0 {
		return nil, errors.New("search term is empty")
	}

	searchUrl, err := url.Parse(search)
	if err != nil {
		return nil, err
	}

	params := GetDefaultKosikSearchParams()
	finalUrl := GetKosikSearchUrl()

	// if the search term is a URL
	if searchUrl.IsAbs() {
		if searchUrl.Hostname() != finalUrl.Hostname() {
			return nil, errors.New("invalid URL: hostname does not match")
		}

		searchParams := searchUrl.Query()
		isCategory := strings.HasPrefix(searchUrl.Path, "/c")

		fmt.Println("params", params, isCategory, searchUrl.Path, strings.Trim(searchUrl.Path, "/"))

		if isCategory {
			slug := strings.Split(strings.Trim(searchUrl.Path, "/"), "/")
			if len(slug) == 0 {
				return nil, errors.New("no category in Path")
			} else if len(slug) > 1 {
				return nil, errors.New("url's path does not match category")
			}
			params.Add(slugParam, slug[0])
		} else if param, ok := searchParams[searchParam]; !ok {
			return nil, errors.New("no search term in URL or category in Path")
		} else {
			params.Add(searchTermParam, param[0])
			params.Add(slugParam, "vyhledavani")
		}
	} else { // if the search term is just a string
		params.Add(searchTermParam, search)
		params.Add(slugParam, "vyhledavani")
	}

	params.Add(orderByParam, orderByDefinitions.UnitPriceAsc)
	finalUrl.RawQuery = params.Encode()

	return &finalUrl, nil
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
		kosikUrl := GetKosikUrl()

		if pathUrl.Hostname() != kosikUrl.Hostname() {
			return nil, errors.New("invalid URL: hostname does not match")
		}

		params := pathUrl.Query()
		isProduct := strings.HasPrefix(pathUrl.Path, "/p")

		if _, ok := params[searchParam]; !ok && !isProduct {
			return nil, errors.New("no search term in URL or category in Path")
		}
	} else {
		kosikUrl := GetKosikUrl()
		pathUrl = &kosikUrl

		pathUrl.Path = path
	}

	return pathUrl, nil
}

func CreateSearchMoreBody(cursor string) (*bytes.Buffer, error) {
	data := map[string]string{
		"limit":  strconv.Itoa(KosikSearchMoreLimit),
		"cursor": cursor,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonData), nil
}

func CreateProductUrl(productPathId string) (*url.URL, error) {
	if len(productPathId) == 0 {
		return nil, errors.New("product id is empty")
	}

	productUrl := GetKosikProductDetailUrl()
	newPath, err := url.JoinPath(productUrl.Path, productPathId)
	if err != nil {
		return nil, err
	}
	productUrl.Path = newPath

	return &productUrl, nil
}
