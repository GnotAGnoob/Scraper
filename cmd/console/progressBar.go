package main

import "github.com/schollz/progressbar/v3"

const termWidthTextSubtraction = 30

func getProgressBar(text string) *progressbar.ProgressBar {
	termSubstract := termWidthTextSubtraction + len(text)
	return progressbar.NewOptions(100, progressbar.OptionSetRenderBlankState(true), progressbar.OptionClearOnFinish(), progressbar.OptionSetDescription(text), progressbar.OptionShowCount(), progressbar.OptionSetElapsedTime(true), progressbar.OptionSetPredictTime(true), progressbar.OptionSetWidth(int(getTermWidth())-termSubstract))
}
