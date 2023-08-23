package search

type Indexer[T Identifiable] interface {
	Index(document T) error
}

type Identifiable interface {
	GetID() string
}
