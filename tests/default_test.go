package tests

import (
	"testing"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func TestDefaultTransform(t *testing.T) {
	input := map[string]interface{}{
		"name": "John",
	}

	spec := map[string]interface{}{
		"status": "active",
		"role":   "user",
	}

	defaultT := transformer.NewDefaultTransform(spec)
	result, err := defaultT.Apply(input)
	if err != nil {
		t.Fatalf("Error applying default transform: %v", err)
	}

	expected := map[string]interface{}{
		"name":   "John",
		"status": "active",
		"role":   "user",
	}

	if !compareJSON(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
