package typesense

import (
	"encoding/json"
	"fmt"
	"github.com/exepirit/report-search/internal/data"
	"github.com/exepirit/report-search/internal/search"
	"github.com/exepirit/report-search/internal/search/index"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"log/slog"
	"time"
)

type typesenseSearchQuery struct {
	client        *typesense.Client
	query         *api.SearchCollectionParams
	highlightHits bool
}

func (q *typesenseSearchQuery) ContainsText(text string) search.ReportSearchQuery {
	q.query.Q = text
	q.query.QueryBy = "subjectName, author.shortName, parts.content"
	return q
}

func (q *typesenseSearchQuery) WrittenInPeriod(startDate, finishDate time.Time) search.ReportSearchQuery {
	filterBy := fmt.Sprintf("period.deadline:>=%d && period.deadline:<=%d", startDate.Unix(), finishDate.Unix())
	q.query.FilterBy = &filterBy
	return q
}

func (q *typesenseSearchQuery) WithHighlights() search.ReportSearchQuery {
	q.highlightHits = true
	return q
}

func (q *typesenseSearchQuery) GetAll() ([]data.Report, error) {
	searchStart := time.Now()
	result, err := q.client.Collection(ReportsCollectionName).Documents().Search(q.query)
	if err != nil {
		return nil, err
	}
	slog.Info("Documents found",
		"latency", time.Now().Sub(searchStart).String(),
		"count", *result.Found,
		"returned", len(*result.Hits))

	reports := make([]data.Report, len(*result.Hits))
	for i, hit := range *result.Hits {
		document := *hit.Document
		if q.highlightHits && hit.Highlight != nil {
			document = ApplyHighlights(document, *hit.Highlight)
		}

		// dirty-dirty hack, but it works
		var indexedReport index.Report
		if err = mapToStruct(*hit.Document, &indexedReport); err != nil {
			return nil, err
		}

		reports[i], err = index.MapReportFromIndex(indexedReport)
		if err != nil {
			return nil, err
		}
	}
	return reports, nil
}

func (q *typesenseSearchQuery) GetN(n int) ([]data.Report, error) {
	page := 1
	perPage := n
	q.query.Page = &page
	q.query.PerPage = &perPage

	return q.GetAll()
}

func mapToStruct(m map[string]any, structRef any) error {
	s, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(s, structRef)
}
