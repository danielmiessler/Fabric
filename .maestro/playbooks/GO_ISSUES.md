---
type: report
title: Go Code Quality Issues for PR 2063
created: 2026-03-25
tags:
  - go
  - pr-review
  - codex
  - quality
related:
  - '[[REVIEW_SCOPE]]'
  - '[[2_GO_QUALITY]]'
---

# Go Code Quality Issues

## Critical Issues

None.

## Major Issues

- Severity: Major
  File: `internal/plugins/ai/codex/errors.go:24`
  Issue: `extractErrorMessage` synthesizes a capitalized error string (`Codex request failed with status %d`), which does not match the repo's documented lowercase error style.
  Suggested fix: Normalize synthesized error messages to lowercase and keep formatting consistent with the rest of the package.

- Severity: Major
  File: `internal/plugins/ai/codex/oauth.go:219`
  Issue: Error wrapping uses `fmt.Errorf(... %w ...)` instead of the project convention `pkg/errors.Wrap`, and the message starts with `Codex`.
  Suggested fix: Switch to `errors.Wrap` or `errors.Wrapf` and normalize the message casing.

- Severity: Major
  File: `internal/plugins/ai/codex/oauth.go:245`
  Issue: The invalid auth base URL error is capitalized (`invalid Codex auth base URL: %w`) and does not follow the documented naming style for initialisms.
  Suggested fix: Lowercase the provider name in the message or rephrase the error to avoid mixed-case provider-specific wording.

- Severity: Major
  File: `internal/plugins/ai/codex/token.go:45`
  Issue: JWT parsing returns `JWT did not include an exp claim`, which is capitalized and inconsistent with standard Go error-string style.
  Suggested fix: Rephrase to lowercase, for example `jwt did not include an exp claim`.

- Severity: Major
  File: `internal/plugins/ai/codex/token.go:61`
  Issue: `invalid JWT format` uses a capitalized acronym in the error string and is inconsistent with the repo's documented lowercase error style.
  Suggested fix: Rephrase to lowercase, for example `invalid jwt format`.

- Severity: Major
  File: `internal/plugins/ai/codex/errors.go:15`
  Issue: Upstream non-401 response text is surfaced through `extractErrorMessage`, then echoed back to the local OAuth callback handler.
  Suggested fix: Sanitize provider-supplied response bodies before returning them to users, and keep detailed upstream content limited to internal diagnostics.

- Severity: Major
  File: `internal/plugins/ai/codex/oauth.go:174`
  Issue: Browser callback failures send `err.Error()` directly to the user via `http.Error`, which can expose provider-supplied detail gathered during token exchange.
  Suggested fix: Return a generic user-facing message and log or wrap the internal cause separately.

- Severity: Major
  File: `internal/plugins/ai/codex/oauth.go:180`
  Issue: The OAuth timeout/error path also returns raw error text to the browser, extending the same disclosure path as the callback failure handler.
  Suggested fix: Replace direct error echoing with a generic callback failure message.

- Severity: Major
  File: `internal/plugins/ai/vendor.go:12`
  Issue: `SendStream` and `ListModels` do not accept `context.Context`, so Codex cannot propagate caller cancellation for streaming or model-list requests.
  Suggested fix: Evolve the shared vendor interface to make these operations context-aware, then thread the caller context through Codex and other vendors.

- Severity: Major
  File: `internal/plugins/ai/codex/codex.go:266`
  Issue: Streaming uses `context.Background()` when opening the upstream SSE stream, so abandoned consumers cannot proactively cancel the request.
  Suggested fix: Update the shared interface to accept context and use the caller's context for stream creation.

- Severity: Major
  File: `internal/plugins/ai/codex/codex.go:273`
  Issue: `SendStream` performs blocking sends to the caller-provided channel, which can stall indefinitely if the channel is unbuffered or the consumer stops reading.
  Suggested fix: Document the buffering requirement explicitly or redesign the streaming contract to support cancellation-aware, non-blocking delivery.

- Severity: Major
  File: `internal/plugins/ai/codex/codex.go:289`
  Issue: Mid-stream failures are only returned from `SendStream()` and are not emitted as an in-band `domain.StreamTypeError` event.
  Suggested fix: Decide on one terminal error-delivery pattern and make it consistent across vendors, ideally including in-band stream error updates for partially streamed responses.

## Minor Issues

- Severity: Minor
  File: `internal/plugins/ai/codex/codex.go:48`
  Issue: The exported `Client` type is missing a doc comment.
  Suggested fix: Add a short doc comment describing the Codex client and its role in the vendor plugin system.

- Severity: Minor
  File: `internal/plugins/ai/codex/codex.go:1`
  Issue: `codex.go` carries multiple responsibilities including setup, configuration, model listing, request execution, and helper logic.
  Suggested fix: Split setup/model management from request execution before additional behavior accumulates in the entrypoint file.

- Severity: Minor
  File: `internal/plugins/ai/openai/openai.go`
  Issue: Embedded shared OpenAI client fields use `ApiBaseURL` and `ApiClient`, which do not follow strict Go initialism casing.
  Suggested fix: Rename these fields to `APIBaseURL` and `APIClient` in a separate cross-provider cleanup if the team wants to enforce initialism style consistently.

## Suggestions

- Severity: Suggestion
  File: `internal/plugins/ai/codex/codex_test.go`
  Issue: Existing streaming coverage exercises the success path, but there is no Codex-specific test for mid-stream transport failures or abandoned consumers.
  Suggested fix: Add tests that simulate transport errors after partial output and verify the chosen terminal error semantics.

- Severity: Suggestion
  File: `internal/plugins/ai/codex/auth_transport.go:90`
  Issue: The package mixes `fmt.Errorf(... %w ...)` with the repo's documented `pkg/errors` wrapping convention across several files.
  Suggested fix: Standardize one wrapping approach across the package, preferably the repository convention if it is still current policy.

- Severity: Suggestion
  File: `internal/plugins/ai/codex/codex.go:179`
  Issue: Debug logging includes the fully expanded `/models` request URL, which is low risk but broader than needed for routine debugging.
  Suggested fix: Log the endpoint path or host instead of the entire URL unless the full query string is required for diagnostics.

## Static Analysis Results

### Modernize

- `go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest ./...`
- Result: clean run, no modernization suggestions.

### Gofmt

- `gofmt -l .`
- Result: clean run, no unformatted files reported.

### Go Vet

- `go vet ./...`
- Result: clean run, no vet issues reported.

## Positive Observations

- The Codex package keeps shared token mutation serialized behind `tokenMu`, and `go test -race ./internal/plugins/ai/codex` completed without reporting a race.
- Request-driven paths consistently propagate caller context where the shared interface allows it, including OAuth token exchange and authenticated request retries.
- The refactor improved package separation by moving auth transport, OAuth flow, token parsing, and API error normalization into focused files.
- The package does not introduce `panic()` in library code and keeps the public `ai.Vendor` surface backward compatible with `origin/main`.
