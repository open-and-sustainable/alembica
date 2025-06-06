package check

import (

	"github.com/open-and-sustainable/alembica/llm/tokens"
	"github.com/open-and-sustainable/alembica/utils/logger"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	openai "github.com/sashabaranov/go-openai"
)

// GetModel selects the appropriate model for the given provider based on user input and internal logic.
//
// Parameters:
//   - prompt: The input prompt provided by the user.
//   - providerName: The name of the AI provider (e.g., "OpenAI", "Cohere").
//   - modelName: The name of the specific model, if any. If empty, cost optimization is attempted.
//   - key: A string representing a key for the provider's service.
//
// Returns:
//   - A string representing the selected model name. An empty string is returned if the model is unsupported.
//
// Example:
//
//	> selectedModel := GetModel("some prompt", "OpenAI", "gpt-4-turbo", "api-key")
//	> if selectedModel == "" {
//	>     log.Println("No supported model selected")
//	> }
func GetModel(prompt string, providerName string, modelName string, key string) string {
	var modelFunc func(string, string, string) string
	switch providerName {
	case "OpenAI":
		modelFunc = getOpenAIModel
	case "GoogleAI":
		modelFunc = getGoogleAIModel
	case "Cohere":
		modelFunc = getCohereModel
	case "Anthropic":
		modelFunc = getAnthropicModel
	case "DeepSeek":
		modelFunc = getDeepSeekModel
	default:
		logger.Error("Unsupported LLM provider: %s", providerName)
		return ""
	}
	return modelFunc(prompt, modelName, key)
}

func getOpenAIModel(prompt string, modelName string, key string) string {
	model := openai.GPT4oMini
	switch modelName {
	case "": // cost optimization
		// old code before GPT 4 Omni mini model availability -- now the only solution to minimize the cost
		/*numTokens := numTokensFromMessages([]openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}, model)
		if numTokens > 16385 {
			model = openai.GPT4o
		}*/
	case "gpt-3.5-turbo":
		model = openai.GPT3Dot5Turbo
	case "gpt-4-turbo":
		model = openai.GPT4Turbo
	case "gpt-4o":
		model = openai.GPT4o
	case "gpt-4o-mini":
		model = openai.GPT4oMini
	case "o1":
		model = openai.O1
	case "o1-mini":
		model = openai.O1Mini
	case "o3":
		model = openai.O3
	case "o3-mini":
		model = openai.O3Mini
	case "o4-mini":
		model = openai.O4Mini
	case "gpt-4.1":
		model = openai.GPT4Dot1
	case "gpt-4.1-mini":
		model = openai.GPT4Dot1Mini
	case "gpt-4.1-nano":
		model = openai.GPT4Dot1Nano
	default:
		logger.Error("Unsupported model: %s", modelName)
		return ""
	}
	return model
}

func getGoogleAIModel(prompt string, modelName string, key string) string {
	model := "gemini-1.5-pro"
	switch modelName {
	case "": // cost optimization, input token limit values: gemini-1.5-flash 1048576, gemini-1.5-pro 2097152
		counter := tokens.RealTokenCounter{}
		numTokens := counter.GetNumTokensFromPrompt(prompt, "GoogleAI", modelName, key)
		if numTokens <= 1048576 {
			model = "gemini-2.0-flash-lite"
		}
	case "gemini-1.0-pro": // deprecated from Feb 15 2025
		logger.Error("Unsupported model: %s", modelName)
		return ""
	case "gemini-1.5-flash":
		model = modelName
	case "gemini-1.5-pro":
		model = modelName
	case "gemini-2.0-flash-lite":
		model = modelName
	case "gemini-2.0-flash":
		model = modelName
	default:
		logger.Error("Unsupported model: %s", modelName)
		return ""
	}
	return model
}

func getCohereModel(prompt string, modelName string, key string) string {
	model := "command-r7b-12-2024"
	switch modelName {
	case "":
		// cost optimization, command-r7b is currently the cheapest and with the most input tokens allowed
	case "command": // leave the model selected by the user, but check if supported
		model = modelName
	case "command-light":
		model = modelName
	case "command-r":
		model = modelName
	case "command-r-08-2024":
		model = modelName
	case "command-r-plus":
		model = modelName
	case "command-r7b-12-2024":
		model = modelName
	case "command-a-03-2025":
		model = modelName
	default:
		logger.Error("Unsupported model: %s", modelName)
		return ""
	}
	return model
}

func getAnthropicModel(prompt string, modelName string, key string) string {
	var model string
	switch modelName {
	case "": // cost optimization
		// all models have the same context window size, hence leave to haiku as the cheapest
		model = string(anthropic.ModelClaude_3_Haiku_20240307)
	case "claude-4-0-opus":
		model = string(anthropic.ModelClaudeOpus4_0)
	case "claude-4-0-sonnet":
		model = string(anthropic.ModelClaudeSonnet4_0)
	case "claude-3-7-sonnet":
		model = string(anthropic.ModelClaude3_7SonnetLatest)
	case "claude-3-5-sonnet":
		model = string(anthropic.ModelClaude3_5SonnetLatest)
	case "claude-3-5-haiku":
		model = string(anthropic.ModelClaude3_5HaikuLatest)
	case "claude-3-opus":
		model = string(anthropic.ModelClaude3OpusLatest)
	case "claude-3-haiku":
		model = string(anthropic.ModelClaude_3_Haiku_20240307)
	default:
		logger.Error("Unsupported model: %s", modelName)
		return ""
	}
	return model
}

func getDeepSeekModel(prompt string, modelName string, key string) string {
	model := "deepseek-chat"
	switch modelName {
	case "":
		// cost optimization: use chat that is cheapest
		model = "deepseek-chat"
	case "deepseek-chat": // leave the model selected by the user, but check if supported
		model = modelName
	case "deepseek-reasoner": // leave the model selected by the user, but check if supported
		model = modelName
	default:
		logger.Error("Unsupported model: %s", modelName)
		return ""
	}
	return model
}
