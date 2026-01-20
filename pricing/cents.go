package pricing

import (
	"github.com/open-and-sustainable/alembica/utils/logger"

	"github.com/shopspring/decimal"

	"github.com/openai/openai-go/v3/shared"
)

// modelRates maps LLM model identifiers to their corresponding pricing rates.
//
// The rates are stored as decimal.Decimal values representing cost per token,
// calculated as USD dollars per million tokens and then divided by 1,000,000.
// This ensures precise cost calculations even for very small per-token amounts.
//
// Supported providers and models include:
// - OpenAI: Various GPT models (O1, O3, O4, GPT-4o, GPT-3.5, etc.)
// - Google AI: Gemini models (1.5 Pro, 1.5 Flash, 2.0 Flash, etc.)
// - Cohere: Command models (Command-R, Command-R+, Command-R7B, etc.)
// - Anthropic: Claude models (3.5 Sonnet, 3 Haiku, 3 Opus, etc.)
// - DeepSeek: DeepSeek Chat and Reasoner models
// - Perplexity: Sonar models (Sonar, Sonar Pro, Sonar Reasoning Pro, etc.)
//
// Note: Some models like Google's Gemini have tiered pricing that depends
// on token count thresholds, which is handled in the numCentsFromTokens function.
var modelRates = map[string]decimal.Decimal{ // dollar prices per input M token
	string(shared.ChatModelO4Mini):      decimal.NewFromFloat(1.10).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelO1):          decimal.NewFromFloat(15).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelO1Mini):      decimal.NewFromFloat(1.10).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelO3):          decimal.NewFromFloat(2.00).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelO3Mini):      decimal.NewFromFloat(1.10).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT5):        decimal.NewFromFloat(1.25).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT5_1):      decimal.NewFromFloat(1.25).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT5_2):      decimal.NewFromFloat(1.75).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT5Mini):    decimal.NewFromFloat(0.25).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT5Nano):    decimal.NewFromFloat(0.05).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT4_1):      decimal.NewFromFloat(2.00).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT4_1Mini):  decimal.NewFromFloat(0.40).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT4_1Nano):  decimal.NewFromFloat(0.10).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT4oMini):   decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT4o):       decimal.NewFromFloat(2.50).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT4Turbo):   decimal.NewFromFloat(10).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT4):        decimal.NewFromFloat(30).Div(decimal.NewFromInt(1000000)),
	string(shared.ChatModelGPT3_5Turbo): decimal.NewFromFloat(0.5).Div(decimal.NewFromInt(1000000)),
	"gemini-3-pro-preview":              decimal.NewFromFloat(2.00).Div(decimal.NewFromInt(1000000)), // $2.00 for prompts <= 200k tokens, $4.00 for > 200k
	"gemini-3-flash-preview":            decimal.NewFromFloat(0.50).Div(decimal.NewFromInt(1000000)),
	"gemini-2.5-pro":                    decimal.NewFromFloat(1.25).Div(decimal.NewFromInt(1000000)), // $1.25 for prompts <= 200k tokens, $2.50 for > 200k
	"gemini-2.5-flash":                  decimal.NewFromFloat(0.30).Div(decimal.NewFromInt(1000000)),
	"gemini-2.5-flash-lite":             decimal.NewFromFloat(0.10).Div(decimal.NewFromInt(1000000)),
	"gemini-2.0-flash-lite":             decimal.NewFromFloat(0.075).Div(decimal.NewFromInt(1000000)),
	"gemini-2.0-flash":                  decimal.NewFromFloat(0.1).Div(decimal.NewFromInt(1000000)),
	"gemini-1.5-flash":                  decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)), // the rate is halved if <= 128K input tokens, fixed below
	"gemini-1.5-pro":                    decimal.NewFromFloat(2.5).Div(decimal.NewFromInt(1000000)),  // the rate is halved if <= 128K input tokens, fixed below
	"gemini-1.0-pro":                    decimal.NewFromFloat(0.5).Div(decimal.NewFromInt(1000000)),
	"command-a-03-2025":                 decimal.NewFromFloat(2.50).Div(decimal.NewFromInt(1000000)),
	"command-a-reasoning-08-2025":       decimal.Zero, // currently free until rate limits reached
	"command-r-08-2024":                 decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)),
	"command-r7b-12-2024":               decimal.NewFromFloat(0.0375).Div(decimal.NewFromInt(1000000)),
	"command-r-plus":                    decimal.NewFromFloat(2.5).Div(decimal.NewFromInt(1000000)),
	"command-r":                         decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)),
	"command-light":                     decimal.NewFromFloat(0.3).Div(decimal.NewFromInt(1000000)),
	"command":                           decimal.NewFromFloat(1).Div(decimal.NewFromInt(1000000)),
	"claude-opus-4-5-20260320":          decimal.NewFromFloat(15).Div(decimal.NewFromInt(1000000)),
	"claude-sonnet-4-5-20260320":        decimal.NewFromFloat(3).Div(decimal.NewFromInt(1000000)),
	"claude-haiku-4-5-20251015":         decimal.NewFromFloat(1).Div(decimal.NewFromInt(1000000)),
	"claude-opus-4-0-20251101":          decimal.NewFromFloat(15).Div(decimal.NewFromInt(1000000)),
	"claude-sonnet-4-0-20250514":        decimal.NewFromFloat(3).Div(decimal.NewFromInt(1000000)),
	"claude-3-7-sonnet-20250219":        decimal.NewFromFloat(3).Div(decimal.NewFromInt(1000000)),
	"claude-3-5-sonnet-20241022":        decimal.NewFromFloat(3).Div(decimal.NewFromInt(1000000)),
	"claude-3-5-haiku-20241022":         decimal.NewFromFloat(0.8).Div(decimal.NewFromInt(1000000)),
	"claude-3-opus-20240229":            decimal.NewFromFloat(15).Div(decimal.NewFromInt(1000000)),
	"claude-3-haiku-20240307":           decimal.NewFromFloat(0.25).Div(decimal.NewFromInt(1000000)),
	"deepseek-chat":                     decimal.NewFromFloat(0.28).Div(decimal.NewFromInt(1000000)), // cache miss rate; cache hit: $0.028/M
	"deepseek-reasoner":                 decimal.NewFromFloat(0.28).Div(decimal.NewFromInt(1000000)), // cache miss rate; cache hit: $0.028/M
	"sonar":                             decimal.NewFromFloat(1.00).Div(decimal.NewFromInt(1000000)),
	"sonar-pro":                         decimal.NewFromFloat(3.00).Div(decimal.NewFromInt(1000000)),
	"sonar-reasoning-pro":               decimal.NewFromFloat(2.00).Div(decimal.NewFromInt(1000000)),
	"sonar-deep-research":               decimal.NewFromFloat(2.00).Div(decimal.NewFromInt(1000000)),
}

// numCentsFromTokens calculates the cost in cents based on token usage and model pricing.
//
// Parameters:
//   - numTokens: The number of tokens used in the request.
//   - model: The model identifier used for processing the request.
//
// Returns:
//   - The computed cost as a decimal.Decimal value.
func numCentsFromTokens(numTokens int, model string) decimal.Decimal {
	rate, ok := modelRates[model]
	if !ok {
		rate = decimal.Zero
		logger.Info("Cost estimation unavailable because model not found: %s", model)
	}
	// halve the rate if the number of tokens is less than 128K and using Google AI Gemini 1.5 flash
	if numTokens <= 128000 && ((model == "gemini-1.5-flash") || (model == "gemini-1.5-pro")) {
		//if model == "gemini-1.5-flash" {
		rate = rate.Div(decimal.NewFromInt(2))
		//} else if model == "gemini-1.5-pro" {
		//	rate = rate.Div(decimal.NewFromInt(4))
		//}
	}
	// double the rate if the number of tokens is greater than 200K for tiered pricing models
	if numTokens > 200000 && (model == "gemini-3-pro-preview" || model == "gemini-2.5-pro") {
		rate = rate.Mul(decimal.NewFromInt(2))
	}
	// Calculate the total cost in cents
	costInCents := decimal.NewFromInt(int64(numTokens)).Mul(rate)

	return costInCents
}
