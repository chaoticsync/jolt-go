package tests

import (
	"testing"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func TestChainTransform(t *testing.T) {
	input := map[string]interface{}{
		"name":     "John",
		"password": "secret",
		"age":      30,
	}

	// Create a chain of transformations
	shiftSpec := map[string]interface{}{
		"name": "person.name",
		"age":  "person.age",
	}
	removeSpec := map[string]interface{}{
		"password": nil,
	}

	shift := transformer.NewShiftTransform(shiftSpec)
	remove := transformer.NewRemoveTransform(removeSpec)

	chain := transformer.NewChainTransform(shift, remove)
	result, err := chain.Apply(input)
	if err != nil {
		t.Fatalf("Error applying chain transform: %v", err)
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
