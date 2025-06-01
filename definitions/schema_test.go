package definitions

import (
	"github.com/xeipuuv/gojsonschema"
	"testing"
)

func TestLoadSchema(t *testing.T) {
	// Test loading a valid schema
	err := LoadSchema("v1", "input")
	if err != nil {
		t.Fatalf("Failed to load input schema: %v", err)
	}

	// Ensure the schema is stored in SchemaStore
	if _, exists := SchemaStore["v1"]["input"]; !exists {
		t.Fatalf("Schema v1/input not found in SchemaStore")
	}

	// Test loading a non-existent schema version
	err = LoadSchema("v99", "input")
	if err == nil {
		t.Fatalf("Expected error for non-existent schema version but got nil")
	}

	// Test loading a non-existent schema type
	err = LoadSchema("v1", "unknown")
	if err == nil {
		t.Fatalf("Expected error for non-existent schema type but got nil")
	}
}

func TestSchemaValidation(t *testing.T) {
	// Ensure the schema is loaded
	LoadSchema("v1", "input")

	schema, exists := SchemaStore["v1"]["input"]
	if !exists {
		t.Fatalf("Schema v1/input not found in SchemaStore")
	}

	documentLoader := gojsonschema.NewStringLoader(`{
		"metadata": {
			"schemaVersion": "v1",
			"timestamp": "2025-02-10T12:00:00Z"
		},
		"models": [
			{
				"provider": "OpenAI",
				"api_key": "test-key",
				"model": "gpt-4o",
				"temperature": 0.7,
				"tpm_limit": 1000,
				"rpm_limit": 2000
			}
		],
		"prompts": [
			{
				"promptContent": "Hello",
				"sequenceId": "123",
				"sequenceNumber": 1
			}
		]
	}`)
	result, err := schema.Validate(documentLoader)
	if err != nil {
		t.Fatalf("Error during schema validation: %v", err)
	}

	if !result.Valid() {
		t.Fatalf("Expected valid document but got validation errors: %v", result.Errors())
	}
}
