package main

import (
	"errors"
	"flag"
	"github.com/exepirit/report-search/internal/data"
	"github.com/exepirit/report-search/internal/fake"
	"github.com/exepirit/report-search/internal/search"
	meilisearch2 "github.com/exepirit/report-search/internal/search/meilisearch"
	"github.com/meilisearch/meilisearch-go"
	"log/slog"
	"os"
	"strconv"
	"time"
)

func main() {
	meilisearchURL := flag.String("meilisearch", "http://meilisearch:7700", "Meilisearch URL")
	meilisearchToken := flag.String("token", "apikey", "Meilisearch token")
	countStr := flag.String("count", "50", "Documents count")
	generatorType := flag.String("generator", "wikipedia", "Document generator type (wikipedia, gofakeit)")
	flag.Parse()

	count, err := strconv.Atoi(*countStr)
	if err != nil {
		slog.Error("Invalid documents count")
		os.Exit(1)
	}

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:    *meilisearchURL,
		APIKey:  *meilisearchToken,
		Timeout: 1 * time.Minute,
	})

	isCollectionExists, err := CheckCollectionExists(client, meilisearch2.ReportIndexKey)
	if err != nil {
		slog.Error("Cannot check index existence", "err", err)
		os.Exit(1)
	}

	if isCollectionExists {
		_, err = client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        meilisearch2.ReportIndexKey,
			PrimaryKey: "id",
		})
		if err != nil {
			slog.Error("Cannot delete index",
				"err", err)
		}
		slog.Info("Collection index")
	}

	_, err = client.Index(meilisearch2.ReportIndexKey).DeleteAllDocuments()
	if err != nil {
		slog.Error("Cannot clear index",
			"err", err)
		os.Exit(1)
	}
	slog.Info("Index created")

	var indexer search.Indexer[data.Report] = &meilisearch2.Indexer[data.Report]{
		Client:   client,
		IndexKey: meilisearch2.ReportIndexKey,
	}

	// fill with fake data
	counter := 0
	generator := SelectGenerator(*generatorType)
	fake.IterateReports(generator, count, func(report data.Report) {
		err = indexer.Index(report)
		if err != nil {
			slog.Error("Cannot index report", "err", err)
		}
		counter++
		slog.Info("Document indexed", "index", counter, "subjectName", report.SubjectName)
	})
	slog.Info("Documents indexed", "count", counter)
}

func CheckCollectionExists(client *meilisearch.Client, name string) (bool, error) {
	_, err := client.GetIndex(name)

	meilisearchErr := new(meilisearch.Error)
	if err != nil {
		if errors.As(err, &meilisearchErr) && meilisearchErr.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
