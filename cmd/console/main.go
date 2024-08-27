package main

import (
	scraperLib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/logger"
	"github.com/rs/zerolog/log"
)

// todo proper user input
// todo proper info output
func main() {
	logger.Init()

	scraper, err := scraperLib.InitScraper()
	if err != nil {
		log.Fatal().Err(err).Msg("error while initializing scraper")
	}
	defer func() {
		err = scraper.Cleanup()
		if err != nil {
			log.Fatal().Err(err).Msg("error while cleaning up scraper")
		}
	}()

	products, err := scraper.GetKosikProducts("majoneza")
	if err != nil {
		log.Fatal().Err(err).Msg("error while getting products")
	}

	for _, product := range products {
		log.Info().Msgf("Product: %+v", product.Value)
		log.Error().Err(product.ScrapeErr)
		log.Info().Msgf("Nutritions: %+v", product.Value.Nutrition.Value)
		log.Error().Err(product.Value.Nutrition.ScrapeErr)
	}
}
