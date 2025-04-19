package transformer

import "errors"

var (
	// ErrInvalidTransformType is returned when an invalid transform type is specified
	ErrInvalidTransformType = errors.New("invalid transform type")

	// ErrInvalidSpec is returned when an invalid specification is provided
	ErrInvalidSpec = errors.New("invalid specification")

	// ErrInvalidInput is returned when the input data is invalid
	ErrInvalidInput = errors.New("invalid input data")

	// ErrInvalidCardinality is returned when an invalid cardinality is specified
	ErrInvalidCardinality = errors.New("invalid cardinality, must be ONE or MANY")

	// ErrInvalidArrayReference is returned when an invalid array reference is specified in the specification
	ErrInvalidArrayReference = errors.New("invalid array reference in specification")
)
