package errors

import (
	"errors"
	"fmt"

	"github.com/go-rod/rod"
)

func ElementNotFoundError(err error, selector string) error {
	if _, ok := err.(*rod.ElementNotFoundError); !ok {
		fmt.Printf("Error: %v\n", err)
		text := fmt.Sprintf("Error: cannot find element with selector: '%s'", selector)
		return errors.New(text)
	}

	return err
}
