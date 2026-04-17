package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func queryPerplexity(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	// Create a new Perplexity client using OpenAI SDK with custom base URL
	client := openai.NewClient(
		option.WithAPIKey(llm.APIKey),
		option.WithBaseURL("https://api.perplexity.ai"),
	)

	// Initialize conversation history
	messages := []openai.ChatCompletionMessageParamUnion{}

	for i, prompt := range prompts {
		// Append user message to conversation history
		messages = append(messages, openai.UserMessage(prompt))

		// Make API call
		resp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
			Model:    openai.ChatModel(llm.Model),
			Messages: messages,
			ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
				OfJSONObject: &openai.ResponseFormatJSONObjectParam{},
			},
			Temperature: openai.Float(llm.Temperature),
		})

		if err != nil {
			logger.Error(fmt.Sprintf("Completion error: %v", err))
			return nil, fmt.Errorf("no response from Perplexity: %v", err)
		}

		// Log full response JSON
		respJSON, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			logger.Error("Failed to marshal response:", err)
			return nil, err
		}
		logger.Info(fmt.Sprintf("Full Perplexity response: %s", string(respJSON)))

		// Extract response text
		if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
			logger.Error("No content found in response")
			return nil, fmt.Errorf("no content in response")
		}

		answer := resp.Choices[0].Message.Content
		answers = append(answers, answer)

		// Append model response to conversation history
		messages = append(messages, openai.AssistantMessage(answer))

		// Call wait for all prompts except the last one
		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
