package kosik

import (
	"fmt"
	"math"
	"net/http"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/kosik/urlParams"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/httpUtils"
	"github.com/rs/zerolog/log"
)

func GetProducts(search string, totalChan chan<- int, productsChan chan<- *shared.ProductResult) error {
	defer func() {
		close(totalChan)
		close(productsChan)
	}()

	client := http.Client{
		Timeout: shared.RequestTimeout,
	}

	searchUrl, err := urlParams.CreateSearchUrl(search)
	if err != nil {
		return err
	}

	// todo handle bad response
	parsedSearchData, err := httpUtils.SendRequest[SearchResponse](&client, http.MethodGet, searchUrl.String(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Found", parsedSearchData.Products.TotalCount, "products")

	searchProductCount := parsedSearchData.Products.TotalCount
	total := int(math.Min(float64(searchProductCount), urlParams.KosikProductLimit))
	log.Debug().Msgf("Found %d products", total)
	totalChan <- total

	fmt.Println("total", total)

	if total == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	for index, product := range parsedSearchData.Products.Items {
		wg.Add(1)
		go func(index int, product *Product) {
			defer wg.Done()

			productResult := transformKosikSearchProductToProduct(index, product)
			productLink := productResult.Result.Value.Link

			fmt.Println("productLink", productLink, index)

			if productLink.ScrapeErr != nil {
				productsChan <- productResult // dat to do fce a vracet error
				return
			}

			parsedNutritionData, err := httpUtils.SendRequest[ProductDetailResponse](&client, http.MethodGet, productLink.Value.String(), nil)
			if err != nil {
				productsChan <- productResult // dat to do fce a vracet error
				return
			}

			productResult.Result.Value.Nutrition.Value = transformKosikSearchProductDetailToNutrition(&parsedNutritionData.Product.Detail)

			productsChan <- productResult
		}(index, &product)
	}

	if searchProductCount <= urlParams.KosikLimit {
		return nil
	}

	searchMoreUrl := urlParams.GetKosikSearchMoreUrl()

	reqBody, err := urlParams.CreateSearchMoreBody(parsedSearchData.Products.Cursor)
	if err != nil {
		return err
	}

	parsedSearchMoreData, err := httpUtils.SendRequest[SearchMoreResponse](&client, http.MethodPost, searchMoreUrl.String(), reqBody)
	if err != nil {
		return err
	}

	for index, product := range parsedSearchMoreData.Products {
		wg.Add(1)
		go func(index int, product *Product) {
			defer wg.Done()

			productResult := transformKosikSearchProductToProduct(index, product)
			productLink := productResult.Result.Value.Link

			if productLink.ScrapeErr != nil {
				productsChan <- productResult // dat to do fce a vracet error
				return
			}

			parsedNutritionData, err := httpUtils.SendRequest[ProductDetailResponse](&client, http.MethodGet, productLink.Value.String(), nil)
			if err != nil {
				productsChan <- productResult // dat to do fce a vracet error
				return
			}

			productResult.Result.Value.Nutrition.Value = transformKosikSearchProductDetailToNutrition(&parsedNutritionData.Product.Detail)

			productsChan <- productResult
		}(index, &product)
	}

	wg.Wait()

	return err
}
