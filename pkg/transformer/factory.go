package transformer

// NewTransform creates a new transform based on the spec
func NewTransform(spec *Spec) (JoltTransform, error) {
	switch spec.Operation {
	case Shift:
		return NewShiftTransform(spec.Spec), nil
	case Default:
		return NewDefaultTransform(spec.Spec), nil
	case Remove:
		return NewRemoveTransform(spec.Spec), nil
	case Cardinality:
		return NewCardinalityTransform(spec.Spec), nil
	case Sort:
		return NewSortTransform(spec.Spec), nil
	case Modify:
		return NewModifyTransform(spec.Spec), nil
	default:
		return nil, ErrInvalidTransformType
	}
}
