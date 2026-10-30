package kosik

import "testing"

func TestSearchProductKeepsStorefrontAndDetailURLsSeparate(t *testing.T) {
	result := transformKosikSearchProductToProduct(0, &Product{URL: "/p12877-butter"})
	product := result.Result.Value
	if product.StorefrontURL.ScrapeErr != nil || product.StorefrontURL.Value.String() != "https://www.kosik.cz/p12877-butter" {
		t.Fatalf("unexpected storefront: %+v", product.StorefrontURL)
	}
	if product.Link.Value.String() != "https://www.kosik.cz/api/front/product/slug/p12877-butter" {
		t.Fatalf("detail link changed: %v", product.Link.Value)
	}
}
