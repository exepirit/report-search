package meilisearch

import (
	"github.com/exepirit/report-search/internal/search"
	"github.com/meilisearch/meilisearch-go"
	"github.com/samber/mo"
)

type ReportSearch struct {
	client *meilisearch.Client
}

func (search *ReportSearch) Query() search.ReportSearchQuery {
	return &ReportSearchQuery{
		client: search.client,
		text:   mo.None[string](),
		searchRequest: &meilisearch.SearchRequest{
			AttributesToHighlight: []string{},
		},
	}
}
