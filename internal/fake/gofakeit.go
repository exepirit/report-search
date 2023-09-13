package fake

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/exepirit/report-search/internal/data"
	"github.com/google/uuid"
	"math/rand"
	"net/url"
	"strings"
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
	parts := make([]data.ReportPart, gofakeit.Number(1, 15))
	for i := 0; i < len(parts); i++ {
		parts[i] = gen.GenerateReportPart()
	}
	return parts
}

func (gen GofakeitGenerator) GenerateReportPart() data.ReportPart {
	switch rand.Int() % 2 {
	case 0:
		return gen.GenerateTextReportPart()
	case 1:
		return gen.GenerateImageReportPart()
	default:
		panic("undefined behaviour")
	}
}

func (GofakeitGenerator) GenerateTextReportPart() data.ReportPart {
	paragraphs := make([]string, 2)
	for i := 0; i < len(paragraphs); i++ {
		paragraphs[i] = gofakeit.Paragraph(
			1,                      // paragraphCount
			gofakeit.Number(1, 5),  // sentenceCount
			gofakeit.Number(3, 10), // wordCount,
			"",                     // separator
		)
		paragraphs[i] = fmt.Sprintf("<p>%s</p>", paragraphs[i])
	}

	return data.TextReportPart{
		BaseReportPart: data.BaseReportPart{
			ID:   uuid.New(),
			Type: data.ReportPartTypeText,
		},
		Content: strings.Join(paragraphs, ""),
	}
}

func (GofakeitGenerator) GenerateImageReportPart() data.ReportPart {
	return data.ImageReportPart{
		BaseReportPart: data.BaseReportPart{
			ID:   uuid.New(),
			Type: data.ReportPartTypeImage,
		},
		URL:   "https://via.placeholder.com/120x120&text=" + url.PathEscape(gofakeit.Sentence(2)),
		Label: gofakeit.Sentence(3),
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
