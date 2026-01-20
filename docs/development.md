---
title: Development
layout: default
---

# Development

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

