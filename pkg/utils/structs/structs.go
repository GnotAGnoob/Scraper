package structs

type ErrorValue[T any] struct {
	Value T
	Err   error
}