package model

import (
	"context"
	"fmt"

	"github.com/open-and-sustainable/alembica/definitions"
	"github.com/open-and-sustainable/alembica/utils/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

func queryAWSBedrock(prompts []string, llm definitions.Model) ([]string, error) {
	answers := []string{}

	if llm.Region == "" {
		return nil, fmt.Errorf("missing region for AWSBedrock provider")
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(llm.Region))
	if err != nil {
		logger.Error(fmt.Sprintf("AWS config error: %v", err))
		return nil, err
	}

	client := bedrockruntime.NewFromConfig(cfg)
	messages := []types.Message{}

	for i, prompt := range prompts {
		messages = append(messages, types.Message{
			Role: types.ConversationRoleUser,
			Content: []types.ContentBlock{
				&types.ContentBlockMemberText{Value: prompt},
			},
		})

		resp, err := client.Converse(ctx, &bedrockruntime.ConverseInput{
			ModelId:  aws.String(llm.Model),
			Messages: messages,
			InferenceConfig: &types.InferenceConfiguration{
				Temperature: aws.Float32(float32(llm.Temperature)),
			},
		})
		if err != nil {
			logger.Error(fmt.Sprintf("Bedrock API error: %v", err))
			return nil, fmt.Errorf("no response from AWS Bedrock: %v", err)
		}

		outputMessage, ok := resp.Output.(*types.ConverseOutputMemberMessage)
		if !ok || outputMessage.Value.Content == nil {
			return nil, fmt.Errorf("empty response from AWS Bedrock")
		}

		answer := extractBedrockText(outputMessage.Value.Content)
		if answer == "" {
			return nil, fmt.Errorf("no content in response")
		}

		answers = append(answers, answer)
		messages = append(messages, types.Message{
			Role: types.ConversationRoleAssistant,
			Content: []types.ContentBlock{
				&types.ContentBlockMemberText{Value: answer},
			},
		})

		if i < len(prompts)-1 {
			Wait(prompt, llm)
		}
	}

	return answers, nil
}

func extractBedrockText(blocks []types.ContentBlock) string {
	for _, block := range blocks {
		switch v := block.(type) {
		case *types.ContentBlockMemberText:
			return v.Value
		}
	}
	return ""
}
