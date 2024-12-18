package test

import (
	"encoding/json"
	"os"
	"testing"
)

// LoadJSON loads a JSON file and unmarshals it into a provided struct
func LoadJSON(t *testing.T, path string, target interface{}) {
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	// Decode the JSON into the provided target struct
	if err := json.NewDecoder(file).Decode(target); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
}
