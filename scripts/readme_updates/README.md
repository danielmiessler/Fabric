# README Update Scripts

This directory contains automation scripts for updating the main README.md file with release information from the changelog database.

## `update_readme_features.py`

A Python script that generates the "Recent Major Features" section for the README by extracting and filtering release information from the changelog SQLite database.

### Usage

```bash
# Generate the Recent Major Features section with default limit (20 releases)
python scripts/readme_updates/update_readme_features.py

# Specify a custom limit
python scripts/readme_updates/update_readme_features.py --limit 15

# Use a custom database path
python scripts/readme_updates/update_readme_features.py --db /path/to/changelog.db
```

### How It Works

1. **Database Connection**: Connects to `cmd/generate_changelog/changelog.db` (or custom path)
2. **Data Extraction**: Queries the `versions` table for release information
3. **Feature Filtering**: Uses heuristics to identify feature/improvement releases
4. **Markdown Generation**: Formats output to match README style

### Feature Detection Heuristics

The script uses keyword-based heuristics to filter releases:

#### Include Keywords (Features/Improvements)
- new, feature, feat, add, introduce, enable, support
- improve, enhance, performance, speed
- option, flag, argument, parameter
- integration, provider, search, tts, audio, model
- cli, ui, web, oauth, sync, database
- notifications, desktop, reasoning, thinking

#### Exclude Keywords (Non-Features)
- fix, bug, hotfix
- ci, cd, pipeline, chore
- docs, readme, refactor, style, typo
- test, bump, deps, dependency
- merge, revert, format, lint, build
- release, prepare, coverage, security

### Integration with README

To update the README with new release features:

```bash
# Generate the features and save to a temporary file
python scripts/readme_updates/update_readme_features.py --limit 20 > /tmp/recent_features.md

# Manually replace the "### Recent Major Features" section in README.md
# with the generated content
```

### Database Schema

The script expects the following SQLite table structure:

```sql
CREATE TABLE versions (
    name TEXT PRIMARY KEY,
    date DATETIME,
    commit_sha TEXT,
    pr_numbers TEXT,
    ai_summary TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### Date Format Support

The script can parse various date formats:
- ISO 8601 with timezone: `2025-08-14 14:11:04+00:00`
- ISO 8601 basic: `2025-08-14T14:11:04`
- Date only: `2025-08-14`
- US format: `08/14/2025`

Output format is standardized to: `Aug 14, 2025`

### Maintenance Notes

- **AI Summary Format Changes**: If the format of AI summaries changes, update the `extract_title_desc()` and `split_summary()` functions
- **Keyword Tuning**: Adjust `INCLUDE_RE` and `EXCLUDE_RE` patterns as needed
- **Title Extraction**: The script attempts to extract concise titles from feature descriptions
- **Description Length**: Descriptions are limited to 200 characters for readability

### Future Enhancements

Potential improvements for automated README updates:
- Add section delimiter markers in README for automated replacement
- Create a GitHub Action to run on new releases
- Add support for categorizing features by type
- Implement confidence scoring for feature detection
