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

## Getting Help

- Check existing issues first
- Ask questions in discussions
- Tag maintainers for urgent issues
- Be patient - maintainers are volunteers

## License

By contributing, you agree your contributions will be licensed under the MIT License.
