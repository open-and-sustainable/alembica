package pricing

import (
	"encoding/json"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/open-and-sustainable/alembica/llm/tokens"
	"github.com/open-and-sustainable/alembica/validation"

	"github.com/shopspring/decimal"
)

// Global TokenCounter for calculating token usage across models.
var tokenCounter tokens.TokenCounter = tokens.RealTokenCounter{}

// ComputeCosts calculates the cost of processing input prompts based on the specified models and providers.
//
// Parameters:
//   - jsonInput: A JSON string containing the input data.
//   - version: (Optional) The schema version to validate against. Defaults to "v1" if not provided.
//
// Returns:
//   - A JSON string containing computed cost details.
//   - An error if input validation, cost computation, or output validation fails.
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

	// Compute costs per sequence
	sequenceCostMap := make(map[string]decimal.Decimal)
	for _, prompt := range input.Prompts {
		sequenceTotalCost := decimal.NewFromInt(0)
		for _, model := range input.Models {
			cost, err := assessPromptCost(prompt.PromptContent, model.Provider, model.Model, model.APIKey)
			if err != nil {
				logger.Error("Error processing cost for Sequence ID:", prompt.SequenceID, "Model:", model.Model, "Error:", err)
				continue
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

		// Accumulate total per sequence
		if existingCost, ok := sequenceCostMap[prompt.SequenceID]; ok {
			sequenceCostMap[prompt.SequenceID] = existingCost.Add(sequenceTotalCost)
		} else {
			sequenceCostMap[prompt.SequenceID] = sequenceTotalCost
		}
	}

	// Append total cost per sequence
	for seqID, total := range sequenceCostMap {
		costOutput.Costs = append(costOutput.Costs, definitions.Cost{
			SequenceID: seqID,
			Provider:   "TOTAL",
			Model:      "TOTAL",
			Cost:       float64(total.InexactFloat64()),
		})
	}

	// Convert costOutput to JSON
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

// assessPromptCost calculates the cost of processing a prompt based on token usage.
//
// Parameters:
//   - prompt: The text prompt whose cost is being assessed.
//   - provider: The LLM provider handling the prompt.
//   - model: The model processing the prompt.
//   - key: The API key for accessing the model.
//
// Returns:
//   - The computed cost as a decimal.Decimal value.
//   - An error if token counting fails.
func assessPromptCost(prompt string, provider string, model string, key string) (decimal.Decimal, error) {
	numTokens := tokenCounter.GetNumTokensFromPrompt(prompt, provider, model, key)
	numCents := numCentsFromTokens(numTokens, model)
	return numCents, nil
}
