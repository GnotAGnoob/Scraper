package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/GnotAGnoob/kosik-scraper/internal/logger"
	scraperLibShared "github.com/GnotAGnoob/kosik-scraper/internal/scraper/shared"
	"github.com/GnotAGnoob/kosik-scraper/internal/utils/structs"
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

		bar := getProgressBar("Scraping...")
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
		if !ok || total == 0 {
			bar.Finish()
			fmt.Println("No products found")
			continue
		}

		progressbar.Bprintf(bar, "Found %d products\n", total)

		barChunk := float64(100) / float64(total)

		products := make([]*scraperLibShared.ReturnProduct, total)
		for i := 0; i < total; i++ {
			fmt.Printf("\rScraping product %d/%d", i+1, total)
			productResult, ok := <-productsChan
			if !ok {
				if err != nil {
					log.Fatal().Err(err).Msg("error while getting products")
				}

				log.Fatal().Msg("channel closed unexpectedly")
			}

			progress := int(math.Ceil((float64(i+1) * barChunk)))
			bar.Set(progress)

			products[productResult.Index] = productResult.Result
		}

		wg.Wait()
		if err != nil {
			log.Fatal().Err(err).Msg("error while getting products")
		}

		tab := NewTable(total)

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
			if product.Value.IsSoldOut.Value {
				availabilityText = "sold out"
			}

			row := table.Row{
				getDisplayText(product.Value.Name.Value, product.Value.Name.ScrapeErr),
				availabilityText,
				getDisplayText(strconv.FormatFloat(product.Value.Price.Value, 'f', 2, 64), product.Value.Price.ScrapeErr),
				getDisplayText(product.Value.Unit.Value, product.Value.Unit.ScrapeErr),
				getDisplayText(strconv.FormatFloat(product.Value.PricePerUnit.Value.Value, 'f', 2, 64), product.Value.PricePerUnit.ScrapeErr),
				getDisplayText(product.Value.PricePerUnit.Value.Unit, product.Value.PricePerUnit.ScrapeErr),
			}

			nutrition := product.Value.Nutrition
			var calories, protein, fat, saturatedFat, carbs, sugar, fiber structs.ScrapeResult[float64]
			if nutrition.Value != nil {
				calories = nutrition.Value.Calories
				protein = nutrition.Value.Protein
				fat = nutrition.Value.Fat
				saturatedFat = nutrition.Value.SaturatedFat
				carbs = nutrition.Value.Carbs
				sugar = nutrition.Value.Sugar
				fiber = nutrition.Value.Fiber
			}
			nutritionFields := []structs.ScrapeResult[float64]{calories, protein, fat, saturatedFat, carbs, sugar, fiber}

			for _, field := range nutritionFields {
				if nutrition.ScrapeErr != nil {
					row = append(row, text.FgRed.Sprint("nutrition error"))
				} else {
					row = append(row, getDisplayText(strconv.FormatFloat(field.Value, 'f', -1, 64), field.ScrapeErr))
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
