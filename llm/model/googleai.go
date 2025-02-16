package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	genai "github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)

func queryGoogleAI(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	// Create a new context for API calls
	ctx := context.Background()

	// Create a new Google Generative AI client using the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(llm.APIKey))
	if err != nil {
		logger.Error(fmt.Sprintf("[GoogleAI] Failed to create client: %v", err))
		return nil, err
	}
	defer client.Close()

	// Log selected model
	logger.Info(fmt.Sprintf("[GoogleAI] Using model: %s", llm.Model))

	// Select and configure the generative model
	model := client.GenerativeModel(llm.Model)
	model.SetTemperature(float32(llm.Temperature))
	model.SetCandidateCount(1)
	model.ResponseMIMEType = "application/json"

	// Start a new chat session
	cs := model.StartChat()

	// Loop over prompts while maintaining chat history
	for i, prompt := range prompts {
		logger.Info(fmt.Sprintf("[GoogleAI] Sending prompt #%d: %s", i+1, prompt))

		// Append user message to conversation history
		cs.History = append(cs.History, &genai.Content{
			Parts: []genai.Part{genai.Text(prompt)},
			Role:  "user",
		})

		// Send message to model
		resp, err := cs.SendMessage(ctx, genai.Text(prompt))
		if err != nil {
			logger.Error(fmt.Sprintf("[GoogleAI] Error on prompt #%d: %v", i+1, err))
			return nil, fmt.Errorf("the Google AI response error: %v", err)
		}
		
		// Ensure response contains candidates
		if len(resp.Candidates) == 0 {
			logger.Error(fmt.Sprintf("[GoogleAI] No candidates received for prompt #%d", i+1))
			return nil, fmt.Errorf("no candidates returned from Google AI")
		}

		// Log full response JSON
		respJSON, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			logger.Error(fmt.Sprintf("[GoogleAI] Failed to marshal response for prompt #%d: %v", i+1, err))
			return nil, err
		}
		logger.Info(fmt.Sprintf("[GoogleAI] Full response for prompt #%d: %s", i+1, string(respJSON)))

		// Extract content from first candidate
		content := resp.Candidates[0].Content
		if content == nil || len(content.Parts) == 0 {
			logger.Error(fmt.Sprintf("[GoogleAI] No content parts in response for prompt #%d", i+1))
			return nil, fmt.Errorf("no content in response")
		}

		// Iterate over parts to extract text
		var resultText string
		for _, part := range content.Parts {
			switch v := part.(type) {
			case genai.Text:
				resultText += string(v)
			default:
				logger.Error(fmt.Sprintf("[GoogleAI] Unhandled response part type: %T", part))
			}
		}

		// Validate extracted text
		if resultText == "" {
			logger.Error(fmt.Sprintf("[GoogleAI] No text content extracted for prompt #%d", i+1))
			return nil, fmt.Errorf("empty response from Google AI")
		}

		// Append response to answers and history
		answers = append(answers, resultText)
		cs.History = append(cs.History, &genai.Content{
			Parts: []genai.Part{genai.Text(resultText)},
			Role:  "assistant",
		})

		logger.Info(fmt.Sprintf("[GoogleAI] Processed response for prompt #%d: %s", i+1, resultText))
	}

	return answers, nil
}
