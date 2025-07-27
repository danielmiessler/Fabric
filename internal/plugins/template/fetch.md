# Fetch Plugin Tests

Simple test file for validating fetch plugin functionality.

## Basic Fetch Operations

```
Raw Content:
{{plugin:fetch:get:https://raw.githubusercontent.com/user/repo/main/README.md}}

JSON API:
{{plugin:fetch:get:https://api.example.com/data.json}}
```

## Error Cases
These should produce appropriate error messages:

```
Invalid Operation:
{{plugin:fetch:invalid:https://example.com}}

Invalid URL:
{{plugin:fetch:get:not-a-url}}

Non-text Content:
{{plugin:fetch:get:https://example.com/image.jpg}}

Server Error:
{{plugin:fetch:get:https://httpstat.us/500}}
```

## Security Considerations

- Only use trusted URLs
- Be aware of rate limits
- Content is limited to 1MB
- Only text content types are allowed
- Consider URL allow listing in production
- Validate and sanitize fetched content before use