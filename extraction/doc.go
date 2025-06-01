/*
Package extraction handles structured information extraction from text inputs.
It processes user prompts, sequences them correctly, and queries large language models (LLMs)
for responses using an abstracted query service.

Core Functionality:
  - Extract: Main function that processes input JSON, sequences prompts, and queries the appropriate LLM.
  - Ensures correct prompt sequencing before calling models.
  - Calls validation on the output to maintain schema integrity.

Features:
  - Supports multiple LLM providers via `model.DefaultQueryService`.
  - Maintains structured responses that match predefined schema.
  - Implements waiting mechanisms to respect rate limits of LLM APIs.

Example Usage:

	package main

	import (
		"fmt"
		"github.com/open-and-sustainable/alembica/extraction"
	)

	func main() {
		inputJSON := `{"metadata": {"schemaVersion": "1.0"}, "models": [...], "prompts": [...]}`
		output, err := extraction.Extract(inputJSON)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Extracted Output:", output)
	}
*/
package extraction
