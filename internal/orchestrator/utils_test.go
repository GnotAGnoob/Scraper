package orchestrator

import (
	scraperShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"net/url"
	"testing"
)

func TestMapProductStorefrontURL(t *testing.T) {
	storefront, _ := url.Parse("https://www.kosik.cz/p12877-butter")
	fields := mapProductFields(&scraperShared.Product{
		StorefrontURL: scraperShared.ScrapeResult[*url.URL]{Value: storefront},
	})
	if fields.StorefrontURL.Value == nil || *fields.StorefrontURL.Value != storefront.String() {
		t.Fatalf("storefront missing from API response: %+v", fields.StorefrontURL)
	}
}
