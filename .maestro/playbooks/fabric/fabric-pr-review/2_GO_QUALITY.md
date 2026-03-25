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

- [ ] **Package structure**: Check:
  - `internal/` packages are truly internal
  - No circular dependencies
  - Clear package boundaries
  - Appropriate file sizes

- [ ] **Naming conventions**: Verify:
  - CamelCase for exported identifiers
  - camelCase for unexported identifiers
  - Meaningful, descriptive names
  - Acronyms are all caps (HTTP, API, ID)

- [ ] **Documentation**: Check:
  - Exported functions have doc comments
  - Package-level documentation exists
  - Complex logic is explained
  - No stale comments

### Task 4: Review Concurrency

- [ ] **Goroutine safety**: Look for:
  - Race conditions on shared state
  - Proper channel usage (closing, direction)
  - Context-aware goroutines
  - No goroutine leaks

- [ ] **Streaming**: For streaming responses:
  - Channels are properly buffered
  - Errors are communicated correctly
  - Cleanup happens on cancellation

### Task 5: Review API Changes

- [ ] **Breaking changes**: If public APIs change:
  - Are changes backward compatible?
  - Are deprecation notices added?
  - Is the CHANGELOG updated?

- [ ] **Function signatures**: Verify:
  - Context is first parameter
  - Options pattern for many parameters
  - Error is last return value

### Task 6: Check Fabric-Specific Patterns

- [ ] **Plugin patterns**: For AI providers:
  - Implement `VendorPlugin` interface
  - Handle streaming via callbacks
  - Support model listing
  - Handle context cancellation

- [ ] **Configuration patterns**: For config changes:
  - Environment variables via `godotenv`
  - Flags via `go-flags`
  - YAML config support

- [ ] **Logging patterns**: For logging:
  - Use standard `log` package
  - Debug levels via `--debug` flag
  - No sensitive data in logs

### Task 7: Run Static Analysis

- [ ] **Check for modernization**: Run:
  ```bash
  go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest ./...
  ```
  Note any suggestions.

- [ ] **Check formatting**: Run:
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
