# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What is Fabric

Fabric is an open-source Go CLI tool and REST API framework for augmenting humans using AI. It applies reusable "patterns" (markdown-based system prompts) to user input and routes them through configurable AI vendors. The main binary is `fabric`.

## Build & Run Commands

```bash
# Build the CLI
go build ./cmd/fabric/...

# Run all Go tests
go test ./...

# Run tests for a specific package
go test ./internal/core/...
go test ./internal/plugins/template/...

# Run a single test
go test ./internal/core/... -run TestChatter

# Build and install the binary
go install ./cmd/fabric/...

# Generate Swagger docs (requires swag CLI)
swag init -g internal/server/serve.go -o docs

# Web UI (in web/ directory)
cd web && npm install && npm run dev   # dev server
cd web && npm run build               # production build
cd web && npm run lint                # ESLint + Prettier check
cd web && npm run check               # Svelte type checking
```

## Architecture Overview

### Entry Point → CLI Layer
`cmd/fabric/main.go` calls `internal/cli.Cli(version)`. The CLI layer (`internal/cli/`) parses flags (using `go-flags`), loads YAML config from `~/.config/fabric/config.yaml`, reads stdin, and dispatches to handlers.

### Plugin Registry (`internal/core/plugin_registry.go`)
The central wiring object. `NewPluginRegistry(db)` instantiates and registers all AI vendor plugins, tools (YouTube, Jina, Spotify, etc.), and managers. Configuration is loaded from `~/.config/fabric/.env` (via godotenv). Vendor plugins that fail `Configure()` are excluded from the active `VendorManager`.

### Chatter (`internal/core/chatter.go`)
The core execution engine. `PluginRegistry.GetChatter(...)` resolves the correct vendor and model. `Chatter.Send(request, opts)` assembles a session, applies pattern/context/strategy system prompts, handles template variable substitution, and calls `vendor.SendStream()` or `vendor.Send()`.

### Plugin System (`internal/plugins/`)
- **`plugin.go`** — `Plugin` interface + `PluginBase` struct with settings/setup questions backed by env vars
- **`ai/`** — AI vendor implementations. Each vendor implements the `Vendor` interface: `ListModels()`, `Send()`, `SendStream()`, `NeedsRawMode()`. Active vendors: OpenAI, Anthropic, Gemini, Azure, Bedrock, Ollama, VertexAI, Perplexity, LMStudio, DigitalOcean, Exolab, Copilot, ClaudeCode, plus any OpenAI-compatible providers from `openai_compatible/`
- **`db/fsdb/`** — Filesystem database for patterns, sessions, and contexts stored under `~/.config/fabric/`
- **`template/`** — Template variable substitution (`{{variable}}`) with extensions for file, fetch, hash, sys, datetime, and text operations
- **`strategy/`** — Strategy files (JSON, under `data/strategies/`) prepended as system prompts

### Data Storage (`~/.config/fabric/`)
- `.env` — API keys and vendor settings (written by `fabric --setup`)
- `patterns/` — Pattern directories, each with a `system.md` file
- `sessions/` — Saved chat sessions as JSON
- `contexts/` — Reusable context text files
- `config.yaml` — Optional YAML config mirroring CLI flags

### REST API Server (`internal/server/`)
Started with `fabric --serve`. Uses Gin + Swagger. Handlers for `/patterns`, `/contexts`, `/sessions`, `/chat`, `/models`, `/strategies`, `/youtube`. Swagger docs at `localhost:8080/swagger/index.html`.

### Web UI (`web/`)
SvelteKit + Tailwind CSS + Skeleton UI. Built separately from the Go backend; connects to the REST API. Uses pnpm (lock file present).

### i18n (`internal/i18n/`)
String translations embedded from `internal/i18n/locales/*.json`. Initialized with `i18n.Init(languageCode)`. Supports de, en, es, fa, fr, it, ja, pt-BR, pt-PT, zh.

## Key Conventions

- AI vendor env var names are auto-generated: `BuildEnvVariablePrefix(name)` uppercases and underscores the vendor name as a prefix (e.g., `OPENAI_API_KEY`)
- Bedrock is only registered when `BEDROCK_AWS_REGION` is set alongside valid AWS credentials
- Pattern files live in a directory named after the pattern; the system prompt is always `system.md`
- Template variables use `{{varname}}` syntax; pattern variables use `-v=#name:value` CLI syntax
- The `NeedsRawMode()` method on vendors controls whether a model skips chat parameters (temperature, etc.)
- Config YAML values are overridden by explicit CLI flags; CLI flag detection uses reflection on struct tags
