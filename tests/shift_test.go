package tests

import (
	"testing"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func TestShiftTransform(t *testing.T) {
	input := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	spec := map[string]interface{}{
		"name": "person.name",
		"age":  "person.age",
	}

	shift := transformer.NewShiftTransform(spec)
	result, err := shift.Apply(input)
	if err != nil {
		t.Fatalf("Error applying shift transform: %v", err)
	}

	expected := map[string]interface{}{
		"person": map[string]interface{}{
			"name": "John",
			"age":  30,
		},
	}

	if !compareJSON(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
