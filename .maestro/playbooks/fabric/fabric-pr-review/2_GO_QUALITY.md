# Document 2: Go Code Quality Review

## Context

- **Playbook**: Fabric PR Review
- **Agent**: Fabric-PR-Review
- **Project**: /Users/kayvan/src/fabric
- **Date**: 2026-03-25
- **Working Folder**: /Users/kayvan/src/fabric/.maestro/playbooks

## Purpose

Perform a Go-specific code review focusing on Fabric's coding conventions, Go idioms, and best practices.

## Prerequisites

- `/Users/kayvan/src/fabric/.maestro/playbooks/REVIEW_SCOPE.md` exists from Document 1

## Tasks

### Task 1: Load Context

- [x] **Read scope**: Loaded `/Users/kayvan/src/fabric/.maestro/playbooks/REVIEW_SCOPE.md` and identified the Go review target as the Codex plugin refactor in `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/codex.go`, `internal/plugins/ai/codex/errors.go`, `internal/plugins/ai/codex/oauth.go`, and `internal/plugins/ai/codex/token.go`. No PR images were attached.

### Task 2: Check Go Idioms

- [x] **Error handling**: Verified `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/errors.go`, `internal/plugins/ai/codex/oauth.go`, and `internal/plugins/ai/codex/token.go` against Fabric's error-handling expectations.
  - Errors are returned, not panicked (no `panic()` in library code)
  - Use `pkg/errors` for wrapping: `errors.Wrap(err, "context")`
  - Error messages are lowercase, no punctuation
  - Errors don't expose sensitive information
  - Review notes:
    - No `panic()` calls were found in the reviewed Codex package, so the refactor still returns errors through normal call chains.
    - Wrapping is inconsistent with the stated Fabric convention. The refactor uses `fmt.Errorf(... %w ...)` in multiple places instead of `pkg/errors.Wrap`, including `internal/plugins/ai/codex/auth_transport.go:90`, `internal/plugins/ai/codex/auth_transport.go:104`, `internal/plugins/ai/codex/oauth.go:59`, `internal/plugins/ai/codex/oauth.go:219`, `internal/plugins/ai/codex/oauth.go:233`, `internal/plugins/ai/codex/oauth.go:245`, `internal/plugins/ai/codex/oauth.go:281`, and `internal/plugins/ai/codex/codex.go:197`.
    - A few synthesized error strings violate the lowercase style requirement, for example `internal/plugins/ai/codex/errors.go:24` (`Codex request failed with status %d`), `internal/plugins/ai/codex/oauth.go:219` (`Codex token exchange failed: %w`), `internal/plugins/ai/codex/oauth.go:245` (`invalid Codex auth base URL: %w`), `internal/plugins/ai/codex/token.go:45` (`JWT did not include an exp claim`), and `internal/plugins/ai/codex/token.go:61` (`invalid JWT format`).
    - Non-401 HTTP failures currently surface upstream response text directly via `extractErrorMessage` in `internal/plugins/ai/codex/errors.go:15-46`, and those errors are echoed back to the local OAuth browser callback in `internal/plugins/ai/codex/oauth.go:174` and `internal/plugins/ai/codex/oauth.go:180`. That can leak provider-supplied detail to the user, so it should be reviewed as a potential sensitive-information exposure path rather than assumed safe.

- [x] **Context usage**: Checked `context.Context` patterns across `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/codex.go`, and `internal/plugins/ai/codex/oauth.go`.
  - Context is the first parameter where applicable
  - Context is propagated through call chains
  - Cancellation is handled for long operations
  - Timeouts are set for external calls
  - Review notes:
    - The request-driven paths do pass `context.Context` first where the package API allows it. `Send(ctx, ...)` accepts the caller context and passes it into `Responses.NewStreaming`, while the auth transport reuses `req.Context()` for initial auth and the forced refresh retry in `internal/plugins/ai/codex/codex.go:212-248` and `internal/plugins/ai/codex/auth_transport.go:139-169`.
    - The OAuth flow also propagates cancellation correctly once setup has started. `Setup()` creates a bounded 5 minute context, `runOAuthFlow()` selects on `ctx.Done()`, and the callback token exchange uses `r.Context()` for the `/oauth/token` POST in `internal/plugins/ai/codex/codex.go:103-124` and `internal/plugins/ai/codex/oauth.go:53-123,197-239`.
    - External calls do have explicit time bounds in the reviewed code: `authHTTPClient` is configured with a 30 second client timeout in `NewClient()` and `configure()`, `ListModels()` creates its own 30 second request context, and the interactive OAuth setup uses the 5 minute `oauthTimeout` in `internal/plugins/ai/codex/codex.go:90,103-124,128,168-169`.
    - The main gap is at the Fabric plugin boundary: `ai.Vendor.SendStream` does not accept a context parameter, so `Codex.SendStream()` has to create the OpenAI stream with `context.Background()` in `internal/plugins/ai/vendor.go:12-17` and `internal/plugins/ai/codex/codex.go:252-289`. That means caller cancellation is not propagated for streaming requests, matching the same known limitation already documented in `internal/plugins/ai/azureaigateway/azureaigateway.go:201-224`.
    - Two additional non-interactive paths also detach from caller cancellation by using `context.Background()`: `configure()` refreshes tokens with a background context before requests are sent, and `ListModels()` cannot accept a caller context because the shared vendor interface exposes `ListModels() ([]string, error)` rather than a context-aware variant in `internal/plugins/ai/codex/codex.go:127-142,160-208`. Those calls are still bounded by timeout, but they cannot stop early if the upstream command is cancelled.

- [x] **Interface compliance**: Verified interfaces across `internal/plugins/ai/codex/*.go` and the shared AI/plugin boundaries in `internal/plugins/ai/vendor.go` and `internal/plugins/plugin.go`.
  - Functions accept interfaces, return concrete types
  - Interfaces are defined where they're used
  - No empty interfaces (`interface{}`) without good reason
  - Review notes:
    - The Codex refactor does not introduce any new package-local interfaces. It continues to satisfy the existing `ai.Vendor` and `plugins.Plugin` contracts through the concrete `*codex.Client` type, which matches Fabric's current plugin architecture rather than adding one-off abstractions.
    - Return types in the reviewed package stay concrete and specific (`*Client`, `oauthTokens`, `modelsResponse`, `tokenClaims`, standard library structs, and plain `error`), so the refactor did not widen APIs with interface-typed returns.
    - Parameter typing is also mostly concrete. The only deliberate interface-style extension points are standard-library seams that are appropriate for the call site: the embedded `http.RoundTripper` in `authTransport` and the injected `openBrowserFn func(string) error` used to make the OAuth flow testable without shelling out.
    - No `interface{}` usages were added in the Codex package. The only `any` usage in scope is the local `map[string]any` decoding in `internal/plugins/ai/codex/errors.go:84-135`, which is justified because upstream error payloads are schema-variable JSON blobs and need dynamic inspection.
    - One design limitation remains at the shared boundary rather than in Codex itself: `internal/plugins/ai/vendor.go:12-17` defines `SendStream` and `ListModels` without `context.Context`, so Codex cannot accept interface-shaped cancellation hooks for those operations even though its concrete request paths are otherwise context-aware.

### Task 3: Review Code Organization

- [x] **Package structure**: Checked the Codex package layout and dependency boundaries across `internal/plugins/ai/codex/*.go`, plus the registration path in `internal/core/plugin_registry.go`.
  - `internal/` packages are truly internal
  - No circular dependencies
  - Clear package boundaries
  - Appropriate file sizes
  - Review notes:
    - `internal/plugins/ai/codex` remains internal-only in practice: the reviewed package is registered from `internal/core/plugin_registry.go:21,86` and no non-internal import sites were found, so the refactor did not create a new externally consumable boundary.
    - `go list -deps ./internal/plugins/ai/codex` completed successfully, and the import graph stays one-directional through shared internals (`internal/chat`, `internal/domain`, `internal/i18n`, `internal/plugins`, and `internal/plugins/ai/openai`) plus the upstream OpenAI SDK. No circular dependency indicators surfaced in the package graph.
    - File responsibilities are mostly separated cleanly by concern: `auth_transport.go` owns token injection and retry behavior, `oauth.go` owns the browser callback flow, `errors.go` centralizes HTTP/API error normalization, and `token.go` keeps JWT parsing/version helpers isolated.
    - The main structural weakness is size and mixed responsibility in `internal/plugins/ai/codex/codex.go`. At 348 lines it currently combines package constants, client construction, setup/configuration, model listing, request/stream execution, and request-shaping helpers, while `oauth.go` is another 304 lines. This is still manageable, but the package is trending toward a "god file" entrypoint; splitting setup/model operations from request execution would make boundaries clearer before more behavior lands there.

- [x] **Naming conventions**: Verified identifier casing and descriptiveness across `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/codex.go`, `internal/plugins/ai/codex/errors.go`, `internal/plugins/ai/codex/oauth.go`, and `internal/plugins/ai/codex/token.go`.
  - CamelCase for exported identifiers
  - camelCase for unexported identifiers
  - Meaningful, descriptive names
  - Acronyms are all caps (HTTP, API, ID)
  - Review notes:
    - Exported identifiers introduced or exercised by the Codex package follow Go casing conventions: `Client`, `NewClient`, `Setup`, `ListModels`, `Send`, and `SendStream` use CamelCase, while unexported helpers like `ensureAccessToken`, `refreshAccessToken`, `buildAuthorizeURL`, `tokenNeedsRefresh`, and `codexInstructionsAndMessages` stay in camelCase.
    - Most names are descriptive at the level of responsibility they carry. Transport/authentication helpers (`authTransport`, `cloneRequest`, `refreshErrorFromResponse`), OAuth helpers (`runOAuthFlow`, `exchangeCodeForTokens`, `publishOAuthResult`), and token parsing helpers (`parseTokenClaims`, `extractAccountIDFromJWT`) are all specific enough that the call sites read clearly without extra comments.
    - Acronym handling is mostly correct in Codex-local names. Examples like `AuthBaseURL`, `AccountID`, `IDToken`, `oauthClientID`, `oauthCallbackPath`, `generatePKCECodes`, and `extractExpiryFromJWT` use standard Go initialism casing consistently.
    - The one naming mismatch against the stated acronym rule is the inherited `ApiBaseURL` and `ApiClient` fields used throughout `internal/plugins/ai/codex/codex.go`. Those names should be `APIBaseURL` and `APIClient` by strict Go style, but they originate from the embedded shared OpenAI vendor client in `internal/plugins/ai/openai/openai.go` rather than being introduced by this Codex refactor. Treat this as an existing cross-provider naming debt, not a Codex-specific regression.

- [x] **Documentation**: Checked package and exported API documentation across `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/codex.go`, `internal/plugins/ai/codex/errors.go`, `internal/plugins/ai/codex/oauth.go`, and `internal/plugins/ai/codex/token.go`.
  - Exported functions have doc comments
  - Package-level documentation exists
  - Complex logic is explained
  - No stale comments
  - Review notes:
    - Package-level documentation exists via the `Package codex ...` comment in `internal/plugins/ai/codex/codex.go:1-2`, so the package is discoverable in Go doc output without adding a separate `doc.go`.
    - Exported functions and methods reviewed in scope are documented: `NewClient`, `Setup`, `ListModels`, `Send`, and `SendStream` in `internal/plugins/ai/codex/codex.go:72,94,160,211,251`, plus the exported `RoundTrip` method in `internal/plugins/ai/codex/auth_transport.go:134`.
    - The exported `Client` type itself is missing a doc comment at `internal/plugins/ai/codex/codex.go:48`, which is the primary documentation gap in the package surface and should be fixed to satisfy standard Go doc expectations for exported identifiers.
    - The most non-obvious flow has at least some targeted inline explanation. In particular, `cloneRequest()` documents why `GetBody` must exist for the one-time authenticated retry in `internal/plugins/ai/codex/auth_transport.go:171-178`.
    - No clearly stale or misleading comments were found in the reviewed files. The existing comments still match current behavior, though complex OAuth callback and error-normalization paths rely more on readable naming than on explanatory comments.

### Task 4: Review Concurrency

- [x] **Goroutine safety**: Checked shared-state and goroutine behavior in `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/codex.go`, and `internal/plugins/ai/codex/oauth.go`.
  - Race conditions on shared state
  - Proper channel usage (closing, direction)
  - Context-aware goroutines
  - No goroutine leaks
  - Review notes:
    - Shared token state is serialized through `Client.tokenMu` in `internal/plugins/ai/codex/auth_transport.go:26-67`, so concurrent request paths cannot race while reading or refreshing `AccessToken`, `RefreshToken`, or `AccountID`. A package-level race run (`go test -race ./internal/plugins/ai/codex`) completed cleanly.
    - The only Codex-local goroutine is the OAuth callback server launched in `internal/plugins/ai/codex/oauth.go:91-99`. Its lifecycle is bounded by the outer `select` in `runOAuthFlow()` and both the success and timeout branches call `server.Shutdown(...)` and wait on `serveDone`, so the listener goroutine is not left running after the flow returns.
    - Channel usage in the OAuth flow is defensive rather than racy. Both `results` and `serveDone` are buffered with capacity 1 in `internal/plugins/ai/codex/oauth.go:84,91`, and `publishOAuthResult()` uses a non-blocking send in `internal/plugins/ai/codex/oauth.go:190-195`, which prevents duplicate callbacks or late error paths from hanging the handler after the first terminal result is delivered.
    - I did not find Codex-specific evidence of a goroutine leak in the reviewed paths. The main concurrency limitation remains outside this checkbox: streaming requests still inherit Fabric's broader `SendStream` interface design, which is reviewed separately under the streaming task below.

- [x] **Streaming**: Reviewed streaming-response behavior in `internal/plugins/ai/codex/codex.go`, the shared `internal/plugins/ai/vendor.go` interface, and the existing Codex streaming coverage in `internal/plugins/ai/codex/codex_test.go`.
  - Channels are properly buffered
  - Errors are communicated correctly
  - Cleanup happens on cancellation
  - Review notes:
    - `SendStream()` always closes the caller-provided channel with `defer close(channel)` and also closes the upstream SDK stream with `defer stream.Close()`, so normal completion and early-return error paths do not leak the local channel or the SDK stream handle in `internal/plugins/ai/codex/codex.go:252-289`.
    - Buffering is delegated to the caller rather than enforced by Codex. The implementation performs direct blocking sends on the provided channel (`channel <- ...` at `internal/plugins/ai/codex/codex.go:273-276` and `internal/plugins/ai/codex/codex.go:283-286`), so the path is safe with the buffered test channel in `internal/plugins/ai/codex/codex_test.go:471` but can stall indefinitely if an unbuffered or unread channel is passed. That matches Fabric's current vendor contract, but it means the provider itself does not guarantee non-blocking streaming behavior.
    - Errors are only returned from `SendStream()` via `c.mapRequestError(stream.Err())` at `internal/plugins/ai/codex/codex.go:289`; no `domain.StreamTypeError` update is emitted on the channel even though the unified stream payload supports that event type in `internal/domain/stream.go:7-16`. Consumers therefore have to watch both the channel and the function return value, and they cannot receive an in-band terminal error after partially streamed content.
    - Cancellation cleanup remains limited by the shared interface rather than the Codex loop itself. `ai.Vendor.SendStream` does not accept `context.Context` in `internal/plugins/ai/vendor.go:12-17`, so Codex starts the upstream SSE request with `context.Background()` in `internal/plugins/ai/codex/codex.go:267`. If the caller disconnects or abandons the read side, the request cannot be cancelled proactively; cleanup happens only when the remote stream ends or the transport errors, which is the same architectural gap already called out in `internal/plugins/ai/azureaigateway/azureaigateway.go:201-206`.
    - Existing coverage exercises the success-path SSE parsing and confirms that the channel closes after the stream is drained in `internal/plugins/ai/codex/codex_test.go:425-493`, but there is no Codex-specific test for mid-stream transport errors or abandoned consumers.

### Task 5: Review API Changes

- [x] **Breaking changes**: Reviewed API compatibility for the Codex refactor against `origin/main`.
  - Are changes backward compatible?
  - Are deprecation notices added?
  - Is the CHANGELOG updated?
  - Review notes:
    - I compared the PR diff against `origin/main` and the change surface is limited to an internal package split plus the changelog fragment in `cmd/generate_changelog/incoming/2063.txt`; there are no edits to shared interfaces in `internal/plugins/ai/vendor.go`, `internal/plugins/plugin.go`, or the embedded OpenAI client contract.
    - The exported Codex package surface remains unchanged. `Client`, `NewClient`, `Setup`, `ListModels`, `Send`, and `SendStream` in `internal/plugins/ai/codex/codex.go` keep the same names, parameters, and return types as the pre-refactor implementation, so this refactor is backward compatible for Fabric callers.
    - No deprecation notices were needed or added because the PR does not remove or replace any public API entrypoints; it only moves internal helpers from `codex.go` into focused files.
    - The repo-level `CHANGELOG.md` is not directly modified in this branch, but the required changelog input was added via `cmd/generate_changelog/incoming/2063.txt`, which is the project’s normal path for changelog generation. That is sufficient for this refactor-sized internal change.

- [x] **Function signatures**: Verified the Codex refactor against Fabric's signature conventions and the existing `ai.Vendor` contract.
  - Context is first parameter
  - Options pattern for many parameters
  - Error is last return value
  - Review notes:
    - The exported Codex surface is still signature-compatible with `ai.Vendor` in `internal/plugins/ai/vendor.go:12-17`. `Send(ctx, msgs, opts)` keeps `context.Context` as the first parameter in `internal/plugins/ai/codex/codex.go:212`, while `ListModels() ([]string, error)` and `SendStream(msgs, opts, channel) error` remain unchanged at `internal/plugins/ai/codex/codex.go:161` and `internal/plugins/ai/codex/codex.go:252`.
    - Package-local helpers that take context also follow the same rule: `runOAuthFlow(ctx, ...)` in `internal/plugins/ai/codex/oauth.go:53`, `exchangeCodeForTokens(ctx, ...)` in `internal/plugins/ai/codex/oauth.go:197`, `ensureAccessToken(ctx, ...)` in `internal/plugins/ai/codex/auth_transport.go:26`, and `refreshAccessToken(ctx)` in `internal/plugins/ai/codex/auth_transport.go:69` all place `context.Context` first.
    - Error returns consistently stay in the final position across the reviewed signatures, including multi-return helpers like `ensureAccessToken(...) (string, string, error)`, `refreshAccessToken(...) (oauthTokens, error)`, and `Send(...) (string, error)` in `internal/plugins/ai/codex/auth_transport.go:26-69` and `internal/plugins/ai/codex/codex.go:212`.
    - I did not find a strong options-pattern violation in the refactor itself. The widest package-local signatures are still small and cohesive, and the user-facing request methods already consolidate tunables into `*domain.ChatOptions` rather than adding more positional parameters.
    - The remaining gap is inherited from the shared interface rather than introduced here: `ai.Vendor` defines `SendStream([]*chat.ChatCompletionMessage, *domain.ChatOptions, chan domain.StreamUpdate) error` and `ListModels() ([]string, error)` without `context.Context` in `internal/plugins/ai/vendor.go:14-16`. Codex therefore cannot expose context-first signatures for those operations even though its internal helpers otherwise follow the convention.

### Task 6: Check Fabric-Specific Patterns

- [x] **Plugin patterns**: Reviewed the Codex refactor against Fabric's AI vendor conventions in `internal/plugins/ai/codex/codex.go`, `internal/plugins/ai/vendor.go`, `internal/plugins/plugin.go`, and the shared OpenAI vendor base. No PR images were attached for this task.
  - Implement `VendorPlugin` interface
  - Handle streaming via callbacks
  - Support model listing
  - Handle context cancellation
  - Review notes:
    - Codex still fits Fabric's vendor architecture by embedding the shared OpenAI vendor client created via `openaivendor.NewClientCompatibleNoSetupQuestions(...)` in `internal/plugins/ai/codex/codex.go:73-91`. That inherited client already carries the standardized `plugins.PluginBase` from `internal/plugins/plugin.go:24-47`, so Codex continues to satisfy the combined `plugins.Plugin` and `ai.Vendor` contract in `internal/plugins/ai/vendor.go:12-17` through method promotion plus its Codex-specific overrides.
    - The provider still supports the required model-listing path. `ListModels()` remains exported on the Codex client and performs an authenticated `/models` request with response filtering in `internal/plugins/ai/codex/codex.go:160-208`, so the refactor did not drop Fabric's model discovery behavior.
    - Streaming remains aligned with Fabric's current vendor pattern, which is channel-based rather than function-callback based. `SendStream(...)` emits `domain.StreamUpdate` values onto the caller-provided channel and closes it on exit in `internal/plugins/ai/codex/codex.go:252-289`, matching the shared `ai.Vendor.SendStream(..., chan domain.StreamUpdate)` signature in `internal/plugins/ai/vendor.go:14-16`.
    - Context handling is only partially compliant because the shared interface still splits behavior: normal requests use caller cancellation through `Send(ctx, ...)` in `internal/plugins/ai/codex/codex.go:211-248`, and interactive setup uses a bounded timeout in `Setup()` in `internal/plugins/ai/codex/codex.go:94-125`, but streaming cannot receive caller cancellation because `SendStream` has no `context.Context` parameter in `internal/plugins/ai/vendor.go:14-16`. Codex therefore starts the upstream stream with `context.Background()` in `internal/plugins/ai/codex/codex.go:266`, which preserves existing Fabric behavior but remains a design gap for abandoned or disconnected stream consumers.
    - I did not find a Codex-specific regression in plugin registration or discovery. The provider is still registered alongside the other AI vendors in `internal/core/plugin_registry.go:71-90`, so the refactor stays within the established plugin-loading path rather than introducing a parallel mechanism.

- [x] **Configuration patterns**: Reviewed the Codex refactor against Fabric's existing configuration paths in `internal/plugins/ai/codex/codex.go`, `internal/plugins/plugin.go`, `internal/core/plugin_registry.go`, `internal/plugins/db/fsdb/db.go`, and `cmd/fabric/main.go`. No PR images were attached for this task.
  - Environment variables via `godotenv`
  - Flags via `go-flags`
  - YAML config support
  - Review notes:
    - The refactor stays on Fabric's standard plugin configuration path rather than introducing a new one. `NewClient()` still declares Codex configuration through `AddSetting(...)` and `AddSetupQuestionWithEnvName(...)` for the access token, refresh token, account ID, API base URL, and auth base URL in `internal/plugins/ai/codex/codex.go:72-91`.
    - Those settings still flow into Fabric's shared `.env` persistence layer unchanged. `PluginBase.SetupFillEnvFileContent()` serializes plugin settings, `PluginRegistry.SaveEnvFile()` writes all vendor settings into the repo-managed env content, and the filesystem DB reloads that file with `godotenv.Load(...)` in `internal/plugins/plugin.go:132`, `internal/core/plugin_registry.go:139-157`, and `internal/plugins/db/fsdb/db.go:82-87`. The Codex split therefore remains compliant with the project's environment-variable configuration pattern.
    - No new CLI flags were added for Codex. Fabric's command-line parsing still lives in `cmd/fabric/main.go:7-17` via `github.com/jessevdk/go-flags`, while the Codex refactor only preserved interactive setup questions and runtime environment-variable updates. That is consistent with the existing vendor pattern, but it also means Codex-specific configuration is not exposed as first-class command flags.
    - The reviewed Codex files do not add YAML parsing or YAML-backed config structures. Fabric does use YAML in other subsystems such as template extensions, but this refactor does not touch those paths and does not introduce a parallel YAML configuration source for Codex.
    - One small consistency gap remains in the env naming itself: the Codex setup uses explicit env names with spaces (`\"Base URL\"`, `\"Auth Base URL\"`), which `BuildEnvVariable(...)` normalizes into `CODEX_BASE_URL` and `CODEX_AUTH_BASE_URL` in `internal/plugins/plugin.go:336-348`. That behavior is inherited from the shared plugin helper and matches the pre-refactor implementation, so it is not a regression.

- [x] **Logging patterns**: Reviewed logging usage across `internal/plugins/ai/codex/auth_transport.go`, `internal/plugins/ai/codex/codex.go`, `internal/plugins/ai/codex/oauth.go`, the shared logger in `internal/log/log.go`, and debug flag wiring in `internal/cli/flags.go`. No PR images were attached for this task.
  - Use standard `log` package
  - Debug levels via `--debug` flag
  - No sensitive data in logs
  - Review notes:
    - The Codex refactor does not use the standard library `log` package directly. Instead it follows Fabric's existing repository-wide convention of routing debug output through `internal/log` via `debuglog.Debug(...)` in `internal/plugins/ai/codex/codex.go:108,155,179`, `internal/plugins/ai/codex/auth_transport.go:64,163`, and `internal/plugins/ai/codex/oauth.go:62`. Relative to the task checklist wording, that is a convention mismatch in the checklist, not a Codex-specific regression.
    - Debug verbosity is still controlled through Fabric's shared `--debug` flag path rather than any Codex-local switch. `internal/cli/flags.go:110-174,240-251` parses `--debug`, maps numeric levels into `internal/log.Level`, and sets the global logger level with `debuglog.SetLevel(...)`, so Codex inherits the same basic/detailed/trace/wire gating used elsewhere in the project.
    - The current Codex debug statements avoid logging bearer tokens, refresh tokens, authorization codes, PKCE values, request bodies, or raw provider response bodies. The logged fields are limited to base URLs, the localhost callback port, retry events, and the resolved account ID, which is materially safer than logging credential material.
    - There is still a small information-exposure concern at trace level: `internal/plugins/ai/codex/codex.go:179` logs the full `/models` request URL including the `client_version` query parameter. That is low risk, but other vendors generally avoid logging fully expanded authenticated request metadata unless needed for troubleshooting.
    - The more important sensitive-data concern remains adjacent to logging rather than inside the debug statements themselves: provider-supplied auth failure text is still propagated back to the browser callback through `http.Error(w, err.Error(), ...)` in `internal/plugins/ai/codex/oauth.go:174,180`. That is not a log leak, but it is still an externally visible disclosure path worth carrying into the issue summary.

### Task 7: Run Static Analysis

- [x] **Check for modernization**: Ran `go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest ./...` from `/Users/kayvan/src/fabric`; it exited cleanly and produced no modernization suggestions.
  ```bash
  go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest ./...
  ```
  Note any suggestions.

- [x] **Check formatting**: Ran `gofmt -l .` from `/Users/kayvan/src/fabric`; it exited cleanly and reported no unformatted files.
  ```bash
  gofmt -l .
  ```
  Flag any unformatted files.

- [ ] **Check vet**: Run:
  ```bash
  go vet ./...
  ```
  Note any issues.

### Task 8: Document Go Issues

- [ ] **Create GO_ISSUES.md**: Write findings to `/Users/kayvan/src/fabric/.maestro/playbooks/GO_ISSUES.md`:

```markdown
# Go Code Quality Issues

## Critical Issues
[Must fix - compiler errors, data races, panics]

## Major Issues
[Should fix - error handling, context misuse, interface violations]

## Minor Issues
[Nice to fix - naming, documentation, style]

## Suggestions
[Optional improvements, modernization opportunities]

## Static Analysis Results

### Modernize
[Results from modernize tool]

### Gofmt
[Unformatted files if any]

### Go Vet
[Vet issues if any]

## Positive Observations
[Good Go practices observed]
```

For each issue include:
- File and line number
- Issue description
- Suggested fix
- Severity: Critical / Major / Minor / Suggestion

## Success Criteria

- All Go files reviewed for idioms
- Error handling verified
- Context usage checked
- Static analysis completed
- GO_ISSUES.md created

## Status

Mark complete when Go review document is created.

---

**Next**: Document 3 will validate Fabric pattern system changes.
