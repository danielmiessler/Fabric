# Changelog

## v1.4.304 (2025-09-30)

### Direct commits

- Create produce_report.sh
Added a script that compiles one  organized file out of all of the folder's text files. The script proceeds to use said file as  psychological data  input for Fabric's "analyze_personality" pattern, and outputs an analysis report.

## v1.4.303 (2025-09-30)

### Direct commits

- Feat: update `anthropic-sdk-go` to v1.13.0 and add new model

- Upgrade `anthropic-sdk-go` to version 1.13.0
- Add `ModelClaudeSonnet4_5` to supported models list
- Fix: improve PT-BR translation naturalness and fluency

- Thanks to @JuracyAmerico for Brazilian Portugese native speaker expertise!
- Replace "dos" with "entre" for better preposition usage

- Add definite articles where natural in Portuguese
- Clarify "configurações padrão" instead of just "padrões"

- Keep technical terms visible like "padrões/patterns"
- Remove unnecessary quotes around "URL"

- Make phrasing more natural "Exportar para arquivo"
- Feat(i18n): add i18n support for language variants (pt-BR/pt-PT)
• Add Brazilian Portuguese (pt-BR) translation file
• Add European Portuguese (pt-PT) translation file
• Implement BCP 47 locale normalization system
• Create fallback chain for language variants
• Add default variant mapping for Portuguese
• Update help text to show variant examples
• Add comprehensive test suite for variants
• Create documentation for i18n variant architecture
- Chore: remove garble installation from release workflow

- Remove garble installation step from release workflow
- Add comment for GoReleaser config file reference link

- The original idea of adding garble was to make it pass virus
  scanning during version upgrades for Winget, and this
  was a failed experiment.
- Chore: update CI workflow and simplify goreleaser build configuration

- Add changelog database to git tracking

- Remove unnecessary goreleaser comments
- Add version metadata to default build

- Rename windows build from garbled to standard
- Remove garble obfuscation from windows build

- Standardize ldflags across all build targets
- Inject version info during compilation
- Feat: migrate Azure client to openai-go/azure and default API version
CHANGES

- switch Azure OpenAI config to openai-go azure helpers
- require API key and base URL during configuration

- default API version to 2024-05-01-preview when unspecified
- trim and parse deployments input into clean slice

- update dependencies to support azure client and authentication flow
- add tests for configuration and default API version behavior

- remove latest-tag boundary logic from changelog walker (revert to the v1.4.213 version)
- simplify version assignment by matching commit messages directly
- Fix: One-time fix for CHANGELOG and changelog cache db
- Feat: add garble obfuscation for Windows builds and fix changelog generation

- Add garble tool installation to release workflow
- Configure garble obfuscation for Windows builds only

- Fix changelog walker to handle unreleased commits
- Update changelog database with latest changes

- Add mvdan to VSCode dictionary settings
- Implement boundary detection for released vs unreleased

- Keep newer commits as "Unreleased" until tagged
- Chore: prevent Go toolchain auto-download by setting GOTOOLCHAIN=local environment variable

- Add GOTOOLCHAIN=local to shell environment

- Configure preBuild hook in fabric package
- Wrap goimports with local toolchain setting

- Wrap gofmt with local toolchain setting
- Update treefmt module import structure

- Add pkgs parameter to treefmt config
- Create shell script wrappers for Go tools
- Chore: Go 1.25 upgrade and critical package updates for AI/ML services

- Upgrade Go from 1.24 to 1.25.1
- Update Anthropic SDK for web fetch tools

- Upgrade AWS Bedrock SDK 12 versions
- Update Azure Core and Identity SDKs

- Fix Nix config for Go version lag
- Update Docker base to golang:1.25-alpine

- Add comprehensive upgrade documentation
- Feat(i18n): add de/fr/ja/pt/zh/fa locales; expand tests, improve changelog spacing
CHANGES

- Add DE, FR, JA, PT, ZH, FA i18n locale files.
- Expand i18n tests with table-driven multilingual coverage.

- Verify 'html_readability_error' translations across all supported languages.
- Update README with release notes for added languages.

- Watch new locale and test files in VSCode.
- Insert blank lines between aggregated PR changelog sections.

- Append direct commits section only when content exists.
- Chore: update changelog formatting and sync changelog database

- Add line breaks to improve changelog readability

- Sync changelog database with latest entries
- Clean up whitespace in version sections

- Maintain consistent formatting across entries
- Chore: add spacing between changelog entries for improved readability

- Add blank lines between PR sections

- Update changelog database with  to correspond with CHANGELOG fix.
- Feat: update Vite and Rollup dependencies to latest versions

- Update Vite to version 5.4.20

- Update Rollup to version 4.50.1
- Add `@eslint-community/eslint-utils` version 4.9.0

- Update `@humanfs/node` to version 0.16.7
- Update `@humanwhocodes/retry` to version 0.4.3

- Update Rollup platform-specific packages to 4.50.1
- Add `@rollup/rollup-openharmony-arm64` version 4.50.1

- Closes Dependabot PR <https://github.com/danielmiessler/Fabric/pull/1763>
- Merge branch 'main' into kayvan/fix/0909-windows-flag-fix
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
- Merge branch 'main' into kayvan/fix/0909-windows-flag-fix
- Fix: Change attribution of PR to <https://github.com/OmriH-Elister>
- Chore: incoming 1762 changelog entry
- Chore: add `create_story_about_people_interaction` pattern for persona analysis

- Add `create_story_about_people_interaction` pattern description

- Include pattern in `ANALYSIS` and `WRITING` categories
- Update `suggest_pattern` system and user documentation

- Modify JSON files to incorporate new pattern details
- Merge branch 'danielmiessler:main' into stick
- Feat: add new pattern that creates story simulating interaction between two people
- Chore: incoming 1759 changelog entry
- Feat: add Windows-style forward slash flag support to CLI argument parser

- Add runtime OS detection for Windows platform
- Support `/flag` syntax for Windows command line

- Handle Windows colon delimiter `/flag:value` format
- Maintain backward compatibility with Unix-style flags

- Add comprehensive test coverage for flag extraction
- Support both `:` and `=` delimiters on Windows

- Preserve existing dash-based flag parsing logic
- Feat: add comprehensive internationalization support with English and Spanish locales

- Replace hardcoded strings with i18n.T translations
- Add en and es JSON locale files

- Implement custom translated help system
- Enable language detection from CLI args

- Add locale download capability
- Localize error messages throughout codebase

- Support TTS and notification translations
- Feat: add i18n support with Spanish localization and documentation improvements

- Add internationalization system with Spanish support
- Create contexts and sessions tutorial documentation

- Fix broken Warp sponsorship image URL
- Add locale detection from environment variables

- Update VSCode settings with new dictionary words
- Exclude VSCode settings from version workflows

- Update pattern descriptions and explanations
- Add comprehensive i18n test coverage
- Update Warp sponsor section with proper formatting

- Replace with correct div structure and styling
- Use proper Warp image URL from brand assets

- Add 'Special thanks to:' text and platform availability
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
- Docs: streamline install process with one-line installer scripts and update documentation

- Add markdown file triggers to GitHub workflow
- Update VSCode settings with new spell entries

- Simplify README installation with one-line installers
- Add bash installer script for Unix systems

- Add PowerShell installer script for Windows
- Create installer documentation with usage examples

- Remove redundant pattern from pattern explanations
- Docs: add Windows install via winget and Docker deployment instructions

- Add winget installation method for Windows
- Add Docker Hub and GHCR image references

- Include docker run examples for setup/patterns
- Remove deprecated PowerShell download link

- Delete unused show_fabric_options_markmap pattern
- Update suggest_pattern with new AI patterns

- Add personal development patterns for storytelling
- Chore: incoming 1741 changelog entry
- Fix: update release workflow to support manual dispatch with custom tag

- Support custom tag from client payload in workflow

- Fallback to github.ref_name when no custom tag provided
- Enable manual release triggers with specified tag parameter
- Ci: add changelog generation step to release workflow and support fork releases

- Add changelog generation step to GitHub release workflow
- Create updateReleaseForRepo helper method for release updates

- Add fork detection logic in UpdateReleaseDescription method
- Implement upstream repository release update for forks

- Add fallback to current repository when upstream fails
- Enhance error handling with detailed repository context

- Remove duplicate success logging from main method
- Ci: harden release pipeline; gate to upstream, migrate tokens, remove docker-on-tag
CHANGES

- Gate release and version workflows to upstream owner only.
- Switch tagging and releases to built-in GITHUB_TOKEN.

- Replace environment passing with step outputs across workflows.
- Remove docker-publish-on-tag workflow to reduce duplication and complexity.

- Add OCI description label to Docker image.
- Document GHCR multi-arch annotations for accurate package descriptions.

- Update README with new ARM binary release announcement.
- Simplify GoReleaser config by removing comments and extras.
- Chore: goreleaser and winget support
- Feat: add 'create_story_about_person' and 'heal_person' patterns; bump devalue
CHANGES

- Add create_story_about_person system pattern with narrative workflow
- Add heal_person system pattern for compassionate healing plans

- Update pattern_explanations to register new patterns and renumber indices
- Extend pattern_descriptions with entries, tags, and concise descriptions

- Add pattern_extracts for both patterns with full instruction blocks
- Bump devalue dependency from 5.1.1 to 5.3.2

- Refresh lockfile snapshots to reference updated devalue version
- Sync web static pattern_descriptions with new patterns
Updates `devalue` from 5.1.1 to 5.3.2

- [Release notes](<https://github.com/sveltejs/devalue/releases)>
- [Changelog](<https://github.com/sveltejs/devalue/blob/main/CHANGELOG.md)>

- [Commits](sveltejs/devalue@v5.1.1...v5.3.2)
updated-dependencies:
- dependency-name: devalue
  dependency-version: 5.3.2
  dependency-type: indirect
  dependency-group: npm_and_yarn
Signed-off-by: dependabot[bot] <support@github.com>
- Fix: update Docker workflow to use specific Dockerfile and monitor markdown file changes
• Add explicit Dockerfile path to Docker build action
• Remove markdown files from workflow paths-ignore filter
• Enable CI triggers for documentation file changes
• Specify Docker build context with custom file location
- Ci: add tag-based multi-arch Docker publish to GHCR and Docker Hub
CHANGES

- Add GitHub Actions workflow to publish Docker images on tags
- Build multi-arch images with Buildx and QEMU across amd64, arm64

- Tag images using semver; push to GHCR and Docker Hub
- Set :latest only for highest semver tag via imagetools

- Gate patterns workflow steps on detected changes instead of failing
- Auto-detect GitHub owner and repo from git remote URL

- Remove hardcoded repository values in changelog release manager
- Normalize image names to lowercase for registry compatibility

- Enable GitHub Actions cache for faster Docker builds
- Add VS Code dictionary entries for Docker-related terms
- Chore: upgrade ollama dependency from v0.9.0 to v0.11.7
• Update ollama package to version 0.11.7
• Refresh go.sum with new dependency checksums

- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2025-0317](https://nvd.nist.gov/vuln/detail/CVE-2025-0317)>
- **CVSS Score**: 7.5 (High)

- **Description**: A vulnerability in ollama/ollama versions <=0.3.14 allows a malicious user to upload and create a customized GGUF model file on the Ollama server. This can lead to a division by zero error in the ggufPadding function, causing the server to crash and resulting in a Denial of Service (DoS) attack.
- **Affected**: Ollama server versions ≤ 0.3.14

- **Impact**: Denial of Service through division by zero error
- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2025-0315](https://nvd.nist.gov/vuln/detail/CVE-2025-0315)>

- **CVSS Score**: 7.5 (High)
- **Description**: Vulnerability allows Denial of Service via customized GGUF model file upload on Ollama server.

- **Affected**: Ollama/ollama versions ≤ 0.3.14
- **Impact**: Denial of Service through malicious GGUF model file uploads

- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2024-12886](https://nvd.nist.gov/vuln/detail/CVE-2024-12886)>
- **CVSS Score**: 7.5 (High)

- **Description**: An Out-Of-Memory (OOM) vulnerability exists in the ollama server version 0.3.14. This vulnerability can be triggered when a malicious API server responds with a gzip bomb HTTP response, leading to the ollama server crashing.
- **Affected**: Ollama server version 0.3.14

- **Impact**: Denial of Service through memory exhaustion via gzip bomb attack
- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2024-8063](https://nvd.nist.gov/vuln/detail/CVE-2024-8063)>

- **CVSS Score**: 7.5 (High)
- **Description**: Security vulnerability with high severity rating

- **Impact**: Requires patching for security compliance
- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2024-12055](https://nvd.nist.gov/vuln/detail/CVE-2024-12055)>

- **CVSS Score**: 7.5 (High)
- **Description**: High-severity security vulnerability requiring immediate attention

- **Impact**: Critical security flaw needing remediation
- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2025-51471](https://nvd.nist.gov/vuln/detail/CVE-2025-51471)>

- **CVSS Score**: 6.9 (Medium)
- **Description**: Medium severity security vulnerability

- **Impact**: Security risk requiring patching as part of comprehensive security updates
- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2025-46394](https://nvd.nist.gov/vuln/detail/CVE-2025-46394)>

- **CVSS Score**: 3.2 (Low)
- **Description**: Low-severity security issue

- **Impact**: Minor security concern addressed as part of comprehensive security maintenance
- **Link**: [<https://nvd.nist.gov/vuln/detail/CVE-2024-58251](https://nvd.nist.gov/vuln/detail/CVE-2024-58251)>

- **CVSS Score**: 2.5 (Low)
- **Description**: Low-severity security vulnerability

- **Impact**: Minimal security risk addressed for comprehensive security posture
This comprehensive security fix addresses **8 CVEs*- total:
- **5 High Severity*- vulnerabilities (CVSS 7.5)

- **1 Medium Severity*- vulnerability (CVSS 6.9)
- **2 Low Severity*- vulnerabilities (CVSS 3.2 and 2.5)
The majority of high-severity issues are related to **Ollama server vulnerabilities*- that could lead to Denial of Service attacks through various vectors including division by zero errors, memory exhaustion, and malicious file uploads. These fixes ensure robust protection against these attack vectors and maintain system availability.
**Priority**: The high-severity Ollama vulnerabilities should be considered critical for any systems running Ollama server components, as they can lead to service disruption and potential system crashes.
- Chore: remove docker-test framework and simplify production docker setup

- Remove entire docker-test directory and testing infrastructure
- Delete complex test runner script and environment files

- Simplify production Dockerfile with multi-stage build optimization
- Remove docker-compose.yml and start-docker.sh helper scripts

- Update README with cleaner Docker usage instructions
- Streamline container build process and reduce image size
- Docs: add contributing, security, support, and code-of-conduct docs; add docs index
CHANGES

- Add CODE_OF_CONDUCT defining respectful, collaborative community behavior
- Add CONTRIBUTING with setup, testing, PR, changelog requirements

- Add SECURITY policy with reporting process and response timelines
- Add SUPPORT guide for bugs, features, discussions, expectations

- Add docs README indexing guides, quick starts, contributor essentials
- Refactor: replace stderr prints with centralized debuglog.Log and improve auth messaging

- Replace fmt.Fprintf/os.Stderr with centralized debuglog.Log across CLI
- Add unconditional Log function to debuglog for important messages

- Improve OAuth flow messaging and token refresh diagnostics
- Update tests to capture debuglog output via SetOutput

- Convert Perplexity streaming errors to unified debug logging
- Emit file write notifications through debuglog instead of stderr

- Warn on ambiguous model selection using centralized logger
- Announce large audio processing steps via debuglog progress messages

- Standardize extension registry and patterns warnings through debuglog
- Refactor: route Anthropic beta failure logs through internal debug logger
CHANGES

- Replace fmt.Fprintf stderr with debuglog.Debug for beta failures
- Import internal log package and remove os dependency

- Standardize logging level to debuglog.Basic for beta errors
- Preserve fallback stream behavior when beta features fail

- Maintain message send fallback when beta options fail
- Docs: update README with Venice AI provider and Windows install script

- Add Venice AI provider configuration with API endpoint
- Document Venice AI as privacy-first open-source provider

- Include PowerShell installation script for Windows users
- Add debug levels section to table of contents

- Update recent major features with v1.4.294 release notes
- Configure Venice AI base URL and response settings
- Feat: add --debug flag with levels and centralized logging
CHANGES

- Add --debug flag controlling runtime logging verbosity levels
- Introduce internal/log package with Off, Basic, Detailed, Trace

- Replace ad-hoc Debugf and globals with centralized debug logger
- Wire debug level during early CLI argument parsing

- Add bash, zsh, fish completions for --debug levels
- Document debug levels in README with usage examples

- Add comprehensive STT guide covering models, flags, workflows
- Simplify splitAudioFile signature and log ffmpeg chunking operations

- Remove FABRIC_STT_DEBUG environment variable and related code
- Clean minor code paths in vendors and template modules
- Feat: highlight default vendor/model in listings, pass registry defaults
CHANGES

- Update PrintWithVendor signature to accept default vendor and model
- Mark default vendor/model with asterisk in non-shell output

- Compare vendor and model case-insensitively when marking
- Pass registry defaults to PrintWithVendor from CLI

- Add test ensuring default selection appears with asterisk
- Keep shell completion output unchanged without default markers
- Docs: update version number in README updates section from v1.4.290 to v1.4.291
- Feat: add speech-to-text via OpenAI with transcription flags and completions
CHANGES

- Add --transcribe-file flag to transcribe audio or video
- Add --transcribe-model flag with model listing and completion

- Add --split-media-file flag to chunk files over 25MB
- Implement OpenAI transcription using Whisper and GPT-4o Transcribe

- Integrate transcription pipeline into CLI before readability processing
- Provide zsh, bash, fish completions for new transcription flags

- Validate media extensions and enforce 25MB upload limits
- Update README with release and corrected pattern link path
- Feat: add per-pattern model mapping support via environment variables
• Add per-pattern model mapping documentation section
• Implement environment variable lookup for pattern-specific models
• Support vendor|model format in environment variable specification
• Check pattern-specific model when no model explicitly set
• Transform pattern names to uppercase environment variable format
• Add table of contents entry for new feature
• Enable shell startup file configuration for patterns
- Feat: add --no-variable-replacement flag to disable pattern variable substitution

- Introduce CLI flag to skip pattern variable replacement.
- Wire flag into domain request and session builder.

- Avoid applying input variables when replacement is disabled.
- Provide PatternsEntity.GetWithoutVariables for input-only pattern processing support.

- Refactor patterns code into reusable load and apply helpers.
- Update bash, zsh, fish completions with new flag.

- Document flag in README and CLI help output.
- Add unit tests covering GetWithoutVariables path and behavior.

- Ensure {{input}} placeholder appends when missing in patterns.
- Fix: improve YouTube subtitle language fallback handling in yt-dlp integration

- Fix typo "Gemmini" to "Gemini" in README
- Add "kballard" and "shellquote" to VSCode dictionary

- Add "YTDLP" to VSCode spell checker
- Enhance subtitle language options with fallback variants

- Build language options string with comma-separated alternatives
- Feat: add release updates section and Gemini thinking support

- Add comprehensive "Recent Major Features" section to README
- Introduce new readme_updates Python script for automation

- Enable Gemini thinking configuration with token budgets
- Update CLI help text for Gemini thinking support

- Add comprehensive test coverage for Gemini thinking
- Create documentation for README update automation

- Reorganize README navigation structure with changelog section
- Refactor: extract token budget constants for thinking levels with validation bounds

- Extract hardcoded token values into named constants

- Add comprehensive documentation for token budget purposes
- Implement token validation bounds (1-10000) in parsing

- Replace magic numbers with semantic constant references
- Improve code maintainability through constant extraction
- Feat: add cross-provider --thinking flag mapping to Anthropic/OpenAI
CHANGES

- Add --thinking flag to set reasoning level cross-vendors
- Map Anthropic thinking levels and token budgets appropriately

- Translate OpenAI reasoning effort from thinking levels
- Propagate Thinking through ChatOptions, server, and dry-run output

- Update zsh, bash, fish completions with thinking choices
- Expand suggest_pattern docs with categories, workflows, usage examples

- Remove outdated suggest_pattern user files to avoid duplication
- Add VSCode dictionary terms: Anki, DMARC, wireframes

- Extend tests to include Thinking defaults in ChatOptions
- Chore: upgrade anthropic-sdk-go to v1.9.1 and add beta feature support for context-1m

- Upgrade anthropic-sdk-go from v1.7.0 to v1.9.1

- Upgrade golang.org/x/crypto from v0.39.0 to v0.40.0
- Add modelBetas map for beta feature configuration

- Implement context-1m-2025-08-07 beta for Claude Sonnet 4
- Add beta header support in streaming requests

- Add beta header support in standard requests
- Implement fallback mechanism when beta features fail

- Preserve existing beta headers in OAuth transport
- Add test coverage for model beta configuration
- Chore: incoming 1695 changelog entry
- Refactor: standardize obtain_completion_files logging; use stderr-only printf
CHANGES

- Replace print_info with tagged printf directed to stderr.
- Replace print_dry_run with tagged printf directed to stderr.

- Add comment enforcing stderr-only output inside this function.
- Preserve dry-run behavior by echoing path only on stdout.

- Retain error handling using print_error for directory creation.
- Normalize log message prefixes to [INFO] and [DRY-RUN].

- Avoid stdout pollution by routing informational messages to stderr.
- Fix: convert GitHub blob/tree URLs to raw and validate completion downloads
CHANGES

- Add helper to translate GitHub blob/tree to raw URLs
- Use effective URL in curl and wget download paths

- Validate downloaded files are non-empty and not HTML
- Redirect info and dry-run messages to standard error

- Relocate temporary directory cleanup trap into main execution
- Improve error messages when completion download sources appear invalid
- Docs: add quick install method for shell completions without cloning repo

- Add one-liner curl install for completions

- Support downloading completions when files missing locally
- Add dry-run option for preview changes

- Enable custom download source via environment variable
- Create temp directory for downloaded completion files

- Add automatic cleanup of temporary files
- Update documentation with new installation methods
- Chore: incoming 1692 changelog entry
- Feat: add -V/--vendor flag and vendor-aware model selection
CHANGES

- Add -V/--vendor flag to specify model vendor
- Implement vendor-aware model resolution and availability validation

- Warn on ambiguous models; suggest --vendor to disambiguate
- Update bash, zsh, fish completions with vendor suggestions

- Extend --listmodels to print vendor|model when interactive
- Add VendorsModels.PrintWithVendor; sort vendors and models alphabetically

- Pass vendor through API; update server chat handler
- Standardize docs and errors to --yt-dlp-args="..." syntax

- Add test covering ambiguous model warning across multiple vendors
- Promote go-shellquote to direct dependency in go.mod
- Feat: enhance completions with 'fabric-ai' alias, dynamic exec, installer
CHANGES

- Support 'fabric-ai' alias across Zsh, Bash, and Fish
- Use invoked command for dynamic completion list queries

- Refactor Fish completions into reusable registrar for multiple commands
- Update Bash completion to reference executable via COMP_WORDS[0]

- Extend Zsh compdef to register fabric and fabric-ai
- Add cross-shell installer script with autodetection and dry-run mode

- Document installation, features, troubleshooting in new completions guide
- Chore: incoming 1687 changelog entry
- Feat(gemini): enable web search, citations, and search-location validation
CHANGES

- Enable Gemini models to use web search tool
- Validate search-location timezone or language code formats

- Normalize language codes from underscores to hyphenated form
- Inject Google Search tool when --search flag enabled

- Append deduplicated web citations under standardized Sources section
- Improve robustness for nil candidates and content parts

- Factor generation config builder for reuse in streaming
- Update CLI help and completions to include Gemini
- Chore: incoming 1686 changelog entry
- Fix: prevent duplicate text output in OpenAI streaming responses

- Skip processing of ResponseOutputTextDone events

- Prevent doubled text in stream output
- Add clarifying comment about API behavior

- Maintain delta chunk streaming functionality
- Fix duplicate content issue in responses
- Chore: incoming 1685 changelog entry
- Fix(gemini): map chat roles to Gemini user/model in convertMessages
CHANGES

- Map assistant role to model per Gemini constraints
- Map system, developer, function, tool roles to user

- Default unrecognized roles to user to preserve instruction context
- Add unit test validating convertMessages role mapping logic

- Import chat package in tests for role constants
- Docs: update release notes
- Chore: incoming 1681 changelog entry
- Refactor: replace custom arg parser with shellquote; precompile regexes
CHANGES

- Precompile regexes for video, playlist, VTT tags, durations.
- Parse yt-dlp additional arguments using shellquote.Split for safety.

- Validate user-provided yt-dlp args and surface quoting errors.
- Reuse compiled regex in GetVideoOrPlaylistId extractions for stability.

- Simplify removeVTTTags by leveraging precompiled VTT tag matcher.
- Parse ISO-8601 durations with precompiled pattern for efficiency.

- Replace inline VTT language regex with cached compiled matcher.
- Remove unused findVTTFiles helper and redundant language checks.

- Add go-shellquote dependency in go.mod and go.sum.
- Reduce allocations by eliminating per-call regexp.MustCompile invocations.
- Docs: update release notes
- Chore: incoming 1681 changelog entry
- Feat: add smart subtitle language fallback when requested locale unavailable
CHANGES

- Introduce findVTTFilesWithFallback to handle subtitle language absence
- Prefer requested language VTT, gracefully fallback to available alternatives

- Auto-detect downloaded subtitle language and proceed without interruption
- Update yt-dlp processing to use fallback-aware VTT discovery

- Document language fallback behavior and provide usage example
- Return first available VTT when no specific language requested

- Detect language-coded filenames using regex for robust matching
- Docs: update YouTube processing documentation for yt-dlp argument precedence control

- Add user argument precedence over built-in flags

- Document argument order and override behavior
- Include new precedence section with detailed explanation

- Add override examples for language and format
- Update tips section with precedence guidance

- Modify Go code to append user args last
- Add testing tip for subtitle language discovery

- Include practical override use case examples
- Feat: add `--yt-dlp-args` flag for custom YouTube downloader options

- Introduce `--yt-dlp-args` flag for advanced control

- Allow passing browser cookies for authentication
- Improve error handling for YouTube rate limits

- Add comprehensive documentation for YouTube processing
- Refactor YouTube methods to accept additional arguments

- Update shell completions to include new flag
- Merge branch 'main' into 0803-youtube-transcript-lang-fix
- Chore: format fix
- Chore: incoming 1679 changelog entry
- Feat(cli): add cross-platform desktop notifications with secure custom commands
CHANGES

- Integrate notification sending into chat processing workflow
- Add --notification and --notification-command CLI flags and help

- Provide cross-platform providers: macOS, Linux, Windows with fallbacks
- Escape shell metacharacters to prevent injection vulnerabilities

- Truncate Unicode output safely for notification message previews
- Update bash, zsh, fish completions with new notification options

- Add docs and YAML examples for configuration and customization
- Add unit tests for providers and notification integration paths
- Ci: add write permissions to update_release_notes job

- Add contents write permission to release notes job

- Enable GitHub Actions to modify repository contents
- Fix potential permission issues during release process
- Chore: incoming 1676 changelog entry
- Feat: add 'gpt-5' to raw-mode models in OpenAI client

- Add gpt-5 to raw mode model requirements list.
- Ensure gpt-5 responses bypass structured chat message formatting.

- Align NeedsRawMode logic with expanded OpenAI model support.
- Docs: document GetTokenFromEnv behavior and token environment fallback
- Docs: document GetTokenFromEnv behavior and token environment fallback
- Chore: incoming 1676 changelog entry
- Refactor: centralize GitHub token retrieval logic into utility function

- Extract token retrieval into `util.GetTokenFromEnv` function

- Support both `GITHUB_TOKEN` and `GH_TOKEN` environment variables
- Replace direct `os.Getenv` calls with utility function

- Add new `util/token.go` file for token handling
- Update walker.go to use centralized token logic

- Update main.go to use token utility function
- Chore: incoming 1673 changelog entry
- Fix: ensure Anthropic client always sets temperature to override API default

- Always set temperature parameter for consistent behavior

- Prioritize TopP over temperature when explicitly set
- Override Anthropic's default 1.0 with Fabric's 0.7

- Add comprehensive tests for parameter precedence logic
- Update VSCode dictionary with Keploy entry

- Simplify conditional logic for temperature/TopP selection
- Chore: incoming 1673 changelog entry
- Refactor: improve chat parameter defaults handling with domain constants

- Add domain constants for default chat parameter values

- Update Anthropic client to check explicitly set parameters
- Add documentation linking CLI flags to domain defaults

- Improve temperature and TopP parameter selection logic
- Ensure consistent default values across CLI and domain

- Replace zero-value checks with explicit default comparisons
- Centralize chat option defaults in domain package
- Chore: incoming 1673 changelog entry
- Ci: refactor release workflow to use shared version job and simplify OS handling
- Chore: incoming 1673 changelog entry
- Fix: update anthropic SDK and refactor release workflow for release notes generation

- Upgrade anthropic-sdk-go from v1.4.0 to v1.7.0

- Move changelog generation to separate workflow job
- Add Claude Opus 4.1 model support

- Fix temperature/topP parameter conflict for models
- Separate release artifact upload from changelog update

- Add dedicated update_release_notes job configuration
- Chore: remove redundant words
Signed-off-by: queryfast <queryfast@outlook.com>
- Ci: separate release notes generation into dedicated job

- Move changelog generation to separate workflow job

- Add fallback logic for YouTube subtitle language detection
- Remove changelog commands from main release job

- Create dedicated update_release_notes job with Go setup
- Implement retry mechanism without language specification

- Improve yt-dlp command argument construction flexibility
- Add proper checkout and Go configuration steps
- Fix typos in t_ patterns
- Chore: incoming 1658 changelog entry
- Chore: Update changelog cache db
- Feat: add database sync before generating changelog in release workflow

- Add database sync command to release workflow
- Ensure changelog generation includes latest database updates
- Chore: incoming 1657 changelog entry
- Feat: add GitHub release description update with AI summary

- Add `--release` flag to command line options documentation

- Enable AI summary updates for GitHub releases
- Support version-specific release description updates

- Reorder internal package imports for consistency
- Feat: add GitHub release description update via `--release` flag

- Add `--release` flag to generate_changelog to update GitHub release

- Implement `ReleaseManager` for managing release descriptions
- Create `release.go` for handling release updates

- Update `release.yml` to run changelog generation
- Ensure mutual exclusivity for `--release` with other flags

- Modify `Config` struct to include `Release` field
- Update `main.go` to handle new release functionality
- Chore: incoming 1654 changelog entry
- Fix: prevent file overwrite and improve output messaging in CreateOutputFile

- Add file existence check before creating output file

- Return error if target file already exists
- Change success message to write to stderr

- Update message format with brackets for clarity
- Prevent accidental file overwrites during output creation
- Chore: incoming 1653 changelog entry
- Docs: update Gemini TTS model references to gemini-2.5-flash-preview-tts

- Update documentation examples to use gemini-2.5-flash-preview-tts

- Replace gemini-2.0-flash-tts references throughout Gemini-TTS.md
- Update voice selection example commands

- Modify CLI help text example command
- Update changelog database binary file
- Chore: differentiate voice descriptions
- Chore: incoming 1652 changelog entry
- Feat: add Gemini TTS voice selection and listing functionality

- Add `--voice` flag for TTS voice selection

- Add `--list-gemini-voices` command for voice discovery
- Implement voice validation for Gemini TTS models

- Update shell completions for voice options
- Add comprehensive Gemini TTS documentation

- Create voice samples directory structure
- Extend spell checker dictionary with voice names
- Fix: correct audio data extraction to avoid double byte conversion

- Remove redundant byte conversion from audio data extraction

- Extract audio data as string before converting once
- Simplify audio data processing in chat handler

- Fix potential data corruption in audio output
- Fix: initialize Parts slice in genai.Content struct to prevent nil pointer errors

- Initialize Parts slice with empty slice in Content struct

- Prevent potential nil pointer dereference during message conversion
- Ensure Parts field is ready for append operations

- Improve robustness of convertMessages function in Gemini client
- Chore: minor format fix
- Chore: more spelling words
- Refactor: extract TTS methods and add audio validation with security limits

- Extract text extraction logic into separate method

- Add GenAI client creation helper function
- Split TTS generation into focused helper methods

- Add audio data size validation with security limits
- Implement MIME type validation for audio responses

- Add WAV file generation input validation checks
- Pre-allocate buffer capacity for better performance

- Define audio constants for reusable configuration
- Add comprehensive error handling for edge cases

- Validate generated WAV data before returning results
- Chore: update changelog generation to sync database

- Add database sync command to changelog workflow

- Remove unnecessary newline addition in changelog processing
- Chore: incoming 1650 changelog entry
- Chore: update Gemini SDK to new genai library and add TTS audio output support

- Replace deprecated generative-ai-go with google.golang.org/genai library

- Add TTS model detection and audio output validation
- Implement WAV file generation for TTS audio responses

- Add audio format checking utilities in CLI output
- Update Gemini client to support streaming with new SDK

- Add "Kore" and "subchunk" to VSCode spell checker dictionary
- Remove extra blank line from changelog formatting

- Update dependency imports and remove unused packages
- Docs: minor formatting of CHANGELOG
- Chore: incoming 1649 changelog entry
- Feat: prevent unconfigured API initialization and add Docker test suite

- Add BEDROCK_AWS_REGION requirement for Bedrock initialization

- Implement IsConfigured check for Ollama API URL
- Create comprehensive Docker testing environment with 6 scenarios

- Add interactive test runner with shell access
- Include environment files for different API configurations

- Update spell checker dictionary with new test terms
- Document testing workflow and expected results
- Chore: incoming 1647 changelog entry
- Chore: replace git tag lookup with version.nix file reading for release workflow

- Remove OS-specific git tag retrieval steps

- Add unified version extraction from nix file
- Include version format validation with regex check

- Add error handling for missing version file
- Consolidate cross-platform version logic into single step

- Use bash shell for consistent version parsing
- Chore: incoming 1642 changelog entry
- Fix: improve error message formatting in version date parsing

- Add actual error details to date parsing failure message

- Include error variable in stderr output formatting
- Enhance debugging information for invalid date formats
- Refactor: simplify merge pattern management by removing unnecessary struct wrapper

- Remove mergePatternManager struct wrapper for patterns

- Replace struct fields with package-level variables
- Simplify getMergePatterns function implementation

- Clean up merge commit detection documentation
- Reduce code complexity in pattern initialization

- Maintain thread-safe lazy initialization with sync.Once
- Chore: standardize logging output format and improve error messages in changelog generator

- Replace emoji prefixes with bracketed text labels

- Standardize synchronization step logging format across methods
- Simplify version existence check error message text

- Update commit author email extraction comment clarity
- Maintain consistent stderr output formatting throughout sync process
- Chore: incoming 1642 changelog entry
- Refactor: improve error handling and simplify merge pattern management in changelog generation

- Remove unused runtime import from processing.go

- Simplify date parsing error messages in cache
- Replace global merge pattern variables with struct

- Use sync.Once for thread-safe pattern initialization
- Remove OS-specific file deletion instructions from errors

- Clean up merge commit detection function documentation
- Eliminate redundant error variable in date parsing
- Chore: incoming 1642 changelog entry
- Fix: improve error reporting in date parsing and merge commit detection

- Capture first RFC3339Nano parsing error for better diagnostics

- Display both RFC3339Nano and RFC3339 errors in output
- Extract merge patterns to variable for cleaner code

- Improve error message clarity in date parsing failures
- Feat: add cross-platform file removal instructions for changelog generation

- Import runtime package for OS detection

- Add Windows-specific file deletion commands in error messages
- Provide both Command Prompt and PowerShell alternatives

- Maintain existing Unix/Linux rm command for non-Windows systems
- Improve user experience across different operating systems
- Chore: incoming 1642 changelog entry
- Refactor: replace sync.Once with mutex for merge patterns initialization

- Replace sync.Once with mutex and boolean flag

- Add thread-safe initialization check for merge patterns
- Remove overly broad merge pattern regex rule

- Improve error messaging for file removal failures
- Clarify filesystem vs git index error contexts

- Add detailed manual intervention instructions for failures
- Chore: incoming 1642 changelog entry
- Perf: optimize merge pattern matching with lazy initialization and sync.Once

- Add sync package import for concurrency safety

- Implement lazy initialization for merge patterns using sync.Once
- Wrap merge patterns in getMergePatterns function

- Replace direct mergePatterns access with function call
- Ensure thread-safe pattern compilation on first use
- Chore: incoming 1642 changelog entry
- Refactor: improve merge commit detection and update error messages

- Move merge patterns to package-level variables

- Update date parsing error message for clarity
- Simplify author email field comment

- Extract regex compilation from function scope
- Improve merge commit detection performance

- Clarify RFC3339 fallback error context
- Fix: improve warning message clarity for invalid commit timestamps

- Simplify warning message for invalid commit timestamps

- Remove parenthetical explanation about git history rewrites
- Make error message more concise and readable
- Chore: optimize error logging and regex pattern matching for better performance

- Remove redundant RFC3339Nano parsing error message

- Enhance RFC3339 error message with version name
- Pre-compile regex patterns for merge commit detection

- Replace regexp.MatchString with compiled pattern matching
- Improve merge commit pattern matching performance

- Add structured regex compilation for better efficiency
- Fix: improve error handling and guidance for file removal failures

- Replace generic warning with detailed error message

- Add step-by-step manual intervention instructions
- Provide multiple recovery options for users

- Separate git and filesystem error reporting
- Include specific commands for manual cleanup
- Docs: improve code comments for version pattern and PR commit fields

- Expand version pattern regex documentation with examples

- Add matching and non-matching commit message examples
- Clarify version pattern behavior with optional prefix

- Update PR commit field comments for clarity
- Document email field availability from GitHub API

- Simplify timestamp and parents field descriptions
- Feat: improve changelog entry creation and error messages

- Rename `changelogDate` to `versionDate` for clarity

- Enhance error message for git index removal failure
- Add comments for `versionPattern` regex in `walker.go`
- Chore: incoming 1642 changelog entry
- Chore: improve error message clarity in changelog generation and cache operations

- Clarify RFC3339Nano date parsing error message

- Improve PR batch cache save error description
- Add context for commit timestamp fallback warning

- Specify git index in file removal error message
- Chore: incoming 1642 changelog entry
- Chore: improve error logging and documentation in changelog generation components

- Add date string to RFC3339 parsing error messages

- Enhance isMergeCommit function documentation with detailed explanation
- Document calculateVersionDate function with comprehensive behavior description

- Improve error context for date parsing failures
- Add implementation details for merge commit detection methods

- Clarify fallback behavior in version date calculation
- Chore: incoming 1642 changelog entry
- Chore: improve error message clarity in version existence check for git history sync

- Enhance warning message with additional context details

- Add guidance for users when version check fails
- Improve error handling feedback in sync operation

- Provide actionable steps for troubleshooting sync issues
- Chore: incoming 1642 changelog entry
- Feat: add email field support and improve error logging in changelog generation

- Add Email field to PRCommit struct for author information

- Extract version date calculation into reusable helper function
- Redirect error messages from stdout to stderr properly

- Populate commit email from GitHub API responses correctly
- Add comprehensive test coverage for email field handling

- Remove duplicate version date calculation code blocks
- Import os package for proper stderr output handling
- Feat: improve timestamp handling and merge commit detection in changelog generator

- Add debug logging for date parsing failures

- Pass forcePRSync parameter explicitly to fetchPRs method
- Implement comprehensive merge commit detection using parents

- Capture actual commit timestamps from GitHub API
- Calculate version dates from most recent commit

- Add parent commit SHAs for merge detection
- Use real commit dates instead of current time

- Add timestamp validation with fallback handling
- Chore: incoming 1642 changelog entry
- Feat: add database synchronization and improve changelog processing workflow

- Add database sync command with comprehensive validation

- Implement version and commit existence checking methods
- Enhance time parsing with RFC3339Nano fallback support

- Cache fetched PRs during changelog entry creation
- Remove individual incoming files using git operations

- Add sync-db flag for database integrity validation
- Improve commit-PR mapping verification process

- Exclude incoming directory from workflow trigger paths
- Docs: clean up duplicate CHANGELOG for v1.4.262
- Docs: Update CHANGELOG after v1.4.263
- Chore: incoming 1641 changelog entry
- Chore: extend proxy timeout in `vite.config.ts` to 15 minutes

- Increase `/api` proxy timeout to 900,000 ms

- Increase `/names` proxy timeout to 900,000 ms
- Docs: Remove duplicated section.
- Chore: fix tests for generate_changelog
- Chore: adjust `insertVersionAtTop` for consistent newline handling
- Chore: adjust newline handling in `insertVersionAtTop` method
- Chore: trim leading newline in changelog entry content
- Chore: simplify direct commits content handling in changelog generation
- Refactor: rename ProcessIncomingPRs to CreateNewChangelogEntry for clarity

- Rename ProcessIncomingPRs method to CreateNewChangelogEntry

- Update method comment to reflect new name
- Update main.go to call renamed method

- Reduce newline spacing in content formatting
- Chore: incoming 1640 changelog entry
- Fix: formatting fixes in tests.
- Feat: enhance changelog generator to accept version parameter for PR processing

- Pass version parameter to changelog generation workflow

- Update ProcessIncomingPRs method to accept version string
- Add commit SHA tracking to prevent duplicate entries

- Modify process-prs flag to require version parameter
- Improve changelog formatting with proper spacing

- Update configuration to use ProcessPRsVersion string field
- Enhance direct commit filtering with SHA exclusion

- Update documentation to reflect version parameter requirement
- Feat: enhance changelog generation to avoid duplicate commit entries

- Extract PR numbers from processed changelog files

- Pass processed PRs map to direct commits function
- Filter out commits already included via PR files

- Reduce extra newlines in changelog version insertion
- Add strconv import for PR number parsing

- Prevent duplicate entries between PRs and direct commits
- Improve changelog formatting consistency
- Fix: ensure the PR#.txt file ends with a newline.
- Feat: change push behavior from opt-out to opt-in with GitHub token auth

- Change `NoPush` config field to `Push` boolean

- Update CLI flag from `--no-push` to `--push`
- Add GitHub token authentication for push operations

- Import `os` and HTTP transport packages
- Reverse push logic to require explicit enable

- Update documentation for new push behavior
- Add automatic GitHub repository detection for auth
- Chore: update gitignore and simplify changelog generator error handling

- Add .claude/ directory to gitignore exclusions

- Update comment clarity for SilenceUsage flag
- Remove redundant error handling in main function

- Simplify command execution without explicit error checking
- Fix CLI error handling and improve git status validation

- Add SilenceUsage to prevent help output on errors
- Add GetStatusDetails method to show which files are dirty

- Include direct commits in ProcessIncomingPRs for complete AI summaries
- Chore: add automated changelog processing for CI/CD integration

- Add incoming PR preprocessing with validation

- Implement release aggregation for incoming files
- Create git operations for staging changes

- Add comprehensive test coverage for processing
- Extend GitHub client with validation methods

- Support version detection from nix files
- Include documentation for automated workflow

- Add command flags for PR processing
- Docs: Update CHANGELOG after v1.4.261
- Chore: update `NeedsRawMode` to include `mistral` prefix

- Add `mistral` to `ollamaPrefixes` list.
- Docs: Update CHANGELOG after v1.4.260
- Chore: add API key setup question to Exolab AI plugin configuration

- Add "openaiapi" to VSCode spell check dictionary

- Include API key setup question in Exolab client
- Configure API key as required field for setup

- Maintain existing API base URL configuration order
- Docs: Update CHANGELOG after v1.4.259
- Feat: improve timestamp parsing to handle fractional seconds in YouTube tool

- Move timestamp regex initialization to init function

- Add parseSeconds helper function for fractional seconds
- Replace direct strconv.Atoi calls with parseSeconds function

- Support decimal seconds in timestamp format parsing
- Extract seconds parsing logic into reusable function
- Fix: Youtube VTT parsing gap test
- Feat: enhance VTT duplicate filtering to allow legitimate repeated content

- Fix regex escape sequence for timestamp parsing

- Add configurable time gap constant for repeat detection
- Track content with timestamps instead of simple deduplication

- Implement time-based repeat inclusion logic for choruses
- Add timestamp parsing helper functions for calculations

- Allow repeated content after significant time gaps
- Preserve legitimate recurring phrases while filtering duplicates
- Chore: refactor timestamp regex and seenSegments logic

- Update `timestampRegex` to support optional seconds/milliseconds

- Change `seenSegments` to use `struct{}` for memory efficiency
- Refactor duplicate check using `struct{}` pattern

- Improve readability by restructuring timestamp logic
- Chore: refactor timestamp regex to global scope and add spell check words

- Move timestamp regex to global package scope

- Remove duplicate regex compilation from isTimeStamp function
- Add "horts", "mbed", "WEBVTT", "youtu" to spell checker

- Improve regex performance by avoiding repeated compilation
- Clean up code organization in YouTube module
- Fix: prevent duplicate segments in VTT file processing

- Add deduplication map to track seen segments
- Skip duplicate text segments in plain VTT processing

- Skip duplicate segments in timestamped VTT processing
- Improve timestamp regex to handle more formats

- Use clean text as deduplication key consistently
- Docs: Update CHANGELOG after v1.4.258
- Chore: define constants for file and directory permissions
- Chore: improve error handling in `ensureEnvFile` function
- Refactor: improve error handling and permissions in `ensureEnvFile`
- Feat: add startup check to initialize config and .env file

- Introduce ensureEnvFile function to create ~/.config/fabric/.env if missing.
- Add directory creation for config path in ensureEnvFile.

- Integrate setup flag in CLI to call ensureEnvFile on demand.
- Handle errors for home directory detection and file operations.
- Docs: Update README and CHANGELOG after v1.4.257
- Feat: add disable-responses-api flag for OpenAI compatibility

- Add disable-responses-api flag to CLI completions

- Update zsh completion with new API flag
- Update bash completion options list

- Add fish shell completion for API flag
- Add testpattern to VSCode spell checker dictionary

- Configure disableResponsesAPI in example YAML config
- Enable flag for llama-server compatibility
- Feat: add OpenAI Responses API configuration control via CLI flag

- Add `--disable-responses-api` CLI flag for OpenAI control

- Implement `SetResponsesAPIEnabled` method in OpenAI client
- Configure OpenAI Responses API setting during CLI initialization

- Update default config path to `~/.config/fabric/config.yaml`
- Add OpenAI import to CLI package dependencies
- Docs: Update CHANGELOG after v1.4.256
- Refactor: extract flag parsing logic into separate extractFlag function
- Fix: improve error handling for default config path resolution

- Update `GetDefaultConfigPath` to return error alongside path

- Add proper error handling in flags initialization
- Include debug logging for config path failures

- Move channel close to defer in dryrun SendStream
- Return wrapped errors with context messages

- Handle non-existent config as valid case
- Fix: improve dry run output formatting and config path error handling

- Remove leading newline from DryRunResponse constant

- Add newline separator in SendStream method output
- Add newline separator in Send method output

- Improve GetDefaultConfigPath error handling logic
- Add stderr error message for config access failures

- Return empty string when config file doesn't exist
- Chore: refactor `constructRequest` method for consistency

- Rename `_ConstructRequest` to `constructRequest` for consistency

- Update `SendStream` to use `constructRequest`
- Update `Send` method to use `constructRequest`
- Chore: remove unneeded parenthesis around function call
- Chore: update `Send` method to append `request` to `DryRunResponse`

- Assign `_ConstructRequest` output to `request` variable

- Concatenate `request` with `DryRunResponse` in `Send` method
- Feat: improve flag handling and add default config support

- Map both short and long flags to yaml tags

- Add support for short flag parsing with dashes
- Implement default ~/.fabric.yaml config file detection

- Fix think block suppression in dry run mode
- Add think options to dry run output formatting

- Refactor dry run response construction into helper method
- Return actual response content from dry run client

- Create utility function for default config path resolution
- Docs: Update CHANGELOG after v1.4.255
- Merge branch 'danielmiessler:main' into main
- Chore: add more paths to update-version-andcreate-tag workflow to reduce unnecessary tagging
- Adds generate code rules pattern
Signed-off-by: Roberto Carvajal <roberto.carvajal@gmail.com>
- Docs: Update CHANGELOG after v1.4.253
- Feat: add 'think' tag options for text suppression and completion

- Remove outdated update notes from README

- Add `--suppress-think` option to suppress 'think' tags
- Introduce `--think-start-tag` and `--think-end-tag` options

- Update bash completion with 'think' tag options
- Update fish completion with 'think' tag options
- Docs: Update CHANGELOG after v.1.4.252
- Perf: add regex caching to `StripThinkBlocks` function for improved performance
- Feat: add suppress-think feature to filter AI reasoning output

- Add suppress-think flag to hide thinking blocks

- Configure customizable start and end thinking tags
- Strip thinking content from final response output

- Update streaming logic to respect suppress-think setting
- Add YAML configuration support for thinking options

- Implement StripThinkBlocks utility function for content filtering
- Add comprehensive tests for thinking suppression functionality
- Chore: Update CHANGELOG after v1.4.251
- Ci: update workflow to ignore additional paths during version updates

- Add `data/strategies/**` to paths-ignore list

- Add `cmd/generate_changelog/*.db` to paths-ignore list
- Prevent workflow triggers from strategy data changes

- Prevent workflow triggers from changelog database files
- Docs: Update changelog with v1.4.249 changes
- Chore: add log message for missing PRs in cache
- Feat: preserve PR numbers during version cache merges

- Enhance changelog to associate PR numbers with version tags

- Improve PR number parsing with proper error handling
- Collect all PR numbers for commits between version tags

- Associate aggregated PR numbers with each version entry
- Update cached versions with newly found PR numbers

- Add check for missing PRs to trigger sync if needed
- Fix: improve PR number parsing with proper error handling
- Feat: enhance changelog to correctly associate PR numbers with version tags

- Collect all PR numbers for commits between version tags.

- Associate aggregated PR numbers with each version entry.
- Update cached versions with newly found PR numbers.

- Attribute all changes in a version to relevant PRs.
- Docs: reorganize v1.4.247 changelog to attribute changes to PR #1613
- Chore: update logging output to use os.Stderr
- Fix: improve error handling in plugin registry configuration
- Chore: remove debug logging and sync custom patterns directory configuration

- Remove debug stderr logging from content summarization

- Add custom patterns directory to PatternsLoader configuration
- Ensure consistent patterns directory setup across components

- Clean up unnecessary console output during summarization
- Feat: improve error handling in `plugin_registry` and `patterns_loader`

- Adjust prompt formatting in `summarize.go`

- Add error handling for `CustomPatterns` configuration
- Enhance error messages in `patterns_loader`

- Check for patterns in multiple directories
- Chore: reorder plugin configuration sequence in PluginRegistry.Configure method

- Move CustomPatterns.Configure() before PatternsLoader.Configure()

- Adjust plugin initialization order in Configure method
- Ensure proper dependency sequence for pattern loading
- Fix: improve git walking termination and error handling in changelog generator

- Add storer import for proper git iteration control

- Use storer.ErrStop instead of nil for commit iteration termination
- Handle storer.ErrStop as expected condition in git walker

- Update cache comment to clarify Unreleased version skipping
- Change custom patterns warning to stderr output

- Add storer to VSCode spell checker dictionary
- Chore: clean up changelog and add debug logging for content length validation
- Feat: enhance changelog generation with incremental caching and improved AI summarization

- Add incremental processing for new Git tags since cache

- Implement `WalkHistorySinceTag` method for efficient history traversal
- Cache new versions and commits after AI processing

- Update AI summarization prompt for better release note formatting
- Remove conventional commit prefix stripping from commit messages

- Add custom patterns directory support to plugin registry
- Generate unique patterns file including custom directory patterns

- Improve session string formatting with switch statement
- Docs: update README for GraphQL optimization and AI summary features

- Detail GraphQL API usage for faster PR fetching

- Introduce AI-powered summaries via Fabric integration
- Explain content-based caching for AI summaries

- Document support for loading secrets from .env files
- Add usage examples for new AI summary feature

- Clarify project license is The MIT License
- Docs: Update CHANGELOG
- Feat: add AI-powered changelog generation with high-performance Go tool and comprehensive caching

- Add high-performance Go changelog generator with GraphQL integration

- Implement SQLite-based persistent caching for incremental updates
- Create one-pass git history walking algorithm with concurrent processing

- Add comprehensive CLI with cobra framework and tag-based caching
- Integrate AI summarization using Fabric CLI for enhanced output

- Support batch PR fetching with GitHub Search API optimization
- Add VSCode configuration with spell checking and markdown linting

- Include extensive documentation with PRD and README files
- Implement commit-PR mapping for lightning-fast git operations

- Add content hashing for change detection and cache optimization
- Chore: optimize model ID extraction and remove redundant comment

- Remove duplicate comment about reading response body

- Preallocate slice capacity in extractModelIDs function
- Initialize modelIDs slice with known capacity
- Fix: improve error message truncation in DirectlyGetModels method

- Add proper bounds checking for response body truncation

- Prevent slice out of bounds errors in error messages
- Add ellipsis indicator when response body is truncated

- Improve error message clarity for debugging purposes
- Refactor: clean up HTTP request handling and improve error response formatting

- Remove unnecessary else block in HTTP request creation

- Move header setting outside conditional block for clarity
- Add TODO comment about reusing HTTP client instance

- Truncate error response output to prevent excessive logging
- Chore: refactor DirectlyGetModels to read response body once

- Read response body once for efficiency
- Use io.ReadAll for response body

- Unmarshal json from bodyBytes
- Return error with raw response bytes

- Improve error handling for json parsing
- Fix: increase error response limit and simplify model extraction logic

- Increase error response limit from 500 to 1024 bytes

- Add documentation comment for ErrorResponseLimit constant
- Remove unnecessary error return from extractModelIDs function

- Fix return statements in DirectlyGetModels parsing logic
- Add TODO comment for proper context handling

- Simplify model ID extraction without error propagation
- Fix: improve error message in DirectlyGetModels to include provider name

- Add provider name to API base URL error message

- Enhance error context for better debugging experience
- Include GetName() method call in error formatting
- Feat: add context support to DirectlyGetModels method

- Add context parameter to DirectlyGetModels method signature

- Add nil context check with Background fallback
- Extract magic number 500 into errorResponseLimit constant

- Update DirectlyGetModels call to pass context.Background
- Import context package in providers_config.go file
- Refactor: replace string manipulation with url.JoinPath for models endpoint construction
- Refactor: improve OpenAI compatible models API client with timeout and cleaner parsing
- Refactor: extract model ID parsing logic into reusable helper function
- Fix: enhance error messages in OpenAI compatible models endpoint with response body details
- Feat: add direct model fetching support for non-standard providers

- Add `DirectlyGetModels` function to handle non-standard API responses
- Add Together provider configuration to ProviderMap

- Implement fallback to direct model fetching when standard method fails
- Fix: broken image link
- Docs: update file paths to reflect new data directory structure

- Move fabric logo image path to docs directory

- Update patterns directory reference to data/patterns location
- Update strategies directory reference to data/strategies location

- Fix create_coding_feature README path reference
- Update code_helper install path to cmd directory
- Shell: fix typo
- Fix: improve error handling and temporary file management in patterns loader

- Replace println with fmt.Fprintln to stderr for errors

- Use os.MkdirTemp for secure temporary directory creation
- Remove unused time import from patterns loader

- Add proper error wrapping for file operations
- Handle RemoveAll errors with warning messages

- Improve error messages with context information
- Add explicit error checking for cleanup operations
- Chore: improve error handling for scraping configuration in `tools.go`
- Chore: enhance error handling and early returns in CLI

- Add early return if registry is nil to prevent panics.
- Introduce early return for non-chat tool operations.

- Update error message to use original input on HTML readability failure.
- Enhance error wrapping for playlist video fetching.

- Modify temp patterns folder name with timestamp for uniqueness.
- Improve error handling for patterns directory access.
- Update-mod: fix generation path
- Shell: rename command
- Nix:pkgs:fabric: use self reference
- Chore: remove fabric binary
- Chore: update command handlers to return 'handled' boolean

- Add `handled` boolean return to command handlers

- Modify `handleSetupAndServerCommands` to use `handled`
- Update `handleConfigurationCommands` with `handled` logic

- Implement `handled` return in `handleExtensionCommands`
- Revise `handleListingCommands` to support `handled` return

- Adjust `handleManagementCommands` to return `handled`
- Feat: refactor CLI to modularize command handling

- Extract chat processing logic into separate function

- Create modular command handlers for setup, configuration, listing, management, and extensions
- Improve patterns loader with migration support and better error handling

- Simplify main CLI logic by delegating to specialized handlers
- Enhance code organization and maintainability

- Add tool processing for YouTube and web scraping functionality
- Chore: update workflow paths to reflect directory structure change

- Modify trigger path to `data/patterns/**`

- Update `git diff` command to new path
- Change zip command to include `data/patterns/` directory
- Docs: Update README with new go install commands
- Fix: minor edit
- Docs: update restructure guide with Homebrew and go install details

- Document required Homebrew formula update for new structure.

- Add new `go install` commands for all tools.
- Specify new build path is `./cmd/fabric`.

- Include link to the draft Homebrew PR.
- Feat: add new patterns for content tagging and cognitive bias analysis

- Fix static directory path in extract_patterns.py script

- Add apply_ul_tags pattern for content categorization
- Add t_check_dunning_kruger pattern for bias analysis

- Update pattern descriptions with new entries
- Sync web static data with latest patterns

- Include pattern extracts for new functionality
- Support standardized content topic classification

- Enable cognitive bias identification capabilities
- Docs: update project restructuring status and reorganize pattern scripts

- Mark all 10 migration steps as completed

- Add restructuring completion status section
- Move pattern generation scripts to pattern_descriptions

- Update completion checkmarks throughout migration plan
- Document remaining external packaging verification tasks

- Consolidate pattern description files under new directory
- Confirm all binaries compile and tests pass

- Note standard Go project layout achieved
- Refactor: move common package to domain and util packages for better organization

- Move domain types from common to domain package

- Move utility functions from common to util package
- Update all import statements across codebase

- Reorganize OAuth storage functionality into util package
- Move file management functions to domain package

- Update test files to use new package structure
- Maintain backward compatibility for existing functionality
- Refactor: alias server package import as restapi for clarity

- Rename the `server` package import to `restapi`.

- Improve code readability and prevent naming collisions.
- Refactor: restructure project to align with standard Go layout

- Introduce `cmd` directory for all main application binaries.

- Move all Go packages into the `internal` directory.
- Rename the `restapi` package to `server` for clarity.

- Consolidate patterns and strategies into a new `data` directory.
- Group all auxiliary scripts into a new `scripts` directory.

- Move all documentation and images into a `docs` directory.
- Update all Go import paths to reflect the new structure.

- Adjust CI/CD workflows and build commands for new layout.
- Chore: refactor token path to use `authTokenIdentifier`
- Test: add comprehensive OAuth testing suite for Anthropic plugin

- Add OAuth test file with 434 lines coverage

- Create mock token server for safe testing
- Implement PKCE generation and validation tests

- Add token expiration logic verification tests
- Create OAuth transport round-trip testing

- Add benchmark tests for performance validation
- Implement helper functions for test token creation

- Add comprehensive error path testing scenarios
- Fix: update `RefreshToken` to use `tokenIdentifier` parameter
- Refactor: replace hardcoded "claude" with configurable `authTokenIdentifier` parameter

- Replace hardcoded "claude" string with `authTokenIdentifier` constant

- Update `RunOAuthFlow` to accept token identifier parameter
- Modify `RefreshToken` to use configurable token identifier

- Update `exchangeToken` to accept token identifier parameter
- Enhance `getValidToken` to use parameterized token identifier

- Add token refresh attempt before full OAuth flow
- Improve OAuth flow with existing token validation
- Chore: improve error comparison in `TestChatter_Send_StreamingErrorPropagation`
- Chore: remove redundant channel closure in `Send` method

- Remove redundant `close(responseChan)` in `Send` method

- Update `SendStream` to close `responseChan` properly
- Modify test to reflect channel closure logic
- Chore: rename `doneChan` to `done` and add streaming aggregation test

- Rename `doneChan` variable to `done` for consistency

- Add `streamChunks` field to mock vendor struct
- Implement chunk sending logic in mock SendStream method

- Add comprehensive streaming success aggregation test case
- Verify message aggregation from multiple stream chunks

- Test assistant response role and content validation
- Ensure proper session handling in streaming scenarios
- Feat: add test for Chatter's Send method error propagation

- Implement mockVendor for testing ai.Vendor interface

- Add TestChatter_Send_StreamingErrorPropagation test case
- Verify error propagation in Chatter's Send method

- Ensure session returns even on streaming error
- Create temporary database for testing Chatter functionality
- Chore: rename channels for clarity in `Send` method

- Rename `done` to `doneChan` for clarity

- Adjust channel closure for `doneChan`
- Update channel listening logic to use `doneChan`
- Refactor: rename `channel` variable to `responseChan` for better clarity in streaming logic

- Rename `channel` variable to `responseChan` for clarity

- Update channel references in goroutine defer statements
- Pass renamed channel to `SendStream` method call

- Maintain consistent naming throughout streaming flow
- Chore: close `channel` after sending stream in `Send`

- Add `channel` closure after sending stream

- Ensure resource cleanup in `Send` method
- Chore: refactor error handling and response aggregation in `Send`

- Simplify response aggregation loop in `Send`

- Remove redundant select case for closed channel
- Streamline error checking from `errChan`

- Ensure goroutine completion before returning
- Chore: enhance `Chatter.Send` method with proper goroutine synchronization

- Add `done` channel to track goroutine completion.
- Replace `errChan` closure with `done` channel closure.

- Ensure main loop waits for goroutine on channel close.
- Synchronize error handling with `done` channel wait.
- Refactor: use select to handle stream and error channels concurrently

- Replace for-range loop with a non-blocking select statement.

- Process message and error channels concurrently for better handling.
- Improve the robustness of streaming error detection.

- Exit loop cleanly when the message channel closes.
- Chore: simplify error handling in streaming chat response by removing unnecessary select statement
- Fix: improve error handling in streaming chat functionality

- Add dedicated error channel for stream operations

- Separate error handling from message streaming logic
- Check for streaming errors after channel closure

- Close error channel properly in goroutine cleanup
- Remove error messages from message stream channel

- Add proper error propagation for stream failures
- Refactor: extract vendor token identifier constant and remove redundant configure call

- Extract vendor token identifier into named constant

- Remove redundant Configure() call from IsConfigured method
- Use constant for token validation consistency

- Improve code maintainability with centralized identifier
- Feat: add vendor configuration validation and OAuth auto-authentication

- Add IsConfigured check to vendor configuration loop

- Implement IsConfigured method for Anthropic client validation
- Remove conditional API key requirement based on OAuth

- Add automatic OAuth flow when no valid token
- Validate both API key and OAuth token configurations

- Simplify API key setup question logic
- Add token expiration checking with 5-minute buffer
- Fix: add conditional check for TopP parameter in OpenAI client

- Add zero-value check before setting TopP parameter

- Prevent sending TopP when value is zero
- Apply fix to both chat completions method

- Apply fix to response parameters method
- Ensure consistent parameter handling across OpenAI calls
- Chore: enhance bug report template with detailed system info and installation method fields

- Add detailed instructions for bug reproduction steps

- Include operating system dropdown with specific architectures
- Add OS version textarea with command examples

- Create installation method dropdown with all options
- Replace version checkbox with proper version output field

- Improve formatting and organization of form sections
- Add helpful links to installation documentation
- Fix: make custom patterns persist correctly
- Chore: improve directory creation logic in `configure` method

- Add `fmt` package for logging errors

- Check directory existence before creating
- Log error without clearing directory value
- Refactor: move custom patterns directory initialization to Configure method

- Move custom patterns directory logic to Configure method
- Initialize CustomPatternsDir after loading .env file

- Add alphabetical sorting to pattern names retrieval
- Override ListNames method for PatternsEntity class

- Improve pattern listing with proper error handling
- Ensure custom patterns loaded after environment configuration
- Docs: add comprehensive custom patterns setup and usage guide

- Add custom patterns directory setup instructions

- Document priority system for custom vs built-in patterns
- Include step-by-step custom pattern creation workflow

- Explain update-safe custom pattern storage
- Add table of contents entries for new sections

- Document seamless integration with existing fabric commands
- Clarify privacy and precedence behavior for custom patterns
- Feat: add custom patterns directory support with environment variable configuration

- Add custom patterns directory support via environment variable

- Implement custom patterns plugin with registry integration
- Override main patterns with custom directory patterns

- Expand home directory paths in custom patterns config
- Add comprehensive test coverage for custom patterns functionality

- Integrate custom patterns into plugin setup workflow
- Support pattern precedence with custom over main patterns
- Fix: remove duplicate API key setup question in Anthropic client
- Refactor: extract OAuth functionality from anthropic client to separate module

- Remove OAuth transport implementation from main client

- Extract OAuth flow functions to separate module
- Remove unused imports and constants from client

- Replace inline OAuth transport with NewOAuthTransport call
- Update runOAuthFlow to exported RunOAuthFlow function

- Clean up token management and refresh logic
- Simplify client configuration by removing OAuth internals
- Feat: add OAuth login support for Anthropic API configuration
- Feat: remove OAuth flow functions for simplified token handling
- Chore: simplify base URL configuration in `configure` method

- Remove redundant base URL trimming logic

- Append base URL directly without modification
- Eliminate conditional check for API version suffix
- Feat: enhance OAuth authentication flow with automatic re-authentication and timeout handling

- Add automatic OAuth flow initiation when no token exists

- Implement fallback re-authentication when token refresh fails
- Add timeout contexts for OAuth and refresh operations

- Create context-aware OAuth flow and token exchange functions
- Enhance error handling with graceful authentication recovery

- Add user input timeout protection for authorization codes
- Preserve refresh tokens during token exchange operations
- Refactor: remove OAuth endpoint logic and standardize on v2 API endpoint

- Remove OAuth-specific v1 endpoint handling logic

- Standardize all API calls to use v2 endpoint
- Simplify baseURL configuration by removing conditional branching

- Update endpoint logic to always append v2 suffix
- Feat: implement OAuth token refresh and persistent storage for Claude authentication

- Add automatic OAuth token refresh when expired

- Implement persistent token storage using common OAuth storage
- Remove deprecated AuthToken setting from client configuration

- Add token validation with 5-minute expiration buffer
- Create refreshToken function for seamless token renewal

- Update OAuth flow to save complete token information
- Enhance error handling for OAuth authentication failures

- Simplify client configuration by removing manual token management
- Feat: add OAuth authentication support for Anthropic Claude

- Move golang.org/x/oauth2 from indirect to direct dependency
- Add OAuth login option for Anthropic client

- Implement PKCE OAuth flow with browser integration
- Add custom HTTP transport for OAuth Bearer tokens

- Support both API key and OAuth authentication methods
- Add Claude Code system message for OAuth sessions

- Update REST API to handle OAuth tokens
- Improve environment variable name sanitization with regex
- Feat: add advanced image generation parameters for OpenAI models

- Add four new image generation CLI flags

- Implement validation for image parameter combinations
- Support size, quality, compression, and background controls

- Add comprehensive test coverage for new parameters
- Update shell completions for new image options

- Enhance README with detailed image generation examples
- Fix PowerShell code block formatting issues
- Refactor: extract supported models list to shared constant for image generation validation

• Extract hardcoded model lists into shared constant
• Create ImageGenerationSupportedModels variable for reusability
• Update supportsImageGeneration function to use shared constant
• Refactor error messages to reference centralized model list
• Add documentation comment for supported models variable
• Import strings package in test file
• Consolidate duplicate model validation logic across files
- Merge branch 'main' into 0704-image-tool-model-validation
- Feat: add model validation for image generation support

- Add model field to `BuildChatOptions` method

- Implement `supportsImageGeneration` function for model checks
- Validate model supports image generation in `sendResponses`

- Remove `mars-colony.png` from repository
- Add tests for `supportsImageGeneration` function

- Validate image file support in `TestModelValidationLogic`
- Feat: add image file validation and format detection for image generation

• Add image file path validation with extension checking
• Implement dynamic output format detection from file extensions
• Update BuildChatOptions method to return error for validation
• Add comprehensive test coverage for image file validation
• Upgrade YAML library from v2 to v3
• Update shell completions to reflect supported image formats
• Add error handling for existing file conflicts
• Support PNG, JPEG, JPG, and WEBP image formats
- Addded tutorial as a tag.
- Chore: refactor image generation constants for clarity and reuse

- Define `ImageGenerationResponseType` constant for response handling

- Define `ImageGenerationToolType` constant for tool type usage
- Update `addImageGenerationTool` to use defined constants

- Refactor `extractAndSaveImages` to use response type constant
- Feat: add web search and image file support to fabric CLI

- Add web search tool for Anthropic and OpenAI models

- Add search location parameter for web search results
- Add image file output option with format support

- Update zsh completion with new search and image flags
- Update bash completion with new option handling logic

- Update fish completion with search and image descriptions
- Support PNG, JPG, JPEG, GIF, BMP image formats
- Feat: add image generation support with OpenAI image generation model

- Add `--image-file` flag for saving generated images

- Implement image generation tool integration with OpenAI
- Support image editing with attachment input files

- Add comprehensive test coverage for image features
- Update documentation with image generation examples

- Fix HTML formatting issues in README
- Improve PowerShell code block indentation

- Clean up help text formatting and spacing
- Fixed ul tag applier.
- Updated ul tag prompt.
- Added the UL tags pattern.
- Docs: update README with new web search feature details
- Feat: add web search tool support for OpenAI models with citation formatting

- Enable web search tool for OpenAI models

- Add location parameter support for search results
- Extract and format citations from search responses

- Implement citation deduplication to avoid duplicates
- Add comprehensive test coverage for search functionality

- Update CLI flag description to include OpenAI
- Format citations as markdown links with sources
- Chore: refactor `Send` method to optimize string building

- Add `sourcesHeader` constant for citation section title.

- Use `strings.Builder` to construct result efficiently.
- Append sources header and citations in result builder.

- Update `ret` to use constructed string from builder.
- Chore: remove unused web-search tool parameters for simplification

- Remove unused `AllowedDomains` and `MaxUses` parameters

- Simplify `webTool` definition in `buildMessageParams` method
- Refactor: extract web search tool constants in anthropic plugin

- Add webSearchToolName constant for tool identification

- Add webSearchToolType constant for tool versioning
- Replace hardcoded string literals with named constants

- Improve code maintainability through constant extraction
- Chore: update `formatOptions` to include search options display

- Add search option status to `formatOptions`

- Include `SearchLocation` in formatted output if specified
- Feat: add web search tool support for Anthropic models

- Add --search flag to enable web search

- Add --search-location for timezone-based results
- Pass search options through ChatOptions struct

- Implement web search tool in Anthropic client
- Format search citations with sources section

- Add comprehensive tests for search functionality
- Remove plugin-level web search configuration
- Merge branch 'main' of <https://github.com/amancioandre/Fabric>
- Fix: sections as heading 1, typos
- Merge branch 'danielmiessler:main' into main
- Feat: adds pattern telos check dunning kruger
- Feat: handle JSONDecodeError in `load_existing_file` gracefully

- Add JSONDecodeError handling with warning message.

- Initialize with empty list on JSON decode failure.
- Reorder pattern processing to reduce redundant logs.

- Remove redundant directory check logging.
- Ensure new pattern processing is logged correctly.
- Feat: add new patterns for code review, alpha extraction, and server analysis

- Add `review_code`, `extract_alpha`, and `extract_mcp_servers` patterns.
- Refactor the pattern extraction script for improved clarity.

- Add docstrings and specific error handling to script.
- Improve formatting in the pattern management README.

- Fix typo in the `analyze_bill_short` pattern description.
- Feat: add comprehensive code review pattern for systematic analysis

- Add new code review system prompt

- Define principal engineer reviewer role
- Include systematic analysis framework

- Specify markdown output format
- Add prioritized recommendations section

- Include detailed feedback structure
- Provide example Python review

- Cover security, performance, readability
- Add error handling guidelines
- Chore: update GitHub Actions to use bash shell in release job

- Adjust repository_dispatch type spacing for consistency

- Use bash shell for creating release if absent
- Updated alpha post.
- Feat(openai): add support for multi-content user messages in chat completions

- Enhance user message conversion to support multi-content.

- Add capability to process image URLs in messages.
- Build multi-part messages with both text and images.
- Chore: update `NewClient` to use `NewClientCompatibleWithResponses`

- Modify `NewClient` to call `NewClientCompatibleWithResponses`

- Add support for response handling in client initialization
- Feat: simplify `supportsResponsesAPI`
- Refactor: extract common message conversion logic to reduce duplication

- Extract shared message conversion to convertMessageCommon

- Reuse logic between chat and response APIs
- Maintain existing text-only behavior for chat

- Support multi-content messages in response API
- Reduce code duplication across converters

- Preserve backward compatibility for both APIs
- Fix: move channel close to defer statement in OpenAI streaming methods

- Move close(channel) to defer statement

- Ensure channel closes even on errors
- Apply fix to sendStreamChatCompletions method

- Apply fix to sendStreamResponses method
- Improve error handling reliability

- Prevent potential channel leaks
- Feat: add chat completions API support for OpenAI-compatible providers

- Add chat completions API fallback for non-Responses API providers

- Implement `sendChatCompletions` and `sendStreamChatCompletions` methods
- Introduce `buildChatCompletionParams` to construct API request parameters

- Add `ImplementsResponses` flag to track provider API capabilities
- Update provider configurations with Responses API support status

- Enhance `Send` and `SendStream` methods to use appropriate API endpoints
- Feat: migrate OpenAI plugin to use new responses API instead of chat completions

- Replace chat completions with responses API
- Update message conversion to new format

- Refactor streaming to handle event types
- Remove frequency and presence penalty params

- Replace seed parameter with max tokens
- Update test cases for new API

- Add response text extraction method
- Updated extract alpha.
- Updated extract alpha.
- Added extract_alpha as kind of an experiment.
- Refactor: abstract chat message structs and migrate to official openai-go SDK

- Introduce local `chat` package for message abstraction

- Replace sashabaranov/go-openai with official openai-go SDK
- Update OpenAI, Azure, and Exolab plugins for new client

- Refactor all AI providers to use internal chat types
- Decouple codebase from third-party AI provider structs

- Replace deprecated `ioutil` functions with `os` equivalents
- Chore: improve release creation to gracefully handle pre-existing tags.

- Check if a release exists before attempting creation.

- Suppress error output from `gh release view` command.
- Add an informative log when release already exists.
- Docs: add DeepWiki badge and fix minor typos in README

- Add DeepWiki badge to README header

- Fix typo "chatbots" to "chat-bots"
- Correct "Perlexity" to "Perplexity"

- Fix "distro" to "Linux distribution"
- Add alt text to contributor images

- Update dependency versions in go.mod
- Remove unused soup dependency
- Fix typos on README.md
- Feat: add support for new OpenAI search and research model variants

- Add slices import for array operations

- Define new search preview model names
- Add mini search preview variants

- Include deep research model support
- Add June 2025 dated model versions

- Replace hardcoded check with slices.Contains
- Support both prefix and exact model matching
- Add extract_mcp_servers pattern
New pattern to extract mentions of MCP (Model Context Protocol) servers from content. Identifies server names, features, capabilities, and usage examples.
🤖 Generated with [Claude Code](<https://claude.ai/code)>
Co-Authored-By: Claude <noreply@anthropic.com>
- Chore: fix endpoint calls from frontend
- Feat: add dedicated YouTube transcript API endpoint

- Add new YouTube handler for transcript requests

- Create `/youtube/transcript` POST endpoint route
- Add request/response types for YouTube API

- Support language and timestamp options
- Update frontend to use new endpoint

- Remove chat endpoint dependency for transcripts
- Validate video vs playlist URLs properly
- Refactor(ai): unify assistant and user message formatting in dryrun

- Unify assistant and user message formatting logic.

- Use `formatMultiContentMessage` for assistant role messages.
- Improve dryrun support for multi-part message content.
- Fix: correctly combine text and attachments in raw mode sessions

- Combine user text and attachments into MultiContent.

- Preserve existing non-text parts like images.
- Use standard content field for text-only messages.
- Feat: add MultiContent support to chat message construction in raw mode
- Refactor: extract message and option formatting logic into reusable methods

- Extract multi-content message formatting to dedicated method

- Create formatMessages method for all message types
- Add formatOptions method for chat options display

- Replace inline formatting with strings.Builder usage
- Reduce code duplication between Send and SendStream

- Improve code organization and maintainability
- Fix(chatter): prevent duplicate user message when applying patterns

- Prevent adding user message twice when using patterns.

- Ensure multi-part content is always included in session.
- Chore: fix formatting.
- Chore: clean up comments in `chatter.go` for clarity
- Chore: simplify user message appending logic in BuildSession

- Remove conditional check for pattern name in message appending.
- Always append user message if it exists in request.
- Feat: enhance dryrun client to display multi-content user messages

- Handle multi-content messages for the user role.

- Display image URLs from user messages in output.
- Update both `Send` and `SendStream` methods.

- Retain existing behavior for simple text messages.
- Feat: allow combining user messages and attachments with patterns

- Allow user messages and attachments with patterns.
- Append user message to session regardless of pattern.

- Refactor chat request builder for improved clarity.
- Chore: reformat `pattern_descriptions.json` to improve readability

- Reformat JSON `tags` array to display on new lines.

- Update `write_essay` pattern description for clarity.
- Apply consistent formatting to both data files.
- Chore: Fixes caught by review
- Chore: rename essay patterns to clarify Paul Graham style and author variable usage

- Rename `write_essay` to `write_essay_pg` for Paul Graham style

- Rename `write_essay_by_author` to `write_essay` with author variable
- Update pattern descriptions to reflect naming changes

- Fix duplicate `write_essay_pg` entry in pattern descriptions
- Feat: add  new pattern and update pattern metadata files.

- Add tags and descriptions for five new creative and analytical patterns.

- Introduce `analyze_terraform_plan` for infrastructure review.
- Add `write_essay_by_author` for stylistic writing.

- Include `summarize_board_meeting` for corporate notes.
- Introduce `create_mnemonic_phrases` for memory aids.

- Update and clean pattern description data files.
- Sort the pattern explanations list alphabetically.
- Merge branch 'danielmiessler:main' into main
- Chore: refactor ProviderMap for dynamic URL template handling

- Add `os` and `strings` packages to imports

- Implement dynamic URL handling with environment variables
- Refactor provider configuration to support URL templates

- Reorder providers for consistent key order in ProviderMap
- Extract and parse template variables from BaseURL

- Use environment variables or default values for templates
- Replace template with actual values in BaseURL
- Chore: refactor ProviderMap for dynamic URL template handling

- Add `os` and `strings` packages to imports

- Implement dynamic URL handling with environment variables
- Refactor provider configuration to support URL templates

- Reorder providers for consistent key order in ProviderMap
- Extract and parse template variables from BaseURL

- Use environment variables or default values for templates
- Replace template with actual values in BaseURL
- Chore: refactor Bedrock client to improve error handling and add interface compliance

- Add ai.Vendor interface implementation check

- Improve error handling with wrapped errors
- Add AWS region validation logic

- Fix resource cleanup in SendStream
- Add nil checks for response parsing

- Update context usage to Background()
- Add user agent constants

- Enhance code documentation
- Bedrock region handling - updated to set region value correctly if it exists in the config
- Bedrock region handling - updated to fix bad pointer reference
- Fixed bedrock region handling
- Updated hasAWSCredentials to also check for AWS_DEFAULT_REGION when access keys are configured in the environment
- Updated paper analyzer.
- Updated paper analyzer. Went back to my own format.
- Chore: removed a directory of raycast scripts sitting in the patterns/ directory
- Refactor(ChatService): clean up message stream and pattern output methods

- Refactor `cleanPatternOutput` to use a dedicated return variable.
- Hoist `processResponse` function for improved stream readability.

- Remove unnecessary whitespace and trailing newlines from file.
- Updated paper analyzer.
- Updated paper analyzer.
- Feat: add `ApplyPattern` route for applying patterns with variables

- Create `PatternApplyRequest` struct for request body parsing
- Implement `ApplyPattern` method for POST /patterns/:name/apply

- Register manual routes for pattern operations in `NewPatternsHandler`
- Refactor `Get` method to return raw pattern content

- Merge query parameters with request body variables in `ApplyPattern`
- Use `StorageHandler` for pattern-related storage operations
- Feat: add pattern variables support to REST API chat endpoint

- Add Variables field to PromptRequest struct

- Pass pattern variables through chat handler
- Create API variables documentation example

- Add pattern variables UI in web interface
- Create pattern variables store in Svelte

- Include variables in chat service requests
- Add JSON textarea for variable input
- Updated sanitization instructions.
- Updated markdown cleaner.
- Updated markdown cleaner.
- Feat: add citation support to perplexity AI responses

- Add citation extraction from API responses

- Append citations section to response content
- Format citations as numbered markdown list

- Handle citations in streaming responses
- Store last response for citation access

- Add citations after stream completion
- Maintain backward compatibility with responses
- Update README.md
- Update README.md
Updated readme.
- Update README.md
An update to the intro text, describing Fabric's utility to most people.
- Chore: update README with Perplexity AI support instructions

- Add instructions for configuring Perplexity AI with Fabric
- Include example command for querying Perplexity AI

- Retain existing instructions for YouTube transcription changes
- Feat: add Perplexity AI provider support with token limits and streaming

- feat: Add `MaxTokens` field to `ChatOptions` struct for response control

- feat: Integrate Perplexity client into core plugin registry initialization
- build: Add perplexity-go/v2 dependency to enable API interactions

- feat: Implement stream handling in Perpexlty client using sync.WaitGroup
- fix: Correct parameter types for penalty options in API requests
-
<<https://github.com/sgaunet/perlexipty-go>> - Client library used
- Check for AWS_PROFILE or AWS_ROLE_SESSION_NAME environment variables
- Refactor: extract common yt-dlp logic to reduce code duplication in YouTube plugin

- Extract shared yt-dlp logic into tryMethodYtDlpInternal helper

- Add processVTTFileFunc parameter for flexible VTT processing
- Implement language matching for 2-char language codes

- Refactor tryMethodYtDlp to use new helper function
- Refactor tryMethodYtDlpWithTimestamps to use helper

- Reduce code duplication between transcript methods
- Maintain existing functionality with cleaner structure
- Updated extract insights.
- Updated extract insights.
- Feat: add AWS credential detection for Bedrock client initialization

- Add hasAWSCredentials helper function

- Check for AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
- Look for AWS shared credentials file

- Support custom AWS_SHARED_CREDENTIALS_FILE path
- Default to ~/.aws/credentials location

- Only initialize Bedrock client if credentials exist
- Prevent AWS SDK credential search failures
- Updated prompt.
- Dynamically fetch and list available foundation models and inference profiles
- Updated markdown sanitizer.
- Chore: remove duplicate/outdated patterns
- Updated markdown cleaner.
- Updated markdown cleaner.
- Ci: improve version update workflow to prevent race conditions

- Add concurrency control to prevent simultaneous runs

- Pull latest main branch changes before tagging
- Fetch all remote tags before calculating version
- Feat: add Save method to PatternsEntity for persisting patterns to filesystem

- Add Save method to PatternsEntity struct

- Create pattern directory with proper permissions
- Write pattern content to system pattern file

- Add comprehensive test for Save functionality
- Verify directory creation and file contents

- Handle errors for directory and file operations
- Feat: create mnemonic phrase pattern
Add a new pattern for generating mnemonic phrases from diceware words. This includes two markdown files defining the user guide, and system implementation details.
- Add Bedrock plugin
This commits adds support for using Amazon Bedrock within fabric.
- *fix: replace Unix-specific file operations with cross-platform alternatives

- Replace hardcoded `/tmp` with `os.TempDir()` for paths

- Use `filepath.Join()` instead of string concatenation
- Remove Unix `find` command dependency completely

- Add new `findVTTFiles()` method using `filepath.Walk()`
- Make VTT file discovery work on Windows

- Improve error handling for file operations
- Maintain backward compatibility with existing functionality
- Feat: cleanup after `yt-dlp` addition

- Update README with yt-dlp requirement for transcripts
- Ensure the errors are clear and actionable.
- Refactor: replace web scraping with yt-dlp for YouTube transcript extraction

- Remove unreliable YouTube API scraping methods

- Add yt-dlp integration for transcript extraction
- Implement VTT subtitle parsing functionality

- Add timestamp preservation for transcripts
- Remove soup HTML parsing dependency

- Add error handling for missing yt-dlp
- Create temporary directory management

- Support multiple subtitle format fallbacks
- Fix: fix web search tool location
- Chore(deps): bump brace-expansion
Bumps the npm_and_yarn group with 1 update in the /web directory: [brace-expansion](<https://github.com/juliangruber/brace-expansion).>

Updates `brace-expansion` from 1.1.11 to 1.1.12

- [Release notes](<https://github.com/juliangruber/brace-expansion/releases)>
- [Commits](<https://github.com/juliangruber/brace-expansion/compare/1.1.11...v1.1.12)>
updated-dependencies:
- dependency-name: brace-expansion
  dependency-version: 1.1.12
  dependency-type: indirect
  dependency-group: npm_and_yarn
Signed-off-by: dependabot[bot] <support@github.com>
- Fix: Add configurable HTTP timeout for Ollama client
Add a new setup question to configure the HTTP timeout duration for
Ollama requests. The default value is set to 20 minutes.
- Feat: search tool result collection
- Feat: search tool working
- GitButler Workspace Commit
This is a merge commit the virtual branches in your workspace.
Due to GitButler managing multiple virtual branches, you cannot switch back and
forth between git branches and virtual branches easily.
If you switch to another branch, GitButler will need to be reinitialized.
If you commit on this branch, GitButler will throw it away.
Here are the branches that are currently applied:
- improve-create-prd (refs/gitbutler/improve-create-prd)
For more information about what we're doing here, check out our docs:
<https://docs.gitbutler.com/features/virtual-branches/integration-branch>
- Feat: Enhance the PRD Generator's identity and purpose
The changes in this commit expand the identity and purpose of the PRD Generator
to provide more clarity on its role and the expected output. The key changes
include:
- Defining the Generator's purpose as transforming product ideas into a
  structured PRD that ensures clarity, alignment, and precision in product
  planning and execution.

- Outlining the key sections typically found in a PRD that the Generator should
  cover, such as Overview, Objectives, Target Audience, Features, User Stories,
  Functional and Non-functional Requirements, Success Metrics, and Timeline.

- Providing more detailed instructions on the expected output format, structure,
  and content, including the use of Markdown, labeled sections, bullet points,
  tables, and highlighting of priorities or MVP features.
- Feat: add Terraform plan analyzer pattern for infrastructure change assessment

- Create new pattern for analyzing Terraform plans
- Add identity defining expert plan analyzer role

- Include focus on security, cost, and compliance
- Define three output sections for summaries

- Specify 20-word sentence summary requirement
- List 10 critical changes with word limits

- Include 5 key takeaways section format
- Add markdown formatting output instructions

- Require numbered lists over bullet points
- Prohibit warnings and duplicate content
- Feat: add AIML provider to OpenAI compatible providers configuration

- Add AIML provider configuration

- Set AIML base URL to api.aimlapi.com/v1
- Expand supported OpenAI compatible providers list

- Enable AIML API integration support
- Updated output.
- Updated output.
- Updated output.
- Updated output.
- Updated output.
- Updated output.
- Updated output.
- Added simpler paper analyzer, updated the output.
- Added simpler paper analyzer.
- Feat: upgrade PDF.js to v4.2 and refactor worker initialization

- Add `.browserslistrc` to define target browser versions.
- Upgrade `pdfjs-dist` dependency from v2.16 to v4.2.67.

- Upgrade `nanoid` dependency from v4.0.2 to v5.0.9.
- Introduce `pdf-config.ts` for centralized PDF.js worker setup.

- Refactor `PdfConversionService` to use new PDF worker configuration.
- Add static `pdf.worker.min.mjs` to serve PDF.js worker.

- Update Vite configuration for ESNext build target and PDF.js.
- Feat: add centralized environment configuration for Fabric base URL

- Create environment config module for URL handling
- Add getFabricBaseUrl() function with server/client support

- Add getFabricApiUrl() helper for API endpoints
- Configure Vite to inject FABRIC_BASE_URL client-side

- Update proxy targets to use environment variable
- Add TypeScript definitions for window config

- Support FABRIC_BASE_URL env var with fallback
- Fix typo in script name
- Docs: reorganize web documentation and add installation scripts

- Move legacy documentation files to web/legacy/

- Update web README with installation instructions
- Add convenience scripts for npm and pnpm installation

- Update all package dependencies to latest versions
- Add PDF-to-Markdown installation steps to README

- Remove duplicate documentation files
- Update meeting summary template with word count requirement
AI:
Add minimum word count for context section in board summary
- Merge branch 'danielmiessler:main' into main
- Add board meeting summary pattern template
- Feat: add automatic raw mode detection for specific AI models

- Add model-specific raw mode detection logic

- Check Ollama llama2/llama3 models for raw mode
- Check OpenAI o1/o3/o4 models for raw mode

- Use model from options or default chatter
- Auto-enable raw mode when vendor requires it

- Import strings package for prefix matching
- Feat: add NeedsRawMode method to AI vendor interface

- Add NeedsRawMode to Vendor interface

- Implement NeedsRawMode in all AI clients
- Return false for all implementations

- Support model-specific raw mode detection
- Enable future raw mode requirements
- Feat: add support for Anthropic Claude 4 models and update SDK to v1.2.0
CHANGES

- Upgrade `anthropic-sdk-go` dependency to version `v1.2.0`.
- Integrate new Anthropic Claude 4 Opus and Sonnet models.

- Remove deprecated Claude 2.0 and 2.1 models from list.
- Adjust model type casting for `anthropic-sdk-go v1.2.0` compatibility.

- Refresh README: announce Claude 4, update date, fix links.
- Refactor: improve raw mode handling in BuildSession

- Fix system message handling with patterns in raw mode

- Prevent duplicate inputs when using patterns
- Add conditional logic for pattern vs non-pattern scenarios

- Simplify message construction with clearer variable names
- Improve code comments for better readability
- Refactor: improve message handling for raw mode and Anthropic client

- Clarify raw mode message handling in BuildSession

- Fix pattern-based message handling in non-raw mode
- Refactor Anthropic client message normalization

- Add proper handling for empty message arrays
- Implement user/assistant message alternation for Anthropic

- Preserve system messages in Anthropic conversations
- Add safeguards for message sequence validation
- Add authentification for ollama instance
- Refactor content structure in system.md for clarity and readability

- Improved formatting of the introduction and content summary sections for better flow.
- Consolidated repetitive sentences and enhanced the overall coherence of the text.

- Adjusted bullet points and numbering for consistency and easier comprehension.
- Ensured that key concepts are clearly articulated and visually distinct to aid understanding.
- Docs: fix grammar in nuclei template instructions
- Docs: correct Anthropic spelling in notes
- Docs: fix typos in web README
- Docs: fix spelling in PR 1284 update notes
- Docs: fix spelling in pattern management guide
- Add completion files to the build output for Nix
- Chore: update .gitignore and remove obsolete files

- Add `coverage.out` to `.gitignore` for ignoring coverage output.
- Remove `Alma.md` documentation file from the repository.

- Delete `rate_ai_result.txt` stitch script from `stitches` folder.
- Remove `readme.md` for `rate_ai_result` stitch documentation.
- Refactor: introduce `getSortedGroupsItems` for consistent sorting logic

- Add `getSortedGroupsItems` to centralize sorting logic.
- Sort groups and items alphabetically, case-insensitive.

- Replace inline sorting in `Print` with new method.
- Update `GetGroupAndItemByItemNumber` to use sorted data.

- Ensure original `GroupsItems` remains unmodified.
- Feat: add shell completion scripts for Zsh, Bash, and Fish
CHANGES:
- Add shell completion support for three major shells

- Create standardized completion scripts in completions/ directory
- Add --shell-complete-list flag for machine-readable output

- Update Print() methods to support plain output format
- Document installation steps for each shell in README

- Replace old fish completion script with improved version
- Chore: fix "nix flake check" errors
- Refactor: centralize Go version definition in flake.nix
CHANGES

- Define `getGoVersion` function in `flake.nix`.
- Use `getGoVersion` to set Go version consistently.

- Pass `goVersion` explicitly into `nix/shell.nix`.
- Remove redundant Go version definition from `shell.nix`.
- Chore: update Go to 1.24.2 and refresh dependencies
Update Go version across Dockerfile, Nix configurations, and Go modules.
Refresh dependencies and Nix flake inputs.
CHANGES:
- Update Go version to 1.24.2 in Dockerfile.

- Set Go version to 1.24.0 and toolchain to 1.24.2.
- Refresh Go module dependencies and sums (go.mod, go.sum).

- Update Nix flake lock file inputs.
- Configure Nix environment and packages for Go 1.24.

- Update gomod2nix lock file with dependency hashes.
- Use Go 1.24 in Nix development shell environment.
- Chore: unify raw mode message handling and preserve env vars in extension executor

- refactor BuildSession raw mode to prepend system to user content

- ensure raw mode messages always have User role
- keep existing user message when no systemMessage provided

- append systemMessage separately in non-raw mode sessions
- store original cmd.Env before context-based exec command creation

- recreate exec command with context then restore originalEnv
- add comments clarifying raw vs non-raw handling behavior
- Chore: update Anthropic SDK to v0.2.0-beta.3 and migrate to V2 API

- Upgrade Anthropic SDK from alpha.11 to beta.3

- Update API endpoint from v1 to v2
- Replace anthropic.F() with direct assignment

- Replace anthropic.F() with anthropic.Opt() for optional params
- Simplify event delta handling in streaming

- Change client type from pointer to value type
- Update comment with SDK changelog reference
- Chore: sort AI models alphabetically for consistent listing
CHANGES

- Import `sort` and `strings` packages for sorting functionality.
- Sort retrieved AI model names alphabetically, ignoring case.

- Ensure consistent ordering of AI models in lists.
- Chore: alphabetize the order of plugin tools
- Add a completion script for fish
- Feat: enhance StrategyMeta with Prompt field and dynamic naming

- Add `Prompt` field to `StrategyMeta` struct.

- Include `strings` package for filename processing.
- Derive strategy name from filename using `strings.TrimSuffix`.

- Store `Prompt` value from JSON data in `StrategyMeta`
- Feat: add alphabetical sorting to groups and items in Print method**

- Import `sort` and `strings` packages for sorting functionality.
- Create a copy of groups for stable sorting.

- Sort groups alphabetically in a case-insensitive manner.
- Create a copy of items within each group for sorting.

- Sort items alphabetically in a case-insensitive manner.
- Iterate over sorted groups and items for display.
- Feat: add `--listvendors` command to list AI vendors

- Introduce `--listvendors` flag to display all AI vendors.
- Refactor OpenAI-compatible providers into a unified configuration.

- Remove individual vendor packages for streamlined management.
- Add sorting for consistent vendor listing output.

- Update documentation to include new `--listvendors` option.
- Feat: add Cerebras AI plugin to plugin registry

- Introduce Cerebras AI plugin import in plugin registry.
- Register Cerebras client in the NewPluginRegistry function.
- Chore: add final newline to aot json file
- Feat: add Atom-of-Thought (AoT) strategy and prompt definition

- add new aot.json for Atom-of-Thought (AoT) prompting

- define AoT strategy description and detailed prompt instructions
- update strategies.json to include AoT in available strategies list

- ensure AoT strategy appears alongside CoD, CoT, and LTM options
- Fix error in deleting patterns due to non empty directory
- Chore(deps): bump golang.org/x/net
Bumps the go_modules group with 1 update in the / directory: [golang.org/x/net](<https://github.com/golang/net).>

Updates `golang.org/x/net` from 0.36.0 to 0.38.0

- [Commits](<https://github.com/golang/net/compare/v0.36.0...v0.38.0)>
updated-dependencies:
- dependency-name: golang.org/x/net
  dependency-version: 0.38.0
  dependency-type: indirect
  dependency-group: go_modules
Signed-off-by: dependabot[bot] <support@github.com>
- Chore: Update README with a note about Grok
- Feat: add Grok AI provider support`
Integrate the Grok AI provider into the Fabric system for AI model interactions.

- Add Grok AI client to the plugin registry.

- Include Grok AI API key in REST API configuration endpoints.
- #### docs: add contributors section to README with contrib.rocks image

- Add contributors section with visual representation

- Include link to project contributors page
- Add attribution to contrib.rocks tool
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Bump golang version to match go.mod
- Update pattern_descriptions.json
- Finalize WEB UI V2 loose endsfixes
- Fix chat history LLM response sequence in ChatInput.svelte
- Update strategies.json
- Integrate in web ui the strategy flag enhancement first developed in fabric cli
- Updated ed
- Added excalidraw pattern.
- Shorter version of analyze bill.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added bill analyzer.
- Refactor: refactor API key middleware based on code review feedback
- Fix: bad format
- Feat: add simple optional api key management for protect routes in --serve mode
- Feat: add it lang to the chat drop down menu lang in web gui
- Refactor: streamline code_helper CLI interface and require explicit instructions

- Require exactly two arguments: directory and instructions

- Remove dedicated help flag, use flag.Usage instead
- Improve directory validation to check if it's a directory

- Inline pattern parsing, removing separate function
- Simplify error messages for better clarity

- Update usage text to reflect required instructions parameter
- Print usage to stderr instead of stdout
- Docs: improve README link

- Fix broken what-and-why link reference
- Fix: enhance JSON string handling with proper control character escaping

- Convert control chars to proper JSON escape sequences

- Prevent invalid JSON due to literal control chars
- Refactor: rename `fabric_code` tool to `code_helper` for clarity

- Rename tool from `fabric_code` to `code_helper`

- Update all documentation references to the tool
- Update installation instructions in README

- Modify usage examples in documentation
- Update tool's self-description and help text
- Refactor: modify `ParseFileChanges` to return summary and changes separately
CHANGES:
- Return summary text from `ParseFileChanges` separately.

- Update `chatter` to use returned summary text.
- Update tests to match new function signature.
- Refactor: replace FILE_CHANGES marker with constant FileChangesMarker

- Add FileChangesMarker constant for file changes section

- Update parser to use new constant marker
- Improve error messages with dynamic marker reference

- Update tests to use new marker format
- Update system documentation with new marker syntax
- Fix: improve JSON parsing in ParseFileChanges to handle invalid escape sequences

- Add dedicated function to fix invalid JSON escapes

- Handle common \C escape sequence issue
- Implement fallback parsing with comprehensive escape fixes

- Track string context for accurate escape detection
- Preserve valid JSON escape sequences
- Feat: add file management system for AI-driven code changes
CHANGES:
- Replace deprecated io/ioutil with modern alternatives

- Add file change parsing and validation system
- Create secure file application mechanism

- Update chatter to process AI file changes
- Improve create_coding_feature pattern documentation
- Feat: add `fabric_code` tool and `create_coding_feature` pattern
This commit introduces the `fabric_code` tool and the `create_coding_feature` pattern, allowing Fabric to modify existing codebases.

- add `fabric_code` tool to generate JSON representation of code projects

- add `create_coding_feature` pattern to apply AI-generated code changes
- update README with `fabric_code` installation and usage

- walk file system with maximum depth and ignore list
- scan directory and return file/dir JSON data for AI model

- provide usage instructions and examples for `fabric_code`
- add file management API to system prompt for code changes
- Docs: improve README formatting and add clipboard support section

- Remove colons from heading anchors

- Fix broken installation link reference
- Replace code tags with backticks

- Improve code block formatting with indentation
- Clarify package manager alias requirements

- Fix environment variables link
- Simplify custom patterns directory instructions
- Fixed processing message not stopping after pattern output completion
- Add flex windows sizing to web interface
- Fix typo on fallacies instruction.
- Add installation instructions for OS package managers
- Updated find prompt.
- Updated find prompt.
- Updated find prompt.
- Updated find prompt.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added find_female_life_partner.
- Fix: improve error handling in ChangeDefaultModel flow and save environment file

- Add early return on setup error
- Save environment file after successful setup

- Maintain proper error propagation
- Chore: Remove redundant file system.md at top level.
CHANGES:
- Removed `system.md` on the top level of the fabric repo.

- system.md was an RPG session summarization prompt.
- There are two other RPM summary patterns created after this file was added: `create_rpg_summary` and `summarize_rpg_session`
- Fix: set percentEncoded to false
If you use a youtube link like `<https://youtu.be/sHIlFKKaq0A`> percentEndcoding encodes the link to `https%3A%2F%2Fyoutu.be%2FsHIlFKKaq0A`, which throws an error in fabric.
With percentEndcoding false, the script receives the link without encoding and works.
- Moved system file to proper directory.
- Moved system file to proper directory.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added activity extractor.
- Standardize sections for no repeat guidelines
- Added flashcard generator.
- Refactor: remove generic type parameters from NewStorageHandler calls

- Remove explicit type parameters from StorageHandler initialization

- Update contexts handler constructor implementation
- Update patterns handler constructor implementation

- Update sessions handler constructor implementation
- Simplify API by relying on type inference
- Chore: remove redundant yt function definition
- Add newline to end of cod.json
- Fix help message when no strategies found.
- Fix: fix handling of the installed strategies dir
- Chore: remove fallback to local strategies directory if missing
- Clipboard operations now work on Mac and PC
- Change [optional] to [required] in strategies
- Feat: add prompt strategies and improve installation documentation

- Add prompt strategies like Chain of Thought (CoT)

- Implement strategy selection with `--strategy` flag
- Improve README with platform-specific installation instructions

- Fix web interface documentation link
- Refactor git operations with new githelper package

- Add `--liststrategies` command to view available strategies
- Support applying strategies to system prompts

- Fix YouTube configuration check
- Improve error handling in session management
- Bump golang.org/x/net in the go_modules group across 1 directory
Bumps the go_modules group with 1 update in the / directory: [golang.org/x/net](<https://github.com/golang/net).>

Updates `golang.org/x/net` from 0.35.0 to 0.36.0

- [Commits](<https://github.com/golang/net/compare/v0.35.0...v0.36.0)>
updated-dependencies:
- dependency-name: golang.org/x/net
  dependency-type: indirect
  dependency-group: go_modules
Signed-off-by: dependabot[bot] <support@github.com>
- Chore: add .vscode to `.gitignore` and fix typos and markdown linting in `Alma.md`
- Update Web V2 Install Guide with improved instructions
- Fix  Chat history window sizing
- Pattern_explanations.md: fix typo
- Implement  column resize functionnality
- Implement Pattern Tile search functionality
- Fix: update Azure client API version access path in tests
- Chore: remove unnecessary `version` variable from `main.go`
- Feat: Add LiteLLM AI plugin support with local endpoint configuration
- Fix: Fix pipe handling
- Rename input.svelte to Input.svelte for proper component naming convention
- Chore: update version
- Feat: update yt commands and docs to support timestamped transcripts
CHANGES

- Add argument validation to yt for usage errors
- Enable -t flag for transcript with timestamps

- Refactor PowerShell yt function with parameter switch
- Update README to dynamically select transcript option

- Document youtube_summary feature in pattern explanations
- Introduce youtube_summary pattern.
- Remove spurious newline
- Feat: update YouTube regex to support live URLs
- Update Web V2 Install Guide layout 2
- Update azure.go
- Update openai.go
- Update azure_test.go
- Update azure.go
- Update Web V2 Install Guide layout
- Fix: Rework LM Studio plugin
- Update QUOTES section to include speaker names for clarity
- Update Web V2 Install Guide with improved instructions V2
- Update Web V2 Install Guide with improved instructions
- Reorganize documentation with consistent directory naming and updated guides
- Update install guide with Plain Text instructions
- Add Web V2 Installation Guide
- Fix: continue fetching models even if some vendors fail
Remove the cancellation of remaining goroutines when a vendor collection fails.
This ensures that other vendor collections continue even if one fails.
Fixes listing models via `fabric -L` and using non-default models via `fabric -m custom_model`,
when localhost models (e.g. Ollama, LM Studio) are not listening on a given port (basically shut down).
- Merge branch 'main' into pdf-integration-clean
- Reinstate file in original location to resolve PR conflict
- Remove pr-1284-update.md from tracking to resolve PR conflict
- Add required UI image assets for feature implementation
- Complete directory reorganization by moving pr-1284-update.md to new location
- Restore file to original location to resolve path conflict
- Remove pr-1284-update.md from PR scope
- Rename pattern descriptions directory to follow consistent naming convention
- Update README files directory structure and naming convention
- Remove pdf-to-markdown folder from PR
- Flake: fix/update
- Upgrade upload artifacts to v4
- Merge branch 'main' into feat/exolab
- Fix: build problems
- Merge branch 'main' into main
- Merge branch 'main' into main
- Don't trigger on PRs
- Update demo video link in PR-1284 documentation
- Add complete PDF to Markdown documentation F
- Add Svelte implementation files for PDF integration
- Add PDF to Markdown integration documentation
- Update version to v..1 and commit
- Delete version.go
- Delete pkgs/fabric/version.nix
- Add PDF to Markdown conversion functionality to the web svelte caht interface
- Update README.md
- Merge remote-tracking branch 'upstream/main'
- Update to upload-artifact@v4 because upload-artifact@v3 is deprecated
- Merge branch 'danielmiessler:main' into main
- Chore: update Anthropic SDK and add Claude 3.7 Sonnet model support

- Updated anthropic-sdk-go from v0.2.0-alpha.4 to v0.2.0-alpha.11

- Added Claude 3.7 Sonnet models to available model list
- Added ModelClaude3_7SonnetLatest to model options

- Added ModelClaude3_7Sonnet20250219 to model options
- Removed ModelClaude_Instant_1_2 from available models
- Resolving a couple of more medium vulnerabilites
- Updated to fix security issues with ollama.go
- Update version to v..1 and commit
- Added create_loe_document prompt
- Added create_loe_document prompt
- Exclude static PNG files from PR
- Remove PNG files from PR scope
- Enhance pattern handling and chat interface improvements
- Update .gitignore to exclude sensitive and generated files
- Remove sensitive and generated files from tracking
- Remove personal development notes from tracking
- Development checkpoint - Web UI enhancements with complete directory structure
- Setup backup configuration and update dependencies
- Update ENV
- Feat: Add LM Studio compatibility

- Added LM Studio as a new plugin, now it can be used with Fabric.
- Updated the plugin registry with the new plugin name

- Updated the configuration with the required base url
- Updated extract domains
- Updated extract domains
- Added extract_domains
- Create pattern_explanations.md
Spent way too long getting chatgpt to give a one-line summary, based on the contents of each prompt.
- Bump github.com/go-git/go-git/v5
Bumps the go_modules group with 1 update in the / directory: [github.com/go-git/go-git/v5](<https://github.com/go-git/go-git).>

Updates `github.com/go-git/go-git/v5` from 5.12.0 to 5.13.0

- [Release notes](<https://github.com/go-git/go-git/releases)>
- [Commits](<https://github.com/go-git/go-git/compare/v5.12.0...v5.13.0)>
updated-dependencies:
- dependency-name: github.com/go-git/go-git/v5
  dependency-type: direct:production
  dependency-group: go_modules
Signed-off-by: dependabot[bot] <support@github.com>
- Feat: Increase unit test coverage from 0 to 100% in the AI module using Keploy's Agent
- Added h3 TELOS pattern.
- Added challenge handling pattern.
- Added year in review pattern.
- Adding more TELOS patterns.
- Added additional Telos patterns.
- Add the ability to grab YouTube video transcript with timestamps
This commit adds the ability to grab the transcript
of a YouTube video with timestamps. The timestamps
are formatted as HH:MM:SS and are prepended to
each line of the transcript. The feature is enabled
by the new `--transcript-with-timestamps` flag,
so it's similar to the existing `--transcript` flag.
Example future use-case:
Providing summary of a video that includes timestamps
for quick navigation to specific parts of the video.
- Updated panel topic extractor
- Added panel topic extractor
- Added intro sentences pattern
- Updated announcement at the top
- Feat(anthropic): enable custom API base URL support

- Enable and improve custom API base URL configuration
- Add proper handling of v1 endpoint for UUID-containing URLs

- Implement URL formatting logic for consistent endpoint structure
- Clean up commented code and improve configuration flow
- Feat: Added Deepseek AI integration
- Added output filename support for to_pdf
- Feat: implement support for <https://github.com/exo-explore/exo>
- Typos correction
- Doc: Add scrape URL example. Fix Example 4
- Doc: Custom patterns also work with Claude models
- Updated conversion post.
- Adding markdown converter.
- Updated prediction creator.
- Updated predictor pattern.
- Added new prediction generator.
- Create system.md
Create pattern to extract commands from videos and threat reports to obtain commands so pentesters or red teams or Threat hunters can use to either threat hunt or simulate the threat actor.
- Update README.md

1. Windows Command: Because actually curl does not exist natively on Windows
2. Syntax: Because like this; it makes the “click, cut and paste” easier
- Fixed few typos that I could find
- Update README.md: Add PowerShell aliases
- Update summaries and add recently added patterns
- Better metadata
- Update README
- Remove inbox note
- Merge branch 'main' of <https://github.com/johnconnor-sec/fabric>
- Update README
- Updates for BUILD
- Fixing indentation again
Removed backup side-nav and terminal.
updated toast for transcripts
- Added: Only dates are required for Posts now.
- Fix: Chat.svelte indentation
Removed backup files
- Fix: NoteDrawer textarea sizing
- Updated tags page to use Frontmatter instead of PostMetadata
buffer issues
- Deleting old files that were moved or renamed
Folders deleted:
- `types`. The folders contained are now `lib/interfaces` and `lib/api`

- `types/markdown` now in `utils/markdown`
- `components/ui/{side-nav,terminal}` now `components/ui/toc` and
`terminal`
- !NOTEDRAWER IS NOW CENTERED IN VIEWPORT!
- Added metadata lookup to youtube helper
- Font familiy changes.
- !!CHAT IS NOW CENTERED IN VIEWPORT!!
!Chat is now centered in the viewport!
- Edit: styling on Posts page
- 16 word summaries.
- Indented Toc
- Added: blinking cursor to Terminal. Removed / added comments
Removed from ChatMessages
Added to api/context
- Update: References
Moved

- `lib/types/interfaces` to `lib/interfaces`.
- `components/ui/side-nav` to `components/ui/toc`.

- `components/ui/terminal` to `components/terminal`.
- `types/markdown` to `utils/markdown`

- `lib/types/chat` to `lib/api`
- Edit: `type/note` to `note`
type/note was causing problems.
- House Keeping: Added missing png. Removed more unused
- Update README to reflect current @12-30-24
- House Keeping: Fixing Indentation
- Update: Contact page
- Updated copy
- Indentation
- Add: Templates for posts
- Fix: Obsidian Card. Indented app.html
- Update: removed grid from PostContent
- Update: ui/button component
- Rename `chat.ts` to `chat-store.ts`
- Rename `noteStore` to `note-store`
- Updated wrapping instructions.
- Enhanced pattern.
- Enhanced enrich pattern.
- Enhanced enrich pattern.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added enrich_blog_post
- Deleted lib/layouts/files. Renamed lib/store/theme
lib/layouts/files are not longer in use. Renamed lib/store/theme to
adhere to current naming convention
- Update version.nix

Update version.go

Update version.nix

Update version.go
Update version.nix
- Update version.nix
- Update version.go
- Update version to v..1 and commit
- Update version.nix
- Update version to v..1 and commit
- Update version.go
- Update version to v..1 and commit
- Update version.nix
- Update: Post page styling and layout
Indented unused Search.svelte file
- Minor styling improvements
- Remove: Docs in Posts
- Update: NotesDrawer now saves notes to lib/content/inbox
- WIP: Restyling Chat page
- Indented Main page
- Moved NotesDrawer to ModelConfig component
- Deleted: Moved Components from /home to respective dirs
- Add: NotesDrawer to header
- Removed styling from /routes/chat/*.svelte
- Update translate pattern to use curly braces
- Updated story to be shorter bullets.
- Updated extension readme
- Fix tests to handle NewPluginRegistry returning error
- Fix:properly instatiated extensionManager var
fix:added timeout value validation
- WIP: Notes Drawer text color
- WIP: Notes Drawer. Updated default theme to rocket
- Updated POSTS to make main 24-12-08
- Update imports
- Merge branch 'main' into feature/template-extensions
- Added Humanize Pattern
- Merge branch 'main' into main
- Update version to v..1 and commit
- Update version.go
- Fix cross-filesystem file move in to_pdf plugin (issue 1221)
- Update version to v..1 and commit
- Don't quite know how I screwed this up, I wasn't even working there.
- Update version to v..1 and commit
- Merge branch 'main' into main
- Merge branch 'main' into main
- Delete patternstudio.py
- Update README.md
- Refactor pattern management and enhance error handling

- Improved pattern creation, editing, and deletion functionalities.
- Enhanced logging configuration for better debugging and user feedback.

- Updated input validation and sanitization processes to ensure safe pattern processing.
- Streamlined session state initialization for improved performance.

- Added new UI components for better user experience in pattern management and output analysis.
- Add files via upload
Streamlit application for managing and executing patterns, with a focus on pattern creation, execution, and analysis. Below is a breakdown of the key components and functionality of the application:
Key Components and Functionality
    Logging Configuration:
        The application sets up logging with both console and file handlers.
        The console logs are color-coded for better readability, and the file logs are more detailed for debugging purposes.
    Session State Initialization:
        The initialize_session_state() function initializes the session state with default values for various configuration and UI states.
        It also loads saved outputs from persistent storage.
    Pattern Management:
        Pattern Creation: The create_pattern() function allows creating new patterns with either simple or advanced editing options.
        Pattern Deletion: The delete_pattern() function allows deleting existing patterns.
        Pattern Editing: The pattern_editor() function provides an interface for editing existing patterns.
    Pattern Execution:
        Pattern Execution: The execute_patterns() function executes selected patterns and captures their outputs.
        Pattern Chain Execution: The execute_pattern_chain() function executes a sequence of patterns in a chain, passing output from each pattern to the next.
    Output Management:
        Saving Outputs: The save_output_log() function saves pattern execution logs.
        Starring Outputs: The star_output() and unstar_output() functions allow users to star/favorite outputs for quick access.
    Configuration and Model Selection:
        Model and Provider Selection: The load_models_and_providers() function fetches and displays available models and providers for selection.
        Configuration Loading: The load_configuration() function loads environment variables and initializes the configuration.
    UI Components:
        Pattern Creation UI: The pattern_creation_ui() and pattern_creation_wizard() functions provide UI components for creating new patterns.
        Pattern Management UI: The pattern_management_ui() function provides UI components for managing patterns.
        Output Analysis UI: The application includes tabs for displaying all outputs and starred outputs, with options to copy or star outputs.
    Error Handling and Validation:
        Input Validation: The validate_input_content() and sanitize_input_content() functions validate and sanitize input content to ensure it is safe for processing.
        Pattern Validation: The validate_pattern() function validates the structure and content of a pattern.
    Main Function:
        The main() function orchestrates the entire application, setting up the Streamlit page, initializing session state, and handling the main navigation between different views (Run Patterns, Pattern Management, Analysis Dashboard).
Usage and Features
    Pattern Creation: Users can create new patterns using either a simple text editor or an advanced wizard.
    Pattern Execution: Users can select patterns to run, provide input, and execute them either individually or in a chain.
    Output Analysis: Users can view and analyze the outputs of executed patterns, star favorite outputs, and copy outputs to the clipboard.
    Pattern Management: Users can edit, delete, and bulk edit patterns.
    Configuration: Users can select different models and providers for pattern execution.
Error Handling and Logging
    The application includes robust error handling and logging to ensure that any issues are logged and displayed to the user.
    Logging is done both to the console and to a file for debugging purposes.
Future Enhancements
    Enhanced Pattern Validation: More comprehensive validation of pattern content and structure.
    Advanced Analysis: Adding more advanced analysis features, such as sentiment analysis or keyword extraction on pattern outputs.
    Integration with External APIs: Integrating with external APIs for additional functionality, such as sending outputs via email or storing them in a database.
- Add Endpoints to facilitate Ollama based chats
Add Endpoints to facilitate Ollama based chats.
Built to use with Open WebUI
- Fix the typo in the sentence
- Spelling fixes in create_quiz pattern
- Spelling fix in READEME
- Spelling fixes in patterns
- Update version to v..1 and commit
- Merge branch 'main' into main
- Significant update
Added specific steps for research, analysis, and code reviews.
- Significant thematic rewrite.
Ingested the following documents, and then extracted themes and examples of how Socrates interacted with those around him.

- Apology by Plato
- Phaedrus by Plato

- Symposium by Plato
- The Republic by Plato

- The Economist by Xenophon
- The Memorabilia by Xenophon

- The Memorable Thoughts of Socrates by Xenophon
- The Symposium by Xenophon
Many thanks to <a href="<https://www.gutenberg.org/">Project> Gutenberg</a> for the source materials.
- Trying an XML-based Markdown converter pattern.
- Trying an XML-based Markdown converter pattern.
- Updates
- Fix: Issue with the custom message and added example config file.
- Feat: Add YAML configuration support
Add support for persistent configuration via YAML files. Users can now specify
common options in a config file while maintaining the ability to override with
CLI flags. Currently supports core options like model, temperature, and pattern
settings.

- Add --config flag for specifying YAML config path
- Support standard option precedence (CLI > YAML > defaults)

- Add type-safe YAML parsing with reflection
- Add tests for YAML config functionality
- Fix: Mask input token to prevent var substitution in patterns
- Added new instruction trick.
- Add --input-has-vars flag to control variable substitution in input

- Add InputHasVars field to ChatRequest struct
- Only process template variables in user input when flag is set

- Fixes issue with Ansible/Jekyll templates that use {{var}} syntax
This change makes template variable substitution in user input opt-in
via the --input-has-vars flag, preserving literal curly braces by
default.
- Update +page.svelte
- Update version to v..1 and commit
- Update version.nix to reflect upstream/main
- Update version to v..1 and commit
- Update version.go to reflect upstream/main
- Removed arcanum gif
- Update version to v..1 and commit
- Merge branch 'danielmiessler-main'
- Merging
- Analyze_risk pattern
Created a pattern to analyze 3rd party vendor risk.
- Revert "Update version to v..1 and commit"
This reverts commit ec5ed689bb3a90bbc74c68b4d8f3f3da6b47d01b.
- Fix #1169: Add robust handling for paths and symlinks in GetAbsolutePath
- Update version to v..1 and commit
- Added test pattern
- Actually added tutorial
- Added example files and tutorial
- Add cards component
- Update: packages, main page, styles
- Check extension names don't have spoaces
- Build(deps-dev): bump @sveltejs/kit
Bumps the npm_and_yarn group with 1 update in the /web directory: [@sveltejs/kit](<https://github.com/sveltejs/kit/tree/HEAD/packages/kit).>

Updates `@sveltejs/kit` from 2.8.4 to 2.9.0

- [Release notes](<https://github.com/sveltejs/kit/releases)>
- [Changelog](<https://github.com/sveltejs/kit/blob/main/packages/kit/CHANGELOG.md)>

- [Commits](<https://github.com/sveltejs/kit/commits/@sveltejs/kit@2.9.0/packages/kit)>
updated-dependencies:
- dependency-name: "@sveltejs/kit"
  dependency-type: direct:development
  dependency-group: npm_and_yarn
Signed-off-by: dependabot[bot] <support@github.com>
- Update version to v..1 and commit
- Merge remote-tracking branch 'refs/remotes/origin/main'
- Style: Reordered columns. Improved responsive layout
- Style: modified chat/+layout display. Update Header buttons
- Style: updates to ui components (components/ui)
- Added a new pattern create_newsletter_entry
- Added tests for extension manager, registration and execution.
- Moved pattern loader to ModelConfig. Editing styles in chat/. Added page fly transitions. Tidying. Removed - ChatHeader, unused modal from Transcripts, FlyandScaleParams from lib/types/utils.
- Fixed : if there is no stdin, then a nil message was passed to pattern.go resulting in segfault.
now we make user input ' ', before processing.
- Revert "Fix pattern file usage without stdin"
This reverts commit 744ec0824be60ab39337d8cc92ea0552fcc2c31c.
- Fix pattern file usage without stdin
When using pattern files with variables but no stdin input, ensure proper
template processing by initializing an empty message. This allows patterns
like:
  ./fabric -p pattern.txt -v=name:value
to work without requiring stdin input, while maintaining compatibility
with existing stdin usage:
  echo "input" | ./fabric -p pattern.txt -v=name:value
Changes:
- Add empty message initialization in BuildSession when Message is nil

- Remove redundant template processing of message content
- Let pattern processing handle all template resolution
This simplifies the template processing flow while supporting both
stdin and non-stdin use cases.
- Merge branch 'curly-brace-templates' into feature/template-extensions
- Fixed : if there is no stdin, then a nil message was passed to pattern.go resulting in segfault.
now we make user input ' ', before processing.
- Revert "Fix pattern file usage without stdin"
This reverts commit 744ec0824be60ab39337d8cc92ea0552fcc2c31c.
- Fix pattern file usage without stdin
When using pattern files with variables but no stdin input, ensure proper
template processing by initializing an empty message. This allows patterns
like:
  ./fabric -p pattern.txt -v=name:value
to work without requiring stdin input, while maintaining compatibility
with existing stdin usage:
  echo "input" | ./fabric -p pattern.txt -v=name:value
Changes:
- Add empty message initialization in BuildSession when Message is nil

- Remove redundant template processing of message content
- Let pattern processing handle all template resolution
This simplifies the template processing flow while supporting both
stdin and non-stdin use cases.
- Added better messages when adding and listing extensions
Fix issuse with listextension where it would fail if any hash had changed, now says hash failed.
- Extension Registry Refinement

- Successfully implemented path-based registry storage
- Moved to storing paths instead of full configurations

- Implemented proper hash verification for both configs and executables
- Registry format now clean and minimal.
File-Based Output Implementation

- Successfully implemented file-based output handling
- Demonstrated clean interface requiring only path output

- Properly handles cleanup of temporary files
- Verified working with both local and remote operations
- Emplemented stdout template extensions
- Fix: close #1173
- Chore: cleanup style
- Chore: cleanup style
- Fix: use the custom message and then piped one
- Fix: use the custom message and then piped one
- Update README.md
- Update README.md
- Updated readme
- Build(deps-dev): bump @sveltejs/kit
Bumps the npm_and_yarn group with 1 update in the /web directory: [@sveltejs/kit](<https://github.com/sveltejs/kit/tree/HEAD/packages/kit).>

Updates `@sveltejs/kit` from 2.6.1 to 2.8.4

- [Release notes](<https://github.com/sveltejs/kit/releases)>
- [Changelog](<https://github.com/sveltejs/kit/blob/main/packages/kit/CHANGELOG.md)>

- [Commits](<https://github.com/sveltejs/kit/commits/@sveltejs/kit@2.8.4/packages/kit)>
updated-dependencies:
- dependency-name: "@sveltejs/kit"
  dependency-type: direct:development
  dependency-group: npm_and_yarn
Signed-off-by: dependabot[bot] <support@github.com>
- Fix: provide default message content to avoid nil pointer dereference
- Update README.md
- Update Obsidian.md
- Update Obsidian.md
- John 2024-11-26 08:53:48
- Update version to v..1 and commit
- John 2024-11-26 08:44:20
- John 2024-11-26 08:40:21
- Update version to v..1 and commit
- Ci: Integrate code formating
- Update version to v..1 and commit
- Merge branch 'main' into main
- Fax: raw mode was doubling user input, because it's now already embeded in pattern
 streamlined some context staging
- Fix : template.go will handle missing var in stdin imput too.
echo 'Hello {{name}}' | ./fabric -v=noname:World
missing required variable: name
- Fix: process template variables in raw input
Process template variables ({{var}}) consistently in both pattern files
and raw input messages. Previously variables were only processed when
using pattern files.

- Add template variable processing for raw input in BuildSession
- Initialize messageContent explicitly

- Remove errantly committed build artifact (fabric binary in previous commit)
- Added analyze_mistakes.
- Feat(template): implement core plugin system and utility plugins
Add initial set of utility plugins for the template system:
- datetime: Date/time formatting and manipulation

- fetch: HTTP content retrieval and processing
- file: File system operations and content handling

- sys: System information and environment access
- text: String manipulation and formatting operations
Each plugin includes:
- Implementation with comprehensive test coverage

- Markdown documentation of capabilities
- Integration with template package
This builds on the template system to provide practical utility functions
while maintaining a focused scope for the initial plugin release.
- Ci: update patterns zip workflow
- Ci: remove patterns zip workflow
- Fix typo in md_callout
Just a small typo in this pattern. Thanks so much for this splendid tool.
- Feat: migrate to official anthropics Go SDK
- Feat(template): introduce template package for variable substitution

- Add new template package to handle variable substitution with {{variable}} syntax
- Move substitution logic from patterns to centralized template system

- Update patterns.go to use template package for variable processing
- Support special {{input}} handling for pattern content

- Update chatter.go and rest API to pass input parameter
- Enable multiple passes to handle nested variables

- Report errors for missing required variables
This change sets up a foundation for future templating features like front matter
and plugin support while keeping the substitution logic centralized.
- Refactor: unify pattern loading and variable handling

- Stronger separation of concerns between chatter.go and patterns.go
- Consolidate pattern loading logic into GetPattern method

- Support both file and database patterns through single interface
- Maintain API compatibility with Storage interface

- Handle variable substitution in one place
- Keep backward compatibility for REST API through Get method
The changes enable cleaner pattern handling while maintaining
existing interfaces and adding file-based pattern support.
- Add summarize_meeting

Add a new pattern to create a meeting summary from an audio transcript.
The pattern outputs the following sections (where relevant):
- Key Points

- Tasks
- Decisions

- Next Steps
- Update docker image
- Feat: add file-based pattern support
Allow patterns to be loaded directly from files using explicit path prefixes
(~/, ./, /, or \). This enables easier testing and iteration of patterns
without requiring installation into the fabric config structure.

- Supports relative paths (./pattern.txt, ../pattern.txt)
- Supports home directory expansion (~/patterns/test.txt)

- Supports absolute paths
- Maintains backwards compatibility with named patterns

- Requires explicit path markers to distinguish from pattern names
Example usage:
  fabric --pattern ./draft-pattern.txt
  fabric --pattern ~/patterns/my-pattern.txt
  fabric --pattern ../../shared-patterns/test.txt
- Improve logging for missing setup steps
- Add extract_recipe to easily extract the necessary information from cooking-videos
- Fix: fix default gin
- Update version to v..1 and commit
- Add a screenshot of fabric
- Added our first formal stitch.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Upgraded AI result rater.
- Flake: add gomod2nix auto-update
- Ci: zip patterns
- Feat: update dependencies; improve vendors setup/default model
- Feat: add claude-3-5-haiku-latest model
- Merge branch 'main' of github.com:danielmiessler/fabric
- Updated README.
- :sparkles: Added unaliasing to pattern setup
In the process of setting up patterns, we've added a step to unalias any existing alias with the same name. This ensures that our dynamically defined functions won't conflict with any pre-existing aliases.
- Create Selemela07 devcontainer.json
Edit, patch diff
- Add auto save to aliases
-Updated the readme with information about autogenerating aliases that
allow autosaving to obsidian like tools
-Updated the table of contents
- Merge branch 'main' into fix/yt-shorts
- Merge branch 'main' into add-aliases-for-patterns
- Fix: short YouTube url patter
- Add alias generation information
-Updated the readme with information about generating aliases for each
prompt including on for youtube transcripts
-Updated the table of contents
- Added create_diy
- [add] VideoID for YT shorts
- Fix: bufio.Scanner message too long
- Add docker
- Feat: impl. Youtube PlayList support
- Fix: close #1103, Update Readme hpt to install to_pdf
- Fix: close #1106, fix pipe reading
- Feat: YouTube PlayList support
- Create user story pattern
- Fix nix package version auto update workflow
- Automate nix package version update
- Modularize nix flake
- Add model context length setting
- Feat: write tools output also to output file if defined; fix XouTube transcript &#39; character
- Ci: deactivate build triggering at changes of patterns or docu
- Feat: split tools messages from use message
- Feat: impl. multi-model / attachments, images
- Feat: impl. multi-model / attachments, images
- Feat: impl. multi-model / attachments, images
- Feat: impl. multi-model / attachments, images
- Feat: impl. multi-model / attachments, images
- Feat: impl. multi-model / attachments, images
- Feat: impl. multi-model / attachments, images
- Feat: impl. multi-model / attachments, images
- Feat: add md_callout pattern
Add a pattern that can convert text into an appropriate markdown callout
- Add Nix Flake
- Chore: simplify isChatRequest
- Add trailing newline
- Ask uncle Duke
Duke is an expert in software development using the Java programing language, especially with the Spring Framework and Maven.
- Dialog with Socrates
Have a conversation with a modern day philosopher who desires to engage in deep, meaningful conversations.
- Merge branch 'main' of <https://github.com/mattjoyce/fabric> into main
- Added metadata and styleguide
- Merge branch 'danielmiessler:main' into main
- Added structure to prompt
- Added headwinds and tailwinds
- Initial draft of s7 Strategy profiling
- Create system.md
- Feat: add pattern refine_design_document
- Added identify_job_stories
- Feat: add review_design pattern
- Feat: create create_design_document pattern
- Update README.md with pbpaste section
- Added system and user prompts
- Added system and user prompts
- Updated the Alma.md file.
- Fix: setup does not overwrites old values
- Merge remote-tracking branch 'origin/main'
- Feat: plugins arch., new setup procedure
- Feat: plugins arch., new setup procedure
- Update patterns/analyze_answers/system.md - Fixed a bunch of typos
- Updated readme
- Updated extract sponsors.
- Merge branch 'main' into feat/rest-api
- Feat: restructure for better reuse
- Feat: restructure for better reuse
- Feat: restructure for better reuse
- Fix: IsChatRequest rule; Close #1042 is
- Added ctw to Raycast.
- Chore: we don't need tp configure DryRun vendor
- Fix: Close #1040. Configure vendors separately that were not configured yet
- Docs: Close #1035, provide better example for pattern variables
- Updated all dsrp prompts to increase divergenct thinking.
- Fmt
- Support set default output language

#### cli/cli.go

#### core/fabric.go
- Fixed mix up with system
- Initial dsrp prompts
- Fix: Close #1036
- Fix: fix NP if response is empty, close #1026, #1027
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added extract_core_message.
- Feat: work on Rest API
- Feat: work on Rest API
- Feat: work on Rest API
- Feat: work on Rest API
- Feat: work on Rest API
- Feat: work on Rest API
- Corrected spelling and grammatical errors for consistency and clarity
Description:
Changed "agreed within the meeting" to "agreed upon within the meeting" to improve grammatical accuracy.
Added missing periods to ensure consistency across list items.
Corrected the spelling of "highliting" to "highlighting."
Fixed the spelling of "exxactly" to "exactly."
Updated phrasing in "Write NEXT STEPS a 2-3 sentences" to "Write NEXT STEPS as 2-3 sentences" for grammatical correctness.
These changes improve the readability and consistency of the document, ensuring all instructions are clear and error-free.
- Fix: tests
- Fix: windows release
- Fix: windows release
- Feat: Add 'meta' role to store meta info to session, like source of input content.
- Feat: Add 'meta' role to store meta info to session, like source of input content.
- Feat: Add 'meta' role to store meta info to session, like source of input content.
- Feat: Close #1018
- Feat: implement print session and context
- Feat: implement print session and context
- Feat: Setup for specific vendor, e.g. --setup-vendor=OpenAI
- Ci: use the latest tag by date
- Ci: use the latest tag by date
- Ci: use the latest tag by date
- Ci: use the latest tag by date
- Ci: trigger release workflow ony tag_created
- Ci: create repo dispatch
- Ci: test tag creation
- Ci: test tag creation
- Ci: commit version changes only if it changed
- Ci: commit version changes only if it changed
- Ci: use TAG_PAT instead of secrets.GITHUB_TOKEN for tag push
- Merge branch 'main' of github.com:danielmiessler/fabric
- Updated predictions pattern.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added redeeming thing.
- Feat: clean up html readability; add autm. tag creation
- Feat: clean up html readability; add autm. tag creation
- Feat: clean up html readability; add autm. tag creation
- Feat: clean up html readability; add autm. tag creation
- Feat: clean up html readability; add autm. tag creation
- Feat: clean up html readability; add autm. tag creation
- Added extract features.
- Added primary solution.
- Ci: update version
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added epp.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added create_story_explanation.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added create_story_explanation.
- Grab transcript from youtube matching the user's language instead of the first one
- Feat: add version updater bot
- Update system.md in transcribe_minutes
- Support turn any web page into clean view content
- Feat: add version updater bot
- Feat: add version updater bot
- Feat: add version updater bot
- Feat: add version updater bot
- Feat: add version updater bot
- Feat: add version updater bot
- Feat: add version updater bot
- Chore: supress printing YouTube transcripts/comments and grabed Web Sites if a pattern is defined
- Update patterns/solve_with_cot/system.md part three
Noticed the opening closing brackets were incorrect < >
- Update patterns/solve_with_cot/system.md typos part two
Forgot the last typos. Sorry.
- Update patterns/solve_with_cot/system.md typos
- Docs: update YouTube example call
- Fix: skip cli test for now
- Fix: cli test
- Feat: extend installation instruction to get the latest release binaries
- Feat: extend installation instruction to get the latest release binaries
- Feat: extend installation instruction to get the latest release binaries
- Add wipe flag for ctx and session
- Doc: update flags order
- Fix: #986 implement --version flag
- Fix: #986 implement --version flag
- Fix: #997 use setting env value over default values
- Add pattern analyze_military_strategy
Use this pattern to analyze real historic, or fictional battle strategy.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added a generic TELOS file, Alma.md to the repo.
- Merge branch 'main' into feature/seed_parameter
- Updating README with the new flag
- Fix GOROOT path for Apple Silicon Macs in setup instructions
The previous instructions incorrectly set GOROOT to '/opt/homebrew/bin/go', which points to the Go binary rather than the Go root directory. This caused errors when running Go commands on Apple Silicon-based Macs.
I updated the instructions to dynamically determine the correct GOROOT path using Homebrew, ensuring compatibility across different environments. This change resolves the 'go: cannot find GOROOT directory' issue on M1/M2 Macs.
- Feat: remove cli list label and indentation
- Adding flag for pinning seed in openai and compatible APIs
- Feat: #979 add support for configurable base url for Anthropic
- Feat: integrate the output language to the system/user prompt
- Merge branch 'main' into specify_language_return
- Chore: update dependencies
- Chore: #975 check choices available
- Chore: #975 check choices available
- Fix: correct changeDefaultModel flag description
- Feat: improve Jina AI impl.
- Made jina api key optional
- Add mistral vendor
- Updated readme.
- Updated readme.
- Updated readme.
- Updated readme.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Updated readme.
- Feat: use -r, --raw: Use defaults of model (don't send temperature etc.) and use the user role instead of the system role.
- Feat: use -r, --raw: Use defaults of model (don't send temperature etc.) and use the user role instead of the system role.
- Feat: implement -u, --user-instead-of-system: Use the user role instead of the system role for the pattern. It is needed for Open AI o1 models for now.
- Updated Path to install to_pdf [Bug Fix]
- Update system.md
Topic should make sense
- Update the discription of language commend
- Use default language param default:"" to avoid changes in the output.
- Merge branch 'main' into patch-1
- Fix: Spelling Grog to Groq. The setup for the vendor must be done, because of ENV prefix for setting.
- Fix: Spelling Grog to Groc. The setup for the vendor must be done, because of ENV prefix for setting.
- Merge branch 'main' into openrouter
- Feat: add cmd -g to select the language would reply
- Ci: work on artifact names for upload
- Ci: work on artifact names for upload
- Ci: work on artifact names for upload
- Ci: create release before upload
- Ci: set GH_TOKEN for release upload
- Feat: improve Youttube support. Print transcript and/or comments. Don't send them to AI if pattern is not defined.
- Added yt raycast script.
- Added yt raycast script.
- Added yt raycast script.
- Added yt raycast script.
- Fixing binary to .go file
- To_pdf readme
- Latex + pdf
- Updated the rpg summary.
- Updated the rpg summary.
- Updated the rpg summary.
- Updated the rpg summary.
- Updated extract_wisdom Raycast frontmatter.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added a Raycast directory with scripts that integrate Fabric with Raycast.
- Added extract_insights_dm.
- Added extract_insights_dm.
- Added extract_insights_dm.
- Move fabric setup after environment setup in readme
- Feat: Add Jina AI integration for web scraping and question search
- Add new pattern extract_skills
- Fix typo in README.md
- Merge branch 'main' into scrape_url
- Add extract_ctf_writeup as a new pattern
- Added CoT experiment.
- Update README.md
Small cleanup suggestion
- Added comment analysis due to a request in Jason Haddix's AI class.
- Add a pattern for extracting minutes from a transcribed meeting
- Update user.md to match current CLI
Update this pattern to match the current fabric command line options Remove --agents, add -S as an alternative to --setup, and replace -c with -C to align with the current cli interface.
- Update README.md
- Describe CLI changes / Update README.md
- Update: add  env variable info for Apple Silicon
-Updated the readme with env variables for Apple Silicon based mac
as the path for Brew installed apps is different there.
- OpenRouter Vendor
- OpenRouter Vendor
- Feat: add SiliconCloud support
- Feat: add support for pattern variables
- Feat: add support for pattern variables
- Feat: add support for pattern variables
- Merge branch 'danielmiessler:main' into add_dry_run
- Added extract primary problem.
- Added extract primary problem.
- Added extract primary problem.
- Revert unneeded DryRun Vendor registration
- Merge branch 'main' into add_dry_run
- Refactor dry run to DryRun Vendor
- Adding new pattern: `analyze_product_feedback`.
This pattern allows you to summarize, rate, and deduplicate feedback
about products. It's very helpful for anyone working in product
management, engineering, etc.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Updated interviwer analysis.
- Add MIT license file
- Fix dry run
- Fix: usage with deprecated elements
- Merge branch 'main' into create_recursive_outline
- Add dry run
- Updated sales analysis.
- Added analyze sales call.
- Refactor: accept context as parameter of Vendor.Send
In golang, contexts should be propagated downwards in order to be able
to provide features such as cancellation.
This commit refactors the Vendor interface to accept a context as a
first parameter so that it can be propagated downwards.
- Merge branch 'main' into create_recursive_outline
- Added create_story_explanation.
- Added create_story_explanation.
- Added create_story_explanation.
- Added create_story_explanation.
- Added create_story_explanation.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added create_story_explanation.
- Update system.md
Corrected grammatical issues and made the list more readable and consistent.
- Update system.md
Replaced the nested parentheses with equals signs for clarity
- Update system.md
Corrected "it's" to "its" to denote possession instead of a contraction.
- Update system.md
Removed extra "the" for grammatical correctness.
- Update system.md
Corrected "upmost" to "at most" for proper expression.
- Update system.md
Changed "highlight" to "highlights" to match subject-verb agreement.
- Update system.md
Update system.md - removed "a" for a better model understanding
- Fix: shadowing original error
This fixes shadowing the original error so that the original error is
propagated upwards
- Fix: correct os.Exit code from -1 to 1 in main.go
As per the os.Exit documentation, the exit code should be in the rage
[0, 125]
- Feat: native integration of yt tp fabric
- Updated question analysis.
- Updated question extractor.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Updated question extractor.
- Test: core
- Updated interviewer analysis.
- Updated interviewer analysis.
- Updated interviewer analysis.
- Test: core
- Updated interviewer analysis name.
- Added interviwer analysis.
- Updated extract_questions.
- Merge branch 'main' into fix_groq_spelling
- Fix: groq spelling
- Feat: simplify setup logic
- Test: implement test for common package
- Test: implement test for common package
- Updated extract_wisdom_dm.
- Create setup_fabric.bat, a batch script to automate setup and running fabric on windows.
- Fixed strange ollama input involving someone named fred
- Fix broken link in table of contents in README.md (Migrating -> Migration)
- Added back some debug statements
- Removed debug statements
- Chore: Add ScrapeURL flag for CLI to scrape website URL to markdown using Jina AI
- Update README.md
- Adding new pattern: create_recursive_outline.
This pattern is actually based on this incredibly great article: <https://learnhowtolearn.org/how-to-build-extremely-quickly/>
The idea is to use this pattern whenever you want to break an idea or
task down into small components, fully fleshing out your own TODO list
of things to implement to get it working.
This applies to things like writing articles/papers, creating
applications, and much more.
- Ci: split ci and release jobs
- Test: Implement db unit tests
- Feat: Base URL Setup for OpenAPI-Compatible and Proxy Providers. Adapt Grocq and Azure Setup for it.
- Fix: YouTube configured is not mandatory
- Fix: YouTube configured is not mandatory
- Update README.md
Updates to the README for legibility, more detail.
- Update README.md
Fixed some formatting in the README.
- Remove duplicate usage
- Update README.md
Fix typo
- Merge branch 'main' into analyze_cfp_submission_pattern
- Update README.md
Updated install instruction formatting.
- Adding new pattern to help analyze CFP submissions for conference organizers.
- Update README.md
Updated migration instructions.
- Update ollama.go
- Patterns fixes
- Fix spelling error in fabric.go
- Added environment variables to setup.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added migration and upgrade instructions.
- Feat: Improve Gemini vendor - message handling and streaming mode
- Fix: Fix YouTube API key env. name
- Feat: Add YouTube config
- Fix(ci): fix names of artifacts to upload
- Fix(ci): fix names of artifacts to upload
- Fix(ci): fix names of artifacts to upload
- Fix(ci): fix names of artifacts to upload
- Fix(ci): remove dmg from upload artifacts
- Fix(ci): Remove DMG for MacOS
- Created new RPG summarizer.
- Added new RPG summarizer.
- Chore(ci): keep macos binary around in release assets
- Feat: improve Gemini model name handling
- Fix(ci): standardise binary names
- Fix(ci): upload built binaries to GitHub Releases on tag creation
- Feat: add YouTube Configurable Support (without setup activation because the key is external)
- Feat: add last changes from fabric-go; fix some Gemini problems
- Added images folder
- Deleted temp readme.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Updated README.
- Trigger at main
- Updated README.
- Updated README.
- Updated Notes in README.
- Updated Notes in README.
- Massive update to README.md after the Go migration.
- Removed add-context
- Massive update to README.md after the Go migration.
- Massive update to README.md after the Go migration.
- Massive update to README.md after the Go migration.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Massive update to README.md after the Go migration.
- Create Go Build
- Fixed readme
- Added patterns folder
- Updated readme
- Initial
- Improved analyze_email_headers pattern
- Add empty user.md for consistency
- New pattern to analyze email headers for SPF, DKIM and DMARC
- Update README.md, missing word for clarity
inserted "you" into "...to forget the stuff read, watch, or listen to" so that it reads "...to forget the stuff you read, watch, or listen to"
- Docs: correct typos in documentation files
fix various spelling and grammatical errors to improve readability and clarity.
- Added ttrc graph.
- Updated critical vulns patterns.
- Added Alma.md context file example.
- Updated critical data outputer.
- Updated critical data outputer.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added critical graph pattern.
- Updated my EW.
- Updated my EW.
- Updated my EW.
- Updated my EW.
- Updated my EW.
- Added recommend_pipeline_upgrades.
- Added recommend_talkpanel_topics
- Updated legislation analysis.
- Updated summarize_legislation.
- Updated summarize_legislation.
- Updated summarize_legislation.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added summarize_legislation.
- Merge remote-tracking branch 'refs/remotes/origin/main'
- Added controversy extractor.
- Updated Alex Hormozi offer Pattern.
- Added Alex Hormozi offer Pattern.
- Docs: update system.md to refine output and instruction clarity
CHANGES:
- Add intro sentence output requirement

- Emphasize succinct bullet points in CHANGES section
- Remove redundant output instruction

- Expand commit prefix list in instructions
- Add imperative mood and present tense guidelines

- Mention Deis Commit Style Guide adherence
- Update system.md
Typooo
- Updatd create graph.
- Updatd create graph.
- Updatd create graph.
- Updatd create graph.
- Updatd create graph.
- Added create graph.
- Added export data as csv.
- Updated extract_sponsors.
- Updated extract_sponsors.
- Updated extract_sponsors.
- Fixed typo 'bolt' to 'bold' in pattern templates
- Add a mermaid diagram that can render in github markdown readmes
- Feat: add a pattern to draft or respond to an email
- Create system.md for extract_instructions
I created this pattern using the official_pattern_template and piping it to improve_prompt multiple times.
- Create explain_math pattern
- Create system.md for extract_jokes pattern
- First attempt at suggest_pattern pattern with user.md generated with script that uses summarize_prompt pattern
- Add summarize_prompt pattern
- Adding write_hackerone_report
- Fix (grammar)
- Update README.md
- Update fabric.py
Removed "NOTE: This will revert the default model to gpt4-turbo. please run --changeDefaultModel to once again set the default model". This note appears to reflect behavior that is no longer happening.
- Update README.md
Removed "NOTE: This will revert the default model to gpt4-turbo. please run --changeDefaultModel to once again set the default model". This note appears to reflect behavior that is no longer happening.
- Fix: typo in system.md
- Update system.md - Typo, language error
- Update system.md
Added missing language back to the HABITS step.
- Fix some typos in create_logo

- remove a stray quotation mark
- add a missing period for consistency

- add a missing hyphen
- change a hyphen to an em-dash
- Fix some typos in analyze_spiritual_text

- tenants -> tenets
- a handful of punctuation edits
- Added my personal pattern for summarize my course lecutres
- Create system.md for create_tags pattern
- Update extract article wisdom README.md
fixed curl commands - are they obsolete?
- Fixed readme
- Merge remote-tracking branch 'origin/create_coding_project' into create_coding_project
- Added create_coding_project pattern
- Merge branch 'danielmiessler:main' into create_coding_project
- Added the 'create_coding_project' directory. Added README.md. Added system.md (prompt).
- Update utils.py to add Claude 3 Sonnet back in for backwards compatibility
- Update utils.py to support Claude 3.5 Sonnet
- Updated cyber summary.
- Update patterns\clean_text\system.md to improve text cleaning instructions
- Fixed minor typo in the summarise_paper pattern.
app4. roach > approach
- Updated analyze thinker.
- Updated summarize_debate.
- Added summarize_debate.
- Create system.md
- Added create_cyber_summary.
- Added create_cyber_summary.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Moved philocapsulate to capture_thinkers_work.
- Update README.md
- Feat: add create_pattern pattern
- Minor spelling
by to be
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added extract_song_meaning.
- Update README.md
typos fixed
- Fix typo in README.md
- Feat: improve create_stride_threat_model pattern
- Fix typos and formatting in nuclei template rule
- Update system.md
- Create system.md
- Fix naming links
- Fix anchor links in README
- Removes .python-version
- Updated my personal ew.
- Updated official template with INPUT section.
- Changed name of extract_wisdom_large.
- Updated new ew.
- Updated official template.
- Added official_pattern_template.
- Add system.md file for analyzing logs and identifying patterns and anomalies
- Refactored save.py
- Added idea compass pattern
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Added README.md for the create_git_diff_commit pattern with usag for it
- Feat: add create_stride_threat_model pattern
- Added create_git_diff_commit pattern
- Update README.md
- Update system.md
- Create system.md
- Updated new extract_wisdom.
- Updated new extract_wisdom.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added new extract_wisdom pattern with a hand-written example for one-shot training.
- Create_Better_Frames - "Our" to "Are"
Line 67 said: "When our frames our different,..."
Changed it to say: "When our frames are different,..."
- Added SOLUTION section to analyze_patent
Added a SOLUTION section; removed empty lines; added instruction to be verbose and detailed
- Bugfix in helper.py
Bugfix for the error:
Traceback (most recent call last):
  File "/home/xxx/.local/bin/fabric", line 8, in <module>
    sys.exit(cli())
             ^^^^^
  File "/home/xxx/.local/share/pipx/venvs/fabric/lib/python3.12/site-packages/installer/client/cli/fabric.py", line 148, in main
    session.list_sessions()
  File "/home/xxx/.local/share/pipx/venvs/fabric/lib/python3.12/site-packages/installer/client/cli/helper.py", line 67, in list_sessions
    most_recent = self.find_most_recent_file().split("/")[-1]
                  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
AttributeError: 'NoneType' object has no attribute 'split'
- Update system.md
- Create system.md
Add pattern analyze_patent
- Update README.md
Updated readme install instructions.
- Update README.md
Updated Readme.
- Updated extract_wisdom_agents.
- Added an agent version of extract_wisdom using our new Agents in Prompt technique.
- Update README.md
- updated-dependencies:
- dependency-name: requests
  dependency-type: direct:development
  dependency-group: pip
Signed-off-by: dependabot[bot] <support@github.com>
- Added analyze_debate
- Bug fixes for tags

- Prevent generation_date tag format from being modified when SAVE_DATE_FORMAT is
  specified

- Prevent NoneType from ending up in the tags (previous fix did not work)
- Add configurable date format for save helper app

- Update DATE_FORMAT to be configurable using the SAVE_DATE_FORMAT environment variable
- Modify target filename generation to handle cases where SAVE_DATE_FORMAT is left blank

- Default to date format "%Y-%m-%d" if SAVE_DATE_FORMAT is not set
- Feat: introduced pattern for Git diff summaries
CHANGES:
- New system.md file created for summarizing git diffs

- Detailed steps for summarizing Git diffs outlined.
- Emphasis on creating concise, impactful update bullets.

- Introduction of conventional commits for clear change tracking.
- Update README.md Typo
- Updated recommend artists.
- Updated recommend artists.
- Updated recommend artists.
- Updated recommend artists.
- Updated recommend artists.
- Added recommended artists.
- Updated explain_terms.
- Updated explain_terms.
- Updated explain_terms.
- Updated explain_terms.
- Added explain_terms.
- Fix small typo in README
- Fixed gemini support
- Added gemini support
- Updated WPM name.
- Updated WPM name.
- Updated get_wpm.
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added get_wpm.
- Merge branch 'danielmiessler:main' into main
- Added patterns/summarize_paper/README.md
- Updated system prompt; added README
- Ensure env file created in setup when no API keys provided
- Implementation of the analyze answers pattern. Updated the create quiz pattern
- Previous link to client was old/broken.
Replaced it with new:
<https://github.com/danielmiessler/fabric/tree/main/installer/client>
- Adding human readable md
- First draft
- Merge branch 'main' of github.com:danielmiessler/fabric
- Added extracted_business_ideas, by Joseph Thacker.
- Updated rate_ai_response.
- Updated analyze_personality.
- Updated analyze_personality.
- Updated analyze_personality.
- Added analyze_personality.
- Updated rate_ai_response.
- Added rate_ai_response.
- Bump the pip group across 1 directory with 4 updates
Bumps the pip group with 4 updates in the / directory: [gunicorn](<https://github.com/benoitc/gunicorn),> [tqdm](<https://github.com/tqdm/tqdm),> [aiohttp](<https://github.com/aio-libs/aiohttp)> and [idna](<https://github.com/kjd/idna).>

Updates `gunicorn` from 21.2.0 to 22.0.0

- [Release notes](<https://github.com/benoitc/gunicorn/releases)>
- [Commits](<https://github.com/benoitc/gunicorn/compare/21.2.0...22.0.0)>
Updates `tqdm` from 4.66.2 to 4.66.3

- [Release notes](<https://github.com/tqdm/tqdm/releases)>
- [Commits](<https://github.com/tqdm/tqdm/compare/v4.66.2...v4.66.3)>
Updates `aiohttp` from 3.9.3 to 3.9.4

- [Release notes](<https://github.com/aio-libs/aiohttp/releases)>
- [Changelog](<https://github.com/aio-libs/aiohttp/blob/master/CHANGES.rst)>

- [Commits](<https://github.com/aio-libs/aiohttp/compare/v3.9.3...v3.9.4)>
Updates `idna` from 3.6 to 3.7

- [Release notes](<https://github.com/kjd/idna/releases)>
- [Changelog](<https://github.com/kjd/idna/blob/master/HISTORY.rst)>

- [Commits](<https://github.com/kjd/idna/compare/v3.6...v3.7)>
updated-dependencies:
- dependency-name: gunicorn
  dependency-type: direct:development
  dependency-group: pip

- dependency-name: tqdm
  dependency-type: direct:development
  dependency-group: pip

- dependency-name: aiohttp
  dependency-type: indirect
  dependency-group: pip

- dependency-name: idna
  dependency-type: indirect
  dependency-group: pip
Signed-off-by: dependabot[bot] <support@github.com>
- Added pattern of summarize_paper
- Feat: add metadata flag to yt cli
Output includes: id, title, channel, and published_at
- Nuclei template
- Update fabric.py
- Update README.md
- Disentangle PraisonAI references in README
- Updated create_5_sentence_summary.
- Updated create_5_sentence_summary.
- Updated create_5_sentence_summary.
- Updated create_5_sentence_summary.
- Added create_5_sentence_summary.
- Fixed update patterns in gui
- Updated extract_extraordinary_claims
- Updated extract_extraordinary_claims
- Updated extract_extraordinary_claims
- Updated extract_extraordinary_claims
- Updated extract_extraordinary_claims
- Updated extract_extraordinary_claims
- Updated extract_extraordinary_claims
- Added extract_extraordinary_claims
- Fix: The variable 'wisdomFilePath' is already a complete path constructed with 'config_directory'. Joining it again with 'current_directory' could lead to an incorrect path.
The variable 'wisdomFilePath' is already a complete path constructed with 'config_directory'. Joining it again with 'current_directory' could lead to an incorrect path.
- Added create_ai_jobs_analysis.
- Added raw_query.
- Updated extract_wisdom to include a one-sentence takeaway.
- Added extract_questions.
- Updated readme.
- Deleted test.yaml
- Updating Readme Quickstart instructions to include required python version
When I attempted to follow these instructions in a windows environment using WSL, I kept running into issues because my python version was too low (3.8). I then was going through hoops trying to upgrade to version 3.12 as the process seems more complicated on windows OS.
To avoid these headaches, I thought it best to warn potential users ahead of time to ensure their environment is running the latest version of Python or at least python 3.10, which seemed to work for me finally.
- Fixed a typo
- Added ability to list sessions and gives the first line
- Added session log to view your sessions
- Removed analyze-paper.txt
- Added ability to delete some or all sessions
- Fixed some broken things about sessions
- Added sessions
- Add answer interview question pattern
because: As a user, I should be able to answer
interview questions quickly and effectively in realtime
this commit: Adds a pattern for answering interview questions
- Upgraded write_essay.
- Updated presentation analysis pattern.
- Added analyze_presentation.
- Fixed copy and output in local models and claude
- Fixed changing default model to ollama
- Add LMStudio
- Updated guidance.
- Changed threat model to threat scenarios.
- Updated threat modeling.
- Updated ask questions.
- Updated ask questions.
- Updated ask questions.
- Changed name of secure_by_default.
- Changed name of secure_by_default.
- Added secure by design pattern.
- Update system.md
- Adding a pattern for malware analysis summary
This is an experimental pattern for creating a summary of a malware report.
- Fixed --listmodels in the situation where there is no claude key
- Merge remote-tracking branch 'upstream/main'
- Fixed the situation where there is no openai api key...again
- Merge remote-tracking branch 'upstream/main'
- Upgraded agents with PraisonAI. the --agents flag will now CREATE an AI agent for you and then perform a task. Enjoy
- Update README.md for #324
Closes #324 , showing how to connect to another server
- Merge branch 'main' of github.com:ksylvan/fabric
- Merge remote-tracking branch 'upstream/main'
- Merge branch 'danielmiessler:main' into main
- Merge branch 'danielmiessler:main' into main
- Added fine tuning to the gui
- Added options to set temperature, top_p, frequency_penelty, presence_penalty
- Merge branch 'main' into main
- Fixed the gui
- Upgraded investigation pattern.
- Added create_investigation_visualization.
- Fixed something
- Fixed gui again
- Fixed the gui
- Updated readme
- Added --gui option to fabric. this will open the gui
- Made gui look a little nicer
- Added functionality to gui to create your own patterns
- Bump follow-redirects from 1.15.5 to 1.15.6 in /installer/client/gui
Bumps [follow-redirects](<https://github.com/follow-redirects/follow-redirects)> from 1.15.5 to 1.15.6.

- [Release notes](<https://github.com/follow-redirects/follow-redirects/releases)>
- [Commits](<https://github.com/follow-redirects/follow-redirects/compare/v1.15.5...v1.15.6)>
updated-dependencies:
- dependency-name: follow-redirects
  dependency-type: indirect
Signed-off-by: dependabot[bot] <support@github.com>
- Fixed even more stuff...trust me you'll love it
- Fixed stuff in the UI that I did badly...more to come im sure
- Updated gui to include local models and claud...more to comee
- Get OLLAMA models to work in Windows, including both native and WSL environments.
- Updated fabric markmap visualizer.
- Added fabric markmap visualizer.
- Added show_fabric_options
- Added extract_wisdom_nometa
- Added rate_ai_result.
- Bump langchain-core from 0.1.31 to 0.1.35
Bumps [langchain-core](<https://github.com/langchain-ai/langchain)> from 0.1.31 to 0.1.35.

- [Release notes](<https://github.com/langchain-ai/langchain/releases)>
- [Commits](<https://github.com/langchain-ai/langchain/commits)>
updated-dependencies:
- dependency-name: langchain-core
  dependency-type: indirect
Signed-off-by: dependabot[bot] <support@github.com>
- Modified:   installer/client/cli/yt.py
- Updated create_upgrade_pack.
- Added create_upgrade_pack.
- Added create_upgrade_pack.
- Added extract_insights.
- Added get_youtube_rss.
- Updated pinker prose.
- Fixed Latin-1 decode problems
Fixes Latin-1 decode problems
- Generate CSV instead of a Markdown table
- Add pattern: to_flashcards
- Updated pinker prose.
- Updated pinker prose.
- Added find_logical_fallacies
- Improved analyze_prose_pinker
- Improved analyze_prose_pinker
- UpdatedPinker prose analysis.
- Added Pinker prose analysis.
- Updated find hidden message.
- Updated find hidden message.
- Added an INSIGHTS section to extract_wisdom.
- Added an INSIGHTS section to extract_wisdom.
- Updated length on extract_ideas
- Updated length on extract_ideas
- Unfucking things
- Unfucking things
- Last min fixes
- Last min fixes
- Last min changes
- Last min changes
- Last min changes
- Added create_branch function to git-cont.py.
- Two new prompts, create and improve report finding for pentest repor finding generation.
- Two new prompts, create and improve report finding for pentest repor finding generation.
- Added micro essay pattern.
- Updated essay pattern.
- Updated essay pattern.
- Updated extract_ideas.
- Updated reading plan pattern.
- Updated reading plan pattern.
- Added create_reading_plan.
- Updated readme.
- Fixed default models once again
- Updated extract_book_recommendations.
- Updated extract_book_recommendations.
- Updated extract_book_recommendations.
- Updated extract_book_ideas.
- Updated extract_book_ideas.
- Added extract_book_recommendations.
- Updated extract_book_ideas.
- Updated extract_book_ideas.
- Updated extract_book_ideas.
- Added extract_book_ideas.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Improved analyze_paper.
- Made extract_wisdom more concise.
- Updated the ai pattern to give slightly longer output.
- Updated the ai pattern to give slightly longer output.
- Updated the ai pattern to give slightly longer output.
- Updated create_show_intro.
- Added create_show_intro.
- Added create_art_prompt.
- Removed helper_file directory.
- Update README.md
- Update README.md
- Added a setup.sh just as an onramp to the new pipx installer.
- Updated create_security_update.
- Updated create_security_update.
- Updated create_security_update.
- Updated create_security_update.
- Updated create_security_update.
- Added create_security_update.
- Fix missing --transcript flag for yt command in example
- Add a great example on extracting wisdom from any Youtube video
- Yt comments includes reply threads. Readme updated.
- Added comment retrieval option to yt.py
- Updated create_better_frame
- Updated create_better_frame
- Updated create_better_frame
- Added create_better_frame
- Updated create_academic_paper.
- Removed user.md
- Added create_academic_paper.
- Removed user.md
- Updated explain_project.
- Added explain_project.
- Fixed situation where there was no default model listed
- Again fixed defaultmodel
- Fixed defaultmodel
- Update system.md
- Modified yt to also accept urls via stdin
- Updated extract_patterns.
- Updated extract_patterns.
- Updated create summary and create micro summary.
- Added create summary and create micro summary.
- Updated readme.
- Fixed yt...again
- Fixed yt
- Fixed version. also removed a redundant reference to pyperlclip in poetry env
- Add system.md file for writing pull requests
- Fixed yt, ts and save
- Fixed something with models i broke yesterday
- New network_threat_landscape pattern to analyse port statistics created by FlyingPhish/Nmap-Analysis or provide two bullet point lists with port and service info.
- Fix grammar in improve_academic_writing
- Fix grammar and add improve_academic_writing
- Fix grammar
- Change improve_writing prompt into md format
- Now fixed something that I myself broke
- Fixed even more stuff that was broken by pull requests
- Fixed lots of things that pull requests broke
- Added copy to local models and claude
- Fixed readme
- Added persistant custom patterns. Anything you add to the .config/fabric/patterns folder will persist
- Fixed yt and ts
- Add support for Claude 3 Haiku
- Merge branch 'main' into output-saver
- Added extrac_main_idea pattern.
- Updated readme.
- Updated poetry installer for yt.
- Updated the readme with better install instructions.
- Updated the readme with better install instructions.
- Tweaked installer.
- Tweaked installer.
- Fixed something
- Updated the readme with better install instructions.
- Updated Matthew Berman video.
- Added Matthew Berman video.
- Added Matthew Berman video.
- Added dependancy
- Deleted setup.sh. its no longer needed because of pipx
- Updated readme
- Initial
- Added pyperclip dependancy to poetry
- Update system.md
minor fix
- Fix bug in sendMessage by moving code
- Updated provide_guidance pattern.
- Updated provide_guidance pattern.
- Updated provide_guidance pattern.
- Updated provide_guidance pattern.
- Updated provide_guidance pattern.
- Updated provide_guidance pattern.
- Added provide_guidance pattern.
- Updated algorithm recommender.
- Added extract_algorithm_update to patterns.
- Add code to use openai_base_url and use OpenAI's model lister function
Signed-off-by: zestysoft <ian@zestysoft.com>
- Merge branch 'main' of github.com:danielmiessler/fabric
fixed youtube
- Added youtube api key to --setup
- Added support for remote ollama instances with --remoteOllamaServer
- Fixed an error with -ChangeDefaultModel with local models
- Fixed a setup.sh error that would occur on macos
- Fixed local models
- Minor typo in extract_predictions
- Use safer method to get data from exception
Signed-off-by: zestysoft <ian@zestysoft.com>
- Add code to use openai_base_url and use OpenAI's model lister function
Signed-off-by: zestysoft <ian@zestysoft.com>
- Assorted typo and spelling corrections.
- Added analyze_tech_impact pattern for assessing the impact of technology
- Add supported Claude models
- Changed how aliases are stored. Intead of the .zshrc etc. aliases now have their own file located at ~/.config/fabric/fabric-bootstrap.inc which is created during setup.sh. Please run ./setup.sh and these changes will be made automatically. your .zshrc/.bashrc will also be automatically updated
- Fixed local
- Fixed something with llama models
- Updated readme.
- Updated extract_predictions.
- Added extract_predictions.
- Updated find_hidden_message pattern.
- Updated find_hidden_message pattern.
- Updated find_hidden_message pattern.
- Updated find_hidden_message pattern.
- Updated find_hidden_message pattern.
- Updated pattern.
- Updated pattern.
- Fixed typo
- Updated rpg_summarizer.
- Updated rpg_summarizer.
- Updated rpg_summarizer.
- Updated extract_patterns.
- Updated extract_patterns.
- Updated extract_patterns.
- Updated extract_patterns.
- Added extract_patterns.
- Fixed even more stuff
- Fixed more
- Fixed stuff
- Fixed stuff
- Fixed something
- Changed some documentation
- Changed some documentation
- Fixed some stuff
- Fixed setup
- Updated the README.md notes.
- Updated the README.md notes.
- Updated the README.md notes.
- Updated the README.md notes.
- Updated the README.md notes.
- Fixed some stuff
- Got rid of --claude and --local. everything is in --model
- Fixed something
- Added an error message
- Added stuff to setup
- Changed readme
- Added persistance
- Trying a thing
- Added --changeDefaultModel to persistantly change default model
- Fixed something
- Changed more documentation
- Added some stuff
- Added support for claude. choose --claude. make sure to run --setup again to enter your claude api key
- Added yet another error message
- Updated readme to add refresh note.
- Updated readme
- Changed an error message
- Fixed the stuff that was broken
- Fixed readme
- Added support for local models
- Just a little faster now
- Made it faster
- Add language option to yt.py
- Updated readme
- Updated agents
- Updated summarize_git_changes.
- Updated summarize_git_changes.
- Updated extract_ideas.
- Updated extract_ideas.
- Added extract_ideas.
- Move usage block
- Use exception messages for a better chance at debugging
- Update design pattern and docs
- Helper utility for saving a Markdown file
'save' can be used to save a Markdown file, with optional frontmatter
and additional tags. By default, if set, `FABRIC_FRONTMATTER_TAGS` will
be placed into the file as it is written. These tags and front matter
are suppressed from STDOUT, which can be piped into other patterns or
programs with no ill effects. This strives to be a version of `tee` that
is enhanced for personal knowledge systems that use frontmatter.
- Updated summarize_git_changes.
- Updated summarize_git_changes.
- Added summarize_git_changes.
- Merge remote-tracking branch 'origin/main'
fixed agents
- Fixed agents
- Added new pattern called create_command
Add New "create_command" Pattern
- Merge remote-tracking branch 'origin/main'
fixed things
- Fixed yt
- Added yt and ts to poetry and to config in setup.sh
- Updated readme
- Updated readme
- Added transcription
- Added vm dependencies to poetry
- Merge branch 'agents'
added agents functionality
- Fix the cat.
- Updated client documentation.
- Removed default context file.
- Added MarkMap visualization.
- Added MarkMap visualization.
- Added MarkMap visualization.
- Updated visualizations.
- Updated visualizations.
- Updated visualizations.
- Updated README.md.
- Added helpers README.md.
- Removed visualize.
- Renamed vm to yt, for youtube.
- Removed temp plot.
- Updated create_keynote.
- Added slide creator.
- Added slide creator.
- Added slide creator.
- Added slide creator.
- Added slide creator.
- Added slide creator.
- Added create_threat_model.
- Added create_threat_model.
- Added create_threat_model.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated create_visualization.
- Updated pattern.
- Updated pattern.
- Updated pattern.
- Updated pattern.
- Updated pattern.
- Updated pattern.
- Updated pattern.
- Updated pattern.
- Added create_visualization.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Updated hidden messages Pattern.
- Added find_hidden_message Pattern.
- Added agents
- Updated typo in README
on-behalf-of: pensivesecurity <luke@pensivesecurity.io>
- Added create_command pattern
on-behalf-of: pensivesecurity <luke@pensivesecurity.io>
- Updated rpg_summarizer.
- Updated rpg_summarizer.
- Fixed more stuff
- Fixed something
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated README.md with video info.
- Updated intro video.
- Updated readme for server instructions.
- Updated create_video_chapters.
- Updated create_video_chapters.
- Updated create_video_chapters.
- Updated create_video_chapters.
- Added create_video_chapters.
- Updated label_and_rate.
- Updated label_and_rate.
- Updated label_and_rate.
- Alphabetically sort patterns list
Ensures that when the users lists the available patterns, they are presented in alphabetical order. Helps find the desired pattern faster.
- Bump the pip group across 1 directories with 1 update
Bumps the pip group with 1 update in the /. directory: [cryptography](<https://github.com/pyca/cryptography).>

Updates `cryptography` from 42.0.2 to 42.0.4

- [Changelog](<https://github.com/pyca/cryptography/blob/main/CHANGELOG.rst)>
- [Commits](<https://github.com/pyca/cryptography/compare/42.0.2...42.0.4)>
updated-dependencies:
- dependency-name: cryptography
  dependency-type: indirect
Signed-off-by: dependabot[bot] <support@github.com>
- Cleanup.
- Updated output instructions.
- Created a STATISTICS version of analyze_threat_report.
- Created a STATISTICS version of analyze_threat_report.
- Created a TRENDS version of analyze_threat_report.
- Improved summary to analyze_threat_report.
- Added summary to analyze_threat_report.
- Added a threat report analysis pattern.
- Changed readme
- Changed another message
- Added a statement
- Added aliases for individual patterns. Also fixed pattern download process
- Updates

- README.md - added instructions to make sure the setup.sh script was executable as this was not explicitly stated
- setup.sh - updated sed to use `sed -i` to be compatible with Linux, MacOSX and other OS versions and added a check in the local directory taht setup.sh executes in for a pyproject.toml file because the script was looking for the .toml file in the user's home directory and throwing an error
- Updated write_essay to be more conversational and less grandiose and pompous.
- Minor README edit for verb form consistency
Change `Create` to `Creating`.
- Remove stray .DS_Store file
- Now context.md is in .config
- Updated readme
- Fixed context
- Added context to cli. edit context.md and add -C to add context to your queries
- Updated analyze_paper with more detail and legibility.
- Updated analyze_paper with more detail and legibility.
- Added a specific version of extract_wisdom just for articles.
- Adds templates on the repo
- Update README.md
- Fixes readme link on CLI instructions
- Update README.md
- Removes initialization of API keys from server
- Update README.md
- Update README.md with new Quickstart note.
- Missing a reference on readme
- New line so that aliases are appended on new lines
- Bash_profile added to aliases
- Typo
- Removes echo
- Updates install instructions after naked debian test
- Added Andre Guerra to credits.
- Added Andre Guerra to primary contributors.
- Single script install instructions added on readme
- Incorporates poetry install and dep setup on a single script
- Renamed package to installer while keeping poetry project as fabric
- Added Dani Goland to credits for enhancing the server.
- Removed helpers2.
- No need to enter installer folder
- Updates readme
- Renames fabric folder into fabric_installer
- Merge branch 'main' into single.poetry
- Updated vm.
- Updated to add better docs.
- Updated vm.
- Added /helpers/vm which downloads youtube transcripts and accurate durations of videos using your own YouTube API key.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated rate_value.
- Updated main readme.
- Updated rate_value.
- Updated rate_value with credits in the README.md.
- Updated rate_value.
- New value rating pattern.
- Updated analyze_prose_json.
- Updated analyze_prose_json.
- Add 2 patterns
Added 1 pattern (improve_writing) which improve the writing and returns it in the native language of the input
Added 1 pattern (analyze_incident) which analyses incident articles and produces a neat and simple output (Taken from the YT Video that Daniel was in by David B
- Added Joseph Thacker to the credits.
- Added Jason Haddix to the credits.
- Conflicts solved
- Merging upstream main and solving conflict
- Fix typo
- Update system.md
Fixed a typo
- Small fix for a problem where the GUI was loading every pattern twice
- Update utils.py to get_cli_input Line 192
Changed sys.stdin.readline().strip() to sys.stdin.read().strip() to allow multiple line input.
- Readded the client folder to the structure
- Add compare and contrast system and user patterns
- Add user story and acceptance criteria for agility story patterns
- Shortened summary sentences.
- Add client dir.
- Added installation note.
- Added video embed.
- Merge branch 'danielmiessler:main' into main
- Update README.md #86 Clarify the Instructions in the README
- Redirects redundant instruction on CLI to main readme
- Single poetry project; script to create aliases in bash and zsh; updates readme
- Update utils.py
fixed something else
- Update fabric.py
fixed an error
- Fixed something
- Fixed something
- Update README.md
Clarified a line in the readme
- Update utils.py with a class to transcribe YouTube Videos
Added Class Transcribe with method youtube which accepts a video id as a parameter.  Returns the transcript.
- Updated gui to include adding API key and updating patterns
- Adds meta back
- Pushes readme updates
- Poetry for server app; readme instructions added
- Fix steps to install
- Added copy to clipboard
- Added drag and drop and updated UI
- Chore: typo
- Feat: mapping path and pattern in the dictionary, allowing to scale the pattern "The Mill" server can use easily
- Updated the readme with credit to Jonathan Dunn for the GUI client.
- Update fabric.py to work with standalone.get_cli_input()
For compatibility with Visual Studio Community Edition
- Update utils.py - This is a utility function standalone.get_cli_input()
This function adds compatibility to Visual Studio Community edition.
- Fixed the README
- Changed name of web_frontend to gui as this is a standalone electron app
- Added a web frontend-electron app
- Fabric as a CLI; poetry for dep management with latest versions; gitignore re-added
- Broke analyze_prose into Markdown and JSON versions.
- Updates to analyze_prose.
- AP.
- AP.
- AP.
- AP.
- AP.
- AP.
- AP.
- AP.
- AP.
- AP.
- Prose analysis upgrade.
- Analyze_prose
- Upgrades to analyze_prose.
- Upgrades to analyze_prose.
- Upgrades to analyze_prose.
- Updated analyze_prose.
- Unscrewed the repo.
- AP.
- Fixed dupes.
- Upgrades to analyze_prose.
- Made analyze_prose more stringent.
- Added analyze_prose.
- Fix some typos and updates gitignore
- Adds more comments to the code.
[Snorkell.ai] Please review the generated documentation
- EW.
- EW.
- EW.
- EW.
- EW.
- EW tweak.
- Tweak to extwis again.
- Slight tweak.
- Slight tweak to extract_wisdom.
- Updated extract_wisdom with tiny tweaks..
- Updated extract_wisdom with insight and surprise.
- Reverted label_and_rate.
- Updated label and rate.
- Removed helpers for now.
- Added a pattern and a new helper directory.
- Test commit
- Fixed some typos.
- Removed an extra print statement, thanks to @rez0.
- Added missing word to prompt instruction
- Update system.md
Fixed Markdown mismatches and added H1 headers to Steps and Output to make more consistent with other patterns
- Correct the configuration to define alias in the shell
- Added a one-sentence summary to label_and_rate.
- Update model list to api dynamic response based on user key
- Updated PR pattern.
- Added summarize PRs.
- Create FUNDING.yml
- Merge branch 'main' into cli-model-version
- Add model and list-model to args
- Added extract_references.
- Updated recommendations output.
- Added extract_recommendations.
- Updated model to GPT-4 preview, which is always the latest.
- Changed model to gpt-4-turbo-preview.
- Added new pattern for labeling and rating content.
- Update README.md - fix contributor URLs
Headshot for Daniel went to Scott, and vice versa.
- Formatting.
- Formatting.
- Demo comment again.
- Added demo comment.
- Uploaded DEMO-Write Essay Movie
- Uploaded the write_essay demo.
- Update README.md
- Formatting.
- Use jwt auth
- Nav.
- Nav.
- Nav.
- Nav.
- Nav again.
- Nav.
- Remove: virtual env source folder and .zshrc
- Nav again.
- Add: add virual env ignore in client's
- Fixed nav, maybe?
- Fixed nav.
- Fixed nav, maybe.
- Updated the readme.
- Updated TOC.
- Updated TOC.
- Updated TOC.
- Added TOC
- Update README.md
- Update README.md
Fix typo in one of the code examples
- Update .gitignore and server configurations
- Updates to the Quickstart documentation.
- Updated the Quickstart instructions.
- Update README.md
- Update README.md
- Update README.md
- Changed model back to 4-preview.
- Updated RPG sumamrizer.
- Updating some help info on the client.
- Updated patterns update message.
- Fixed Python path in client.
- Added --setup
- Fixed how the API key was loaded to be sent to OpenAI.
- Added contributor avatars.
- Update README.md
- Update README.md
- Update README.md
- Added contributors.
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Major README updates.
- Major updates to the main README.
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Update README.md
- Updating the readme.
- Minor change
- Fixed a typo
- Fixed --update
- Updated the standalone app
- Removes .venv from gitignore as it appears to be used on purpose
- Adds python's gitignore for common files and removes py chaches
- Readme update.
- Client docs update.
- Some icon stuff.
- Hitcount badge.
- Some icon stuff.
- Some icon stuff.
- More README restructuring.
- Readme restructuring.
- More documentation updates.
- Move documentation updates.
- Move documentation updates.
- Some formatting cleanup.
- Updated docs links.
- Added Quickstart and a patterns screenshot.
- Moved standalone client examples under /client.
- Updated main docs with Quickstart.
- Added an early version of the main client. Still buggy. Let us know what issues you have.
- Updated extract_wisdom to be a compound word.
- Some additional documentation and instructions.
- Added some additional documentation, but the real stuff will come with a usage GIF of the universal client.
- Added some additional documentation, but the real stuff will come with a usage GIF of the universal client.
- Added more examples.
- Added link to Bombal video
- Added an input section to the bottom of extract_wisdom.
- Fixed confusing language in prompt.
- Updated prompt to be more liberal.
- Updated prompt to be more liberal.
- Added explain_docs.
- Fix typo in Markdown bold
- Fixes apikey bug in client code
- Adds missing requirement flask
- Update README.md
Miss Click?
- Update README.md
ouptut -> output
- Added note to top of project about early release.
- Fixed extwis path in readme.
- Added Jonathan Dunn to the credits for his speactular work on the soon-to-be released standalone client.
- Added more analysis to analyze_claims.
- Added analyze_claims, in hopes that it might help us in 2024.
- Updated tagline.
- Updated extract_sponsors.
- Updated extract_sponsors.
- Updated extract_sponsors.
- Updated extract_sponsors.
- Updated extract_sponsors to include potential sponsors.
- Updated extract_sponsors.
- Updated extract_sponsors.
- Added sponsor extraction.
- Tweaked output.
- Added more detail to the output.
- Added more detail to the output.
- Added final intro sentence.
- Updated the output for create_show_intro.
- Updated the output for create_show_intro.
- Changed the name of the pattern and added guest info as well.
- Updated output instructions.
- Updated output format.
- Added topics creator.
- Added philocapsulate.
- Updated summarize_micro.
- Updated summarize_micro.
- Fixed API server keyfile definition.
- Added claims to analize spiritual text.
- Updated analyze_spiritual_text to put the examples in the points.
- Updated analyze_spiritual_text to not include similarties, and to include examples.
- Updated analyze_spiritual_text to not include similarties, and to include examples.
- Added clean_text.
- Removed silly directory.
- Added analyze_spiritual_text.
- Added summarize_newsletter.
- Added summarize_rpg_session.
- Added create_npc.
- Added improve_prompt.
- Added extract_poc
- Added aphorism.
- Added extract_videoid.
- Added Semgrep rule creator.
- Updated pattern again.
- Added a readme and updated the pattern.
- Added write_essay.
- Updated references section on ExtWis.
- Updated output formatting.
- Added rate content (used to be labelandrate.
- Added explain code pattern.
- Added requirements.txt file and updated how the OpenAI library is used.
- Added create_toc
- Updated create logo.
- Updated create logo.
- Added create logo.
- Added analyze paper.
- Added openai key file.
- Added server infra.
- Added MIT license to the project.
- Added system and user to base.
- Added extract wisdom back.
- Added check agreement.
- Adding agreement check.
- Made them all H1s.
- Made them all H1s.
- Updated to new format.
- Optimizing.
- Optimizing summarize.
- Back to original format.
- Back to recommended format.
- Removed dmiessler from a pattern.
- Added patterns directory.
- GPT-4 recommended format.
- Fixing micro formatting.
- Updated system prompt for summarize_micro.
- Updated dash to underscore.
- Added Summarize-micro
- Tweaking prompt.
- Testing something.
- Tweaking prompt.
- Tweaking prompt.
- Tweaking prompt.
- Tweaking prompt.
- Tweaking prompt.
- Made the pattern use bullets.
- Added files to root of Summarize.
- Added Summarize.
- Updated extwis instructions and headers.
- Updated extwis system prompt formatting.
- Added Caleb to credits.
- Revised main description.
- Removed main logo.
- Removed main logo.
- Removed main logo.
- Added gif.
- Added logo file.
- Copied readme from subdirectory.
- Added additional comment to output.
- Removed alert from first extwis pattern.
- Cleanup.
- Cleanup.
- Fixed quotes.
- Fixed code.
- Added examples.
- Added auto call sexiness.
- Updated usage.
- Added usage.
- Formatting.
- Added Joel to credits.
- Tweaks.
- Updated alert.
- Added second caution emoji.
- Updated alert message.
- Added credits.
- More verbiage addition.
- Updated example.
- Updated main description again.
- Updated Fabric main description.
- Updated functionality description for the Fabric readme.
- Transparent image.
- Updated logo image.
- First commit of main Fabric readme.
- Updated use cases.
- Italics.
- Italics, not bold.
- Capitalized second sentence.
- Capitalized Fabric.
- Removed extra space on top.
- Changed to h4.
- Added h4 on the main statement.
- Removed colons.
- Formatting for use cases.
- Cleaned up header.
- More cleanup.
- Updated meta.
- Added example.
- Adding more content.
- Adding more content.
- Adding more content.
- Adding more content.
- Adding more content.
- Adding more content.
- Adding sections and content to the readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adding formatting stuff to readme.
- Adjusted extwis logo location.
- Adjusted extwis logo size.
- Added a CONTENT: piece to user.md.
- Added new file structure.
- Removed extwis file.
- Changed JSON format.
- Added first pattern, extwis.
- Initial commit

## v1.4.319 (2025-09-30)

### PR [#1783](https://github.com/danielmiessler/Fabric/pull/1783) by [ksylvan](https://github.com/ksylvan): Update anthropic-sdk-go and add claude-sonnet-4-5

- Feat: update `anthropic-sdk-go` to v1.13.0 and add new model
- Upgrade `anthropic-sdk-go` to version 1.13.0
- Add `ModelClaudeSonnet4_5` to supported models list

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

- Remove garble installation step from release workflow
- Add comment for GoReleaser config file reference link
- The original idea of adding garble was to make it pass
  virus scanning during version upgrades for Winget, and
  this was a failed experiment.

## v1.4.315 (2025-09-20)

### Direct commits

- Chore: update CI workflow and simplify goreleaser build configuration

- Add changelog database to git tracking

- Remove unnecessary goreleaser comments
- Add version metadata to default build

- Rename windows build from garbled to standard
- Remove garble obfuscation from windows build

- Standardize ldflags across all build targets
- Inject version info during compilation

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
