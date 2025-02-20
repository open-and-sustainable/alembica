/*
Package definitions provides core data structures and schema validation mechanisms
for LLM input and output formats. It ensures consistency in model requests and responses.

Supported Schema Types:
  - Input Schema: Defines metadata, models, and prompts used for LLM requests.
  - Output Schema: Defines metadata and responses generated by LLMs.

Core Components:
  - Data Structures:
    - `Input`: Represents an AI model request with metadata, models, and prompts.
    - `Output`: Represents an AI model response with metadata and generated answers.
  - Schema Management:
    - `LoadSchema`: Loads and stores JSON schemas for input/output validation.
    - `SchemaStore`: Holds different versions of schema files.
  - Token Counting:
    - `RealTokenCounter`: Calculates token counts using provider-specific APIs.

Features:
  - Uses `gojsonschema` to validate JSON structures.
  - Supports multiple schema versions dynamically.
  - Defines standard input and output formats for AI model interactions.

Example Usage:
	package main

	import (
		"fmt"
		"github.com/open-and-sustainable/alembica/definitions"
	)

	func main() {
		// Load the input schema for version v1
		err := definitions.LoadSchema("v1", "input")
		if err != nil {
			fmt.Println("Schema loading failed:", err)
			return
		}

		// Example input structure
		input := definitions.Input{
			Metadata: definitions.InputMetadata{
				Version:       "1.0",
				SchemaVersion: "v1",
				Timestamp:     "2025-02-10T12:00:00Z",
			},
			Models: []definitions.Model{
				{
					Provider:    "OpenAI",
					APIKey:      "your-api-key",
					Model:       "gpt-4o",
					Temperature: 0.7,
					TPMLimit:    1000,
					RPMLimit:    2000,
				},
			},
			Prompts: []definitions.Prompt{
				{
					PromptContent: "Hello, AI!",
					SequenceID:    "123",
					SequenceNumber: 1,
				},
			},
		}

		// Print input structure
		fmt.Printf("Generated Input: %+v\n", input)
	}

*/
package definitions
