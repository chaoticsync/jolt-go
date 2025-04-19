package api

import (
	"encoding/json"
	"net/http"

	"github.com/chaoticsync/jolt-go/pkg/transformer"
)

// TransformHandler processes the JOLT transformation request
func TransformHandler(w http.ResponseWriter, r *http.Request) {
	// Define request structure
	var request struct {
		Input string `json:"input"`
		Spec  string `json:"spec"`
	}

	// Decode JSON request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	chainr, err := transformer.NewChainr(request.Spec)
	if err != nil {
		http.Error(w, "Invalid transformation spec: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Apply transformation
	result, err := chainr.Apply(request.Input)
	if err != nil {
		http.Error(w, "Transformation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers and return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
