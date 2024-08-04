package scraping

import (
	"github.com/go-rod/rod"
)

type scrapeElement interface {
	ElementR(string, string) (*rod.Element, error)
	Element(string) (*rod.Element, error)
}

func GetTextFromTable(element scrapeElement, searchText string, isDirect bool) (string, error) {
	searchedElement, err := element.ElementR("td", searchText)
	if err != nil {
		return "", nil
	}

	if !isDirect {
		searchedElement, err = searchedElement.Next()
		if err != nil {
			return "", err
		}
	}

	text, err := searchedElement.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}

func getImageResource(element scrapeElement, selector string) (*[]byte, error) {
	imageElement, err := element.Element(selector)
	if err != nil {
		return nil, err
	}

	image, err := imageElement.Resource()
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func GetPageImageResource(element *rod.Page, selector string) (*[]byte, error) {
	return getImageResource(element.Sleeper(rod.NotFoundSleeper), selector)
}

func GetElementImageResource(element *rod.Element, selector string) (*[]byte, error) {
	return getImageResource(element.Sleeper(rod.NotFoundSleeper), selector)
}

func getText(element scrapeElement, selector string) (string, error) {
	textElement, err := element.Element(selector)
	if err != nil {
		return "", err
	}

	text, err := textElement.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}

func GetPageText(element *rod.Page, selector string) (string, error) {
	return getText(element.Sleeper(rod.NotFoundSleeper), selector)
}

func GetElementText(element *rod.Element, selector string) (string, error) {
	return getText(element.Sleeper(rod.NotFoundSleeper), selector)
}
