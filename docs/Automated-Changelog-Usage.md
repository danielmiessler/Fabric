# Automated Changelog System - Developer Guide

This guide explains how to use the new automated changelog system for the Fabric project.

## Overview

The automated changelog system allows developers to pre-process their PR changelog entries during development, which are then automatically aggregated during the release process. This eliminates manual CHANGELOG.md editing and reduces merge conflicts.

## Developer Workflow

### Step 1: Create Your Feature Branch and PR

Work on your feature as usual and create a pull request.

### Step 2: Generate Changelog Entry

Once your PR is ready for review, generate a changelog entry:

```bash
cd cmd/generate_changelog
go build -o generate_changelog .
./generate_changelog --incoming-pr YOUR_PR_NUMBER
```

For example, if your PR number is 1672:

```bash
./generate_changelog --incoming-pr 1672
```

### Step 3: Validation

The tool will validate:

- ✅ PR exists and is open
- ✅ PR is mergeable (no conflicts)
- ✅ Your working directory is clean

If any validation fails, fix the issues and try again.

### Step 4: Review Generated Entry

The tool will:

1. Create `./cmd/generate_changelog/incoming/1672.txt`
2. Generate an AI-enhanced summary (if `--ai-summarize` is enabled)
3. Auto-commit the file to your branch (use `--push` to also push to remote)

Review the generated file and edit if needed:

```bash
cat ./cmd/generate_changelog/incoming/1672.txt
```

### Step 5: Include in PR

The incoming changelog entry is now part of your PR and will be reviewed along with your code changes.

## Example Generated Entry

```markdown
### PR [#1672](https://github.com/danielmiessler/fabric/pull/1672) by [ksylvan](https://github.com/ksylvan): Changelog Generator Enhancement

- Added automated CI/CD integration for changelog generation
- Implemented pre-processing of PR entries during development
- Enhanced caching system for better performance
- Added validation for mergeable PR states
```

## Command Options

### `--incoming-pr`

Pre-process a specific PR for changelog generation.

**Usage**: `./generate_changelog --incoming-pr PR_NUMBER`

**Requirements**:

- PR must be open
- PR must be mergeable (no conflicts)
- Working directory must be clean (no uncommitted changes)
- GitHub token must be available (`GITHUB_TOKEN` env var or `--token` flag)

**Mutual Exclusivity**: Cannot be used with `--process-prs` flag

### `--incoming-dir`

Specify custom directory for incoming PR files (default: `./cmd/generate_changelog/incoming`).

**Usage**: `./generate_changelog --incoming-pr 1672 --incoming-dir ./custom/path`

### `--process-prs`

Process all incoming PR files for release aggregation. Used by CI/CD during release creation.

**Usage**: `./generate_changelog --process-prs {new_version_string}`

**Mutual Exclusivity**: Cannot be used with `--incoming-pr` flag

### `--ai-summarize`

Enable AI-enhanced summaries using Fabric integration.

**Usage**: `./generate_changelog --incoming-pr 1672 --ai-summarize`

### `--push`

Enable automatic git push after creating an incoming entry. By default, the commit is created locally but not pushed to the remote repository.

**Usage**: `./generate_changelog --incoming-pr 1672 --push`

**Note**: When using `--push`, ensure you have proper authentication configured (SSH keys or GITHUB_TOKEN environment variable).

## Troubleshooting

### "PR is not open"

Your PR has been closed or merged. Only open PRs can be processed.

### "PR is not mergeable"

Your PR has merge conflicts or other issues preventing it from being merged. Resolve conflicts and ensure the PR is in a mergeable state.

### "Working directory is not clean"

You have uncommitted changes. Commit or stash them before running the tool.

### "Failed to fetch PR"

Check your GitHub token and network connection. Ensure the PR number exists.

## CI/CD Integration

The system automatically processes all incoming PR files during the release workflow. No manual intervention is required.

When a release is created:

1. All `incoming/*.txt` files are aggregated using `--process-prs`
2. Version is detected from `version.nix` or latest git tag
3. A new version entry is created in CHANGELOG.md
4. Incoming files are cleaned up (removed)
5. Changes are staged for the release commit (CHANGELOG.md and cache file)

## Best Practices

1. **Run early**: Generate your changelog entry as soon as your PR is ready for review
2. **Review content**: Always review the generated entry and edit if necessary
3. **Keep it updated**: If you make significant changes to your PR, regenerate the entry
4. **Use AI summaries**: Enable `--ai-summarize` for more professional, consistent formatting

## Advanced Usage

### Custom GitHub Token

```bash
./generate_changelog --incoming-pr 1672 --token YOUR_GITHUB_TOKEN
```

### Custom Repository Path

```bash
./generate_changelog --incoming-pr 1672 --repo /path/to/repo
```

### Disable Caching

```bash
./generate_changelog --incoming-pr 1672 --no-cache
```

### Enable Auto-Push

```bash
./generate_changelog --incoming-pr 1672 --push
```

This creates the commit locally and pushes it to the remote repository. By default, commits are only created locally, allowing you to review changes before pushing manually.

**Authentication**: The tool automatically detects GitHub repositories and uses the GITHUB_TOKEN environment variable for authentication when pushing. For SSH repositories, ensure your SSH keys are properly configured.

## Integration with Existing Workflow

This system is fully backward compatible. The existing changelog generation continues to work unchanged. The new features are opt-in and only activated when using the new flags.

## Support

If you encounter issues:

1. Check this documentation
2. Verify your GitHub token has appropriate permissions
3. Ensure your PR meets the validation requirements
4. Check the tool's help: `./generate_changelog --help`

For bugs or feature requests, please create an issue in the repository.
