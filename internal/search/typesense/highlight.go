package typesense

import "fmt"

func ApplyHighlights(document map[string]any, highlights map[string]any) map[string]any {
	return applyStructHighlights(document, highlights)
}

func applyStructHighlights(document map[string]any, highlights map[string]any) map[string]any {
	for key, nestedHighlight := range highlights {
		docValue, ok := document[key]
		if !ok {
			continue
		}

		switch v := docValue.(type) {
		case map[string]any:
			document[key] = applyStructHighlights(v, nestedHighlight.(map[string]any))
		case []any:
			document[key] = applyArrayHighlights(v, nestedHighlight.([]any))
		case string:
			document[key] = applyStringHighlights(v, nestedHighlight.(map[string]any))
		default:
			panic(fmt.Errorf("highlight for %T is not supported yet", v))
		}
	}

	return document
}

func applyArrayHighlights(documents []any, highlights []any) []any {
	for i := 0; i < len(documents); i++ {
		docValue := documents[i]
		highlight := highlights[i]

		switch v := docValue.(type) {
		case map[string]any:
			documents[i] = applyStructHighlights(v, highlight.(map[string]any))
		case []any:
			documents[i] = applyArrayHighlights(v, highlight.([]any))
		case string:
			documents[i] = applyStringHighlights(v, highlight.(map[string]any))
		default:
			panic(fmt.Errorf("highlight for %T is not supported yet", v))
		}
	}
	return documents
}

func applyStringHighlights(document string, highlight map[string]any) string {
	matchedTokens := highlight["matched_tokens"].([]any)
	if len(matchedTokens) == 0 {
		return document
	}

	return highlight["snippet"].(string)
}
