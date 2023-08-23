package meilisearch

import "encoding/json"

func unmarshalHit[T any](hit interface{}, highlight bool) (T, error) {
	const highlightedAttribute = "_formatted"
	hitMap := hit.(map[string]any)

	if _, haveHighlight := hitMap[highlightedAttribute]; haveHighlight && highlight {
		hitMap = hitMap[highlightedAttribute].(map[string]any)
	}

	rawHit, _ := json.Marshal(hit)
	val := new(T)
	return *val, json.Unmarshal(rawHit, val)
}
