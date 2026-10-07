# File Plugin Tests

Simple test file for validating file plugin functionality.

Each path is relative to the working directory of the fabric process.

## Basic File Operations

```
Read File:
{{plugin:file:read:docs/file.txt}}

Last 5 Lines:
{{plugin:file:tail:logs/app.log|5}}

Check Existence:
{{plugin:file:exists:docs/file.txt}}

Get Size:
{{plugin:file:size:docs/file.txt}}

Last Modified:
{{plugin:file:modified:docs/file.txt}}
```

## Error Cases
These should produce appropriate error messages:

```
Invalid Operation:
{{plugin:file:invalid:docs/file.txt}}

Non-existent File:
{{plugin:file:read:docs/nonexistent.txt}}

Path Traversal Attempt:
{{plugin:file:read:../../../etc/passwd}}

Absolute Path:
{{plugin:file:read:/path/to/file.txt}}

Home Directory Path:
{{plugin:file:read:~/file.txt}}

Invalid Tail Format:
{{plugin:file:tail:docs/file.txt}}

Large File:
{{plugin:file:read:data/huge.iso}}
```

## Security Considerations

- The plugin rejects an absolute path, a path that starts with `~`, and a path that contains `..`
- A symbolic link must be relative and must stay in the working directory
- The plugin can read all files in the working directory and its subdirectories. Run `fabric` in a directory that has only the files that patterns can read
- The REST server (`fabric --serve`) does not run the file plugin
- Be aware of file size limits (1MB max)
- All paths are cleaned and normalized
