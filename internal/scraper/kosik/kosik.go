package kosik

import (
	"context"
	"math"
	"net/http"
	"sync"

	applogger "github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/kosik/urlParams"
	"github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/httpUtils"
)

func scrapeProductAsync(ctx context.Context, index int, product *Product, client *http.Client, ch chan<- *shared.ProductResult, wg *sync.WaitGroup) {
	defer wg.Done()

	logger := applogger.LoggerFromContext(ctx)

	productResult := transformKosikSearchProductToProduct(index, product)
	defer func() {
		ch <- productResult
	}()
	productLink := productResult.Result.Value.Link

	if productLink.ScrapeErr != nil {
		return
	}

	logger.Debug().Str("product", product.Name).Int("index", index).Msg("scraping nutritions")
	parsedNutritionData, err := httpUtils.SendRequest[ProductDetailResponse](ctx, client, http.MethodGet, productLink.Value.String(), nil)
	if err != nil {
		productResult.Result.Value.Nutrition.ScrapeErr = err
		return
	}

	productResult.Result.Value.Nutrition.Value = transformKosikSearchProductDetailToNutrition(&parsedNutritionData.Product.Detail)
}

func GetProducts(ctx context.Context, search string, totalChan chan<- int, productsChan chan<- *shared.ProductResult) error {
	defer func() {
		close(totalChan)
		close(productsChan)
	}()

	logger := applogger.LoggerFromContext(ctx)
	client := http.Client{
		Timeout: shared.RequestTimeout,
	}

	searchURL, err := urlParams.CreateSearchUrl(search)
	if err != nil {
		return err
	}

	logger.Debug().Str("search_url", searchURL.String()).Msg("searching kosik")
	parsedSearchData, err := httpUtils.SendRequest[SearchResponse](ctx, &client, http.MethodGet, searchURL.String(), nil)
	if err != nil {
		return err
	}

	searchProductCount := parsedSearchData.Products.TotalCount
	total := int(math.Min(float64(searchProductCount), urlParams.KosikProductLimit))
	logger.Debug().Int("total", total).Msg("products found")
	totalChan <- total

	if total == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	for index, product := range parsedSearchData.Products.Items {
		wg.Add(1)
		go scrapeProductAsync(ctx, index, &product, &client, productsChan, &wg)
	}

	if searchProductCount <= urlParams.KosikLimit {
		return nil
	}

	searchMoreURL := urlParams.GetKosikSearchMoreUrl()
	logger.Debug().Str("search_url", searchMoreURL.String()).Msg("fetching additional products")

	reqBody, err := urlParams.CreateSearchMoreBody(parsedSearchData.Products.Cursor)
	if err != nil {
		return err
	}

	parsedSearchMoreData, err := httpUtils.SendRequest[SearchMoreResponse](ctx, &client, http.MethodPost, searchMoreURL.String(), reqBody)
	if err != nil {
		return err
	}

	for index, product := range parsedSearchMoreData.Products {
		wg.Add(1)
		go scrapeProductAsync(ctx, index+len(parsedSearchData.Products.Items), &product, &client, productsChan, &wg)
	}

	return nil
}
