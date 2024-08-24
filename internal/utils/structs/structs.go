package structs

type ScrapeResult[T any] struct {
	Value     T
	ScrapeErr error
}
