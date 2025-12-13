# GitHub Models Setup Guide for Fabric

This guide will walk you through setting up and using GitHub Models with Fabric CLI. GitHub Models provides free access to multiple AI models from OpenAI, Meta, Microsoft, DeepSeek, xAI, and other providers using only your GitHub credentials.

## Table of Contents

- [What are GitHub Models?](#what-are-github-models)
- [Getting Your GitHub Models API Key](#getting-your-github-models-api-key)
- [Configuring Fabric for GitHub Models](#configuring-fabric-for-github-models)
- [Testing Your Setup](#testing-your-setup)
- [Available Models](#available-models)
- [Rate Limits & Free Tier](#rate-limits--free-tier)
- [Troubleshooting](#troubleshooting)
- [Advanced Usage](#advanced-usage)

---

## What are GitHub Models?

**GitHub Models** is a free AI inference API platform that allows you to access multiple AI models using only your GitHub account. It's powered by Azure AI infrastructure and provides:

- **Unified Access**: Single API endpoint for models from multiple providers
- **No Extra API Keys**: Uses GitHub Personal Access Tokens (no separate OpenAI, Anthropic, etc. keys needed)
- **Free Tier**: Rate-limited free access perfect for prototyping and personal projects
- **Web Playground**: Test models directly at [github.com/marketplace/models](https://github.com/marketplace/models)
- **Compatible Format**: Works with OpenAI SDK standards

### Why Use GitHub Models with Fabric?

- **No Cost for Testing**: Free tier allows 50-150 requests/day depending on model
- **Multiple Providers**: Access OpenAI, Meta Llama, Microsoft Phi, DeepSeek, and more
- **Easy Setup**: Just one GitHub token instead of managing multiple API keys
- **Great for Learning**: Experiment with different models without financial commitment

---

## Getting Your GitHub Models API Key

GitHub Models uses **Personal Access Tokens (PAT)** instead of separate API keys.

### Step-by-Step Instructions

1. **Sign in to GitHub** at [github.com](https://github.com)

2. **Navigate to Token Settings:**
   - Click your profile picture (upper-right corner)
   - Click **Settings**
   - Scroll down the left sidebar to **Developer settings** (at the bottom)
   - Click **Personal access tokens** → **Fine-grained tokens** (recommended)

3. **Generate New Token:**
   - Click **Generate new token**
   - Give it a descriptive name: `Fabric CLI - GitHub Models`
   - Set expiration (recommended: 90 days or custom)
   - **Repository access**: Select "Public Repositories (read-only)" or "All repositories" (your choice)
   - **Permissions**:
     - Scroll down to **Account permissions**
     - Find **AI Models** and set to **Read-only** ✓
     - This grants the `models:read` scope
   - Click **Generate token** at the bottom

4. **Save Your Token:**
   - **IMPORTANT**: Copy the token immediately (starts with `github_pat_` or `ghp_`)
   - You won't be able to see it again!
   - Store it securely - this will be your `GITHUB_TOKEN`

### Security Best Practices

- ✅ Use fine-grained tokens with minimal permissions
- ✅ Set an expiration date (rotate tokens regularly)
- ✅ Never commit tokens to Git repositories
- ✅ Store in environment variables or secure credential managers
- ❌ Don't share tokens in chat, email, or screenshots

---

## Configuring Fabric for GitHub Models

### Method 1: Using Fabric Setup (Recommended)

This is the easiest and safest method:

1. **Run Fabric Setup:**

   ```bash
   fabric --setup
   ```

2. **Select GitHub from the Menu:**
   - You'll see a numbered list of AI vendors
   - Find `[8] GitHub (configured)` or similar
   - Enter the number (e.g., `8`) and press Enter

3. **Enter Your GitHub Token:**
   - When prompted for "API Key", paste your GitHub Personal Access Token
   - The token you created earlier (starts with `github_pat_` or `ghp_`)
   - Press Enter

4. **Verify Base URL (Optional):**
   - You'll be asked for "API Base URL"
   - Press Enter to use the default: `https://models.github.ai/inference`
   - Or customize if needed (advanced use only)

5. **Save and Exit:**
   - The setup wizard will save your configuration
   - You should see "GitHub (configured)" next time

### Method 2: Manual Configuration (Advanced)

If you prefer to manually edit the configuration file:

1. **Edit Environment File:**

   ```bash
   nano ~/.config/fabric/.env
   ```

2. **Add GitHub Configuration:**

   ```bash
   # GitHub Models API Key (your Personal Access Token)
   GITHUB_API_KEY=github_pat_YOUR_TOKEN_HERE

   # GitHub Models API Base URL (default, usually don't need to change)
   GITHUB_API_BASE_URL=https://models.github.ai/inference
   ```

   Save and exit (Ctrl+X, then Y, then Enter)

**Note**: The environment variable is `GITHUB_API_KEY`, not `GITHUB_TOKEN`.

### Verify Configuration

Check that your configuration is properly set:

```bash
grep GITHUB_API_KEY ~/.config/fabric/.env
```

You should see:

```text
GITHUB_API_KEY=github_pat_...
```

Or run setup again to verify:

```bash
fabric --setup
```

Look for `[8] GitHub (configured)` in the list.

---

## Testing Your Setup

### 1. List Available Models

Verify that Fabric can connect to GitHub Models and fetch the model list:

```bash
fabric --listmodels | grep GitHub
```

**Expected Output:**

```text
Available models:
...
$ fabric -L | grep GitHub
        [65]    GitHub|ai21-labs/ai21-jamba-1.5-large
        [66]    GitHub|cohere/cohere-command-a
        [67]    GitHub|cohere/cohere-command-r-08-2024
        [68]    GitHub|cohere/cohere-command-r-plus-08-2024
        [69]    GitHub|deepseek/deepseek-r1
        [70]    GitHub|deepseek/deepseek-r1-0528
        [71]    GitHub|deepseek/deepseek-v3-0324
        [72]    GitHub|meta/llama-3.2-11b-vision-instruct
        [73]    GitHub|meta/llama-3.2-90b-vision-instruct
... (and more)
```

### 2. Simple Chat Test

Test a basic chat completion with a small, fast model:

```bash
# Use gpt-4o-mini (fast and has generous rate limits)
fabric --vendor GitHub -m openai/gpt-4o-mini 'Why is th
e sky blue?'
```

**Expected**: You should see a response explaining Rayleigh scattering.

**Tip**: Model names from `--listmodels` can be used directly (e.g., `openai/gpt-4o-mini`, `openai/gpt-4o`, `meta/llama-4-maverick-17b-128e-instruct-fp8`).

### 3. Test with a Pattern

Use one of Fabric's built-in patterns:

```bash
echo "Artificial intelligence is transforming how we work and live." | \
  fabric --pattern summarize --vendor GitHub --model "openai/gpt-4o-mini"
```

### 4. Test Streaming

Verify streaming responses work:

```bash
echo "Count from 1 to 100" | \
  fabric --vendor GitHub --model "openai/gpt-4o-mini" --stream
```

You should see the response appear progressively, word by word.

### 5. Test with Different Models

Try a Meta Llama model:

```bash
# Use a Llama model
echo "Explain quantum computing" | \
  fabric --vendor GitHub --model "meta/Meta-Llama-3.1-8B-Instruct"
```

### Quick Validation Checklist

- [x] `--listmodels` shows GitHub models
- [x] Basic chat completion works
- [x] Patterns work with GitHub vendor
- [x] Streaming responses work
- [x] Can switch between different models

---

## Available Models

GitHub Models provides access to models from multiple providers. Models use the format: `{publisher}/{model-name}`

### OpenAI Models

| Model ID | Description | Tier | Best For |
|----------|-------------|------|----------|
| `openai/gpt-4.1` | Latest flagship GPT-4 | High | Complex tasks, reasoning |
| `openai/gpt-4o` | Optimized GPT-4 | High | General purpose, fast |
| `openai/gpt-4o-mini` | Compact, cost-effective | Low | Quick tasks, high volume |
| `openai/o1` | Advanced reasoning | High | Complex problem solving |
| `openai/o3` | Next-gen reasoning | High | Cutting-edge reasoning |

### Meta Llama Models

| Model ID | Description | Tier | Best For |
|----------|-------------|------|----------|
| `meta/llama-3.1-405b` | Largest Llama model | High | Complex tasks, accuracy |
| `meta/llama-3.1-70b` | Mid-size Llama | Low | Balanced performance |
| `meta/llama-3.1-8b` | Compact Llama | Low | Fast, efficient tasks |

### Microsoft Phi Models

| Model ID | Description | Tier | Best For |
|----------|-------------|------|----------|
| `microsoft/phi-4` | Latest Phi generation | Low | Efficient reasoning |
| `microsoft/phi-3-medium` | Mid-size variant | Low | General tasks |
| `microsoft/phi-3-mini` | Smallest Phi | Low | Quick, simple tasks |

### DeepSeek Models

| Model ID | Description | Tier | Special |
|----------|-------------|------|---------|
| `deepseek/deepseek-r1` | Reasoning model | Very Limited | 8 requests/day |
| `deepseek/deepseek-r1-0528` | Updated version | Very Limited | 8 requests/day |

### xAI Models

| Model ID | Description | Tier | Special |
|----------|-------------|------|---------|
| `xai/grok-3` | Latest Grok | Very Limited | 15 requests/day |
| `xai/grok-3-mini` | Smaller Grok | Very Limited | 15 requests/day |

### Getting the Full List

To see all currently available models:

```bash
fabric --listmodels | grep GitHub
```

Or for a formatted list with details, you can query the GitHub Models API directly:

```bash
curl -H "Authorization: Bearer $GITHUB_TOKEN" \
     -H "X-GitHub-Api-Version: 2022-11-28" \
     https://models.github.ai/catalog/models | jq '.[] | {id, publisher, tier: .rate_limit_tier}'
```

---

## Rate Limits & Free Tier

GitHub Models has tiered rate limits based on model complexity. Understanding these helps you use the free tier effectively.

### Low Tier Models (Recommended for High Volume)

**Models**: `gpt-4o-mini`, `llama-3.1-*`, `phi-*`

- **Requests per minute**: 15
- **Requests per day**: 150
- **Tokens per request**: 8,000 input / 4,000 output
- **Concurrent requests**: 5

**Best practices**: Use these for most Fabric patterns and daily tasks.

### High Tier Models (Use Sparingly)

**Models**: `gpt-4.1`, `gpt-4o`, `o1`, `o3`, `llama-3.1-405b`

- **Requests per minute**: 10
- **Requests per day**: 50
- **Tokens per request**: 8,000 input / 4,000 output
- **Concurrent requests**: 2

**Best practices**: Save for complex tasks, important queries, or when you need maximum quality.

### Very Limited Models

**Models**: `deepseek-r1`, `grok-3`

- **Requests per minute**: 1
- **Requests per day**: 8-15 (varies by model)
- **Tokens per request**: 4,000 input / 4,000 output
- **Concurrent requests**: 1

**Best practices**: Use only for special experiments or when you specifically need these models.

### Rate Limit Reset Times

- **Per-minute limits**: Reset every 60 seconds
- **Daily limits**: Reset at midnight UTC
- **Per-user**: Limits are tied to your GitHub account, not the token

### Enhanced Limits with GitHub Copilot

If you have a GitHub Copilot subscription, you get higher limits:

- **Copilot Business**: 2× daily request limits
- **Copilot Enterprise**: 3× daily limits + higher token limits

### What Happens When You Hit Limits?

You'll receive an HTTP 429 error with a message like:

```text
Rate limit exceeded. Try again in X seconds.
```

Fabric will display this error. Wait for the reset time and try again.

### Tips for Staying Within Limits

1. **Use low-tier models** for most tasks (`gpt-4o-mini`, `llama-3.1-8b`)
2. **Batch your requests** - process multiple items together when possible
3. **Cache results** - save responses for repeated queries
4. **Monitor usage** - keep track of daily request counts
5. **Set per-pattern models** - configure specific models for specific patterns (see Advanced Usage)

---

## Troubleshooting

### Error: "Authentication failed" or "Unauthorized"

**Cause**: Invalid or missing GitHub token

**Solutions**:

1. Verify token is in `.env` file:

   ```bash
   grep GITHUB_API_KEY ~/.config/fabric/.env
   ```

2. Check token has `models:read` permission:
   - Go to GitHub Settings → Developer settings → Personal access tokens
   - Click on your token
   - Verify "AI Models: Read-only" is checked

3. Re-run setup to reconfigure:

   ```bash
   fabric --setup
   # Select GitHub (number 8 or similar)
   # Enter your token again
   ```

4. Generate a new token if needed (tokens expire)

### Error: "Rate limit exceeded"

**Cause**: Too many requests in a short time period

**Solutions**:

1. Check which tier your model is in (see [Rate Limits](#rate-limits--free-tier))
2. Wait for the reset (check error message for wait time)
3. Switch to a lower-tier model:

   ```bash
   # Instead of gpt-4.1 (high tier)
   fabric --vendor GitHub --model openai/gpt-4.1 ...

   # Use gpt-4o-mini (low tier)
   fabric --vendor GitHub --model openai/gpt-4o-mini ...
   ```

### Error: "Model not found" or "Invalid model"

**Cause**: Model name format incorrect or model not available

**Solutions**:

1. Use correct format: `{publisher}/{model-name}`, e.g., `openai/gpt-4o-mini`

   ```bash
   # ❌ Wrong
   fabric --vendor GitHub --model gpt-4o-mini

   # ✅ Correct
   fabric --vendor GitHub --model openai/gpt-4o-mini
   ```

2. List available models to verify name:

   ```bash
   fabric --listmodels --vendor GitHub | grep -i "gpt-4"
   ```

### Error: "Cannot list models" or Empty model list

**Cause**: API endpoint issue or authentication problem

**Solutions**:

1. Test direct API access:

   ```bash
   curl -H "Authorization: Bearer $GITHUB_TOKEN" \
        -H "X-GitHub-Api-Version: 2022-11-28" \
        https://models.github.ai/catalog/models
   ```

2. If curl works but Fabric doesn't, rebuild Fabric:

   ```bash
   cd /path/to/fabric
   go build ./cmd/fabric
   ```

3. Check for network/firewall issues blocking `models.github.ai`

### Error: "Response format not supported"

**Cause**: This should be fixed in the latest version with direct fetch fallback

**Solutions**:

1. Update to the latest Fabric version with PR #1839 merged
2. Verify you're on a version that includes the `FetchModelsDirectly` fallback

### Models are slow to respond

**Cause**: High tier models have limited concurrency, or GitHub Models API congestion

**Solutions**:

1. Switch to faster models:
   - `openai/gpt-4o-mini` instead of `gpt-4.1`
   - `meta/llama-3.1-8b` instead of `llama-3.1-405b`

2. Check your internet connection

3. Try again later (API may be experiencing high traffic)

### Token expires or becomes invalid

**Cause**: Tokens have expiration dates or can be revoked

**Solutions**:

1. Generate a new token (see [Getting Your GitHub Models API Key](#getting-your-github-models-api-key))
2. Update `.env` file with new token
3. Set longer expiration when creating tokens (e.g., 90 days)

---

## Advanced Usage

### Using Specific Models with Patterns

You can specify which model to use with any pattern:

```bash
# Use GPT-4.1 with the analyze_claims pattern
cat article.txt | fabric --pattern analyze_claims \
  --vendor GitHub --model openai/gpt-4.1

# Use Llama for summarization
cat document.txt | fabric --pattern summarize \
  --vendor GitHub --model meta/llama-3.1-70b
```

### Per-Pattern Model Mapping

Set default models for specific patterns using environment variables:

Edit `~/.config/fabric/.env`:

```bash
# Use GPT-4.1 for complex analysis
FABRIC_MODEL_analyze_claims=GitHub|openai/gpt-4.1
FABRIC_MODEL_extract_wisdom=GitHub|openai/gpt-4.1

# Use GPT-4o-mini for simple tasks
FABRIC_MODEL_summarize=GitHub|openai/gpt-4o-mini
FABRIC_MODEL_extract_article_wisdom=GitHub|openai/gpt-4o-mini

# Use Llama for code tasks
FABRIC_MODEL_explain_code=GitHub|meta/llama-3.1-70b
```

Now when you run:

```bash
cat article.txt | fabric --pattern analyze_claims
```

It will automatically use `GitHub|openai/gpt-4.1` without needing to specify the vendor and model.

### Comparing Responses Across Providers

Compare how different models respond to the same input:

```bash
# OpenAI GPT-4o-mini
echo "Explain quantum computing" | \
  fabric --vendor GitHub --model openai/gpt-4o-mini > response_openai.txt

# Meta Llama
echo "Explain quantum computing" | \
  fabric --vendor GitHub --model meta/llama-3.1-70b > response_llama.txt

# Microsoft Phi
echo "Explain quantum computing" | \
  fabric --vendor GitHub --model microsoft/phi-4 > response_phi.txt

# Compare
diff response_openai.txt response_llama.txt
```

### Testing Different Models for a Pattern

Find the best model for your use case:

```bash
# Create a test script
cat > test_models.sh << 'EOF'
#!/bin/bash

INPUT="Explain the concept of recursion in programming"
PATTERN="explain_code"

for MODEL in "openai/gpt-4o-mini" "meta/llama-3.1-8b" "microsoft/phi-4"; do
  echo "=== Testing $MODEL ==="
  echo "$INPUT" | fabric --pattern "$PATTERN" --vendor GitHub --model "$MODEL"
  echo ""
done
EOF

chmod +x test_models.sh
./test_models.sh
```

### Quick Test Without Setup

If you want to quickly test without running full setup, you can set the environment variable directly:

```bash
# Temporary test (this session only)
export GITHUB_API_KEY=github_pat_YOUR_TOKEN_HERE

# Test immediately
fabric --listmodels --vendor GitHub
```

This is useful for quick tests, but we recommend using `fabric --setup` for permanent configuration.

### Streaming for Long Responses

For long-form content, use streaming to see results as they generate:

```bash
cat long_article.txt | \
  fabric --pattern summarize \
  --vendor GitHub --model openai/gpt-4o-mini \
  --stream
```

### Saving Token Usage

Monitor your usage to stay within rate limits:

```bash
# Create a simple usage tracker
echo "$(date): Used gpt-4.1 for analyze_claims" >> ~/.config/fabric/usage.log

# Check daily usage
grep "$(date +%Y-%m-%d)" ~/.config/fabric/usage.log | wc -l
```

### Environment-Based Configuration

Create different profiles for different use cases:

```bash
# Development profile (uses free GitHub Models)
cat > ~/.config/fabric/.env.dev << EOF
GITHUB_TOKEN=github_pat_dev_token_here
DEFAULT_VENDOR=GitHub
DEFAULT_MODEL=openai/gpt-4o-mini
EOF

# Production profile (uses paid OpenAI)
cat > ~/.config/fabric/.env.prod << EOF
OPENAI_API_KEY=sk-prod-key-here
DEFAULT_VENDOR=OpenAI
DEFAULT_MODEL=gpt-4
EOF

# Switch profiles
ln -sf ~/.config/fabric/.env.dev ~/.config/fabric/.env
```

---

## Additional Resources

### Official Documentation

- [GitHub Models Quickstart](https://docs.github.com/en/github-models/quickstart)
- [GitHub Models API Reference](https://docs.github.com/en/rest/models)
- [GitHub Models Marketplace](https://github.com/marketplace/models)

### Fabric Documentation

- [Fabric README](../README.md)
- [Contexts and Sessions Tutorial](./contexts-and-sessions-tutorial.md)
- [Using Speech-to-Text](./Using-Speech-To-Text.md)

### Community

- [Fabric GitHub Repository](https://github.com/danielmiessler/fabric)
- [Fabric Issues](https://github.com/danielmiessler/fabric/issues)
- [Fabric Discussions](https://github.com/danielmiessler/fabric/discussions)

---

## Summary

GitHub Models provides an excellent way to experiment with AI models through Fabric without managing multiple API keys or incurring costs. Key points:

✅ **Free to start**: No credit card required, 50-150 requests/day
✅ **Multiple providers**: OpenAI, Meta, Microsoft, DeepSeek, xAI
✅ **Simple setup**: Just one GitHub token via `fabric --setup`
✅ **Great for learning**: Try different models and patterns
✅ **Production path**: Can upgrade to paid tier when ready

### Quick Start Commands

```bash
# 1. Get GitHub token with models:read scope from:
#    https://github.com/settings/tokens

# 2. Configure Fabric
fabric --setup
# Select [8] GitHub
# Paste your token when prompted

# 3. List available models
fabric --listmodels --vendor GitHub | grep gpt-4o

# 4. Try it out with gpt-4o-mini
echo "What is AI?" | fabric --vendor GitHub --model "gpt-4o-mini"
```

**Recommended starting point**: Use `gpt-4o-mini` for most patterns - it's fast, capable, and has generous rate limits (150 requests/day).

**Available Models**: `gpt-4o`, `gpt-4o-mini`, `Meta-Llama-3.1-8B-Instruct`, `Meta-Llama-3.1-70B-Instruct`, `Mistral-large-2407`, and more. Use `--listmodels` to see the complete list.

Happy prompting! 🚀
