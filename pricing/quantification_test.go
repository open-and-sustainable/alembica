package pricing

import (
    "testing"
    "github.com/shopspring/decimal"
)

// MockTokenCounter implements the TokenCounter interface for testing
type MockTokenCounter struct {
    TokensPerPrompt map[string]int
}

func (m *MockTokenCounter) GetNumTokensFromPrompt(prompt string, provider string, model string, key string) int {
    if tokens, ok := m.TokensPerPrompt[prompt]; ok {
        return tokens
    }
    return 0 // Default to 0 if not specified
}

func TestComputeCosts(t *testing.T) {
    // Save the original tokenCounter and restore it after the test
    originalTokenCounter := tokenCounter
    defer func() { tokenCounter = originalTokenCounter }()

    // Set up the mock token counter
    mockCounter := &MockTokenCounter{
        TokensPerPrompt: map[string]int{
            "Test prompt one": 5,
            "Test prompt two": 6,
        },
    }

    // Replace the package-level tokenCounter with the mock
    tokenCounter = mockCounter

    // Create JSON input
    jsonInput := `{
        "metadata": {
            "schemaVersion": "v1",
            "timestamp": "2025-02-15T12:00:00Z"
        },
        "models": [
            {
                "provider": "OpenAI",
                "api_key": "test-api-key",
                "model": "gpt-4"
            }
        ],
        "prompts": [
            {
                "promptContent": "Test prompt one",
                "sequenceId": "1",
                "sequenceNumber": 1
            },
            {
                "promptContent": "Test prompt two",
                "sequenceId": "2",
                "sequenceNumber": 2
            }
        ]
    }`

    // Expected total cost calculation
    numTokens1 := mockCounter.TokensPerPrompt["Test prompt one"]
    numTokens2 := mockCounter.TokensPerPrompt["Test prompt two"]

    cost1 := numCentsFromTokens(numTokens1, "gpt-4")
    cost2 := numCentsFromTokens(numTokens2, "gpt-4")
    expectedTotalCost := cost1.Add(cost2)

    // Call ComputeCosts
    totalCostStr := ComputeCosts(jsonInput)

    totalCost, err := decimal.NewFromString(totalCostStr)
    if err != nil {
        t.Fatalf("Failed to parse total cost: %v", err)
    }

    if !totalCost.Equal(expectedTotalCost) {
        t.Errorf("Total cost mismatch. Expected: %s, Got: %s", expectedTotalCost.String(), totalCost.String())
    }
}
