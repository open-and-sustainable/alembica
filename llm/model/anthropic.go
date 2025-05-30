package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func queryAnthropic(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}
	var messages []anthropic.MessageParam

	client := anthropic.NewClient(
		option.WithAPIKey(llm.APIKey),
	)

	for i, prompt := range prompts {
		// Append new user message to the conversation history
		messages = append(messages, anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)))

		// Send the updated conversation history to the model
		message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
			Model:       anthropic.Model(llm.Model),
			MaxTokens:   4096,
			Temperature: anthropic.Float(llm.Temperature),
			Messages:    messages,
			System: []anthropic.TextBlockParam{
				{Text: "Respond with properly formatted JSON."},
			},
		})
		if err != nil {
			logger.Error("Anthropic API error: %v", err)
			return nil, fmt.Errorf("[Anthropic] API error: %v", err)
		}

		// Check if response content is valid
		if message == nil || len(message.Content) == 0 {
			logger.Error("Received nil or empty response from Anthropic API")
			return nil, fmt.Errorf("nil or empty response from Anthropic API")
		}

		// Log the response from Anthropic
		logger.Info("Anthropic response first block: %s", message.Content[0].Text)

		// Extract response text
		textBlock := extractTextBlock(message.Content)

		// Append assistant response to history
		messages = append(messages, anthropic.NewAssistantMessage(anthropic.NewTextBlock(textBlock)))

		// Extract valid JSON from response
		answer, err := extractJSONString(textBlock)
		if err != nil {
			logger.Error("Failed to extract JSON from response: %v", err)
			return nil, fmt.Errorf("no valid JSON review response from Anthropic: %v", err)
		}
		answers = append(answers, answer)

		// Call wait for all prompts except the last one
		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}

// extractTextBlock extracts the first text block from the model's response.
func extractTextBlock(content []anthropic.ContentBlockUnion) string {
	for _, block := range content {
		if block.Type == "text" {
			return block.Text
		}
	}
	return ""
}

// extractJSONString wraps extractSubstring to return properly formatted JSON.
func extractJSONString(text string) (string, error) {
	extracted, err := extractSubstring(text, "{", "}")
	if err != nil {
		return "", err
	}
	return "{\n" + extracted + "\n}", nil
}

// Function to extract substring between first occurrences of two delimiters
func extractSubstring(s, startDelim, endDelim string) (string, error) {
	// Find the index of the first occurrence of the start delimiter
	startIndex := strings.Index(s, startDelim)
	if startIndex == -1 {
		return "", fmt.Errorf("start delimiter not found")
	}

	// Adjust the start index to skip over the start delimiter
	startIndex += len(startDelim)

	// Find the index of the first occurrence of the end delimiter after the start delimiter
	endIndex := strings.Index(s[startIndex:], endDelim)
	if endIndex == -1 {
		return "", fmt.Errorf("end delimiter not found")
	}

	// Adjust the endIndex relative to the original string
	endIndex += startIndex

	// Extract the substring between the two delimiters
	return s[startIndex:endIndex], nil
}
