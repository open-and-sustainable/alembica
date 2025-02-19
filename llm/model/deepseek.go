package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	deepseek "github.com/cohesion-org/deepseek-go"
	"github.com/cohesion-org/deepseek-go/constants"
)

func queryDeepSeek(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	client := deepseek.NewClient(llm.APIKey)
	messages := []deepseek.ChatCompletionMessage{}

	for i, prompt := range prompts {
		messages = append(messages, deepseek.ChatCompletionMessage{Role: constants.ChatMessageRoleUser, Content: prompt})

		completionParams := &deepseek.ChatCompletionRequest{
			Model:    llm.Model,
			Messages: messages,
			ResponseFormat: &deepseek.ResponseFormat{Type: "json_object"},
			TopP:        float32(1.0),
			MaxTokens:   8192,
			Temperature: float32(llm.Temperature),
		}

		resp, err := client.CreateChatCompletion(context.Background(), completionParams)
		if err != nil {
			if apiErr, ok := err.(*deepseek.APIError); ok {
				logger.Error("API Error: HTTP %d, Code %d, Message: %s", apiErr.StatusCode, apiErr.APICode, apiErr.Message)
			} else {
				logger.Error("Unexpected error: %v", err)
			}
			return nil, fmt.Errorf("no response from deepseek: %v", err)
		}

		respJSON, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			logger.Error("Failed to marshal response: %v", err)
			return nil, err
		}
		logger.Info("Full deepseek response: %s\n", string(respJSON))

		if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
			logger.Error("No content found in response")
			return nil, fmt.Errorf("no content in response")
		}

		answer := resp.Choices[0].Message.Content
		answers = append(answers, answer)
		messages = append(messages, deepseek.ChatCompletionMessage{Role: constants.ChatMessageRoleAssistant, Content: answer})
	
		// Call wait for all prompts except the last one
		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
