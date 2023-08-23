package meilisearch

import "encoding/json"

func unmarshalHit[T any](hit interface{}) (T, error) {
	rawHit, _ := json.Marshal(hit)
	val := new(T)
	return *val, json.Unmarshal(rawHit, val)
}
