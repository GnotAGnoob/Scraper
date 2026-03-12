package kosik

import (
	"context"
	"math"
	"net/http"
	"net/url"
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
	apiLink, err := urlParams.CreateProductUrl(product.URL)

	if err != nil {
		return
	}

	logger.Debug().Str("product", product.Name).Int("index", index).Msg("scraping nutritions")
	parsedNutritionData, err := httpUtils.SendRequestWithRetry[ProductDetailResponse](ctx, client, http.MethodGet, apiLink.String(), nil)
	if err != nil {
		productResult.Result.Value.Nutrition.ScrapeErr = err
		return
	}

	productResult.Result.Value.Nutrition.Value = transformKosikSearchProductDetailToNutrition(&parsedNutritionData.Product.Detail)
}

func GetProducts(ctx context.Context, search string, maxProducts int, infoChan chan<- shared.SearchInfo, productsChan chan<- *shared.ProductResult) error {
	defer func() {
		close(infoChan)
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

	searchURLString := searchURL.String()
	logger.Debug().Str("search_url", searchURLString).Msg("searching kosik")

	searchTerm := urlParams.ExtractSearchTerm(search)
	var userSearchURL string
	if parsedURL, err := url.Parse(searchTerm); err == nil && parsedURL.IsAbs() {
		userSearchURL = searchTerm
	} else {
		userSearchURLObj := urlParams.GetKosikUserSearchUrl(searchTerm)
		userSearchURL = userSearchURLObj.String()
	}

	parsedSearchData, err := httpUtils.SendRequestWithRetry[SearchResponse](ctx, &client, http.MethodGet, searchURLString, nil)
	if err != nil {
		infoChan <- shared.SearchInfo{
			Total:     0,
			SearchURL: userSearchURL,
		}
		return err
	}

	searchProductCount := parsedSearchData.Products.TotalCount
	total := int(math.Min(float64(searchProductCount), urlParams.KosikProductLimit))
	logger.Debug().Int("total", total).Int("max_products", maxProducts).Msg("products found")

	infoChan <- shared.SearchInfo{
		Total:     total,
		SearchURL: userSearchURL,
	}

	if total == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	scrapedCount := 0
	for index, product := range parsedSearchData.Products.Items {
		if scrapedCount >= maxProducts {
			break
		}
		wg.Add(1)
		scrapedCount++
		go scrapeProductAsync(ctx, index, &product, &client, productsChan, &wg)
	}

	if scrapedCount >= maxProducts || searchProductCount <= urlParams.KosikLimit {
		return nil
	}

	searchMoreURL := urlParams.GetKosikSearchMoreUrl()
	logger.Debug().Str("search_url", searchMoreURL.String()).Msg("fetching additional products")

	reqBody, err := urlParams.CreateSearchMoreBody(parsedSearchData.Products.Cursor)
	if err != nil {
		return err
	}

	parsedSearchMoreData, err := httpUtils.SendRequestWithRetry[SearchMoreResponse](ctx, &client, http.MethodPost, searchMoreURL.String(), reqBody.Bytes())
	if err != nil {
		return err
	}

	for index, product := range parsedSearchMoreData.Products {
		if scrapedCount >= maxProducts {
			break
		}
		wg.Add(1)
		scrapedCount++
		go scrapeProductAsync(ctx, index+len(parsedSearchData.Products.Items), &product, &client, productsChan, &wg)
	}

	return nil
}
