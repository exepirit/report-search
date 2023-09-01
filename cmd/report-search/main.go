package main

import (
	"flag"
	"github.com/exepirit/report-search/internal/api"
	"github.com/exepirit/report-search/internal/infrastructure"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	typesenseURL := flag.String("typesense", "http://typesense:80", "Typesense entrypoint URL")
	apiToken := flag.String("token", "apikey", "API token")
	meilisearchURL := flag.String("meilisearch", "http://meilisearch:7700", "Meilisearch URL")
	listenAddress := flag.String("listen", ":80", "")
	flag.Parse()

	infrastructure.SetupTypesenseClient(*typesenseURL, *apiToken)
	infrastructure.SetupMeilisearchClient(*meilisearchURL, *apiToken)

	app := fiber.New(fiber.Config{
		Prefork:       false,
		ServerHeader:  "Fiber",
		StrictRouting: false,
		CaseSensitive: false,
		AppName:       "Report Search",
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Static("/", "./web/build")
	api.Bind(app.Group("/api"))

	err := app.Listen(*listenAddress)
	if err != nil {
		panic(err)
	}
}
