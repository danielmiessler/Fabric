# Changelog

## v1.4.390 (2026-01-24)

MiniMax unexpected status code: 404 from provider MiniMax, response body: 404 page not found
### PR [#1964](https://github.com/danielmiessler/Fabric/pull/1964) by [jessesep](https://github.com/jessesep): feat: add design system, golden rules, and discord structure patterns

- Added create_design_system pattern to generate CSS design systems with tokens, typography scales, and dark/light mode support from requirements
- Added create_golden_rules pattern to extract implicit and explicit rules from codebases into testable, enforceable guidelines  
- Added analyze_discord_structure pattern to audit Discord server organization, permissions, and naming conventions

### PR [#1967](https://github.com/danielmiessler/Fabric/pull/1967) by [ksylvan](https://github.com/ksylvan): chore: add MiniMax provider support and update API endpoints

- Add MiniMax provider configuration with API endpoint updates
- Implement NeedsRawMode method for MiniMax model handling
- Define static MiniMax model list with M2 variants
- Add Infermatic and Novita to VS Code extensions
- Configure MiniMax models as static discovery

## v1.4.389 (2026-01-23)

### PR [#1960](https://github.com/danielmiessler/Fabric/pull/1960) by [ksylvan](https://github.com/ksylvan): fix: consume all positional arguments as input

- Fix: consume all positional arguments as input by joining all positional arguments with spaces instead of using only the last argument, allowing commands to process entire phrases correctly

## v1.4.388 (2026-01-23)

### PR [#1957](https://github.com/danielmiessler/Fabric/pull/1957) by [ksylvan](https://github.com/ksylvan): Add Novita AI as a new OpenAI-compatible provider

- Add Novita AI as a new OpenAI-compatible provider
- Add Novita AI provider configuration with API endpoint
- Update README to include Novita AI in supported providers list
- Configure Novita AI to use OpenAI-compatible interface

## v1.4.387 (2026-01-21)

### PR [#1950](https://github.com/danielmiessler/Fabric/pull/1950) by [cleong14](https://github.com/cleong14): feat: add suggest_gt_command pattern

- Added new suggest_gt_command pattern that suggests Gas Town (gt) CLI commands based on natural language descriptions
- Covers 85+ commands across work management, agents, communication, services, diagnostics, and recovery
- Includes pipe-friendly output with raw command on first line for easy command extraction
- Integrates with Gas Town CLI tool for enhanced command-line workflow automation

### PR [#1951](https://github.com/danielmiessler/Fabric/pull/1951) by [ksylvan](https://github.com/ksylvan): Add `extract_wisdom_with_attribution` pattern for speaker-attributed quotes

- Add new `extract_wisdom_with_attribution` pattern that extends `extract_wisdom` functionality with speaker attribution capabilities
- Create comprehensive documentation including README and system.md files for the new pattern
- Update pattern_explanations.md with detailed entry for the new pattern
- Add pattern to suggest_pattern category lists for improved discoverability
- Update pattern metadata files including pattern_descriptions.json and pattern_extracts.json with relevant tags and content

### PR [#1952](https://github.com/danielmiessler/Fabric/pull/1952) by [ksylvan](https://github.com/ksylvan): Fix: using attachments with Anthropic models

- Feat: add multi-content support for images and PDFs in Anthropic client
- Update toMessages to handle multi-content messages with text and attachments
- Add contentBlocksFromMessage to convert message parts to Anthropic blocks
- Implement support for image URLs including data URLs and base64 images
- Add PDF attachment handling via data URLs and URL-based PDFs

## v1.4.386 (2026-01-21)

### PR [#1945](https://github.com/danielmiessler/Fabric/pull/1945) by [ksylvan](https://github.com/ksylvan): feat: Add Spotify API integration for podcast metadata retrieval

- Add Spotify metadata retrieval via --spotify flag
- Add Spotify plugin with OAuth token handling and metadata
- Wire --spotify flag into CLI processing and output
- Register Spotify in plugin setup, env, and registry
- Update shell completions to include --spotify option

## v1.4.385 (2026-01-20)

### PR [#1947](https://github.com/danielmiessler/Fabric/pull/1947) by [cleong14](https://github.com/cleong14): feat(patterns): add extract_bd_ideas pattern

- Added extract_bd_ideas pattern that extracts actionable ideas from content and transforms them into well-structured bd issue tracker commands
- Implemented identification system for tasks, problems, ideas, improvements, bugs, and features
- Added actionability evaluation and appropriate scoping functionality
- Integrated priority assignment system (P0-P4) with relevant labels
- Created ready-to-execute bd create commands output format

### PR [#1948](https://github.com/danielmiessler/Fabric/pull/1948) by [cleong14](https://github.com/cleong14): feat(patterns): add create_bd_issue pattern

- Added create_bd_issue pattern that transforms natural language issue descriptions into optimal bd (Beads) issue tracker commands
- Implemented comprehensive bd create flag reference for better command generation
- Added intelligent type detection system that automatically categorizes issues as bug, feature, task, epic, or chore
- Included priority assessment capability that assigns P0-P4 priority levels based on urgency signals in descriptions
- Integrated smart label selection feature that automatically chooses 1-4 relevant labels for each issue

### PR [#1949](https://github.com/danielmiessler/Fabric/pull/1949) by [ksylvan](https://github.com/ksylvan): Fix #1931 - Image Generation Feature should warn if the model is not capable of Image Generation

- Add image generation compatibility warnings for unsupported models
- Add warning to stderr when using incompatible models with image generation
- Add GPT-5, GPT-5-nano, and GPT-5.2 to supported image generation models
- Create `checkImageGenerationCompatibility` function in OpenAI plugin
- Add comprehensive tests for image generation compatibility warnings

## v1.4.384 (2026-01-19)

### PR [#1944](https://github.com/danielmiessler/Fabric/pull/1944) by [ksylvan](https://github.com/ksylvan): Add Infermatic AI Provider Support

- Add Infermatic provider to ProviderMap as part of Phase 1 implementation for issue #1033
- Add test coverage for the Infermatic AI provider in TestCreateClient to verify provider exists and creates valid client
- Replace go-git status API with native `git status --porcelain` command to fix worktree compatibility issues
- Simplify `IsWorkingDirectoryClean` and `GetStatusDetails` functions to use CLI output parsing instead of go-git library
- Use native `git rev-parse HEAD` to get commit hash after commit and remove unused imports from walker.go

## v1.4.383 (2026-01-18)

### PR [#1943](https://github.com/danielmiessler/Fabric/pull/1943) by [ksylvan](https://github.com/ksylvan): fix: Ollama server now respects the default context window

- Fix: Ollama server now respects the default context window instead of using hardcoded 2048 tokens
- Add parseOllamaNumCtx() function with type-safe extraction supporting 6 numeric types and platform-aware integer overflow protection
- Extract num_ctx from client request options and add ModelContextLength field to ChatRequest struct
- Implement DoS protection via 1,000,000 token maximum limit with sanitized error messages
- Add comprehensive unit tests for parseOllamaNumCtx function covering edge cases including overflow and invalid types

## v1.4.382 (2026-01-17)

### PR [#1941](https://github.com/danielmiessler/Fabric/pull/1941) by [ksylvan](https://github.com/ksylvan): Add `greybeard_secure_prompt_engineer` to metadata, also remove duplicate json data file

- Add greybeard_secure_prompt_engineer pattern to metadata (pattern explanations and json index)
- Refactor build process to use npm hooks for copying JSON files instead of manual copying
- Update .gitignore to exclude generated data and tmp directories
- Modify suggest_pattern categories to include new security pattern
- Delete redundant web static data file and rely on build hooks

## v1.4.381 (2026-01-17)

### PR [#1940](https://github.com/danielmiessler/Fabric/pull/1940) by [ksylvan](https://github.com/ksylvan): Rewrite Ollama chat handler to support proper streaming responses

- Refactor Ollama chat handler to support proper streaming responses with real-time SSE data parsing
- Replace single-read body parsing with streaming bufio.Scanner approach and implement writeOllamaResponse helper function
- Add comprehensive error handling improvements including proper HTTP error responses instead of log.Fatal to prevent server crashes
- Fix upstream error handling to return stringified error payloads and validate Fabric chat URL hosts
- Implement proper request context propagation and align duration fields to int64 nanosecond precision for consistency

## v1.4.380 (2026-01-16)

### PR [#1936](https://github.com/danielmiessler/Fabric/pull/1936) by [ksylvan](https://github.com/ksylvan): New Vendor: Microsoft Copilot

- Add Microsoft 365 Copilot integration as a new AI vendor with OAuth2 authentication for delegated user permissions
- Enable querying of Microsoft 365 data including emails, documents, and chats with both synchronous and streaming response support
- Provide comprehensive setup instructions for Azure AD app registration and detail licensing, technical, and permission requirements
- Add troubleshooting steps for common authentication and API errors with current API limitations documentation
- Fix SendStream interface to use domain.StreamUpdate instead of chan string to match current Vendor interface requirements

## v1.4.379 (2026-01-15)

### PR [#1935](https://github.com/danielmiessler/Fabric/pull/1935) by [dependabot](https://github.com/apps/dependabot): chore(deps): bump the npm_and_yarn group across 1 directory with 2 updates

- Updated @sveltejs/kit from version 2.21.1 to 2.49.5
- Updated devalue dependency from version 5.3.2 to 5.6.2

## v1.4.378 (2026-01-14)

### PR [#1933](https://github.com/danielmiessler/Fabric/pull/1933) by [ksylvan](https://github.com/ksylvan): Add DigitalOcean Gradient AI support

- Feat: add DigitalOcean Gradient AI Agents as a new vendor
- Add DigitalOcean as a new AI provider in plugin registry
- Implement DigitalOcean client with OpenAI-compatible inference endpoint
- Support model access key authentication for inference requests
- Add optional control plane token for model discovery

### Direct commits

- Chore: Update README with a links to other docs

## v1.4.377 (2026-01-12)

### PR [#1929](https://github.com/danielmiessler/Fabric/pull/1929) by [ksylvan](https://github.com/ksylvan): Add Mammouth as new OpenAI-compatible AI provider

- Feat: add Mammouth as new OpenAI-compatible AI provider
- Add Mammouth provider configuration with API base URL
- Configure Mammouth to use standard OpenAI-compatible interface
- Disable Responses API implementation for Mammouth provider
- Add "Mammouth" to VSCode spell check dictionary

## v1.4.376 (2026-01-12)

### PR [#1928](https://github.com/danielmiessler/Fabric/pull/1928) by [ksylvan](https://github.com/ksylvan): Eliminate repetitive boilerplate across eight vendor implementations

- Refactor: add NewVendorPluginBase factory function to reduce duplication
- Update 8 vendor files (anthropic, bedrock, gemini, lmstudio, ollama, openai, perplexity, vertexai) to use the factory function
- Add 3 test cases for the new factory function
- Add centralized factory function for AI vendor plugin initialization
- Chore: exempt json files from VSCode format-on-save

### Direct commits

- Docs: Add GitHub sponsor section to README
I spend hundreds of hours a year on open source. If you'd like to help support this project, you can sponsor me here.
Co-Authored-By: Claude Opus 4.5 <noreply@anthropic.com>

## v1.4.375 (2026-01-08)

### PR [#1925](https://github.com/danielmiessler/Fabric/pull/1925) by [ksylvan](https://github.com/ksylvan): docs: update README to document new AI providers and features

- Docs: update README to document new AI providers and features
- List supported native and OpenAI-compatible AI provider integrations
- Document dry run mode for previewing prompt construction
- Explain Ollama compatibility mode for exposing API endpoints
- Detail available prompt strategies like chain-of-thought and reflexion

### PR [#1926](https://github.com/danielmiessler/Fabric/pull/1926) by [henricook](https://github.com/henricook) and [ksylvan](https://github.com/ksylvan): feat(vertexai): add dynamic model listing and multi-model support

- Dynamic model listing from Vertex AI Model Garden API
- Support for both Gemini (genai SDK) and Claude (Anthropic SDK) models
- Curated Gemini model list with web search support for Gemini models
- Thinking/extended thinking support for Gemini
- TopP parameter support for Claude models

## v1.4.374 (2026-01-05)

### PR [#1924](https://github.com/danielmiessler/Fabric/pull/1924) by [ksylvan](https://github.com/ksylvan): Rename `code_helper` to `code2context` across documentation and CLI

- Rename `code_helper` command to `code2context` throughout codebase
- Update README.md table of contents and references
- Update installation instructions with new binary name
- Update all usage examples in main.go help text
- Update create_coding_feature pattern documentation

## v1.4.373 (2026-01-04)

### PR [#1914](https://github.com/danielmiessler/Fabric/pull/1914) by [majiayu000](https://github.com/majiayu000): feat(code_helper): add stdin support for piping file lists

- Added stdin support for piping file lists to code_helper, enabling commands like `find . -name '*.go' | code_helper "instructions"` and `git ls-files '*.py' | code_helper "Add type hints"`
- Implemented automatic detection of stdin pipe mode with single argument (instructions) support
- Enhanced tool to read file paths from stdin line by line while maintaining backward compatibility with existing directory scanning functionality

### PR [#1915](https://github.com/danielmiessler/Fabric/pull/1915) by [majiayu000](https://github.com/majiayu000): feat: parallelize audio chunk transcription for improved performance

- Parallelize audio chunk transcription using goroutines for improved performance

## v1.4.372 (2026-01-04)

### PR [#1913](https://github.com/danielmiessler/Fabric/pull/1913) by [majiayu000](https://github.com/majiayu000): fix: REST API /chat endpoint doesn't pass 'search' parameter to ChatOptions

- Fix: REST API /chat endpoint now properly passes Search and SearchLocation parameters to ChatOptions

## v1.4.371 (2026-01-04)

### PR [#1923](https://github.com/danielmiessler/Fabric/pull/1923) by [ksylvan](https://github.com/ksylvan): ChangeLog Generation stability

- Fix: improve date parsing and prevent early return when PR numbers exist
- Add SQLite datetime formats to version date parsing logic
- Loop through multiple date formats until one succeeds
- Include SQLite fractional seconds format support
- Prevent early return when version has PR numbers to output

## v1.4.370 (2026-01-04)

### PR [#1921](https://github.com/danielmiessler/Fabric/pull/1921) by [ksylvan](https://github.com/ksylvan): chore: remove redundant `--sync-db` step from changelog workflow

- Remove redundant `--sync-db` step from changelog workflow
- Remove duplicate database sync command from version workflow
- Simplify changelog generation to single process-prs step
- Clean up `heal_person` pattern by removing duplicate content sections
- Remove duplicate IDENTITY, PURPOSE, STEPS, and OUTPUT INSTRUCTIONS from pattern file

## v1.4.369 (2026-01-04)

### PR [#1919](https://github.com/danielmiessler/Fabric/pull/1919) by [ksylvan](https://github.com/ksylvan): Fix the `last_pr_sync` setting during PR incoming processing

- Fix: update `SetLastPRSync` to use version date instead of current time
- Change last_pr_sync to use versionDate instead of time.Now()
- Ensure future runs fetch PRs merged after the version date
- Add clarifying comments explaining the sync timing logic

## v1.4.368 (2026-01-04)

### PR [#1918](https://github.com/danielmiessler/Fabric/pull/1918) by [ksylvan](https://github.com/ksylvan): Maintenance: Fix  ChangeLog Generation during CI/CD

- Refactor CHANGELOG.md entries with improved formatting and conventional commit prefixes
- Consolidate git worktree fixes into single PR #1917 entry
- Reorder PR entries chronologically within version sections
- Add cache metadata update step before staging release changes
- Update changelog database binary with new entry formatting

## v1.4.367 (2026-01-03)

### PR [#1912](https://github.com/danielmiessler/Fabric/pull/1912) by [berniegreen](https://github.com/berniegreen): refactor: implement structured streaming and metadata support

- Feat: add domain types for structured streaming (Phase 1)
- Refactor: update Vendor interface and Chatter for structured streaming (Phase 2)
- Refactor: implement structured streaming in all AI vendors (Phase 3)
- Feat: implement CLI support for metadata display (Phase 4)
- Feat: implement REST API support for metadata streaming (Phase 5)

## v1.4.366 (2026-01-03)

### PR [#1917](https://github.com/danielmiessler/Fabric/pull/1917) by [ksylvan](https://github.com/ksylvan): Fix: generate_changelog now works in Git Work Trees

- Fix: improve git worktree status detection to ignore staged-only files and check worktree status codes instead of using IsClean method
- Fix: use native git CLI for add/commit in worktrees to resolve go-git issues with shared object databases
- Check filesystem existence of staged files to handle worktree scenarios and ignore files staged in main repo that don't exist in worktree
- Update GetStatusDetails to only include worktree-modified files and ignore unmodified and untracked files in clean check
- Allow staged files that exist in worktree to be committed normally and fix 'cannot create empty commit: clean working tree' errors

### PR [#1909](https://github.com/danielmiessler/Fabric/pull/1909) by [copyleftdev](https://github.com/copyleftdev): feat: add greybeard_secure_prompt_engineer pattern

- Feat: add greybeard_secure_prompt_engineer pattern

### Direct commits

- Feat: implement REST API support for metadata streaming (Phase 5)
- Feat: implement CLI support for metadata display (Phase 4)
- Refactor: implement structured streaming in all AI vendors (Phase 3)

## v1.4.365 (2025-12-30)

### PR [#1908](https://github.com/danielmiessler/Fabric/pull/1908) by [rodaddy](https://github.com/rodaddy): feat(ai): add VertexAI provider for Claude models

- Add support for Google Cloud Vertex AI as a provider to access Claude models using Application Default Credentials (ADC)
- Enable routing of Fabric requests through Google Cloud Platform instead of directly to Anthropic for GCP billing
- Support for Claude models (Sonnet 4.5, Opus 4.5, Haiku 4.5, etc.) via Vertex AI with configurable project ID and region
- Implement full streaming and non-streaming request capabilities with complete ai.Vendor interface
- Extract message conversion logic to dedicated `toMessages` helper method with proper role handling and validation

## v1.4.364 (2025-12-28)

### PR [#1907](https://github.com/danielmiessler/Fabric/pull/1907) by [majiayu000](https://github.com/majiayu000): feat(gui): add Session Name support for multi-turn conversations

- Add Session Name support for multi-turn conversations in GUI chat interface, enabling persistent conversations similar to CLI's --session flag
- Extract session UI into dedicated SessionSelector component with proper Select component integration
- Add session message loading functionality when selecting existing sessions
- Fix session input handling to prevent resetting on each keystroke and improve layout with vertical stacking
- Implement proper error handling for session loading and two-way binding with Select component

## v1.4.363 (2025-12-25)

### PR [#1906](https://github.com/danielmiessler/Fabric/pull/1906) by [ksylvan](https://github.com/ksylvan): Code Quality: Optimize HTTP client reuse + simplify error formatting

- Refactor: optimize HTTP client reuse and simplify error formatting
- Simplify error wrapping by removing redundant Sprintf calls in CLI
- Pass HTTP client to FetchModelsDirectly to enable connection reuse
- Store persistent HTTP client instance inside the OpenAI provider struct
- Update compatible AI providers to match the new function signature

## v1.4.362 (2025-12-25)

### PR [#1904](https://github.com/danielmiessler/Fabric/pull/1904) by [majiayu000](https://github.com/majiayu000): fix: resolve WebUI tooltips not rendering due to overflow clipping

- Fix WebUI tooltips not rendering due to overflow clipping by using position: fixed and getBoundingClientRect() for dynamic positioning
- Extract positioning calculations into dedicated `positioning.ts` module for better code organization
- Add reactive tooltip position updates on scroll and resize events for improved user experience
- Improve accessibility with `aria-describedby` attributes and unique IDs for better screen reader support
- Update unit tests to use extracted functions and add test coverage for style formatting function

## v1.4.361 (2025-12-25)

### PR [#1905](https://github.com/danielmiessler/Fabric/pull/1905) by [majiayu000](https://github.com/majiayu000): fix: optimize oversized logo images reducing package size by 93%

- Fix: optimize oversized logo images reducing package size by 93%
- Replace 42MB favicon.png with proper 64x64 PNG (4.7KB)
- Replace 42MB fabric-logo.png with static PNG from first GIF frame (387KB)
- Optimize animated GIF from 42MB to 5.4MB (half resolution, 12fps, 128 colors)
- Chore: incoming 1905 changelog entry

### Direct commits

- Fix: resolve WebUI tooltips not rendering due to overflow clipping

## v1.4.360 (2025-12-23)

### PR [#1903](https://github.com/danielmiessler/Fabric/pull/1903) by [ksylvan](https://github.com/ksylvan): Update project dependencies and core SDK versions

- Chore: update project dependencies and core SDK versions
- Upgrade AWS SDK v2 components to latest stable versions
- Update Ollama library to version 0.13.5 for improvements
- Bump Google API and GenAI dependencies to newer releases
- Refresh Cobra CLI framework and Pflag to latest versions

## v1.4.359 (2025-12-23)

### PR [#1902](https://github.com/danielmiessler/Fabric/pull/1902) by [ksylvan](https://github.com/ksylvan): Code Cleanup and Simplification

- Chore: simplify error formatting and clean up model assignment logic
- Remove redundant fmt.Sprintf calls from error formatting logic
- Simplify model assignment to always use normalized model names
- Remove unused variadic parameter from the VendorsManager Clear method
- Chore: incoming 1902 changelog entry

## v1.4.358 (2025-12-23)

### PR [#1901](https://github.com/danielmiessler/Fabric/pull/1901) by [orbisai0security](https://github.com/orbisai0security): sexurity fix: Ollama update: CVE-2025-63389

- Chore: incoming 1901 changelog entry
- Fix: resolve critical vulnerability CVE-2025-63389

## v1.4.357 (2025-12-22)

### PR [#1897](https://github.com/danielmiessler/Fabric/pull/1897) by [ksylvan](https://github.com/ksylvan): feat: add MiniMax provider support to OpenAI compatible plugin

- Add MiniMax provider support to OpenAI compatible plugin
- Add MiniMax provider configuration to ProviderMap with base URL set to api.minimaxi.com/v1
- Configure MiniMax with ImplementsResponses as false and add test case for provider validation

### Direct commits

- Add v1.4.356 release note highlighting complete internationalization support across 10 languages
- Highlight full setup prompt i18n and intelligent environment variable handling for consistency

## v1.4.356 (2025-12-22)

### PR [#1895](https://github.com/danielmiessler/Fabric/pull/1895) by [ksylvan](https://github.com/ksylvan): Localize setup process and add funding configuration

- Localize setup prompts and error messages across multiple languages for improved user experience
- Add GitHub and Buy Me a Coffee funding configuration to support project development
- Implement helper for localized questions with static environment keys to streamline internationalization
- Update environment variable builder to handle hyphenated plugin names properly
- Replace hardcoded console output with localized i18n translation strings throughout the application

## v1.4.355 (2025-12-20)

### PR [#1890](https://github.com/danielmiessler/Fabric/pull/1890) by [ksylvan](https://github.com/ksylvan): Bundle yt-dlp with fabric in Nix flake, introduce slim variant

- Added bundled yt-dlp with fabric package in Nix flake configuration
- Introduced fabric-slim variant as a lightweight alternative without yt-dlp
- Renamed original fabric package to fabricSlim for better organization
- Created new fabric package as symlinkJoin of fabricSlim and yt-dlp
- Updated default package to point to the bundled fabric version with yt-dlp

## v1.4.354 (2025-12-19)

### PR [#1889](https://github.com/danielmiessler/Fabric/pull/1889) by [ksylvan](https://github.com/ksylvan): docs: Add a YouTube transcript endpoint to the Swagger UI

- Add `/youtube/transcript` POST endpoint to Swagger docs
- Define `YouTubeRequest` schema with URL, language, timestamps fields
- Define `YouTubeResponse` schema with transcript and metadata fields
- Add API security requirement using ApiKeyAuth
- Document 200, 400, and 500 response codes

## v1.4.353 (2025-12-19)

### PR [#1887](https://github.com/danielmiessler/Fabric/pull/1887) by [bvandevliet](https://github.com/bvandevliet): feat: correct video title and added description to yt transcript api response

- Feat: correct video title (instead of id) and added description to yt transcript api response
- Updated API documentation
- Chore: incoming 1887 changelog entry

## v1.4.352 (2025-12-18)

### PR [#1886](https://github.com/danielmiessler/Fabric/pull/1886) by [ksylvan](https://github.com/ksylvan): Enhanced Onboarding and Setup Experience

- User Experience: implement automated first-time setup and improved configuration validation
- Add automated first-time setup for patterns and strategies
- Implement configuration validation to warn about missing required components
- Update setup menu to group plugins into required and optional
- Provide helpful guidance when no patterns are found in listing

### Direct commits

- Chore: update README with new interactive Swagger available in v.1.4.350

## v1.4.351 (2025-12-18)

### PR [#1882](https://github.com/danielmiessler/Fabric/pull/1882) by [bvandevliet](https://github.com/bvandevliet): Added yt-dlp package to docker image

- Added yt-dlp package to docker image.
- Chore: incoming 1882 changelog entry

## v1.4.350 (2025-12-18)

### PR [#1884](https://github.com/danielmiessler/Fabric/pull/1884) by [ksylvan](https://github.com/ksylvan): Implement interactive Swagger API documentation and automated OpenAPI specification generation

- Add Swagger UI at `/swagger/index.html` endpoint
- Generate OpenAPI spec files (JSON and YAML)
- Document chat, patterns, and models endpoints
- Update contributing guide with Swagger annotation instructions
- Add swaggo dependencies to project

### PR [#1880](https://github.com/danielmiessler/Fabric/pull/1880) by [ksylvan](https://github.com/ksylvan): docs: add REST API server section and new endpoint reference

- Add README table-of-contents link for REST API
- Document REST API server startup and capabilities
- Add endpoint overview for chat, patterns, contexts
- Describe sessions management and model listing endpoints
- Provide curl examples for key API workflows

## v1.4.349 (2025-12-16)

### PR [#1877](https://github.com/danielmiessler/Fabric/pull/1877) by [ksylvan](https://github.com/ksylvan): modernize: update GitHub Actions and modernize Go code

- Modernize: update GitHub Actions and modernize Go code with latest stdlib features
- Upgrade GitHub Actions to latest versions (v6, v21)
- Add modernization check step in CI workflow
- Replace strings manipulation with `strings.CutPrefix` and `strings.CutSuffix`
- Replace manual loops with `slices.Contains` for validation

## v1.4.348 (2025-12-16)

### PR [#1876](https://github.com/danielmiessler/Fabric/pull/1876) by [ksylvan](https://github.com/ksylvan): modernize Go code with TypeFor and range loops

- Replace reflect.TypeOf with TypeFor generic syntax for improved type safety
- Convert traditional for loops to range-based iterations for cleaner code
- Simplify reflection usage in CLI flag handling
- Update test loops to use range over integers
- Refactor string processing loops in template plugin

## v1.4.347 (2025-12-16)

### PR [#1875](https://github.com/danielmiessler/Fabric/pull/1875) by [ksylvan](https://github.com/ksylvan): modernize: update benchmarks to use b.Loop and refactor map copying

- Update benchmark loops to use cleaner `b.Loop()` syntax
- Remove unnecessary `b.ResetTimer()` call in token benchmark
- Use `maps.Copy` for merging variables in patterns handler
- Update benchmarks to use b.Loop and refactor map copying

## v1.4.346 (2025-12-16)

### PR [#1874](https://github.com/danielmiessler/Fabric/pull/1874) by [ksylvan](https://github.com/ksylvan): refactor: replace interface{} with any across codebase

- Replace `interface{}` with `any` in slice type declarations
- Update map types from `map[string]interface{}` to `map[string]any`
- Change variadic function parameters to use `...any` instead of `...interface{}`
- Modernize JSON unmarshaling variables to `any` for consistency
- Update struct fields and method signatures to prefer `any` alias

## v1.4.345 (2025-12-15)

### PR [#1870](https://github.com/danielmiessler/Fabric/pull/1870) by [ksylvan](https://github.com/ksylvan): Web UI: upgrade pdfjs and add SSR-safe dynamic PDF worker init

- Upgrade `pdfjs-dist` to v5 with new engine requirement
- Dynamically import PDF.js to avoid SSR import-time crashes
- Configure PDF worker via CDN using runtime PDF.js version
- Update PDF conversion pipeline to use lazy initialization
- Guard chat message localStorage persistence behind browser checks

## v1.4.344 (2025-12-14)

### PR [#1867](https://github.com/danielmiessler/Fabric/pull/1867) by [jaredmontoya](https://github.com/jaredmontoya): chore: update flake

- Chore: update flake
- Merge branch 'main' into update-flake
- Chore: incoming 1867 changelog entry

## v1.4.343 (2025-12-14)

### PR [#1829](https://github.com/danielmiessler/Fabric/pull/1829) by [dependabot[bot]](https://github.com/apps/dependabot): chore(deps): bump js-yaml from 4.1.0 to 4.1.1 in /web in the npm_and_yarn group across 1 directory

- Updated js-yaml dependency from version 4.1.0 to 4.1.1 in the web directory
- Added changelog entry for incoming PR #1829

### Direct commits

- Updated flake configuration

## v1.4.342 (2025-12-13)

### PR [#1866](https://github.com/danielmiessler/Fabric/pull/1866) by [ksylvan](https://github.com/ksylvan): fix: write CLI and streaming errors to stderr

- Fix: write CLI and streaming errors to stderr
- Route CLI execution errors to standard error output
- Print Anthropic stream errors to stderr consistently
- Add os import to support stderr error writes
- Preserve help-output suppression and exit behavior

## v1.4.341 (2025-12-11)

### PR [#1860](https://github.com/danielmiessler/Fabric/pull/1860) by [ksylvan](https://github.com/ksylvan): fix: allow resetting required settings without validation errors

- Fix: allow resetting required settings without validation errors
- Update `Ask` to detect reset command and bypass validation
- Refactor `OnAnswer` to support new `isReset` parameter logic
- Invoke `ConfigureCustom` in `Setup` to avoid redundant re-validation
- Add unit tests ensuring required fields can be reset

## v1.4.340 (2025-12-08)

### PR [#1856](https://github.com/danielmiessler/Fabric/pull/1856) by [ksylvan](https://github.com/ksylvan): Add support for new ClaudeHaiku 4.5 models

- Added support for new ClaudeHaiku 4.5 models in client
- Added `ModelClaudeHaiku4_5` to supported models list
- Added `ModelClaudeHaiku4_5_20251001` to supported models list

## v1.4.339 (2025-12-08)

### PR [#1855](https://github.com/danielmiessler/Fabric/pull/1855) by [ksylvan](https://github.com/ksylvan): feat: add image attachment support for Ollama vision models

- Add multi-modal image support to Ollama client
- Add base64 and io imports for image handling
- Store httpClient separately in Client struct for reuse
- Convert createChatRequest to return error for validation
- Implement convertMessage to handle multi-content chat messages

## v1.4.338 (2025-12-04)

### PR [#1852](https://github.com/danielmiessler/Fabric/pull/1852) by [ksylvan](https://github.com/ksylvan): Add Abacus vendor for ChatLLM models with static model list

- Add static model support and register Abacus provider
- Detect modelsURL starting with 'static:' and route appropriately
- Implement getStaticModels returning curated Abacus model list
- Register Abacus provider with ModelsURL 'static:abacus'
- Extend provider tests to include Abacus existence

## v1.4.337 (2025-12-04)

### PR [#1851](https://github.com/danielmiessler/Fabric/pull/1851) by [ksylvan](https://github.com/ksylvan): Add Z AI provider and glm model support

- Add Z AI provider configuration to ProviderMap
- Include BaseURL for Z AI API endpoint
- Add test case for Z AI provider existence
- Add glm to OpenAI model prefixes list
- Support new Z AI provider in OpenAI compatible plugins

## v1.4.336 (2025-12-01)

### PR [#1848](https://github.com/danielmiessler/Fabric/pull/1848) by [zeddy303](https://github.com/zeddy303): Fix localStorage SSR error in favorites-store

- Fix localStorage SSR error in favorites-store by using SvelteKit's browser constant instead of typeof localStorage check to properly handle server-side rendering and prevent 'localStorage.getItem is not a function' error when running dev server
- Add changelog entry for incoming PR #1848

## v1.4.335 (2025-11-28)

### PR [#1847](https://github.com/danielmiessler/Fabric/pull/1847) by [ksylvan](https://github.com/ksylvan): Improve model name matching for NeedsRaw in Ollama plugin

- Improved model name matching in Ollama plugin by replacing prefix matching with substring matching
- Enhanced Ollama model name detection by enabling substring-based search instead of prefix-only matching
- Added "conceptmap" to VSCode dictionary settings for better development experience
- Fixed typo in README documentation
- Renamed `ollamaPrefixes` variable to `ollamaSearchStrings` for better code clarity

## v1.4.334 (2025-11-26)

### PR [#1845](https://github.com/danielmiessler/Fabric/pull/1845) by [ksylvan](https://github.com/ksylvan): Add Claude Opus 4.5 Support

- Add Claude Opus 4.5 model variants to Anthropic client
- Upgrade anthropic-sdk-go from v1.16.0 to v1.19.0
- Update golang.org/x/crypto from v0.41.0 to v0.45.0
- Upgrade golang.org/x/net from v0.43.0 to v0.47.0
- Bump golang.org/x/text from v0.28.0 to v0.31.0

## v1.4.333 (2025-11-25)

### PR [#1844](https://github.com/danielmiessler/Fabric/pull/1844) by [ksylvan](https://github.com/ksylvan): Correct directory name from `concall_summery` to `concall_summary`

- Fix: correct directory name from `concall_summery` to `concall_summary`
- Rename pattern directory to fix spelling error
- Update suggest_pattern system with concall_summary references
- Add concall_summary to BUSINESS and SUMMARIZE category listings
- Add user documentation for earnings call analysis

### PR [#1833](https://github.com/danielmiessler/Fabric/pull/1833) by [junaid18183](https://github.com/junaid18183): Added concall_summery

- Added concall_summery

## v1.4.332 (2025-11-24)

### PR [#1843](https://github.com/danielmiessler/Fabric/pull/1843) by [ksylvan](https://github.com/ksylvan): Implement case-insensitive vendor and model name matching

- Fix: implement case-insensitive vendor and model name matching across the application
- Add case-insensitive vendor lookup in VendorsManager
- Implement model name normalization in GetChatter method
- Add FilterByVendor method with case-insensitive matching
- Add FindModelNameCaseInsensitive helper for model queries

## v1.4.331 (2025-11-23)

### PR [#1839](https://github.com/danielmiessler/Fabric/pull/1839) by [ksylvan](https://github.com/ksylvan): Add GitHub Models Provider and Refactor Fetching Fallback Logic

- Feat: add GitHub Models provider and refactor model fetching with direct API fallback
- Add GitHub Models to supported OpenAI-compatible providers list
- Implement direct HTTP fallback for non-standard model responses
- Centralize model fetching logic in openai package
- Upgrade openai-go SDK dependency from v1.8.2 to v1.12.0

## v1.4.330 (2025-11-23)

### PR [#1840](https://github.com/danielmiessler/Fabric/pull/1840) by [ZackaryWelch](https://github.com/ZackaryWelch): Replace deprecated bash function in completion script

- Replace deprecated bash function in completion script to use `_comp_get_words` instead of the removed `__get_comp_words_by_ref` function
- Fix compatibility issues with latest bash version 5.2 and newer distributions like Fedora 42+

## v1.4.329 (2025-11-20)

### PR [#1838](https://github.com/danielmiessler/Fabric/pull/1838) by [ksylvan](https://github.com/ksylvan): refactor: implement i18n support for YouTube tool error messages

- Refactor: implement i18n support for YouTube tool error messages
- Replace hardcoded error strings with i18n translation calls
- Add localization keys for YouTube errors to all locale files
- Introduce `extractAndValidateVideoId` helper to reduce code duplication
- Update timestamp parsing logic to handle localized error formats

## v1.4.328 (2025-11-18)

### PR [#1836](https://github.com/danielmiessler/Fabric/pull/1836) by [ksylvan](https://github.com/ksylvan): docs: clarify `--raw` flag behavior for OpenAI and Anthropic providers

- Updated documentation to clarify `--raw` flag behavior across OpenAI and Anthropic providers
- Documented that Anthropic models use smart parameter selection instead of raw flag behavior
- Updated CLI help text and shell completion descriptions for better clarity
- Translated updated flag descriptions to all supported locales
- Removed outdated references to system/user role changes

### Direct commits

- Added concall_summery

## v1.4.327 (2025-11-16)

### PR [#1832](https://github.com/danielmiessler/Fabric/pull/1832) by [ksylvan](https://github.com/ksylvan): Improve channel management in Gemini provider

- Fix: improve channel management in Gemini streaming method
- Add deferred channel close at function start
- Return error immediately instead of breaking loop
- Remove redundant channel close statements from loop
- Ensure channel closes on all exit paths consistently

### PR [#1831](https://github.com/danielmiessler/Fabric/pull/1831) by [ksylvan](https://github.com/ksylvan): Remove `get_youtube_rss` pattern

- Chore: remove `get_youtube_rss` pattern from multiple files
- Remove `get_youtube_rss` from `pattern_explanations.md`
- Delete `get_youtube_rss` entry in `pattern_descriptions.json`
- Delete `get_youtube_rss` entry in `pattern_extracts.json`
- Remove `get_youtube_rss` from `suggest_pattern/system.md`

## v1.4.326 (2025-11-16)

### PR [#1830](https://github.com/danielmiessler/Fabric/pull/1830) by [ksylvan](https://github.com/ksylvan): Ensure final newline in model generated outputs

- Add newline to `CreateOutputFile` if missing and improve tests with `t.Cleanup` for file removal
- Add test for message with trailing newline and introduce `printedStream` flag in `Chatter.Send`
- Print newline if stream printed without trailing newline

### Direct commits

- Add v1.4.322 release with concept maps and introduce WELLNESS category with psychological analysis
- Upgrade to Claude Sonnet 4.5 and add Portuguese language variants with BCP 47 support
- Migrate to `openai-go/azure` SDK for Azure integration
- Update README with recent features and extensions, including new Extensions section navigation
- General repository maintenance and feature documentation updates

## v1.4.325 (2025-11-15)

### PR [#1828](https://github.com/danielmiessler/Fabric/pull/1828) by [ksylvan](https://github.com/ksylvan): Fix empty string detection in chatter and AI clients

- Chore: improve message handling by trimming whitespace in content checks
- Remove default space in `BuildSession` message content
- Trim whitespace in `anthropic` message content check
- Trim whitespace in `gemini` message content check
- Chore: incoming 1828 changelog entry

## v1.4.324 (2025-11-14)

### PR [#1827](https://github.com/danielmiessler/Fabric/pull/1827) by [ksylvan](https://github.com/ksylvan): Make YouTube API key optional in setup

- Made YouTube API key optional during setup process
- Changed API key setup question to be optional rather than required
- Added test coverage for optional API key behavior
- Ensured plugin configuration works without API key
- Added changelog entry for the changes

## v1.4.323 (2025-11-12)

### PR [#1802](https://github.com/danielmiessler/Fabric/pull/1802) by [nickarino](https://github.com/nickarino): fix: improve template extension handling for {{input}} and add examples

- Fix: improve template extension handling for {{input}} and add examples
- Extract InputSentinel constant to shared constants.go file and remove duplicate inputSentinel definitions from template.go and patterns.go
- Create withTestExtension helper function to reduce test code duplication and refactor 3 test functions to use the helper
- Fix shell script to use $@ instead of $- for proper argument quoting
- Add prominent warning at top of Extensions guide with visual indicators and update main README with brief Extensions section

### PR [#1823](https://github.com/danielmiessler/Fabric/pull/1823) by [ksylvan](https://github.com/ksylvan): Add missing patterns and renumber pattern explanations list

- Add `apply_ul_tags` pattern for content categorization
- Add `extract_mcp_servers` pattern for MCP server identification
- Add `generate_code_rules` pattern for AI coding guardrails
- Add `t_check_dunning_kruger` pattern for competence assessment
- Renumber all patterns from 37-226 to 37-230 and insert new patterns at positions 37, 129, 153, 203

## v1.4.322 (2025-11-05)

### PR [#1816](https://github.com/danielmiessler/Fabric/pull/1816) by [ksylvan](https://github.com/ksylvan): Update `anthropic-sdk-go` to v1.16.0 and update models

- Upgrade `anthropic-sdk-go` to version 1.16.0
- Remove outdated model `ModelClaude3_5SonnetLatest`
- Add new model `ModelClaudeSonnet4_5_20250929`
- Include `ModelClaudeSonnet4_5_20250929` in `modelBetas` map

### PR [#1814](https://github.com/danielmiessler/Fabric/pull/1814) by [ksylvan](https://github.com/ksylvan): Add Concept Map in html

- Add `create_conceptmap` for interactive HTML concept maps using Vis.js
- Add `fix_typos` for proofreading and correcting text errors
- Introduce `model_as_sherlock_freud` for psychological modeling and behavior analysis
- Implement `predict_person_actions` for behavioral response predictions
- Add `recommend_yoga_practice` for personalized yoga guidance

## v1.4.321 (2025-11-03)

### PR [#1803](https://github.com/danielmiessler/Fabric/pull/1803) by [dependabot[bot]](https://github.com/apps/dependabot): chore(deps-dev): bump vite from 5.4.20 to 5.4.21 in /web in the npm_and_yarn group across 1 directory

- Bumped vite dependency from 5.4.20 to 5.4.21 in the /web directory

### PR [#1805](https://github.com/danielmiessler/Fabric/pull/1805) by [OmriH-Elister](https://github.com/OmriH-Elister): Added several new patterns

- Added new WELLNESS category with four patterns including yoga practice recommendations
- Introduced psychological analysis patterns: `model_as_sherlock_freud` and `predict_person_actions`
- Added `fix_typos` pattern for proofreading and text corrections
- Updated ANALYSIS and SELF categories to include new wellness-related patterns

### PR [#1808](https://github.com/danielmiessler/Fabric/pull/1808) by [sluosapher](https://github.com/sluosapher): Updated create_newsletter_entry pattern to generate more factual titles

- Updated title generation style for more factual newsletter entries and added output example

## v1.4.320 (2025-10-28)

### PR [#1810](https://github.com/danielmiessler/Fabric/pull/1810) by [tonymet](https://github.com/tonymet): improve subtitle lang, retry, debugging & error handling

- Improve subtitle lang, retry, debugging & error handling

### PR [#1780](https://github.com/danielmiessler/Fabric/pull/1780) by [marcas756](https://github.com/marcas756): feat: add extract_characters pattern

- Add extract_characters pattern for detailed character analysis and identification
- Define character extraction goals with canonical naming and deduplication rules
- Include output schema with formatting guidelines and positive/negative examples

### PR [#1794](https://github.com/danielmiessler/Fabric/pull/1794) by [productStripesAdmin](https://github.com/productStripesAdmin): Enhance web app docs

- Remove duplicate content from main readme and link to web app readme
- Update table of contents with proper nesting and fix minor formatting issues

### Direct commits

- Add new patterns and update title generation style with output examples
- Fix template extension handling for {{input}} and add examples

## v1.4.319 (2025-09-30)

### PR [#1783](https://github.com/danielmiessler/Fabric/pull/1783) by [ksylvan](https://github.com/ksylvan): Update anthropic-sdk-go and add claude-sonnet-4-5

- Updated `anthropic-sdk-go` to version 1.13.0 for improved compatibility and performance
- Added support for `ModelClaudeSonnet4_5` to the list of available AI models

### Direct commits

- Added new `extract_characters` system definition with comprehensive character extraction capabilities
- Implemented canonical naming and deduplication rules for consistent character identification
- Created structured output schema with detailed formatting guidelines and examples
- Established interaction mapping functionality to track character relationships and narrative importance
- Added fallback handling for scenarios where no characters are found in the content

## v1.4.318 (2025-09-24)

### PR [#1779](https://github.com/danielmiessler/Fabric/pull/1779) by [ksylvan](https://github.com/ksylvan): Improve pt-BR Translation - Thanks to @JuracyAmerico

- Fix: improve PT-BR translation naturalness and fluency
- Replace "dos" with "entre" for better preposition usage
- Add definite articles where natural in Portuguese
- Clarify "configurações padrão" instead of just "padrões"
- Keep technical terms visible like "padrões/patterns"

## v1.4.317 (2025-09-21)

### PR [#1778](https://github.com/danielmiessler/Fabric/pull/1778) by [ksylvan](https://github.com/ksylvan): Add Portuguese Language Variants Support (pt-BR and pt-PT)

- Add Brazilian Portuguese (pt-BR) translation file
- Add European Portuguese (pt-PT) translation file
- Implement BCP 47 locale normalization system
- Create fallback chain for language variants
- Add default variant mapping for Portuguese

## v1.4.316 (2025-09-20)

### PR [#1777](https://github.com/danielmiessler/Fabric/pull/1777) by [ksylvan](https://github.com/ksylvan): chore: remove garble installation from release workflow

- Remove garble installation step from release workflow to simplify the build process
- Add comment with GoReleaser config file reference link for better documentation
- Discontinue failed experiment with garble that was intended to improve Windows package manager virus scanning compatibility

## v1.4.315 (2025-09-20)

### PR [#1776](https://github.com/danielmiessler/Fabric/pull/1776) by [ksylvan](https://github.com/ksylvan): Remove garble from the build process for Windows

- Remove garble obfuscation from windows build
- Standardize ldflags across all build targets
- Inject version info during compilation
- Update CI workflow and simplify goreleaser build configuration
- Add changelog database to git tracking

## v1.4.314 (2025-09-17)

### PR [#1774](https://github.com/danielmiessler/Fabric/pull/1774) by [ksylvan](https://github.com/ksylvan): Migrate Azure client to openai-go/azure and default API version

- Migrated Azure client to openai-go/azure and default API version
- Switched Azure OpenAI config to openai-go azure helpers and now require API key and base URL during configuration
- Set default API version to 2024-05-01-preview when unspecified
- Updated dependencies to support azure client and authentication flow
- Removed latest-tag boundary logic from changelog walker and simplified version assignment by matching commit messages directly

### Direct commits

- Fix: One-time fix for CHANGELOG and changelog cache db

## v1.4.313 (2025-09-16)

### PR [#1773](https://github.com/danielmiessler/Fabric/pull/1773) by [ksylvan](https://github.com/ksylvan): Add Garble Obfuscation for Windows Builds

- Add garble obfuscation for Windows builds and fix changelog generation
- Add garble tool installation to release workflow
- Configure garble obfuscation for Windows builds only
- Fix changelog walker to handle unreleased commits
- Implement boundary detection for released vs unreleased commits

## v1.4.312 (2025-09-14)

### PR [#1769](https://github.com/danielmiessler/Fabric/pull/1769) by [ksylvan](https://github.com/ksylvan): Go 1.25.1 Upgrade & Critical SDK Updates

- Upgrade Go from 1.24 to 1.25.1
- Update Anthropic SDK for web fetch tools
- Upgrade AWS Bedrock SDK 12 versions
- Update Azure Core and Identity SDKs
- Fix Nix config for Go version lag

## v1.4.311 (2025-09-13)

### PR [#1767](https://github.com/danielmiessler/Fabric/pull/1767) by [ksylvan](https://github.com/ksylvan): feat(i18n): add de, fr, ja, pt, zh, fa locales; expand tests

- Add DE, FR, JA, PT, ZH, FA i18n locale files
- Expand i18n tests with table-driven multilingual coverage
- Verify 'html_readability_error' translations across all supported languages
- Update README with release notes for added languages
- Insert blank lines between aggregated PR changelog sections

### Direct commits

- Chore: update changelog formatting and sync changelog database

- Add line breaks to improve changelog readability

- Sync changelog database with latest entries
- Clean up whitespace in version sections

- Maintain consistent formatting across entries
- Chore: add spacing between changelog entries for improved readability

- Add blank lines between PR sections

- Update changelog database with  to correspond with CHANGELOG fix.

## v1.4.310 (2025-09-11)

### PR [#1759](https://github.com/danielmiessler/Fabric/pull/1759) by [ksylvan](https://github.com/ksylvan): Add Windows-style Flag Support for Language Detection

- Feat: add Windows-style forward slash flag support to CLI argument parser
- Add runtime OS detection for Windows platform
- Support `/flag` syntax for Windows command line
- Handle Windows colon delimiter `/flag:value` format
- Maintain backward compatibility with Unix-style flags

### PR [#1762](https://github.com/danielmiessler/Fabric/pull/1762) by [OmriH-Elister](https://github.com/OmriH-Elister): New pattern for writing interaction between two characters

- Feat: add new pattern that creates story simulating interaction between two people
- Chore: add `create_story_about_people_interaction` pattern for persona analysis
- Add `create_story_about_people_interaction` pattern description
- Include pattern in `ANALYSIS` and `WRITING` categories
- Update `suggest_pattern` system and user documentation

### Direct commits

- Chore: update alias creation to use consistent naming

- Remove redundant prefix from `pattern_name` variable

- Add `alias_name` variable for consistent alias creation
- Update alias command to use `alias_name`

- Modify PowerShell function to use `aliasName`
- Docs: add optional prefix support for fabric pattern aliases via FABRIC_ALIAS_PREFIX env var

- Add FABRIC_ALIAS_PREFIX environment variable support

- Update bash/zsh alias generation with prefix
- Update PowerShell alias generation with prefix

- Improve readability of alias setup instructions
- Enable custom prefixing for pattern commands

- Maintain backward compatibility without prefix

## v1.4.309 (2025-09-09)

### PR [#1756](https://github.com/danielmiessler/Fabric/pull/1756) by [ksylvan](https://github.com/ksylvan): Add Internationalization Support with Custom Help System

- Add comprehensive internationalization support with English and Spanish locales
- Replace hardcoded strings with i18n.T translations and add en and es JSON locale files
- Implement custom translated help system with language detection from CLI args
- Add locale download capability and localize error messages throughout codebase
- Support TTS and notification translations

## v1.4.308 (2025-09-05)

### PR [#1755](https://github.com/danielmiessler/Fabric/pull/1755) by [ksylvan](https://github.com/ksylvan): Add i18n Support for Multi-Language Fabric Experience

- Add Spanish localization support with i18n
- Create contexts and sessions tutorial documentation
- Fix broken Warp sponsorship image URL
- Remove solve_with_cot pattern from codebase
- Update pattern descriptions and explanations

### Direct commits

- Update Warp sponsor section with proper formatting

- Replace with correct div structure and styling
- Use proper Warp image URL from brand assets

- Add "Special thanks to:" text and platform availability
- Maintains proper spacing and alignment
- Fix unclosed div tag in README causing display issues

- Close the main div container properly after fabric screenshot
- Fix HTML structure that was causing repetitive content display

- Ensure proper markdown rendering on GitHub
🤖 Generated with [Claude Code](<https://claude.ai/code)>
Co-Authored-By: Claude <noreply@anthropic.com>

- Update Warp sponsor section with new banner and branding

- Replace old banner with new warp-banner-light.png image
- Update styling to use modern p tags with proper centering

- Maintain existing go.warp.dev/fabric redirect URL
- Add descriptive alt text and emphasis text for accessibility
🤖 Generated with [Claude Code](<https://claude.ai/code)>
Co-Authored-By: Claude <noreply@anthropic.com>

## v1.4.307 (2025-09-01)

### PR [#1745](https://github.com/danielmiessler/Fabric/pull/1745) by [ksylvan](https://github.com/ksylvan): Fabric Installation Improvements and Automated Release Updates

- Streamlined install process with one-line installer scripts and updated documentation
- Added bash installer script for Unix systems
- Added PowerShell installer script for Windows
- Created installer documentation with usage examples
- Simplified README installation with one-line installers

## v1.4.306 (2025-09-01)

### PR [#1742](https://github.com/danielmiessler/Fabric/pull/1742) by [ksylvan](https://github.com/ksylvan): Documentation and Pattern Updates

- Add winget installation method for Windows users
- Include Docker Hub and GHCR image references with docker run examples
- Remove deprecated PowerShell download link and unused show_fabric_options_markmap pattern
- Update suggest_pattern with new AI patterns
- Add personal development patterns for storytelling

## v1.4.305 (2025-08-31)

### PR [#1741](https://github.com/danielmiessler/Fabric/pull/1741) by [ksylvan](https://github.com/ksylvan): CI: Fix Release Description Update

- Fix: update release workflow to support manual dispatch with custom tag
- Support custom tag from client payload in workflow
- Fallback to github.ref_name when no custom tag provided
- Enable manual release triggers with specified tag parameter

## v1.4.304 (2025-08-31)

### PR [#1740](https://github.com/danielmiessler/Fabric/pull/1740) by [ksylvan](https://github.com/ksylvan): Restore our custom Changelog Updates in GitHub Actions

- Add changelog generation step to GitHub release workflow
- Create updateReleaseForRepo helper method for release updates
- Add fork detection logic in UpdateReleaseDescription method
- Implement upstream repository release update for forks
- Enhance error handling with detailed repository context

## v1.4.303 (2025-08-28)

### PR [#1736](https://github.com/danielmiessler/Fabric/pull/1736) by [tonymet](https://github.com/tonymet): Winget Publishing and GoReleaser

- Added GoReleaser support for improved package distribution
- Winget and Docker publishing moved to ksylvan/fabric-packager GitHub repo
- Hardened release pipeline by gating workflows to upstream owner only
- Migrated from custom tokens to built-in GITHUB_TOKEN for enhanced security
- Removed docker-publish-on-tag workflow to reduce duplication and complexity
- Added ARM binary release support with updated documentation

## v1.4.302 (2025-08-28)

### PR [#1737](https://github.com/danielmiessler/Fabric/pull/1737) by [ksylvan](https://github.com/ksylvan) and [OmriH-Elister](https://github.com/OmriH-Elister): Add New Psychological Analysis Patterns + devalue version bump

- Add create_story_about_person system pattern with narrative workflow
- Add heal_person system pattern for compassionate healing plans
- Update pattern_explanations to register new patterns and renumber indices
- Extend pattern_descriptions with entries, tags, and concise descriptions
- Bump devalue dependency from 5.1.1 to 5.3.2

## v1.4.301 (2025-08-28)

### PR [#1735](https://github.com/danielmiessler/Fabric/pull/1735) by [ksylvan](https://github.com/ksylvan): Fix Docker Build Path Configuration

- Fix: update Docker workflow to use specific Dockerfile and monitor markdown file changes
- Add explicit Dockerfile path to Docker build action
- Remove markdown files from workflow paths-ignore filter
- Enable CI triggers for documentation file changes
- Specify Docker build context with custom file location

## v1.4.300 (2025-08-28)

### PR [#1732](https://github.com/danielmiessler/Fabric/pull/1732) by [ksylvan](https://github.com/ksylvan): CI Infra: Changelog Generation Tool + Docker Image Pubishing

- Add GitHub Actions workflow to publish Docker images on tags
- Build multi-arch images with Buildx and QEMU across amd64, arm64
- Tag images using semver; push to GHCR and Docker Hub
- Gate patterns workflow steps on detected changes instead of failing
- Auto-detect GitHub owner and repo from git remote URL

## v1.4.299 (2025-08-27)

### PR [#1731](https://github.com/danielmiessler/Fabric/pull/1731) by [ksylvan](https://github.com/ksylvan): chore: upgrade ollama dependency from v0.9.0 to v0.11.7

- Updated ollama package from version 0.9.0 to 0.11.7
- Fixed 8 security vulnerabilities including 5 high-severity CVEs that could cause denial of service attacks
- Patched Ollama server vulnerabilities related to division by zero errors and memory exhaustion
- Resolved security flaws that allowed malicious GGUF model file uploads to crash the server
- Enhanced system stability and security posture through comprehensive dependency upgrade

## v1.4.298 (2025-08-27)

### PR [#1730](https://github.com/danielmiessler/Fabric/pull/1730) by [ksylvan](https://github.com/ksylvan): Modernize Dockerfile with Best Practices Implementation

- Remove docker-test framework and simplify production docker setup by eliminating complex testing infrastructure
- Delete entire docker-test directory including test runner scripts and environment configuration files
- Implement multi-stage build optimization in production Dockerfile to improve build efficiency
- Remove docker-compose.yml and start-docker.sh helper scripts to streamline container workflow
- Update README documentation with cleaner Docker usage instructions and reduced image size benefits

## v1.4.297 (2025-08-26)

### PR [#1729](https://github.com/danielmiessler/Fabric/pull/1729) by [ksylvan](https://github.com/ksylvan): Add GitHub Community Health Documents

- Add CODE_OF_CONDUCT defining respectful, collaborative community behavior
- Add CONTRIBUTING with setup, testing, PR, changelog requirements
- Add SECURITY policy with reporting process and response timelines
- Add SUPPORT guide for bugs, features, discussions, expectations
- Add docs README indexing guides, quick starts, contributor essentials

## v1.4.296 (2025-08-26)

### PR [#1728](https://github.com/danielmiessler/Fabric/pull/1728) by [ksylvan](https://github.com/ksylvan): Refactor Logging System to Use Centralized Debug Logger

- Replace fmt.Fprintf/os.Stderr with centralized debuglog.Log across CLI and add unconditional Log function for important messages
- Improve OAuth flow messaging and token refresh diagnostics with better error handling
- Update tests to capture debuglog output via SetOutput for better test coverage
- Convert Perplexity streaming errors to unified debug logging and emit file write notifications through debuglog
- Standardize extension registry warnings and announce large audio processing steps via centralized logger

## v1.4.295 (2025-08-24)

### PR [#1727](https://github.com/danielmiessler/Fabric/pull/1727) by [ksylvan](https://github.com/ksylvan): Standardize Anthropic Beta Failure Logging

- Refactor: route Anthropic beta failure logs through internal debug logger
- Replace fmt.Fprintf stderr with debuglog.Debug for beta failures
- Import internal log package and remove os dependency
- Standardize logging level to debuglog.Basic for beta errors
- Preserve fallback stream behavior when beta features fail

## v1.4.294 (2025-08-20)

### PR [#1723](https://github.com/danielmiessler/Fabric/pull/1723) by [ksylvan](https://github.com/ksylvan): docs: update README with Venice AI provider and Windows install script

- Add Venice AI provider configuration with API endpoint
- Document Venice AI as privacy-first open-source provider
- Include PowerShell installation script for Windows users
- Add debug levels section to table of contents
- Update recent major features with v1.4.294 release notes

## v1.4.293 (2025-08-19)

### PR [#1718](https://github.com/danielmiessler/Fabric/pull/1718) by [ksylvan](https://github.com/ksylvan): Implement Configurable Debug Logging Levels

- Add --debug flag controlling runtime logging verbosity levels
- Introduce internal/log package with Off, Basic, Detailed, Trace
- Replace ad-hoc Debugf and globals with centralized debug logger
- Wire debug level during early CLI argument parsing
- Add bash, zsh, fish completions for --debug levels

## v1.4.292 (2025-08-18)

### PR [#1717](https://github.com/danielmiessler/Fabric/pull/1717) by [ksylvan](https://github.com/ksylvan): Highlight default vendor/model in model listing

- Update PrintWithVendor signature to accept default vendor and model
- Mark default vendor/model with asterisk in non-shell output
- Compare vendor and model case-insensitively when marking
- Pass registry defaults to PrintWithVendor from CLI
- Add test ensuring default selection appears with asterisk

### Direct commits

- Docs: update version number in README updates section from v1.4.290 to v1.4.291

## v1.4.291 (2025-08-18)

### PR [#1715](https://github.com/danielmiessler/Fabric/pull/1715) by [ksylvan](https://github.com/ksylvan): feat: add speech-to-text via OpenAI with transcription flags and comp…

- Add --transcribe-file flag to transcribe audio or video
- Add --transcribe-model flag with model listing and completion
- Add --split-media-file flag to chunk files over 25MB
- Implement OpenAI transcription using Whisper and GPT-4o Transcribe
- Integrate transcription pipeline into CLI before readability processing

## v1.4.290 (2025-08-17)

### PR [#1714](https://github.com/danielmiessler/Fabric/pull/1714) by [ksylvan](https://github.com/ksylvan): feat: add per-pattern model mapping support via environment variables

- Add per-pattern model mapping support via environment variables
- Implement environment variable lookup for pattern-specific models
- Support vendor|model format in environment variable specification
- Enable shell startup file configuration for patterns
- Transform pattern names to uppercase environment variable format

## v1.4.289 (2025-08-16)

### PR [#1710](https://github.com/danielmiessler/Fabric/pull/1710) by [ksylvan](https://github.com/ksylvan): feat: add --no-variable-replacement flag to disable pattern variable …

- Add --no-variable-replacement flag to disable pattern variable substitution
- Introduce CLI flag to skip pattern variable replacement and wire it into domain request and session builder
- Provide PatternsEntity.GetWithoutVariables for input-only pattern processing support
- Refactor patterns code into reusable load and apply helpers
- Update bash, zsh, fish completions with new flag and document in README and CLI help output

## v1.4.288 (2025-08-16)

### PR [#1709](https://github.com/danielmiessler/Fabric/pull/1709) by [ksylvan](https://github.com/ksylvan): Enhanced YouTube Subtitle Language Fallback Handling

- Fix: improve YouTube subtitle language fallback handling in yt-dlp integration
- Fix typo "Gemmini" to "Gemini" in README
- Add "kballard" and "shellquote" to VSCode dictionary
- Add "YTDLP" to VSCode spell checker
- Enhance subtitle language options with fallback variants

## v1.4.287 (2025-08-14)

### PR [#1706](https://github.com/danielmiessler/Fabric/pull/1706) by [ksylvan](https://github.com/ksylvan): Gemini Thinking Support and README (New Features) automation

- Add comprehensive "Recent Major Features" section to README
- Introduce new readme_updates Python script for automation
- Enable Gemini thinking configuration with token budgets
- Update CLI help text for Gemini thinking support
- Add comprehensive test coverage for Gemini thinking

## v1.4.286 (2025-08-14)

### PR [#1700](https://github.com/danielmiessler/Fabric/pull/1700) by [ksylvan](https://github.com/ksylvan): Introduce Thinking Config Across Anthropic and OpenAI Providers

- Add --thinking CLI flag for configurable reasoning levels across providers
- Implement Anthropic ThinkingConfig with standardized budgets and tokens
- Map OpenAI reasoning effort from thinking levels
- Show thinking level in dry-run formatted options
- Overhaul suggest_pattern docs with categories, workflows, usage examples

## v1.4.285 (2025-08-13)

### PR [#1698](https://github.com/danielmiessler/Fabric/pull/1698) by [ksylvan](https://github.com/ksylvan): Enable One Million Token Context Beta Feature for Sonnet-4

- Chore: upgrade anthropic-sdk-go to v1.9.1 and add beta feature support for context-1m
- Add modelBetas map for beta feature configuration
- Implement context-1m-2025-08-07 beta for Claude Sonnet 4
- Add beta header support with fallback handling
- Preserve existing beta headers in OAuth transport

## v1.4.284 (2025-08-12)

### PR [#1695](https://github.com/danielmiessler/Fabric/pull/1695) by [ksylvan](https://github.com/ksylvan): Introduce One-Liner Curl Install for Completions

- Add one-liner curl install method for shell completions without requiring repository cloning
- Support downloading completions when files are missing locally with dry-run option for previewing changes
- Enable custom download source via environment variable and create temporary directory for downloaded completion files
- Add automatic cleanup of temporary files and validate downloaded files are non-empty and not HTML
- Improve error handling and standardize logging by routing informational messages to stderr to avoid stdout pollution

## v1.4.283 (2025-08-12)

### PR [#1692](https://github.com/danielmiessler/Fabric/pull/1692) by [ksylvan](https://github.com/ksylvan): Add Vendor Selection Support for Models

- Add -V/--vendor flag to specify model vendor
- Implement vendor-aware model resolution and availability validation
- Warn on ambiguous models; suggest --vendor to disambiguate
- Update bash, zsh, fish completions with vendor suggestions
- Extend --listmodels to print vendor|model when interactive

## v1.4.282 (2025-08-11)

### PR [#1689](https://github.com/danielmiessler/Fabric/pull/1689) by [ksylvan](https://github.com/ksylvan): Enhanced Shell Completions for Fabric CLI Binaries

- Add 'fabric-ai' alias support across all shell completions
- Use invoked command name for dynamic completion list queries
- Refactor fish completions into reusable registrar for multiple commands
- Update Bash completion to reference executable via COMP_WORDS[0]
- Install completions automatically with new cross-shell setup script

## v1.4.281 (2025-08-11)

### PR [#1687](https://github.com/danielmiessler/Fabric/pull/1687) by [ksylvan](https://github.com/ksylvan): Add Web Search Tool Support for Gemini Models

- Enable Gemini models to use web search tool with --search flag
- Add validation for search-location timezone and language code formats
- Normalize language codes from underscores to hyphenated form
- Append deduplicated web citations under standardized Sources section
- Improve robustness for nil candidates and content parts

## v1.4.280 (2025-08-10)

### PR [#1686](https://github.com/danielmiessler/Fabric/pull/1686) by [ksylvan](https://github.com/ksylvan): Prevent duplicate text output in OpenAI streaming responses

- Fix: prevent duplicate text output in OpenAI streaming responses
- Skip processing of ResponseOutputTextDone events
- Prevent doubled text in stream output
- Add clarifying comment about API behavior
- Maintain delta chunk streaming functionality

## v1.4.279 (2025-08-10)

### PR [#1685](https://github.com/danielmiessler/Fabric/pull/1685) by [ksylvan](https://github.com/ksylvan): Fix Gemini Role Mapping for API Compatibility

- Fix Gemini role mapping to ensure proper API compatibility by converting chat roles to Gemini's user/model format
- Map assistant role to model role per Gemini API constraints
- Map system, developer, function, and tool roles to user role for proper handling
- Default unrecognized roles to user role to preserve instruction context
- Add comprehensive unit tests to validate convertMessages role mapping logic

## v1.4.278 (2025-08-09)

### PR [#1681](https://github.com/danielmiessler/Fabric/pull/1681) by [ksylvan](https://github.com/ksylvan): Enhance YouTube Support with Custom yt-dlp Arguments

- Add `--yt-dlp-args` flag for custom YouTube downloader options with advanced control capabilities
- Implement smart subtitle language fallback system when requested locale is unavailable
- Add fallback logic for YouTube subtitle language detection with auto-detection of downloaded languages
- Replace custom argument parser with shellquote and precompile regexes for improved performance and safety

## v1.4.277 (2025-08-08)

### PR [#1679](https://github.com/danielmiessler/Fabric/pull/1679) by [ksylvan](https://github.com/ksylvan): Add cross-platform desktop notifications to Fabric CLI

- Add cross-platform desktop notifications with secure custom commands
- Integrate notification sending into chat processing workflow
- Add --notification and --notification-command CLI flags and help
- Provide cross-platform providers: macOS, Linux, Windows with fallbacks
- Escape shell metacharacters to prevent injection vulnerabilities

## v1.4.276 (2025-08-08)

### Direct commits

- Ci: add write permissions to update_release_notes job

- Add contents write permission to release notes job

- Enable GitHub Actions to modify repository contents
- Fix potential permission issues during release process

## v1.4.275 (2025-08-07)

### PR [#1676](https://github.com/danielmiessler/Fabric/pull/1676) by [ksylvan](https://github.com/ksylvan): Refactor authentication to support GITHUB_TOKEN and GH_TOKEN

- Refactor: centralize GitHub token retrieval logic into utility function
- Support both GITHUB_TOKEN and GH_TOKEN environment variables with fallback handling
- Add new util/token.go file for centralized token handling across the application
- Update walker.go and main.go to use the new centralized token utility function
- Feat: add 'gpt-5' to raw-mode models in OpenAI client to bypass structured chat message formatting

## v1.4.274 (2025-08-07)

### PR [#1673](https://github.com/danielmiessler/Fabric/pull/1673) by [ksylvan](https://github.com/ksylvan): Add Support for Claude Opus 4.1 Model

- Add Claude Opus 4.1 model support
- Upgrade anthropic-sdk-go from v1.4.0 to v1.7.0
- Fix temperature/topP parameter conflict for models
- Refactor release workflow to use shared version job and simplify OS handling
- Improve chat parameter defaults handling with domain constants

## v1.4.273 (2025-08-05)

### Direct commits

- Remove redundant words from codebase
- Fix typos in t_ patterns

## v1.4.272 (2025-07-28)

### PR [#1658](https://github.com/danielmiessler/Fabric/pull/1658) by [ksylvan](https://github.com/ksylvan): Update Release Process for Data Consistency

- Add database sync before generating changelog in release workflow
- Ensure changelog generation includes latest database updates
- Update changelog cache database

## v1.4.271 (2025-07-28)

### PR [#1657](https://github.com/danielmiessler/Fabric/pull/1657) by [ksylvan](https://github.com/ksylvan): Add GitHub Release Description Update Feature

- Add GitHub release description update via `--release` flag
- Implement `ReleaseManager` for managing release descriptions
- Create `release.go` for handling release updates
- Update `release.yml` to run changelog generation
- Enable AI summary updates for GitHub releases

## v1.4.270 (2025-07-27)

### PR [#1654](https://github.com/danielmiessler/Fabric/pull/1654) by [ksylvan](https://github.com/ksylvan): Refine Output File Handling for Safety

- Fix: prevent file overwrite and improve output messaging in CreateOutputFile
- Add file existence check before creating output file
- Return error if target file already exists
- Change success message to write to stderr
- Update message format with brackets for clarity

## v1.4.269 (2025-07-26)

### PR [#1653](https://github.com/danielmiessler/Fabric/pull/1653) by [ksylvan](https://github.com/ksylvan): docs: update Gemini TTS model references to gemini-2.5-flash-preview-tts

- Updated Gemini TTS model references from gemini-2.0-flash-tts to gemini-2.5-flash-preview-tts throughout documentation
- Modified documentation examples to use the new gemini-2.5-flash-preview-tts model
- Updated voice selection example commands in Gemini-TTS.md
- Revised CLI help text example commands to reflect model changes
- Updated changelog database binary file

## v1.4.268 (2025-07-26)

### PR [#1652](https://github.com/danielmiessler/Fabric/pull/1652) by [ksylvan](https://github.com/ksylvan): Implement Voice Selection for Gemini Text-to-Speech

- Feat: add Gemini TTS voice selection and listing functionality
- Add `--voice` flag for TTS voice selection
- Add `--list-gemini-voices` command for voice discovery
- Implement voice validation for Gemini TTS models
- Update shell completions for voice options

## v1.4.267 (2025-07-26)

### PR [#1650](https://github.com/danielmiessler/Fabric/pull/1650) by [ksylvan](https://github.com/ksylvan): Update Gemini Plugin to New SDK with TTS Support

- Update Gemini SDK to new genai library and add TTS audio output support
- Replace deprecated generative-ai-go with google.golang.org/genai library
- Add TTS model detection and audio output validation
- Implement WAV file generation for TTS audio responses
- Add audio format checking utilities in CLI output

## v1.4.266 (2025-07-25)

### PR [#1649](https://github.com/danielmiessler/Fabric/pull/1649) by [ksylvan](https://github.com/ksylvan): Fix Conditional API Initialization to Prevent Unnecessary Error Messages

- Prevent unconfigured API initialization and add Docker test suite
- Add BEDROCK_AWS_REGION requirement for Bedrock initialization
- Implement IsConfigured check for Ollama API URL
- Create comprehensive Docker testing environment with 6 scenarios
- Add interactive test runner with shell access

## v1.4.265 (2025-07-25)

### PR [#1647](https://github.com/danielmiessler/Fabric/pull/1647) by [ksylvan](https://github.com/ksylvan): Simplify Workflow with Single Version Retrieval Step

- Replace git tag lookup with version.nix file reading for release workflow
- Remove OS-specific git tag retrieval steps and add unified version extraction from nix file
- Include version format validation with regex check
- Add error handling for missing version file
- Consolidate cross-platform version logic into single step with bash shell for consistent version parsing

## v1.4.264 (2025-07-22)

### PR [#1642](https://github.com/danielmiessler/Fabric/pull/1642) by [ksylvan](https://github.com/ksylvan): Add --sync-db to `generate_changelog`, plus many fixes

- Add database synchronization command with comprehensive validation and sync-db flag for database integrity validation
- Implement version and commit existence checking methods with enhanced time parsing using RFC3339Nano fallback support
- Improve timestamp handling and merge commit detection in changelog generator with comprehensive merge commit detection using parents
- Add email field support to PRCommit struct for author information and improve error logging throughout changelog generation
- Optimize merge pattern matching with lazy initialization and thread-safe pattern compilation for better performance

### Direct commits

- Chore: incoming 1642 changelog entry
- Fix: improve error message formatting in version date parsing

- Add actual error details to date parsing failure message

- Include error variable in stderr output formatting
- Enhance debugging information for invalid date formats
- Docs: Update CHANGELOG after v1.4.263

## v1.4.263 (2025-07-21)

### PR [#1641](https://github.com/danielmiessler/Fabric/pull/1641) by [ksylvan](https://github.com/ksylvan): Fix Fabric Web timeout error

- Chore: extend proxy timeout in `vite.config.ts` to 15 minutes
- Increase `/api` proxy timeout to 900,000 ms
- Increase `/names` proxy timeout to 900,000 ms

## v1.4.262 (2025-07-21)

### PR [#1640](https://github.com/danielmiessler/Fabric/pull/1640) by [ksylvan](https://github.com/ksylvan): Implement Automated Changelog System for CI/CD Integration

- Add automated changelog processing for CI/CD integration with comprehensive test coverage and GitHub client validation methods
- Implement release aggregation for incoming files with git operations for staging changes and support for version detection from nix files
- Change push behavior from opt-out to opt-in with GitHub token authentication and automatic repository detection
- Enhance changelog generation to avoid duplicate commit entries by extracting PR numbers and filtering commits already included via PR files
- Add version parameter requirement for PR processing with commit SHA tracking to prevent duplicate entries and improve formatting consistency

### Direct commits

- Docs: Update CHANGELOG after v1.4.261

## v1.4.261 (2025-07-19)

### PR [#1637](https://github.com/danielmiessler/Fabric/pull/1637) by [ksylvan](https://github.com/ksylvan): chore: update `NeedsRawMode` to include `mistral` prefix for Ollama

- Updated `NeedsRawMode` to include `mistral` prefix for Ollama compatibility
- Added `mistral` to `ollamaPrefixes` list for improved model support

### Direct commits

- Updated CHANGELOG after v1.4.260 release

## v1.4.260 (2025-07-18)

### PR [#1634](https://github.com/danielmiessler/Fabric/pull/1634) by [ksylvan](https://github.com/ksylvan): Fix abort in Exo-Labs provider plugin; with credit to @sakithahSenid

- Fix abort issue in Exo-Labs provider plugin
- Add API key setup question to Exolab AI plugin configuration
- Include API key setup question in Exolab client with required field validation
- Add "openaiapi" to VSCode spell check dictionary
- Maintain existing API base URL configuration order

### Direct commits

- Update CHANGELOG after v1.4.259

## v1.4.259 (2025-07-18)

### PR [#1633](https://github.com/danielmiessler/Fabric/pull/1633) by [ksylvan](https://github.com/ksylvan): YouTube VTT Processing Enhancement

- Fix: prevent duplicate segments in VTT file processing by adding deduplication map to track seen segments
- Feat: enhance VTT duplicate filtering to allow legitimate repeated content with configurable time gap detection
- Feat: improve timestamp parsing to handle fractional seconds and optional seconds/milliseconds formats
- Chore: refactor timestamp regex to global scope and improve performance by avoiding repeated compilation
- Fix: Youtube VTT parsing gap test and extract seconds parsing logic into reusable function

### Direct commits

- Docs: Update CHANGELOG after v1.4.258

## v1.4.258 (2025-07-17)

### PR [#1629](https://github.com/danielmiessler/Fabric/pull/1629) by [ksylvan](https://github.com/ksylvan): Create Default (empty) .env in ~/.config/fabric on Demand

- Add startup check to initialize config and .env file automatically
- Introduce ensureEnvFile function to create ~/.config/fabric/.env if missing
- Add directory creation for config path in ensureEnvFile
- Integrate setup flag in CLI to call ensureEnvFile on demand
- Improve error handling and permissions in ensureEnvFile function

### Direct commits

- Update README and CHANGELOG after v1.4.257

## v1.4.257 (2025-07-17)

### PR [#1628](https://github.com/danielmiessler/Fabric/pull/1628) by [ksylvan](https://github.com/ksylvan): Introduce CLI Flag to Disable OpenAI Responses API

- Add `--disable-responses-api` CLI flag for OpenAI control and llama-server compatibility
- Implement `SetResponsesAPIEnabled` method in OpenAI client with configuration control
- Update default config path to `~/.config/fabric/config.yaml`
- Add CLI completions for new API flag across zsh, bash, and fish shells
- Update CHANGELOG after v1.4.256 release

## v1.4.256 (2025-07-17)

### PR [#1624](https://github.com/danielmiessler/Fabric/pull/1624) by [ksylvan](https://github.com/ksylvan): Feature: Add Automatic ~/.fabric.yaml Config Detection

- Implement default ~/.fabric.yaml config file detection
- Add support for short flag parsing with dashes
- Improve dry run output formatting and config path error handling
- Refactor dry run response construction into helper method
- Extract flag parsing logic into separate extractFlag function

### Direct commits

- Docs: Update CHANGELOG after v1.4.255

## v1.4.255 (2025-07-16)

### Direct commits

- Merge branch 'danielmiessler:main' into main
- Chore: add more paths to update-version-andcreate-tag workflow to reduce unnecessary tagging

## v1.4.254 (2025-07-16)

### PR [#1621](https://github.com/danielmiessler/Fabric/pull/1621) by [robertocarvajal](https://github.com/robertocarvajal): Adds generate code rules pattern

- Adds generate code rules pattern

### Direct commits

- Docs: Update CHANGELOG after v1.4.253

## v1.4.253 (2025-07-16)

### PR [#1620](https://github.com/danielmiessler/Fabric/pull/1620) by [ksylvan](https://github.com/ksylvan): Update Shell Completions for New Think-Block Suppression Options

- Add `--suppress-think` option to suppress 'think' tags
- Introduce `--think-start-tag` and `--think-end-tag` options for text suppression and completion
- Update bash completion with 'think' tag options
- Update fish completion with 'think' tag options
- Update CHANGELOG after v.1.4.252

## v1.4.252 (2025-07-16)

### PR [#1619](https://github.com/danielmiessler/Fabric/pull/1619) by [ksylvan](https://github.com/ksylvan): Feature: Optional Hiding of Model Thinking Process with Configurable Tags

- Add suppress-think flag to hide thinking blocks from AI reasoning output
- Configure customizable start and end thinking tags for content filtering
- Update streaming logic to respect suppress-think setting with YAML configuration support
- Implement StripThinkBlocks utility function with comprehensive testing for thinking suppression
- Performance improvement: add regex caching to StripThinkBlocks function

### Direct commits

- Update CHANGELOG after v1.4.251

## v1.4.251 (2025-07-16)

### PR [#1618](https://github.com/danielmiessler/Fabric/pull/1618) by [ksylvan](https://github.com/ksylvan): Update GitHub Workflow to Ignore Additional File Paths

- Ci: update workflow to ignore additional paths during version updates
- Add `data/strategies/**` to paths-ignore list
- Add `cmd/generate_changelog/*.db` to paths-ignore list
- Prevent workflow triggers from strategy data changes
- Prevent workflow triggers from changelog database files

## v1.4.250 (2025-07-16)

### Direct commits

- Docs: Update changelog with v1.4.249 changes

## v1.4.249 (2025-07-16)

### PR [#1617](https://github.com/danielmiessler/Fabric/pull/1617) by [ksylvan](https://github.com/ksylvan): Improve PR Sync Logic for Changelog Generator

- Preserve PR numbers during version cache merges
- Enhance changelog to associate PR numbers with version tags
- Improve PR number parsing with proper error handling
- Collect all PR numbers for commits between version tags
- Associate aggregated PR numbers with each version entry

## v1.4.248 (2025-07-16)

### PR [#1616](https://github.com/danielmiessler/Fabric/pull/1616) by [ksylvan](https://github.com/ksylvan): Preserve PR Numbers During Version Cache Merges

- Feat: enhance changelog to correctly associate PR numbers with version tags
- Fix: improve PR number parsing with proper error handling
- Collect all PR numbers for commits between version tags
- Associate aggregated PR numbers with each version entry
- Update cached versions with newly found PR numbers

### Direct commits

- Docs: reorganize v1.4.247 changelog to attribute changes to PR #1613

## v1.4.247 (2025-07-15)

### PR [#1613](https://github.com/danielmiessler/Fabric/pull/1613) by [ksylvan](https://github.com/ksylvan): Improve AI Summarization for Consistent Professional Changelog Entries

- Feat: enhance changelog generation with incremental caching and improved AI summarization
- Add incremental processing for new Git tags since cache
- Implement `WalkHistorySinceTag` method for efficient history traversal
- Add custom patterns directory support to plugin registry
- Feat: improve error handling in `plugin_registry` and `patterns_loader`

### Direct commits

- Docs: update README for GraphQL optimization and AI summary features

## v1.4.246 (2025-07-14)

### PR [#1611](https://github.com/danielmiessler/Fabric/pull/1611) by [ksylvan](https://github.com/ksylvan): Changelog Generator: AI-Powered Automation for Fabric Project

- Add AI-powered changelog generation with high-performance Go tool and comprehensive caching
- Implement SQLite-based persistent caching for incremental updates with one-pass git history walking algorithm
- Create comprehensive CLI with cobra framework and tag-based caching integration
- Integrate AI summarization using Fabric CLI with batch PR fetching and GitHub Search API optimization
- Add extensive documentation with PRD and README files, including commit-PR mapping for optimized git operations

## v1.4.245 (2025-07-11)

### PR [#1603](https://github.com/danielmiessler/Fabric/pull/1603) by [ksylvan](https://github.com/ksylvan): Together AI Support with OpenAI Fallback Mechanism Added

- Added direct model fetching support for non-standard providers with fallback mechanism
- Enhanced error messages in OpenAI compatible models endpoint with response body details
- Improved OpenAI compatible models API client with timeout and cleaner parsing
- Added context support to DirectlyGetModels method with proper error handling
- Optimized HTTP request handling and improved error response formatting

### PR [#1599](https://github.com/danielmiessler/Fabric/pull/1599) by [ksylvan](https://github.com/ksylvan): Update file paths to reflect new data directory structure

- Updated file paths to reflect new data directory structure including patterns and strategies locations

### Direct commits

- Fixed broken image link

## v1.4.244 (2025-07-09)

### PR [#1598](https://github.com/danielmiessler/Fabric/pull/1598) by [jaredmontoya](https://github.com/jaredmontoya): flake: fixes and enhancements

- Nix:pkgs:fabric: use self reference
- Shell: rename command
- Update-mod: fix generation path
- Shell: fix typo

## v1.4.243 (2025-07-09)

### PR [#1597](https://github.com/danielmiessler/Fabric/pull/1597) by [ksylvan](https://github.com/ksylvan): CLI Refactoring: Modular Command Processing and Pattern Loading Improvements

- Refactor CLI to modularize command handling with specialized handlers for setup, configuration, listing, management, and extensions
- Improve patterns loader with migration support and better error handling
- Add tool processing for YouTube and web scraping functionality
- Enhance error handling and early returns in CLI to prevent panics
- Improve error handling and temporary file management in patterns loader with secure temporary directory creation

### Direct commits

- Nix:pkgs:fabric: use self reference
- Update-mod: fix generation path
- Shell: rename command

## v1.4.242 (2025-07-09)

### PR [#1596](https://github.com/danielmiessler/Fabric/pull/1596) by [ksylvan](https://github.com/ksylvan): Fix patterns zipping workflow

- Chore: update workflow paths to reflect directory structure change
- Modify trigger path to `data/patterns/**`
- Update `git diff` command to new path
- Change zip command to include `data/patterns/` directory

## v1.4.241 (2025-07-09)

### PR [#1595](https://github.com/danielmiessler/Fabric/pull/1595) by [ksylvan](https://github.com/ksylvan): Restructure project to align with standard Go layout

- Restructure project to align with standard Go layout by introducing `cmd` directory for binaries and moving packages to `internal` directory
- Consolidate patterns and strategies into new `data` directory and group auxiliary scripts into `scripts` directory
- Move documentation and images into `docs` directory and update all Go import paths to reflect new structure
- Rename `restapi` package to `server` for clarity and reorganize OAuth storage functionality into util package
- Add new patterns for content tagging and cognitive bias analysis including apply_ul_tags and t_check_dunning_kruger

### PR [#1594](https://github.com/danielmiessler/Fabric/pull/1594) by [amancioandre](https://github.com/amancioandre): Adds check Dunning-Kruger Telos self-evaluation pattern

- Add pattern telos check dunning kruger for cognitive bias self-evaluation

## v1.4.240 (2025-07-07)

### PR [#1593](https://github.com/danielmiessler/Fabric/pull/1593) by [ksylvan](https://github.com/ksylvan): Refactor: Generalize OAuth flow for improved token handling

- Refactor: replace hardcoded "claude" with configurable `authTokenIdentifier` parameter for improved flexibility
- Update `RunOAuthFlow` and `RefreshToken` functions to accept token identifier parameter instead of hardcoded values
- Add token refresh attempt before full OAuth flow to improve authentication efficiency
- Test: add comprehensive OAuth testing suite with 434 lines coverage including mock token server and PKCE validation
- Chore: refactor token path to use `authTokenIdentifier` for consistent token handling across the system

## v1.4.239 (2025-07-07)

### PR [#1592](https://github.com/danielmiessler/Fabric/pull/1592) by [ksylvan](https://github.com/ksylvan): Fix Streaming Error Handling in Chatter

- Fix: improve error handling in streaming chat functionality
- Add dedicated error channel for stream operations
- Refactor: use select to handle stream and error channels concurrently
- Feat: add test for Chatter's Send method error propagation
- Chore: enhance `Chatter.Send` method with proper goroutine synchronization

## v1.4.238 (2025-07-07)

### PR [#1591](https://github.com/danielmiessler/Fabric/pull/1591) by [ksylvan](https://github.com/ksylvan): Improved Anthropic Plugin Configuration Logic

- Add vendor configuration validation and OAuth auto-authentication
- Implement IsConfigured method for Anthropic client validation with automatic OAuth flow when no valid token
- Add token expiration checking with 5-minute buffer for improved reliability
- Extract vendor token identifier into named constant for better code maintainability
- Remove redundant Configure() call from IsConfigured method to improve performance

## v1.4.237 (2025-07-07)

### PR [#1590](https://github.com/danielmiessler/Fabric/pull/1590) by [ksylvan](https://github.com/ksylvan): Do not pass non-default TopP values

- Fix: add conditional check for TopP parameter in OpenAI client
- Add zero-value check before setting TopP parameter
- Prevent sending TopP when value is zero
- Apply fix to both chat completions method
- Apply fix to response parameters method

## v1.4.236 (2025-07-06)

### PR [#1587](https://github.com/danielmiessler/Fabric/pull/1587) by [ksylvan](https://github.com/ksylvan): Enhance bug report template

- Chore: enhance bug report template with detailed system info and installation method fields
- Add detailed instructions for bug reproduction steps
- Include operating system dropdown with specific architectures
- Add OS version textarea with command examples
- Create installation method dropdown with all options

## v1.4.235 (2025-07-06)

### PR [#1586](https://github.com/danielmiessler/Fabric/pull/1586) by [ksylvan](https://github.com/ksylvan): Fix to persist the CUSTOM_PATTERNS_DIRECTORY variable

- Fix: make custom patterns persist correctly

## v1.4.234 (2025-07-06)

### PR [#1581](https://github.com/danielmiessler/Fabric/pull/1581) by [ksylvan](https://github.com/ksylvan): Fix Custom Patterns Directory Creation Logic

- Chore: improve directory creation logic in `configure` method
- Add `fmt` package for logging errors
- Check directory existence before creating
- Log error without clearing directory value

## v1.4.233 (2025-07-06)

### PR [#1580](https://github.com/danielmiessler/Fabric/pull/1580) by [ksylvan](https://github.com/ksylvan): Alphabetical Pattern Sorting and Configuration Refactor

- Refactor: move custom patterns directory initialization to Configure method
- Add alphabetical sorting to pattern names retrieval
- Improve pattern listing with proper error handling
- Ensure custom patterns loaded after environment configuration

### PR [#1578](https://github.com/danielmiessler/Fabric/pull/1578) by [ksylvan](https://github.com/ksylvan): Document Custom Patterns Directory Support

- Add comprehensive custom patterns setup and usage guide

## v1.4.232 (2025-07-06)

### PR [#1577](https://github.com/danielmiessler/Fabric/pull/1577) by [ksylvan](https://github.com/ksylvan): Add Custom Patterns Directory Support

- Add custom patterns directory support via environment variable configuration
- Implement custom patterns plugin with registry integration and pattern precedence
- Override main patterns with custom directory patterns for enhanced flexibility
- Expand home directory paths in custom patterns config for better usability
- Add comprehensive test coverage for custom patterns functionality

## v1.4.231 (2025-07-05)

### PR [#1565](https://github.com/danielmiessler/Fabric/pull/1565) by [ksylvan](https://github.com/ksylvan): OAuth Authentication Support for Anthropic

- Feat: add OAuth authentication support for Anthropic Claude
- Implement PKCE OAuth flow with browser integration
- Add automatic OAuth token refresh when expired
- Implement persistent token storage using common OAuth storage
- Refactor: extract OAuth functionality from anthropic client to separate module

## v1.4.230 (2025-07-05)

### PR [#1575](https://github.com/danielmiessler/Fabric/pull/1575) by [ksylvan](https://github.com/ksylvan): Advanced image generation parameters for OpenAI models

- Add advanced image generation parameters for OpenAI models with four new CLI flags
- Implement validation for image parameter combinations with size, quality, compression, and background controls
- Add comprehensive test coverage for new image generation parameters
- Update shell completions to support new image options
- Enhance README with detailed image generation examples and fix PowerShell code block formatting issues

## v1.4.229 (2025-07-05)

### PR [#1574](https://github.com/danielmiessler/Fabric/pull/1574) by [ksylvan](https://github.com/ksylvan): Add Model Validation for Image Generation and Fix CLI Flag Mapping

- Add model validation for image generation support with new `supportsImageGeneration` function
- Implement model field in `BuildChatOptions` method for proper CLI flag mapping
- Refactor model validation logic by extracting supported models list to shared constant `ImageGenerationSupportedModels`
- Add comprehensive tests for model validation logic in `TestModelValidationLogic`
- Remove unused `mars-colony.png` file from repository

## v1.4.228 (2025-07-05)

### PR [#1573](https://github.com/danielmiessler/Fabric/pull/1573) by [ksylvan](https://github.com/ksylvan): Add Image File Validation and Dynamic Format Support

- Add image file path validation with extension checking
- Implement dynamic output format detection from file extensions
- Update BuildChatOptions method to return error for validation
- Add comprehensive test coverage for image file validation
- Upgrade YAML library from v2 to v3

### Direct commits

- Added tutorial as a tag

## v1.4.227 (2025-07-04)

### PR [#1572](https://github.com/danielmiessler/Fabric/pull/1572) by [ksylvan](https://github.com/ksylvan): Add Image Generation Support to Fabric

- Add image generation support with OpenAI image generation model and `--image-file` flag for saving generated images
- Implement web search tool for Anthropic and OpenAI models with search location parameter support
- Add comprehensive test coverage for image features and update documentation with image generation examples
- Support multiple image formats (PNG, JPG, JPEG, GIF, BMP) and image editing with attachment input files
- Refactor image generation constants for clarity and reuse with defined response type and tool type constants

### Direct commits

- Fixed ul tag applier and updated ul tag prompt
- Added the UL tags pattern

## v1.4.226 (2025-07-04)

### PR [#1569](https://github.com/danielmiessler/Fabric/pull/1569) by [ksylvan](https://github.com/ksylvan): OpenAI Plugin Now Supports Web Search Functionality

- Feat: add web search tool support for OpenAI models with citation formatting
- Enable web search tool for OpenAI models
- Add location parameter support for search results
- Extract and format citations from search responses
- Implement citation deduplication to avoid duplicates

## v1.4.225 (2025-07-04)

### PR [#1568](https://github.com/danielmiessler/Fabric/pull/1568) by [ksylvan](https://github.com/ksylvan): Runtime Web Search Control via Command-Line Flag

- Add web search tool support for Anthropic models with --search flag to enable web search functionality
- Add --search-location flag for timezone-based search results and pass search options through ChatOptions struct
- Implement web search tool in Anthropic client with formatted search citations and sources section
- Add comprehensive tests for search functionality and remove plugin-level web search configuration
- Refactor web search tool constants in anthropic plugin to improve code maintainability through constant extraction

### Direct commits

- Fix: sections as heading 1, typos
- Feat: adds pattern telos check dunning kruger

## v1.4.224 (2025-07-01)

### PR [#1564](https://github.com/danielmiessler/Fabric/pull/1564) by [ksylvan](https://github.com/ksylvan): Add code_review pattern and updates in Pattern_Descriptions

- Added comprehensive code review pattern with systematic analysis framework and principal engineer reviewer role
- Introduced new patterns for code review, alpha extraction, and server analysis (`review_code`, `extract_alpha`, `extract_mcp_servers`)
- Enhanced pattern extraction script with improved clarity, docstrings, and specific error handling
- Implemented graceful JSONDecodeError handling in `load_existing_file` function with warning messages
- Fixed typo in `analyze_bill_short` pattern description and improved formatting in pattern management README

## v1.4.223 (2025-07-01)

### PR [#1563](https://github.com/danielmiessler/Fabric/pull/1563) by [ksylvan](https://github.com/ksylvan): Fix Cross-Platform Compatibility in Release Workflow

- Chore: update GitHub Actions to use bash shell in release job
- Adjust repository_dispatch type spacing for consistency
- Use bash shell for creating release if absent

## v1.4.222 (2025-07-01)

### PR [#1559](https://github.com/danielmiessler/Fabric/pull/1559) by [ksylvan](https://github.com/ksylvan): OpenAI Plugin Migrates to New Responses API

- Migrate OpenAI plugin to use new responses API instead of chat completions
- Add chat completions API fallback for non-Responses API providers
- Fix channel close handling in OpenAI streaming methods to prevent potential leaks
- Extract common message conversion logic to reduce code duplication
- Add support for multi-content user messages including image URLs in chat completions

## v1.4.221 (2025-06-28)

### PR [#1556](https://github.com/danielmiessler/Fabric/pull/1556) by [ksylvan](https://github.com/ksylvan): feat: Migrate to official openai-go SDK

- Refactor: abstract chat message structs and migrate to official openai-go SDK
- Introduce local `chat` package for message abstraction
- Replace sashabaranov/go-openai with official openai-go SDK
- Update OpenAI, Azure, and Exolab plugins for new client
- Refactor all AI providers to use internal chat types

## v1.4.220 (2025-06-28)

### PR [#1555](https://github.com/danielmiessler/Fabric/pull/1555) by [ksylvan](https://github.com/ksylvan): fix: Race condition in GitHub actions release flow

- Chore: improve release creation to gracefully handle pre-existing tags.
- Check if a release exists before attempting creation.
- Suppress error output from `gh release view` command.
- Add an informative log when release already exists.

## v1.4.219 (2025-06-28)

### PR [#1553](https://github.com/danielmiessler/Fabric/pull/1553) by [ksylvan](https://github.com/ksylvan): docs: add DeepWiki badge and fix minor typos in README

- Add DeepWiki badge to README header
- Fix typo "chatbots" to "chat-bots"
- Correct "Perlexity" to "Perplexity"
- Fix "distro" to "Linux distribution"
- Add alt text to contributor images

### PR [#1552](https://github.com/danielmiessler/Fabric/pull/1552) by [nawarajshahi](https://github.com/nawarajshahi): Fix typos in README.md

- Fix typos on README.md

## v1.4.218 (2025-06-27)

### PR [#1550](https://github.com/danielmiessler/Fabric/pull/1550) by [ksylvan](https://github.com/ksylvan): Add Support for OpenAI Search and Research Model Variants

- Add support for new OpenAI search and research model variants
- Define new search preview model names and mini search preview variants
- Include deep research model support with June 2025 dated model versions
- Replace hardcoded check with slices.Contains for better array operations
- Support both prefix and exact model matching functionality

## v1.4.217 (2025-06-26)

### PR [#1546](https://github.com/danielmiessler/Fabric/pull/1546) by [ksylvan](https://github.com/ksylvan): New YouTube Transcript Endpoint Added to REST API

- Added dedicated YouTube transcript API endpoint with `/youtube/transcript` POST route
- Implemented YouTube handler for transcript requests with language and timestamp options
- Updated frontend to use new endpoint and removed chat endpoint dependency for transcripts
- Added proper validation for video vs playlist URLs
- Fixed endpoint calls from frontend

### Direct commits

- Added extract_mcp_servers pattern to identify MCP (Model Context Protocol) servers from content, including server names, features, capabilities, and usage examples

## v1.4.216 (2025-06-26)

### PR [#1545](https://github.com/danielmiessler/Fabric/pull/1545) by [ksylvan](https://github.com/ksylvan): Update Message Handling for Attachments and Multi-Modal content

- Allow combining user messages and attachments with patterns
- Enhance dryrun client to display multi-content user messages including image URLs
- Prevent duplicate user message when applying patterns while ensuring multi-part content is included
- Extract message and option formatting logic into reusable methods to reduce code duplication
- Add MultiContent support to chat message construction in raw mode with proper text and attachment combination

## v1.4.215 (2025-06-25)

### PR [#1543](https://github.com/danielmiessler/Fabric/pull/1543) by [ksylvan](https://github.com/ksylvan): fix: Revert multiline tags in generated json files

- Chore: reformat `pattern_descriptions.json` to improve readability
- Reformat JSON `tags` array to display on new lines
- Update `write_essay` pattern description for clarity
- Apply consistent formatting to both data files

## v1.4.214 (2025-06-25)

### PR [#1542](https://github.com/danielmiessler/Fabric/pull/1542) by [ksylvan](https://github.com/ksylvan): Add `write_essay_by_author` and update Pattern metadata

- Refactor ProviderMap for dynamic URL template handling with environment variables
- Add new pattern `write_essay_by_author` for stylistic writing with author variable usage
- Introduce `analyze_terraform_plan` pattern for infrastructure review
- Add `summarize_board_meeting` pattern for corporate notes
- Rename `write_essay` to `write_essay_pg` for Paul Graham style clarity

## v1.4.213 (2025-06-23)

### PR [#1538](https://github.com/danielmiessler/Fabric/pull/1538) by [andrewsjg](https://github.com/andrewsjg): Bug/bedrock region handling

- Updated hasAWSCredentials to also check for AWS_DEFAULT_REGION when access keys are configured in the environment
- Fixed bedrock region handling with corrected pointer reference and proper region value setting
- Refactored Bedrock client to improve error handling and add interface compliance
- Added AWS region validation logic and enhanced error handling with wrapped errors
- Improved resource cleanup in SendStream with nil checks for response parsing

## v1.4.212 (2025-06-23)

### PR [#1540](https://github.com/danielmiessler/Fabric/pull/1540) by [ksylvan](https://github.com/ksylvan): Add Langdock AI and enhance generic OpenAI compatible support

- Implement dynamic URL handling with environment variables for provider configuration
- Refactor ProviderMap to support URL templates with template variable parsing
- Extract and parse template variables from BaseURL with fallback to default values
- Add `os` and `strings` packages to imports for enhanced functionality
- Reorder providers for consistent key order in ProviderMap

### Direct commits

- Improve Bedrock client error handling with wrapped errors and AWS region validation
- Add ai.Vendor interface implementation check for better compliance
- Fix resource cleanup in SendStream with proper nil checks for response parsing
- Update AWS credentials checking to include AWS_DEFAULT_REGION environment variable
- Update paper analyzer functionality

## v1.4.211 (2025-06-19)

### PR [#1533](https://github.com/danielmiessler/Fabric/pull/1533) by [ksylvan](https://github.com/ksylvan): REST API and Web UI Now Support Dynamic Pattern Variables

- Added pattern variables support to REST API chat endpoint with Variables field in PromptRequest struct
- Implemented pattern variables UI in web interface with JSON textarea for variable input and dedicated Svelte store
- Created new `ApplyPattern` route for POST /patterns/:name/apply with `PatternApplyRequest` struct for request body parsing
- Refactored chat service to clean up message stream and pattern output methods with improved stream readability
- Merged query parameters with request body variables in `ApplyPattern` method using `StorageHandler` for pattern operations

## v1.4.210 (2025-06-18)

### PR [#1530](https://github.com/danielmiessler/Fabric/pull/1530) by [ksylvan](https://github.com/ksylvan): Add Citation Support to Perplexity Response

- Add citation support to Perplexity AI responses with automatic extraction from API responses
- Append citations section to response content formatted as numbered markdown list
- Handle citations in streaming responses while maintaining backward compatibility
- Store last response for citation access and add citations after stream completion

### Direct commits

- Update README.md with improved intro text describing Fabric's utility to most people

## v1.4.208 (2025-06-17)

### PR [#1527](https://github.com/danielmiessler/Fabric/pull/1527) by [ksylvan](https://github.com/ksylvan): Add Perplexity AI Provider with Token Limits Support

- Add Perplexity AI provider support with token limits and streaming capabilities
- Add `MaxTokens` field to `ChatOptions` struct for response control
- Integrate Perplexity client into core plugin registry initialization
- Implement stream handling in Perplexity client using sync.WaitGroup
- Update README with Perplexity AI support instructions and configuration examples

### PR [#1526](https://github.com/danielmiessler/Fabric/pull/1526) by [ConnorKirk](https://github.com/ConnorKirk): Check for AWS_PROFILE or AWS_ROLE_SESSION_NAME environment variables

- Check for AWS_PROFILE or AWS_ROLE_SESSION_NAME environment variables

## v1.4.207 (2025-06-17)

### PR [#1525](https://github.com/danielmiessler/Fabric/pull/1525) by [ksylvan](https://github.com/ksylvan): Refactor yt-dlp Transcript Logic and Fix Language Bug

- Refactored yt-dlp logic to reduce code duplication in YouTube plugin by extracting shared logic into tryMethodYtDlpInternal helper
- Added processVTTFileFunc parameter for flexible VTT processing and implemented language matching for 2-character language codes
- Improved transcript methods structure while maintaining existing functionality
- Updated extract insights functionality

## v1.4.206 (2025-06-16)

### PR [#1523](https://github.com/danielmiessler/Fabric/pull/1523) by [ksylvan](https://github.com/ksylvan): Conditional AWS Bedrock Plugin Initialization

- Add AWS credential detection for Bedrock client initialization
- Check for AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
- Look for AWS shared credentials file with support for custom AWS_SHARED_CREDENTIALS_FILE path
- Only initialize Bedrock client if credentials exist to prevent AWS SDK credential search failures
- Updated prompt

## v1.4.205 (2025-06-16)

### PR [#1519](https://github.com/danielmiessler/Fabric/pull/1519) by [ConnorKirk](https://github.com/ConnorKirk): feat: Dynamically list AWS Bedrock models

- Dynamically fetch and list available foundation models and inference profiles

### PR [#1518](https://github.com/danielmiessler/Fabric/pull/1518) by [ksylvan](https://github.com/ksylvan): chore: remove duplicate/outdated patterns

- Chore: remove duplicate/outdated patterns

### Direct commits

- Updated markdown sanitizer
- Updated markdown cleaner

## v1.4.204 (2025-06-15)

### PR [#1517](https://github.com/danielmiessler/Fabric/pull/1517) by [ksylvan](https://github.com/ksylvan): Fix: Prevent race conditions in versioning workflow

- Ci: improve version update workflow to prevent race conditions
- Add concurrency control to prevent simultaneous runs
- Pull latest main branch changes before tagging
- Fetch all remote tags before calculating version

## v1.4.203 (2025-06-14)

### PR [#1512](https://github.com/danielmiessler/Fabric/pull/1512) by [ConnorKirk](https://github.com/ConnorKirk): feat:Add support for Amazon Bedrock

- Add Bedrock plugin for using Amazon Bedrock within fabric

### PR [#1513](https://github.com/danielmiessler/Fabric/pull/1513) by [marcas756](https://github.com/marcas756): feat: create mnemonic phrase pattern

- Add new pattern for generating mnemonic phrases from diceware words with user guide and system implementation details

### PR [#1516](https://github.com/danielmiessler/Fabric/pull/1516) by [ksylvan](https://github.com/ksylvan): Fix REST API pattern creation

- Add Save method to PatternsEntity for persisting patterns to filesystem
- Create pattern directory with proper permissions and write pattern content to system pattern file
- Add comprehensive test for Save functionality with directory creation and file contents verification
- Handle errors for directory and file operations

## v1.4.202 (2025-06-12)

### PR [#1510](https://github.com/danielmiessler/Fabric/pull/1510) by [ksylvan](https://github.com/ksylvan): Cross-Platform fix for Youtube Transcript extraction

- Replace hardcoded `/tmp` with `os.TempDir()` for cross-platform temporary directory handling
- Use `filepath.Join()` instead of string concatenation for proper path construction
- Remove Unix `find` command dependency and replace with native Go `filepath.Walk()` method
- Add new `findVTTFiles()` method to make VTT file discovery work on Windows
- Improve error handling for file operations while maintaining backward compatibility

## v1.4.201 (2025-06-12)

### PR [#1503](https://github.com/danielmiessler/Fabric/pull/1503) by [dependabot[bot]](https://github.com/apps/dependabot): chore(deps): bump brace-expansion from 1.1.11 to 1.1.12 in /web in the npm_and_yarn group across 1 directory

- Updated brace-expansion dependency from version 1.1.11 to 1.1.12 in the web directory

### PR [#1508](https://github.com/danielmiessler/Fabric/pull/1508) by [ksylvan](https://github.com/ksylvan): feat: cleanup after `yt-dlp` addition

- Updated README documentation to include yt-dlp requirement for transcripts
- Improved error messages to be clearer and more actionable

## v1.4.200 (2025-06-11)

### PR [#1507](https://github.com/danielmiessler/Fabric/pull/1507) by [ksylvan](https://github.com/ksylvan): Refactor: No more web scraping, just use yt-dlp

- Refactor: replace web scraping with yt-dlp for YouTube transcript extraction
- Remove unreliable YouTube API scraping methods
- Add yt-dlp integration for transcript extraction
- Implement VTT subtitle parsing functionality
- Add timestamp preservation for transcripts

## v1.4.199 (2025-06-11)

### PR [#1506](https://github.com/danielmiessler/Fabric/pull/1506) by [eugeis](https://github.com/eugeis): fix: fix web search tool location

- Fix: fix web search tool location

## v1.4.198 (2025-06-11)

### PR [#1504](https://github.com/danielmiessler/Fabric/pull/1504) by [marcas756](https://github.com/marcas756): fix: Add configurable HTTP timeout for Ollama client

- Fix: Add configurable HTTP timeout for Ollama client with default value set to 20 minutes

## v1.4.197 (2025-06-11)

### PR [#1502](https://github.com/danielmiessler/Fabric/pull/1502) by [eugeis](https://github.com/eugeis): Feat/antropic tool

- Feat: search tool working
- Feat: search tool result collection

### PR [#1499](https://github.com/danielmiessler/Fabric/pull/1499) by [noamsiegel](https://github.com/noamsiegel): feat: Enhance the PRD Generator's identity and purpose

- Feat: Enhance the PRD Generator's identity and purpose with expanded role definition and structured output format
- Add comprehensive PRD sections including Overview, Objectives, Target Audience, Features, User Stories, and Success Metrics
- Provide detailed instructions for Markdown formatting with labeled sections, bullet points, and priority highlighting

### PR [#1497](https://github.com/danielmiessler/Fabric/pull/1497) by [ksylvan](https://github.com/ksylvan): feat: add Terraform plan analyzer pattern for infrastructure changes

- Feat: add Terraform plan analyzer pattern for infrastructure change assessment
- Create expert plan analyzer role with focus on security, cost, and compliance evaluation
- Include structured output format with 20-word summaries, critical changes list, and key takeaways section

### Direct commits

- Fix: Add configurable HTTP timeout for Ollama client with default 20-minute duration
- Chore(deps): bump brace-expansion from 1.1.11 to 1.1.12 in npm_and_yarn group

## v1.4.196 (2025-06-07)

### PR [#1495](https://github.com/danielmiessler/Fabric/pull/1495) by [ksylvan](https://github.com/ksylvan): Add AIML provider configuration

- Add AIML provider to OpenAI compatible providers configuration
- Set AIML base URL to api.aimlapi.com/v1 and expand supported providers list
- Enable AIML API integration support

### Direct commits

- Add simpler paper analyzer functionality
- Update output formatting across multiple components

## v1.4.195 (2025-05-24)

### PR [#1487](https://github.com/danielmiessler/Fabric/pull/1487) by [ksylvan](https://github.com/ksylvan): Dependency Updates and PDF Worker Refactoring

- Feat: upgrade PDF.js to v4.2 and refactor worker initialization
- Add `.browserslistrc` to define target browser versions
- Upgrade `pdfjs-dist` dependency from v2.16 to v4.2.67
- Upgrade `nanoid` dependency from v4.0.2 to v5.0.9
- Introduce `pdf-config.ts` for centralized PDF.js worker setup

## v1.4.194 (2025-05-24)

### PR [#1485](https://github.com/danielmiessler/Fabric/pull/1485) by [ksylvan](https://github.com/ksylvan): Web UI: Centralize Environment Configuration and Make Fabric Base URL Configurable

- Feat: add centralized environment configuration for Fabric base URL
- Create environment config module for URL handling
- Add getFabricBaseUrl() function with server/client support
- Add getFabricApiUrl() helper for API endpoints
- Configure Vite to inject FABRIC_BASE_URL client-side

## v1.4.193 (2025-05-24)

### PR [#1484](https://github.com/danielmiessler/Fabric/pull/1484) by [ksylvan](https://github.com/ksylvan): Web UI update all packages, reorganize docs, add install scripts

- Reorganize web documentation and add installation scripts
- Update all package dependencies to latest versions
- Add PDF-to-Markdown installation steps to README
- Move legacy documentation files to web/legacy/
- Add convenience scripts for npm and pnpm installation

### PR [#1481](https://github.com/danielmiessler/Fabric/pull/1481) by [skibum1869](https://github.com/skibum1869): Add board meeting summary pattern template

- Add board meeting summary pattern template
- Update meeting summary template with word count requirement
- Add minimum word count for context section in board summary

### Direct commits

- Add centralized environment configuration for Fabric base URL
- Create environment config module for URL handling with server/client support
- Configure Vite to inject FABRIC_BASE_URL client-side
- Update proxy targets to use environment variable
- Add TypeScript definitions for window config

## v1.4.192 (2025-05-23)

### PR [#1480](https://github.com/danielmiessler/Fabric/pull/1480) by [ksylvan](https://github.com/ksylvan): Automatic setting of "raw mode" for some models

- Added NeedsRawMode method to AI vendor interface to support model-specific raw mode detection
- Implemented automatic raw mode detection for specific AI models including Ollama llama2/llama3 and OpenAI o1/o3/o4 models
- Enhanced vendor interface with NeedsRawMode implementation across all AI clients
- Added model-specific raw mode detection logic with prefix matching capabilities
- Enabled automatic raw mode activation when vendor requirements are detected

## v1.4.191 (2025-05-22)

### PR [#1478](https://github.com/danielmiessler/Fabric/pull/1478) by [ksylvan](https://github.com/ksylvan): Claude 4 Integration and README Updates

- Add support for Anthropic Claude 4 models and update SDK to v1.2.0
- Upgrade `anthropic-sdk-go` dependency to version `v1.2.0`
- Integrate new Anthropic Claude 4 Opus and Sonnet models
- Remove deprecated Claude 2.0 and 2.1 models from list
- Adjust model type casting for `anthropic-sdk-go v1.2.0` compatibility

## v1.4.190 (2025-05-20)

### PR [#1475](https://github.com/danielmiessler/Fabric/pull/1475) by [ksylvan](https://github.com/ksylvan): refactor: improve raw mode handling in BuildSession

- Refactor: improve raw mode handling in BuildSession
- Fix system message handling with patterns in raw mode
- Prevent duplicate inputs when using patterns
- Add conditional logic for pattern vs non-pattern scenarios
- Simplify message construction with clearer variable names

## v1.4.189 (2025-05-19)

### PR [#1473](https://github.com/danielmiessler/Fabric/pull/1473) by [roumy](https://github.com/roumy): add authentification for ollama instance

- Add authentification for ollama instance

## v1.4.188 (2025-05-19)

### PR [#1474](https://github.com/danielmiessler/Fabric/pull/1474) by [ksylvan](https://github.com/ksylvan): feat: update `BuildSession` to handle message appending logic

- Refactor message handling for raw mode and Anthropic client with improved logic
- Add proper handling for empty message arrays and user/assistant message alternation
- Implement safeguards for message sequence validation and preserve system messages
- Fix pattern-based message handling in non-raw mode with better normalization

### PR [#1467](https://github.com/danielmiessler/Fabric/pull/1467) by [joshuafuller](https://github.com/joshuafuller): Typos, spelling, grammar and other minor updates

- Fix spelling and grammar issues across documentation including pattern management guide, PR notes, and web README

### PR [#1468](https://github.com/danielmiessler/Fabric/pull/1468) by [NavNab](https://github.com/NavNab): Refactor content structure in create_hormozi_offer system.md for clarity and readability

- Improve formatting and content structure in system.md for better flow and readability
- Consolidate repetitive sentences and enhance overall text coherence with consistent bullet points

### Direct commits

- Add authentication for Ollama instance

## v1.4.187 (2025-05-10)

### PR [#1463](https://github.com/danielmiessler/Fabric/pull/1463) by [CodeCorrupt](https://github.com/CodeCorrupt): Add completion to the build output for Nix

- Add completion files to the build output for Nix

## v1.4.186 (2025-05-06)

### PR [#1459](https://github.com/danielmiessler/Fabric/pull/1459) by [ksylvan](https://github.com/ksylvan): chore: Repository cleanup and .gitignore Update

- Add `coverage.out` to `.gitignore` for ignoring coverage output
- Remove `Alma.md` documentation file from the repository
- Delete `rate_ai_result.txt` stitch script from `stitches` folder
- Remove `readme.md` for `rate_ai_result` stitch documentation

## v1.4.185 (2025-04-28)

### PR [#1453](https://github.com/danielmiessler/Fabric/pull/1453) by [ksylvan](https://github.com/ksylvan): Fix for default model setting

- Refactor: introduce `getSortedGroupsItems` for consistent sorting logic
- Add `getSortedGroupsItems` to centralize sorting logic
- Sort groups and items alphabetically, case-insensitive
- Replace inline sorting in `Print` with new method
- Update `GetGroupAndItemByItemNumber` to use sorted data

## v1.4.184 (2025-04-25)

### PR [#1447](https://github.com/danielmiessler/Fabric/pull/1447) by [ksylvan](https://github.com/ksylvan): More shell completion scripts: Zsh, Bash, and Fish

- Add shell completion support for three major shells (Zsh, Bash, and Fish)
- Create standardized completion scripts in completions/ directory
- Add --shell-complete-list flag for machine-readable output
- Update Print() methods to support plain output format
- Replace old fish completion script with improved version

## v1.4.183 (2025-04-23)

### PR [#1431](https://github.com/danielmiessler/Fabric/pull/1431) by [KenMacD](https://github.com/KenMacD): Add a completion script for fish

- Add a completion script for fish

## v1.4.182 (2025-04-23)

### PR [#1441](https://github.com/danielmiessler/Fabric/pull/1441) by [ksylvan](https://github.com/ksylvan): Update go toolchain and go module packages to latest versions

- Updated Go version to 1.24.2 across Dockerfile, Nix configurations, and Go modules
- Refreshed Go module dependencies and updated go.mod and go.sum files
- Updated Nix flake lock file inputs and configured Nix environment for Go 1.24
- Centralized Go version definition by creating `getGoVersion` function in flake.nix for consistent version management
- Fixed "nix flake check" errors and removed redundant Go version definitions

## v1.4.181 (2025-04-22)

### PR [#1433](https://github.com/danielmiessler/Fabric/pull/1433) by [ksylvan](https://github.com/ksylvan): chore: update Anthropic SDK to v0.2.0-beta.3 and migrate to V2 API

- Upgrade Anthropic SDK from alpha.11 to beta.3
- Update API endpoint from v1 to v2
- Replace anthropic.F() with direct assignment for required parameters
- Replace anthropic.F() with anthropic.Opt() for optional parameters
- Simplify event delta handling in streaming responses

## v1.4.180 (2025-04-22)

### PR [#1435](https://github.com/danielmiessler/Fabric/pull/1435) by [ksylvan](https://github.com/ksylvan): chore: Fix user input handling when using raw mode and `--strategy` flag

- Fixed user input handling when using raw mode and `--strategy` flag by unifying raw mode message handling and preserving environment variables in extension executor
- Refactored BuildSession raw mode to prepend system to user content and ensure raw mode messages always have User role
- Improved session handling by appending systemMessage separately in non-raw mode sessions and storing original command environment before context-based execution
- Added comments clarifying raw vs non-raw handling behavior for better code maintainability

### Direct commits

- Updated Anthropic SDK to v0.2.0-beta.3 and migrated to V2 API, including endpoint changes from v1 to v2 and replacement of anthropic.F() with direct assignment and anthropic.Opt() for optional parameters

## v1.4.179 (2025-04-21)

### PR [#1432](https://github.com/danielmiessler/Fabric/pull/1432) by [ksylvan](https://github.com/ksylvan): chore: fix fabric setup mess-up introduced by sorting lists (tools and models)

- Chore: alphabetize the order of plugin tools
- Chore: sort AI models alphabetically for consistent listing
- Import `sort` and `strings` packages for sorting functionality
- Sort retrieved AI model names alphabetically, ignoring case
- Add a completion script for fish

## v1.4.178 (2025-04-21)

### PR [#1427](https://github.com/danielmiessler/Fabric/pull/1427) by [ksylvan](https://github.com/ksylvan): Refactor OpenAI-compatible AI providers and add `--listvendors` flag

- Add `--listvendors` command to list all available AI vendors
- Refactor OpenAI-compatible providers into a unified configuration system
- Remove individual vendor packages for streamlined management
- Add sorting functionality for consistent vendor listing output
- Update documentation to include new `--listvendors` option

## v1.4.177 (2025-04-21)

### PR [#1428](https://github.com/danielmiessler/Fabric/pull/1428) by [ksylvan](https://github.com/ksylvan): feat: Alphabetical case-insensitive sorting for groups and items

- Added alphabetical case-insensitive sorting for groups and items in Print method
- Imported `sort` and `strings` packages to enable sorting functionality
- Implemented stable sorting by creating copies of groups and items before sorting
- Enhanced display organization by sorting both groups and their contained items alphabetically
- Improved user experience through consistent case-insensitive alphabetical ordering

## v1.4.176 (2025-04-21)

### PR [#1429](https://github.com/danielmiessler/Fabric/pull/1429) by [ksylvan](https://github.com/ksylvan): feat: enhance StrategyMeta with Prompt field and dynamic naming

- Add `Prompt` field to `StrategyMeta` struct for storing JSON prompt data
- Implement dynamic strategy naming by deriving names from filenames using `strings.TrimSuffix`
- Include `strings` package for enhanced filename processing capabilities

### Direct commits

- Add alphabetical sorting to groups and items in Print method with case-insensitive ordering
- Introduce `--listvendors` command to display all available AI vendors with sorted output
- Refactor OpenAI-compatible providers into unified configuration and remove individual vendor packages
- Import `sort` and `strings` packages to enable sorting functionality across the application
- Update documentation to include the new `--listvendors` option for improved user guidance

## v1.4.175 (2025-04-19)

### PR [#1418](https://github.com/danielmiessler/Fabric/pull/1418) by [dependabot[bot]](https://github.com/apps/dependabot): chore(deps): bump golang.org/x/net from 0.36.0 to 0.38.0 in the go_modules group across 1 directory

- Updated golang.org/x/net dependency from version 0.36.0 to 0.38.0

## v1.4.174 (2025-04-19)

### PR [#1425](https://github.com/danielmiessler/Fabric/pull/1425) by [ksylvan](https://github.com/ksylvan): feat: add Cerebras AI plugin to plugin registry

- Add Cerebras AI plugin to plugin registry
- Introduce Cerebras AI plugin import in plugin registry
- Register Cerebras client in the NewPluginRegistry function

## v1.4.173 (2025-04-18)

### PR [#1420](https://github.com/danielmiessler/Fabric/pull/1420) by [sherif-fanous](https://github.com/sherif-fanous): Fix error in deleting patterns due to non empty directory

- Fix error in deleting patterns due to non empty directory

### PR [#1421](https://github.com/danielmiessler/Fabric/pull/1421) by [ksylvan](https://github.com/ksylvan): feat: add Atom-of-Thought (AoT) strategy and prompt definition

- Add new Atom-of-Thought (AoT) strategy and prompt definition
- Add new aot.json for Atom-of-Thought (AoT) prompting
- Define AoT strategy description and detailed prompt instructions
- Update strategies.json to include AoT in available strategies list
- Ensure AoT strategy appears alongside CoD, CoT, and LTM options

### Direct commits

- Bump golang.org/x/net from 0.36.0 to 0.38.0

## v1.4.172 (2025-04-16)

### PR [#1415](https://github.com/danielmiessler/Fabric/pull/1415) by [ksylvan](https://github.com/ksylvan): feat: add Grok AI provider support

- Add Grok AI provider support to integrate with the Fabric system for AI model interactions
- Add Grok AI client to the plugin registry
- Include Grok AI API key in REST API configuration endpoints
- Update README with documentation about Grok integration

### PR [#1411](https://github.com/danielmiessler/Fabric/pull/1411) by [ksylvan](https://github.com/ksylvan): docs: add contributors section to README with contrib.rocks image

- Add contributors section to README with visual representation using contrib.rocks image

## v1.4.171 (2025-04-15)

### PR [#1407](https://github.com/danielmiessler/Fabric/pull/1407) by [sherif-fanous](https://github.com/sherif-fanous): Update Dockerfile so that Go image version matches go.mod version

- Bump golang version to match go.mod

### Direct commits

- Update README.md

## v1.4.170 (2025-04-13)

### PR [#1406](https://github.com/danielmiessler/Fabric/pull/1406) by [jmd1010](https://github.com/jmd1010): Fix chat history LLM response sequence in ChatInput.svelte

- Fix chat history LLM response sequence in ChatInput.svelte
- Finalize WEB UI V2 loose ends fixes
- Update pattern_descriptions.json

### Direct commits

- Bump golang version to match go.mod

## v1.4.169 (2025-04-11)

### PR [#1403](https://github.com/danielmiessler/Fabric/pull/1403) by [jmd1010](https://github.com/jmd1010): Strategy flag enhancement - Web UI implementation

- Integrate in web ui the strategy flag enhancement first developed in fabric cli
- Update strategies.json

### Direct commits

- Added excalidraw pattern
- Added bill analyzer
- Shorter version of analyze bill
- Updated ed

## v1.4.168 (2025-04-02)

### PR [#1399](https://github.com/danielmiessler/Fabric/pull/1399) by [HaroldFinchIFT](https://github.com/HaroldFinchIFT): feat: add simple optional api key management for protect routes in --serve mode

- Added optional API key management for protecting routes in --serve mode
- Fixed formatting issues
- Refactored API key middleware based on code review feedback

## v1.4.167 (2025-03-31)

### PR [#1397](https://github.com/danielmiessler/Fabric/pull/1397) by [HaroldFinchIFT](https://github.com/HaroldFinchIFT): feat: add it lang to the chat drop down menu lang in web gui

- Feat: add it lang to the chat drop down menu lang in web gui

## v1.4.166 (2025-03-29)

### PR [#1392](https://github.com/danielmiessler/Fabric/pull/1392) by [ksylvan](https://github.com/ksylvan): chore: enhance argument validation in `code_helper` tool

- Refactor: streamline code_helper CLI interface and require explicit instructions
- Require exactly two arguments: directory and instructions
- Remove dedicated help flag, use flag.Usage instead
- Improve directory validation to check if it's a directory
- Inline pattern parsing, removing separate function

### PR [#1390](https://github.com/danielmiessler/Fabric/pull/1390) by [PatrickCLee](https://github.com/PatrickCLee): docs: improve README link

- Fix broken what-and-why link reference

## v1.4.165 (2025-03-26)

### PR [#1389](https://github.com/danielmiessler/Fabric/pull/1389) by [ksylvan](https://github.com/ksylvan): Create Coding Feature

- Feat: add `fabric_code` tool and `create_coding_feature` pattern allowing Fabric to modify existing codebases
- Add file management system for AI-driven code changes with secure file application mechanism
- Fix: improve JSON parsing in ParseFileChanges to handle invalid escape sequences and control characters
- Refactor: rename `fabric_code` tool to `code_helper` for clarity and update all documentation references
- Update chatter to process AI file changes and improve create_coding_feature pattern documentation

### Direct commits

- Docs: improve README link by fixing broken what-and-why link reference

## v1.4.164 (2025-03-22)

### PR [#1380](https://github.com/danielmiessler/Fabric/pull/1380) by [jmd1010](https://github.com/jmd1010): Add flex windows sizing to web interface + raw text input fix

- Add flex windows sizing to web interface
- Fixed processing message not stopping after pattern output completion

### PR [#1379](https://github.com/danielmiessler/Fabric/pull/1379) by [guilhermechapiewski](https://github.com/guilhermechapiewski): Fix typo on fallacies instruction

- Fix typo on fallacies instruction

### PR [#1382](https://github.com/danielmiessler/Fabric/pull/1382) by [ksylvan](https://github.com/ksylvan): docs: improve README formatting and fix some broken links

- Improve README formatting and add clipboard support section
- Fix broken installation link reference and environment variables link
- Improve code block formatting with indentation and clarify package manager alias requirements

### PR [#1376](https://github.com/danielmiessler/Fabric/pull/1376) by [vaygr](https://github.com/vaygr): Add installation instructions for OS package managers

- Add installation instructions for OS package managers

### Direct commits

- Added find_female_life_partner pattern

## v1.4.163 (2025-03-19)

### PR [#1362](https://github.com/danielmiessler/Fabric/pull/1362) by [dependabot[bot]](https://github.com/apps/dependabot): Bump golang.org/x/net from 0.35.0 to 0.36.0 in the go_modules group across 1 directory

- Bump golang.org/x/net from 0.35.0 to 0.36.0 in the go_modules group

### PR [#1372](https://github.com/danielmiessler/Fabric/pull/1372) by [rube-de](https://github.com/rube-de): fix: set percentEncoded to false

- Fix: set percentEncoded to false to prevent YouTube link encoding errors

### PR [#1373](https://github.com/danielmiessler/Fabric/pull/1373) by [ksylvan](https://github.com/ksylvan): Remove unnecessary `system.md` file at top level

- Remove redundant system.md file at top level of the fabric repository

## v1.4.162 (2025-03-19)

### PR [#1374](https://github.com/danielmiessler/Fabric/pull/1374) by [ksylvan](https://github.com/ksylvan): Fix Default Model Change Functionality

- Fix: improve error handling in ChangeDefaultModel flow and save environment file
- Add early return on setup error and save environment file after successful setup
- Maintain proper error propagation

### Direct commits

- Chore: Remove redundant file system.md at top level
- Fix: set percentEncoded to false to prevent YouTube link encoding errors that break fabric functionality

## v1.4.161 (2025-03-17)

### PR [#1363](https://github.com/danielmiessler/Fabric/pull/1363) by [garkpit](https://github.com/garkpit): clipboard operations now work on Mac and PC

- Clipboard operations now work on Mac and PC

## v1.4.160 (2025-03-17)

### PR [#1368](https://github.com/danielmiessler/Fabric/pull/1368) by [vaygr](https://github.com/vaygr): Standardize sections for no repeat guidelines

- Standardize sections for no repeat guidelines

### Direct commits

- Moved system file to proper directory
- Added activity extractor

## v1.4.159 (2025-03-16)

### Direct commits

- Added flashcard generator.

## v1.4.158 (2025-03-16)

### PR [#1367](https://github.com/danielmiessler/Fabric/pull/1367) by [ksylvan](https://github.com/ksylvan): Remove Generic Type Parameters from StorageHandler Initialization

- Refactor: remove generic type parameters from NewStorageHandler calls
- Remove explicit type parameters from StorageHandler initialization
- Update contexts handler constructor implementation
- Update patterns handler constructor implementation
- Update sessions handler constructor implementation

## v1.4.157 (2025-03-16)

### PR [#1365](https://github.com/danielmiessler/Fabric/pull/1365) by [ksylvan](https://github.com/ksylvan): Implement Prompt Strategies in Fabric

- Add prompt strategies like Chain of Thought (CoT) with `--strategy` flag for strategy selection
- Implement `--liststrategies` command to view available strategies and support applying strategies to system prompts
- Improve README with platform-specific installation instructions and fix web interface documentation link
- Refactor git operations with new githelper package and improve error handling in session management
- Fix YouTube configuration check and handling of the installed strategies directory

### Direct commits

- Clipboard operations now work on Mac and PC
- Bump golang.org/x/net from 0.35.0 to 0.36.0 in the go_modules group

## v1.4.156 (2025-03-11)

### PR [#1356](https://github.com/danielmiessler/Fabric/pull/1356) by [ksylvan](https://github.com/ksylvan): chore: add .vscode to `.gitignore` and fix typos and markdown linting  in `Alma.md`

- Add .vscode to `.gitignore` and fix typos and markdown linting in `Alma.md`

### PR [#1352](https://github.com/danielmiessler/Fabric/pull/1352) by [matmilbury](https://github.com/matmilbury): pattern_explanations.md: fix typo

- Fix typo in pattern_explanations.md

### PR [#1354](https://github.com/danielmiessler/Fabric/pull/1354) by [jmd1010](https://github.com/jmd1010): Fix Chat history window scrolling behavior

- Fix chat history window sizing
- Update Web V2 Install Guide with improved instructions

## v1.4.155 (2025-03-09)

### PR [#1350](https://github.com/danielmiessler/Fabric/pull/1350) by [jmd1010](https://github.com/jmd1010): Implement Pattern Tile search functionality

- Implement Pattern Tile search functionality
- Implement  column resize functionnality

## v1.4.154 (2025-03-09)

### PR [#1349](https://github.com/danielmiessler/Fabric/pull/1349) by [ksylvan](https://github.com/ksylvan): Fix: v1.4.153 does not compile because of extra version declaration

- Chore: remove unnecessary `version` variable from `main.go`
- Fix: update Azure client API version access path in tests

### Direct commits

- Implement column resize functionality
- Implement Pattern Tile search functionality

## v1.4.153 (2025-03-08)

### PR [#1348](https://github.com/danielmiessler/Fabric/pull/1348) by [liyuankui](https://github.com/liyuankui): feat: Add LiteLLM AI plugin support with local endpoint configuration

- Feat: Add LiteLLM AI plugin support with local endpoint configuration

## v1.4.152 (2025-03-07)

### Direct commits

- Fix: Fix pipe handling

## v1.4.151 (2025-03-07)

### PR [#1339](https://github.com/danielmiessler/Fabric/pull/1339) by [Eckii24](https://github.com/Eckii24): Feature/add azure api version

- Update azure.go
- Update azure_test.go
- Update openai.go

## v1.4.150 (2025-03-07)

### PR [#1343](https://github.com/danielmiessler/Fabric/pull/1343) by [jmd1010](https://github.com/jmd1010): Rename input.svelte to Input.svelte for proper component naming convention

- Rename input.svelte to Input.svelte for proper component naming convention

## v1.4.149 (2025-03-05)

### PR [#1340](https://github.com/danielmiessler/Fabric/pull/1340) by [ksylvan](https://github.com/ksylvan): Fix for youtube live links plus new youtube_summary pattern

- Update YouTube regex to support live URLs and add timestamped transcript functionality
- Add argument validation to yt command for usage errors and enable -t flag for transcript with timestamps
- Refactor PowerShell yt function with parameter switch and update README for dynamic transcript selection
- Document youtube_summary feature in pattern explanations and introduce new youtube_summary pattern
- Update version

### PR [#1338](https://github.com/danielmiessler/Fabric/pull/1338) by [jmd1010](https://github.com/jmd1010): Update Web V2 Install Guide layout

- Update Web V2 Install Guide layout with improved formatting and structure

### PR [#1330](https://github.com/danielmiessler/Fabric/pull/1330) by [jmd1010](https://github.com/jmd1010): Fixed ALL CAP DIR as requested and processed minor updates to documentation

- Reorganize documentation with consistent directory naming and updated installation guides

### PR [#1333](https://github.com/danielmiessler/Fabric/pull/1333) by [asasidh](https://github.com/asasidh): Update QUOTES section to include speaker names for clarity

- Update QUOTES section to include speaker names for improved clarity

### Direct commits

- Update Azure and OpenAI Go modules with bug fixes and improvements

## v1.4.148 (2025-03-03)

- Fix: Rework LM Studio plugin
- Update QUOTES section to include speaker names for clarity
- Update Web V2 Install Guide with improved instructions V2
- Update Web V2 Install Guide with improved instructions
- Reorganize documentation with consistent directory naming and updated guides

## v1.4.147 (2025-02-28)

### PR [#1326](https://github.com/danielmiessler/Fabric/pull/1326) by [pavdmyt](https://github.com/pavdmyt): fix: continue fetching models even if some vendors fail

- Fix: continue fetching models even if some vendors fail by removing cancellation of remaining goroutines when a vendor collection fails
- Ensure other vendor collections continue even if one fails
- Fix listing models via `fabric -L` and using non-default models via `fabric -m custom_model` when localhost models are not listening

### PR [#1329](https://github.com/danielmiessler/Fabric/pull/1329) by [jmd1010](https://github.com/jmd1010): Svelte Web V2 Installation Guide

- Add Web V2 Installation Guide
- Update install guide with Plain Text instructions

## v1.4.146 (2025-02-27)

### PR [#1319](https://github.com/danielmiessler/Fabric/pull/1319) by [jmd1010](https://github.com/jmd1010): Enhancement: PDF to Markdown Conversion Functionality to the Web Svelte Chat Interface

- Add PDF to Markdown conversion functionality to the web svelte chat interface
- Add PDF to Markdown integration documentation
- Add Svelte implementation files for PDF integration
- Update README files directory structure and naming convention
- Add required UI image assets for feature implementation

## v1.4.145 (2025-02-26)

### PR [#1324](https://github.com/danielmiessler/Fabric/pull/1324) by [jaredmontoya](https://github.com/jaredmontoya): flake: fix/update and enhance

- Flake: fix/update

## v1.4.144 (2025-02-26)

### Direct commits

- Upgrade upload artifacts to v4

## v1.4.143 (2025-02-26)

### PR [#1264](https://github.com/danielmiessler/Fabric/pull/1264) by [eugeis](https://github.com/eugeis): feat: implement support for exolab

- Feat: implement support for <https://github.com/exo-explore/exo>
- Merge branch 'main' into feat/exolab

## v1.4.142 (2025-02-25)

### Direct commits

- Fix: build problems

## v1.4.141 (2025-02-25)

### PR [#1260](https://github.com/danielmiessler/Fabric/pull/1260) by [bluPhy](https://github.com/bluPhy): Fixing typo

- Typos correction
- Update version to v1.4.80 and commit

## v1.4.140 (2025-02-25)

### PR [#1313](https://github.com/danielmiessler/Fabric/pull/1313) by [cx-ken-swain](https://github.com/cx-ken-swain): Updated ollama.go to fix a couple of potential DoS issues

- Updated ollama.go to fix security issues and resolve potential DoS vulnerabilities
- Resolved additional medium severity vulnerabilities in the codebase
- Updated application version and committed changes
- Cleaned up version-related files including pkgs/fabric/version.nix and version.go

## v1.4.139 (2025-02-25)

### PR [#1321](https://github.com/danielmiessler/Fabric/pull/1321) by [jmd1010](https://github.com/jmd1010): Update demo video link in PR-1309 documentation

- Update demo video link in PR-1284 documentation

### Direct commits

- Add complete PDF to Markdown documentation
- Add Svelte implementation files for PDF integration
- Add PDF to Markdown integration documentation
- Add PDF to Markdown conversion functionality to the web svelte chat interface
- Update version to v..1 and commit

## v1.4.138 (2025-02-24)

### PR [#1317](https://github.com/danielmiessler/Fabric/pull/1317) by [ksylvan](https://github.com/ksylvan): chore: update Anthropic SDK and add Claude 3.7 Sonnet model support

- Updated anthropic-sdk-go from v0.2.0-alpha.4 to v0.2.0-alpha.11
- Added Claude 3.7 Sonnet models to available model list
- Added ModelClaude3_7SonnetLatest to model options
- Added ModelClaude3_7Sonnet20250219 to model options
- Removed ModelClaude_Instant_1_2 from available models

## v1.4.80 (2025-02-24)

### Direct commits

- Feat: impl. multi-model / attachments, images

## v1.4.79 (2025-02-24)

### PR [#1257](https://github.com/danielmiessler/Fabric/pull/1257) by [jessefmoore](https://github.com/jessefmoore): Create analyze_threat_report_cmds

- Create system.md pattern to extract commands from videos and threat reports for pentesters, red teams, and threat hunters to simulate threat actors

### PR [#1256](https://github.com/danielmiessler/Fabric/pull/1256) by [JOduMonT](https://github.com/JOduMonT): Update README.md

- Update README.md with Windows Command improvements and syntax enhancements for easier copy-paste functionality

### PR [#1247](https://github.com/danielmiessler/Fabric/pull/1247) by [kevnk](https://github.com/kevnk): Update suggest_pattern: refine summaries and add recently added patterns

- Update summaries and add recently added patterns to suggest_pattern

### PR [#1252](https://github.com/danielmiessler/Fabric/pull/1252) by [jeffmcjunkin](https://github.com/jeffmcjunkin): Update README.md: Add PowerShell aliases

- Add PowerShell aliases to README.md

### PR [#1253](https://github.com/danielmiessler/Fabric/pull/1253) by [abassel](https://github.com/abassel): Fixed few typos that I could find

- Fixed multiple typos throughout the codebase

## v1.4.137 (2025-02-24)

### PR [#1296](https://github.com/danielmiessler/Fabric/pull/1296) by [dependabot[bot]](https://github.com/apps/dependabot): Bump github.com/go-git/go-git/v5 from 5.12.0 to 5.13.0 in the go_modules group across 1 directory

- Updated github.com/go-git/go-git/v5 dependency from version 5.12.0 to 5.13.0

## v1.4.136 (2025-02-24)

- Update to upload-artifact@v4 because upload-artifact@v3 is deprecated
- Merge branch 'danielmiessler:main' into main
- Updated anthropic-sdk-go from v0.2.0-alpha.4 to v0.2.0-alpha.11
- Added Claude 3.7 Sonnet models to available model list
- Removed ModelClaude_Instant_1_2 from available models

## v1.4.135 (2025-02-24)

### PR [#1309](https://github.com/danielmiessler/Fabric/pull/1309) by [jmd1010](https://github.com/jmd1010): Feature/Web Svelte GUI Enhancements: Pattern Descriptions, Tags, Favorites, Search Bar, Language Integration, PDF file conversion, etc

- Enhanced pattern handling and chat interface improvements
- Updated .gitignore to exclude sensitive and generated files
- Setup backup configuration and update dependencies

### PR [#1312](https://github.com/danielmiessler/Fabric/pull/1312) by [junaid18183](https://github.com/junaid18183): Added Create LOE Document Prompt

- Added create_loe_document prompt

### PR [#1302](https://github.com/danielmiessler/Fabric/pull/1302) by [verebes1](https://github.com/verebes1): feat: Add LM Studio compatibility

- Added LM Studio as a new plugin, now it can be used with Fabric
- Updated the plugin registry with the new plugin name

### PR [#1297](https://github.com/danielmiessler/Fabric/pull/1297) by [Perchycs](https://github.com/Perchycs): Create pattern_explanations.md

- Create pattern_explanations.md

### Direct commits

- Added extract_domains functionality
- Resolved security vulnerabilities in ollama.go

## v1.4.134 (2025-02-11)

### PR [#1289](https://github.com/danielmiessler/Fabric/pull/1289) by [thevops](https://github.com/thevops): Add the ability to grab YouTube video transcript with timestamps

- Add the ability to grab YouTube video transcript with timestamps using the new `--transcript-with-timestamps` flag
- Format timestamps as HH:MM:SS and prepend them to each line of the transcript
- Enable quick navigation to specific parts of videos when creating summaries

## v1.4.133 (2025-02-11)

### PR [#1294](https://github.com/danielmiessler/Fabric/pull/1294) by [TvisharajiK](https://github.com/TvisharajiK): Improved unit-test coverage from 0 to 100 (AI module) using Keploy's agent

- Feat: Increase unit test coverage from 0 to 100% in the AI module using Keploy's Agent

### Direct commits

- Bump github.com/go-git/go-git/v5 from 5.12.0 to 5.13.0 in the go_modules group
- Add the ability to grab YouTube video transcript with timestamps using the new `--transcript-with-timestamps` flag
- Added multiple TELOS patterns including h3 TELOS pattern, challenge handling pattern, year in review pattern, and additional Telos patterns
- Added panel topic extractor for improved content analysis
- Added intro sentences pattern for better content structuring

## v1.4.132 (2025-02-02)

### PR [#1278](https://github.com/danielmiessler/Fabric/pull/1278) by [aicharles](https://github.com/aicharles): feat(anthropic): enable custom API base URL support

- Enable custom API base URL configuration for Anthropic integration
- Add proper handling of v1 endpoint for UUID-containing URLs
- Implement URL formatting logic for consistent endpoint structure
- Clean up commented code and improve configuration flow

## v1.4.131 (2025-01-30)

### PR [#1270](https://github.com/danielmiessler/Fabric/pull/1270) by [wmahfoudh](https://github.com/wmahfoudh): Added output filename support for to_pdf

- Added output filename support for to_pdf

### PR [#1271](https://github.com/danielmiessler/Fabric/pull/1271) by [wmahfoudh](https://github.com/wmahfoudh): Adding deepseek support

- Feat: Added Deepseek AI integration

### PR [#1258](https://github.com/danielmiessler/Fabric/pull/1258) by [tuergeist](https://github.com/tuergeist): Minor README fix and additional Example

- Doc: Custom patterns also work with Claude models
- Doc: Add scrape URL example. Fix Example 4

### Direct commits

- Feat: implement support for <https://github.com/exo-explore/exo>

## v1.4.130 (2025-01-03)

### PR [#1240](https://github.com/danielmiessler/Fabric/pull/1240) by [johnconnor-sec](https://github.com/johnconnor-sec): Updates: ./web

- Moved pattern loader to ModelConfig and added page fly transitions with improved responsive layout
- Updated UI components and chat layout display with reordered columns and improved Header buttons
- Added NotesDrawer component to header that saves notes to lib/content/inbox
- Centered chat interface in viewport and improved Post page styling and layout
- Updated project structure by moving and renaming components from lib/types to lib/interfaces and lib/api

## v1.4.129 (2025-01-03)

### PR [#1242](https://github.com/danielmiessler/Fabric/pull/1242) by [CuriouslyCory](https://github.com/CuriouslyCory): Adding youtube --metadata flag

- Added metadata lookup to youtube helper
- Better metadata

### PR [#1230](https://github.com/danielmiessler/Fabric/pull/1230) by [iqbalabd](https://github.com/iqbalabd): Update translate pattern to use curly braces

- Update translate pattern to use curly braces

### Direct commits

- Added enrich_blog_post pattern for enhanced blog post processing
- Enhanced enrich pattern with improved functionality
- Centered chat and note drawer components in viewport for better user experience
- Updated post page styling and layout with improved visual design
- Added templates for posts and improved content management structure

## v1.4.128 (2024-12-26)

### PR [#1227](https://github.com/danielmiessler/Fabric/pull/1227) by [mattjoyce](https://github.com/mattjoyce): Feature/template extensions

- Implemented stdout template extensions with path-based registry storage and proper hash verification for both configs and executables
- Successfully implemented file-based output handling with clean interface requiring only path output and proper cleanup of temporary files
- Fixed pattern file usage without stdin by initializing empty message when Message is nil, allowing patterns like `./fabric -p pattern.txt -v=name:value` to work without requiring stdin input
- Added comprehensive tests for extension manager, registration and execution with validation for extension names and timeout values
- Enhanced extension functionality with example files, tutorial documentation, and improved error handling for hash verification failures

### Direct commits

- Updated story to be shorter bullets and improved formatting
- Updated POSTS to make main 24-12-08 and refreshed imports
- WIP: Notes Drawer text color improvements and updated default theme to rocket

## v1.4.127 (2024-12-23)

### PR [#1218](https://github.com/danielmiessler/Fabric/pull/1218) by [sosacrazy126](https://github.com/sosacrazy126): streamlit ui

- Add Streamlit application for managing and executing patterns with comprehensive pattern creation, execution, and analysis capabilities
- Refactor pattern management and enhance error handling with improved logging configuration for better debugging and user feedback
- Improve pattern creation, editing, and deletion functionalities with streamlined session state initialization for enhanced performance
- Update input validation and sanitization processes to ensure safe pattern processing
- Add new UI components for better user experience in pattern management and output analysis

### PR [#1225](https://github.com/danielmiessler/Fabric/pull/1225) by [wmahfoudh](https://github.com/wmahfoudh): Added Humanize Pattern

- Added Humanize Pattern

## v1.4.126 (2024-12-22)

### PR [#1212](https://github.com/danielmiessler/Fabric/pull/1212) by [wrochow](https://github.com/wrochow): Significant updates to Duke and Socrates

- Significant thematic rewrite incorporating classical philosophical texts including Plato's Apology, Phaedrus, Symposium, and The Republic, plus Xenophon's works on Socrates
- Added specific steps for research, analysis, and code reviews
- Updated version to v1.1 with associated code changes

## v1.4.125 (2024-12-22)

### PR [#1222](https://github.com/danielmiessler/Fabric/pull/1222) by [wmahfoudh](https://github.com/wmahfoudh): Fix cross-filesystem file move in to_pdf plugin (issue 1221)

- Fix cross-filesystem file move in to_pdf plugin (issue 1221)

### Direct commits

- Update version to v..1 and commit

## v1.4.124 (2024-12-21)

### PR [#1215](https://github.com/danielmiessler/Fabric/pull/1215) by [infosecwatchman](https://github.com/infosecwatchman): Add Endpoints to facilitate Ollama based chats

- Add Endpoints to facilitate Ollama based chats

### PR [#1214](https://github.com/danielmiessler/Fabric/pull/1214) by [iliaross](https://github.com/iliaross): Fix the typo in the sentence

- Fix the typo in the sentence

### PR [#1213](https://github.com/danielmiessler/Fabric/pull/1213) by [AnirudhG07](https://github.com/AnirudhG07): Spelling Fixes

- Spelling fixes in patterns

- Refactor pattern management and enhance error handling
- Improved pattern creation, editing, and deletion functionalities

## v1.4.123 (2024-12-20)

### PR [#1208](https://github.com/danielmiessler/Fabric/pull/1208) by [mattjoyce](https://github.com/mattjoyce): Fix: Issue with the custom message and added example config file

- Fix: Issue with the custom message and added example config file

### Direct commits

- Add comprehensive Streamlit application for managing and executing patterns with pattern creation, execution, analysis, and robust logging capabilities
- Add endpoints to facilitate Ollama based chats for integration with Open WebUI
- Significant thematic rewrite incorporating Socratic interaction themes from classical texts including Plato's Apology, Phaedrus, Symposium, and The Republic
- Add XML-based Markdown converter pattern for improved document processing
- Update version to v1.1 and fix various spelling errors across patterns and documentation

## v1.4.122 (2024-12-14)

### PR [#1201](https://github.com/danielmiessler/Fabric/pull/1201) by [mattjoyce](https://github.com/mattjoyce): feat: Add YAML configuration support

- Add support for persistent configuration via YAML files with ability to override using CLI flags
- Add --config flag for specifying YAML configuration file path
- Implement standard option precedence system (CLI > YAML > defaults)
- Add type-safe YAML parsing with reflection for robust configuration handling
- Add comprehensive tests for YAML configuration functionality

## v1.4.121 (2024-12-13)

### PR [#1200](https://github.com/danielmiessler/Fabric/pull/1200) by [mattjoyce](https://github.com/mattjoyce): Fix: Mask input token to prevent var substitution in patterns

- Fix: Mask input token to prevent var substitution in patterns

### Direct commits

- Added new instruction trick.

## v1.4.120 (2024-12-10)

### PR [#1189](https://github.com/danielmiessler/Fabric/pull/1189) by [mattjoyce](https://github.com/mattjoyce): Add --input-has-vars flag to control variable substitution in input

- Add --input-has-vars flag to control variable substitution in input
- Add InputHasVars field to ChatRequest struct
- Only process template variables in user input when flag is set
- Fixes issue with Ansible/Jekyll templates that use {{var}} syntax

### PR [#1182](https://github.com/danielmiessler/Fabric/pull/1182) by [jessefmoore](https://github.com/jessefmoore): analyze_risk pattern

- Created a pattern to analyze 3rd party vendor risk

## v1.4.119 (2024-12-07)

### PR [#1181](https://github.com/danielmiessler/Fabric/pull/1181) by [mattjoyce](https://github.com/mattjoyce): Bugfix/1169 symlinks

- Fix #1169: Add robust handling for paths and symlinks in GetAbsolutePath

### Direct commits

- Added tutorial with example files
- Add cards component
- Update: packages, main page, styles
- Check extension names don't have spaces
- Added test pattern

## v1.4.118 (2024-12-05)

### PR [#1174](https://github.com/danielmiessler/Fabric/pull/1174) by [mattjoyce](https://github.com/mattjoyce): Curly brace templates

- Fix pattern file usage without stdin by initializing empty message when Message is nil, allowing patterns to work with variables but no stdin input
- Remove redundant template processing of message content and let pattern processing handle all template resolution
- Simplify template processing flow while supporting both stdin and non-stdin use cases

### PR [#1179](https://github.com/danielmiessler/Fabric/pull/1179) by [sluosapher](https://github.com/sluosapher): added a new pattern create_newsletter_entry

- Added a new pattern create_newsletter_entry

### Direct commits

- Update @sveltejs/kit dependency from version 2.8.4 to 2.9.0 in web directory
- Implement extension registry refinement with path-based storage and proper hash verification for configurations and executables
- Add file-based output implementation with clean interface and proper cleanup of temporary files

## v1.4.117 (2024-11-30)

### Direct commits

- Fix: close #1173

## v1.4.116 (2024-11-28)

### Direct commits

- Chore: cleanup style

## v1.4.115 (2024-11-28)

### PR [#1168](https://github.com/danielmiessler/Fabric/pull/1168) by [johnconnor-sec](https://github.com/johnconnor-sec): Update README.md

- Update README.md

### Direct commits

- Chore: cleanup style
- Updated readme
- Fix: use the custom message and then piped one

## v1.4.114 (2024-11-26)

### PR [#1164](https://github.com/danielmiessler/Fabric/pull/1164) by [MegaGrindStone](https://github.com/MegaGrindStone): fix: provide default message content to avoid nil pointer dereference

- Fix: provide default message content to avoid nil pointer dereference

## v1.4.113 (2024-11-26)

### PR [#1166](https://github.com/danielmiessler/Fabric/pull/1166) by [dependabot[bot]](https://github.com/apps/dependabot): build(deps-dev): bump @sveltejs/kit from 2.6.1 to 2.8.4 in /web in the npm_and_yarn group across 1 directory

- Updated @sveltejs/kit dependency from version 2.6.1 to 2.8.4 in the web directory

## v1.4.112 (2024-11-26)

### PR [#1165](https://github.com/danielmiessler/Fabric/pull/1165) by [johnconnor-sec](https://github.com/johnconnor-sec): feat: Fabric Web UI

- Added new Fabric Web UI feature
- Updated version to v1.1 and committed changes
- Updated Obsidian.md documentation
- Updated README.md with new information

### Direct commits

- Fixed nil pointer dereference by providing default message content

## v1.4.111 (2024-11-26)

### Direct commits

- Ci: Integrate code formating

## v1.4.110 (2024-11-26)

### PR [#1135](https://github.com/danielmiessler/Fabric/pull/1135) by [mrtnrdl](https://github.com/mrtnrdl): Add `extract_recipe`

- Update version to v..1 and commit
- Add extract_recipe to easily extract the necessary information from cooking-videos
- Merge branch 'main' into main

## v1.4.109 (2024-11-24)

### PR [#1157](https://github.com/danielmiessler/Fabric/pull/1157) by [mattjoyce](https://github.com/mattjoyce): fix: process template variables in raw input

- Fix: process template variables in raw input - Process template variables ({{var}}) consistently in both pattern files and raw input messages, as variables were previously only processed when using pattern files
- Add template variable processing for raw input in BuildSession with explicit messageContent initialization
- Remove errantly committed build artifact (fabric binary from previous commit)
- Fix template.go to handle missing variables in stdin input with proper error messaging
- Fix raw mode doubling user input issue by streamlining context staging since input is now already embedded in pattern

### Direct commits

- Added analyze_mistakes

## v1.4.108 (2024-11-21)

### PR [#1155](https://github.com/danielmiessler/Fabric/pull/1155) by [mattjoyce](https://github.com/mattjoyce): Curly brace templates and plugins

- Introduced new template package for variable substitution with {{variable}} syntax
- Moved substitution logic from patterns to centralized template system for better organization
- Updated patterns.go to use template package for variable processing with special {{input}} handling
- Implemented core plugin system with utility plugins including datetime, fetch, file, sys, and text operations
- Added comprehensive test coverage and markdown documentation for all plugins

## v1.4.107 (2024-11-19)

### PR [#1149](https://github.com/danielmiessler/Fabric/pull/1149) by [mathisto](https://github.com/mathisto): Fix typo in md_callout

- Fix typo in md_callout pattern

### Direct commits

- Update patterns zip workflow in CI
- Remove patterns zip workflow from CI

## v1.4.106 (2024-11-19)

### Direct commits

- Feat: migrate to official anthropics Go SDK

## v1.4.105 (2024-11-19)

### PR [#1147](https://github.com/danielmiessler/Fabric/pull/1147) by [mattjoyce](https://github.com/mattjoyce): refactor: unify pattern loading and variable handling

- Refactored pattern loading and variable handling to improve separation of concerns between chatter.go and patterns.go
- Consolidated pattern loading logic into unified GetPattern method supporting both file and database patterns
- Implemented single interface for pattern handling while maintaining API compatibility with Storage interface
- Centralized variable substitution processing to maintain backward compatibility for REST API
- Enhanced pattern handling architecture while preserving existing interfaces and adding file-based pattern support

### PR [#1146](https://github.com/danielmiessler/Fabric/pull/1146) by [mrwadams](https://github.com/mrwadams): Add summarize_meeting

- Added new summarize_meeting pattern for creating meeting summaries from audio transcripts with structured output including Key Points, Tasks, Decisions, and Next Steps sections

### Direct commits

- Introduced new template package for variable substitution with {{variable}} syntax and centralized substitution logic
- Updated patterns.go to use template package for variable processing with special {{input}} handling for pattern content
- Enhanced chatter.go and REST API to support input parameter passing and multiple passes for nested variables
- Implemented error reporting for missing required variables to establish foundation for future templating features

## v1.4.104 (2024-11-18)

### PR [#1142](https://github.com/danielmiessler/Fabric/pull/1142) by [mattjoyce](https://github.com/mattjoyce): feat: add file-based pattern support

- Add file-based pattern support allowing patterns to be loaded directly from files using explicit path prefixes (~/, ./, /, or \)
- Support relative paths (./pattern.txt, ../pattern.txt) and home directory expansion (~/patterns/test.txt)
- Support absolute paths while maintaining backwards compatibility with named patterns
- Require explicit path markers to distinguish from pattern names

### Direct commits

- Add summarize_meeting pattern to create meeting summaries from audio transcripts with sections for Key Points, Tasks, Decisions, and Next Steps

## v1.4.103 (2024-11-18)

### PR [#1133](https://github.com/danielmiessler/Fabric/pull/1133) by [igophper](https://github.com/igophper): fix: fix default gin

- Fix: fix default gin

### PR [#1129](https://github.com/danielmiessler/Fabric/pull/1129) by [xyb](https://github.com/xyb): add a screenshot of fabric

- Add a screenshot of fabric

## v1.4.102 (2024-11-18)

### PR [#1143](https://github.com/danielmiessler/Fabric/pull/1143) by [mariozig](https://github.com/mariozig): Update docker image

- Update docker image

### Direct commits

- Add file-based pattern support allowing patterns to be loaded directly from files using explicit path prefixes (~/, ./, /, or \)
- Support relative paths (./pattern.txt, ../pattern.txt) for easier pattern testing and iteration
- Support home directory expansion (~/patterns/test.txt) for user-specific pattern locations
- Support absolute paths for system-wide pattern access
- Maintain backwards compatibility with existing named patterns while requiring explicit path markers to distinguish from pattern names

## v1.4.101 (2024-11-15)

### Direct commits

- Improve logging for missing setup steps
- Add extract_recipe to easily extract the necessary information from cooking-videos
- Fix: fix default gin
- Update version to v..1 and commit
- Add a screenshot of fabric

## v1.4.100 (2024-11-13)

- Added our first formal stitch.
- Upgraded AI result rater.

## v1.4.99 (2024-11-10)

### PR [#1126](https://github.com/danielmiessler/Fabric/pull/1126) by [jaredmontoya](https://github.com/jaredmontoya): flake: add gomod2nix auto-update

- Flake: add gomod2nix auto-update

### Direct commits

- Upgraded AI result rater

## v1.4.98 (2024-11-09)

### Direct commits

- Ci: zip patterns

## v1.4.97 (2024-11-09)

### Direct commits

- Feat: update dependencies; improve vendors setup/default model

## v1.4.96 (2024-11-09)

### PR [#1060](https://github.com/danielmiessler/Fabric/pull/1060) by [noamsiegel](https://github.com/noamsiegel): Analyze Candidates Pattern

- Added system and user prompts

### Direct commits

- Feat: add claude-3-5-haiku-latest model

## v1.4.95 (2024-11-09)

### PR [#1123](https://github.com/danielmiessler/Fabric/pull/1123) by [polyglotdev](https://github.com/polyglotdev): :sparkles: Added unaliasing to pattern setup

- Added unaliasing functionality to pattern setup process to prevent conflicts between dynamically defined functions and pre-existing aliases

### PR [#1119](https://github.com/danielmiessler/Fabric/pull/1119) by [verebes1](https://github.com/verebes1): Add auto save functionality

- Added auto save functionality to aliases for integration with tools like Obsidian
- Updated README with information about autogenerating aliases that support auto-saving features
- Updated table of contents in documentation

### Direct commits

- Updated README documentation
- Created Selemela07 devcontainer.json configuration file

## v1.4.94 (2024-11-06)

### PR [#1108](https://github.com/danielmiessler/Fabric/pull/1108) by [butterflyx](https://github.com/butterflyx): [add] RegEx for YT shorts

- Added VideoID support for YouTube shorts

### PR [#1117](https://github.com/danielmiessler/Fabric/pull/1117) by [verebes1](https://github.com/verebes1): Add alias generation information

- Added alias generation information to README including YouTube transcript aliases
- Updated table of contents

### PR [#1115](https://github.com/danielmiessler/Fabric/pull/1115) by [ignacio-arce](https://github.com/ignacio-arce): Added create_diy

- Added create_diy functionality

## v1.4.93 (2024-11-06)

## PR #123: Fix YouTube URL Pattern and Add Alias Generation

- Fix: short YouTube URL pattern
- Add alias generation information
- Updated the readme with information about generating aliases for each prompt including one for YouTube transcripts
- Updated the table of contents
- Added create_diy feature
- [add] VideoID for YT shorts

## v1.4.92 (2024-11-05)

### PR [#1109](https://github.com/danielmiessler/Fabric/pull/1109) by [leonsgithub](https://github.com/leonsgithub): Add docker

- Add docker

## v1.4.91 (2024-11-05)

### Direct commits

- Fix: bufio.Scanner message too long
- Add docker

## v1.4.90 (2024-11-04)

### Direct commits

- Feat: impl. Youtube PlayList support
- Fix: close #1103, Update Readme hpt to install to_pdf

## v1.4.89 (2024-11-04)

### PR [#1102](https://github.com/danielmiessler/Fabric/pull/1102) by [jholsgrove](https://github.com/jholsgrove): Create user story pattern

- Create user story pattern

### Direct commits

- Fix: close #1106, fix pipe reading
- Feat: YouTube PlayList support

## v1.4.88 (2024-10-30)

### PR [#1098](https://github.com/danielmiessler/Fabric/pull/1098) by [jaredmontoya](https://github.com/jaredmontoya): Fix nix package update workflow

- Fix nix package version auto update workflow

## v1.4.87 (2024-10-30)

### PR [#1096](https://github.com/danielmiessler/Fabric/pull/1096) by [jaredmontoya](https://github.com/jaredmontoya): Implement automated ci nix package version update

- Modularize nix flake
- Automate nix package version update

## v1.4.86 (2024-10-30)

### PR [#1088](https://github.com/danielmiessler/Fabric/pull/1088) by [jaredmontoya](https://github.com/jaredmontoya): feat: add DEFAULT_CONTEXT_LENGTH setting

- Add model context length setting

## v1.4.85 (2024-10-30)

### Direct commits

- Feat: write tools output also to output file if defined; fix XouTube transcript &#39; character

## v1.4.84 (2024-10-30)

### Direct commits

- Ci: deactivate build triggering at changes of patterns or docu

## v1.4.83 (2024-10-30)

### PR [#1089](https://github.com/danielmiessler/Fabric/pull/1089) by [jaredmontoya](https://github.com/jaredmontoya): Introduce Nix to the project

- Add trailing newline
- Add Nix Flake

## v1.4.82 (2024-10-30)

### PR [#1094](https://github.com/danielmiessler/Fabric/pull/1094) by [joshmedeski](https://github.com/joshmedeski): feat: add md_callout pattern

- Feat: add md_callout pattern
Add a pattern that can convert text into an appropriate markdown callout

## v1.4.81 (2024-10-29)

### Direct commits

- Feat: split tools messages from use message

## v1.4.78 (2024-10-28)

### PR [#1059](https://github.com/danielmiessler/Fabric/pull/1059) by [noamsiegel](https://github.com/noamsiegel): Analyze Proposition Pattern

- Added system and user prompts

## v1.4.77 (2024-10-28)

### PR [#1073](https://github.com/danielmiessler/Fabric/pull/1073) by [mattjoyce](https://github.com/mattjoyce): Five patterns to explore a project, opportunity or brief

- Added five new DSRP (Distinctions, Systems, Relationships, Perspectives) patterns for project exploration with enhanced divergent thinking capabilities
- Implemented identify_job_stories pattern for user story identification and analysis
- Created S7 Strategy profiling pattern with structured approach for strategic analysis
- Added headwinds and tailwinds analysis functionality for comprehensive project assessment
- Enhanced all DSRP prompts with improved metadata and style guide compliance

### Direct commits

- Add Nix Flake

## v1.4.76 (2024-10-28)

### Direct commits

- Chore: simplify isChatRequest

## v1.4.75 (2024-10-28)

### PR [#1090](https://github.com/danielmiessler/Fabric/pull/1090) by [wrochow](https://github.com/wrochow): A couple of patterns

- Added "Dialog with Socrates" pattern for engaging in deep, meaningful conversations with a modern day philosopher
- Added "Ask uncle Duke" pattern for Java software development expertise, particularly with Spring Framework and Maven

### Direct commits

- Add trailing newline

## v1.4.74 (2024-10-27)

### PR [#1077](https://github.com/danielmiessler/Fabric/pull/1077) by [xvnpw](https://github.com/xvnpw): feat: add pattern refine_design_document

- Feat: add pattern refine_design_document

## v1.4.73 (2024-10-27)

### PR [#1086](https://github.com/danielmiessler/Fabric/pull/1086) by [NuCl34R](https://github.com/NuCl34R): Create a basic translator pattern, edit file to add desired language

- Create system.md

### Direct commits

- Added metadata and styleguide
- Added structure to prompt
- Added headwinds and tailwinds
- Initial draft of s7 Strategy profiling

## v1.4.72 (2024-10-25)

### PR [#1070](https://github.com/danielmiessler/Fabric/pull/1070) by [xvnpw](https://github.com/xvnpw): feat: create create_design_document pattern

- Feat: create create_design_document pattern

## v1.4.71 (2024-10-25)

### PR [#1072](https://github.com/danielmiessler/Fabric/pull/1072) by [xvnpw](https://github.com/xvnpw): feat: add review_design pattern

- Feat: add review_design pattern

## v1.4.70 (2024-10-25)

### PR [#1064](https://github.com/danielmiessler/Fabric/pull/1064) by [rprouse](https://github.com/rprouse): Update README.md with pbpaste section

- Update README.md with pbpaste section

### Direct commits

- Added new pattern: refine_design_document for improving design documentation
- Added identify_job_stories pattern for user story identification
- Added review_design pattern for design review processes
- Added create_design_document pattern for generating design documentation
- Added system and user prompts for enhanced functionality

## v1.4.69 (2024-10-21)

### Direct commits

- Updated the Alma.md file.

## v1.4.68 (2024-10-21)

### Direct commits

- Fix: setup does not overwrites old values

## v1.4.67 (2024-10-19)

### Direct commits

- Merge remote-tracking branch 'origin/main'
- Feat: plugins arch., new setup procedure

## v1.4.66 (2024-10-19)

### Direct commits

- Feat: plugins arch., new setup procedure

## v1.4.65 (2024-10-16)

### PR [#1045](https://github.com/danielmiessler/Fabric/pull/1045) by [Fenicio](https://github.com/Fenicio): Update patterns/analyze_answers/system.md - Fixed a bunch of typos

- Update patterns/analyze_answers/system.md - Fixed a bunch of typos

## v1.4.64 (2024-10-14)

### Direct commits

- Updated readme

## v1.4.63 (2024-10-13)

### PR [#862](https://github.com/danielmiessler/Fabric/pull/862) by [Thepathakarpit](https://github.com/Thepathakarpit): Create setup_fabric.bat, a batch script to automate setup and running…

- Create setup_fabric.bat, a batch script to automate setup and running fabric on windows.
- Merge branch 'main' into patch-1

## v1.4.62 (2024-10-13)

### PR [#1044](https://github.com/danielmiessler/Fabric/pull/1044) by [eugeis](https://github.com/eugeis): Feat/rest api

- Feat: work on Rest API
- Feat: restructure for better reuse
- Merge branch 'main' into feat/rest-api

## v1.4.61 (2024-10-13)

### Direct commits

- Updated extract sponsors.
- Merge branch 'main' into feat/rest-api
- Feat: restructure for better reuse
- Feat: restructure for better reuse
- Feat: restructure for better reuse

## v1.4.60 (2024-10-12)

### Direct commits

- Fix: IsChatRequest rule; Close #1042 is

## v1.4.59 (2024-10-11)

### Direct commits

- Added ctw to Raycast.

## v1.4.58 (2024-10-11)

### Direct commits

- Chore: we don't need tp configure DryRun vendor
- Fix: Close #1040. Configure vendors separately that were not configured yet

## v1.4.57 (2024-10-11)

### Direct commits

- Docs: Close #1035, provide better example for pattern variables

## v1.4.56 (2024-10-11)

### PR [#1039](https://github.com/danielmiessler/Fabric/pull/1039) by [hallelujah-shih](https://github.com/hallelujah-shih): Feature/set default lang

- Support set default output language

### Direct commits

- Updated all dsrp prompts to increase divergent thinking
- Fixed mix up with system
- Initial dsrp prompts

## v1.4.55 (2024-10-09)

### Direct commits

- Fix: Close #1036

## v1.4.54 (2024-10-07)

### PR [#1021](https://github.com/danielmiessler/Fabric/pull/1021) by [joshuafuller](https://github.com/joshuafuller): Corrected spelling and grammatical errors for consistency and clarity for transcribe_minutes

- Fixed spelling errors including "highliting" to "highlighting" and "exxactly" to "exactly"
- Improved grammatical accuracy by changing "agreed within the meeting" to "agreed upon within the meeting"
- Added missing periods to ensure consistency across list items
- Updated phrasing from "Write NEXT STEPS a 2-3 sentences" to "Write NEXT STEPS as 2-3 sentences" for grammatical correctness
- Enhanced overall readability and consistency of the transcribe_minutes document

## v1.4.53 (2024-10-07)

### Direct commits

- Fix: fix NP if response is empty, close #1026, #1027

## v1.4.52 (2024-10-06)

### Direct commits

- Added extract_core_message functionality
- Feat: Enhanced Rest API development with multiple improvements
- Corrected spelling and grammatical errors for consistency and clarity, including fixes to "agreed upon within the meeting", "highlighting", "exactly", and "Write NEXT STEPS as 2-3 sentences"
- Merged latest changes from main branch

## v1.4.51 (2024-10-05)

### Direct commits

- Fix: tests

## v1.4.50 (2024-10-05)

### Direct commits

- Fix: windows release

## v1.4.49 (2024-10-05)

### Direct commits

- Fix: windows release

## v1.4.48 (2024-10-05)

### Direct commits

- Feat: Add 'meta' role to store meta info to session, like source of input content.

## v1.4.47 (2024-10-05)

### Direct commits

- Feat: Add 'meta' role to store meta info to session, like source of input content.
- Feat: Add 'meta' role to store meta info to session, like source of input content.

## v1.4.46 (2024-10-04)

### Direct commits

- Feat: Close #1018
- Feat: implement print session and context
- Feat: implement print session and context

## v1.4.45 (2024-10-04)

### Direct commits

- Feat: Setup for specific vendor, e.g. --setup-vendor=OpenAI

## v1.4.44 (2024-10-03)

### Direct commits

- Ci: use the latest tag by date

## v1.4.43 (2024-10-03)

### Direct commits

- Ci: use the latest tag by date

## v1.4.42 (2024-10-03)

### Direct commits

- Ci: use the latest tag by date
- Ci: use the latest tag by date

## v1.4.41 (2024-10-03)

### Direct commits

- Ci: trigger release workflow ony tag_created

## v1.4.40 (2024-10-03)

### Direct commits

- Ci: create repo dispatch

## v1.4.39 (2024-10-03)

### Direct commits

- Ci: test tag creation

## v1.4.38 (2024-10-03)

- Ci: test tag creation
- Ci: commit version changes only if it changed
- Ci: use TAG_PAT instead of secrets.GITHUB_TOKEN for tag push
- Updated predictions pattern

## v1.4.36 (2024-10-03)

### Direct commits

- Merge branch 'main' of github.com:danielmiessler/fabric
- Added redeeming thing.

## v1.4.35 (2024-10-02)

### Direct commits

- Feat: clean up html readability; add autm. tag creation

## v1.4.34 (2024-10-02)

### Direct commits

- Feat: clean up html readability; add autm. tag creation

## v1.4.33 (2024-10-02)

### Direct commits

- Feat: clean up html readability; add autm. tag creation
- Feat: clean up html readability; add autm. tag creation
- Feat: clean up html readability; add autm. tag creation

## v1.5.0 (2024-10-02)

### Direct commits

- Feat: clean up html readability; add autm. tag creation

## v1.4.32 (2024-10-02)

### PR [#1007](https://github.com/danielmiessler/Fabric/pull/1007) by [hallelujah-shih](https://github.com/hallelujah-shih): support turn any web page into clean view content

- Support turn any web page into clean view content

### PR [#1005](https://github.com/danielmiessler/Fabric/pull/1005) by [fn5](https://github.com/fn5): Update patterns/solve_with_cot/system.md typos

- Update patterns/solve_with_cot/system.md typos

### PR [#962](https://github.com/danielmiessler/Fabric/pull/962) by [alucarded](https://github.com/alucarded): Update prompt in agility_story

- Update system.md

### PR [#994](https://github.com/danielmiessler/Fabric/pull/994) by [OddDuck11](https://github.com/OddDuck11): Add pattern analyze_military_strategy

- Add pattern analyze_military_strategy

### PR [#1008](https://github.com/danielmiessler/Fabric/pull/1008) by [MattBash17](https://github.com/MattBash17): Update system.md in transcribe_minutes

- Update system.md in transcribe_minutes

## v1.4.31 (2024-10-01)

### PR [#987](https://github.com/danielmiessler/Fabric/pull/987) by [joshmedeski](https://github.com/joshmedeski): feat: remove cli list label and indentation

- Remove CLI list label and indentation for cleaner interface

### PR [#1011](https://github.com/danielmiessler/Fabric/pull/1011) by [fooman[org]](https://github.com/fooman): Grab transcript from youtube matching the user's language

- Grab transcript from YouTube matching the user's language instead of the first one

### Direct commits

- Add version updater bot functionality
- Add create_story_explanation pattern
- Support turning any web page into clean view content
- Update system.md in transcribe_minutes pattern
- Add epp pattern

## v1.4.30 (2024-09-29)

### Direct commits

- Feat: add version updater bot

## v1.4.29 (2024-09-29)

### PR [#996](https://github.com/danielmiessler/Fabric/pull/996) by [hallelujah-shih](https://github.com/hallelujah-shih): add wipe flag for ctx and session

- Add wipe flag for ctx and session

### PR [#967](https://github.com/danielmiessler/Fabric/pull/967) by [akashkankariya](https://github.com/akashkankariya): Updated Path to install to_pdf in readme[Bug Fix]

- Updated Path to install to_pdf [Bug Fix]

### PR [#984](https://github.com/danielmiessler/Fabric/pull/984) by [riccardo1980](https://github.com/riccardo1980): adding flag for pinning seed in openai and compatible APIs

- Adding flag for pinning seed in openai and compatible APIs

### PR [#991](https://github.com/danielmiessler/Fabric/pull/991) by [aculich](https://github.com/aculich): Fix GOROOT path for Apple Silicon Macs

- Fix GOROOT path for Apple Silicon Macs in setup instructions

### PR [#976](https://github.com/danielmiessler/Fabric/pull/976) by [pavdmyt](https://github.com/pavdmyt): fix: correct changeDefaultModel flag description

- Fix: correct changeDefaultModel flag description
