package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	genai "cloud.google.com/go/vertexai/genai"
)

func queryVertexAI(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	if llm.ProjectID == "" || llm.Location == "" {
		return nil, fmt.Errorf("missing project_id or location for VertexAI provider")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, llm.ProjectID, llm.Location)
	if err != nil {
		logger.Error(fmt.Sprintf("[VertexAI] Failed to create client: %v", err))
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel(llm.Model)
	model.SetTemperature(float32(llm.Temperature))
	model.SetCandidateCount(1)
	model.ResponseMIMEType = "application/json"

	cs := model.StartChat()

	for i, prompt := range prompts {
		logger.Info(fmt.Sprintf("[VertexAI] Sending prompt #%d: %s", i+1, prompt))

		cs.History = append(cs.History, &genai.Content{
			Parts: []genai.Part{genai.Text(prompt)},
			Role:  "user",
		})

		resp, err := cs.SendMessage(ctx, genai.Text(prompt))
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

		var resultText string
		for _, part := range content.Parts {
			switch v := part.(type) {
			case genai.Text:
				resultText += string(v)
			default:
				logger.Error(fmt.Sprintf("[VertexAI] Unhandled response part type: %T", part))
			}
		}

		if resultText == "" {
			logger.Error(fmt.Sprintf("[VertexAI] No text content extracted for prompt #%d", i+1))
			return nil, fmt.Errorf("empty response from Vertex AI")
		}

		answers = append(answers, resultText)
		cs.History = append(cs.History, &genai.Content{
			Parts: []genai.Part{genai.Text(resultText)},
			Role:  "model",
		})

		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
