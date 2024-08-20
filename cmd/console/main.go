package main

import (
	"fmt"
	"log"

	scraperLib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
)

// todo proper user input
// todo proper info output
func main() {
	scraper, err := scraperLib.InitScraper()
	if err != nil {
		log.Fatalf("Error while initializing scraper: %v", err)
	}
	defer scraper.Cleanup()

	products, err := scraper.GetKosikProducts("majoneza")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, product := range *products {
		fmt.Printf("Product: %+v, Error:%v\n", product.Value, product.Err)
		fmt.Printf("Nutrition: %+v, Error:%v\n", product.Value.Nutrition, product.Value.Nutrition.Err)
	}
}
