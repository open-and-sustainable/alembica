package test

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
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

	// Setup test data and run extraction
	_, output, err := setupMultiSequenceTest(t)
	if err != nil {
		t.Fatalf("Test setup failed: %v", err)
	}

	// Group responses by model and sequence for verification
	responsesByModelAndSequence := groupResponsesByModelAndSequence(output.Responses)

	// Verify the test results
	verifyResponseCounts(t, output.Responses)
	verifyContextMaintenance(t, responsesByModelAndSequence)
	verifyContextIsolation(t, responsesByModelAndSequence)
}

// setupMultiSequenceTest prepares and runs the test, returning the input and output data
func setupMultiSequenceTest(t *testing.T) (definitions.Input, definitions.Output, error) {
	// Load API keys from ENV or a file
	apiKeys, err := loadAPIKeys("test_keys.txt")
	if err != nil {
		return definitions.Input{}, definitions.Output{}, err
	}
	t.Logf("Loaded API Keys: %v\n", apiKeys)

	// Define input structure with multiple chat sequences that test contextual awareness
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
				Model:       check.GetModel("", "Anthropic", "claude-3-5-sonnet", ""),
				Temperature: 0.7,
			},
		},
		Prompts: []definitions.Prompt{
			// Sequence 1: Requires contextual memory
			{PromptContent: "My name is TestUser. Respond with a JSON object with greeting field.", SequenceID: "chat1", SequenceNumber: 1},
			{PromptContent: "What's my name? Include a 'remembered_name' field in your JSON response.", SequenceID: "chat1", SequenceNumber: 2},

			// Sequence 2: Different context
			{PromptContent: "The capital of France is Paris. Respond with a JSON object.", SequenceID: "chat2", SequenceNumber: 1},
			{PromptContent: "What's the capital of France? Include a 'capital' field in your response.", SequenceID: "chat2", SequenceNumber: 2},
		},
	}

	// Convert input to JSON
	inputJSON, err := json.Marshal(input)
	if err != nil {
		return definitions.Input{}, definitions.Output{}, fmt.Errorf("failed to marshal input JSON: %v", err)
	}
	t.Logf("Marshalled Input JSON: %s\n", inputJSON)

	// Run the extraction function
	outputJSON, err := extraction.Extract(string(inputJSON))
	if err != nil {
		return definitions.Input{}, definitions.Output{}, fmt.Errorf("failed to extract responses: %v", err)
	}
	t.Logf("Output JSON: %s\n", outputJSON)

	// Parse output
	var output definitions.Output
	err = json.Unmarshal([]byte(outputJSON), &output)
	if err != nil {
		return definitions.Input{}, definitions.Output{}, fmt.Errorf("failed to parse output JSON: %v", err)
	}
	t.Logf("Parsed Output: %+v\n", output)

	return input, output, nil
}

// groupResponsesByModelAndSequence organizes responses by model and sequence for easier verification
func groupResponsesByModelAndSequence(responses []definitions.Response) map[string]map[string][]definitions.Response {
	responsesByModelAndSequence := make(map[string]map[string][]definitions.Response)

	for _, response := range responses {
		modelKey := response.Provider + "-" + response.Model

		// Initialize nested maps if they don't exist
		if _, exists := responsesByModelAndSequence[modelKey]; !exists {
			responsesByModelAndSequence[modelKey] = make(map[string][]definitions.Response)
		}

		responsesByModelAndSequence[modelKey][response.SequenceID] = append(
			responsesByModelAndSequence[modelKey][response.SequenceID],
			response,
		)
	}

	// Sort responses by sequence number for each model and sequence
	for _, sequenceMap := range responsesByModelAndSequence {
		for _, responses := range sequenceMap {
			sort.Slice(responses, func(i, j int) bool {
				return responses[i].SequenceNumber < responses[j].SequenceNumber
			})
		}
	}

	return responsesByModelAndSequence
}

// verifyResponseCounts checks that we have the expected number of responses for each sequence
func verifyResponseCounts(t *testing.T, responses []definitions.Response) {
	expectedSequences := map[string]int{"chat1": 4, "chat2": 4} // 2 models x 2 prompts each = 4 responses per sequence
	actualSequences := make(map[string]int)

	// Print each response to examine what you received
	for _, response := range responses {
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

// verifyContextMaintenance checks that the model maintains context within a sequence
func verifyContextMaintenance(t *testing.T, responsesByModelAndSequence map[string]map[string][]definitions.Response) {
	for modelKey, sequenceMap := range responsesByModelAndSequence {
		// Check chat1 context (name remembering)
		if responses, ok := sequenceMap["chat1"]; ok && len(responses) >= 2 {
			verifyNameRemembered(t, modelKey, responses[1].ModelResponses[0])
		}

		// Check chat2 context (capital remembering)
		if responses, ok := sequenceMap["chat2"]; ok && len(responses) >= 2 {
			verifyCapitalRemembered(t, modelKey, responses[1].ModelResponses[0])
		}
	}
}

// verifyNameRemembered checks if the model remembered the name "TestUser"
func verifyNameRemembered(t *testing.T, modelKey, responseText string) {
	var responseObj map[string]interface{}

	// Try to parse the JSON response
	if err := json.Unmarshal([]byte(responseText), &responseObj); err == nil {
		// Check if remembered_name field contains TestUser
		if name, ok := responseObj["remembered_name"].(string); ok {
			if !strings.Contains(strings.ToLower(name), "testuser") {
				t.Errorf("[%s] Context not maintained in chat1: remembered_name=%s doesn't match 'TestUser'",
					modelKey, name)
			} else {
				t.Logf("[%s] Context successfully maintained in chat1", modelKey)
			}
		} else {
			t.Errorf("[%s] Missing 'remembered_name' field in response for chat1", modelKey)
			t.Logf("Response content: %s", responseText)
		}
	} else {
		// Fallback to checking raw text if JSON parsing fails
		if !strings.Contains(strings.ToLower(responseText), "testuser") {
			t.Errorf("[%s] Context not maintained in chat1: response doesn't contain 'TestUser'",
				modelKey)
			t.Logf("Response content: %s", responseText)
		}
	}
}

// verifyCapitalRemembered checks if the model remembered "Paris" as the capital
func verifyCapitalRemembered(t *testing.T, modelKey, responseText string) {
	var responseObj map[string]interface{}

	// Try to parse the JSON response
	if err := json.Unmarshal([]byte(responseText), &responseObj); err == nil {
		// Check if capital field contains Paris
		if capital, ok := responseObj["capital"].(string); ok {
			if !strings.Contains(strings.ToLower(capital), "paris") {
				t.Errorf("[%s] Context not maintained in chat2: capital=%s doesn't match 'Paris'",
					modelKey, capital)
			} else {
				t.Logf("[%s] Context successfully maintained in chat2", modelKey)
			}
		} else {
			t.Errorf("[%s] Missing 'capital' field in response for chat2", modelKey)
			t.Logf("Response content: %s", responseText)
		}
	} else {
		// Fallback to checking raw text if JSON parsing fails
		if !strings.Contains(strings.ToLower(responseText), "paris") {
			t.Errorf("[%s] Context not maintained in chat2: response doesn't contain 'Paris'",
				modelKey)
			t.Logf("Response content: %s", responseText)
		}
	}
}

// verifyContextIsolation checks that the sequences don't leak context between each other
func verifyContextIsolation(t *testing.T, responsesByModelAndSequence map[string]map[string][]definitions.Response) {
	for modelKey, sequenceMap := range responsesByModelAndSequence {
		if responses, ok := sequenceMap["chat2"]; ok {
			for _, response := range responses {
				responseText := response.ModelResponses[0]
				if strings.Contains(strings.ToLower(responseText), "testuser") {
					t.Errorf("[%s] Context leakage between sequences: chat2 knows about TestUser from chat1", modelKey)
					t.Logf("Response content: %s", responseText)
				}
			}
		}
	}
}
