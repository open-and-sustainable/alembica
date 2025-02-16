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

	for _, prompt := range prompts {
		chatRequest := &cohere.ChatRequest{
			Message:        prompt,
			Model:          &llm.Model,
			ConversationId: &chatID, // Maintain conversation ID
			Temperature:    &llm.Temperature,
		}

		// Make API call
		response, err := client.Chat(context.TODO(), chatRequest)
		if err != nil {
			logger.Error(fmt.Sprintf("Completion error: err:%v len(text):%v\n", err, len(response.Text)))
			return nil, fmt.Errorf("no response from Cohere: %v", err)
		}

		// Log full response JSON
		respJSON, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to marshal response: %v", err))
			return nil, err
		}
		logger.Info(fmt.Sprintf("Full Cohere response: %s\n", string(respJSON)))

		// Ensure valid response
		if len(response.Text) == 0 {
			logger.Error("No content found in response")
			return nil, fmt.Errorf("no content in response")
		}

		// Append response to answers slice
		answers = append(answers, response.Text)
	}

	return answers, nil
}
