package main

import (
	"github.com/schollz/progressbar/v3"
)

func getProgressBar(text string, logLevel string) *progressbar.ProgressBar {
	if logLevel == "debug" {
		return progressbar.DefaultSilent(100) // return a silent progress bar because it is annoying when debugging
	}

	return progressbar.NewOptions(100, progressbar.OptionSetRenderBlankState(true), progressbar.OptionClearOnFinish(), progressbar.OptionSetDescription(text), progressbar.OptionShowCount(), progressbar.OptionSetElapsedTime(true), progressbar.OptionSetPredictTime(true))
}
