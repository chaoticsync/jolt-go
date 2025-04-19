package transformer

// ChainTransform implements the chain transformation
type ChainTransform struct {
	transforms []JoltTransform
}

// NewChainTransform creates a new chain transformation
func NewChainTransform(transforms ...JoltTransform) *ChainTransform {
	return &ChainTransform{
		transforms: transforms,
	}
}

// Apply applies the chain of transformations to the input
func (t *ChainTransform) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{} = input
	var err error

	for _, transform := range t.transforms {
		result, err = transform.Apply(result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
