package typesense

import (
	"github.com/exepirit/report-search/internal/search"
	"github.com/exepirit/report-search/pkg/ref"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

type ReportSearch struct {
	Client             *typesense.Client
	HighlightThreshold int
}

func (search ReportSearch) Query() search.ReportSearchQuery {
	return &typesenseSearchQuery{
		client: search.Client,
		query: &api.SearchCollectionParams{
			SnippetThreshold: ref.Ref(search.HighlightThreshold),
		},
	}
}
