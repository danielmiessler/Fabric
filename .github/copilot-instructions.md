# Copilot instructions for this repository

Purpose: Help future Copilot sessions navigate, build, test, and make repository-specific changes efficiently.

---

## Build, test, and lint (exact commands)

Root (Go CLI / API):
- Build binary: `go build -o fabric ./cmd/fabric`
- Install (local): `go install github.com/danielmiessler/fabric/cmd/fabric@latest`
- Run server: `go run ./cmd/fabric --serve`
- All tests: `go test ./...`
- Verbose: `go test -v ./...`
- Single package: `go test ./<package/path>` (example: `go test ./internal/cli`)
- Coverage: `go test -cover ./...`
- Format: `gofmt -w .` (or `gofmt` on specific files)
- Vet/modernize: `go vet ./...` and CI uses: `go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest ./...`
- CI uses `nix flake check` as an additional formatting/validation step

Web (SvelteKit) — located in `/web`:
- Install (preferred): `cd web && pnpm install` (npm/yarn acceptable)
- Dev server: `cd web && npm run dev`
- Build: `cd web && npm run build`
- Preview: `cd web && npm run preview`
- Tests (vitest): `cd web && npm run test`
- Run a single frontend test: `cd web && npm run test -- -t "<test name|pattern>"`
- Lint: `cd web && npm run lint` (runs Prettier + ESLint)
- Format: `cd web && npm run format`

Notes:
- `web`'s `prebuild`/`predev` copy `scripts/pattern_descriptions/pattern_descriptions.json` into `web/static/data/`; keep that in mind when editing pattern descriptions.

---

## High-level architecture (big picture)

- Monorepo centered on a Go application + web UI.
  - `cmd/fabric` — CLI entrypoint(s) and the main binary.
  - `internal/` — application code (server handlers, providers, patterns runtime, utilities).
  - `data/patterns/` — pattern content repository (one pattern per subdirectory).
  - `web/` — SvelteKit frontend for Fabric's web UI and documentation site.
  - `docs/` — docs and developer guides (CONTRIBUTING.md, Swagger outputs, images).
  - `scripts/` — helper scripts (installer, docker, pattern extraction/generation, changelog generation).
  - `.github/workflows/` — CI and release pipelines (CI runs Go tests, modernize, and `nix flake check`).

- REST API: implemented using Gin (`internal/server`). Endpoints are annotated for swaggo; Swagger UI is generated and committed to `docs/` when handlers change.

- Provider integrations: multiple vendor adapters (OpenAI, Anthropic, Ollama, Bedrock, Azure, etc.) are implemented as internal/provider modules and invoked by pattern runtime.

---

## Key repository conventions (non-obvious, important)

- Patterns format and placement:
  - Each pattern lives under `data/patterns/<pattern-name>/`.
  - The canonical pattern file is `system.md` with top-level sections like `# IDENTITY and PURPOSE`, `# STEPS`, `# OUTPUT`, `# EXAMPLE`.
  - The `web` app and scripts assume `scripts/pattern_descriptions/pattern_descriptions.json` exists and is copied into `web/static/data/` during `prebuild`/`predev`.

- Swagger / API docs:
  - Handlers must include swaggo annotations (see docs/CONTRIBUTING.md examples).
  - After changing handlers, run `swag init -g internal/server/serve.go -o docs` and commit `docs/swagger.json`, `docs/swagger.yaml`, and `docs/docs.go`.

- Changelogs and PR requirements:
  - Every PR should include a changelog entry generated with the helper: `go run ./cmd/generate_changelog --ai-summarize --incoming-pr YOUR_PR_NUMBER` (see docs).
  - Keep PRs small and focused; large, sweeping PRs should be discussed via an issue first.

- Tests and tooling:
  - Backend tests use standard `go test` and `stretchr/testify` where applicable.
  - CI runs `go test -v ./...`, modernize analyzer, and `nix flake check` — reproducing CI locally may require nix.
  - Frontend tests use `vitest`. Use the `-t` flag to filter test names.

- Go version and modules:
  - `go.mod` pins Go 1.25.1; target that runtime for local builds.

- Releasing / workflows:
  - Release and tag automation live in `.github/workflows/release.yml` and `update-version-and-create-tag.yml`. Avoid manually duplicating release metadata when possible.

---

## Files / assistant configs checked

Searched for and considered repository assistant rules or AI assistant files (CLAUDE.md, AGENTS.md, .cursorrules, .windsurfrules, CONVENTIONS.md, AIDER_CONVENTIONS.md). None of those custom assistant-config files were present or applicable.

---

If useful, Copilot sessions should begin by running the repository tests (`go test ./...`) and `cd web && pnpm install && npm run test` to ensure local dev health before proposing code changes.

---

Would you like me to configure any MCP servers relevant to this repo (examples: Playwright for web E2E tests, or a browser-based test server)? If yes, specify which server(s) to configure and whether to add CI steps or local dev scripts.

Summary: created .github/copilot-instructions.md capturing build/test/lint commands, high-level architecture, and repository-specific conventions. Want any adjustments or additional coverage (examples, shortcuts, or package manager preference)?
