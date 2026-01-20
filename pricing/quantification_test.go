package pricing

import (
	"encoding/json"
	"testing"
)

func TestComputeCosts(t *testing.T) {
	// Sample input JSON
	inputJSON := `{
		"metadata": {
			"schemaVersion": "v1",
			"timestamp": "2025-02-17T12:00:00Z"
		},
		"models": [
			{
				"provider": "OpenAI",
				"model": "gpt-4-turbo",
				"temperature": 1.0
			}
		],
		"prompts": [
			{
				"promptContent": "What is AI?",
				"sequenceId": "seq1",
				"sequenceNumber": 1
			}
		]
	}`

	// Execute function
	resultJSON, err := ComputeCosts(inputJSON)
	if err != nil {
		t.Errorf("ComputeCosts failed: %v", err)
	}
	// Ensure output is valid JSON
	var result map[string]interface{}
	if err = json.Unmarshal([]byte(resultJSON), &result); err != nil {
		t.Errorf("Invalid JSON output: %v", err)
	}
	// Verify metadata section exists and has the correct schema version
	metadata, ok := result["metadata"].(map[string]interface{})
	if !ok {
		t.Fatalf("metadata missing or invalid")
	}
	if metadata["schemaVersion"] != "v1" {
		t.Errorf("expected schemaVersion v1, got %v", metadata["schemaVersion"])
	}

	// Ensure costs array is present and not empty
	costs, ok := result["costs"].([]interface{})
	if !ok {
		t.Fatalf("costs missing or invalid")
	}
	if len(costs) == 0 {
		t.Errorf("expected costs array to be non-empty")
	}

	// Optionally verify each cost entry contains expected keys
	for i, c := range costs {
		costMap, ok := c.(map[string]interface{})
		if !ok {
			t.Fatalf("cost entry %d has invalid type", i)
		}
		if _, ok := costMap["sequenceId"]; !ok {
			t.Errorf("cost entry %d missing sequenceId", i)
		}
		if _, ok := costMap["provider"]; !ok {
			t.Errorf("cost entry %d missing provider", i)
		}
		if _, ok := costMap["model"]; !ok {
			t.Errorf("cost entry %d missing model", i)
		}
		if _, ok := costMap["cost"]; !ok {
			t.Errorf("cost entry %d missing cost", i)
		}
	}
}
