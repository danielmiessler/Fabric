# Fabric REST API

Fabric's REST API provides HTTP access to all core functionality: chat completions, pattern management, contexts, sessions, and more.

## Quick Start

Start the server:

```bash
fabric --serve
```

The server runs on `http://localhost:8080` by default.

Test it:

```bash
curl http://localhost:8080/patterns/names
```

## Interactive API Documentation

Fabric includes Swagger/OpenAPI documentation with an interactive UI:

- **Swagger UI**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- **OpenAPI JSON**: [http://localhost:8080/swagger/doc.json](http://localhost:8080/swagger/doc.json)
- **OpenAPI YAML**: [http://localhost:8080/swagger/swagger.yaml](http://localhost:8080/swagger/swagger.yaml)

The Swagger UI lets you:

- Browse all available endpoints
- View request/response schemas
- Test API calls directly in your browser
- See authentication requirements

**Note:** Swagger documentation endpoints are publicly accessible even when API key authentication is enabled. Only the actual API endpoints require authentication

## Server Options

| Flag | Description | Default |
| ------ | ------------- | --------- |
| `--serve` | Start the REST API server | - |
| `--address` | Server address and port | `:8080` |
| `--api-key` | Enable API key authentication | (none) |
| `--cors-origins` | Browser origins that can call the server. Repeat the flag for each origin. `*` opens it to all. | (none) |

Example with custom configuration:

```bash
fabric --serve --address :9090 --api-key my_secret_key
```

## Authentication

When you set an API key with `--api-key`, all requests must include:

```http
X-API-Key: your-api-key-here
```

Example:

```bash
curl -H "X-API-Key: my_secret_key" http://localhost:8080/patterns/names
```

Use a random key of 16 or more characters. The server warns at startup when the key is shorter. The server has no limit on wrong keys.

The API key is an admin key. A client with the key can read and change patterns, contexts and sessions, and can change the configuration, for example the Ollama URL. Do not give the key to a user that must not see the prompts of other users.

The server uses plain HTTP. On a non-loopback address, the key and the prompts go over the network as clear text. Put a reverse proxy with TLS in front of the server.

Without an API key, the server logs a warning and examines the request headers:

- The `Host` header must be `localhost`, a loopback IP address or the host of a `--cors-origins` entry. If you open the server through a LAN address or another host name, add that origin with `--cors-origins`.
- A `POST`, `PUT`, `DELETE` or `PATCH` request with an `Origin` header must come from the same origin (host and port), a `--cors-origins` entry or the web UI dev or preview server (`localhost` or `127.0.0.1` on port `5173` or `4173`) with `Content-Type: application/json`. A page on a different loopback port is a different origin, and the server rejects it. A client that sends no `Origin` header, such as `curl`, can send these requests.

The server rejects other requests with `403 Forbidden`. With an API key, the server does not do these checks.

## CORS

By default, the server sends no CORS headers. A browser or WebView client on a different origin cannot call it. To let a client call the server, list its origin:

```bash
fabric --serve --cors-origins http://localhost:1420 --cors-origins tauri://localhost
```

You can also set `FABRIC_CORS_ORIGINS` to a comma-separated list, in the shell or in `~/.config/fabric/.env`. The flag applies to `--serve` and `--serveOllama`.

The server answers `OPTIONS` preflight requests before the API key check. It accepts the request headers `Content-Type` and `X-API-Key`.

`--cors-origins '*'` lets all origins call the server. The server logs a warning at startup. The server does not start when you set `*` without `--api-key`. The server never allows the origin `null`.

## Content Type

The endpoints that read a JSON body need the header `Content-Type: application/json`. These endpoints are `POST /chat`, `POST /patterns/:name/apply`, `POST /config/update`, `POST /youtube/transcript` and `POST /api/chat`. The server rejects other content types with `415 Unsupported Media Type`.

The save endpoints (`POST /patterns/:name`, `POST /contexts/:name` and `POST /sessions/:name`) take the raw body with any content type.

## Endpoints

### Chat Completions

Stream AI responses using Server-Sent Events (SSE).

**Endpoint:** `POST /chat`

**Request:**

```json
{
  "prompts": [
    {
      "userInput": "Explain quantum computing",
      "vendor": "openai",
      "model": "gpt-5.2",
      "patternName": "explain",
      "contextName": "",
      "strategyName": "",
      "variables": {}
    }
  ],
  "language": "en",
  "temperature": 0.7,
  "topP": 0.9,
  "frequencyPenalty": 0,
  "presencePenalty": 0,
  "thinking": 0
}
```

**Prompt Fields:**

| Field | Required | Default | Description |
| ------- | ---------- | --------- | ------------- |
| `userInput` | **Yes** | - | Your message or question |
| `vendor` | **Yes** | - | AI provider: `openai`, `anthropic`, `gemini`, `ollama`, etc. |
| `model` | **Yes** | - | Model name: `gpt-5.2`, `claude-sonnet-4.5`, `gemini-2.0-flash-exp`, etc. |
| `patternName` | No | `""` | Pattern to apply (from `~/.config/fabric/patterns/`) |
| `contextName` | No | `""` | Context to prepend (from `~/.config/fabric/contexts/`) |
| `strategyName` | No | `""` | Strategy to use (from `~/.config/fabric/strategies/`) |
| `variables` | No | `{}` | Variable substitutions for patterns (e.g., `{"role": "expert"}`) |

**Chat Options:**

| Field | Required | Default | Description |
| ------- | ---------- | --------- | ------------- |
| `language` | No | `"en"` | Language code for responses |
| `temperature` | No | `0.7` | Randomness (0.0-1.0) |
| `topP` | No | `0.9` | Nucleus sampling (0.0-1.0) |
| `frequencyPenalty` | No | `0.0` | Reduce repetition (-2.0 to 2.0) |
| `presencePenalty` | No | `0.0` | Encourage new topics (-2.0 to 2.0) |
| `thinking` | No | `0` | Reasoning level (0=off, or numeric for tokens) |

**Response:**

Server-Sent Events stream with `Content-Type: text/readystream`. Each line contains JSON:

```json
{"type": "content", "format": "markdown", "content": "Quantum computing uses..."}
{"type": "content", "format": "markdown", "content": " quantum mechanics..."}
{"type": "complete", "format": "markdown", "content": ""}
```

**Types:**

- `content` - Response chunk
- `error` - Error message
- `complete` - Stream finished

**Formats:**

- `markdown` - Standard text
- `mermaid` - Mermaid diagram
- `plain` - Plain text

**Example:**

```bash
curl -X POST http://localhost:8080/chat \
  -H "Content-Type: application/json" \
  -d '{
    "prompts": [{
      "userInput": "What is Fabric?",
      "vendor": "openai",
      "model": "gpt-5.2",
      "patternName": "explain"
    }]
  }'
```

### Patterns

Manage reusable AI prompts.

| Method | Endpoint | Description |
| -------- | ---------- | ------------- |
| `GET` | `/patterns/names` | List all pattern names |
| `GET` | `/patterns/:name` | Get pattern content |
| `GET` | `/patterns/exists/:name` | Check if pattern exists |
| `POST` | `/patterns/:name` | Create or update pattern |
| `DELETE` | `/patterns/:name` | Delete pattern |
| `PUT` | `/patterns/rename/:oldName/:newName` | Rename pattern |
| `POST` | `/patterns/:name/apply` | Apply pattern with variables |

The server runs only the `text` and `datetime` template plugins in a pattern. A pattern with a `sys`, `file` or `fetch` plugin token or an `{{ext:...}}` token gives an error. A client can save a pattern, so this stops a client from reading the server environment or files, fetching a URL through the server, or running an extension. The `/chat` endpoint applies patterns in the same way. The CLI runs all plugins.

**Example - Get pattern:**

```bash
curl http://localhost:8080/patterns/summarize
```

**Example - Apply pattern with variables:**

```bash
curl -X POST http://localhost:8080/patterns/translate/apply \
  -H "Content-Type: application/json" \
  -d '{
    "input": "Hello world",
    "variables": {"lang_code": "es"}
  }'
```

**Example - Create pattern:**

```bash
curl -X POST http://localhost:8080/patterns/my_custom_pattern \
  -H "Content-Type: text/plain" \
  -d "You are an expert in explaining complex topics simply..."
```

### Contexts

Manage context snippets that prepend to prompts.

| Method | Endpoint | Description |
| -------- | ---------- | ------------- |
| `GET` | `/contexts/names` | List all context names |
| `GET` | `/contexts/:name` | Get context content |
| `GET` | `/contexts/exists/:name` | Check if context exists |
| `POST` | `/contexts/:name` | Create or update context |
| `DELETE` | `/contexts/:name` | Delete context |
| `PUT` | `/contexts/rename/:oldName/:newName` | Rename context |

### Sessions

Manage chat conversation history.

| Method | Endpoint | Description |
| -------- | ---------- | ------------- |
| `GET` | `/sessions/names` | List all session names |
| `GET` | `/sessions/:name` | Get session messages (JSON array) |
| `GET` | `/sessions/exists/:name` | Check if session exists |
| `POST` | `/sessions/:name` | Save session messages |
| `DELETE` | `/sessions/:name` | Delete session |
| `PUT` | `/sessions/rename/:oldName/:newName` | Rename session |

### Models

List available AI models.

**Endpoint:** `GET /models/names`

**Response:**

```json
{
  "models": ["gpt-5.2", "gpt-5-mini", "claude-sonnet-4.5", "gemini-2.0-flash-exp"],
  "vendors": {
    "openai": ["gpt-5.2", "gpt-5-mini"],
    "anthropic": ["claude-sonnet-4.5", "claude-opus-4.5"],
    "gemini": ["gemini-2.0-flash-exp", "gemini-2.0-flash-thinking-exp"]
  }
}
```

### Strategies

List available prompt strategies (Chain of Thought, etc.).

**Endpoint:** `GET /strategies`

**Response:**

```json
[
  {
    "name": "chain_of_thought",
    "description": "Think step by step",
    "prompt": "Let's think through this step by step..."
  }
]
```

### YouTube Transcripts

Extract transcripts from YouTube videos.

**Endpoint:** `POST /youtube/transcript`

**Request:**

```json
{
  "url": "https://youtube.com/watch?v=dQw4w9WgXcQ",
  "timestamps": false
}
```

**Response:**

```json
{
  "videoId": "Video ID",
  "title": "Video Title",
  "description" : "Video description...",
  "transcript": "Full transcript text..."
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/youtube/transcript \
  -H "Content-Type: application/json" \
  -d '{"url": "https://youtube.com/watch?v=dQw4w9WgXcQ", "timestamps": true}'
```

### Configuration

Manage API keys and environment settings.

**Get configuration:**

`GET /config`

Returns API keys and URLs for all configured vendors.

**Update configuration:**

`POST /config/update`

```json
{
  "OPENAI_API_KEY": "sk-...",
  "ANTHROPIC_API_KEY": "sk-ant-..."
}
```

Updates `~/.config/fabric/.env` with new values.

## Complete Workflow Examples

### Example: Summarize a YouTube Video

This example shows how to extract a YouTube transcript and summarize it using the `youtube_summary` pattern. This requires two API calls:

#### Step 1: Extract the transcript

```bash
curl -X POST http://localhost:8080/youtube/transcript \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://youtube.com/watch?v=dQw4w9WgXcQ",
    "timestamps": false
  }' > transcript.json
```

Response:

```json
{
  "videoId": "dQw4w9WgXcQ",
  "title": "Rick Astley - Never Gonna Give You Up (Official Video)",
  "description": "The official video for “Never Gonna Give You Up” by Rick Astley...",
  "transcript": "We're no strangers to love. You know the rules and so do I..."
}
```

#### Step 2: Summarize the transcript

Extract the transcript text and send it to the chat endpoint with the `youtube_summary` pattern:

```bash
# Extract transcript text from JSON
TRANSCRIPT=$(cat transcript.json | jq -r '.transcript')

# Send to chat endpoint with pattern
curl -X POST http://localhost:8080/chat \
  -H "Content-Type: application/json" \
  -d "{
    \"prompts\": [{
      \"userInput\": \"$TRANSCRIPT\",
      \"vendor\": \"openai\",
      \"model\": \"gpt-5.2\",
      \"patternName\": \"youtube_summary\"
    }]
  }"
```

#### Combined one-liner (using jq)

```bash
curl -s -X POST http://localhost:8080/youtube/transcript \
  -H "Content-Type: application/json" \
  -d '{"url": "https://youtube.com/watch?v=dQw4w9WgXcQ", "timestamps": false}' | \
jq -r '.transcript' | \
xargs -I {} curl -X POST http://localhost:8080/chat \
  -H "Content-Type: application/json" \
  -d "{\"prompts\":[{\"userInput\":\"{}\",\"vendor\":\"openai\",\"model\":\"gpt-5.2\",\"patternName\":\"youtube_summary\"}]}"
```

#### Alternative: Using a script

```bash
#!/bin/bash
YOUTUBE_URL="https://youtube.com/watch?v=dQw4w9WgXcQ"
API_BASE="http://localhost:8080"

# Step 1: Get transcript
echo "Extracting transcript..."
TRANSCRIPT=$(curl -s -X POST "$API_BASE/youtube/transcript" \
  -H "Content-Type: application/json" \
  -d "{\"url\":\"$YOUTUBE_URL\",\"timestamps\":false}" | jq -r '.transcript')

# Step 2: Summarize with pattern
echo "Generating summary..."
curl -X POST "$API_BASE/chat" \
  -H "Content-Type: application/json" \
  -d "{
    \"prompts\": [{
      \"userInput\": $(echo "$TRANSCRIPT" | jq -Rs .),
      \"vendor\": \"openai\",
      \"model\": \"gpt-5.2\",
      \"patternName\": \"youtube_summary\"
    }]
  }"
```

#### Comparison with CLI

The CLI combines these steps automatically:

```bash
# CLI version (single command)
fabric -y "https://youtube.com/watch?v=dQw4w9WgXcQ" --pattern youtube_summary
```

The API provides more flexibility by separating transcript extraction and summarization, allowing you to:

- Extract the transcript once and process it multiple ways
- Apply different patterns to the same transcript
- Store the transcript for later use
- Use different models or vendors for summarization

## Docker Usage

Run the server in Docker:

```bash
# Setup (first time)
mkdir -p $HOME/.fabric-config
docker run --rm -it \
  -v $HOME/.fabric-config:/root/.config/fabric \
  kayvan/fabric:latest --setup

# Start server
docker run --rm -it \
  -p 8080:8080 \
  -v $HOME/.fabric-config:/root/.config/fabric \
  kayvan/fabric:latest --serve

# With authentication
docker run --rm -it \
  -p 8080:8080 \
  -v $HOME/.fabric-config:/root/.config/fabric \
  kayvan/fabric:latest --serve --api-key my_secret_key
```

## Ollama Compatibility Mode

Fabric can emulate Ollama's API endpoints:

```bash
fabric --serveOllama --address :11434
```

This mode provides:

- `GET /api/tags` - Lists patterns as models
- `GET /api/version` - Server version
- `POST /api/chat` - Ollama-compatible chat endpoint

`POST /api/chat` needs the header `Content-Type: application/json`. The Ollama client libraries send it. Add `-H "Content-Type: application/json"` to a `curl` example from the Ollama documentation.

This mode also serves the `/patterns`, `/contexts`, `/sessions`, `/chat`, `/config` and `/models` endpoints. Patterns have the same plugin limit as in the REST API: the server runs only the `text` and `datetime` template plugins.

## Error Handling

All endpoints return standard HTTP status codes:

- `200 OK` - Success
- `400 Bad Request` - Invalid input
- `401 Unauthorized` - Missing or invalid API key
- `403 Forbidden` - The server has no API key, and it does not accept the `Host` or `Origin` header
- `404 Not Found` - Resource not found
- `413 Content Too Large` - The request body is more than 50 MiB
- `415 Unsupported Media Type` - A JSON endpoint got a body that is not `application/json`
- `500 Internal Server Error` - Server error

Error responses include JSON with details:

```json
{
  "error": "Pattern not found: nonexistent"
}
```

## Rate Limiting

The server does not implement rate limiting. When deploying publicly, use a reverse proxy (nginx, Caddy) with rate limiting enabled.
