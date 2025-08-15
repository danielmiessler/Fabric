# Docker Test Environment for API Configuration Fix

This directory contains a Docker-based testing setup for fixing the issue where Fabric calls Ollama and Bedrock APIs even when not configured. This addresses the problem where unconfigured services show error messages during model listing.

## Quick Start

```bash
# Run all tests
./scripts/docker-test/test-runner.sh

# Interactive mode - pick which test to run
./scripts/docker-test/test-runner.sh -i

# Run specific test case
./scripts/docker-test/test-runner.sh gemini-only

# Shell into test environment
./scripts/docker-test/test-runner.sh -s gemini-only

# Build image only (for development)
./scripts/docker-test/test-runner.sh -b

# Show help
./scripts/docker-test/test-runner.sh -h
```

## Test Cases

1. **no-config**: No APIs configured
2. **gemini-only**: Only Gemini configured (reproduces original issue #1195)
3. **openai-only**: Only OpenAI configured
4. **ollama-only**: Only Ollama configured
5. **bedrock-only**: Only Bedrock configured
6. **mixed**: Multiple APIs configured (Gemini + OpenAI + Ollama)

## Environment Files

Each test case has a corresponding environment file in `scripts/docker-test/env/`:

- `env.no-config` - Empty configuration
- `env.gemini-only` - Only Gemini API key
- `env.openai-only` - Only OpenAI API key
- `env.ollama-only` - Only Ollama URL
- `env.bedrock-only` - Only Bedrock configuration
- `env.mixed` - Multiple API configurations

These files are volume-mounted into the Docker container and persist changes made with `fabric -S`.

## Interactive Mode & Shell Access

The interactive mode (`-i`) provides several options:

```text
Available test cases:

1) No APIs configured (no-config)
2) Only Gemini configured (gemini-only)
3) Only OpenAI configured (openai-only)
4) Only Ollama configured (ollama-only)
5) Only Bedrock configured (bedrock-only)
6) Mixed configuration (mixed)
7) Run all tests
0) Exit

Add '!' after number to shell into test environment (e.g., '1!' to shell into no-config)
```

### Shell Mode

- Use `1!`, `2!`, etc. to shell into any test environment
- Run `fabric -S` to configure APIs interactively
- Run `fabric --listmodels` or `fabric -L` to test model listing
- Changes persist in the environment files
- Type `exit` to return to test runner

## Expected Results

**Before Fix:**

- `no-config` and `gemini-only` tests show Ollama connection errors
- Tests show Bedrock authentication errors when BEDROCK_AWS_REGION not set
- Error: `Ollama Get "http://localhost:11434/api/tags": dial tcp...`
- Error: `Bedrock failed to list foundation models...`

**After Fix:**

- Clean output with no error messages for unconfigured services
- Only configured services appear in model listings
- Ollama only initialized when `OLLAMA_API_URL` is set
- Bedrock only initialized when `BEDROCK_AWS_REGION` is set

## Implementation Details

- **Volume-mounted configs**: Environment files are mounted to `/home/testuser/.config/fabric/.env`
- **Persistent state**: Configuration changes survive between test runs
- **Single Docker image**: Built once from `scripts/docker-test/base/Dockerfile`, reused for all tests
- **Isolated environments**: Each test uses its own environment file
- **Cross-platform**: Works on macOS, Linux, and Windows with Docker

## Development Workflow

1. Make code changes to fix API initialization logic
2. Run `./scripts/docker-test/test-runner.sh no-config` to test the main issue
3. Use `./scripts/docker-test/test-runner.sh -i` for interactive testing
4. Shell into environments (`1!`, `2!`, etc.) to debug specific configurations
5. Run all tests before submitting PR: `./scripts/docker-test/test-runner.sh`

## Architecture

The fix involves:

1. **Ollama**: Override `IsConfigured()` method to check for `OLLAMA_API_URL` env var
2. **Bedrock**: Modify `hasAWSCredentials()` to require `BEDROCK_AWS_REGION`
3. **Plugin Registry**: Only initialize providers when properly configured

This prevents unnecessary API calls and eliminates confusing error messages for users.
