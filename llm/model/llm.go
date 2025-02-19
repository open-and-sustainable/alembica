package model

import (
	"fmt"
	"github.com/open-and-sustainable/alembica/definitions"
)

// QueryService defines the interface for querying various LLM providers.
type QueryService interface {
	// QueryLLM sends a list of prompts to a specified LLM provider and returns responses.
	//
	// Parameters:
	//   - prompts: A list of string prompts to be processed by the LLM.
	//   - llm: The model configuration containing provider details and parameters.
	//
	// Returns:
	//   - A list of responses from the model.
	//   - An error if the request fails.
	QueryLLM(prompts []string, llm definitions.Model) ([]string, error)
}

// DefaultQueryService implements the QueryService interface and routes queries to the appropriate LLM provider.
type DefaultQueryService struct{}

// QueryLLM determines the correct function to use based on the LLM provider and queries the model.
//
// Parameters:
//   - prompts: A list of string prompts to be processed by the LLM.
//   - llm: The model configuration containing provider details and parameters.
//
// Returns:
//   - A list of responses from the model.
//   - An error if the provider is not supported or the query fails.
func (dqs DefaultQueryService) QueryLLM(prompts []string, llm definitions.Model) ([]string, error) {
	var queryFunc func([]string, definitions.Model) ([]string, error)

	switch llm.Provider {
	case "OpenAI":
		queryFunc = queryOpenAI
	case "GoogleAI":
		queryFunc = queryGoogleAI
	case "Cohere":
		queryFunc = queryCohere
	case "Anthropic":
		queryFunc = queryAnthropic
	case "DeepSeek":
		queryFunc = queryDeepSeek
	default:
		return nil, fmt.Errorf("unsupported LLM provider: %s", llm.Provider)
	}

	return queryFunc(prompts, llm)
}
