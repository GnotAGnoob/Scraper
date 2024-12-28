package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	scraperLibShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"

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

// todo handle errors gracefully
func main() {
	logLevel := flag.String("log-level", "info", "sets log level")
	flag.Parse()

	logger.Init(*logLevel)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter query or full url: ")
		scanner.Scan()
		search := strings.TrimSpace(scanner.Text())

		if len(search) == 0 {
			break
		}

		bar := getProgressBar("Scraping...", *logLevel == "debug")
		totalChan := make(chan int)
		productsChan := make(chan *scraperLibShared.ProductResult)

		var err error
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = scraperLib.GetProducts(search, totalChan, productsChan)
		}()

		total, ok := <-totalChan
		if err != nil {
			bar.Finish()
			log.Error().Err(err).Msg("Error while getting products")
			continue
		}

		if !ok {
			bar.Finish()
			log.Error().Msg("Unexpected error while getting total products")
			continue
		}

		if total == 0 {
			bar.Finish()
			log.Error().Msg("No products found")
			continue
		}

		progressbar.Bprintf(bar, "Found %d products\n", total)

		barChunk := float64(100) / float64(total)

		products := make([]*scraperLibShared.ReturnProduct, total)
		isAtleastOneProduct := false
		for i := 0; i < total; i++ {
			fmt.Printf("\rScraping product %d/%d", i+1, total)
			productResult, ok := <-productsChan
			if !ok {
				if err == nil {
					log.Fatal().Msg("channel closed unexpectedly")
				}

				break
			}
			isAtleastOneProduct = true

			progress := int(math.Ceil((float64(i+1) * barChunk)))
			bar.Set(progress)

			products[productResult.Index] = productResult.Result
		}

		wg.Wait()
		if !isAtleastOneProduct {
			bar.Finish()
			log.Error().Err(err).Msg("error while getting products")
			continue
		}

		tab := newTable(total)

		for i, product := range products {
			if product == nil {
				log.Error().Msg(fmt.Sprintf("product at index %d is nil", i))
				continue
			}
			if product.ScrapeErr != nil {
				log.Error().Err(product.ScrapeErr).Msg(fmt.Sprintf("product error at index %d is nil", i))
				continue
			}
			if product.Value == nil {
				log.Error().Msg(fmt.Sprintf("product value at index %d is nil", i))
				continue
			}

			availabilityText := "available"
			if product.Value.IsSoldOut.Value {
				availabilityText = "sold out"
			}

			row := table.Row{
				getDisplayText(product.Value.Name.Value, product.Value.Name.ScrapeErr),
				availabilityText,
				getDisplayText(formatFloatUnitToString(&product.Value.Price.Value, "Kč"), product.Value.Price.ScrapeErr),
				getDisplayText(product.Value.Unit.Value, product.Value.Unit.ScrapeErr),
				getDisplayText(formatFloatUnitToString(&product.Value.PricePerUnit.Value.Value, "Kč"), product.Value.PricePerUnit.ScrapeErr),
				getDisplayText(product.Value.PricePerUnit.Value.Unit, product.Value.PricePerUnit.ScrapeErr),
			}

			nutrition := product.Value.Nutrition
			var calories, protein, fat, saturatedFat, carbs, sugar, fiber string
			if nutrition.Value != nil {
				calories = getDisplayText(formatFloatUnitToString(nutrition.Value.Calories.Value, "kcal"), nutrition.Value.Calories.ScrapeErr)
				protein = getDisplayText(formatFloatUnitToString(nutrition.Value.Protein.Value, "g"), nutrition.Value.Protein.ScrapeErr)
				fat = getDisplayText(formatFloatUnitToString(nutrition.Value.Fat.Value, "g"), nutrition.Value.Fat.ScrapeErr)
				saturatedFat = getDisplayText(formatFloatUnitToString(nutrition.Value.SaturatedFat.Value, "g"), nutrition.Value.SaturatedFat.ScrapeErr)
				carbs = getDisplayText(formatFloatUnitToString(nutrition.Value.Carbs.Value, "g"), nutrition.Value.Carbs.ScrapeErr)
				sugar = getDisplayText(formatFloatUnitToString(nutrition.Value.Sugar.Value, "g"), nutrition.Value.Sugar.ScrapeErr)
				fiber = getDisplayText(formatFloatUnitToString(nutrition.Value.Fiber.Value, "g"), nutrition.Value.Fiber.ScrapeErr)
			}
			nutritionFields := []string{calories, protein, fat, saturatedFat, carbs, sugar, fiber}

			for _, field := range nutritionFields {
				if nutrition.ScrapeErr != nil {
					row = append(row, text.FgRed.Sprint("nutrition error"))
				} else {
					row = append(row, field)
				}
			}

			tab.AppendRow(row)
		}

		bar.Clear()
		tab.Render()
		fmt.Println()
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("error while scanning input")
	}
}
