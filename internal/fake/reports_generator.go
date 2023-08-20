package fake

import (
	"github.com/exepirit/report-search/internal/data"
	"log/slog"
)

type ReportsGenerator interface {
	Generate() (data.Report, error)
}

func IterateReports(generator ReportsGenerator, n int, cb func(report data.Report)) {
	for i := 0; i < n; i++ {
		report, err := generator.Generate()
		if err != nil {
			slog.Warn("Reports generation interrupted", "err", err)
			return
		}

		cb(report)
	}
}
