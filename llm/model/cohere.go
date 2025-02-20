package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
	uuid "github.com/google/uuid"
)

func queryCohere(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}
	chatID := uuid.New().String()

	// Create a new Cohere client
	client := cohereclient.NewClient(cohereclient.WithToken(llm.APIKey))

	for i, prompt := range prompts {
		chatRequest := &cohere.ChatRequest{
			Message:        prompt,
			Model:          &llm.Model,
			ConversationId: &chatID, // Maintain conversation ID
			Temperature:    &llm.Temperature,
		}

		// Log request for debugging
		reqJSON, _ := json.MarshalIndent(chatRequest, "", "  ")
		logger.Info("Sending Cohere request: %s\n", string(reqJSON))

		// Make API call
		response, err := client.Chat(context.TODO(), chatRequest)
		if err != nil {
			logger.Error("Cohere API error: %v", err)
			return nil, fmt.Errorf("[Cohere] API error: %v", err)
		}

		// Check if response is nil before accessing its fields
		if response == nil {
			logger.Error("Received nil response from Cohere API")
			return nil, fmt.Errorf("nil response from Cohere API")
		}

		// Log full response JSON
		respJSON, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			logger.Error("Failed to marshal response: %v", err)
			return nil, err
		}
		logger.Info("Full Cohere response: %s\n", string(respJSON))

		// Ensure valid response
		if len(response.Text) == 0 {
			logger.Error("No content found in response")
			return nil, fmt.Errorf("no content in response from Cohere")
		}

		// Append response to answers slice
		answers = append(answers, response.Text)

		// Call wait for all prompts except the last one
		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
