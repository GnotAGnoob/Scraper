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

var headers = []headerCategory{
	{Name: "Product", WidthWeight: 4},
	{Name: "Status", WidthWeight: 1},
	{Name: "Price", WidthWeight: 1},
	{Name: "Price per kg", WidthWeight: 1.75},
	{Name: "Unit", WidthWeight: 1},
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
const extraItemWidth = 2 // 2 because of a padding on each side

func setStyle(tab table.Writer) {
	tab.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	tab.Style().Color.Header = text.Colors{text.BgRed, text.FgBlack, text.Bold}
	tab.Style().Format.Header = text.FormatTitle
	tab.Style().Color.IndexColumn = text.Colors{text.BgRed, text.FgBlack}
	tab.Style().Color.Row = text.Colors{text.BgBlack, text.FgWhite}
	tab.Style().Color.RowAlternate = text.Colors{text.BgHiBlack, text.FgHiWhite}
}

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
	indexDigitsCount := int(math.Log10(float64(itemsCount)) + 1)                      // number of digits in itemsCount
	width -= float64(len(headers)*extraItemWidth + extraItemWidth + indexDigitsCount) // subtracting extra non-item width
	widthFragment := float64(width) / float64(sumWidthWeight)

	var columnConfigs []table.ColumnConfig
	var headerRow table.Row
	var overflowWidth float64
	// the width for each column is calculated based on the width of the terminal and the weight of each column
	// the width of the terminal is divided into fragments based on the sum of the weight of all columns
	// allocation of the width for each column is based on the weight defined in headers variable
	// the width of each column is then rounded down to the nearest integer
	// the remainder of the width is stored in overflowWidth and is added to the next column
	// if the overflowWidth is greater than or equal to 1, the width of the column is increased by 1 and the overflowWidth is decreased by 1
	// this is done to ensure that the width of the table is equal to the width of the terminal. The overflowing text is trimmed
	// full width might be missing one pixel due to the rounding error of the float division (dont want to deal with that)
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
	setStyle(tab)

	return tab
}
