package pricing

import (
	"fmt"

	"github.com/open-and-sustainable/alembica/utils/logger"

	"github.com/shopspring/decimal"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	openai "github.com/sashabaranov/go-openai"
)

// Define a map to hold the rates for each model
var modelRates = map[string]decimal.Decimal{ // dollar prices per input M token
	openai.O4Mini:                          decimal.NewFromFloat(1.10).Div(decimal.NewFromInt(1000000)),
	openai.O1:                              decimal.NewFromFloat(15).Div(decimal.NewFromInt(1000000)),
	openai.O1Mini:                          decimal.NewFromFloat(1.10).Div(decimal.NewFromInt(1000000)),
	openai.O3:                              decimal.NewFromFloat(10).Div(decimal.NewFromInt(1000000)),
	openai.O3Mini:                          decimal.NewFromFloat(1.10).Div(decimal.NewFromInt(1000000)),
	openai.GPT4Dot1:                        decimal.NewFromFloat(2.00).Div(decimal.NewFromInt(1000000)),
	openai.GPT4Dot1Mini:                    decimal.NewFromFloat(0.40).Div(decimal.NewFromInt(1000000)),
	openai.GPT4Dot1Nano:                    decimal.NewFromFloat(0.10).Div(decimal.NewFromInt(1000000)),
	openai.GPT4oMini:                       decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)),
	openai.GPT4o:                           decimal.NewFromFloat(5).Div(decimal.NewFromInt(1000000)),
	openai.GPT4Turbo:                       decimal.NewFromFloat(10).Div(decimal.NewFromInt(1000000)),
	openai.GPT4:                            decimal.NewFromFloat(30).Div(decimal.NewFromInt(1000000)),
	openai.GPT432K:                         decimal.NewFromFloat(60).Div(decimal.NewFromInt(1000000)),
	openai.GPT3Dot5Turbo:                   decimal.NewFromFloat(0.5).Div(decimal.NewFromInt(1000000)),
	openai.GPT3Dot5TurboInstruct:           decimal.NewFromFloat(1.5).Div(decimal.NewFromInt(1000000)),
	"gemini-2.0-flash-lite":                decimal.NewFromFloat(0.075).Div(decimal.NewFromInt(1000000)),
	"gemini-2.0-flash":                     decimal.NewFromFloat(0.1).Div(decimal.NewFromInt(1000000)),
	"gemini-1.5-flash":                     decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)), // the rate is halved if <= 128K input tokens, fixed below
	"gemini-1.5-pro":                       decimal.NewFromFloat(2.5).Div(decimal.NewFromInt(1000000)),  // the rate is halved if <= 128K input tokens, fixed below
	"gemini-1.0-pro":                       decimal.NewFromFloat(0.5).Div(decimal.NewFromInt(1000000)),
	"command-a-03-2025":                    decimal.NewFromFloat(2.50).Div(decimal.NewFromInt(1000000)),
	"command-r-08-2024":                    decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)),
	"command-r7b-12-2024":                  decimal.NewFromFloat(0.0375).Div(decimal.NewFromInt(1000000)),
	"command-r-plus":                       decimal.NewFromFloat(2.5).Div(decimal.NewFromInt(1000000)),
	"command-r":                            decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)),
	"command-light":                        decimal.NewFromFloat(0.3).Div(decimal.NewFromInt(1000000)),
	"command":                              decimal.NewFromFloat(1).Div(decimal.NewFromInt(1000000)),
	string(anthropic.ModelClaudeOpus4_0):   decimal.NewFromFloat(15).Div(decimal.NewFromInt(1000000)),
	string(anthropic.ModelClaudeSonnet4_0): decimal.NewFromFloat(3).Div(decimal.NewFromInt(1000000)),
	string(anthropic.ModelClaude3_7SonnetLatest):   decimal.NewFromFloat(3).Div(decimal.NewFromInt(1000000)),
	string(anthropic.ModelClaude3_5SonnetLatest):   decimal.NewFromFloat(3).Div(decimal.NewFromInt(1000000)),
	string(anthropic.ModelClaude3_5HaikuLatest):    decimal.NewFromFloat(0.8).Div(decimal.NewFromInt(1000000)),
	string(anthropic.ModelClaude3OpusLatest):       decimal.NewFromFloat(15).Div(decimal.NewFromInt(1000000)),
	string(anthropic.ModelClaude_3_Haiku_20240307): decimal.NewFromFloat(0.25).Div(decimal.NewFromInt(1000000)),
	"deepseek-chat":     decimal.NewFromFloat(0.27).Div(decimal.NewFromInt(1000000)),
	"deepseek-reasoner": decimal.NewFromFloat(0.55).Div(decimal.NewFromInt(1000000)),
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
		logger.Error(fmt.Println("Cost estimation unavailable because model not found:", model))
	}
	// halve the rate if the number of tokens is less than 128K and using Google AI Gemini 1.5 flash
	if numTokens <= 128000 && ((model == "gemini-1.5-flash") || (model == "gemini-1.5-pro")) {
		//if model == "gemini-1.5-flash" {
		rate = rate.Div(decimal.NewFromInt(2))
		//} else if model == "gemini-1.5-pro" {
		//	rate = rate.Div(decimal.NewFromInt(4))
		//}
	}
	// Calculate the total cost in cents
	costInCents := decimal.NewFromInt(int64(numTokens)).Mul(rate)

	return costInCents
}
