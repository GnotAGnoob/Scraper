package service

import "errors"

var (
	ErrEmptyRecipeText = errors.New("recipe text is required")
	ErrNoProductFound  = errors.New("no product found")
)
