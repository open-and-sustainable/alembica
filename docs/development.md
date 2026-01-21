---
title: Development
layout: default
---

# Development

`alembica` is an open-source tool written in pure Go, designed to bridge the gap between unstructured text and structured data through LLM-powered extraction.

## Architecture

The project follows a modular architecture with clear separation of concerns:

- **`definitions/`**: Core data structures and input/output schemas (supports v1 and v2 schema versions)
- **`validation/`**: Schema validation helpers ensuring data integrity
- **`extraction/`**: Prompt sequencing engine and model execution orchestration
- **`llm/`**: Provider integrations (OpenAI, Anthropic, Google AI, Cohere, DeepSeek, Perplexity, AWS Bedrock, Azure AI, Vertex AI, Self-Hosted)
- **`pricing/`**: Token-based cost estimation for cloud providers
- **`utils/`**: Logging utilities and shared library exports for cross-language interoperability

The architecture enables flexible extraction workflows where users define prompts, specify models, and chain sequential operations to transform unstructured text into structured JSON datasets.

## User Approach

`alembica` can be used in multiple ways to fit different workflows:

1. **As a Go Package**: Import directly into Go applications for native integration
2. **As a C-Shared Library**: Use from Python, R, C#, or other languages via FFI bindings
3. **Via MCP Server**: Integrate with AI agents and autonomous systems through the Model Context Protocol

Users define extraction tasks through JSON input files that specify:
- Schema version and metadata
- Model configurations (provider, model ID, temperature, optional endpoints)
- Prompt sequences with content and ordering
- Optional output validation schemas

## MCP Server Integration

The optional `alembica-mcp` server exposes core functionality as tools for AI agents:

- **`alembica_validate_input`**: Validates input schema before processing
- **`alembica_validate_output`**: Ensures extracted data matches expected schema
- **`alembica_extract`**: Executes the full extraction pipeline
- **`alembica_compute_costs`**: Estimates token costs for planned operations
- **`alembica_list_schemas`**: Lists available schema versions

The MCP server uses stdio transport and follows JSON-RPC 2.0 protocol, supporting only schema version `v2`. This enables agents to autonomously perform semantic extraction tasks as part of larger workflows.

Install with: `go install github.com/open-and-sustainable/alembica/cmd/alembica-mcp@latest`

## Possible Development Directions

Future enhancements being considered:

1. **Enhanced Provider Support**: Adding more LLM providers and keeping up with new model releases
2. **Streaming Support**: Real-time extraction for large documents with progressive output
3. **Batch Processing**: Optimized handling of multiple documents in parallel
4. **Schema Evolution**: Tools for migrating between schema versions and managing backwards compatibility
5. **Caching Layer**: Reduce redundant API calls by caching intermediate results
6. **Advanced Validation**: Richer output schema validation with custom rules and constraints
7. **Observability**: Enhanced logging, metrics, and tracing for production deployments
8. **Template Library**: Pre-built extraction templates for common use cases (citations, entities, summaries)
9. **Multi-modal Support**: Extending extraction capabilities to images and PDFs
10. **Fine-tuning Integration**: Tools to generate training data from extraction results for model improvement

Contributions addressing these or other improvements are welcome!

## Tests
```sh
go test ./...
```

## Project Layout
- `definitions/`: input/output schema and core structures
- `validation/`: schema validation helpers
- `extraction/`: prompt sequencing and model execution
- `llm/`: provider integrations and token checks
- `pricing/`: cost estimation
- `utils/`: logging and shared library exports


<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>
