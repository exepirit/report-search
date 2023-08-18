package search

import "github.com/typesense/typesense-go/typesense/api"

const ReportsCollectionName = "reports"

var ReportsCollectionSchema = &api.CollectionSchema{
	EnableNestedFields: boolRef(true),
	Fields: []api.Field{
		{
			Name: "id",
			Type: "string",
		},
		{
			Name: "subjectId",
			Type: "string",
		},
		{
			Name: "subjectName",
			Type: "string",
		},
		{
			Name: "period",
			Type: "object",
		},
		{
			Name: "author",
			Type: "object",
		},
		{
			Name: "parts",
			Type: "object[]",
		},
	},
	Name: "reports",
}

func boolRef(val bool) *bool {
	return &val
}
