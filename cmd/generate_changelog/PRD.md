# Product Requirements Document: Changelog Generator

## Overview

The Changelog Generator is a high-performance Go tool that automatically generates comprehensive changelogs from git history and GitHub pull requests.

## Goals

1. **Performance**: Very fast. Efficient enough to be used in CI/CD as part of release process.
2. **Completeness**: Capture ALL commits including unreleased changes
3. **Efficiency**: Minimize API calls through caching and batch operations
4. **Reliability**: Handle errors gracefully with proper Go error handling
5. **Simplicity**: Single binary with no runtime dependencies

## Key Features

### 1. One-Pass Git History Algorithm

- Walk git history once from newest to oldest
- Start with "Unreleased" bucket for all new commits
- Switch buckets when encountering version commits
- No need to calculate ranges between versions

### 2. Native Library Integration

- **go-git**: Pure Go git implementation (no git binary required)
- **go-github**: Official GitHub Go client library
- Benefits: Type safety, better error handling, no subprocess overhead

### 3. Smart Caching System

- SQLite-based persistent cache
- Stores: versions, commits, PR details, last processed commit
- Enables incremental updates on subsequent runs
- Instant changelog regeneration from cache

### 4. Concurrent Processing

- Parallel GitHub API calls (up to 10 concurrent)
- Batch PR fetching with deduplication
- Rate limiting awareness

### 5. Enhanced Output

- "Unreleased" section for commits since last version
- Clean markdown formatting
- Configurable version limiting
- Direct commit tracking (non-PR commits)

## Technical Architecture

### Module Structure

```text
cmd/generate_changelog/
├── main.go              # CLI entry point with cobra
├── internal/
│   ├── git/            # Git operations (go-git)
│   ├── github/         # GitHub API client (go-github)
│   ├── cache/          # SQLite caching layer
│   ├── changelog/      # Core generation logic
│   └── config/         # Configuration management
└── changelog.db        # SQLite cache (generated)
```

### Data Flow

1. Git walker collects all commits in one pass
2. Commits bucketed by version (starting with "Unreleased")
3. PR numbers extracted from merge commits
4. GitHub API batch-fetches PR details
5. Cache stores everything for future runs
6. Formatter generates markdown output

### Cache Schema

- **metadata**: Last processed commit SHA
- **versions**: Version names, dates, commit SHAs
- **commits**: Full commit details with version associations
- **pull_requests**: PR details including commits
- Indexes on version and PR number for fast lookups

### Features

- **Unreleased section**: Shows all new commits
- **Better caching**: SQLite vs JSON, incremental updates
- **Smarter deduplication**: Removes consecutive duplicate commits
- **Direct commit tracking**: Shows non-PR commits

### Reliability

- **No subprocess errors**: Direct library usage
- **Type safety**: Compile-time checking
- **Better error handling**: Go's explicit error returns

### Deployment

- **Single binary**: No Python/pip/dependencies
- **Cross-platform**: Compile for any OS/architecture
- **No git CLI required**: Uses go-git library

## Configuration

### Environment Variables

- `GITHUB_TOKEN`: GitHub API authentication token

### Command Line Flags

- `--repo, -r`: Repository path (default: current directory)
- `--output, -o`: Output file (default: stdout)
- `--limit, -l`: Version limit (default: all)
- `--version, -v`: Target specific version
- `--save-data`: Export debug JSON
- `--cache`: Cache file location
- `--no-cache`: Disable caching
- `--rebuild-cache`: Force cache rebuild
- `--token`: GitHub token override

## Success Metrics

1. **Performance**: Generate full changelog in <5 seconds for fabric repo
2. **Completeness**: 100% commit coverage including unreleased
3. **Accuracy**: Correct PR associations and change extraction
4. **Reliability**: Handle network failures gracefully
5. **Usability**: Simple CLI with sensible defaults

## Future Enhancements

1. **Multiple output formats**: JSON, HTML, etc.
2. **Custom version patterns**: Configurable regex
3. **Change categorization**: feat/fix/docs auto-grouping
4. **Conventional commits**: Full support for semantic versioning
5. **GitLab/Bitbucket**: Support other platforms
6. **Web UI**: Interactive changelog browser
7. **Incremental updates**: Update existing CHANGELOG.md file
8. **Breaking change detection**: Highlight breaking changes

## Implementation Status

- ✅ Core architecture and modules
- ✅ One-pass git walking algorithm
- ✅ GitHub API integration with concurrency
- ✅ SQLite caching system
- ✅ Changelog formatting and generation
- ✅ CLI with all planned flags
- ✅ Documentation (README and PRD)

## Conclusion

This Go implementation provides a modern, efficient, and feature-rich changelog generator.
