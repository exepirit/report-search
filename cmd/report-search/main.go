package main

import (
	"flag"
	"github.com/exepirit/report-search/internal/api"
	"github.com/exepirit/report-search/internal/infrastructure"
	"github.com/gofiber/fiber/v2"
)

func main() {
	typesenseURL := flag.String("typesense", "http://typesense:80", "Typesense entrypoint URL")
	typesenseToken := flag.String("token", "apikey", "Typesense token")
	listenAddress := flag.String("listen", ":80", "")
	flag.Parse()

	infrastructure.SetupTypesenseClient(*typesenseURL, *typesenseToken)

	app := fiber.New(fiber.Config{
		Prefork:       false,
		ServerHeader:  "Fiber",
		StrictRouting: false,
		CaseSensitive: false,
		AppName:       "Report Search",
	})
	api.Bind(app.Group("/api"))

	err := app.Listen(*listenAddress)
	if err != nil {
		panic(err)
	}
}
