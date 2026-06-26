package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	"google.golang.org/genai"
)

func queryVertexAI(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	if llm.ProjectID == "" || llm.Location == "" {
		return nil, fmt.Errorf("missing project_id or location for VertexAI provider")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Project:  llm.ProjectID,
		Location: llm.Location,
		Backend:  genai.BackendVertexAI,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("[VertexAI] Failed to create client: %v", err))
		return nil, err
	}

	config := &genai.GenerateContentConfig{
		Temperature:      genai.Ptr(float32(llm.Temperature)),
		CandidateCount:   1,
		ResponseMIMEType: "application/json",
	}

	// Start a new chat session; history is maintained automatically by SendMessage
	cs, err := client.Chats.Create(ctx, llm.Model, config, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("[VertexAI] Failed to create chat: %v", err))
		return nil, err
	}

	for i, prompt := range prompts {
		logger.Info(fmt.Sprintf("[VertexAI] Sending prompt #%d: %s", i+1, prompt))

		resp, err := cs.SendMessage(ctx, genai.Part{Text: prompt})
		if err != nil {
			logger.Error(fmt.Sprintf("[VertexAI] Error on prompt #%d: %v", i+1, err))
			return nil, fmt.Errorf("the Vertex AI response error: %v", err)
		}

		if len(resp.Candidates) == 0 {
			logger.Error(fmt.Sprintf("[VertexAI] No candidates received for prompt #%d", i+1))
			return nil, fmt.Errorf("no candidates returned from Vertex AI")
		}

		respJSON, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			logger.Error(fmt.Sprintf("[VertexAI] Failed to marshal response for prompt #%d: %v", i+1, err))
			return nil, err
		}
		logger.Info(fmt.Sprintf("[VertexAI] Full response for prompt #%d: %s", i+1, string(respJSON)))

		content := resp.Candidates[0].Content
		if content == nil || len(content.Parts) == 0 {
			logger.Error(fmt.Sprintf("[VertexAI] No content parts in response for prompt #%d", i+1))
			return nil, fmt.Errorf("no content in response")
		}

		resultText := resp.Text()
		if resultText == "" {
			logger.Error(fmt.Sprintf("[VertexAI] No text content extracted for prompt #%d", i+1))
			return nil, fmt.Errorf("empty response from Vertex AI")
		}

		answers = append(answers, resultText)

		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
