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

Non-public Address:
{{plugin:fetch:get:http://127.0.0.1:8080/status}}
```

## Security Considerations

- Only use trusted URLs
- Be aware of rate limits
- Content is limited to 1MB
- Only text content types are allowed
- The plugin connects only to public IP addresses. It refuses loopback, private, shared (`100.64.0.0/10`), link-local, multicast and unspecified addresses. It checks the address after DNS resolution, and again for each redirect
- A fetch stops after 30 seconds or after 5 redirects
- The plugin does not use `HTTP_PROXY` or `HTTPS_PROXY`
- Consider URL allow listing in production
- Validate and sanitize fetched content before use