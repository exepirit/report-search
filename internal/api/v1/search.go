package v1

import (
	"github.com/exepirit/report-search/internal/infrastructure"
	"github.com/exepirit/report-search/internal/search"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"net/http"
	"time"
)

func SearchReport(ctx *fiber.Ctx) error {
	text := ctx.Query("text")
	if text == "" {
		return ctx.Status(http.StatusBadRequest).
			JSON(map[string]any{
				"error": "query text must be provided",
			})
	}

	var reportSearch search.ReportSearch = &search.TypesenseReportSearch{
		Client:             infrastructure.GetTypesenseClient(),
		HighlightThreshold: 2000,
	}

	searchStart := time.Now()

	reports, err := reportSearch.
		Query().
		ContainsText(text).
		WithHighlights().
		GetN(20)
	if err != nil {
		return err
	}

	slog.Info("Fulltext search completed",
		"latency", time.Now().Sub(searchStart).String())

	return ctx.Status(http.StatusOK).JSON(reports)
}
