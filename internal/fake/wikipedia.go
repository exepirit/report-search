package fake

import (
	"cgt.name/pkg/go-mwclient"
	"errors"
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/exepirit/report-search/internal/data"
	"github.com/google/uuid"
	"strconv"
)

type WikipediaGenerator struct {
	Lang              string
	ArticlesNamespace int

	client *mwclient.Client
}

func (gen WikipediaGenerator) Generate() (data.Report, error) {
	gen.client = gen.makeClient()

	pages, err := gen.selectRandomPages(1)
	if err != nil {
		return data.Report{}, nil
	}
	pageTitle, _ := pages[0].GetString("title")

	article, err := gen.downloadArticleContent(pageTitle)
	if err != nil {
		return data.Report{}, err
	}

	return gen.makeReport(pageTitle, article), nil
}

func (gen WikipediaGenerator) selectRandomPages(n int) ([]*jason.Object, error) {
	query := map[string]string{
		"action":      "query",
		"list":        "random",
		"rnnamespace": strconv.Itoa(gen.ArticlesNamespace),
		"rnlimit":     strconv.Itoa(n),
	}

	resp, err := gen.client.Get(query)
	if err != nil {
		return nil, err
	}

	return resp.GetObjectArray("query", "random")
}

func (gen WikipediaGenerator) downloadArticleContent(title string) ([]string, error) {
	query := map[string]string{
		"action": "parse",
		"prop":   "sections",
		"page":   title,
	}
	resp, err := gen.client.Get(query)
	if err != nil {
		return nil, err
	}
	sections, _ := resp.GetObjectArray("parse", "sections")

	if len(sections) == 0 {
		return nil, errors.New("empty page sections")
	}
	firstSectionId, _ := sections[0].GetString("index")

	query = map[string]string{
		"action":             "parse",
		"prop":               "text",
		"page":               title,
		"section":            firstSectionId,
		"disableeditsection": "true",
		"disablelimitreport": "true",
	}
	resp, err = gen.client.Get(query)
	if err != nil {
		return nil, err
	}

	content, err := resp.GetString("parse", "text")
	return []string{content}, err
}

func (WikipediaGenerator) makeReport(title string, content []string) data.Report {
	reportParts := make([]data.ReportPart, len(content))
	for i, text := range content {
		reportParts[i] = data.ReportPart{
			ID:      uuid.New(),
			Content: text,
		}
	}

	gofakeitGen := GofakeitGenerator{}

	return data.Report{
		ID:          uuid.New(),
		SubjectID:   uuid.New(),
		SubjectName: title,
		Period:      gofakeitGen.GenerateReportPeriod(),
		Author:      gofakeitGen.GenerateUser(),
		Parts:       reportParts,
	}
}

func (gen WikipediaGenerator) makeClient() *mwclient.Client {
	url := fmt.Sprintf("https://%s.wikipedia.org/w/api.php", gen.Lang)
	client, _ := mwclient.New(url, "go-mwclient")
	return client
}
