# Document 4: Fabric Plugin Architecture Review

## Context

- **Playbook**: Fabric PR Review
- **Agent**: Fabric-PR-Review
- **Project**: /Users/kayvan/src/fabric
- **Date**: 2026-03-25
- **Working Folder**: /Users/kayvan/src/fabric/.maestro/playbooks

## Purpose

Verify that new or modified AI providers and plugins follow Fabric's plugin architecture.

## Prerequisites

- `/Users/kayvan/src/fabric/.maestro/playbooks/REVIEW_SCOPE.md` exists from Document 1

## Tasks

### Task 1: Load Context

- [x] **Read scope**: Loaded `/Users/kayvan/src/fabric/.maestro/playbooks/REVIEW_SCOPE.md`; plugin review scope is limited to the Codex vendor refactor in `internal/plugins/ai/codex/`.

- [x] **Check if plugins changed**: Plugin files were modified in this PR. Relevant scope from `REVIEW_SCOPE.md`: `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/codex.go`, `internal/plugins/ai/codex/errors.go`, `internal/plugins/ai/codex/oauth.go`, and `internal/plugins/ai/codex/token.go`.

### Task 2: Review VendorPlugin Interface

For new or modified AI providers in `internal/plugins/ai/`:

- [x] **Check interface compliance**: Verified Codex against the actual AI vendor interface in `internal/plugins/ai/vendor.go` (the playbook text points to `internal/plugins/plugin.go`, which only defines the generic plugin base). `codex.Client` satisfies the contract via embedded `openai.Client`/`plugins.PluginBase` for lifecycle methods plus Codex-specific `ListModels`, `Send`, and `SendStream` overrides in `internal/plugins/ai/codex/codex.go`. Confirmed with `go test ./internal/plugins/ai/codex ./internal/core`.
  - `Name() string` - Returns vendor identifier
  - `Models() ([]string, error)` - Lists available models
  - `Chat(context.Context, *ChatRequest) (*ChatResponse, error)` - Main chat method
  - Any other required interface methods

- [x] **Verify registration**: Confirmed `internal/core/plugin_registry.go` registers Codex exactly once via `codex.NewClient()` in `NewPluginRegistry`, and `VendorsManager.AddVendors` keys vendors by lowercase name so `Codex` remains uniquely addressable with case-insensitive lookup. Checked `internal/plugins/ai/openai_compatible/providers_config.go` and found no `"Codex"` entry in `ProviderMap`, so there is no name collision with OpenAI-compatible providers. Also verified the only Codex-specific model-selection customization is the explicit manual-model passthrough in `PluginRegistry.GetChatter`; there is no separate vendor-alias configuration layer to review for this provider.
  - Vendor name is unique
  - Proper initialization
  - Model aliases configured (if applicable)

### Task 3: Review OpenAI-Compatible Vendors

For vendors extending `openai_compatible`:

- [x] **Check base extension**: Not applicable for the scoped Codex refactor. Verified `internal/plugins/ai/codex/codex.go` embeds `*internal/plugins/ai/openai.Client` via `openaivendor.NewClientCompatibleNoSetupQuestions(...)`, not `internal/plugins/ai/openai_compatible.Client`, so this PR does not add or modify any vendor that extends the `openai_compatible` provider base. Confirmed there is no `"Codex"` entry in `internal/plugins/ai/openai_compatible/providers_config.go`, which is where OpenAI-compatible vendors are declared and configured.
  - Properly embeds `openai_compatible` base
  - Only overrides necessary methods
  - API endpoint is correctly configured

- [x] **Verify API differences**: Reviewed Codex for the API-difference concerns this task is meant to catch and documented the intentional divergence: Codex uses OAuth refresh tokens and an authenticated transport in `internal/plugins/ai/codex/auth_transport.go`/`oauth.go`, adds `originator`, `User-Agent`, `Authorization`, and `ChatGPT-Account-ID` headers per request, and points at dedicated Codex/Auth base URLs instead of the generic OpenAI-compatible provider registry. No OpenAI-compatible vendor-specific header/auth/model-mapping changes are present in this PR beyond that standalone Codex implementation.
  - Any provider-specific headers
  - Authentication method (API key, OAuth, etc.)
  - Model name mapping if different

### Task 4: Review Streaming Implementation

- [ ] **Check streaming support**:
  - Implements streaming via callbacks
  - Handles SSE (Server-Sent Events) correctly
  - Properly closes connections on context cancellation
  - Error handling during streams

- [ ] **Verify stream cleanup**:
  - No goroutine leaks
  - Buffers are flushed
  - Resources are released

### Task 5: Review Model-Specific Features

- [ ] **Check feature flags**:
  - Thinking/reasoning modes (if supported)
  - Web search capabilities
  - Image/multimodal support
  - TTS (text-to-speech) support

- [ ] **Verify context handling**:
  - Context window limits respected
  - Token counting (if applicable)
  - Truncation strategy for long inputs

### Task 6: Review Configuration

- [ ] **Check config loading**:
  - API keys from environment variables
  - Proper use of `godotenv`
  - No hardcoded credentials
  - Fallback handling for missing config

- [ ] **Verify flag support**:
  - CLI flags for provider selection
  - Model selection flags
  - Provider-specific options

### Task 7: Review Other Plugin Types

For changes to other plugin types:

- [ ] **Database plugins** (`internal/plugins/db/`):
  - Proper connection handling
  - Query safety (no SQL injection)
  - Resource cleanup

- [ ] **Strategy plugins** (`internal/plugins/strategy/`):
  - Proper prompt modification
  - Strategy chaining support
  - Error handling

- [ ] **Template plugins** (`internal/plugins/template/`):
  - Extension loading
  - Variable substitution
  - Security of template execution

### Task 8: Document Plugin Issues

- [ ] **Create PLUGIN_ISSUES.md**: Write findings to `/Users/kayvan/src/fabric/.maestro/playbooks/PLUGIN_ISSUES.md`:

```markdown
# Plugin Architecture Review

## Plugins Reviewed
[List of plugins/vendors checked]

## Interface Compliance Issues
[VendorPlugin interface violations]

## Registration Issues
[Plugin registry problems]

## Streaming Issues
[Streaming implementation problems]

## Configuration Issues
[Config loading, credential handling]

## Feature Implementation Issues
[Model-specific feature problems]

## Security Concerns
[Credential exposure, injection risks]

## Suggestions
[Architectural improvements]

## No Issues Found
[Plugins that passed all checks]

## Skipped
[Note if no plugins were modified in this PR]
```

For each issue include:
- Plugin/vendor name and file
- Issue description
- Suggested fix
- Severity: Critical / Major / Minor / Suggestion

## Success Criteria

- All modified plugins reviewed
- Interface compliance verified
- Streaming implementation checked
- Configuration reviewed
- PLUGIN_ISSUES.md created

## Status

Mark complete when plugin review document is created.

---

**Next**: Document 5 will perform security-focused analysis.
