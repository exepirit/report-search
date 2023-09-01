package meilisearch

import (
	"github.com/exepirit/report-search/internal/search"
	"github.com/meilisearch/meilisearch-go"
	"github.com/samber/mo"
)

type ReportSearch struct {
	Client *meilisearch.Client
}

func (search *ReportSearch) Query() search.ReportSearchQuery {
	return &ReportSearchQuery{
		client: search.Client,
		text:   mo.None[string](),
		searchRequest: &meilisearch.SearchRequest{
			AttributesToHighlight: []string{},
		},
	}
}
