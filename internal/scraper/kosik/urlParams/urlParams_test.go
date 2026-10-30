package urlParams

import "testing"

func TestProductURLs(t *testing.T) {
	// Product.URL returned by the Kosik search API.
	const productPath = "/p12877-madeta-jihoceske-maslo-82-100x10g"
	for _, path := range []string{productPath, productPath[1:]} {
		storefront, err := CreateStorefrontURL(path)
		if err != nil || storefront.String() != "https://www.kosik.cz"+productPath {
			t.Fatalf("storefront URL: %v, %v", storefront, err)
		}
		detail, err := CreateProductUrl(path)
		if err != nil || detail.String() != "https://www.kosik.cz/api/front/product/slug"+productPath {
			t.Fatalf("detail API URL changed: %v, %v", detail, err)
		}
	}
	for _, path := range []string{"", "https://example.com/p123", "//example.com/p123", "/api/front/product/slug/p123", "%"} {
		if _, err := CreateStorefrontURL(path); err == nil {
			t.Errorf("expected invalid storefront path %q to fail", path)
		}
	}
}
