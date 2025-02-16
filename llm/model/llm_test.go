package model

import (
	"errors"
	"testing"

	"github.com/open-and-sustainable/alembica/definitions"
)

// MockQueryService implements QueryService for testing
type MockQueryService struct {
	MockResponse []string
	MockError    error
}

// Implements QueryLLM method for mocking
func (mqs MockQueryService) QueryLLM(prompts []string, llm definitions.Model) ([]string, error) {
	if mqs.MockError != nil {
		return nil, mqs.MockError
	}
	return mqs.MockResponse, nil
}

func TestQueryLLM_MockService(t *testing.T) {
	// Define test cases
	tests := []struct {
		name        string
		prompts     []string
		llm         definitions.Model
		mockResp    []string
		mockErr     error
		expectResp  []string
		expectError bool
	}{
		{
			name:       "Successful query",
			prompts:    []string{"Hello, AI!"},
			llm:        definitions.Model{Provider: "MockProvider"},
			mockResp:   []string{"Hello, human!"},
			mockErr:    nil,
			expectResp: []string{"Hello, human!"},
		},
		{
			name:        "LLM returns an error",
			prompts:     []string{"Error test"},
			llm:         definitions.Model{Provider: "MockProvider"},
			mockResp:    nil,
			mockErr:     errors.New("API failure"),
			expectResp:  nil,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Use MockQueryService
			mockService := MockQueryService{
				MockResponse: tc.mockResp,
				MockError:    tc.mockErr,
			}

			// Call QueryLLM
			resp, err := mockService.QueryLLM(tc.prompts, tc.llm)

			// Check if error expectation matches
			if tc.expectError && err == nil {
				t.Errorf("%s: expected error, got nil", tc.name)
			} else if !tc.expectError && err != nil {
				t.Errorf("%s: unexpected error: %v", tc.name, err)
			}

			// Check response match
			if len(resp) != len(tc.expectResp) {
				t.Errorf("%s: expected response length %d, got %d", tc.name, len(tc.expectResp), len(resp))
			} else {
				for i := range resp {
					if resp[i] != tc.expectResp[i] {
						t.Errorf("%s: expected response %q, got %q", tc.name, tc.expectResp[i], resp[i])
					}
				}
			}
		})
	}
}
