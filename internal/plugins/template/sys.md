# System Plugin Tests

Simple test file for validating system plugin functionality.

## Basic System Information

```
Hostname: {{plugin:sys:hostname}}
Username: {{plugin:sys:user}}
Operating System: {{plugin:sys:os}}
Architecture: {{plugin:sys:arch}}
```

## Paths and Directories

```
Current Directory: {{plugin:sys:pwd}}
Home Directory: {{plugin:sys:home}}
```

## Environment Variables

```
Path: {{plugin:sys:env:PATH}}
Home: {{plugin:sys:env:HOME}}
Shell: {{plugin:sys:env:SHELL}}
```

## Error Cases
These should produce appropriate error messages:

```
Invalid Operation: {{plugin:sys:invalid}}
Missing Env Var: {{plugin:sys:env:}}
Non-existent Env Var: {{plugin:sys:env:NONEXISTENT_VAR_123456}}
```

## Security Note

Be careful when exposing system information in templates, especially:
- Environment variables that might contain sensitive data
- Full paths that reveal system structure
- Username/hostname information in public templates