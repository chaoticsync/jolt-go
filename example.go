package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func main() {
	// Read input JSON
	inputData, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Read spec JSON
	specData, err := os.ReadFile("spec.json")
	if err != nil {
		fmt.Printf("Error reading spec file: %v\n", err)
		return
	}

	// Initialize Chainr with the spec
	chainr, err := transformer.NewChainr(string(specData))
	if err != nil {
		fmt.Printf("Error creating transformer: %v\n", err)
		return
	}

	// Apply transformations
	output, err := chainr.Apply(string(inputData))
	if err != nil {
		fmt.Printf("Error applying transformations: %v\n", err)
		return
	}

	// Pretty print the output
	var prettyOutput interface{}
	if err := json.Unmarshal([]byte(output), &prettyOutput); err != nil {
		fmt.Println(output) // If not valid JSON, print as is
	} else {
		prettyJSON, _ := json.MarshalIndent(prettyOutput, "", "  ")
		fmt.Println(string(prettyJSON))
	}
}
