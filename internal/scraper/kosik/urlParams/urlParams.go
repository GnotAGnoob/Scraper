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

	// if the search term was a full URL
	if searchUrl.IsAbs() {
		if searchUrl.Hostname() != finalUrl.Hostname() {
			return nil, errors.New("invalid URL: hostname does not match")
		}

		searchParams := searchUrl.Query()
		isCategory := strings.HasPrefix(searchUrl.Path, "/c")

		if isCategory {
			slug := strings.Split(strings.Trim(searchUrl.Path, "/"), "/")
			if len(slug) == 0 {
				return nil, errors.New("no category in Path")
			} else if len(slug) > 1 {
				return nil, errors.New("url's path does not match category")
			}
			params.Set(slugParam, slug[0])
		} else if param, ok := searchParams[searchParam]; !ok {
			return nil, errors.New("no search term in URL or category in Path")
		} else {
			params.Set(searchTermParam, param[0])
			params.Set(slugParam, searchSlug)
		}
	} else { // if the search term was just a string
		params.Set(searchTermParam, search)
		params.Set(slugParam, searchSlug)
	}

	params.Set(orderByParam, orderByDefinitions.UnitPriceAsc)
	fmt.Println("PARAMS", params)
	finalUrl.RawQuery = params.Encode()

	return &finalUrl, nil
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

func CreateProductUrl(productPath string) (*url.URL, error) {
	if len(productPath) == 0 {
		return nil, errors.New("product path is empty")
	}

	productUrl := GetKosikProductDetailUrl()
	newPath, err := url.JoinPath(productUrl.Path, productPath)
	if err != nil {
		return nil, err
	}
	productUrl.Path = newPath

	return &productUrl, nil
}
