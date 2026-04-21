package orchestrator

import (
	"net/url"
	"strings"

	scraperLibShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
)

func transformProducts(products []*scraperLibShared.ReturnProduct) []*productFields {
	transformed := make([]*productFields, len(products))

	for index, product := range products {
		if product == nil || product.ScrapeErr != nil || product.Value == nil {
			continue
		}

		fields := mapProductFields(product.Value)
		transformed[index] = fields
	}

	return transformed
}

func mapProductFields(product *scraperLibShared.Product) *productFields {
	fields := &productFields{
		Name:            mapField(product.Name),
		Price:           mapField(product.Price),
		Currency:        mapField(product.Currency),
		Unit:            mapField(product.Unit),
		PricePerUnit:    mapPricePerUnitField(product.PricePerUnit),
		ProductQuantity: mapProductQuantityField(product.ProductQuantity),
		Link:            mapURLField(product.Link),
		ImageURL:        mapImageURLField(product.ImageUrl),
		IsSoldOut:       mapField(product.IsSoldOut),
		Sales:           mapSalesField(product.Sales),
	}

	fields.Nutrition = mapNutritionField(product.Nutrition)

	return fields
}

func mapSalesField(result scraperLibShared.ScrapeResult[*scraperLibShared.Sales]) fieldResult[*salesFields] {
	res := fieldResult[*salesFields]{Error: errorString(result.ScrapeErr)}
	if result.Value == nil {
		return res
	}

	sales := &salesFields{
		RecommendedPrice:   result.Value.RecommendedPrice,
		PercentageDiscount: result.Value.PercentageDiscount,
		ActionLabel:        result.Value.ActionLabel,
		MembershipPrice:    result.Value.MembershipPrice,
	}

	if len(result.Value.Labels) > 0 {
		labels := make([]saleLabel, len(result.Value.Labels))
		for i, l := range result.Value.Labels {
			labels[i] = saleLabel{
				Name:       l.Name,
				Background: l.Background,
				StyleKey:   l.StyleKey,
				Priority:   l.Priority,
			}
		}
		sales.Labels = labels
	}

	if len(result.Value.CumulativePrices) > 0 {
		cumulative := make([]cumulativePriceFields, len(result.Value.CumulativePrices))
		for i, cp := range result.Value.CumulativePrices {
			cumulative[i] = cumulativePriceFields{
				Quantity: cp.Quantity,
				Price:    cp.Price,
				PricePerUnit: pricePerUnit{
					Value: cp.PricePerUnit.Value,
					Unit:  cp.PricePerUnit.Unit,
				},
			}
		}
		sales.CumulativePrices = cumulative
	}

	res.Value = sales
	return res
}

func mapNutritionField(nutrition scraperLibShared.ScrapeResult[*scraperLibShared.Nutrition]) fieldResult[*nutritionFields] {
	result := fieldResult[*nutritionFields]{Error: errorString(nutrition.ScrapeErr)}

	if nutrition.Value == nil {
		return result
	}

	fields := &nutritionFields{
		Calories:     mapField(nutrition.Value.Calories),
		Protein:      mapField(nutrition.Value.Protein),
		Fat:          mapField(nutrition.Value.Fat),
		SaturatedFat: mapField(nutrition.Value.SaturatedFat),
		Carbs:        mapField(nutrition.Value.Carbs),
		Sugar:        mapField(nutrition.Value.Sugar),
		Fiber:        mapField(nutrition.Value.Fiber),
		Ingredients:  mapField(nutrition.Value.Ingredients),
	}

	result.Value = fields
	return result
}

func mapField[T any](result scraperLibShared.ScrapeResult[T]) fieldResult[T] {

	return fieldResult[T]{
		Value: result.Value,
		Error: errorString(result.ScrapeErr),
	}
}

func mapURLField(result scraperLibShared.ScrapeResult[*url.URL]) fieldResult[*string] {
	return fieldResult[*string]{
		Value: urlToStringPtr(result.Value),
		Error: errorString(result.ScrapeErr),
	}
}

func mapImageURLField(result scraperLibShared.ScrapeResult[*url.URL]) fieldResult[*string] {
	return fieldResult[*string]{
		Value: imageURLToStringPtr(result.Value),
		Error: errorString(result.ScrapeErr),
	}
}

func mapPricePerUnitField(result scraperLibShared.ScrapeResult[scraperLibShared.PricePerUnit]) fieldResult[pricePerUnit] {
	return fieldResult[pricePerUnit]{
		Value: pricePerUnit{
			Value: result.Value.Value,
			Unit:  result.Value.Unit,
		},
		Error: errorString(result.ScrapeErr),
	}
}

func mapProductQuantityField(result scraperLibShared.ScrapeResult[*scraperLibShared.ProductQuantity]) fieldResult[*productQuantity] {
	res := fieldResult[*productQuantity]{Error: errorString(result.ScrapeErr)}
	if result.Value != nil {
		res.Value = &productQuantity{
			Prefix: result.Value.Prefix,
			Value:  result.Value.Value,
			Unit:   result.Value.Unit,
		}
	}
	return res
}

func urlToStringPtr(u *url.URL) *string {
	if u == nil {
		return nil
	}
	str := u.String()
	return &str
}

func imageURLToStringPtr(u *url.URL) *string {
	if u == nil {
		return nil
	}
	str := u.String()
	// Replace WIDTH and HEIGHT placeholders with actual dimensions
	// Using 400x400 for good quality thumbnails
	str = strings.ReplaceAll(str, "WIDTH", "400")
	str = strings.ReplaceAll(str, "HEIGHT", "400")
	return &str
}

func errorString(err error) *string {
	if err == nil {
		return nil
	}
	msg := err.Error()
	return &msg
}
