package search

type Result[T any] struct {
	Highlight string
	Value     T
}
