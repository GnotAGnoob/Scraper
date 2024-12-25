package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
	"github.com/GnotAGnoob/kosik-scraper/internal/utils/urlParams"
	"github.com/rs/zerolog/log"
)

type ReturnProduct struct {
	structs.ScrapeResult[*product]
}

type ProductResult struct {
	Index  int
	Result *ReturnProduct
}

func sendRequest[T any](client *http.Client, method string, url string, body io.Reader) (t T, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return t, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return t, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(resBody, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

// todo add the setting of the address for kosik site
// todo debug logs
// todo handle timeout => send what was found and errors for the rest
// todo parse things to floats / ints
// todo sometimes the calories are not found but there are joules instead
// todo caching?
func GetKosikProducts(search string, totalChan chan<- int, productsChan chan<- *ProductResult) error {
	defer func() {
		close(totalChan)
		close(productsChan)
	}()

	client := http.Client{
		Timeout: time.Second * 10,
	}

	searchUrl, err := urlParams.CreateSearchUrl(search)
	if err != nil {
		return err
	}

	// todo handle bad response
	parsedSearchData, err := sendRequest[SearchResponse](&client, http.MethodGet, searchUrl.String(), nil)
	if err != nil {
		return err
	}

	searchProductCount := parsedSearchData.Products.TotalCount
	total := int(math.Min(float64(searchProductCount), urlParams.KosikProductLimit))
	log.Debug().Msgf("Found %d products", total)
	totalChan <- total

	if total == 0 {
		return nil
	}

	wg := sync.WaitGroup{}

	for index, product := range parsedSearchData.Products.Items {
		wg.Add(1)
		go func(index int, productId string) {
			defer wg.Done()

			fmt.Println("Product", productId)

			productUrl, err := urlParams.CreateProductUrl(productId)
			if err != nil {
				// todo
				return
			}
			fmt.Println("Productxxx", productUrl)

			_, err = sendRequest[ProductDetailResponse](&client, http.MethodGet, productUrl.String(), nil)
			if err != nil {
				// todo
				return
			}

			productsChan <- &ProductResult{}
		}(index, product.URL)
	}

	if searchProductCount <= urlParams.KosikLimit {
		return nil
	}

	searchMoreUrl := urlParams.GetKosikSearchMoreUrl()

	reqBody, err := urlParams.CreateSearchMoreBody(parsedSearchData.Products.Cursor)
	if err != nil {
		return err
	}

	parsedSearchMoreData, err := sendRequest[SearchMoreResponse](&client, http.MethodPost, searchMoreUrl.String(), reqBody)
	if err != nil {
		return err
	}

	// todo these data for loop

	for index, product := range parsedSearchMoreData.Products {
		wg.Add(1)
		go func(index int, productId string) {
			defer wg.Done()

			fmt.Println("Product", productId)

			productUrl, err := urlParams.CreateProductUrl(productId)
			if err != nil {
				// todo
				return
			}
			fmt.Println("Productxxx", productUrl)

			_, err = sendRequest[ProductDetailResponse](&client, http.MethodGet, productUrl.String(), nil)
			if err != nil {
				// todo
				return
			}

			fmt.Println("Product", productId)
			productsChan <- &ProductResult{}

		}(index, product.URL)
	}

	wg.Wait()

	return err
}

func main() {
	totalChan := make(chan int)
	productsChan := make(chan *ProductResult)

	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = GetKosikProducts("https://www.kosik.cz/c1046-uzeniny-a-lahudky?orderBy=unit-price-asc", totalChan, productsChan)
	}()

	total, ok := <-totalChan
	if !ok || total == 0 {
		fmt.Println("No products found")
		return
	}

	fmt.Println("Found", total, "products")

	wg.Wait()

	if err != nil {
		fmt.Println(err)
	}
}
