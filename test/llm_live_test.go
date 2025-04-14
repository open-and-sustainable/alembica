package test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/llm/check"
	"github.com/open-and-sustainable/alembica/llm/model"
	"github.com/open-and-sustainable/alembica/utils/logger"
)

// loadAPIKeys reads API keys from environment variables first, and if any are missing,
// it falls back to reading from the provided file (which is expected to contain lines in KEY=VALUE format).
func loadAPIKeys(filename string) (map[string]string, error) {
	// Map the expected API key environment variable names to the provider names.
	envKeyToProvider := map[string]string{
		"OPENAI_API_KEY":    "OpenAI",
		"GOOGLE_AI_API_KEY": "GoogleAI",
		"CO_API_KEY":        "Cohere",
		"ANTHROPIC_API_KEY": "Anthropic",
		"DEEPSEEK_API_KEY":  "DeepSeek",
	}

	// Initialize a map to store the API keys indexed by provider name.
	providerKeys := make(map[string]string)

	// Retrieve keys from environment variables and map them to providers.
	for envKey, provider := range envKeyToProvider {
		if value := os.Getenv(envKey); value != "" {
			providerKeys[provider] = value
		}
	}

	// Check if we need to load keys from the file due to any missing keys.
	needFile := false
	for _, provider := range envKeyToProvider {
		if _, exists := providerKeys[provider]; !exists {
			needFile = true
			break
		}
	}

	// If we have all keys or no filename, return the current keys.
	if !needFile || filename == "" {
		return providerKeys, nil
	}

	// Attempt to open the file to read additional keys.
	file, err := os.Open(filename)
	if err != nil {
		// If file doesn't exist, return the keys from the environment.
		if os.IsNotExist(err) {
			return providerKeys, nil
		}
		return nil, fmt.Errorf("failed to open API key file: %v", err)
	}
	defer file.Close()

	// Parse the file line by line for KEY=VALUE pairs.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if provider, exists := envKeyToProvider[key]; exists {
				// Only add the key from the file if it's not already set.
				if _, exists := providerKeys[provider]; !exists {
					providerKeys[provider] = value
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading API key file: %v", err)
	}

	return providerKeys, nil
}

func TestLiveQueryLLM(t *testing.T) {
	logger.SetupLogging(logger.Silent, "")
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
