# Changelog Generator

A high-performance changelog generator for Git repositories that automatically creates comprehensive, well-formatted changelogs from your git history and GitHub pull requests.

## Features

- **One-pass git history walking**: Efficiently processes entire repository history in a single pass
- **Automatic PR detection**: Extracts pull request information from merge commits
- **GitHub API integration**: Fetches detailed PR information including commits, authors, and descriptions
- **Smart caching**: SQLite-based caching for instant incremental updates
- **Unreleased changes**: Tracks all commits since the last release
- **Concurrent processing**: Parallel GitHub API calls for improved performance
- **Flexible output**: Generate complete changelogs or target specific versions
- **GraphQL optimization**: Ultra-fast PR fetching using GitHub GraphQL API (~5-10 calls vs 1000s)
- **Intelligent sync**: Automatically syncs new PRs every 24 hours or when missing PRs are detected
- **AI-powered summaries**: Optional Fabric integration for enhanced changelog summaries
- **Advanced caching**: Content-based change detection for AI summaries with hash comparison
- **Author type detection**: Distinguishes between users, bots, and organizations
- **Lightning-fast incremental updates**: SHA→PR mapping for instant git operations

## Installation

```bash
go install github.com/danielmiessler/fabric/cmd/generate_changelog@latest
```

## Usage

### Basic usage (generate complete changelog)

```bash
generate_changelog
```

### Save to file

```bash
generate_changelog -o CHANGELOG.md
```

### Generate for specific version

```bash
generate_changelog -v v1.4.244
```

### Limit to recent versions

```bash
generate_changelog -l 10
```

### Using GitHub token for private repos or higher rate limits

```bash
export GITHUB_TOKEN=your_token_here
generate_changelog

# Or pass directly
generate_changelog --token your_token_here
```

### AI-enhanced summaries

```bash
# Enable AI summaries using Fabric
generate_changelog --ai-summarize

# Use custom model for AI summaries
FABRIC_CHANGELOG_SUMMARIZE_MODEL=claude-opus-4 generate_changelog --ai-summarize
```

### Cache management

```bash
# Rebuild cache from scratch
generate_changelog --rebuild-cache

# Force a full PR sync from GitHub
generate_changelog --force-pr-sync

# Disable cache usage
generate_changelog --no-cache

# Use custom cache location
generate_changelog --cache /path/to/cache.db
```

## Command Line Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--repo` | `-r` | Repository path | `.` (current directory) |
| `--output` | `-o` | Output file | stdout |
| `--limit` | `-l` | Limit number of versions | 0 (all) |
| `--version` | `-v` | Generate for specific version | |
| `--save-data` | | Save version data to JSON | false |
| `--cache` | | Cache database file | `./cmd/generate_changelog/changelog.db` |
| `--no-cache` | | Disable cache usage | false |
| `--rebuild-cache` | | Rebuild cache from scratch | false |
| `--force-pr-sync` | | Force a full PR sync from GitHub | false |
| `--token` | | GitHub API token | `$GITHUB_TOKEN` |
| `--ai-summarize` | | Generate AI-enhanced summaries using Fabric | false |
| `--release` | | Update GitHub release description with AI summary for version | |

## Output Format

The generated changelog follows this structure:

```markdown
# Changelog

## Unreleased

### PR [#1601](url) by [author](profile): PR Title
- Change description 1
- Change description 2

### Direct commits
- Direct commit message 1
- Direct commit message 2

## v1.4.244 (2025-07-09)

### PR [#1598](url) by [author](profile): PR Title
- Change description
...
```

## How It Works

1. **Git History Walking**: The tool walks through your git history from newest to oldest commits
2. **Version Detection**: Identifies version bump commits (pattern: "Update version to vX.Y.Z")
3. **PR Extraction**: Detects merge commits and extracts PR numbers
4. **GitHub API Calls**: Fetches detailed PR information in parallel batches
5. **Change Extraction**: Extracts changes from PR commit messages or PR body
6. **Formatting**: Generates clean, organized markdown output

## Performance

- **Native Go libraries**: Uses go-git and go-github for maximum performance
- **Concurrent API calls**: Processes up to 10 GitHub API requests in parallel
- **Smart caching**: SQLite cache eliminates redundant API calls
- **Incremental updates**: Only processes new commits on subsequent runs
- **GraphQL optimization**: Uses GitHub GraphQL API to fetch all PR data in ~5-10 calls
- **AI-powered summaries**: Optional Fabric integration with intelligent caching
- **Content-based change detection**: AI summaries only regenerated when content changes
- **Lightning-fast git operations**: SHA→PR mapping stored in database for instant lookups

### Major Optimization: GraphQL + Advanced Caching

The tool has been optimized to drastically reduce GitHub API calls and improve performance:

**Previous approach**: Individual API calls for each PR (2 API calls per PR)

- For a repo with 500 PRs: 1,000 API calls

**Current approach**: GraphQL batch fetching with intelligent caching

- For a repo with 500 PRs: ~5-10 GraphQL calls (initial fetch) + 0 calls (subsequent runs with cache)
- **99%+ reduction in API calls after initial run!**

The optimization includes:

1. **GraphQL Batch Fetch**: Uses GitHub's GraphQL API to fetch all merged PRs with commits in minimal calls
2. **Smart Caching**: Stores complete PR data, commits, and SHA mappings in SQLite
3. **Incremental Sync**: Only fetches PRs merged after the last sync timestamp
4. **Automatic Refresh**: PRs are synced every 24 hours or when missing PRs are detected
5. **AI Summary Caching**: Content-based change detection prevents unnecessary AI regeneration
6. **Fallback Support**: If GraphQL fails, falls back to REST API batch fetching
7. **Lightning Git Operations**: Pre-computed SHA→PR mappings for instant commit association

## Requirements

- Go 1.24+ (for installation from source)
- Git repository
- GitHub token (optional, for private repos or higher rate limits)
- Fabric CLI (optional, for AI-enhanced summaries)

## Authentication

The tool supports GitHub authentication via:

1. Environment variable: `export GITHUB_TOKEN=your_token`
2. Command line flag: `--token your_token`
3. `.env` file in the same directory as the binary

### Environment File Support

Create a `.env` file next to the `generate_changelog` binary:

```bash
GITHUB_TOKEN=your_github_token_here
FABRIC_CHANGELOG_SUMMARIZE_MODEL=claude-sonnet-4-20250514
```

The tool automatically loads `.env` files for convenient configuration management.

Without authentication, the tool is limited to 60 GitHub API requests per hour.

## Caching

The SQLite cache stores:

- Version information and commit associations
- Pull request details (title, body, commits, authors)
- Last processed commit SHA for incremental updates
- Last PR sync timestamp for intelligent refresh
- AI summaries with content-based change detection
- SHA→PR mappings for lightning-fast git operations

Cache benefits:

- Instant changelog regeneration
- Drastically reduced GitHub API usage (99%+ reduction after initial run)
- Offline changelog generation (after initial cache build)
- Automatic PR data refresh every 24 hours
- Batch database transactions for better performance
- Content-aware AI summary regeneration

## AI-Enhanced Summaries

The tool can generate AI-powered summaries using Fabric for more polished, professional changelogs:

```bash
# Enable AI summarization
generate_changelog --ai-summarize

# Custom model (default: claude-sonnet-4-20250514)
FABRIC_CHANGELOG_SUMMARIZE_MODEL=claude-opus-4 generate_changelog --ai-summarize
```

### AI Summary Features

- **Content-based change detection**: AI summaries are only regenerated when version content changes
- **Intelligent caching**: Preserves existing summaries and only processes changed versions
- **Content hash comparison**: Uses SHA256 hashing to detect when "Unreleased" content changes
- **Automatic fallback**: Falls back to raw content if AI processing fails
- **Error detection**: Identifies and handles AI processing errors gracefully
- **Minimum content filtering**: Skips AI processing for very brief content (< 256 characters)

### AI Model Configuration

Set the model via environment variable:

```bash
export FABRIC_CHANGELOG_SUMMARIZE_MODEL=claude-opus-4
# or
export FABRIC_CHANGELOG_SUMMARIZE_MODEL=gpt-4
```

AI summaries are cached and only regenerated when:

- Version content changes (detected via hash comparison)
- No existing AI summary exists for the version
- Force rebuild is requested

## Contributing

This tool is part of the Fabric project. Contributions are welcome!

## License

The MIT License. Same as the Fabric project.
