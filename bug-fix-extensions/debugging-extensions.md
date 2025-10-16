# Debugging Fabric Extensions

> **Note:** For information about the `{{input}}` in extensions bug fix, see [input-extension-bug-fix.md](./input-extension-bug-fix.md).

## Method 1: Built-in Debug Logging

Use the `--debug` flag with levels 0-3:

```bash
# Basic debugging
echo "{{ext:word-generator:generate:3}}" | fabric --debug=1

# Detailed debugging  
echo "{{ext:word-generator:generate:3}}" | fabric --debug=2

# Trace level (most verbose)
echo "{{ext:word-generator:generate:3}}" | fabric --debug=3
```

## Method 2: Go Debugger (Delve)

```bash
# Build with debug symbols
cd /Users/ourdecisions/devHelp/Fabric-fix
go build -gcflags="all=-N -l" -o fabric-debug ./cmd/fabric

# Run with debugger
echo "{{ext:word-generator:generate:3}}" | dlv exec ./fabric-debug --
```

In the debugger:
```
(dlv) break internal/plugins/template/template.go:122
(dlv) continue
(dlv) print name
(dlv) print operation  
(dlv) print value
(dlv) step
```

## Method 3: Add Custom Debug Logging

To add more logging to the extension manager, edit:

`internal/plugins/template/extension_manager.go`:
```go
func (em *ExtensionManager) ProcessExtension(name, operation, value string) (string, error) {
    fmt.Printf("DEBUG: ProcessExtension called with name=%s, operation=%s, value=%s\n", name, operation, value)
    result, err := em.executor.Execute(name, operation, value)
    if err != nil {
        fmt.Printf("DEBUG: Extension execution failed: %v\n", err)
    } else {
        fmt.Printf("DEBUG: Extension result: %s\n", result)
    }
    return result, err
}
```

`internal/plugins/template/extension_executor.go`:
```go
func (e *ExtensionExecutor) Execute(name, operation, value string) (string, error) {
    fmt.Printf("DEBUG: ExtensionExecutor.Execute called\n")
    
    // Get extension definition
    ext, err := e.registry.GetExtension(name)
    if err != nil {
        fmt.Printf("DEBUG: Failed to get extension %s: %v\n", name, err)
        return "", fmt.Errorf("extension %s not found: %w", name, err)
    }
    
    fmt.Printf("DEBUG: Found extension %s, executable: %s\n", name, ext.Executable)
    
    // Rest of method...
}
```

## Method 4: Check Extension Status

First verify your extension is properly registered:

```bash
fabric --listextensions
```

## Method 5: Test Extension Directly

Test the Python script directly:

```bash
/Users/ourdecisions/.config/fabric/extensions/bin/word-generator.py generate 3
```

## Common Issues to Check

1. **Extension not registered**: Use `fabric --listextensions` to verify
2. **Hash mismatch**: Re-register with `fabric --addextension`
3. **Python script permissions**: Ensure executable with `chmod +x`
4. **Python script path**: Check if Python interpreter is correct
5. **Template syntax**: Verify `{{ext:name:operation:value}}` format

## Debug Output Example

With `--debug=3`, you should see output like:

```
DEBUG: Starting template processing
DEBUG: 
Extension call:
  Name: word-generator
  Operation: generate  
  Value: 3
DEBUG: Extension result: word1 word2 word3
```