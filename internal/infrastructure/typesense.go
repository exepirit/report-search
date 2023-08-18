package infrastructure

import "github.com/typesense/typesense-go/typesense"

var typesenseClient *typesense.Client = nil

func GetTypesenseClient() *typesense.Client {
	if typesenseClient == nil {
		panic("Typesense client is not set up yet")
	}
	return typesenseClient
}

func SetupTypesenseClient(url string, apiKey string) {
	typesenseClient = typesense.NewClient(
		typesense.WithServer(url),
		typesense.WithAPIKey(apiKey),
	)
}
