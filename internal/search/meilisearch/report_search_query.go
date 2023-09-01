package meilisearch

import (
	"errors"
	"fmt"
	"github.com/exepirit/report-search/internal/data"
	"github.com/exepirit/report-search/internal/search"
	"github.com/exepirit/report-search/internal/search/index"
	"github.com/meilisearch/meilisearch-go"
	"github.com/samber/mo"
	"log/slog"
	"time"
)

type ReportSearchQuery struct {
	client        *meilisearch.Client
	text          mo.Option[string]
	searchRequest *meilisearch.SearchRequest
}

func (query *ReportSearchQuery) ContainsText(text string) search.ReportSearchQuery {
	query.text = mo.Some(text)
	return query
}

func (query *ReportSearchQuery) WrittenInPeriod(startDate, finishDate time.Time) search.ReportSearchQuery {
	query.searchRequest.Filter = fmt.Sprintf(
		"period.deadline >= %d AND period.deadline <= %d",
		startDate.Unix(), finishDate.Unix(),
	)
	return query
}

func (query *ReportSearchQuery) WithHighlights() search.ReportSearchQuery {
	query.searchRequest.AttributesToHighlight = []string{
		"subjectName",
		"author.shortName",
		"parts.content",
	}
	query.searchRequest.HighlightPreTag = "<mark>"
	query.searchRequest.HighlightPostTag = "</mark>"
	return query
}

func (query *ReportSearchQuery) GetAll() ([]data.Report, error) {
	if query.text.IsAbsent() {
		return nil, errors.New("search input text is required")
	}

	result, err := query.client.Index(ReportIndexKey).
		Search(query.text.MustGet(), query.searchRequest)
	if err != nil {
		return nil, err
	}

	withHighlight := query.searchRequest.AttributesToHighlight != nil
	reports := make([]data.Report, 0, len(result.Hits))
	for _, hit := range result.Hits {
		indexedReport, err := unmarshalHit[index.Report](hit, withHighlight)
		if err != nil {
			slog.Error("Invalid data in index", "err", err)
			return nil, errors.New("invalid data in index")
		}

		report, err := index.MapReportFromIndex(indexedReport)
		if err != nil {
			slog.Error("Invalid data in index", "err", err)
			return nil, errors.New("invalid data in index")
		}

		reports = append(reports, report)
	}
	return reports, nil
}

func (query *ReportSearchQuery) GetN(n int) ([]data.Report, error) {
	query.searchRequest.Limit = int64(n)
	return query.GetAll()
}
