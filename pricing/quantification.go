package pricing

import (
	"encoding/json"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/open-and-sustainable/alembica/llm/tokens"
	"github.com/open-and-sustainable/alembica/validation"

	"github.com/shopspring/decimal"
)

// Declare a package-level TokenCounter variable
var tokenCounter tokens.TokenCounter = tokens.RealTokenCounter{}

// ComputeCosts processes a list of input prompts and calculates the total cost of input tokensbased on the specified 
// model and provider. 
func ComputeCosts(jsonInput string, version ...string) (string, error) {
	// Set default version if not provided
	v := "v1"
	if len(version) > 0 {
		v = version[0]
	}

	// Validate input JSON
	if err := validation.ValidateInput(jsonInput, v); err != nil {
		logger.Error("Invalid input JSON:", err)
		return "", err
	}

	// Parse the JSON input
	var input definitions.Input
	err := json.Unmarshal([]byte(jsonInput), &input)
	if err != nil {
		logger.Error("Failed to parse JSON input:", err)
		return "", err
	}

	// Initialize cost tracking structure
	costOutput := definitions.CostOutput{
		Metadata: definitions.CostMetadata{
			SchemaVersion: v,
			Currency:      "USD",
		},
		Costs: []definitions.Cost{},
	}

	// Compute costs by sequence
	sequenceCostMap := make(map[string]decimal.Decimal)
	for _, prompt := range input.Prompts {
		sequenceTotalCost := decimal.NewFromInt(0)
		for _, model := range input.Models {
			cost, err := assessPromptCost(prompt.PromptContent, model.Provider, model.Model, model.APIKey)
			if err != nil {
				logger.Error("Error processing cost for Sequence ID:", prompt.SequenceID, "Model:", model.Model, "Error:", err)
				continue // Skip this model and move to the next one
			}

			logger.Info("Sequence ID:", prompt.SequenceID, "Provider:", model.Provider, "Model:", model.Model, "Cost:", cost)
			sequenceTotalCost = sequenceTotalCost.Add(cost)

			// Store cost details
			costOutput.Costs = append(costOutput.Costs, definitions.Cost{
				SequenceID: prompt.SequenceID,
				Provider:   model.Provider,
				Model:      model.Model,
				Cost:       cost.InexactFloat64(),
			})
		}

		// Ensure per-sequence cost accumulates correctly
		if existingCost, ok := sequenceCostMap[prompt.SequenceID]; ok {
			sequenceCostMap[prompt.SequenceID] = existingCost.Add(sequenceTotalCost)
		} else {
			sequenceCostMap[prompt.SequenceID] = sequenceTotalCost
		}
	}

	for seqID, total := range sequenceCostMap {
		costOutput.Costs = append(costOutput.Costs, definitions.Cost{
			SequenceID: seqID,
			Provider:   "TOTAL",
			Model:      "TOTAL",
			Cost: float64(total.InexactFloat64()),
		})
	}

	// Convert costOutput to JSON for returning
	costOutputJSON, err := json.Marshal(costOutput)
	if err != nil {
		logger.Error("Failed to marshal cost output JSON:", err)
		return "", err
	}

	// Validate output JSON
	if err := validation.ValidateCost(string(costOutputJSON), v); err != nil {
		logger.Error("Invalid output JSON:", err)
		return "", err
	}

	return string(costOutputJSON), nil
}

func assessPromptCost(prompt string, provider string, model string, key string) (decimal.Decimal, error) {
	numTokens := tokenCounter.GetNumTokensFromPrompt(prompt, provider, model, key)
	numCents := numCentsFromTokens(numTokens, model)
	return numCents, nil
}


