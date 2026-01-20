# `alembica` Changelog
All notable changes to this project will be documented in this file.
Releases use semantic versioning as in 'MAJOR.MINOR.PATCH'.
## Change entries
Added: For new features that have been added.
Changed: For changes in existing functionality.
Deprecated: For once-stable features removed in upcoming releases.
Removed: For features removed in this release.
Fixed: For any bug fixes.
Security: For vulnerabilities.

## [0.1.1] - 2026-01-20
### Fixed
- Fixed CI/CD build failure for macOS Intel (AMD64) shared libraries by migrating from deprecated macos-13 runner to macos-15-intel
- Updated macOS ARM64 runner to macos-latest for automatic future updates

## [0.1.0] - 2026-01-20
### Added
- Support for OpenAI GPT-5 series models (GPT-5, GPT-5.1, GPT-5.2, GPT-5-mini, GPT-5-nano)
- Support for Anthropic Claude 4.5 series models (Claude 4.5 Opus, Claude 4.5 Sonnet, Claude 4.5 Haiku)
- Support for Google AI Gemini 2.5 series models (Gemini 2.5 Pro, Gemini 2.5 Flash, Gemini 2.5 Flash-Lite)
- Support for Google AI Gemini 3 series preview models (Gemini 3 Pro Preview, Gemini 3 Flash Preview)
- Support for Perplexity Sonar models (Sonar, Sonar Pro, Sonar Reasoning Pro, Sonar Deep Research)
- Support for Cohere Command A Reasoning model (command-a-reasoning-08-2025)
### Changed
- **BREAKING**: Migrated from community OpenAI SDK (github.com/sashabaranov/go-openai) to official OpenAI SDK (github.com/openai/openai-go/v3)
- Updated Anthropic SDK from v1.2.1 to v1.19.0 with breaking changes requiring versioned model identifiers
- Updated Cohere SDK from v2.14.1 to v2.16.1
- Updated tiktoken-go from v0.1.7 to v0.1.8
- Updated OpenTelemetry dependencies from v1.36.0 to v1.39.0
- Updated golang.org/x dependencies (crypto, net, oauth2, sync, sys, text, time)
- Updated gRPC from v1.72.2 to v1.78.0 and protobuf from v1.36.6 to v1.36.11
- Updated AWS SDK from v1.36.3 to v1.41.1
- Updated Go version from 1.24.2 to 1.25
- Updated OpenAI model pricing: o3 ($10.00 → $2.00/M), gpt-4o ($5.00 → $2.50/M)
- Updated DeepSeek V3.2 pricing: deepseek-chat ($0.27 → $0.28/M), deepseek-reasoner ($0.55 → $0.28/M)
- Updated DeepSeek context window from 64K to 128K tokens
- DeepSeek models now support automatic context caching (cache hit: $0.028/M, 90% cheaper than cache miss)
### Deprecated
- Claude 3.5 Sonnet (v1 and v2) - end of life March 2026
- Claude 3.5 Haiku - end of life July 2026
- Gemini 1.5 Pro 001, Gemini 1.5 Flash 001 - end of life May 27, 2025
- Gemini 1.5 Pro 002, Gemini 1.5 Flash 002, Gemini 1.5 Flash-8B 001 - end of life September 24, 2025
- OpenAI chatgpt-4o-latest - end of life February 17, 2026
- OpenAI gpt-4-0314, gpt-4-1106-preview, gpt-4-0125-preview - end of life March 26, 2026
- OpenAI gpt-3.5-turbo-instruct, babbage-002, davinci-002, gpt-3.5-turbo-1106 - end of life September 28, 2026

## [0.0.8] - 2025-06-01
### Fixed
- Formatting issues and spelling errors.
- Filled in docstrings and tests.

## [0.0.7] - 2025-05-30
### Removed
- removed "claude-3-sonnet" because deprecated by provider and with end of life in July 2025.
- removed support for Gemini 1.0 Pro because of EoL.
### Added
- Support for Claude 3.7 Sonnet model
- Support for Cohere Command-R August 2024 (command-r-08-2024) and Command A (command-a-03-2025)
- Support for DeepSeek-V3-Reasoner
- Support for Google AI Gemini 2.0 Flash and Gemini 2.0 Flash Lite
- Support for Claude 4.0 Sonnet and Claude 4.0 Opus model
- Support for OpenAI models O4Mini, O1, O1Mini, O3, O3Mini, GPT4.1, GPT4.1Mini, GPT4.1Nano

## [0.0.6] - 2025-04-14
### Fixed
- Prompts in sequences where sent to models separately, now they are aware of queries and answers
- Test to verify independence across sequences and contextual awareness within them

## [0.0.5] - 2025-04-14
### Added
- Real test of multi prompt queries
### Changed
- Upgraded all dependencies and Go version
- API keys loading for real API tests
### Fixed
- Anthropic model source code because of breaking chnages in lib not compatible with previous versions

## [0.0.4] - 2025-03-06
### Added
- Sequence number in output schema and object

## [0.0.3] - 2025-03-02
### Fixed
- Schema versioning and loading for validation
- Documentation and Readme

## [0.0.2] - 2025-03-01
### Added
- CI-CD for deployment of C-Shared library
- Documentation on using C-Shared lib in different languages

## [0.0.1] - 2025-02-20
### Added
- Initial project structure
- Code reuse from prismAId tool
- Testing design
