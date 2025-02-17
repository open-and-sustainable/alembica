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
	resultJSON := ComputeCosts(inputJSON)

	// Ensure output is valid JSON
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(resultJSON), &result); err != nil {
		t.Errorf("Invalid JSON output: %v", err)
	}
}
