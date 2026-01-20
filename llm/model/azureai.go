package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func queryAzureAI(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	if llm.BaseURL == "" {
		return nil, fmt.Errorf("missing base_url for AzureAI provider")
	}
	if llm.APIVersion == "" {
		return nil, fmt.Errorf("missing api_version for AzureAI provider")
	}

	endpoint := strings.TrimRight(llm.BaseURL, "/")
	deployment := url.PathEscape(llm.Model)
	baseURL := fmt.Sprintf("%s/openai/deployments/%s", endpoint, deployment)

	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithHeader("Api-Key", llm.APIKey),
		option.WithQuery("api-version", llm.APIVersion),
	)

	messages := []openai.ChatCompletionMessageParamUnion{}

	for i, prompt := range prompts {
		messages = append(messages, openai.UserMessage(prompt))

		resp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
			Model:    openai.ChatModel(llm.Model),
			Messages: messages,
			ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
				OfJSONObject: &openai.ResponseFormatJSONObjectParam{},
			},
			Temperature: openai.Float(llm.Temperature),
		})
		if err != nil {
			logger.Error("Completion error: %v", err)
			return nil, fmt.Errorf("no response from AzureAI: %v", err)
		}

		respJSON, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			logger.Error("Failed to marshal response: %v", err)
			return nil, err
		}
		logger.Info("Full AzureAI response: %s\n", string(respJSON))

		if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
			logger.Error("No content found in response")
			return nil, fmt.Errorf("no content in response")
		}

		answer := resp.Choices[0].Message.Content
		answers = append(answers, answer)
		messages = append(messages, openai.AssistantMessage(answer))

		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}
