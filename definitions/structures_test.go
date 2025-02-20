package definitions

import (
	"encoding/json"
	"testing"
)

// Test marshaling and unmarshaling of Input structure
func TestInputJSON(t *testing.T) {
	inputJSON := `{
		"metadata": {
			"version": "1.0",
			"schemaVersion": "v1",
			"timestamp": "2025-02-15T12:00:00Z"
		},
		"models": [
			{
				"provider": "OpenAI",
				"api_key": "test-key",
				"model": "gpt-4o",
				"temperature": 0.7,
				"tpm_limit": 1000,
				"rpm_limit": 2000
			}
		],
		"prompts": [
			{
				"promptContent": "Hello",
				"sequenceId": "123",
				"sequenceNumber": 1
			}
		]
	}`

	var input Input
	err := json.Unmarshal([]byte(inputJSON), &input)
	if err != nil {
		t.Fatalf("Failed to unmarshal Input JSON: %v", err)
	}

	// Validate key fields
	if input.Metadata.Version != "1.0" || input.Metadata.SchemaVersion != "v1" {
		t.Errorf("Unexpected metadata values: %+v", input.Metadata)
	}

	if len(input.Models) != 1 || input.Models[0].Provider != "OpenAI" {
		t.Errorf("Unexpected model values: %+v", input.Models)
	}

	if len(input.Prompts) != 1 || input.Prompts[0].PromptContent != "Hello" {
		t.Errorf("Unexpected prompt values: %+v", input.Prompts)
	}

	// Marshal back to JSON to test correctness
	marshaled, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal Input back to JSON: %v", err)
	}

	// Ensure no data loss after roundtrip
	var inputCheck Input
	err = json.Unmarshal(marshaled, &inputCheck)
	if err != nil {
		t.Fatalf("Failed to unmarshal marshaled Input JSON: %v", err)
	}
}

// Test marshaling and unmarshaling of Output structure
func TestOutputJSON(t *testing.T) {
	outputJSON := `{
		"metadata": {
			"schemaVersion": "v1"
		},
		"responses": [
			{
				"sequenceId": "123",
				"sequenceNumber": 1,
				"modelResponse": "This is a response"
			},
			{
				"sequenceId": "456",
				"sequenceNumber": 2,
				"modelResponse": "Another response",
				"error": {
					"code": 500,
					"message": "Internal error"
				}
			}
		]
	}`

	var output Output
	err := json.Unmarshal([]byte(outputJSON), &output)
	if err != nil {
		t.Fatalf("Failed to unmarshal Output JSON: %v", err)
	}

	// Validate metadata
	if output.Metadata.SchemaVersion != "v1" {
		t.Errorf("Unexpected metadata value: %+v", output.Metadata)
	}

	// Validate responses
	if len(output.Responses) != 2 {
		t.Fatalf("Expected 2 responses, got %d", len(output.Responses))
	}

	if output.Responses[1].Error == nil || output.Responses[1].Error.Code != 500 {
		t.Errorf("Expected error in second response but got: %+v", output.Responses[1].Error)
	}

	// Marshal back to JSON
	marshaled, err := json.Marshal(output)
	if err != nil {
		t.Fatalf("Failed to marshal Output back to JSON: %v", err)
	}

	// Ensure no data loss after roundtrip
	var outputCheck Output
	err = json.Unmarshal(marshaled, &outputCheck)
	if err != nil {
		t.Fatalf("Failed to unmarshal marshaled Output JSON: %v", err)
	}
}
