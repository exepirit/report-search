package search

import (
	"github.com/exepirit/report-search/internal/data"
	"github.com/exepirit/report-search/internal/search/index"
	"github.com/typesense/typesense-go/typesense"
)

type ReportIndexer interface {
	Index(report data.Report) error
}

type TypesenseReportIndexer struct {
	Client *typesense.Client
}

func (idx *TypesenseReportIndexer) Index(report data.Report) error {
	indexReport := index.MapReportToIndex(report)
	_, err := idx.Client.Collection(ReportsCollectionName).
		Documents().
		Create(indexReport)
	return err
}
