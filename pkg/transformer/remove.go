package transformer

// RemoveTransform implements the remove transformation
type RemoveTransform struct {
	spec map[string]interface{}
}

// NewRemoveTransform creates a new remove transformation
func NewRemoveTransform(spec map[string]interface{}) *RemoveTransform {
	return &RemoveTransform{
		spec: spec,
	}
}

// Apply applies the remove transformation to the input
func (t *RemoveTransform) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	return t.remove(input, result, t.spec)
}

func (t *RemoveTransform) remove(input, output map[string]interface{}, spec map[string]interface{}) (map[string]interface{}, error) {
	// First copy all values from input to output
	for k, v := range input {
		output[k] = v
	}

	// Then remove specified fields
	for key, value := range spec {
		if key == "*" {
			// Remove all keys in input
			for k := range input {
				delete(output, k)
			}
		} else if _, exists := output[key]; exists {
			if value == nil {
				// Simple key removal
				delete(output, key)
			} else if specMap, ok := value.(map[string]interface{}); ok {
				// Nested removal
				if inputMap, ok := output[key].(map[string]interface{}); ok {
					subOutput := make(map[string]interface{})
					if _, err := t.remove(inputMap, subOutput, specMap); err != nil {
						return nil, err
					}
					output[key] = subOutput
				}
			} else if specArray, ok := value.([]interface{}); ok {
				// Array removal
				if inputArray, ok := output[key].([]interface{}); ok {
					subOutput := make([]interface{}, len(inputArray))
					for i, item := range inputArray {
						if itemMap, ok := item.(map[string]interface{}); ok {
							subMap := make(map[string]interface{})
							if _, err := t.remove(itemMap, subMap, specArray[0].(map[string]interface{})); err != nil {
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
	}
	return output, nil
}
