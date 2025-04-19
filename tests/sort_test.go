package tests

import (
	"testing"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func TestSortTransform(t *testing.T) {
	input := map[string]interface{}{
		"numbers": []interface{}{3, 1, 4, 1, 5, 9},
		"names":   []interface{}{"Charlie", "Alice", "Bob"},
	}

	spec := map[string]interface{}{
		"numbers": map[string]interface{}{},
		"names":   map[string]interface{}{},
	}

	sortT := transformer.NewSortTransform(spec)
	result, err := sortT.Apply(input)
	if err != nil {
		t.Fatalf("Error applying sort transform: %v", err)
	}

	expected := map[string]interface{}{
		"numbers": []interface{}{1, 1, 3, 4, 5, 9},
		"names":   []interface{}{"Alice", "Bob", "Charlie"},
	}

	if !compareJSON(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
