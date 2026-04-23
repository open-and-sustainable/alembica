# ![logo](https://raw.githubusercontent.com/open-and-sustainable/alembica/main/docs/assets/images/logo.png) `alembica`

**Open Science Software for Semantic Synthesis and Extraction of Information from Unstructured Sources.**

[![Go Reference](https://pkg.go.dev/badge/github.com/open-and-sustainable/alembica.svg)](https://pkg.go.dev/github.com/open-and-sustainable/alembica) [![Go Report Card](https://goreportcard.com/badge/github.com/open-and-sustainable/alembica)](https://goreportcard.com/report/github.com/open-and-sustainable/alembica) [![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.14899666.svg)](https://doi.org/10.5281/zenodo.14899666)
[![MCP Registry](https://img.shields.io/badge/MCP%20Registry-Alembica%20MCP-0A7B83)](https://registry.modelcontextprotocol.io/?q=alembica)

---

## About

`alembica` simplifies the use of **Large Language Models (LLMs)** to extract structured datasets from unstructured corpora of text.
It provides a **flexible and scalable framework** to process, synthesize, and transform textual information into structured formats suitable for analysis and further processing.

Supports **OpenAI, Google AI, Anthropic, Cohere, DeepSeek, Perplexity, AWS Bedrock, Azure AI, Vertex AI, and Self-Hosted OpenAI-compatible** providers.

---

## Installation (Go)

To install `alembica` in Go, run:

```sh
go get github.com/open-and-sustainable/alembica
```

If you want to use `alembica` in **other programming languages**, check out the C-Shared Library in the [User Guide](https://open-and-sustainable.github.io/alembica/).

---

## Documentation

**[User Guide](https://open-and-sustainable.github.io/alembica/)** – Learn how to use `alembica` in different programming languages.
**[API Reference](https://pkg.go.dev/github.com/open-and-sustainable/alembica)** – Explore the Go package documentation.

## MCP Server
`alembica` includes an optional MCP server for agent tool access.

The MCP server can be used from a locally built Go binary, the GHCR container image, or the MCP Registry.

Available in the official MCP Registry:
- https://registry.modelcontextprotocol.io/?q=alembica
- Server ID: `io.github.open-and-sustainable/alembica-mcp`

See the User Guide MCP page for installation options, run commands, and tool schemas.

## MCP Registry publishing
The MCP server is published to the MCP Registry by GitHub Actions when a version tag such as `v1.2.3` is pushed.

The workflow uses GitHub OIDC for authentication, and each new registry version is published explicitly by CI rather than implicitly from GitHub Releases.

---

## Features

- **Validation of Input** – Ensures that queries are correctly formatted to support proper interaction with models.
- **Cost Assessment** – Calculates token costs based on the requested extraction and different model pricing.
- **Data Extraction** – Processes unstructured text and transforms it into structured datasets for further analysis.

Note: Cost estimation is not supported for Self-Hosted, AWS Bedrock, Azure AI, or Vertex AI providers and will return zero.

Optional model fields for cloud/local providers:
- `base_url` and `api_version` for Azure/OpenAI-compatible endpoints
- `region` for AWS Bedrock
- `project_id` and `location` for Vertex AI

Use `schemaVersion: "v2"` when you need these cloud/local provider fields or non-enumerated model IDs.

---

## Authors & Contributions

**Author**: Riccardo Boero - [ribo@nilu.no](mailto:ribo@nilu.no)

Contributions are welcome!

---

## License

`alembica` is licensed under the **GNU AFFERO GENERAL PUBLIC LICENSE, Version 3**.

![AGPL License](https://www.gnu.org/graphics/agplv3-155x51.png)

---

## Citation

> **Boero, R. (2025).** *`alembica` - Open Science Software for Semantic Synthesis and Extraction of Information from Unstructured Sources.* Zenodo.
> [https://doi.org/10.5281/zenodo.14899666](https://doi.org/10.5281/zenodo.14899666)
