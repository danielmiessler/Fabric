# Fabric Template System

## Quick Start
echo "Hello {{name}}!" | fabric -v=name:World

## Overview

The Fabric Template System provides a powerful and extensible way to handle variable substitution and dynamic content generation through a plugin architecture. It uses a double-brace syntax (`{{}}`) for variables and plugin operations, making it both readable and flexible.



## Basic Usage

### Variable Substitution

The template system supports basic variable substitution using double braces:

```markdown
Hello {{name}}!
Current role: {{role}}
```

Variables can be provided via:
- Command line arguments: `-v=name:John -v=role:admin`
- YAML front matter in input files
- Environment variables (when configured)

### Special Variables

- `{{input}}`: Represents the main input content
  ```markdown
  Here is the analysis:
  {{input}}
  End of analysis.
  ```

## Nested Tokens and Resolution

### Basic Nesting

The template system supports nested tokens, where inner tokens are resolved before outer ones. This enables complex, dynamic template generation.

#### Simple Variable Nesting
```markdown
{{outer{{inner}}}}

Example:
Variables: {
  "inner": "name",
  "john": "John Doe"
}
{{{{inner}}}} -> {{name}} -> John Doe
```

#### Nested Plugin Calls
```markdown
{{plugin:text:upper:{{plugin:sys:env:USER}}}}
First resolves: {{plugin:sys:env:USER}} -> "john"
Then resolves: {{plugin:text:upper:john}} -> "JOHN"
```

### How Nested Resolution Works

1. **Iterative Processing**
   - The engine processes the template in multiple passes
   - Each pass identifies all `{{...}}` patterns
   - Processing continues until no more replacements are needed

2. **Resolution Order**
   ```markdown
   Original: {{plugin:text:upper:{{user}}}}
   Step 1: Found {{user}} -> "john"
   Step 2: Now have {{plugin:text:upper:john}}
   Step 3: Final result -> "JOHN"
   ```

3. **Complex Nesting Example**
   ```markdown
   {{plugin:text:{{case}}:{{plugin:sys:env:{{varname}}}}}}
   
   With variables:
   {
     "case": "upper",
     "varname": "USER"
   }
   
   Resolution steps:
   1. {{varname}} -> "USER"
   2. {{plugin:sys:env:USER}} -> "john"
   3. {{case}} -> "upper"
   4. {{plugin:text:upper:john}} -> "JOHN"
   ```

### Important Considerations

1. **Depth Limitations**
   - While nesting is supported, avoid excessive nesting for clarity
   - Complex nested structures can be hard to debug
   - Consider breaking very complex templates into smaller parts

2. **Variable Resolution**
   - Inner variables must resolve to valid values for outer operations
   - Error messages will point to the innermost failed resolution
   - Debug logs show the step-by-step resolution process

3. **Plugin Nesting**
   ```markdown
   # Valid:
   {{plugin:text:upper:{{plugin:sys:env:USER}}}}
   
   # Also Valid:
   {{plugin:text:{{operation}}:{{value}}}}
   
   # Invalid (plugin namespace cannot be dynamic):
   {{plugin:{{namespace}}:operation:value}}
   ```

4. **Debugging Nested Templates**
   ```go
   Debug = true  // Enable debug logging
   
   Template: {{plugin:text:upper:{{user}}}}
   Debug output:
   > Processing variable: user
   > Replacing {{user}} with john
   > Plugin call:
   >   Namespace: text
   >   Operation: upper
   >   Value: john
   > Plugin result: JOHN
   ```

### Examples

1. **Dynamic Operation Selection**
   ```markdown
   {{plugin:text:{{operation}}:hello}}
   
   With variables:
   {
     "operation": "upper"
   }
   
   Result: HELLO
   ```

2. **Dynamic Environment Variable Lookup**
   ```markdown
   {{plugin:sys:env:{{env_var}}}}
   
   With variables:
   {
     "env_var": "HOME"
   }
   
   Result: /home/user
   ```

3. **Nested Date Formatting**
   ```markdown
   {{plugin:datetime:{{format}}:{{plugin:datetime:now}}}}
   
   With variables:
   {
     "format": "full"
   }
   
   Result: Wednesday, November 20, 2024
   ```





## Plugin System

### Plugin Syntax

Plugins use the following syntax:
```
{{plugin:namespace:operation:value}}
```

- `namespace`: The plugin category (e.g., text, datetime, sys)
- `operation`: The specific operation to perform
- `value`: Optional value for the operation

### Built-in Plugins

#### Text Plugin
Text manipulation operations:
```markdown
{{plugin:text:upper:hello}}  -> HELLO
{{plugin:text:lower:HELLO}}  -> hello
{{plugin:text:title:hello world}} -> Hello World
```

#### DateTime Plugin
Time and date operations:
```markdown
{{plugin:datetime:now}}       -> 2024-11-20T15:04:05Z
{{plugin:datetime:today}}     -> 2024-11-20
{{plugin:datetime:rel:-1d}}   -> 2024-11-19
{{plugin:datetime:month}}     -> November
```

#### System Plugin
System information:
```markdown
{{plugin:sys:hostname}}   -> server1
{{plugin:sys:user}}       -> currentuser
{{plugin:sys:os}}         -> linux
{{plugin:sys:env:HOME}}   -> /home/user
```

## Developing Plugins

### Plugin Interface

To create a new plugin, implement the following interface:

```go
type Plugin interface {
    Apply(operation string, value string) (string, error)
}
```

### Example Plugin Implementation

Here's a simple plugin that performs basic math operations:

```go
package template

type MathPlugin struct{}

func (p *MathPlugin) Apply(operation string, value string) (string, error) {
    switch operation {
    case "add":
        // Parse value as "a,b" and return a+b
        nums := strings.Split(value, ",")
        if len(nums) != 2 {
            return "", fmt.Errorf("add requires two numbers")
        }
        a, err := strconv.Atoi(nums[0])
        if err != nil {
            return "", err
        }
        b, err := strconv.Atoi(nums[1])
        if err != nil {
            return "", err
        }
        return fmt.Sprintf("%d", a+b), nil
    
    default:
        return "", fmt.Errorf("unknown math operation: %s", operation)
    }
}
```

### Registering a New Plugin

1. Add your plugin struct to the template package
2. Register it in template.go:

```go
var (
    // Existing plugins
    textPlugin = &TextPlugin{}
    datetimePlugin = &DateTimePlugin{}
    
    // Add your new plugin
    mathPlugin = &MathPlugin{}
)

// Update the plugin handler in ApplyTemplate
switch namespace {
    case "text":
        result, err = textPlugin.Apply(operation, value)
    case "datetime":
        result, err = datetimePlugin.Apply(operation, value)
    // Add your namespace
    case "math":
        result, err = mathPlugin.Apply(operation, value)
    default:
        return "", fmt.Errorf("unknown plugin namespace: %s", namespace)
}
```

### Plugin Development Guidelines

1. **Error Handling**
   - Return clear error messages
   - Validate all inputs
   - Handle edge cases gracefully

2. **Debugging**
   - Use the `debugf` function for logging
   - Log entry and exit points
   - Log intermediate calculations

```go
func (p *MyPlugin) Apply(operation string, value string) (string, error) {
    debugf("MyPlugin operation: %s value: %s\n", operation, value)
    // ... plugin logic ...
    debugf("MyPlugin result: %s\n", result)
    return result, nil
}
```

3. **Security Considerations**
   - Validate and sanitize inputs
   - Avoid shell execution
   - Be careful with file operations
   - Limit resource usage

4. **Performance**
   - Cache expensive computations
   - Minimize allocations
   - Consider concurrent access

### Testing Plugins

Create tests for your plugin in `plugin_test.go`:

```go
func TestMathPlugin(t *testing.T) {
    plugin := &MathPlugin{}
    
    tests := []struct {
        operation string
        value     string
        expected  string
        wantErr   bool
    }{
        {"add", "5,3", "8", false},
        {"add", "bad,input", "", true},
        {"unknown", "value", "", true},
    }
    
    for _, tt := range tests {
        result, err := plugin.Apply(tt.operation, tt.value)
        if (err != nil) != tt.wantErr {
            t.Errorf("MathPlugin.Apply(%s, %s) error = %v, wantErr %v",
                tt.operation, tt.value, err, tt.wantErr)
            continue
        }
        if result != tt.expected {
            t.Errorf("MathPlugin.Apply(%s, %s) = %v, want %v",
                tt.operation, tt.value, result, tt.expected)
        }
    }
}
```

## Best Practices

1. **Namespace Selection**
   - Choose clear, descriptive names
   - Avoid conflicts with existing plugins
   - Group related operations together

2. **Operation Names**
   - Use lowercase names
   - Keep names concise but clear
   - Be consistent with similar operations

3. **Value Format**
   - Document expected formats
   - Use common separators consistently
   - Provide examples in comments

4. **Error Messages**
   - Be specific about what went wrong
   - Include valid operation examples
   - Help users fix the problem

## Common Issues and Solutions

1. **Missing Variables**
   ```
   Error: missing required variables: [name]
   Solution: Provide all required variables using -v=name:value
   ```

2. **Invalid Plugin Operations**
   ```
   Error: unknown operation 'invalid' for plugin 'text'
   Solution: Check plugin documentation for supported operations
   ```

3. **Plugin Value Format**
   ```
   Error: invalid format for datetime:rel, expected -1d, -2w, etc.
   Solution: Follow the required format for plugin values
   ```




## Contributing

1. Fork the repository
2. Create your plugin branch
3. Implement your plugin following the guidelines
4. Add comprehensive tests
5. Submit a pull request

## Support

For issues and questions:
1. Check the debugging output (enable with Debug=true)
2. Review the plugin documentation
3. Open an issue with:
   - Template content
   - Variables used
   - Expected vs actual output
   - Debug logs