---
title: Getting Started
layout: default
---

# Getting Started

## Installation in Go

### Prerequisites
Before installing `alembica`, ensure you have:
- **Go (latest stable version)** – Required for using `alembica` as a library.
- **API Keys** – Necessary for accessing external LLM providers (OpenAI, Google, Cohere, etc.).
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

## API Reference
Avalable at [https://pkg.go.dev/github.com/open-and-sustainable/alembica](https://pkg.go.dev/github.com/open-and-sustainable/alembica).

## Configuration
`alembica` requires API keys to interact with LLM providers. API keys should be included within the JSON input data structures instead of environment variables or separate configuration files.


<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>