package tokens

import (
	"context"
	"github.com/open-and-sustainable/alembica/utils/logger"

	genai "github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)

func numTokensFromPromptGoogleAI(prompt string, modelName string, key string) (numTokens int) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(key))
	if err != nil {
		logger.Error("[GoogleAI] Failed to create client: %v", err)
		return 0
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	tokResp, err := model.CountTokens(ctx, genai.Text(prompt))
	if err != nil {
		logger.Error("[GoogleAI] Failed to count tokens: %v", err)
		return 0 // ✅ Do NOT stop execution
	}

	return int(tokResp.TotalTokens)
}
