package main

import (
	"fmt"
	"log"

	scraperLib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
)

func main() {
	scraper := scraperLib.NewScraper()
	products, err := scraper.GetKosikProducts("omacka k masu")
	fmt.Println("end: ", products, err)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, product := range *products {
		fmt.Println("kek: ", product)
		if len(product.Errors) > 0 {
			for _, err := range product.Errors {
				fmt.Println(err)
			}
		}

		fmt.Printf("Product: %+v\n", product.Product)
		fmt.Printf("Nutrition: %+v\n", product.Product.Nutrition)
	}
}
