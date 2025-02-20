package test

import (
	"encoding/json"
	"testing"
	"time"
	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/extraction"
)

func TestLiveExtraction(t *testing.T) {
	if testing.Short() {
        t.Skip("Skipping live extraction tests in short mode")
    }
	apiKeys, err := loadAPIKeys("test_keys.txt")
	if err != nil {
		t.Fatalf("Failed to load API keys: %v", err)
	}

	if err := definitions.LoadSchema("v1", "output"); err != nil {
		t.Fatalf("Failed to load output schema: %v", err)
	}

	providers := []string{"OpenAI", "GoogleAI", "Cohere", "Anthropic", "DeepSeek"}

	for _, provider := range providers {
		apiKey, exists := apiKeys[provider]
		if !exists || apiKey == "" {
			t.Logf("Skipping %s: No API key found", provider)
			continue
		}

		input := definitions.Input{
			Metadata: definitions.InputMetadata{
				SchemaVersion: "v1",
			},
			Models: []definitions.Model{
				{
					Provider:    provider,
					APIKey:      apiKey,
					Model:       "gpt-4o", // Replace with actual model lookup if needed
					Temperature: 0.7,
				},
			},
			Prompts: []definitions.Prompt{
				{PromptContent: "Respond strictly in JSON format: { \"request\": \"Test prompt\" }", SequenceID: "seq1", SequenceNumber: 1},
			},
		}

		inputJSON, err := json.Marshal(input)
		if err != nil {
			t.Fatalf("Failed to marshal input: %v", err)
		}

		t.Run(provider, func(t *testing.T) {
			t.Logf("Testing %s...", provider)
			start := time.Now()
			output, err := extraction.Extract(string(inputJSON))
			elapsed := time.Since(start)

			if err != nil {
				t.Errorf("Error querying %s: %v", provider, err)
			} else {
				t.Logf("%s Response (%.2fs): %v", provider, elapsed.Seconds(), output)
			}
		})
	}
}
