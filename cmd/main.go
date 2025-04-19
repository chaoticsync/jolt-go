package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: jolt-go <input-file> <spec-file>")
		fmt.Println("Example: jolt-go input.json spec.json")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	specFile := os.Args[2]

	// Read input JSON
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Read spec JSON
	specData, err := os.ReadFile(specFile)
	if err != nil {
		fmt.Printf("Error reading spec file: %v\n", err)
		os.Exit(1)
	}

	// Initialize Chainr
	chainr, err := transformer.NewChainr(string(specData))
	if err != nil {
		fmt.Printf("Error creating transformer: %v\n", err)
		os.Exit(1)
	}

	// Apply transformations
	output, err := chainr.Apply(string(inputData))
	if err != nil {
		fmt.Printf("Error applying transformations: %v\n", err)
		os.Exit(1)
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
