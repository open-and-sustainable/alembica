package tokens

import (
	"context"
	"fmt"
	"github.com/open-and-sustainable/alembica/utils/logger"

	"google.golang.org/genai"
)

func numTokensFromPromptGoogleAI(prompt string, modelName string, key string) (numTokens int) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("[GoogleAI] Failed to create client: %v", err))
		return 0
	}

	tokResp, err := client.Models.CountTokens(ctx, modelName, genai.Text(prompt), nil)
	if err != nil {
		logger.Error(fmt.Sprintf("[GoogleAI] Failed to count tokens: %v", err))
		return 0 // ✅ Do NOT stop execution
	}

	return int(tokResp.TotalTokens)
}
