package main

import "github.com/exepirit/report-search/internal/fake"

func SelectGenerator(generatorType string) fake.ReportsGenerator {
	switch generatorType {
	case "wikipedia":
		return fake.WikipediaGenerator{
			Lang:              "ru",
			ArticlesNamespace: 0,
		}
	case "gofakeit":
		fallthrough
	default:
		return fake.GofakeitGenerator{}
	}
}
