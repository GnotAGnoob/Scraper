package errorsUtils

import (
	"errors"
	"fmt"

	"github.com/go-rod/rod"
)

func ElementNotFoundError(err error, selector string) error {
	fmt.Printf("Error: %v %v %T\n", err, selector, err)
	if _, ok := err.(*rod.ElementNotFoundError); !ok {
		text := fmt.Sprintf("Error: cannot find element with selector: '%s'", selector)
		return errors.New(text)
	}

	return err
}
