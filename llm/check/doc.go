/*
Package check provides utilities for selecting appropriate AI models, enforcing token limits,
and ensuring compatibility with various LLM providers.

Supported LLM Providers:
  - OpenAI (GPT-4o, GPT-4 Turbo, GPT-3.5 Turbo)
  - GoogleAI (Gemini-1.5-Pro, Gemini-1.5-Flash)
  - Cohere (Command-R, Command-R+, Command-R7B)
  - Anthropic (Claude-3.5-Sonnet, Claude-3-Haiku, Claude-3-Opus)
  - DeepSeek (DeepSeek-Chat)

Core Components:
  - Model Selection:
  - `GetModel`: Selects an optimized model based on the provider and user input.
  - Token Limit Checking:
  - `RunInputLimitsCheck`: Verifies if a prompt exceeds the maximum token limit for a given model.
  - Model-Specific Limits:
  - `ModelMaxTokens`: Stores the maximum token capacities for supported models.

Features:
  - **Dynamic model selection** based on token constraints and cost optimization.
  - **Enforces per-model token limits** to prevent exceeding API constraints.
  - **Supports structured logging** to debug model selection and validation.

Example Usage:

	package main

	import (
		"fmt"
		"github.com/open-and-sustainable/alembica/check"
	)

	func main() {
		model := check.GetModel("Hello, AI!", "OpenAI", "", "api-key")
		fmt.Println("Selected model:", model)

		err := check.RunInputLimitsCheck("This is a test prompt.", "OpenAI", model, "api-key", nil)
		if err != nil {
			fmt.Println("Limit check failed:", err)
		} else {
			fmt.Println("Prompt is within token limits.")
		}
	}
*/
package check
