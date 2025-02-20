package definitions

import (
)

// Define input structures
type InputMetadata struct {
    Version      string `json:"version"`
    SchemaVersion string `json:"schemaVersion"`
    Timestamp    string `json:"timestamp"`
}

type Model struct {
    Provider        string `json:"provider"`
    APIKey          string `json:"api_key"`
    Model           string `json:"model"`
    Temperature     float64 `json:"temperature"`
    TPMLimit        int `json:"tpm_limit"`
    RPMLimit        int `json:"rpm_limit"`
}

type Prompt struct {
    PromptContent string `json:"promptContent"`
    SequenceID    string `json:"sequenceId"`
    SequenceNumber int    `json:"sequenceNumber"`
}

type Input struct {
    Metadata InputMetadata `json:"metadata"`
    Models   []Model       `json:"models"`
    Prompts  []Prompt      `json:"prompts"`
}

// Define output structures
type OutputMetadata struct {
    SchemaVersion string `json:"schemaVersion"`
}

type Response struct {
    Provider       string      `json:"provider"`
    Model          string      `json:"model"`
    SequenceID     string      `json:"sequenceId"`
    ModelResponses []string    `json:"modelResponses"`
    Error          *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type Output struct {
    Metadata  OutputMetadata `json:"metadata"`
    Responses []Response     `json:"responses"`
}

// Define cost structures
type CostMetadata struct {
    SchemaVersion string `json:"schemaVersion"`
    Currency      string `json:"currency"`
}

type Cost struct {
    SequenceID string  `json:"sequenceId"`
    Provider   string  `json:"provider"`
    Model      string  `json:"model"`
    Cost       float64 `json:"cost"`
}

type CostOutput struct {
    Metadata CostMetadata `json:"metadata"`
    Costs    []Cost       `json:"costs"`
}
