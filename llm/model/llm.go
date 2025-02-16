package model

import (
	"fmt"

	"github.com/open-and-sustainable/alembica/definitions"
)

type QueryService interface {
    QueryLLM(prompts []string, llm definitions.Model) ([]string, error)
}

type DefaultQueryService struct{}

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

