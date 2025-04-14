package extraction

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/llm/model"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/open-and-sustainable/alembica/validation"
)

// Extract processes input JSON, queries LLMs, and returns structured responses.
//
// Parameters:
//   - inputJSON: JSON string containing metadata, models, and prompts.
//
// Returns:
//   - A JSON string with responses from the models, or an error if processing fails.
func Extract(inputJSON string) (string, error) {
	var inputData definitions.Input
	err := json.Unmarshal([]byte(inputJSON), &inputData)
	if err != nil {
		logger.Error("error parsing input JSON: %v", err)
		return "", err
	}

	outputData := definitions.Output{
		Metadata: definitions.OutputMetadata{
			SchemaVersion: inputData.Metadata.SchemaVersion,
		},
		Responses: []definitions.Response{},
	}

	promptsBySequence := make(map[string][]definitions.Prompt)
	sequenceIDs := []string{}

	for _, prompt := range inputData.Prompts {
		if _, exists := promptsBySequence[prompt.SequenceID]; !exists {
			sequenceIDs = append(sequenceIDs, prompt.SequenceID)
		}
		promptsBySequence[prompt.SequenceID] = append(promptsBySequence[prompt.SequenceID], prompt)
	}

	// Sort prompts within each sequence by sequence number
	for seqID := range promptsBySequence {
		sort.SliceStable(promptsBySequence[seqID], func(i, j int) bool {
			return promptsBySequence[seqID][i].SequenceNumber < promptsBySequence[seqID][j].SequenceNumber
		})
	}

	queryService := model.DefaultQueryService{}

	for _, modelInstance := range inputData.Models {
		for _, sequenceID := range sequenceIDs {
			prompts := promptsBySequence[sequenceID]

			// Extract all prompt contents in correct sequence order
			var promptContents []string
			for _, p := range prompts {
				promptContents = append(promptContents, p.PromptContent)
			}

			// Query the model with all prompts in the sequence at once
			responses, err := queryService.QueryLLM(promptContents, modelInstance)
			if err != nil {
				logger.Error("error querying LLM: %v", err)
				continue
			}

			// Process responses (they should be in the same order as prompts)
			for i, p := range prompts {
				if i < len(responses) {
					outputResponse := definitions.Response{
						Provider:       modelInstance.Provider,
						Model:          modelInstance.Model,
						SequenceID:     sequenceID,
						SequenceNumber: p.SequenceNumber,
						ModelResponses: []string{responses[i]}, // Ensure this matches your structure
					}

					outputData.Responses = append(outputData.Responses, outputResponse)
				}
			}
		}
	}

	outputJSON, err := json.Marshal(outputData)
	if err != nil {
		logger.Error("error generating output JSON: %v", err)
		return "", fmt.Errorf("error generating output JSON: %v", err)
	}

	if err := validation.ValidateOutput(string(outputJSON), inputData.Metadata.SchemaVersion); err != nil {
		logger.Error("error validating output JSON: %v", err)
		return "", err
	}

	return string(outputJSON), nil
}
