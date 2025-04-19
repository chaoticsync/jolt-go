package transformer

// DefaultTransform implements the default transformation
type DefaultTransform struct {
	spec map[string]interface{}
}

// NewDefaultTransform creates a new default transformation
func NewDefaultTransform(spec map[string]interface{}) *DefaultTransform {
	return &DefaultTransform{
		spec: spec,
	}
}

// Apply applies the default transformation to the input
func (t *DefaultTransform) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	return t.applyDefaults(input, result, t.spec)
}

func (t *DefaultTransform) applyDefaults(input, output map[string]interface{}, spec map[string]interface{}) (map[string]interface{}, error) {
	// First copy all existing values from input to output
	for k, v := range input {
		output[k] = v
	}

	// Then apply defaults for missing values
	for key, defaultValue := range spec {
		if _, exists := output[key]; !exists {
			output[key] = defaultValue
		} else if inputMap, ok := output[key].(map[string]interface{}); ok {
			if specMap, ok := defaultValue.(map[string]interface{}); ok {
				subOutput := make(map[string]interface{})
				if _, err := t.applyDefaults(inputMap, subOutput, specMap); err != nil {
					return nil, err
				}
				output[key] = subOutput
			}
		} else if inputArray, ok := output[key].([]interface{}); ok {
			if specArray, ok := defaultValue.([]interface{}); ok && len(specArray) > 0 {
				subOutput := make([]interface{}, len(inputArray))
				for i, item := range inputArray {
					if itemMap, ok := item.(map[string]interface{}); ok {
						subMap := make(map[string]interface{})
						if _, err := t.applyDefaults(itemMap, subMap, specArray[0].(map[string]interface{})); err != nil {
							return nil, err
						}
						subOutput[i] = subMap
					} else {
						subOutput[i] = item
					}
				}
				output[key] = subOutput
			}
		}
	}
	return output, nil
}
