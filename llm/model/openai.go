package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	openai "github.com/sashabaranov/go-openai"
)

func queryOpenAI(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	// Create a new OpenAI client
	client := openai.NewClient(llm.APIKey)

	// Initialize conversation history
	messages := []openai.ChatCompletionMessage{}

	for i, prompt := range prompts {
		// Append user message to conversation history
		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: prompt})

		completionParams := openai.ChatCompletionRequest{
			Model:    llm.Model,
			Messages: messages,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
			Temperature: float32(llm.Temperature),
		}

		// Make API call
		resp, err := client.CreateChatCompletion(context.Background(), completionParams)
		if err != nil || len(resp.Choices) != 1 {
			logger.Error("Completion error: err:%v len(choices):%v\n", err, len(resp.Choices))
			return nil, fmt.Errorf("no response from OpenAI: %v", err)
		}

		// Log full response JSON
		respJSON, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			logger.Error("Failed to marshal response:", err)
			return nil, err
		}
		logger.Info("Full OpenAI response: %s\n", string(respJSON))

		// Extract response text
		if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
			logger.Error("No content found in response")
			return nil, fmt.Errorf("no content in response")
		}

		answer := resp.Choices[0].Message.Content
		answers = append(answers, answer)

		// Append model response to conversation history
		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: answer})


		// Call wait for all prompts except the last one
		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
