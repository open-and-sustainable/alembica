---
title: Rate Limits
layout: default
---

# Rate Limits

`alembica` supports management of rate limitations imposed by various AI service providers. Understanding these limits is essential for efficient application development and deployment.

## Understanding Rate Limits

Rate limits control how frequently you can access AI models and how much data you can process in a given timeframe. These limits vary by provider and subscription tier, and typically include:

- **RPM**: Requests per minute
- **RPD**: Requests per day
- **TPM**: Tokens per minute
- **TPD**: Tokens per day

Exceeding these limits may result in request throttling or errors. `alembica` helps manage these constraints by providing appropriate fallback mechanisms and retry strategies.

**Disclaimer:** Daily limits (RPD and TPD) are not currently supported by `alembica`. Users are responsible for implementing and respecting these constraints on their own within their applications.

## OpenAI
**(August 2024, tier 1 users)**

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">RPD</th>
            <th style="text-align: right;">TPM</th>
            <th style="text-align: right;">Batch Queue Limit</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">gpt-4o</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">90,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-4o-mini</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">10,000</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-4-turbo</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">90,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-3.5-turbo</td>
            <td style="text-align: right;">3,500</td>
            <td style="text-align: right;">10,000</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
    </tbody>
</table>


## GoogleAI
**(October 2024)**

**Free Tier**:
<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">RPD</th>
            <th style="text-align: right;">TPM</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Flash</td>
            <td style="text-align: right;">15</td>
            <td style="text-align: right;">1,500</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Pro</td>
            <td style="text-align: right;">2</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">32,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.0 Pro</td>
            <td style="text-align: right;">15</td>
            <td style="text-align: right;">1,500</td>
            <td style="text-align: right;">32,000</td>
        </tr>
    </tbody>
</table>

**Pay-as-you-go**:
<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">RPD</th>
            <th style="text-align: right;">TPM</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Flash</td>
            <td style="text-align: right;">2000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Pro</td>
            <td style="text-align: right;">1000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.0 Pro</td>
            <td style="text-align: right;">360</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">120,000</td>
        </tr>
    </tbody>
</table>

## Cohere
Cohere production keys have no limit, but trial keys are limited to 20 API calls per minute.

## Anthropic
**(November 2024, tier 1 users)**
<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">TPM</th>
            <th style="text-align: right;">TPD</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Claude 3.5 Sonnet</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">40,000</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3.5 Haiku</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">50,000</td>
            <td style="text-align: right;">5,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Opus</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">20,000</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Sonnet</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">40,000</td>
            <td style="text-align: right;">1,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Haiku</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">50,000</td>
            <td style="text-align: right;">5,000,000</td>
        </tr>
    </tbody>
</table>

## DeepSeek
DeepSeek does not impose rate limits.

<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>
