package pricing

import (
	"encoding/json"
	"fmt"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/open-and-sustainable/alembica/llm/tokens"

	"github.com/shopspring/decimal"
)

// Declare a package-level TokenCounter variable
var tokenCounter tokens.TokenCounter = tokens.RealTokenCounter{}

// ComputeCosts processes a list of input prompts and calculates the total cost of input tokensbased on the specified 
// model and provider. 
func ComputeCosts(jsonInput string) string {
	// Parse the JSON input
	var input definitions.Input
	err := json.Unmarshal([]byte(jsonInput), &input)
	if err != nil {
		logger.Error(fmt.Println("Failed to parse JSON input:", err))
		return "0"
	}

	// Extract relevant fields
	provider := input.Models[0].Provider // Assuming at least one model exists
	model := input.Models[0].Model
	key := input.Models[0].APIKey
	prompts := make([]string, len(input.Prompts))
	for i, prompt := range input.Prompts {
		prompts[i] = prompt.PromptContent
	}
	
	// assess and report costs
	totalCost := decimal.NewFromInt(0)
	counter := 0
	for _, prompt := range prompts {
		counter++
		cost, err := assessPromptCost(prompt, provider, model, key)
		if err == nil {
			logger.Info(fmt.Println("File: ", counter, "Model: ", model, " Cost: ", cost))
			totalCost = totalCost.Add(cost)
		} else {
			logger.Error(fmt.Println("Error: ", err))
		}
	}
	return totalCost.String()
}

func assessPromptCost(prompt string, provider string, model string, key string) (decimal.Decimal, error) {
	numTokens := tokenCounter.GetNumTokensFromPrompt(prompt, provider, model, key)
	numCents := numCentsFromTokens(numTokens, model)
	return numCents, nil
}


