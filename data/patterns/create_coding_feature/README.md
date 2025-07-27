# Create Coding Feature

Generate code changes to an existing coding project using AI.

## Installation

After installing the `code_helper` binary:

```bash
go install github.com/danielmiessler/fabric/cmd/code_helper@latest
```

## Usage

The create_coding_feature allows you to apply AI-suggested code changes directly to your project files. Use it like this:

```bash
code_helper [project_directory] "[instructions for code changes]" | fabric --pattern create_coding_feature
```

For example:

```bash
code_helper . "Create a simple Hello World C program in file main.c" | fabric --pattern create_coding_feature
```

## How It Works

1. `code_helper` scans your project directory and creates a JSON representation
2. The AI model analyzes your project structure and instructions
3. AI generates file changes in a standard format
4. Fabric parses these changes and prompts you to confirm
5. If confirmed, changes are applied to your project files

## Example Workflow

```bash
# Request AI to create a Hello World program
code_helper . "Create a simple Hello World C program in file main.c" | fabric --pattern create_coding_feature

# Review the changes made to your project
git diff

# Run/test the code
make check

# If satisfied, commit the changes
git add <changed files>
git commit -s -m "Add Hello World program"
```

### Security Enhancement Example

```bash
code_helper . "Ensure that all user input is validated and sanitized before being used in the program." | fabric --pattern create_coding_feature
git diff
make check
git add <changed files>
git commit -s -m "Security fixes: Input validation"
```

## Important Notes

- **Always run from project root**: File changes are applied relative to your current directory
- **Use with version control**: It's highly recommended to use this feature in a clean git repository so you can review and revert
  changes. You will *not* be asked to approve each change.

## Security Features

- Path validation to prevent directory traversal attempts
- File size limits to prevent excessive file generation
- Operation validation (only create/update operations allowed)
- User confirmation required before applying changes

## Suggestions for Future Improvements

- Add a dry-run mode to show changes without applying them
- Enhance reporting with detailed change summaries
- Support for file deletions with safety checks
- Add configuration options for project-specific rules
- Provide rollback capability for applied changes
- Add support for project-specific validation rules
- Enhance script generation with conditional logic
- Include detailed logging for API responses
- Consider adding a GUI for ease of use
