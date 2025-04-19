package transformer

// CardinalityTransform implements the cardinality transformation
type CardinalityTransform struct {
	Spec map[string]string `json:"spec"`
}

// NewCardinalityTransform creates a new cardinality transformation
func NewCardinalityTransform(spec map[string]interface{}) *CardinalityTransform {
	// Convert spec to map[string]string
	strSpec := make(map[string]string)
	for k, v := range spec {
		if str, ok := v.(string); ok {
			strSpec[k] = str
		}
	}
	return &CardinalityTransform{
		Spec: strSpec,
	}
}

// Apply applies the cardinality transformation to the input
func (t *CardinalityTransform) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	for k, v := range input {
		output[k] = v
	}

	for path, cardinality := range t.Spec {
		if value, exists := output[path]; exists {
			if slice, ok := value.([]interface{}); ok {
				switch cardinality {
				case "ONE":
					if len(slice) > 0 {
						output[path] = []interface{}{slice[0]}
					}
				case "MANY":
					// Keep all elements
					output[path] = slice
				}
			} else if m, ok := value.(map[string]interface{}); ok {
				// Handle nested maps
				for k, v := range m {
					if slice, ok := v.([]interface{}); ok {
						switch cardinality {
						case "ONE":
							if len(slice) > 0 {
								m[k] = []interface{}{slice[0]}
							}
						case "MANY":
							// Keep all elements
							m[k] = slice
						}
					}
				}
				output[path] = m
			}
		}
	}

	return output, nil
}

func (t *CardinalityTransform) applyCardinality(input, output map[string]interface{}, spec map[string]interface{}) (map[string]interface{}, error) {
	for key, value := range spec {
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

func (t *CardinalityTransform) processValue(key string, value interface{}, spec interface{}, output map[string]interface{}) error {
	switch v := value.(type) {
	case []interface{}:
		cardinality := spec.(string)
		switch cardinality {
		case "ONE":
			if len(v) > 0 {
				output[key] = v[0]
			}
		case "MANY":
			output[key] = v
		}
	case map[string]interface{}:
		if specMap, ok := spec.(map[string]interface{}); ok {
			subOutput := make(map[string]interface{})
			if _, err := t.applyCardinality(v, subOutput, specMap); err != nil {
				return err
			}
			output[key] = subOutput
		} else {
			output[key] = v
		}
	default:
		output[key] = v
	}
	return nil
}
