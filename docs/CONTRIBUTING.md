# Contributing to Fabric

Thanks for contributing to Fabric! Here's what you need to know to get started quickly.

## Quick Setup

### Prerequisites

- Go 1.24+ installed
- Git configured with your details
- GitHub CLI (`gh`)

### Getting Started

```bash
# Clone your fork (upstream is set automatically)
gh repo clone YOUR_GITHUB_USER/fabric
cd fabric
go build -o fabric ./cmd/fabric
./fabric --setup

# Run tests
go test ./...
```

## Development Guidelines

### Code Style

- Follow standard Go conventions (`gofmt`, `golint`)
- Use meaningful variable and function names
- Write tests for new functionality
- Keep functions focused and small

### Commit Messages

Use descriptive commit messages:

```text
feat: add new pattern for code analysis
fix: resolve OAuth token refresh issue
docs: update installation instructions
```

### Project Structure

- `cmd/` - Executable commands
- `internal/` - Private application code
- `data/patterns/` - AI patterns
- `docs/` - Documentation

## Pull Request Process

### Pull Request Guidelines

**Keep pull requests focused and minimal.**

PRs that touch a large number of files (50+) without clear functional justification will likely be rejected without detailed review.

#### Why we enforce this

- **Reviewability**: Large PRs are effectively un-reviewable. Studies show reviewer effectiveness drops significantly after ~200-400 lines of code. A 93-file "cleanup" PR cannot receive meaningful review.
- **Git history**: Sweeping changes pollute `git blame`, making it harder to trace when and why functional changes were made.
- **Merge conflicts**: Large PRs increase the likelihood of conflicts with other contributors' work.
- **Risk**: More changed lines means more opportunities for subtle bugs, even in "safe" refactors.

#### What to do instead

If you have a large change in mind, break it into logical, independently-mergeable slices. For example:

- ✅ "Replace `interface{}` with `any` across codebase" (single mechanical change, easy to verify)
- ✅ "Migrate to `strings.CutPrefix` in `internal/cli`" (scoped to one package)
- ❌ "Modernize codebase with multiple idiom updates" (too broad, impossible to review)

For sweeping refactors or style changes, **open an issue first** to discuss the approach with maintainers before investing time in the work.

### Changelog Generation (REQUIRED)

After opening your PR, generate a changelog entry:

```bash
go run ./cmd/generate_changelog --ai-summarize --incoming-pr YOUR_PR_NUMBER
```

**Requirements:**

- PR must be open and mergeable
- Working directory must be clean
- GitHub token available (GITHUB_TOKEN env var)

**Optional flags:**

- `--ai-summarize` - Enhanced AI-generated summaries
- `--push` - Auto-push the changelog commit

### PR Guidelines

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Write/update tests
5. Generate changelog entry (see above)
6. Submit PR with clear description

### Review Process

- PRs require maintainer review
- Address feedback promptly
- Keep PRs focused on single features/fixes
- Update changelog if you make significant changes

## Testing

### Run Tests

```bash
# All tests
go test ./...

# Specific package
go test ./internal/cli

# With coverage
go test -cover ./...
```

### Test Requirements

- Unit tests for core functionality
- Integration tests for external dependencies
- Examples in documentation

### Live Pipeline Smoke (Model-Backed)

For model-backed built-in pipeline verification, run the bounded smoke suite locally:

```bash
go build -o fabric ./cmd/fabric
OPENAI_API_KEY=... \
python3 scripts/pipelines/live_smoke.py \
  --fabric-bin ./fabric \
  --vendor OpenAI \
  --model gpt-4.1-mini
```

Optional subset run:

```bash
python3 scripts/pipelines/live_smoke.py --fabric-bin ./fabric --pipelines technical-study-guide,zoom-tech-note,zoom-tech-note-deep-pass
```

CI also runs the same class of checks in:

- `.github/workflows/pipeline-live-smoke.yml`

## Patterns

### Creating Patterns

Patterns go in `data/patterns/[pattern-name]/system.md`:

```markdown
# IDENTITY and PURPOSE
You are an expert at...

# STEPS
- Step 1
- Step 2

# OUTPUT
- Output format requirements

# EXAMPLE
Example output here
```

### Pattern Guidelines

- Use clear, actionable language
- Provide specific output formats
- Include examples when helpful
- Test with multiple AI providers

## Documentation

- Update README.md for new features
- Add docs to `docs/` for complex features
- Include usage examples
- Keep documentation current

### REST API Documentation

When adding or modifying REST API endpoints, you must update the Swagger documentation:

**1. Add Swagger annotations to your handler:**

```go
// HandlerName godoc
// @Summary Short description of what this endpoint does
// @Description Detailed description of the endpoint's functionality
// @Tags category-name
// @Accept json
// @Produce json
// @Param name path string true "Parameter description"
// @Param body body RequestType true "Request body description"
// @Success 200 {object} ResponseType "Success description"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Server error"
// @Security ApiKeyAuth
// @Router /endpoint/path [get]
func (h *Handler) HandlerName(c *gin.Context) {
    // Implementation
}
```

**2. Regenerate Swagger documentation:**

```bash
# Install swag CLI if you haven't already
go install github.com/swaggo/swag/cmd/swag@latest

# Generate updated documentation
swag init -g internal/server/serve.go -o docs
```

**3. Commit the generated files:**

The following files will be updated and should be committed:

- `docs/swagger.json`
- `docs/swagger.yaml`
- `docs/docs.go`

**4. Test your changes:**

Start the server and verify your endpoint appears in Swagger UI:

```bash
go run ./cmd/fabric --serve
# Open http://localhost:8080/swagger/index.html
```

**Examples to follow:**

- Chat endpoint: `internal/server/chat.go:58-68`
- Patterns endpoint: `internal/server/patterns.go:36-45`
- Models endpoint: `internal/server/models.go:20-28`

**Common annotation tags:**

- `@Summary` - One-line description (required)
- `@Description` - Detailed explanation
- `@Tags` - Logical grouping (e.g., "patterns", "chat", "models")
- `@Accept` - Input content type (e.g., "json")
- `@Produce` - Output content type (e.g., "json", "text/event-stream")
- `@Param` - Request parameters (path, query, body)
- `@Success` - Successful response (include status code and type)
- `@Failure` - Error responses
- `@Security` - Authentication requirement (use "ApiKeyAuth" for API key)
- `@Router` - Endpoint path and HTTP method

For complete Swagger annotation syntax, see the [swaggo documentation](https://github.com/swaggo/swag#declarative-comments-format)

## Getting Help

- Check existing issues first
- Ask questions in discussions
- Tag maintainers for urgent issues
- Be patient - maintainers are volunteers

## License

By contributing, you agree your contributions will be licensed under the MIT License.
