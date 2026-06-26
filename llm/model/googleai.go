package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	"google.golang.org/genai"
)

func queryGoogleAI(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	// Create a new context for API calls
	ctx := context.Background()

	// Create a new Google Gemini API client using the API key
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  llm.APIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("[GoogleAI] Failed to create client: %v", err))
		return nil, err
	}

	// Log selected model
	logger.Info(fmt.Sprintf("[GoogleAI] Using model: %s", llm.Model))

	// Configure the generative model
	config := &genai.GenerateContentConfig{
		Temperature:      genai.Ptr(float32(llm.Temperature)),
		CandidateCount:   1,
		ResponseMIMEType: "application/json",
	}

	// Start a new chat session; history is maintained automatically by SendMessage
	cs, err := client.Chats.Create(ctx, llm.Model, config, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("[GoogleAI] Failed to create chat: %v", err))
		return nil, err
	}

	// Loop over prompts while maintaining chat history
	for i, prompt := range prompts {
		logger.Info(fmt.Sprintf("[GoogleAI] Sending prompt #%d: %s", i+1, prompt))

		// Send message to model
		resp, err := cs.SendMessage(ctx, genai.Part{Text: prompt})
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

		// Concatenate all text parts of the response
		resultText := resp.Text()

		// Validate extracted text
		if resultText == "" {
			logger.Error(fmt.Sprintf("[GoogleAI] No text content extracted for prompt #%d", i+1))
			return nil, fmt.Errorf("empty response from Google AI")
		}

		// Append response to answers
		answers = append(answers, resultText)

		logger.Info(fmt.Sprintf("[GoogleAI] Processed response for prompt #%d: %s", i+1, resultText))

		// Call wait for all prompts except the last one
		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
