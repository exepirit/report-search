package fake

import (
	"github.com/exepirit/report-search/internal/data"
	"log/slog"
)

type ReportsGenerator interface {
	Generate() (data.Report, error)
}

func IterateReports(generator ReportsGenerator, n int, cb func(report data.Report)) {
	errorsCount := 0
	for i := 0; i < n && errorsCount < 3; {

		report, err := generator.Generate()
		if err != nil {
			slog.Warn("Error occurred while report generation", "err", err)
			errorsCount++
			continue
		}

		cb(report)
		errorsCount = 0
		i++
	}
}
