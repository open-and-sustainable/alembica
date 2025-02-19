/*
Package model provides query functionalities for various large language model (LLM) providers.
It abstracts API interactions, maintains chat history, and ensures JSON-formatted responses.

Supported LLM Providers:
  - OpenAI (GPT-4o, GPT-4 Turbo, GPT-3.5 Turbo)
  - GoogleAI (Gemini-1.5-Pro, Gemini-1.5-Flash)
  - Cohere (Command-R, Command-R+, Command-R7B)
  - Anthropic (Claude-3.5-Sonnet, Claude-3-Haiku, Claude-3-Opus)
  - DeepSeek (DeepSeek-Chat)

Core Functions:
  - QueryLLM: A generic interface for querying different LLMs.
  - queryOpenAI: Handles interactions with OpenAI's GPT models.
  - queryGoogleAI: Handles requests to Google's Gemini models.
  - queryCohere: Processes queries for Cohere's Command models.
  - queryAnthropic: Manages interactions with Anthropic's Claude models.
  - queryDeepSeek: Sends queries to DeepSeek's models.
  - Wait: Ensures compliance with TPM and RPM limits for API queries.

Features:
  - Supports multi-turn chat history for context-aware responses.
  - Ensures all responses are in structured JSON format.
  - Implements automatic model selection and error handling.
  - Enforces API rate limits using Wait function.

Example Usage:
	package main
	
	import (
		"fmt"
		"github.com/open-and-sustainable/alembica/definitions"
		"github.com/open-and-sustainable/alembica/model"
	)

	func main() {
		llm := definitions.Model{
			Provider:    "OpenAI",
			APIKey:      "your-api-key",
			Model:       "gpt-4o",
			Temperature: 0.7,
		}

		prompts := []string{"Hello, AI!", "What is the capital of France?"}
		answers, err := model.DefaultQueryService{}.QueryLLM(prompts, llm)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for i, answer := range answers {
			fmt.Printf("Response %d: %s\n", i+1, answer)
		}
	}
*/
package model

