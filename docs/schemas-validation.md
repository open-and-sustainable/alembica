---
title: Schemas & Validation
layout: default
---

# Schemas & Validation

`alembica` validates input and output JSON using versioned JSON Schemas.

## Schema Versions
- `v1`: legacy providers and enumerated model IDs.
- `v2`: cloud/local providers and non-enumerated model IDs.

Set the schema version in your input JSON:
```json
{ "metadata": { "schemaVersion": "v2", "timestamp": "2026-01-20T00:00:00Z" } }
```

## Validation APIs
- `validation.ValidateInput(json, version)`
- `validation.ValidateOutput(json, version)`
- `validation.ValidateCost(json, version)`

