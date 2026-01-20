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

**Cloud/local note:** AWS Bedrock, Azure AI, Vertex AI, and SelfHosted deployments have provider-specific rate limits that are not documented here. Set `tpm_limit` and `rpm_limit` in your input JSON when you need client-side throttling.

## Anthropic
**(January 2026, Tier 1 users)**

Anthropic uses a tiered usage system (Tier 1-4) where rate limits apply at the organization level across all models. The limits below represent Tier 1 thresholds. Higher tiers provide increased limits and are automatically granted based on cumulative API credit purchases and usage history.

**Tier 1 Requirements:**
- Credit Purchase: $5
- Maximum Spend Limit: $100/month

**Tier 1 Rate Limits (applies to all Claude models):**
- **RPM (Requests Per Minute):** 50
- **ITPM (Input Tokens Per Minute):** Varies by model class (see below)
- **OTPM (Output Tokens Per Minute):** Varies by model class (see below)

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model Class</th>
            <th style="text-align: right;">RPM</th>
            <th style="text-align: right;">ITPM</th>
            <th style="text-align: right;">OTPM</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Opus 4.x (4.0, 4.5)</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">20,000</td>
            <td style="text-align: right;">8,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Sonnet 4.x (4.0, 4.5)</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">20,000</td>
            <td style="text-align: right;">8,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Haiku 4.5</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">25,000</td>
            <td style="text-align: right;">10,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3.7 Sonnet</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">40,000</td>
            <td style="text-align: right;">8,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3.5 Sonnet</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">40,000</td>
            <td style="text-align: right;">8,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3.5 Haiku</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">50,000</td>
            <td style="text-align: right;">10,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Opus</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">20,000</td>
            <td style="text-align: right;">8,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Sonnet</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">40,000</td>
            <td style="text-align: right;">8,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Haiku</td>
            <td style="text-align: right;">50</td>
            <td style="text-align: right;">50,000</td>
            <td style="text-align: right;">10,000</td>
        </tr>
    </tbody>
</table>

**Note:** Only uncached input tokens and cache creation tokens count towards ITPM limits for most models. Cached tokens (cache reads) do not count, effectively allowing 5-10x higher throughput when using prompt caching. For detailed information about Anthropic's tiered system, visit their [official rate limits documentation](https://platform.claude.com/docs/en/api/rate-limits).

## Cohere
Cohere production keys have no limit, but trial keys are limited to 20 API calls per minute.

## Perplexity
**(January 2026, Tier 1 users)**

Perplexity uses a tiered usage system (Tier 0-5) where rate limits increase based on cumulative API credit purchases. Tier 1 requires $50+ in lifetime purchases.

**Tier 1 Rate Limits:**

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">RPM</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Sonar Deep Research</td>
            <td style="text-align: right;">10</td>
        </tr>
        <tr>
            <td style="text-align: left;">Sonar Reasoning Pro</td>
            <td style="text-align: right;">150</td>
        </tr>
        <tr>
            <td style="text-align: left;">Sonar Pro</td>
            <td style="text-align: right;">150</td>
        </tr>
        <tr>
            <td style="text-align: left;">Sonar</td>
            <td style="text-align: right;">150</td>
        </tr>
    </tbody>
</table>

**Note:** Tiers are based on cumulative purchases. Higher tiers (2-5) provide significantly increased rate limits. For detailed information, visit [Perplexity's rate limits documentation](https://docs.perplexity.ai/guides/rate-limits-usage-tiers).

## DeepSeek
DeepSeek does not impose rate limits.

## GoogleAI
**(May 2025)**

**Tier 1**:
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
            <td style="text-align: left;">Gemini 2.0 Flash</td>
            <td style="text-align: right;">2,000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 2.0 Flash Lite</td>
            <td style="text-align: right;">4,000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Flash</td>
            <td style="text-align: right;">2,000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Pro</td>
            <td style="text-align: right;">1,000</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">4,000,000</td>
        </tr>
    </tbody>
</table>

## OpenAI
**(May 2025, tier 1 users)**

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
            <td style="text-align: left;">o4-mini</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">o3-mini</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">o3</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">90,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">o1-mini</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">o1</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">90,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-4.1-nano</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-4.1-mini</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
        <tr>
            <td style="text-align: left;">gpt-4.1</td>
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">-</td>
            <td style="text-align: right;">30,000</td>
            <td style="text-align: right;">900,000</td>
        </tr>
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
            <td style="text-align: right;">500</td>
            <td style="text-align: right;">10,000</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">2,000,000</td>
        </tr>
    </tbody>
</table>

<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>
