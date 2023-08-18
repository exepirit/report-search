package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/exepirit/report-search/internal/data"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func IterateGeneratedReports(count int, cb func(report data.Report)) {
	for i := 0; i < count; i++ {
		report := GenerateReport()
		cb(report)
		if i%200 == 0 {
			slog.Info("Report generating is pending", "count", i)
			slog.Info("Generated report example", "example", report)
		}
	}
}

func GenerateReport() data.Report {
	return data.Report{
		ID:          uuid.New(),
		SubjectID:   uuid.New(),
		SubjectName: gofakeit.Word(),
		Period:      GenerateReportPeriod(),
		Author:      GenerateUser(),
		Parts:       GenerateReportParts(),
	}
}

func GenerateReportParts() []data.ReportPart {
	parts := make([]data.ReportPart, gofakeit.Number(1, 20))
	for i := 0; i < len(parts); i++ {
		parts[i] = GenerateReportPart()
	}
	return parts
}

func GenerateReportPart() data.ReportPart {
	return data.ReportPart{
		ID: uuid.New(),
		Content: gofakeit.Paragraph(
			gofakeit.Number(1, 4),  // paragraphCount
			gofakeit.Number(1, 5),  // sentenceCount
			gofakeit.Number(3, 15), // wordCount
			".",                    // separator
		),
	}
}

func GenerateReportPeriod() data.ReportPeriod {
	startDate := gofakeit.Date()
	finishDate := startDate.AddDate(0, 0, 7).Add(-time.Second)
	deadline := gofakeit.DateRange(startDate, finishDate)
	return data.ReportPeriod{
		ID:         uuid.New(),
		StartDate:  startDate,
		FinishDate: finishDate,
		Deadline:   deadline,
	}
}

func GenerateUser() data.User {
	return data.User{
		ID:        int(gofakeit.Int16()),
		ShortName: ShortenName(gofakeit.FirstName(), gofakeit.LastName()),
	}
}

func ShortenName(firstName string, lastName string) string {
	abbr := firstName[0]
	return fmt.Sprintf("%s %c.", lastName, abbr)
}
