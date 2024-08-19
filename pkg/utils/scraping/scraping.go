package scraping

import (
	"time"

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

func GetImageResource(element scrapeElement, selector string) (*[]byte, error) {
	imageElement, err := element.Element(selector)
	if err != nil {
		return nil, err
	}

	image, err := imageElement.Timeout(200 * time.Millisecond).Resource()
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func GetText(element scrapeElement, selector string) (string, error) {
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
