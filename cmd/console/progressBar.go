package main

import "github.com/schollz/progressbar/v3"

func getProgressBar(text string) *progressbar.ProgressBar {
	return progressbar.NewOptions(100, progressbar.OptionSetRenderBlankState(true), progressbar.OptionClearOnFinish(), progressbar.OptionSetDescription(text), progressbar.OptionShowCount(), progressbar.OptionSetElapsedTime(true), progressbar.OptionSetPredictTime(true))
}
