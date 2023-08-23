package search

import (
	"github.com/exepirit/report-search/internal/data"
	"time"
)

type ReportSearch interface {
	Query() ReportSearchQuery
}

type ReportSearchQuery interface {
	ContainsText(text string) ReportSearchQuery
	WrittenInPeriod(startDate, finishDate time.Time) ReportSearchQuery
	WithHighlights() ReportSearchQuery
	GetAll() ([]data.Report, error)
	GetN(n int) ([]data.Report, error)
}
