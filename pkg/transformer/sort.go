package transformer

import (
	"sort"
)

// SortTransform implements the sort transformation
type SortTransform struct {
	spec map[string]interface{}
}

// NewSortTransform creates a new sort transformation
func NewSortTransform(spec map[string]interface{}) *SortTransform {
	return &SortTransform{
		spec: spec,
	}
}

// Apply applies the sort transformation to the input
func (t *SortTransform) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	order := "asc"
	if orderVal, ok := t.spec["order"]; ok {
		if orderStr, ok := orderVal.(string); ok {
			order = orderStr
		}
	}

	// Copy all values from input to result
	for k, v := range input {
		result[k] = v
	}

	// If spec is empty or has "order" key, sort all arrays
	if len(t.spec) == 0 || t.spec["order"] != nil {
		for k, v := range result {
			if arr, ok := v.([]interface{}); ok {
				sortedArr := make([]interface{}, len(arr))
				copy(sortedArr, arr)
				sort.Slice(sortedArr, func(i, j int) bool {
					if order == "asc" {
						return compareValues(sortedArr[i], sortedArr[j]) < 0
					}
					return compareValues(sortedArr[i], sortedArr[j]) > 0
				})
				result[k] = sortedArr
			}
		}
	} else {
		// Sort only specified fields
		for k := range t.spec {
			if v, ok := result[k]; ok {
				if arr, ok := v.([]interface{}); ok {
					sortedArr := make([]interface{}, len(arr))
					copy(sortedArr, arr)
					sort.Slice(sortedArr, func(i, j int) bool {
						if order == "asc" {
							return compareValues(sortedArr[i], sortedArr[j]) < 0
						}
						return compareValues(sortedArr[i], sortedArr[j]) > 0
					})
					result[k] = sortedArr
				}
			}
		}
	}

	return result, nil
}

func compareValues(a, b interface{}) int {
	switch aVal := a.(type) {
	case string:
		if bVal, ok := b.(string); ok {
			if aVal < bVal {
				return -1
			} else if aVal > bVal {
				return 1
			}
			return 0
		}
	case float64:
		if bVal, ok := b.(float64); ok {
			if aVal < bVal {
				return -1
			} else if aVal > bVal {
				return 1
			}
			return 0
		}
	case bool:
		if bVal, ok := b.(bool); ok {
			if !aVal && bVal {
				return -1
			} else if aVal && !bVal {
				return 1
			}
			return 0
		}
	case int:
		if bVal, ok := b.(int); ok {
			if aVal < bVal {
				return -1
			} else if aVal > bVal {
				return 1
			}
			return 0
		}
	}
	return 0
}
