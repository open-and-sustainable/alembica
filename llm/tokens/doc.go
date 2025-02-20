/*
Package tokens provides token quantification utilities for multiple AI providers.
It allows estimating the number of tokens in text prompts using provider-specific APIs.

Supported LLM Providers:
  - OpenAI (GPT-4o, GPT-4 Turbo, GPT-3.5 Turbo)
  - GoogleAI (Gemini-1.5-Pro, Gemini-1.5-Flash)
  - Cohere (Command-R, Command-R+, Command-R7B)
  - Anthropic (via OpenAI token counting)
  - DeepSeek (via OpenAI token counting)

Core Components:
  - TokenCounter Interface:
    - Defines a method `GetNumTokensFromPrompt` for retrieving token counts.
  - RealTokenCounter:
    - Implements `TokenCounter` using real API calls.
  - numTokensFromPrompt* Functions:
    - `numTokensFromPromptOpenAI`: Uses OpenAI’s `tiktoken` for token estimation.
    - `numTokensFromPromptGoogleAI`: Calls Google Gemini API for token counting.
    - `numTokensFromPromptCohere`: Queries Cohere API for token quantification.
    - `numTokensFromPromptAnthropic`: Currently mapped to OpenAI token counting.
    - `numTokensFromPromptDeepSeek`: Uses OpenAI token counting as a fallback.

Features:
  - Uses provider-specific APIs to ensure accurate token counts.
  - Supports model-specific tokenization methods.
  - Handles unsupported providers by returning zero tokens.

Example Usage:
	package main

	import (
		"fmt"
		"github.com/open-and-sustainable/alembica/tokens"
	)

	func main() {
		counter := tokens.RealTokenCounter{}
		numTokens := counter.GetNumTokensFromPrompt(
			"Hello, AI!",
			"OpenAI",
			"gpt-4o",
			"your-api-key",
		)

		fmt.Printf("Token count: %d\n", numTokens)
	}

*/
package tokens
