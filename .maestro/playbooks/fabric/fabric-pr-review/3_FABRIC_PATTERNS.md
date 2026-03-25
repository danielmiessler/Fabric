# Document 3: Fabric Patterns Validation

## Context

- **Playbook**: Fabric PR Review
- **Agent**: Fabric-PR-Review
- **Project**: /Users/kayvan/src/fabric
- **Date**: 2026-03-25
- **Working Folder**: /Users/kayvan/src/fabric/.maestro/playbooks

## Purpose

Validate that new or modified patterns follow Fabric's pattern system conventions.

## Prerequisites

- `/Users/kayvan/src/fabric/.maestro/playbooks/REVIEW_SCOPE.md` exists from Document 1

## Tasks

### Task 1: Load Context

- [x] **Read scope**: Loaded `/Users/kayvan/src/fabric/.maestro/playbooks/REVIEW_SCOPE.md`; it reports no files changed in `data/patterns/` or `data/strategies/` and marks pattern validation as not needed for PR 2063.

- [ ] **Check if patterns changed**: If no files in `data/patterns/` or `data/strategies/` were modified, skip to Task 7 and note "No pattern changes in this PR."

### Task 2: Validate Pattern Structure

For each new or modified pattern directory in `data/patterns/`:

- [ ] **Check required files**:
  - `system.md` must exist (the main prompt)
  - `user.md` is optional (user prompt section)
  - No other unexpected files

- [ ] **Verify directory naming**:
  - Lowercase with underscores
  - Descriptive of the pattern's purpose
  - No spaces or special characters

### Task 3: Validate Pattern Content

- [ ] **Check system.md structure**:
  - Uses Markdown formatting for readability
  - Has clear sections/headings
  - Instructions are explicit
  - No ambiguous directives

- [ ] **Verify variable syntax**:
  - Variables use `{{.variable}}` Go template syntax
  - No invalid template syntax
  - Variables are documented if used
  - Common variables: `{{.input}}`, `{{.role}}`, `{{.points}}`

- [ ] **Check for hardcoded values**:
  - No API keys or secrets
  - No user-specific paths
  - No hardcoded model names (should be configurable)

### Task 4: Validate Pattern Quality

- [ ] **Prompt engineering best practices**:
  - Clear, specific instructions
  - Output format is defined
  - Edge cases considered
  - Appropriate for multiple LLM providers

- [ ] **Content quality**:
  - No typos or grammar issues
  - Professional tone
  - Consistent with existing patterns
  - Appropriate length (not too verbose)

### Task 5: Validate Strategy Changes

For changes to `data/strategies/`:

- [ ] **Check JSON structure**:
  - Valid JSON format
  - Required fields present
  - Strategy type is valid (CoT, ToT, etc.)

- [ ] **Verify strategy prompt**:
  - Modifies system prompt appropriately
  - Clear reasoning instructions
  - Compatible with various patterns

### Task 6: Test Pattern Loading

- [ ] **Verify pattern loads**: Test that the pattern can be listed:
  ```bash
  ./fabric --listpatterns | grep pattern_name
  ```

- [ ] **Check variable substitution**: If pattern uses variables, test:
  ```bash
  echo "test" | ./fabric --dry-run --pattern pattern_name -v=#var:value
  ```

### Task 7: Document Pattern Issues

- [ ] **Create PATTERN_ISSUES.md**: Write findings to `/Users/kayvan/src/fabric/.maestro/playbooks/PATTERN_ISSUES.md`:

```markdown
# Pattern Validation Results

## Patterns Reviewed
[List of patterns checked]

## Pattern Structure Issues
[Missing files, naming issues]

## Variable Syntax Issues
[Invalid template syntax, undocumented variables]

## Content Quality Issues
[Prompt engineering concerns, clarity issues]

## Strategy Issues
[JSON errors, invalid strategy types]

## Security Concerns
[Hardcoded values, potential secrets]

## Suggestions
[Pattern improvements, best practice recommendations]

## No Issues Found
[Patterns that passed all checks]

## Skipped
[Note if no patterns were modified in this PR]
```

For each issue include:
- Pattern name and file
- Issue description
- Suggested fix
- Severity: Critical / Major / Minor / Suggestion

## Success Criteria

- All modified patterns reviewed
- Structure validated
- Variable syntax verified
- Content quality checked
- Strategy changes validated
- PATTERN_ISSUES.md created

## Status

Mark complete when pattern review document is created.

---

**Next**: Document 4 will validate plugin architecture compliance.
