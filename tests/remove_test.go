package tests

import (
	"testing"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func TestRemoveTransform(t *testing.T) {
	input := map[string]interface{}{
		"name":     "John",
		"password": "secret",
		"age":      30,
	}

	spec := map[string]interface{}{
		"password": nil,
	}

	remove := transformer.NewRemoveTransform(spec)
	result, err := remove.Apply(input)
	if err != nil {
		t.Fatalf("Error applying remove transform: %v", err)
	}

	expected := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	if !compareJSON(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
