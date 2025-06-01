/*
Package pricing provides cost estimation utilities for different AI providers based on input token usage.

Supported LLM Providers:
  - OpenAI (GPT-4o, GPT-4 Turbo, GPT-3.5 Turbo)
  - GoogleAI (Gemini-1.5-Pro, Gemini-1.5-Flash)
  - Cohere (Command-R, Command-R+, Command-R7B)
  - Anthropic (Claude-3.5-Sonnet, Claude-3-Haiku, Claude-3-Opus)
  - DeepSeek (DeepSeek-Chat)

Core Components:
  - Cost Calculation:
  - `ComputeCosts`: Processes prompts and calculates their associated costs.
  - `assessPromptCost`: Computes the cost of an individual prompt.
  - Token-Based Pricing:
  - `numCentsFromTokens`: Converts token counts into cost estimates.
  - Model Pricing Rates:
  - `modelRates`: Stores per-model pricing for supported providers.

Features:
  - **Uses per-model pricing rates** to compute input costs dynamically.
  - **Supports batch cost estimation** for multiple prompts.
  - **Handles pricing adjustments** (e.g., discounted rates for Google Gemini under 128K tokens).

Example Usage:

	package main

	import (
		"fmt"
		"github.com/open-and-sustainable/alembica/pricing"
	)

	func main() {
		jsonInput := `{
			"metadata": {
				"schemaVersion": "v1",
				"timestamp": "2025-02-15T12:00:00Z"
			},
			"models": [
				{
					"provider": "OpenAI",
					"api_key": "your-api-key",
					"model": "gpt-4o"
				}
			],
			"prompts": [
				{
					"promptContent": "Hello, AI!",
					"sequenceId": "123",
					"sequenceNumber": 1
				}
			]
		}`

		totalCost := pricing.ComputeCosts(jsonInput)
		fmt.Println("Total estimated cost (in cents):", totalCost)
	}
*/
package pricing
