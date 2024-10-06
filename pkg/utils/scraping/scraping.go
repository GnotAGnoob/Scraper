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

// todo fix - possibly out of memory or something. the exceeded deadline always happens to the last products
// func GetImageResource(element scrapeElement, selector string) (*[]byte, error) {
// 	imageElement, err := element.Element(selector)
// 	if err != nil {
// 		return nil, err
// 	}

// 	image, err := imageElement.Timeout(200 * time.Millisecond).Resource()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &image, nil
// }

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

func IsElementNotFound(err error) bool {
	_, ok := err.(*rod.ElementNotFoundError)

	return err != nil && !ok
}

func RaceSelectors(page *rod.Page, timeout time.Duration, selectors ...string) (*rod.Element, error) {
	race := page.Timeout(timeout).Race()

	for _, selector := range selectors {
		race.Element(selector).Handle(func(_ *rod.Element) error {
			return nil
		})
	}

	return race.Do()
}
