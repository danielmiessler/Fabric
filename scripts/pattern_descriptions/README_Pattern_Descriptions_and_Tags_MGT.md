# Pattern Descriptions and Tags Management

This document explains the complete workflow for managing pattern descriptions and tags, including how to process new patterns and maintain metadata.

## System Overview

The pattern system follows this hierarchy:

1. `~/.config/fabric/patterns/` directory: The source of truth for available patterns
2. `pattern_extracts.json`: Contains first 500 words of each pattern for reference
3. `pattern_descriptions.json`: Stores pattern metadata (descriptions and tags)
4. `web/static/data/pattern_descriptions.json`: Web-accessible copy for the interface

## Pattern Processing Workflow

### 1. Adding New Patterns

- Add patterns to `~/.config/fabric/patterns/`
- Run extract_patterns.py to process new additions:

  ```bash
  python extract_patterns.py

The Python Script automatically:

- Creates pattern extracts for reference
- Adds placeholder entries in descriptions file
- Syncs to web interface

### 2. Pattern Extract Creation

The script extracts first 500 words from each pattern's system.md file to:

- Provide context for writing descriptions
- Maintain reference material
- Aid in pattern categorization

### 3. Description and Tag Management

Pattern descriptions and tags are managed in pattern_descriptions.json:

{
  "patterns": [
    {
      "patternName": "pattern_name",
      "description": "[Description pending]",
      "tags": []
    }
  ]
}

## Completing Pattern Metadata

### Writing Descriptions

1. Check pattern_descriptions.json for "[Description pending]" entries
2. Reference pattern_extracts.json for context

3. How to update Pattern short descriptions (one sentence).

You can update your descriptions in pattern_descriptions.json manually or using LLM assistance (preferred approach).

Tell AI to look for "Description pending" entries in this file and write a short description based on the extract info in the pattern_extracts.json file. You can also ask your LLM to add tags for those newly added patterns, using other patterns tag assignments as example.

### Managing Tags

1. Add appropriate tags to new patterns
2. Update existing tags as needed
3. Tags are stored as arrays: ["TAG1", "TAG2"]
4. Edit pattern_descriptions.json directly to modify tags
5. Make tags your own. You can delete, replace, amend existing tags.

## File Synchronization

The script maintains synchronization between:

- Local pattern_descriptions.json
- Web interface copy in static/data/
- No manual file copying needed

## Best Practices

1. Run extract_patterns.py when:
   - Adding new patterns
   - Updating existing patterns
   - Modifying pattern structure

2. Description Writing:
   - Use pattern extracts for context
   - Keep descriptions clear and concise
   - Focus on pattern purpose and usage

3. Tag Management:
   - Use consistent tag categories
   - Apply multiple tags when relevant
   - Update tags to reflect pattern evolution

## Troubleshooting

If patterns are not showing in the web interface:

1. Verify pattern_descriptions.json format
2. Check web static copy exists
3. Ensure proper file permissions
4. Run extract_patterns.py to resync

## File Structure

fabric/
├── patterns/                     # Pattern source files
├── PATTERN_DESCRIPTIONS/
│   ├── extract_patterns.py      # Pattern processing script
│   ├── pattern_extracts.json    # Pattern content references
│   └── pattern_descriptions.json # Pattern metadata
└── web/
    └── static/
        └── data/
            └── pattern_descriptions.json # Web interface copy
