package fake

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/exepirit/report-search/internal/data"
	"github.com/google/uuid"
	"time"
)

type GofakeitGenerator struct{}

func (gen GofakeitGenerator) Generate() (data.Report, error) {
	return data.Report{
		ID:        uuid.New(),
		SubjectID: uuid.New(),
		SubjectName: gofakeit.RandString([]string{
			gofakeit.BeerName(),
			gofakeit.City(),
			gofakeit.Company(),
		}),
		Period: gen.GenerateReportPeriod(),
		Author: gen.GenerateUser(),
		Parts:  gen.GenerateReportParts(),
	}, nil
}

func (gen GofakeitGenerator) GenerateReportParts() []data.ReportPart {
	parts := make([]data.ReportPart, gofakeit.Number(1, 20))
	for i := 0; i < len(parts); i++ {
		parts[i] = gen.GenerateReportPart()
	}
	return parts
}

func (GofakeitGenerator) GenerateReportPart() data.ReportPart {
	return data.ReportPart{
		ID: uuid.New(),
		Content: gofakeit.Paragraph(
			gofakeit.Number(1, 4),  // paragraphCount
			gofakeit.Number(1, 5),  // sentenceCount
			gofakeit.Number(3, 15), // wordCount
			"\n\n",                 // separator
		),
	}
}

func (GofakeitGenerator) GenerateReportPeriod() data.ReportPeriod {
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

func (GofakeitGenerator) GenerateUser() data.User {
	return data.User{
		ID:        int(gofakeit.Int16()),
		ShortName: shortenName(gofakeit.FirstName(), gofakeit.LastName()),
	}
}

func shortenName(firstName string, lastName string) string {
	abbr := firstName[0]
	return fmt.Sprintf("%s %c.", lastName, abbr)
}
