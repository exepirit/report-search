package main

import (
	"flag"
	"github.com/exepirit/report-search/internal/data"
	"github.com/exepirit/report-search/internal/fake"
	"github.com/exepirit/report-search/internal/search"
	"github.com/typesense/typesense-go/typesense"
	"log/slog"
	"os"
	"strconv"
)

func main() {
	typesenseURL := flag.String("typesense", "http://typesense:80", "Typesense entrypoint URL")
	typesenseToken := flag.String("token", "apikey", "Typesense token")
	countStr := flag.String("count", "2000", "Documents count")
	flag.Parse()

	count, err := strconv.Atoi(*countStr)
	if err != nil {
		slog.Error("Invalid documents count")
		os.Exit(1)
	}

	client := typesense.NewClient(
		typesense.WithServer(*typesenseURL),
		typesense.WithAPIKey(*typesenseToken),
	)

	isCollectionExists, err := CheckCollectionExists(client, search.ReportsCollectionName)
	if err != nil {
		slog.Error("Cannot check collection existence", "err", err)
		os.Exit(1)
	}

	if isCollectionExists {
		_, err = client.Collection(search.ReportsCollectionName).Delete()
		if err != nil {
			slog.Error("Cannot delete collection",
				"err", err)
		}
		slog.Info("Collection deleted")
	}

	_, err = client.Collections().Create(search.ReportsCollectionSchema)
	if err != nil {
		slog.Error("Cannot clear collection",
			"err", err)
		os.Exit(1)
	}
	slog.Info("Collection created")

	var indexer search.ReportIndexer = &search.TypesenseReportIndexer{
		Client: client,
	}

	// fill with fake data
	generator := fake.GofakeitGenerator{}
	fake.IterateReports(generator, count, func(report data.Report) {
		err = indexer.Index(report)
		if err != nil {
			slog.Error("Cannot index report", "err", err)
		}
	})
	slog.Info("Documents indexed", "count", count)
}

func CheckCollectionExists(client *typesense.Client, name string) (bool, error) {
	collections, err := client.Collections().Retrieve()
	if err != nil {
		return false, err
	}

	for _, collection := range collections {
		if collection.Name == name {
			return true, nil
		}
	}
	return false, nil
}
