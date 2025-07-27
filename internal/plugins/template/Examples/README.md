
# Fabric Extensions: Complete Guide

## Understanding Extension Architecture

### Registry Structure
The extension registry is stored at `~/.config/fabric/extensions/extensions.yaml` and tracks registered extensions:

```yaml
extensions:
    extension-name:
        config_path: /path/to/config.yaml
        config_hash: <sha256>
        executable_hash: <sha256>
```

The registry maintains security through hash verification of both configs and executables.

### Extension Configuration
Each extension requires a YAML configuration file with the following structure:

```yaml
name: "extension-name"          # Unique identifier
executable: "/path/to/binary"   # Full path to executable
type: "executable"             # Type of extension
timeout: "30s"                 # Execution timeout
description: "Description"     # What the extension does
version: "1.0.0"              # Version number
env: []                       # Optional environment variables

operations:                   # Defined operations
  operation-name:
    cmd_template: "{{executable}} {{operation}} {{value}}"

config:                      # Output configuration
  output:
    method: "stdout"         # or "file"
    file_config:            # Optional, for file output
      cleanup: true
      path_from_stdout: true
      work_dir: "/tmp"
```

### Directory Structure
Recommended organization:
```
~/.config/fabric/extensions/
├── bin/           # Extension executables
├── configs/       # Extension YAML configs
└── extensions.yaml # Registry file
```

## Example 1: Python Wrapper (Word Generator)
A simple example wrapping a Python script.

### 1. Position Files
```bash
# Create directories
mkdir -p ~/.config/fabric/extensions/{bin,configs}

# Install script
cp word-generator.py ~/.config/fabric/extensions/bin/
chmod +x ~/.config/fabric/extensions/bin/word-generator.py
```

### 2. Configure
Create `~/.config/fabric/extensions/configs/word-generator.yaml`:
```yaml
name: word-generator
executable: "~/.config/fabric/extensions/bin/word-generator.py"
type: executable
timeout: "5s"
description: "Generates random words based on count parameter"
version: "1.0.0"

operations:
  generate:
    cmd_template: "{{executable}} {{value}}"

config:
  output:
    method: stdout
```

### 3. Register & Run
```bash
# Register
fabric --addextension ~/.config/fabric/extensions/configs/word-generator.yaml

# Run (generate 3 random words)
echo "{{ext:word-generator:generate:3}}" | fabric
```

## Example 2: Direct Executable (SQLite3)
Using a system executable directly.

copy the memories to your home directory
 ~/memories.db

### 1. Configure
Create `~/.config/fabric/extensions/configs/memory-query.yaml`:
```yaml
name: memory-query
executable: "/usr/bin/sqlite3"
type: executable
timeout: "5s"
description: "Query memories database"
version: "1.0.0"

operations:
  goal:
    cmd_template: "{{executable}} -json ~/memories.db \"select * from memories where type= 'goal'\""
  value:
    cmd_template: "{{executable}} -json ~/memories.db \"select * from memories where type= 'value'\""
  byid:
    cmd_template: "{{executable}} -json ~/memories.db \"select * from memories where uid= {{value}}\""
  all:
    cmd_template: "{{executable}} -json ~/memories.db \"select * from memories\""

config:
  output:
    method: stdout
```

### 2. Register & Run
```bash
# Register
fabric --addextension ~/.config/fabric/extensions/configs/memory-query.yaml

# Run queries
echo  "{{ext:memory-query:all}}" | fabric
echo  "{{ext:memory-query:byid:3}}" | fabric
```


## Extension Management Commands

### Add Extension
```bash
fabric --addextension ~/.config/fabric/extensions/configs/memory-query.yaml
```

Note : if the executable or config file changes, you must re-add the extension.
This will recompute the hash for the extension.


### List Extensions
```bash
fabric --listextensions
```
Shows all registered extensions with their status and configuration details.

### Remove Extension
```bash
fabric --rmextension <extension-name>
```
Removes an extension from the registry.


## Extensions in patterns

```
Create a pattern that use multiple extensions.

These are my favorite
{{ext:word-generator:generate:3}}

These are my least favorite
{{ext:word-generator:generate:2}}

what does this say about me?
```

```bash
./fabric -p ./plugins/template/Examples/test_pattern.md
```

## Security Considerations

1. **Hash Verification**
   - Both configs and executables are verified via SHA-256 hashes
   - Changes to either require re-registration
   - Prevents tampering with registered extensions

2. **Execution Safety**
   - Extensions run with user permissions
   - Timeout constraints prevent runaway processes
   - Environment variables can be controlled via config

3. **Best Practices**
   - Review extension code before installation
   - Keep executables in protected directories
   - Use absolute paths in configurations
   - Implement proper error handling in scripts
   - Regular security audits of registered extensions

## Troubleshooting

### Common Issues
1. **Registration Failures**
   - Verify file permissions
   - Check executable paths
   - Validate YAML syntax

2. **Execution Errors**
   - Check operation exists in config
   - Verify timeout settings
   - Monitor system resources
   - Check extension logs

3. **Output Issues**
   - Verify output method configuration
   - Check file permissions for file output
   - Monitor disk space for file operations

### Debug Tips
1. Enable verbose logging when available
2. Check system logs for execution errors
3. Verify extension dependencies
4. Test extensions with minimal configurations first


Would you like me to expand on any particular section or add more examples?