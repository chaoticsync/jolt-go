package transformer

import (
	"strconv"
	"strings"
)

// ShiftTransform implements the shift transformation
type ShiftTransform struct {
	spec map[string]interface{}
}

// NewShiftTransform creates a new shift transformation
func NewShiftTransform(spec map[string]interface{}) *ShiftTransform {
	return &ShiftTransform{
		spec: spec,
	}
}

// Apply applies the shift transformation to the input
func (t *ShiftTransform) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	for key, value := range t.spec {
		if key == "*" {
			// Process all keys in input
			for k, v := range input {
				if err := t.processValue(k, v, value, output); err != nil {
					return nil, err
				}
			}
		} else if v, exists := input[key]; exists {
			if err := t.processValue(key, v, value, output); err != nil {
				return nil, err
			}
		}
	}
	return output, nil
}

func (t *ShiftTransform) processValue(key string, value interface{}, spec interface{}, output map[string]interface{}) error {
	switch v := value.(type) {
	case []interface{}:
		if specMap, ok := spec.(map[string]interface{}); ok {
			// Handle array transformation with wildcard
			if wildcardSpec, ok := specMap["*"].(map[string]interface{}); ok {
				result := make([]interface{}, len(v))
				for i, item := range v {
					if itemMap, ok := item.(map[string]interface{}); ok {
						resultItem := make(map[string]interface{})
						for specKey, specVal := range wildcardSpec {
							if _, ok := specVal.(string); ok {
								if val, exists := itemMap[specKey]; exists {
									resultItem[specKey] = val
								}
							}
						}
						result[i] = resultItem
					}
				}
				output["contactInfo"] = result
			} else {
				output[key] = v
			}
		} else {
			output[key] = v
		}
	case map[string]interface{}:
		if specMap, ok := spec.(map[string]interface{}); ok {
			subOutput := make(map[string]interface{})
			if err := t.processMap(v, specMap, subOutput); err != nil {
				return err
			}
			output[key] = subOutput
		} else if targetPath, ok := spec.(string); ok {
			t.setNestedValue(output, targetPath, v)
		}
	default:
		if targetPath, ok := spec.(string); ok {
			t.setNestedValue(output, targetPath, v)
		}
	}
	return nil
}

func (t *ShiftTransform) processMap(input, spec map[string]interface{}, output map[string]interface{}) error {
	for key, value := range spec {
		if key == "*" {
			// Process all keys in input
			for k, v := range input {
				if err := t.processValue(k, v, value, output); err != nil {
					return err
				}
			}
		} else if v, exists := input[key]; exists {
			if err := t.processValue(key, v, value, output); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *ShiftTransform) setNestedValue(output map[string]interface{}, path string, value interface{}) {
	// Check if the path contains array indexing
	if strings.Contains(path, "[") && strings.Contains(path, "]") {
		// Extract array path and index
		parts := strings.Split(path, "[")
		arrayPath := parts[0]
		indexStr := strings.TrimRight(parts[1], "]")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return
		}

		// Create or get the array
		var arr []interface{}
		if val, exists := output[arrayPath]; exists {
			if existingArr, ok := val.([]interface{}); ok {
				arr = existingArr
			} else {
				arr = make([]interface{}, index+1)
			}
		} else {
			arr = make([]interface{}, index+1)
		}

		// Ensure array has enough capacity
		if index >= len(arr) {
			newArr := make([]interface{}, index+1)
			copy(newArr, arr)
			arr = newArr
		}

		// Set the value at the specified index
		if index < len(arr) {
			if existingMap, ok := arr[index].(map[string]interface{}); ok {
				// If there's already a map at this index, merge the new value
				if newMap, ok := value.(map[string]interface{}); ok {
					for k, v := range newMap {
						existingMap[k] = v
					}
				} else {
					arr[index] = value
				}
			} else {
				arr[index] = value
			}
		}

		output[arrayPath] = arr
	} else {
		// Handle regular nested paths
		parts := strings.Split(path, ".")
		current := output
		for i, part := range parts {
			if i == len(parts)-1 {
				current[part] = value
			} else {
				if _, exists := current[part]; !exists {
					current[part] = make(map[string]interface{})
				}
				current = current[part].(map[string]interface{})
			}
		}
	}
}
