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
