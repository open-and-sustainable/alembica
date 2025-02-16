/*
Package validation provides JSON schema validation for AI model inputs and outputs.
It ensures that provided JSON structures conform to predefined schema versions.

Supported Schema Types:
  - Input Schema: Validates AI model requests, including metadata, models, and prompts.
  - Output Schema: Ensures AI-generated responses follow expected formatting.

Core Functions:
  - ValidateJSON:
    - Validates a JSON string against a specified schema version and type.
  - ValidateInput:
    - Checks if input data conforms to the defined input schema.

Features:
  - Uses `gojsonschema` for schema-based validation.
  - Supports dynamic schema loading for different versions.
  - Returns detailed validation errors if fields are missing or incorrect.

Example Usage:
	package main

	import (
		"fmt"
		"github.com/open-and-sustainable/alembica/validation"
	)

	func main() {
		jsonInput := `{
			"metadata": {
				"schemaVersion": "v1",
				"timestamp": "2025-02-10T12:00:00Z"
			},
			"models": [
				{
					"provider": "OpenAI",
					"model": "gpt-4o",
					"temperature": 0.7
				}
			],
			"prompts": [
				{
					"promptContent": "Hello",
					"sequenceId": "123",
					"sequenceNumber": 1
				}
			]
		}`

		err := validation.ValidateInput(jsonInput, "v1")
		if err != nil {
			fmt.Println("Validation failed:", err)
		} else {
			fmt.Println("Validation successful!")
		}
	}

*/
package validation
