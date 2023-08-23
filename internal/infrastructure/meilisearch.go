package infrastructure

import (
	"github.com/meilisearch/meilisearch-go"
	"time"
)

var meilisearchClient *meilisearch.Client = nil

func GetMeilisearchClient() *meilisearch.Client {
	if meilisearchClient == nil {
		panic("Meilisearch client is not set up yet")
	}
	return meilisearchClient
}

func SetupMeilisearchClient(url string, apiKey string) {
	meilisearchClient = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:    url,
		APIKey:  apiKey,
		Timeout: 5 * time.Second,
	})
}
