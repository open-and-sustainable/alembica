package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/extraction"
)

// TestMultiSequenceChat verifies the handling of multiple chat sequences with sequential prompts.
func TestMultiSequenceChat(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping multi-sequence chat tests in short mode")
	}

	// Load API keys from a file
	apiKeys, err := loadAPIKeys("test_keys.txt")
	if err != nil {
		t.Fatalf("Failed to load API keys: %v", err)
	}

	// Define input structure with multiple chat sequences
	input := definitions.Input{
		Metadata: definitions.InputMetadata{
			Version:       "1.0",
			SchemaVersion: "2024-03-16",
			Timestamp:     time.Now().Format(time.RFC3339),
		},
		Models: []definitions.Model{
			{
				Provider:    "OpenAI",
				APIKey:      apiKeys["OpenAI"],
				Model:       "gpt-4-turbo",
				Temperature: 0.7,
			},
			{
				Provider:    "Anthropic",
				APIKey:      apiKeys["Anthropic"],
				Model:       "claude-3-sonnet",
				Temperature: 0.7,
			},
		},
		Prompts: []definitions.Prompt{
			{PromptContent: "Hello, how are you?", SequenceID: "chat1", SequenceNumber: 1},
			{PromptContent: "What's your name?", SequenceID: "chat1", SequenceNumber: 2},
			{PromptContent: "Tell me a joke.", SequenceID: "chat2", SequenceNumber: 1},
			{PromptContent: "Give me another joke.", SequenceID: "chat2", SequenceNumber: 2},
		},
	}

	// Convert input to JSON
	inputJSON, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal input JSON: %v", err)
	}

	// Run the extraction function
	outputJSON, err := extraction.Extract(string(inputJSON))
	if err != nil {
		t.Fatalf("Failed to extract responses: %v", err)
	}

	// Parse output
	var output definitions.Output
	err = json.Unmarshal([]byte(outputJSON), &output)
	if err != nil {
		t.Fatalf("Failed to parse output JSON: %v", err)
	}

	// Verify output contains correct sequences and numbers
	expectedSequences := map[string]int{"chat1": 2, "chat2": 2}
	actualSequences := make(map[string]int)

	for _, response := range output.Responses {
		if _, exists := expectedSequences[response.SequenceID]; exists {
			actualSequences[response.SequenceID]++
		}
		if response.SequenceNumber < 1 {
			t.Errorf("Invalid sequence number for %s: %d", response.SequenceID, response.SequenceNumber)
		}
	}

	// Ensure expected sequence counts match
	for seq, expectedCount := range expectedSequences {
		if actualSequences[seq] != expectedCount {
			t.Errorf("Expected %d responses for %s, got %d", expectedCount, seq, actualSequences[seq])
		}
	}
}
