package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/logger"
	"github.com/rs/zerolog/log"

	scraperLib "github.com/GnotAGnoob/kosik-scraper/internal/scraper"
	"github.com/jedib0t/go-pretty/v6/table"
)

func getDisplayText(text string, err error) string {
	if err != nil {
		return "error"
	}
	return text
}

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

		tab := NewTable(len(products))

		fmt.Println("Found products:")
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

			name := getDisplayText(product.Value.Name.Value, product.Value.Name.ScrapeErr)
			price := getDisplayText(product.Value.Price.Value, product.Value.Price.ScrapeErr)
			pricePerKg := getDisplayText(product.Value.PricePerKg.Value, product.Value.PricePerKg.ScrapeErr)
			unit := getDisplayText(product.Value.Unit.Value, product.Value.Unit.ScrapeErr)

			soldOutText := "available"
			if product.Value.IsSoldOut {
				soldOutText = "sold out"
			}

			nutrition := product.Value.Nutrition
			nutritionErr := "none" // i would like to use longer text but the fucking table uses width of the first merge column instead of the whole space
			if nutrition.ScrapeErr != nil {
				nutritionErr = nutrition.ScrapeErr.Error()
			}

			rowConfig := table.RowConfig{AutoMerge: true}
			calories := nutritionErr
			protein := nutritionErr
			fat := nutritionErr
			saturatedFat := nutritionErr
			carbs := nutritionErr
			sugar := nutritionErr
			fiber := nutritionErr
			ingredients := nutritionErr

			if nutrition.ScrapeErr == nil && nutrition.Value != nil {
				rowConfig = table.RowConfig{AutoMerge: false}
				calories = getDisplayText(nutrition.Value.Calories.Value, nutrition.Value.Calories.ScrapeErr)
				protein = getDisplayText(nutrition.Value.Protein.Value, nutrition.Value.Protein.ScrapeErr)
				fat = getDisplayText(nutrition.Value.Fat.Value, nutrition.Value.Fat.ScrapeErr)
				saturatedFat = getDisplayText(nutrition.Value.SaturatedFat.Value, nutrition.Value.SaturatedFat.ScrapeErr)
				carbs = getDisplayText(nutrition.Value.Carbs.Value, nutrition.Value.Carbs.ScrapeErr)
				sugar = getDisplayText(nutrition.Value.Sugar.Value, nutrition.Value.Sugar.ScrapeErr)
				fiber = getDisplayText(nutrition.Value.Fiber.Value, nutrition.Value.Fiber.ScrapeErr)
				ingredients = getDisplayText(nutrition.Value.Ingredients.Value, nutrition.Value.Ingredients.ScrapeErr)
			}

			tab.AppendRow(table.Row{
				name,
				soldOutText,
				price,
				pricePerKg,
				unit,
				calories,
				protein,
				fat,
				saturatedFat,
				carbs,
				sugar,
				fiber,
				ingredients,
			}, rowConfig)
		}

		tab.Render()
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("error while scanning input")
	}
}
