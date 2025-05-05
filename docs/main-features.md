---
title: Main Features
layout: default
---

# Main Features

`alembica` provides a robust framework for structured data extraction using Large Language Models (LLMs). It simplifies interaction with AI models by ensuring validation, cost assessment, and structured output generation.

## Validation of Input
Ensures that input queries are properly formatted before being processed by LLMs, preventing errors and improving response accuracy.

- Uses **JSON schema validation** to enforce structured input.
- Helps avoid malformed queries that could lead to incorrect or costly API calls.
- Supports multiple schema versions for flexibility.

## Cost Assessment
Calculates token-based processing costs before submitting queries to LLM providers, helping users optimize their usage.

- Estimates costs **based on token consumption** per model and provider.
- Supports **OpenAI, GoogleAI, Cohere, Anthropic, and DeepSeek** pricing models.
- Enables informed decision-making by providing **real-time cost estimates**.

## Data Extraction
Processes unstructured text and converts it into structured datasets, making it easier to analyze and integrate into workflows.

- Extracts named entities, structured responses, and key insights.
- Allows **sequenced query processing**, maintaining logical context across interactions.
- Ensures **schema-compliant structured output**, making data ready for storage or analysis.


<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>
