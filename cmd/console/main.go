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

	products, err := scraper.GetKosikProducts("https://www.kosik.cz/c3154-skyry?orderBy=unit-price-asc")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, product := range *products {
		if len(*product.Errors) > 0 {
			for _, err := range *product.Errors {
				fmt.Println(err)
			}
		}

		fmt.Printf("Product: %+v\n", product.Product)
		fmt.Printf("Nutrition: %+v\n", product.Product.Nutrition)
	}
}
