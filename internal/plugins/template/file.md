# File Plugin Tests

Simple test file for validating file plugin functionality.

## Basic File Operations

```
Read File:
{{plugin:file:read:/path/to/file.txt}}

Last 5 Lines:
{{plugin:file:tail:/path/to/log.txt|5}}

Check Existence:
{{plugin:file:exists:/path/to/file.txt}}

Get Size:
{{plugin:file:size:/path/to/file.txt}}

Last Modified:
{{plugin:file:modified:/path/to/file.txt}}
```

## Error Cases
These should produce appropriate error messages:

```
Invalid Operation:
{{plugin:file:invalid:/path/to/file.txt}}

Non-existent File:
{{plugin:file:read:/path/to/nonexistent.txt}}

Path Traversal Attempt:
{{plugin:file:read:../../../etc/passwd}}

Invalid Tail Format:
{{plugin:file:tail:/path/to/file.txt}}

Large File:
{{plugin:file:read:/path/to/huge.iso}}
```

## Security Considerations

- Carefully control which paths are accessible
- Consider using path allow lists in production
- Be aware of file size limits (1MB max)
- No directory traversal is allowed
- Home directory (~/) expansion is supported
- All paths are cleaned and normalized