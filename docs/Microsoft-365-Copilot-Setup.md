# Microsoft 365 Copilot Setup Guide for Fabric

This guide walks you through setting up and using Microsoft 365 Copilot with Fabric CLI. Microsoft 365 Copilot provides AI capabilities grounded in your organization's Microsoft 365 data, including emails, documents, meetings, and more.

> NOTE: As per the conversation in [discussion 1853](https://github.com/danielmiessler/Fabric/discussions/1853) - enterprise users with restrictive consent policies will probably need their IT admin to either create an app registration with the required permissions, or grant admin consent for an existing app like Graph Explorer.

## Table of Contents

- [What is Microsoft 365 Copilot?](#what-is-microsoft-365-copilot)
- [Requirements](#requirements)
- [Azure AD App Registration](#azure-ad-app-registration)
- [Obtaining Access Tokens](#obtaining-access-tokens)
- [Configuring Fabric for Copilot](#configuring-fabric-for-copilot)
- [Testing Your Setup](#testing-your-setup)
- [Usage Examples](#usage-examples)
- [Troubleshooting](#troubleshooting)
- [API Limitations](#api-limitations)

---

## What is Microsoft 365 Copilot?

**Microsoft 365 Copilot** is an AI-powered assistant that works across Microsoft 365 applications. When integrated with Fabric, it allows you to:

- **Query your organization's data**: Ask questions about emails, documents, calendars, and Teams chats
- **Grounded responses**: Get AI responses that are based on your actual Microsoft 365 content
- **Enterprise compliance**: All interactions respect your organization's security policies, permissions, and sensitivity labels

### Why Use Microsoft 365 Copilot with Fabric?

- **Enterprise-ready**: Built for organizations with compliance requirements
- **Data grounding**: Responses are based on your actual organizational data
- **Unified access**: Single integration for all Microsoft 365 content
- **Security**: Respects existing permissions and access controls

---

## Requirements

Before you begin, ensure you have:

### Licensing Requirements

1. **Microsoft 365 Copilot License**: Required for each user accessing the API
2. **Microsoft 365 E3 or E5 Subscription** (or equivalent): Foundation for Copilot services

### Technical Requirements

1. **Azure AD Tenant**: Your organization's Azure Active Directory
2. **Azure AD App Registration**: To authenticate with Microsoft Graph
3. **Delegated Permissions**: The Chat API only supports delegated (user) permissions, not application permissions

### Permissions Required

The following Microsoft Graph permissions are needed:

| Permission | Type | Description |
|------------|------|-------------|
| `Sites.Read.All` | Delegated | Read SharePoint sites |
| `Mail.Read` | Delegated | Read user's email |
| `People.Read.All` | Delegated | Read organization's people directory |
| `OnlineMeetingTranscript.Read.All` | Delegated | Read meeting transcripts |
| `Chat.Read` | Delegated | Read Teams chat messages |
| `ChannelMessage.Read.All` | Delegated | Read Teams channel messages |
| `ExternalItem.Read.All` | Delegated | Read external content connectors |

---

## Azure AD App Registration

### Step 1: Create the App Registration

1. Go to the [Azure Portal](https://portal.azure.com)
2. Navigate to **Azure Active Directory** > **App registrations**
3. Click **New registration**
4. Configure the application:
   - **Name**: `Fabric CLI - Copilot`
   - **Supported account types**: Select "Accounts in this organizational directory only"
   - **Redirect URI**: Select "Public client/native (mobile & desktop)" and enter `http://localhost:8400/callback`
5. Click **Register**

### Step 2: Note Your Application IDs

After registration, note these values from the **Overview** page:

- **Application (client) ID**: e.g., `12345678-1234-1234-1234-123456789abc`
- **Directory (tenant) ID**: e.g., `abcdef12-3456-7890-abcd-ef1234567890`

### Step 3: Configure API Permissions

1. Go to **API permissions** in your app registration
2. Click **Add a permission**
3. Select **Microsoft Graph**
4. Select **Delegated permissions**
5. Add the following permissions:
   - `Sites.Read.All`
   - `Mail.Read`
   - `People.Read.All`
   - `OnlineMeetingTranscript.Read.All`
   - `Chat.Read`
   - `ChannelMessage.Read.All`
   - `ExternalItem.Read.All`
   - `offline_access` (for refresh tokens)
6. Click **Add permissions**
7. **Important**: Click **Grant admin consent for [Your Organization]** (requires admin privileges)

### Step 4: Configure Authentication (Optional - For Confidential Clients)

If you want to use client credentials for token refresh:

1. Go to **Certificates & secrets**
2. Click **New client secret**
3. Add a description and select an expiration
4. Click **Add**
5. **Important**: Copy the secret value immediately (it won't be shown again)

---

## Obtaining Access Tokens

The Microsoft 365 Copilot Chat API requires **delegated permissions**, meaning you need to authenticate as a user. There are several ways to obtain tokens:

### Option 1: Using Azure CLI (Recommended for Development)

```bash
# Install Azure CLI if not already installed
# https://docs.microsoft.com/en-us/cli/azure/install-azure-cli

# Login with your work account
az login --tenant YOUR_TENANT_ID

# Get an access token for Microsoft Graph
az account get-access-token --resource https://graph.microsoft.com --query accessToken -o tsv
```

### Option 2: Using Device Code Flow

For headless environments or when browser authentication isn't possible:

```bash
# Request device code
curl -X POST "https://login.microsoftonline.com/YOUR_TENANT_ID/oauth2/v2.0/devicecode" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "client_id=YOUR_CLIENT_ID&scope=Sites.Read.All Mail.Read People.Read.All OnlineMeetingTranscript.Read.All Chat.Read ChannelMessage.Read.All ExternalItem.Read.All offline_access"

# Follow the instructions to authenticate in a browser
# Then poll for the token using the device_code from the response
```

### Option 3: Using Microsoft Graph Explorer (For Testing)

1. Go to [Microsoft Graph Explorer](https://developer.microsoft.com/en-us/graph/graph-explorer)
2. Sign in with your work account
3. Click the gear icon > "Select permissions"
4. Enable the required permissions
5. Use the access token from the "Access token" tab

### Option 4: Using MSAL Libraries

For production applications, use Microsoft Authentication Library (MSAL):

```go
// Example using Azure Identity SDK for Go
import "github.com/Azure/azure-sdk-for-go/sdk/azidentity"

cred, err := azidentity.NewInteractiveBrowserCredential(&azidentity.InteractiveBrowserCredentialOptions{
    TenantID: "YOUR_TENANT_ID",
    ClientID: "YOUR_CLIENT_ID",
})
```

---

## Configuring Fabric for Copilot

### Method 1: Using Fabric Setup (Recommended)

1. **Run Fabric Setup:**

   ```bash
   fabric --setup
   ```

2. **Select Copilot from the menu:**
   - Find `Copilot` in the numbered list
   - Enter the number and press Enter

3. **Enter Configuration Values:**

   ```
   [Copilot] Enter your Azure AD Tenant ID:
   > contoso.onmicrosoft.com

   [Copilot] Enter your Azure AD Application (Client) ID:
   > 12345678-1234-1234-1234-123456789abc

   [Copilot] Enter your Azure AD Client Secret (optional):
   > (press Enter to skip, or enter secret for token refresh)

   [Copilot] Enter a pre-obtained OAuth2 Access Token:
   > eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs...

   [Copilot] Enter a pre-obtained OAuth2 Refresh Token (optional):
   > (press Enter to skip, or enter refresh token)

   [Copilot] Enter your timezone:
   > America/New_York
   ```

### Method 2: Manual Configuration

Edit `~/.config/fabric/.env`:

```bash
# Microsoft 365 Copilot Configuration
COPILOT_TENANT_ID=contoso.onmicrosoft.com
COPILOT_CLIENT_ID=12345678-1234-1234-1234-123456789abc
COPILOT_CLIENT_SECRET=your-client-secret-if-applicable
COPILOT_ACCESS_TOKEN=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIs...
COPILOT_REFRESH_TOKEN=your-refresh-token-if-available
COPILOT_API_BASE_URL=https://graph.microsoft.com/beta/copilot
COPILOT_TIME_ZONE=America/New_York
```

### Verify Configuration

```bash
fabric --listmodels | grep -i copilot
```

Expected output:

```
        [X]    Copilot|microsoft-365-copilot
```

---

## Testing Your Setup

### Basic Test

```bash
# Simple query
echo "What meetings do I have tomorrow?" | fabric --vendor Copilot

# With explicit model (though there's only one)
echo "Summarize my recent emails" | fabric --vendor Copilot --model microsoft-365-copilot
```

### Test with Streaming

```bash
echo "What are the key points from my last team meeting?" | \
  fabric --vendor Copilot --stream
```

### Test with Patterns

```bash
# Use a pattern with Copilot
echo "Find action items from my recent emails" | \
  fabric --pattern extract_wisdom --vendor Copilot
```

---

## Usage Examples

### Query Calendar

```bash
echo "What meetings do I have scheduled for next week?" | fabric --vendor Copilot
```

### Summarize Emails

```bash
echo "Summarize the emails I received yesterday from my manager" | fabric --vendor Copilot
```

### Search Documents

```bash
echo "Find documents about the Q4 budget proposal" | fabric --vendor Copilot
```

### Team Collaboration

```bash
echo "What were the main discussion points in the engineering standup channel this week?" | fabric --vendor Copilot
```

### Meeting Insights

```bash
echo "What action items came out of the project review meeting on Monday?" | fabric --vendor Copilot
```

### Using with Fabric Patterns

```bash
# Extract wisdom from organizational content
echo "What are the key decisions from last month's leadership updates?" | \
  fabric --pattern extract_wisdom --vendor Copilot

# Summarize with a specific pattern
echo "Summarize the HR policy document about remote work" | \
  fabric --pattern summarize --vendor Copilot
```

---

## Troubleshooting

### Error: "Authentication failed" or "401 Unauthorized"

**Cause**: Invalid or expired access token

**Solutions**:

1. Obtain a fresh access token:

   ```bash
   az account get-access-token --resource https://graph.microsoft.com --query accessToken -o tsv
   ```

2. Update your configuration:

   ```bash
   fabric --setup
   # Select Copilot and enter the new token
   ```

3. Check token hasn't expired (tokens typically expire after 1 hour)

### Error: "403 Forbidden"

**Cause**: Missing permissions or admin consent not granted

**Solutions**:

1. Verify all required permissions are added to your app registration
2. Ensure admin consent has been granted
3. Check that your user has a Microsoft 365 Copilot license

### Error: "Failed to create conversation"

**Cause**: API access issues or service unavailable

**Solutions**:

1. Verify the API base URL is correct: `https://graph.microsoft.com/beta/copilot`
2. Check Microsoft 365 service status
3. Ensure your organization has Copilot enabled

### Error: "Rate limit exceeded"

**Cause**: Too many requests

**Solutions**:

1. Wait a few minutes before retrying
2. Reduce request frequency
3. Consider batching queries

### Token Refresh Not Working

**Cause**: Missing client secret or refresh token

**Solutions**:

1. Ensure you have both a refresh token and client secret configured
2. Re-authenticate to get new tokens
3. Check that your app registration supports refresh tokens (public client)

---

## API Limitations

### Current Limitations

1. **Preview API**: The Chat API is currently in preview (`/beta` endpoint) and subject to change
2. **Delegated Only**: Only delegated (user) permissions are supported, not application permissions
3. **Single Model**: Copilot exposes a single unified model, unlike other vendors with multiple model options
4. **Enterprise Only**: Requires Microsoft 365 work or school accounts
5. **Licensing**: Requires Microsoft 365 Copilot license per user

### Rate Limits

The Microsoft Graph API has rate limits that apply:

- Per-app limits
- Per-user limits
- Tenant-wide limits

Consult [Microsoft Graph throttling guidance](https://docs.microsoft.com/en-us/graph/throttling) for details.

### Data Freshness

Copilot indexes data from Microsoft 365 services. There may be a delay between when content is created and when it becomes available in Copilot responses.

---

## Additional Resources

### Microsoft Documentation

- [Microsoft 365 Copilot APIs Overview](https://learn.microsoft.com/en-us/microsoft-365-copilot/extensibility/copilot-apis-overview)
- [Chat API Documentation](https://learn.microsoft.com/en-us/microsoft-365-copilot/extensibility/api/ai-services/chat/overview)
- [Microsoft Graph Authentication](https://learn.microsoft.com/en-us/graph/auth/)
- [Azure AD App Registration](https://learn.microsoft.com/en-us/azure/active-directory/develop/quickstart-register-app)

### Fabric Documentation

- [Fabric README](../README.md)
- [Contexts and Sessions Tutorial](./contexts-and-sessions-tutorial.md)

---

## Summary

Microsoft 365 Copilot integration with Fabric provides enterprise-ready AI capabilities grounded in your organization's data. Key points:

- **Enterprise compliance**: Works within your organization's security and compliance policies
- **Data grounding**: Responses are based on your actual Microsoft 365 content
- **Single model**: Exposes one unified AI model (`microsoft-365-copilot`)
- **Delegated auth**: Requires user authentication (OAuth2 with delegated permissions)
- **Preview API**: Currently in beta; expect changes

### Quick Start Commands

```bash
# 1. Set up Azure AD app registration (see guide above)

# 2. Get access token
az login --tenant YOUR_TENANT_ID
ACCESS_TOKEN=$(az account get-access-token --resource https://graph.microsoft.com --query accessToken -o tsv)

# 3. Configure Fabric
fabric --setup
# Select Copilot, enter tenant ID, client ID, and access token

# 4. Test it
echo "What meetings do I have this week?" | fabric --vendor Copilot
```

Happy prompting with Microsoft 365 Copilot!
