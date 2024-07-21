package main

import (
	"log"

	"github.com/GnotAGnoob/kosik-scraper/internal/scraper"
)

func main() {
	_, err := scraper.GetKosikItems("https://www.kosik.cz/c9600-salatova-rajcata")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
