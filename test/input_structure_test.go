package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/extraction"
	"github.com/open-and-sustainable/alembica/llm/check"
	"github.com/open-and-sustainable/alembica/utils/logger"
)

// TestMultiSequenceChat verifies the handling of multiple chat sequences with sequential prompts.
func TestMultiSequenceChat(t *testing.T) {
	// switch between logger.Silent and logger.Stdout
	logger.SetupLogging(logger.Silent, "")
	if testing.Short() {
		t.Skip("Skipping multi-sequence chat tests in short mode")
	}

	// Load API keys from ENV or  a file
	apiKeys, err := loadAPIKeys("test_keys.txt")
	if err != nil {
		t.Fatalf("Failed to load API keys: %v", err)
	}
	t.Logf("Loaded API Keys: %v\n", apiKeys)

	// Define input structure with multiple chat sequences
	input := definitions.Input{
		Metadata: definitions.InputMetadata{
			Version:       "1.0",
			SchemaVersion: "v1",
			Timestamp:     time.Now().Format(time.RFC3339),
		},
		Models: []definitions.Model{
			{
				Provider:    "OpenAI",
				APIKey:      apiKeys["OpenAI"],
				Model:       check.GetModel("", "OpenAI", "gpt-4-turbo", ""),
				Temperature: 0.7,
			},
			{
				Provider:    "Anthropic",
				APIKey:      apiKeys["Anthropic"],
				Model:       check.GetModel("", "Anthropic", "claude-3-sonnet", ""),
				Temperature: 0.7,
			},
		},
		Prompts: []definitions.Prompt{
			{PromptContent: "Hello, how are you? Respond with a JSON object.", SequenceID: "chat1", SequenceNumber: 1},
			{PromptContent: "What's your name? Respond with a JSON object.", SequenceID: "chat1", SequenceNumber: 2},
			{PromptContent: "Tell me a joke. Respond with a JSON object.", SequenceID: "chat2", SequenceNumber: 1},
			{PromptContent: "Give me another joke. Respond with a JSON object.", SequenceID: "chat2", SequenceNumber: 2},
		},
	}

	// Convert input to JSON
	inputJSON, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Failed to marshal input JSON: %v", err)
	}
	t.Logf("Marshalled Input JSON: %s\n", inputJSON)

	// Run the extraction function
	outputJSON, err := extraction.Extract(string(inputJSON))
	if err != nil {
		t.Fatalf("Failed to extract responses: %v", err)
	}
	t.Logf("Output JSON: %s\n", outputJSON)

	// Parse output
	var output definitions.Output
	err = json.Unmarshal([]byte(outputJSON), &output)
	if err != nil {
		t.Fatalf("Failed to parse output JSON: %v", err)
	}
	t.Logf("Parsed Output: %+v\n", output)

	// Verify output contains correct sequences and numbers
	expectedSequences := map[string]int{"chat1": 4, "chat2": 4}
	actualSequences := make(map[string]int)

	// Print each response to examine what you received
	for _, response := range output.Responses {
		t.Logf("Response received: %+v\n", response)
		if _, exists := expectedSequences[response.SequenceID]; exists {
			actualSequences[response.SequenceID]++
		}
		if response.SequenceNumber < 1 {
			t.Errorf("Invalid sequence number for %s: %d", response.SequenceID, response.SequenceNumber)
		}
	}

	// Ensure expected sequence counts match
	for seq, expectedCount := range expectedSequences {
		actualCount := actualSequences[seq]
		t.Logf("Expected responses for %s: %d, Actual: %d\n", seq, expectedCount, actualCount)
		if actualCount != expectedCount {
			t.Errorf("Expected %d responses for %s, got %d", expectedCount, seq, actualCount)
		}
	}
}
