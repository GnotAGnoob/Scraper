package kosik

import (
	"math"
	"net/http"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/kosik/urlParams"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/httpUtils"
	"github.com/rs/zerolog/log"
)

func scrapeProductAsync(index int, product *Product, client *http.Client, ch chan<- *shared.ProductResult, wg *sync.WaitGroup) {
	defer wg.Done()

	productResult := transformKosikSearchProductToProduct(index, product)
	defer func() {
		ch <- productResult
	}()
	productLink := productResult.Result.Value.Link

	if productLink.ScrapeErr != nil {
		return
	}
	log.Debug().Msgf("Scraping nutritions for %d %s at %s", index, product.Name, productLink.Value.String())

	parsedNutritionData, err := httpUtils.SendRequest[ProductDetailResponse](client, http.MethodGet, productLink.Value.String(), nil)
	if err != nil {
		productResult.Result.Value.Nutrition.ScrapeErr = err
		return
	}

	productResult.Result.Value.Nutrition.Value = transformKosikSearchProductDetailToNutrition(&parsedNutritionData.Product.Detail)
}

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
	log.Debug().Msgf("Searching for %s", searchUrl.String())

	parsedSearchData, err := httpUtils.SendRequest[SearchResponse](&client, http.MethodGet, searchUrl.String(), nil)
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
	defer wg.Wait()

	for index, product := range parsedSearchData.Products.Items {
		wg.Add(1)
		go scrapeProductAsync(index, &product, &client, productsChan, &wg)
	}

	if searchProductCount <= urlParams.KosikLimit {
		return nil
	}

	searchMoreUrl := urlParams.GetKosikSearchMoreUrl()
	log.Debug().Msgf("Searching for more products at %s", searchMoreUrl.String())

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
		go scrapeProductAsync(index+len(parsedSearchData.Products.Items), &product, &client, productsChan, &wg)
	}

	wg.Wait()

	return err
}
