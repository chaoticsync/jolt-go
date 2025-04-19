package transformer

import (
	"encoding/json"
	"strings"
)

// Chainr holds a sequence of transformations
type Chainr struct {
	Transforms []JoltTransform
}

// NewChainr initializes a new Chainr with a list of transformations
func NewChainr(specss string) (*Chainr, error) {
	var specs []map[string]interface{}
	err := json.Unmarshal([]byte(specss), &specs)
	if err != nil {
		return nil, err
	}
	var transforms []JoltTransform

	for _, spec := range specs {
		operation := TransformType(spec["operation"].(string))
		specValue := spec["spec"]
		var transformSpec *Spec

		switch operation {
		case Remove:
			// Convert array to map for remove operation
			if arr, ok := specValue.([]interface{}); ok {
				removeMap := make(map[string]interface{})
				for _, key := range arr {
					if strKey, ok := key.(string); ok {
						removeMap[strKey] = nil
					}
				}
				transformSpec = NewSpec(operation, removeMap)
			}
		case Sort:
			// Convert string to map for sort operation
			if str, ok := specValue.(string); ok {
				sortMap := make(map[string]interface{})
				sortMap["order"] = str
				transformSpec = NewSpec(operation, sortMap)
			} else if mapSpec, ok := specValue.(map[string]interface{}); ok {
				transformSpec = NewSpec(operation, mapSpec)
			}
		case Cardinality:
			// Convert string to map for cardinality operation
			if str, ok := specValue.(string); ok {
				cardMap := make(map[string]interface{})
				cardMap["*"] = str
				transformSpec = NewSpec(operation, cardMap)
			} else if mapSpec, ok := specValue.(map[string]interface{}); ok {
				transformSpec = NewSpec(operation, mapSpec)
			}
		default:
			if mapSpec, ok := specValue.(map[string]interface{}); ok {
				transformSpec = NewSpec(operation, mapSpec)
			}
		}

		if transformSpec == nil {
			return nil, ErrInvalidSpec
		}

		transform, err := NewTransform(transformSpec)
		if err != nil {
			return nil, err
		}
		transforms = append(transforms, transform)
	}

	return &Chainr{Transforms: transforms}, nil
}

// Apply executes transformations in sequence
func (c *Chainr) Apply(input string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", err
	}

	for _, transform := range c.Transforms {
		data, err = transform.Apply(data)
		if err != nil {
			return "", err
		}
	}

	// Convert array indices in map keys to proper array elements
	data = convertArrayIndices(data)

	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// convertArrayIndices converts map keys with array indices to proper array elements
func convertArrayIndices(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range data {
		if strings.Contains(key, "[") && strings.Contains(key, "]") {
			// Handle array index in key
			parts := strings.Split(key, "[")
			arrayName := parts[0]
			indexStr := strings.TrimRight(parts[1], "]")
			index := int(indexStr[0] - '0')

			// Create array if it doesn't exist
			if _, exists := result[arrayName]; !exists {
				result[arrayName] = make([]interface{}, index+1)
			} else if arr, ok := result[arrayName].([]interface{}); ok {
				// Extend array if needed
				if index >= len(arr) {
					newArr := make([]interface{}, index+1)
					copy(newArr, arr)
					result[arrayName] = newArr
				}
			}

			// Set value at index
			if arr, ok := result[arrayName].([]interface{}); ok {
				if index < len(arr) {
					if mapValue, ok := value.(map[string]interface{}); ok {
						arr[index] = convertArrayIndices(mapValue)
					} else {
						arr[index] = value
					}
				}
			}
		} else {
			// Handle non-array values
			if mapValue, ok := value.(map[string]interface{}); ok {
				result[key] = convertArrayIndices(mapValue)
			} else {
				result[key] = value
			}
		}
	}
	return result
}

func ParseInput(specss string) (map[string]interface{}, error) {
	specs := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(specss), &specs)
	if err != nil {
		return specs, err
	}
	return specs, nil
}

func ParseSpecs(specss string) ([]map[string]interface{}, error) {
	specs := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(specss), &specs)
	if err != nil {
		return specs, err
	}
	return specs, nil
}
