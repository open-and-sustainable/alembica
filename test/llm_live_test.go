package test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/llm/model"
	"github.com/open-and-sustainable/alembica/llm/check"
)

// Reads API keys from a file
func loadAPIKeys(filename string) (map[string]string, error) {
	keys := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open API key file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) == 2 {
			keys[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading API key file: %v", err)
	}

	return keys, nil
}

func TestLiveQueryLLM(t *testing.T) {
	if testing.Short() {
        t.Skip("Skipping live LLM query tests in short mode")
    }
	apiKeys, err := loadAPIKeys("test_keys.txt")
	if err != nil {
		t.Fatalf("Failed to load API keys: %v", err)
	}

	providers := []string{"OpenAI", "GoogleAI", "Cohere", "Anthropic", "DeepSeek"}

	prompts := []string{
		"Please provide a JSON response: { \"question\": \"What is the capital of France?\" }",
		"Respond only in JSON format: { \"request\": \"Tell me a joke.\" }",
	}
	
	queryService := model.DefaultQueryService{}

	for _, provider := range providers {
		apiKey, exists := apiKeys[provider]

		if !exists || apiKey == "" {
			t.Logf("Skipping %s: No API key found", provider)
			continue
		}

		// **Select the correct model for the provider**
		modelName := check.GetModel(prompts[0], provider, "", apiKey) 

		if modelName == "" {
			t.Logf("Skipping %s: No supported model found", provider)
			continue
		}

		t.Run(provider, func(t *testing.T) {
			llm := definitions.Model{
				Provider:    provider,
				APIKey:      apiKey,
				Model:       modelName, // ✅ **Uses dynamically selected model**
				Temperature: 0.7,
			}

			t.Logf("Testing %s with model %s...", provider, modelName)
			start := time.Now()
			responses, err := queryService.QueryLLM(prompts, llm)
			elapsed := time.Since(start)

			if err != nil {
				t.Errorf("Error querying %s: %v", provider, err)
			} else {
				t.Logf("%s Response (%.2fs): %v", provider, elapsed.Seconds(), responses)
			}
		})
	}
}

