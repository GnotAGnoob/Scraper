package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	scraperLib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	"github.com/jedib0t/go-pretty/v6/table"
)

// todo proper user input
// todo proper info output
// todo think about error handling
// retry on error input
// input validation
// todo progress bar -> need channel
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

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter query or full url: ")
		scanner.Scan()
		search := scanner.Text()

		if len(search) == 0 {
			break
		}

		products, err := scraper.GetKosikProducts("majoneza")
		if err != nil {
			log.Fatal().Err(err).Msg("error while getting products")
		}

		tab := NewTable()

		for _, product := range *products {
			log.Debug().Msgf("Product: %+v", product.Value)
			log.Debug().Err(product.ScrapeErr)
			log.Debug().Msgf("Nutritions: %+v", product.Value.Nutrition.Value)
			log.Debug().Err(product.Value.Nutrition.ScrapeErr)
			log.Debug().
				tab.AppendRow(table.Row{
				product.Value.Name.Value,
				product.Value.Price.Value,
				product.Value.PricePerKg.Value,
				product.Value.Unit.Value,
				product.Value.Nutrition.Value.Calories.Value,
				product.Value.Nutrition.Value.Protein.Value,
				product.Value.Nutrition.Value.Fat.Value,
				product.Value.Nutrition.Value.SaturatedFat.Value,
				product.Value.Nutrition.Value.Carbs.Value,
				product.Value.Nutrition.Value.Sugar.Value,
				product.Value.Nutrition.Value.Fiber.Value,
				product.Value.Nutrition.Value.Ingredients.Value,
			})
		}

		tab.Render()
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalln(err)
	}
}
