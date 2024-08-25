package main

import (
	"math"
	"os"
	"reflect"

	"github.com/jedib0t/go-pretty/text"
	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/term"
)

type headerCategory struct {
	Name        string
	WidthWeight int
}

type header struct {
	Product      headerCategory
	Price        headerCategory
	PricePerKg   headerCategory
	Unit         headerCategory
	Calories     headerCategory
	Protein      headerCategory
	Fat          headerCategory
	SaturatedFat headerCategory
	Carbs        headerCategory
	Sugar        headerCategory
	Fiber        headerCategory
	Ingredients  headerCategory
}

var headerDefinition = header{
	Product:      headerCategory{Name: "Product", WidthWeight: 3},
	Price:        headerCategory{Name: "Price", WidthWeight: 1},
	PricePerKg:   headerCategory{Name: "Price per kg", WidthWeight: 1},
	Unit:         headerCategory{Name: "Unit", WidthWeight: 1},
	Calories:     headerCategory{Name: "Calories", WidthWeight: 1},
	Protein:      headerCategory{Name: "Protein", WidthWeight: 1},
	Fat:          headerCategory{Name: "Fat", WidthWeight: 1},
	SaturatedFat: headerCategory{Name: "Saturated Fat", WidthWeight: 1},
	Carbs:        headerCategory{Name: "Carbs", WidthWeight: 1},
	Sugar:        headerCategory{Name: "Sugar", WidthWeight: 1},
	Fiber:        headerCategory{Name: "Fiber", WidthWeight: 1},
	Ingredients:  headerCategory{Name: "Ingredients", WidthWeight: 3},
}

const MAX_TABLE_WIDTH = 250
const MIN_TABLE_WIDTH = 80
const INDEX_WIDTH = 5 // 3 for the index + 2 for the borders of index
const ITEM_WIDTH = 3  // 3 because padding on each side (2) + 1 border

func NewTable() table.Writer {
	tab := table.NewWriter()
	tab.SetAutoIndex(true)
	tab.SetOutputMirror(os.Stdout)

	sumWidthWeight := 0
	v := reflect.ValueOf(headerDefinition)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		sumWidthWeight += int(field.FieldByName("WidthWeight").Int())
	}

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = MIN_TABLE_WIDTH
	}
	width := math.Min(float64(termWidth), MAX_TABLE_WIDTH)
	width -= float64(v.NumField()*ITEM_WIDTH + INDEX_WIDTH)
	widthFragment := float64(width) / float64(sumWidthWeight)

	var columnConfigs []table.ColumnConfig
	var headerRow table.Row
	var overflowWidth float64
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := field.FieldByName("Name").String()
		fieldWeight := int(field.FieldByName("WidthWeight").Int())

		widthMax := widthFragment * float64(fieldWeight)
		widthRemainder := math.Mod(widthMax, 1)
		widthMax -= widthRemainder
		overflowWidth += widthRemainder
		if overflowWidth >= 1 {
			add := math.Floor(overflowWidth)
			overflowWidth -= add
			widthMax += add
		}

		config := table.ColumnConfig{
			Name:     fieldName,
			WidthMax: int(widthMax),
			WidthMin: int(widthMax),
		}

		if fieldName == "Product" || fieldName == "Ingredients" {
			config.WidthMaxEnforcer = text.Trim
		}

		columnConfigs = append(columnConfigs, config)
		headerRow = append(headerRow, fieldName)
	}
	tab.AppendHeader(headerRow)
	tab.SetColumnConfigs(columnConfigs)

	// tab.SetStyle(table.Style{

	// })

	return tab
}
