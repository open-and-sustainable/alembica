---
title: Getting Started
layout: default
---

# Getting Started

## Installation in Go

### Prerequisites
Before installing `alembica`, ensure you have:
- **Go (latest stable version)** – Required for using `alembica` as a library.
- **API Keys** – Necessary for accessing external LLM providers (OpenAI, Google, Cohere, Anthropic, DeepSeek, Perplexity, AWS Bedrock, Azure AI, Vertex AI, etc.).
- **Git** – Recommended for managing the source code **if developing `alembica`**.

### Install `alembica`
```sh
go get github.com/open-and-sustainable/alembica
```

## Using alembica in Your Go Project
To use `alembica` as a library in your Go code:
```go
package main

import (
    "fmt"
    "github.com/open-and-sustainable/alembica/extraction"
)

func main() {
    inputJSON := `{
        "metadata": { "schemaVersion": "v1" },
        "models": [{ "provider": "OpenAI", "model": "gpt-4o", "api_key": "your-openai-key" }],
        "prompts": [{ "sequenceId": "1", "sequenceNumber": 1, "promptContent": "Extract structured data from this text." }]
    }`

    result, err := extraction.Extract(inputJSON)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Extraction Result:", result)
    }
}
```

## Schema Versioning
Use `schemaVersion: "v2"` when you need cloud/local providers (AWS Bedrock, Azure AI, Vertex AI, SelfHosted) or non-enumerated model IDs. Existing `v1` inputs remain supported.

## Cloud and Local Providers (Examples)
Use the same `extraction.Extract` call; only the JSON fields change. These examples use `schemaVersion: "v2"`.

### Self-Hosted (OpenAI-Compatible)
```json
{
  "metadata": { "schemaVersion": "v2", "timestamp": "2026-01-20T00:00:00Z" },
  "models": [
    {
      "provider": "SelfHosted",
      "model": "llama3.1:70b",
      "api_key": "",
      "base_url": "http://localhost:11434/v1",
      "temperature": 0.7
    }
  ],
  "prompts": [
    { "sequenceId": "1", "sequenceNumber": 1, "promptContent": "Respond with JSON: {\"hello\":\"world\"}" }
  ]
}
```

### AWS Bedrock
```json
{
  "metadata": { "schemaVersion": "v2", "timestamp": "2026-01-20T00:00:00Z" },
  "models": [
    {
      "provider": "AWSBedrock",
      "model": "meta.llama3-70b-instruct-v1:0",
      "region": "us-east-1",
      "temperature": 0.7
    }
  ],
  "prompts": [
    { "sequenceId": "1", "sequenceNumber": 1, "promptContent": "Respond with JSON: {\"hello\":\"world\"}" }
  ]
}
```

### Azure AI (Azure OpenAI)
```json
{
  "metadata": { "schemaVersion": "v2", "timestamp": "2026-01-20T00:00:00Z" },
  "models": [
    {
      "provider": "AzureAI",
      "model": "my-deployment-name",
      "api_key": "your-azure-key",
      "base_url": "https://your-resource.openai.azure.com",
      "api_version": "2024-06-01",
      "temperature": 0.7
    }
  ],
  "prompts": [
    { "sequenceId": "1", "sequenceNumber": 1, "promptContent": "Respond with JSON: {\"hello\":\"world\"}" }
  ]
}
```

### Vertex AI (Model Garden)
```json
{
  "metadata": { "schemaVersion": "v2", "timestamp": "2026-01-20T00:00:00Z" },
  "models": [
    {
      "provider": "VertexAI",
      "model": "meta/llama3-70b-instruct",
      "project_id": "your-gcp-project",
      "location": "us-central1",
      "temperature": 0.7
    }
  ],
  "prompts": [
    { "sequenceId": "1", "sequenceNumber": 1, "promptContent": "Respond with JSON: {\"hello\":\"world\"}" }
  ]
}
```

## API Reference
Available at [https://pkg.go.dev/github.com/open-and-sustainable/alembica](https://pkg.go.dev/github.com/open-and-sustainable/alembica).

## Configuration
`alembica` requires API keys to interact with LLM providers. API keys should be included within the JSON input data structures instead of environment variables or separate configuration files.


<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>
