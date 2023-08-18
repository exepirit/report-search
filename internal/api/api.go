package api

import (
	v1 "github.com/exepirit/report-search/internal/api/v1"
	"github.com/gofiber/fiber/v2"
)

func Bind(router fiber.Router) {
	v1api := router.Group("/v1")
	v1.Bind(v1api)
}
