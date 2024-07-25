package main

import (
	"log"

	scraperLib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
)

func main() {
	scraper := scraperLib.NewScraper()
	_, err := scraper.GetKosikItems("banán")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
