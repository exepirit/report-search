package meilisearch

import (
	"errors"
	"github.com/exepirit/report-search/internal/search"
	"github.com/meilisearch/meilisearch-go"
)

type Indexer[T search.Identifiable] struct {
	Client   *meilisearch.Client
	IndexKey string
}

func (indexer *Indexer[T]) Index(document T) error {
	taskInfo, err := indexer.Client.Index(indexer.IndexKey).
		AddDocuments(document, "id")

	indexTask, err := indexer.Client.WaitForTask(taskInfo.TaskUID)
	if err != nil {
		return err
	}

	if indexTask.Error.Code != "" {
		return errors.New(indexTask.Error.Message)
	}

	return nil
}
