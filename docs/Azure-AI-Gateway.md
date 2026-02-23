# Azure AI Gateway Plugin

The Azure AI Gateway plugin enables Fabric to access multiple AI providers through a single Azure API Management (APIM) Gateway endpoint. This allows organizations using Azure APIM as a central AI gateway to leverage Fabric with any supported backend provider using a single subscription key.

## Overview

Azure AI Gateway acts as a unified proxy that routes requests to different AI providers:

- **AWS Bedrock** - Claude models via Bedrock inference profiles
- **Azure OpenAI** - GPT-4o, GPT-4 Turbo, o1, DeepSeek-R1 models
- **Google Vertex AI** - Gemini model family

All backends share the same authentication mechanism (Azure APIM subscription key) and gateway endpoint, simplifying credential management and access control.

## Prerequisites

1. **Azure APIM Gateway** - A configured Azure API Management instance
2. **Gateway Subscription Key** - APIM subscription key with access to AI backends
3. **Backend Access** - Your APIM gateway must be configured to proxy to at least one of:
   - AWS Bedrock (requires AWS credentials configured in APIM)
   - Azure OpenAI (requires Azure OpenAI deployment)
   - Google Vertex AI (requires GCP credentials configured in APIM)

## Configuration

Run `fabric --setup` and select `AzureAIGateway` from the vendor list.

You'll be prompted for:

### Required Fields

1. **Backend Type** (`backend`)
   - Options: `bedrock`, `azure-openai`, `vertex-ai`
   - Default: `bedrock`
   - Choose based on which AI provider your APIM gateway is configured to access

2. **Gateway URL** (`gateway_url`)
   - Your Azure APIM Gateway base URL
   - Example: `https://gateway.company.com`
   - Must use HTTPS

3. **Subscription Key** (`subscription_key`)
   - Your Azure APIM subscription key
   - Used for authentication to the gateway

### Optional Fields

4. **API Version** (`api_version`) - **Azure OpenAI backend only**
   - Azure OpenAI API version to use
   - Default: `2025-04-01-preview`
   - Leave empty to use the default
   - Custom versions: `2024-08-01-preview`, `2024-10-21`, etc.
   - See [Azure OpenAI API Reference](https://learn.microsoft.com/azure/ai-services/openai/reference)

## Backend-Specific Configuration

### AWS Bedrock

**Authentication:** `Authorization: Bearer <subscription-key>`

**API Format:** Anthropic Messages API

**Models:**
```bash
fabric --listmodels
# Returns Claude models available via Bedrock:
# - us.anthropic.claude-3-5-sonnet-20241022-v2:0
# - us.anthropic.claude-3-5-haiku-20241022-v1:0
# - us.anthropic.claude-3-opus-20240229-v1:0
# - etc.
```

**Endpoint Pattern:** `/model/{model-id}/invoke`

**Configuration:**
```bash
fabric --setup
# Select: AzureAIGateway
# Backend: bedrock
# Gateway URL: https://gateway.company.com
# Subscription Key: your-apim-key
```

### Azure OpenAI

**Authentication:** `api-key: <subscription-key>`

**API Format:** OpenAI Chat Completions API

**Models:**
```bash
fabric --listmodels
# Returns Azure OpenAI deployment names:
# - gpt-4o
# - gpt-4o-mini
# - gpt-4-turbo
# - gpt-35-turbo
# - o1
# - o1-mini
# - DeepSeek-R1
```

**Endpoint Pattern:** `/openai/deployments/{deployment-name}/chat/completions?api-version={version}`

**Configuration:**
```bash
fabric --setup
# Select: AzureAIGateway
# Backend: azure-openai
# Gateway URL: https://gateway.company.com
# Subscription Key: your-apim-key
# API Version: (press Enter for default 2025-04-01-preview)
```

**Custom API Version:**
```bash
# During setup, specify a custom version:
# API Version: 2024-10-21
```

### Google Vertex AI

**Authentication:** `x-goog-api-key: <subscription-key>`

**API Format:** Gemini API

**Models:**
```bash
fabric --listmodels
# Returns Gemini models:
# - gemini-2.0-flash-exp
# - gemini-1.5-pro
# - gemini-1.5-flash
# - gemini-pro
# - gemini-pro-vision
```

**Endpoint Pattern:** `/publishers/google/models/{model-id}:generateContent`

**Note:** The endpoint path differs from direct Vertex AI API (`/v1beta/models/...`) because Azure APIM Gateway uses publisher-based routing.

**Configuration:**
```bash
fabric --setup
# Select: AzureAIGateway
# Backend: vertex-ai
# Gateway URL: https://gateway.company.com
# Subscription Key: your-apim-key
```

## Usage Examples

### Basic Usage

```bash
# Bedrock (Claude)
echo "Explain quantum computing" | fabric --model us.anthropic.claude-3-5-sonnet-20241022-v2:0 --pattern explain

# Azure OpenAI (GPT-4o)
echo "Explain quantum computing" | fabric --model gpt-4o --pattern explain

# Vertex AI (Gemini)
echo "Explain quantum computing" | fabric --model gemini-2.0-flash-exp --pattern explain
```

### Using Patterns

```bash
# Extract wisdom from a YouTube video (Bedrock)
fabric --youtube "https://youtube.com/watch?v=example" --model us.anthropic.claude-3-5-sonnet-20241022-v2:0 --pattern extract_wisdom

# Summarize an article (Azure OpenAI)
curl -s https://example.com/article | fabric --model gpt-4o --pattern summarize

# Create content from a prompt (Vertex AI)
fabric --model gemini-1.5-pro --pattern write_essay
```

### Switching Between Backends

```bash
# Reconfigure to use a different backend
fabric --setup
# Select: AzureAIGateway
# Change Backend from bedrock to azure-openai
```

## Troubleshooting

### Authentication Errors (401 Unauthorized)

**Symptom:** `Request failed with status 401`

**Causes:**
- Invalid subscription key
- Subscription key doesn't have access to the selected backend
- APIM subscription expired or disabled

**Solutions:**
1. Verify subscription key in Azure Portal → APIM → Subscriptions
2. Check subscription scope includes your gateway API
3. Regenerate key if compromised: `fabric --setup` and enter new key

### Model Not Found (404)

**Symptom:** `Request failed with status 404`

**Causes:**
- **Bedrock:** Model ID doesn't match inference profile names
- **Azure OpenAI:** Deployment name doesn't exist in your Azure OpenAI resource
- **Vertex AI:** Model not available in your region

**Solutions:**
1. List available models: `fabric --listmodels`
2. **Azure OpenAI:** Verify deployment exists in Azure Portal
3. **Bedrock:** Use full inference profile IDs (e.g., `us.anthropic.claude-3-5-sonnet-20241022-v2:0`)

### Connection Errors

**Symptom:** `connection refused` or timeout errors

**Causes:**
- Gateway URL incorrect or unreachable
- Network/firewall blocking access
- APIM gateway down

**Solutions:**
1. Verify gateway URL is accessible: `curl -I https://gateway.company.com`
2. Check APIM gateway health in Azure Portal
3. Verify HTTPS scheme (HTTP is rejected)

### API Version Errors (Azure OpenAI)

**Symptom:** `API version not supported` or `400 Bad Request`

**Causes:**
- Custom API version not supported by your APIM gateway
- Version mismatch between Azure OpenAI deployment and API version

**Solutions:**
1. Use default version: `fabric --setup` and leave API version empty
2. Check supported versions: [Azure OpenAI API Versions](https://learn.microsoft.com/azure/ai-services/openai/reference)
3. Update APIM gateway to support newer API versions

### Backend Mismatch

**Symptom:** Unexpected response format or empty responses

**Causes:**
- Selected backend doesn't match APIM gateway configuration
- Using wrong model names for the backend

**Solutions:**
1. Reconfigure: `fabric --setup` and select correct backend type
2. **Bedrock:** Use Bedrock inference profile IDs
3. **Azure OpenAI:** Use deployment names from your Azure OpenAI resource
4. **Vertex AI:** Use Gemini model IDs

### Request Timeout

**Symptom:** Request times out after 5 minutes

**Causes:**
- Model inference taking too long
- APIM gateway timeout settings
- Network latency

**Solutions:**
1. Use faster models (e.g., Claude 3.5 Haiku, gpt-4o-mini, gemini-1.5-flash)
2. Reduce input size or complexity
3. Check APIM gateway timeout policies

## Limitations

### No Streaming Support

Azure APIM Gateway doesn't support Server-Sent Events (SSE) pass-through for streaming responses. The plugin automatically falls back to buffered responses.

**Impact:** `--stream` flag is ignored; full response is returned after model completes.

**Workaround:** None - this is an APIM Gateway architectural limitation.

### Request Timeout

Maximum request timeout is 300 seconds (5 minutes). Long-running model inference may timeout.

**Solutions:**
- Use faster models
- Reduce prompt complexity
- Configure APIM gateway timeout policies (requires gateway admin access)

### Model Name Requirements

Model names must exactly match backend expectations:

- **Bedrock:** Full inference profile IDs (e.g., `us.anthropic.claude-3-5-sonnet-20241022-v2:0`)
- **Azure OpenAI:** Exact deployment names configured in your Azure OpenAI resource
- **Vertex AI:** Gemini model IDs (e.g., `gemini-2.0-flash-exp`)

Use `fabric --listmodels` to see available options for your configured backend.

## Security Considerations

### Subscription Key Protection

- **Never commit** subscription keys to version control
- Use Fabric's secure configuration storage (keys stored in `~/.config/fabric/.env`)
- Rotate keys regularly via Azure Portal

### HTTPS Enforcement

The plugin rejects HTTP gateway URLs to prevent plaintext credential transmission. Your gateway URL must use HTTPS.

### Response Size Limits

Responses are limited to 10MB to prevent memory exhaustion attacks. This is sufficient for all normal AI model responses.

## Advanced Configuration

### Multiple Gateway Configurations

To use multiple APIM gateways or backends, run `fabric --setup` and reconfigure when switching contexts.

### Environment Variables

Fabric configuration is stored in `~/.config/fabric/.env`. Manual editing is supported but not recommended:

```bash
# Example configuration
AZUREAIGATEWAY_BACKEND=bedrock
AZUREAIGATEWAY_GATEWAY_URL=https://gateway.company.com
AZUREAIGATEWAY_SUBSCRIPTION_KEY=your-key-here
AZUREAIGATEWAY_API_VERSION=2025-04-01-preview
```

## Support

For issues specific to the Azure AI Gateway plugin:
1. Check this documentation first
2. Verify APIM gateway configuration in Azure Portal
3. Test direct APIM gateway access with curl
4. File issues at: https://github.com/danielmiessler/fabric/issues

For APIM gateway configuration issues, consult:
- [Azure APIM GenAI Gateway Capabilities](https://learn.microsoft.com/azure/api-management/genai-gateway-capabilities)
- [Azure API Management Documentation](https://learn.microsoft.com/azure/api-management/)
