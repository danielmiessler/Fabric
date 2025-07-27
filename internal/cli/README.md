# YAML Configuration Support

## Overview

Fabric now supports YAML configuration files for commonly used options. This allows users to persist settings and share configurations across multiple runs.

## Usage

Use the `--config` flag to specify a YAML configuration file:

```bash
fabric --config ~/.config/fabric/config.yaml "Tell me about APIs"
```

## Configuration Precedence

1. CLI flags (highest priority)
2. YAML config values
3. Default values (lowest priority)

## Supported Configuration Options

```yaml
# Model selection
model: gpt-4
modelContextLength: 4096

# Model parameters
temperature: 0.7
topp: 0.9
presencepenalty: 0.0
frequencypenalty: 0.0
seed: 42

# Pattern selection
pattern: analyze  # Use pattern name or filename

# Feature flags
stream: true
raw: false
```

## Rules and Behavior

- Only long flag names are supported in YAML (e.g., `temperature` not `-t`)
- CLI flags always override YAML values
- Unknown YAML declarations are ignored
- If a declaration appears multiple times in YAML, the last one wins
- The order of YAML declarations doesn't matter

## Type Conversions

The following string-to-type conversions are supported:

- String to number: `"42"` → `42`
- String to float: `"42.5"` → `42.5`
- String to boolean: `"true"` → `true`

## Example Config

```yaml
# ~/.config/fabric/config.yaml
model: gpt-4
temperature: 0.8
pattern: analyze
stream: true
topp: 0.95
presencepenalty: 0.1
frequencypenalty: 0.2
```

## CLI Override Example

```bash
# Override temperature from config
fabric --config ~/.config/fabric/config.yaml --temperature 0.9 "Query"
```
