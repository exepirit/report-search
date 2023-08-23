package meilisearch

import (
	"github.com/exepirit/report-search/internal/search"
	"github.com/meilisearch/meilisearch-go"
)

type Indexer[T search.Identifiable] struct {
	Client   *meilisearch.Client
	IndexKey string
}

func (indexer *Indexer[T]) Index(document T) error {
	_, err := indexer.Client.Index(indexer.IndexKey).AddDocuments(document, document.GetID())
	return err
}
