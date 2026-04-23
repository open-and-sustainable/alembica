---
title: MCP Server
layout: default
---

# MCP Server

The optional MCP server exposes `alembica` tools to agents over `stdio`.

The MCP server supports three usage patterns:

- local installation from the Go source
- container-based execution from the GHCR image
- registry-based use through the MCP Registry

The MCP server only supports schema version `v2`.

## Available Tools

- `alembica_validate_input`
- `alembica_validate_output`
- `alembica_extract`
- `alembica_compute_costs`
- `alembica_list_schemas`

Agents discover tool schemas via `tools/list` and call them with `tools/call`.

## Use from Go Source

Use this when you want a local binary built directly from the project source.

```sh
go install github.com/open-and-sustainable/alembica/cmd/alembica-mcp@latest
alembica-mcp
```

## Use from the GHCR Container Image

Use this when you want to run the MCP server without a local Go toolchain.

```sh
docker pull ghcr.io/open-and-sustainable/alembica-mcp:0.3.3
docker run --rm -i ghcr.io/open-and-sustainable/alembica-mcp:0.3.3
```

Replace `0.3.3` with the released version you want to run.

## Use from the MCP Registry

Use this when your agent platform supports MCP Registry server discovery and installation.

The Alembica MCP server is published explicitly by GitHub Actions on pushed version tags such as `v0.3.2`, using GitHub OIDC authentication and the registry publisher CLI.

The registry entry points to the published OCI package for the MCP server, so agents can resolve a versioned package rather than a repository source tree.

Registry references:

- Discovery page: `https://registry.modelcontextprotocol.io/?q=alembica`
- Server ID: `io.github.open-and-sustainable/alembica-mcp`
- Versioned API example: `https://registry.modelcontextprotocol.io/v0.1/servers/io.github.open-and-sustainable%2Falembica-mcp/versions/0.3.3`

## Example Requests
Example `tools/list` request:
```json
{ "jsonrpc": "2.0", "id": 1, "method": "tools/list" }
```

Example `tools/call` request:
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/call",
  "params": {
    "name": "alembica_extract",
    "arguments": {
      "input_json": "{\"metadata\":{\"schemaVersion\":\"v2\",\"timestamp\":\"2026-01-20T00:00:00Z\"},\"models\":[{\"provider\":\"SelfHosted\",\"model\":\"llama3.1:70b\",\"base_url\":\"http://localhost:11434/v1\",\"temperature\":0.7}],\"prompts\":[{\"sequenceId\":\"1\",\"sequenceNumber\":1,\"promptContent\":\"Respond with JSON: {\\\"hello\\\":\\\"world\\\"}\"}]}"
    }
  }
}
```

In all cases, clients interact with the server through standard MCP `tools/list` and `tools/call` requests.

<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>
