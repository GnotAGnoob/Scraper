package main

import (
	"math"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/term"
)

type headerCategory struct {
	Name        string
	WidthWeight float64
}

// Define a slice of headerCategory directly
var headers = []headerCategory{
	{Name: "Product", WidthWeight: 4},
	{Name: "Status", WidthWeight: 1},
	{Name: "Price", WidthWeight: 1},
	{Name: "Price per kg", WidthWeight: 1.75},
	{Name: "Unit", WidthWeight: 0.75},
	{Name: "Calories", WidthWeight: 1},
	{Name: "Protein", WidthWeight: 0.8},
	{Name: "Fat", WidthWeight: 0.75},
	{Name: "Saturated Fat", WidthWeight: 0.8},
	{Name: "Carbs", WidthWeight: 0.75},
	{Name: "Sugar", WidthWeight: 0.75},
	{Name: "Fiber", WidthWeight: 0.75},
}

const maxTableWidth = 250
const minTableWidth = 80
const indexWidth = 3
const extraItemWidth = 2 // 2 because padding on each side

func NewTable(itemsCount int) table.Writer {
	tab := table.NewWriter()
	tab.SetAutoIndex(true)
	tab.SetOutputMirror(os.Stdout)

	var sumWidthWeight float64
	for _, header := range headers {
		sumWidthWeight += header.WidthWeight
	}

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = minTableWidth
	}
	width := math.Min(float64(termWidth), maxTableWidth)
	digitsCount := int(math.Log10(float64(itemsCount))) // number of digits in itemsCount - 1
	width -= float64(len(headers)*extraItemWidth + indexWidth + digitsCount)
	widthFragment := float64(width) / float64(sumWidthWeight)

	var columnConfigs []table.ColumnConfig
	var headerRow table.Row
	var overflowWidth float64
	for _, header := range headers {
		widthMax := widthFragment * header.WidthWeight
		widthRemainder := math.Mod(widthMax, 1)
		widthMax -= widthRemainder
		overflowWidth += widthRemainder
		if overflowWidth >= 1 {
			add := math.Floor(overflowWidth)
			overflowWidth -= add
			widthMax += add
		}

		config := table.ColumnConfig{
			Name:             header.Name,
			WidthMax:         int(widthMax),
			WidthMin:         int(widthMax),
			WidthMaxEnforcer: text.Trim,
		}

		columnConfigs = append(columnConfigs, config)
		headerRow = append(headerRow, header.Name)
	}
	tab.AppendHeader(headerRow)
	tab.SetColumnConfigs(columnConfigs)
	tab.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	tab.Style().Color.Header = text.Colors{text.BgRed, text.FgBlack, text.Bold}
	tab.Style().Format.Header = text.FormatTitle
	tab.Style().Color.IndexColumn = text.Colors{text.BgRed, text.FgBlack}
	tab.Style().Color.Row = text.Colors{text.BgBlack, text.FgWhite}
	tab.Style().Color.RowAlternate = text.Colors{text.BgHiBlack, text.FgHiWhite}

	return tab
}
