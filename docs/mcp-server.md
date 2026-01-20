---
title: MCP Server
layout: default
---

# MCP Server

The optional MCP server exposes `alembica` tools to agents over stdio.

## Install and Run
```sh
go install github.com/open-and-sustainable/alembica/cmd/alembica-mcp@latest
alembica-mcp
```

## Tools
- `alembica_validate_input`
- `alembica_validate_output`
- `alembica_extract`
- `alembica_compute_costs`
- `alembica_list_schemas`

Agents discover tool schemas via `tools/list` and call them with `tools/call`. The MCP server only supports schema version `v2`.

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
