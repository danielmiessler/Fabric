# Azure OpenAI Troubleshooting

This document describes a known issue with Azure OpenAI integration and its fix.

## Issue: DeploymentNotFound Error (404)

### Symptoms

When using Fabric with Azure OpenAI, you may encounter this error:

```
POST "https://{resource}.cognitiveservices.azure.com/openai/chat/completions?api-version=...": 404 Not Found
{
  "code": "DeploymentNotFound",
  "message": "The API deployment for this resource does not exist..."
}
```

### Root Cause

Azure OpenAI requires the deployment name in the URL path for Chat Completions:

```
✅ Correct:  /openai/deployments/{deployment-name}/chat/completions
❌ Incorrect: /openai/chat/completions
```

Older versions of the OpenAI Go SDK did not add the deployment name to the path. Fabric used custom middleware to change the path.

## Fix

Fabric now uses OpenAI Go SDK v3. The SDK `azure.WithEndpoint()` option does the routing, and Fabric does not use custom middleware:

- **Chat Completions** and other deployment-scoped calls go to `/openai/deployments/{deployment-name}/...`. The SDK reads the deployment name from the `model` field.
- **Responses** calls go to the resource-scoped path `/openai/responses`. The deployment name stays in the `model` field of the request body.

### Endpoint Requirements

The SDK v3 has these transport rules for Azure credentials:

- Use an HTTPS endpoint. The SDK does not send Azure credentials over plain HTTP.
- The SDK does not follow a redirect to a different origin with Azure credentials. If you use a gateway or a proxy, set `AZURE_API_BASE_URL` to the final trusted origin. Do not set it to a URL that redirects to a different host.

## Additional Fix: StreamOptions Error

### Symptom

```
400 Bad Request
{
  "message": "The 'stream_options' parameter is only allowed when 'stream' is enabled."
}
```

### Cause

The Chat Completions API was sending `stream_options` for all requests, but Azure only accepts this parameter when `stream: true` is also set.

### Fix

Moved `StreamOptions` to only be set for streaming requests in `internal/plugins/ai/openai/chat_completions.go`.

## Configuration

Ensure your Azure OpenAI configuration is correct:

```bash
# In ~/.config/fabric/.env
AZURE_API_KEY=your-api-key
AZURE_API_BASE_URL=https://{your-resource}.cognitiveservices.azure.com/  # Must use HTTPS
AZURE_DEPLOYMENTS=your-deployment-1,your-deployment-2  # Comma-separated deployment names
AZURE_API_VERSION=2025-04-01-preview  # Optional, defaults to 2025-04-01-preview
```

**Note:** The deployment name is what you specified when deploying a model in Azure AI Foundry (formerly Azure OpenAI Studio), not the model name itself (e.g., `my-gpt4-deployment` rather than `gpt-4`).

## Verification

Test your Azure OpenAI setup:

```bash
fabric --model <your-deployment-name> --pattern summarize "Hello world"
```

Replace `<your-deployment-name>` with the actual deployment name from your Azure configuration.

You should see a successful response from your Azure OpenAI deployment.

## References

- GitHub Issue: [#1954](https://github.com/danielmiessler/fabric/issues/1954)
- Pull Request: [#1965](https://github.com/danielmiessler/fabric/pull/1965)
- [Azure OpenAI REST API Reference](https://learn.microsoft.com/en-us/azure/ai-services/openai/reference)
