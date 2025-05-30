---
title: Supported Models
layout: default
---

# Supported Models

`alembica` supports a variety of AI models from leading providers including OpenAI, GoogleAI, Cohere, Anthropic, and DeepSeek.

The table below provides an overview of all supported models, organized by provider. For each model, you can find its maximum input token capacity and the cost per million input tokens. This information helps you select the appropriate model based on your context length requirements and budget considerations.

Each model has specific limits for input size and costs, as summarized below:

## Anthropic

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">Maximum Input Tokens</th>
            <th style="text-align: right;">Cost of 1M Input Tokens</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Claude 4.0 Opus</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$15.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 4.0 Sonnet</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$3.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3.7 Sonnet</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$3.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3.5 Haiku</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$0.80</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Opus</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$15.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3.5 Sonnet</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$3.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">Claude 3 Haiku</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$0.25</td>
        </tr>
    </tbody>
</table>

## Cohere

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">Maximum Input Tokens</th>
            <th style="text-align: right;">Cost of 1M Input Tokens</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Command A</td>
            <td style="text-align: right;">256,000</td>
            <td style="text-align: right;">$2.5</td>
        </tr>
        <tr>
            <td style="text-align: left;">Command R+</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$2.50</td>
        </tr>
        <tr>
            <td style="text-align: left;">Command R7B</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$0.0375</td>
        </tr>
        <tr>
            <td style="text-align: left;">Command R, August 2024</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$0.15</td>
        </tr>
        <tr>
            <td style="text-align: left;">Command R</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$0.15</td>
        </tr>
        <tr>
            <td style="text-align: left;">Command Light</td>
            <td style="text-align: right;">4,096</td>
            <td style="text-align: right;">$0.30</td>
        </tr>
        <tr>
            <td style="text-align: left;">Command</td>
            <td style="text-align: right;">4,096</td>
            <td style="text-align: right;">$1.00</td>
        </tr>
    </tbody>
</table>

## DeepSeek
DeepSeek V3 models cost depends on cache and time of use. Here are the pricing details of maximum rates of input tokens. Please consider also that DeepSeek cost is significantly higher for output tokens.

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">Maximum Input Tokens</th>
            <th style="text-align: right;">Cost of 1M Input Tokens</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">DeepSeek-V3-Chat</td>
            <td style="text-align: right;">64,000</td>
            <td style="text-align: right;">$0.27</td>
        </tr>
        <tr>
            <td style="text-align: left;">DeepSeek-V3-Reasoner</td>
            <td style="text-align: right;">64,000</td>
            <td style="text-align: right;">$0.55</td>
        </tr>
    </tbody>
</table>

## GoogleAI

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">Maximum Input Tokens</th>
            <th style="text-align: right;">Cost of 1M Input Tokens</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Flash</td>
            <td style="text-align: right;">1,048,576</td>
            <td style="text-align: right;">$0.15</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.5 Pro</td>
            <td style="text-align: right;">2,097,152</td>
            <td style="text-align: right;">$2.50</td>
        </tr>
        <tr>
            <td style="text-align: left;">Gemini 1.0 Pro</td>
            <td style="text-align: right;">32,760</td>
            <td style="text-align: right;">$0.50</td>
        </tr>
    </tbody>
</table>

## OpenAI

<table class="table-spacing">
    <thead>
        <tr>
            <th style="text-align: left;">Model</th>
            <th style="text-align: right;">Maximum Input Tokens</th>
            <th style="text-align: right;">Cost of 1M Input Tokens</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td style="text-align: left;">o4 Mini</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$1.10</td>
        </tr>
        <tr>
            <td style="text-align: left;">o3 Mini</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$1.10</td>
        </tr>
        <tr>
            <td style="text-align: left;">o3</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$10.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">o1 Mini</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$1.10</td>
        </tr>
        <tr>
            <td style="text-align: left;">o1</td>
            <td style="text-align: right;">200,000</td>
            <td style="text-align: right;">$15.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">GPT-4.1 Nano</td>
            <td style="text-align: right;">1,000,000</td>
            <td style="text-align: right;">$0.10</td>
        </tr>
        <tr>
            <td style="text-align: left;">GPT-4.1 Mini</td>
            <td style="text-align: right;">1,000,000</td>
            <td style="text-align: right;">$0.40</td>
        </tr>
        <tr>
            <td style="text-align: left;">GPT-4.1</td>
            <td style="text-align: right;">1,000,000</td>
            <td style="text-align: right;">$2.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">GPT-4o Mini</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$0.15</td>
        </tr>
        <tr>
            <td style="text-align: left;">GPT-4o</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$5.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">GPT-4 Turbo</td>
            <td style="text-align: right;">128,000</td>
            <td style="text-align: right;">$10.00</td>
        </tr>
        <tr>
            <td style="text-align: left;">GPT-3.5 Turbo</td>
            <td style="text-align: right;">16,385</td>
            <td style="text-align: right;">$0.50</td>
        </tr>
    </tbody>
</table>

<div id="wcb" class="carbonbadge"></div>
<script src="https://unpkg.com/website-carbon-badges@1.1.3/b.min.js" defer></script>
