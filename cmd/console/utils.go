package main

import (
	"math"
	"os"

	"golang.org/x/term"
)

const maxTableWidth = 250
const minTableWidth = 80

func getTermWidth() float64 {
	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = minTableWidth
	}

	return math.Min(float64(termWidth), maxTableWidth)
}
