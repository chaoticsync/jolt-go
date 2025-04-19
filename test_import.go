package main

import (
	"encoding/json"
	"fmt"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func main() {
	// Define input JSON
	input := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	// Define shift specification
	shiftSpec := &transformer.Spec{
		Operation: transformer.Shift,
		Spec: map[string]interface{}{
			"name": "person.name",
			"age":  "person.age",
		},
	}

	// Create shift transform
	transform, err := transformer.NewTransform(shiftSpec)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Apply transformation
	result, err := transform.Apply(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print result
	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}
