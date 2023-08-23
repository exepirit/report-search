package meilisearch

import (
	"errors"
	"fmt"
	"github.com/exepirit/report-search/internal/data"
	"github.com/exepirit/report-search/internal/search"
	"github.com/meilisearch/meilisearch-go"
	"github.com/samber/mo"
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

	reports := make([]data.Report, 0, len(result.Hits))
	for _, hit := range result.Hits {
		report, err := unmarshalHit[data.Report](hit)
		if err != nil {
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
