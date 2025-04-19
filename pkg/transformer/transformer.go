package transformer

// JoltTransform represents a Jolt transformation operation
type JoltTransform interface {
	Apply(input map[string]interface{}) (map[string]interface{}, error)
}

// TransformType represents the type of Jolt transformation
type TransformType string

const (
	Shift       TransformType = "shift"
	Default     TransformType = "default"
	Remove      TransformType = "remove"
	Cardinality TransformType = "cardinality"
	Sort        TransformType = "sort"
	Modify      TransformType = "modify"
)

// Spec represents a Jolt transformation specification
type Spec struct {
	Operation TransformType
	Spec      map[string]interface{}
}

// NewSpec creates a new Jolt transformation specification
func NewSpec(operation TransformType, spec map[string]interface{}) *Spec {
	return &Spec{
		Operation: operation,
		Spec:      spec,
	}
}

// Apply applies the transformation specification to the input
func (s *Spec) Apply(input map[string]interface{}) (map[string]interface{}, error) {
	switch s.Operation {
	case Shift:
		return NewShiftTransform(s.Spec).Apply(input)
	case Default:
		return NewDefaultTransform(s.Spec).Apply(input)
	case Remove:
		return NewRemoveTransform(s.Spec).Apply(input)
	case Cardinality:
		return NewCardinalityTransform(s.Spec).Apply(input)
	case Sort:
		return NewSortTransform(s.Spec).Apply(input)
	default:
		return nil, ErrInvalidTransformType
	}
}
