package transformer

// ModifyTransform implements the modify transformation
type ModifyTransform struct {
	Spec map[string]interface{}
}

// NewModifyTransform creates a new modify transformation
func NewModifyTransform(spec map[string]interface{}) *ModifyTransform {
	return &ModifyTransform{
		Spec: spec,
	}
}

// Apply applies the modify transformation to the input
func (t *ModifyTransform) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for k, v := range input {
		result[k] = v
	}

	for key, operation := range t.Spec {
		if key == "*" {
			// Process all keys
			for k, v := range result {
				if err := t.processValue(k, v, operation, result); err != nil {
					return nil, err
				}
			}
		} else if val, exists := result[key]; exists {
			if err := t.processValue(key, val, operation, result); err != nil {
				return nil, err
			}
		} else if arrKey, isArray := t.isArrayKey(key); isArray {
			// Handle array wildcard syntax (e.g., "items[*].price")
			if arr, ok := result[arrKey].([]interface{}); ok {
				for i, item := range arr {
					if itemMap, ok := item.(map[string]interface{}); ok {
						if err := t.processValue(key, itemMap, operation, itemMap); err != nil {
							return nil, err
						}
						arr[i] = itemMap
					}
				}
				result[arrKey] = arr
			}
		}
	}

	return result, nil
}

func (t *ModifyTransform) isArrayKey(key string) (string, bool) {
	if len(key) > 3 && key[len(key)-3:] == "[*]" {
		return key[:len(key)-3], true
	}
	return "", false
}

func (t *ModifyTransform) processValue(key string, value interface{}, operation interface{}, result map[string]interface{}) error {
	switch op := operation.(type) {
	case string:
		if op == "@double" {
			if num, ok := value.(float64); ok {
				result[key] = num * 2
			} else if num, ok := value.(int); ok {
				result[key] = num * 2
			}
		}
	}
	return nil
}
