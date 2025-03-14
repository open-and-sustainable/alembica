package extraction

import (
    "encoding/json"
    "fmt"
    "sort"

    "github.com/open-and-sustainable/alembica/definitions"
    "github.com/open-and-sustainable/alembica/utils/logger"
    "github.com/open-and-sustainable/alembica/validation"
    "github.com/open-and-sustainable/alembica/llm/model"
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

    for _, prompts := range promptsBySequence {
        sort.SliceStable(prompts, func(i, j int) bool {
            return prompts[i].SequenceNumber < prompts[j].SequenceNumber
        })
    }

    queryService := model.DefaultQueryService{}

    for _, modelInstance := range inputData.Models {
        for _, sequenceID := range sequenceIDs {
            prompts := promptsBySequence[sequenceID]

            for _, p := range prompts {
                // Query model with each individual prompt
                responses, err := queryService.QueryLLM([]string{p.PromptContent}, modelInstance)
                if err != nil {
                    logger.Error("error querying LLM: %v", err)
                    continue
                }

                outputResponses := definitions.Response{
                    Provider:       modelInstance.Provider,
                    Model:          modelInstance.Model,
                    SequenceID:     sequenceID,
                    SequenceNumber: p.SequenceNumber,
                    ModelResponses: responses,
                }

                outputData.Responses = append(outputData.Responses, outputResponses)
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
