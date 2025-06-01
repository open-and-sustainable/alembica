/*
Package test provides live integration tests for querying various large language model (LLM) providers.
It ensures correct API interaction, dynamic model selection, and response validation.

Live tests included:
  - TestLiveExtraction: Runs extraction-based queries using real API keys.
  - TestLiveQueryLLM: Sends direct prompts to multiple LLM providers and validates responses.

Features:
  - Dynamically loads API keys from a file.
  - Supports multiple LLM providers, including OpenAI, GoogleAI, Cohere, Anthropic, and DeepSeek.
  - Selects the appropriate model dynamically based on the provider.
  - Measures response time and logs results.
  - Ensures JSON-compliant responses where required.

Requirements:
  - Users must provide a text file named `test_keys.txt` in the same directory.
  - The file should follow the format given in `test_keys.template`, with API keys assigned as `PROVIDER_NAME=API_KEY`.

Example Usage:

	go test -run TestLiveExtraction
	go test -run TestLiveQueryLLM
*/
package test
