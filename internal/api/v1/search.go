package v1

import (
	"github.com/exepirit/report-search/internal/infrastructure"
	"github.com/exepirit/report-search/internal/search"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
		Client: infrastructure.GetTypesenseClient(),
	}

	reports, err := reportSearch.
		Query().
		ContainsText(text).
		GetAll()
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(reports)
}
