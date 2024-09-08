package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
	"github.com/rs/zerolog/log"

	scraperLib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func getDisplayText(value string, err error) string {
	if err != nil {
		return text.FgRed.Sprint("error")
	}
	return value
}

// todo progress bar -> need channel
func main() {
	isDebug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	logger.Init(*isDebug)

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
		search := strings.TrimSpace(scanner.Text())

		if len(search) == 0 {
			break
		}

		products, err := scraper.GetKosikProducts(search)
		if err != nil {
			log.Fatal().Err(err).Msg("error while getting products")
		}

		tab := NewTable(len(products))

		for _, product := range products {
			if product == nil || product.Value == nil || product.ScrapeErr != nil {
				log.Error().Err(product.ScrapeErr).Msg("error while getting product")
				continue
			}

			log.Debug().Msgf("Product: %+v", product.Value)
			log.Debug().Err(product.ScrapeErr)
			log.Debug().Msgf("Nutritions: %+v", product.Value.Nutrition.Value)
			log.Debug().Err(product.Value.Nutrition.ScrapeErr)
			log.Debug().Msg("\n")

			availabilityText := "available"
			if product.Value.IsSoldOut {
				availabilityText = "sold out"
			}

			row := table.Row{
				getDisplayText(product.Value.Name.Value, product.Value.Name.ScrapeErr),
				availabilityText,
				getDisplayText(product.Value.Price.Value, product.Value.Price.ScrapeErr),
				getDisplayText(product.Value.PricePerKg.Value, product.Value.PricePerKg.ScrapeErr),
				getDisplayText(product.Value.Unit.Value, product.Value.Unit.ScrapeErr),
			}

			nutrition := product.Value.Nutrition
			var calories, protein, fat, saturatedFat, carbs, sugar, fiber structs.ScrapeResult[string]
			if nutrition.Value != nil {
				calories = nutrition.Value.Calories
				protein = nutrition.Value.Protein
				fat = nutrition.Value.Fat
				saturatedFat = nutrition.Value.SaturatedFat
				carbs = nutrition.Value.Carbs
				sugar = nutrition.Value.Sugar
				fiber = nutrition.Value.Fiber
			}
			nutritionFields := []structs.ScrapeResult[string]{calories, protein, fat, saturatedFat, carbs, sugar, fiber}

			nutritionErr := ""
			if nutrition.ScrapeErr != nil {
				nutritionErr = text.FgRed.Sprint("nutrition error")
			}

			for _, field := range nutritionFields {
				if len(nutritionErr) > 0 {
					row = append(row, nutritionErr)
				} else {
					row = append(row, getDisplayText(field.Value, field.ScrapeErr))
				}
			}

			tab.AppendRow(row)
		}

		tab.Render()
		fmt.Println()
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("error while scanning input")
	}
}
