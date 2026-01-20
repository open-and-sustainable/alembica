---
title: Research & Development
layout: default
---

# Research & Development

## Scope
- **Objective**: `alembica` leverages Large Language Models (LLMs) for systematic synthesis of unstrucutred data sources such as text reported in news corpora.
- **Replicability**: Addresses the challenge of consistent, unbiased analysis, countering the subjective nature of human information extraction.
- **Cost**: More economical than custom AI solutions.

## Contributing

### How to Contribute
We welcome contributions to improve `alembica`, whether you're fixing bugs, adding features, or enhancing documentation:
- **Branching Strategy**: Create a new branch for each set of related changes and submit a pull request via GitHub.
- **Code Reviews**: All submissions undergo thorough review to maintain code quality.

#### Repository Overview
Before contributing, it's helpful to understand how the repository is organized:

- `definitions/` – contains JSON schemas and data structures used for validation.
- `validation/` – validates input, output, and cost calculations using those schemas.
- `llm/` – includes model selection, token counting, and provider-specific query functions.
- `extraction/` – the main entry point that sequences prompts and validates responses.
- `pricing/` – computes token costs for a request.
- `utils/logger/` – provides logging utilities.
- `utils/sharedlib/` – exposes C-shared library functions for use from other languages.
- `test/` – integration tests with real model providers.

Check the [README](../README.md) and documentation under `docs/` for additional details on these packages and how they work together.

### Guidelines
For detailed contribution guidelines, see our [`CONTRIBUTING.md`](CONTRIBUTING.md) and [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md).

## Software Stack

`alembica` is developed in Go, selected for its simplicity and efficiency with concurrent operations. We prioritize the latest stable Go releases to incorporate improvements.

## Open Science Support
`alembica` actively supports Open Science principles.


<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>