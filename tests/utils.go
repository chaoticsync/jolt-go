package tests

import "encoding/json"

// compareJSON compares two JSON objects for equality
func compareJSON(a, b interface{}) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}
