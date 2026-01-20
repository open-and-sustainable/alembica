package check

import (
	"github.com/open-and-sustainable/alembica/llm/tokens"
	"github.com/open-and-sustainable/alembica/utils/logger"

	"github.com/openai/openai-go/v3/shared"
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
	case "Perplexity":
		modelFunc = getPerplexityModel
	default:
		logger.Error("Unsupported LLM provider: %s", providerName)
		return ""
	}
	return modelFunc(prompt, modelName, key)
}

func getOpenAIModel(prompt string, modelName string, key string) string {
	model := string(shared.ChatModelGPT4oMini)
	switch modelName {
	case "": // cost optimization
		// old code before GPT 4 Omni mini model availability -- now the only solution to minimize the cost
		/*numTokens := numTokensFromMessages([]openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}, model)
		if numTokens > 16385 {
			model = string(shared.ChatModelGPT4o)
		}*/
	case "gpt-3.5-turbo":
		model = string(shared.ChatModelGPT3_5Turbo)
	case "gpt-4-turbo":
		model = string(shared.ChatModelGPT4Turbo)
	case "gpt-4o":
		model = string(shared.ChatModelGPT4o)
	case "gpt-4o-mini":
		model = string(shared.ChatModelGPT4oMini)
	case "o1":
		model = string(shared.ChatModelO1)
	case "o1-mini":
		model = string(shared.ChatModelO1Mini)
	case "o3":
		model = string(shared.ChatModelO3)
	case "o3-mini":
		model = string(shared.ChatModelO3Mini)
	case "o4-mini":
		model = string(shared.ChatModelO4Mini)
	case "gpt-4.1":
		model = string(shared.ChatModelGPT4_1)
	case "gpt-4.1-mini":
		model = string(shared.ChatModelGPT4_1Mini)
	case "gpt-4.1-nano":
		model = string(shared.ChatModelGPT4_1Nano)
	case "gpt-5":
		model = string(shared.ChatModelGPT5)
	case "gpt-5.1":
		model = string(shared.ChatModelGPT5_1)
	case "gpt-5.2":
		model = string(shared.ChatModelGPT5_2)
	case "gpt-5-mini":
		model = string(shared.ChatModelGPT5Mini)
	case "gpt-5-nano":
		model = string(shared.ChatModelGPT5Nano)
	default:
		logger.Error("Unsupported model: %s", modelName)
		return ""
	}
	return model
}

func getGoogleAIModel(prompt string, modelName string, key string) string {
	model := "gemini-2.5-flash-lite"
	switch modelName {
	case "": // cost optimization, default to most cost-effective model
		counter := tokens.RealTokenCounter{}
		numTokens := counter.GetNumTokensFromPrompt(prompt, "GoogleAI", modelName, key)
		if numTokens <= 1048576 {
			model = "gemini-2.5-flash-lite"
		}
	case "gemini-1.0-pro": // deprecated from Feb 15 2025
		logger.Error("Unsupported model: %s", modelName)
		return ""
	case "gemini-3-pro-preview":
		model = modelName
	case "gemini-3-flash-preview":
		model = modelName
	case "gemini-2.5-pro":
		model = modelName
	case "gemini-2.5-flash":
		model = modelName
	case "gemini-2.5-flash-lite":
		model = modelName
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
	case "command-a-reasoning-08-2025":
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
		model = "claude-haiku-4-5-20251015"
	case "claude-4-5-opus":
		model = "claude-opus-4-5-20260320"
	case "claude-4-5-sonnet":
		model = "claude-sonnet-4-5-20260320"
	case "claude-4-5-haiku":
		model = "claude-haiku-4-5-20251015"
	case "claude-4-0-opus":
		model = "claude-opus-4-0-20251101"
	case "claude-4-0-sonnet":
		model = "claude-sonnet-4-0-20250514"
	case "claude-3-7-sonnet":
		model = "claude-3-7-sonnet-20250219"
	case "claude-3-5-sonnet":
		model = "claude-3-5-sonnet-20241022"
	case "claude-3-5-haiku":
		model = "claude-3-5-haiku-20241022"
	case "claude-3-opus":
		model = "claude-3-opus-20240229"
	case "claude-3-haiku":
		model = "claude-3-haiku-20240307"
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

func getPerplexityModel(prompt string, modelName string, key string) string {
	model := "sonar"
	switch modelName {
	case "":
		// cost optimization: use sonar as the cheapest
		model = "sonar"
	case "sonar":
		model = modelName
	case "sonar-pro":
		model = modelName
	case "sonar-reasoning-pro":
		model = modelName
	case "sonar-deep-research":
		model = modelName
	default:
		logger.Error("Unsupported model: %s", modelName)
		return ""
	}
	return model
}
