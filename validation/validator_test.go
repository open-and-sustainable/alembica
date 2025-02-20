package validation

import (
	"testing"
	"strings"
    "github.com/open-and-sustainable/alembica/definitions"
)

func TestValidateJSON(t *testing.T) {
    // Load schemas before running tests
    definitions.LoadSchema("v1", "input")
    definitions.LoadSchema("v1", "output")

	tests := []struct {
		name        string
		jsonInput   string
		version     string
		schemaType  string
		expectsError bool
		errorMsg    string
	}{
		{
			name:       "Valid Input for Input Schema v1",
			jsonInput:  `{"metadata": {"schemaVersion": "v1", "timestamp": "2025-02-10T12:00:00Z"}, "models": [{"provider": "OpenAI", "model": "gpt-4o", "temperature": 0.7}], "prompts": [{"promptContent": "Hello", "sequenceId": "123", "sequenceNumber": 1}]}`,
			version:    "v1",
			schemaType: "input",
			expectsError: false,
		},
		{
			name:       "Invalid Input Missing Required Field",
			jsonInput:  `{"metadata": {"schemaVersion": "v1"}, "models": [{"provider": "OpenAI", "model": "gpt-4o", "temperature": 0.7}], "prompts": [{"promptContent": "Hello", "sequenceId": "123", "sequenceNumber": 1}]}`,
			version:    "v1",
			schemaType: "input",
			expectsError: true,
			errorMsg:   "validation errors: metadata: timestamp is required",
		},
		{
			name:       "Valid Output Schema v1",
			jsonInput:  `{"metadata": {"schemaVersion": "v1"}, "responses": [{"sequenceId": "123", "sequenceNumber": 1, "modelResponse": "This is a response"}]}`,
			version:    "v1",
			schemaType: "output",
			expectsError: false,
		},
		{
			name:       "Invalid Output Schema Missing Required Field",
			jsonInput:  `{"metadata": {"schemaVersion": "v1"}, "responses": [{"sequenceNumber": 1, "modelResponse": "This is a response"}]}`,
			version:    "v1",
			schemaType: "output",
			expectsError: true,
			errorMsg:   "validation errors: responses.0: sequenceId is required",
		},
		{
			name:       "Non-existent Schema Version",
			jsonInput:  `{"metadata": {"schemaVersion": "v99"}, "models": [{"provider": "OpenAI", "model": "gpt-4o", "temperature": 0.7}], "prompts": [{"promptContent": "Hello", "sequenceId": "123", "sequenceNumber": 1}]}`,
			version:    "v99",
			schemaType: "input",
			expectsError: true,
			errorMsg:   "no schemas found for version v99",
		},
		{
			name:       "Non-existent Schema Type",
			jsonInput:  `{"metadata": {"schemaVersion": "v1"}, "models": [{"provider": "OpenAI", "model": "gpt-4o", "temperature": 0.7}], "prompts": [{"promptContent": "Hello", "sequenceId": "123", "sequenceNumber": 1}]}`,
			version:    "v1",
			schemaType: "unknown",
			expectsError: true,
			errorMsg:   "no schema found for type unknown in version v1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateJSON(tt.jsonInput, tt.version, tt.schemaType)
			if tt.expectsError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %s", err)
				}
			}
		})
	}
}

func TestValidateInput(t *testing.T) {
    tests := []struct {
        name        string
        jsonInput   string
        version     string
        expectsError bool
        errorMsg    string
    }{
        {
            name: "Valid Input JSON",
            jsonInput: `{
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
            }`,
            version: "v1",
            expectsError: false,
        },
        {
			name: "Invalid Input - Missing Required Fields",
			jsonInput: `{
				"metadata": {
					"schemaVersion": "v1"
				}
			}`,
			version: "v1",
			expectsError: true,
			errorMsg: "validation errors: (root): models is required; (root): prompts is required; metadata: timestamp is required",
		},
		{
			name: "Invalid Version - Schema Not Found",
			jsonInput: `{
				"metadata": {
					"schemaVersion": "v99",
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
			}`,
			version: "v99",
			expectsError: true,
			errorMsg: "no schema file found for version v99 and type input at v99/input_schema.json",
		},		
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateInput(tt.jsonInput, tt.version)

            if tt.expectsError {
                if err == nil {
                    t.Errorf("Expected error but got nil")
                } else if !strings.Contains(err.Error(), tt.errorMsg) {
                    t.Errorf("Expected error containing '%s', got '%s'", tt.errorMsg, err.Error())
                }
            } else {
                if err != nil {
                    t.Errorf("Expected no error but got: %s", err)
                }
            }
        })
    }
}

