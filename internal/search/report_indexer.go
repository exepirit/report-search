package search

import (
	"github.com/exepirit/report-search/internal/data"
)

type ReportIndexer interface {
	Index(report data.Report) error
}
