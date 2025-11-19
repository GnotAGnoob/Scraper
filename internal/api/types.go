package api

import "github.com/GnotAGnoob/kosik-scraper/internal/service"

type RecipeRequest struct {
	Text string `json:"text"`
}

type RecipeResponse struct {
	Products []service.MatchedProduct `json:"products"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

