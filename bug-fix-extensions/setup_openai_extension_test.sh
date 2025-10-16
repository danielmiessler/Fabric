#!/usr/bin/env bash
set -euo pipefail

# === Check dependencies ===
echo "✅ Checking dependencies..."
if ! command -v jq &> /dev/null; then
    echo "❌ Error: jq is required but not installed"
    echo "   Install with: brew install jq (macOS) or apt-get install jq (Linux)"
    exit 1
fi

if ! command -v curl &> /dev/null; then
    echo "❌ Error: curl is required but not installed"
    exit 1
fi

# === Build the fixed code ===
echo "✅ building the fixed Fabric binary in top level dir as fabric-fix ..."
cd "$(dirname "$0")/.." || exit 1
go build -o fabric-fix ./cmd/fabric
cd - > /dev/null || exit 1

# === Paths ===
FABRIC_CONFIG_DIR="$HOME/.config/fabric"
EXT_BIN_DIR="$FABRIC_CONFIG_DIR/extensions/bin"
EXT_CONFIG_DIR="$FABRIC_CONFIG_DIR/extensions/configs"
PATTERN_DIR="$FABRIC_CONFIG_DIR/patterns/ai_echo"
WRAPPER="$EXT_BIN_DIR/openai-chat.sh"
YAML="$EXT_CONFIG_DIR/openai.yaml"
SYSTEM_MD="$PATTERN_DIR/system.md"

# === Ensure dirs exist ===
mkdir -p "$EXT_BIN_DIR" "$EXT_CONFIG_DIR" "$PATTERN_DIR"

# === Load environment variables from .env ===
ENV_FILE="$FABRIC_CONFIG_DIR/.env"
echo "✅ Loading environment variables from $ENV_FILE ..."
if [[ ! -f "$ENV_FILE" ]]; then
    echo "❌ Error: .env file not found at $ENV_FILE"
    echo "   Please run 'fabric --setup' first to create the configuration"
    exit 1
fi

# Source the .env file
set -a  # automatically export all variables
source "$ENV_FILE"
set +a

# Verify required variables are now set
if [[ -z "${OPENAI_API_KEY:-}" ]]; then
    echo "❌ Error: OPENAI_API_KEY not found in $ENV_FILE"
    echo "   Please add it to your .env file"
    exit 1
fi

if [[ -z "${ANTHROPIC_API_KEY:-}" ]]; then
    echo "❌ Error: ANTHROPIC_API_KEY not found in $ENV_FILE"
    echo "   Please add it to your .env file"
    exit 1
fi

# Set default base URLs if not provided
if [[ -z "${OPENAI_API_BASE_URL:-}" ]]; then
    echo "⚠️  Warning: OPENAI_API_BASE_URL not set in .env, using default"
    export OPENAI_API_BASE_URL="https://api.openai.com/v1"
fi

if [[ -z "${ANTHROPIC_API_BASE_URL:-}" ]]; then
    echo "⚠️  Warning: ANTHROPIC_API_BASE_URL not set in .env, using default"
    export ANTHROPIC_API_BASE_URL="https://api.anthropic.com/v1"
fi

# === Wrapper script ===
echo "✅ Creating OpenAI wrapper script at $WRAPPER ..."
cat > "$WRAPPER" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail

INPUT=$(jq -R -s '.' <<< "$*")
RESPONSE=$(curl "$OPENAI_API_BASE_URL/chat/completions" \
  -s -w "\n%{http_code}" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -d "{\"model\":\"gpt-4o-mini\",\"messages\":[{\"role\":\"user\",\"content\":$INPUT}]}")

HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')

if [[ "$HTTP_CODE" -ne 200 ]]; then
    echo "Error: HTTP $HTTP_CODE" >&2
    echo "$BODY" | jq -r '.error.message // "Unknown error"' >&2
    exit 1
fi

echo "$BODY" | jq -r '.choices[0].message.content'
EOF

chmod +x "$WRAPPER"

# === Claude/Anthropic wrapper script ===
CLAUDE_WRAPPER="$EXT_BIN_DIR/claude-chat.sh"
echo "✅ Creating Claude wrapper script at $CLAUDE_WRAPPER ..."
cat > "$CLAUDE_WRAPPER" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail

INPUT="$*"
RESPONSE=$(curl "$ANTHROPIC_API_BASE_URL/messages" \
  -s -w "\n%{http_code}" \
  -H "Content-Type: application/json" \
  -H "x-api-key: $ANTHROPIC_API_KEY" \
  -H "anthropic-version: 2023-06-01" \
  -d "{\"model\":\"claude-3-5-sonnet-20241022\",\"max_tokens\":1024,\"messages\":[{\"role\":\"user\",\"content\":\"$INPUT\"}]}")

HTTP_CODE=$(echo "$RESPONSE" | tail -n1)
BODY=$(echo "$RESPONSE" | sed '$d')

if [[ "$HTTP_CODE" -ne 200 ]]; then
    echo "Error: HTTP $HTTP_CODE" >&2
    echo "$BODY" | jq -r '.error.message // "Unknown error"' >&2
    exit 1
fi

echo "$BODY" | jq -r '.content[0].text'
EOF

chmod +x "$CLAUDE_WRAPPER"

# === Fabric extension config ===
echo "✅ Writing Fabric extension config at $YAML ..."
cat > "$YAML" <<EOF
name: openai
executable: "$WRAPPER"
type: executable
timeout: "30s"
description: "Call OpenAI Chat Completions API"
version: "1.0.0"

operations:
  chat:
    cmd_template: "{{executable}} {{value}}"

config:
  output:
    method: stdout
EOF

# === Claude/Anthropic extension config ===
CLAUDE_YAML="$EXT_CONFIG_DIR/claude.yaml"
echo "✅ Writing Claude extension config at $CLAUDE_YAML ..."
cat > "$CLAUDE_YAML" <<EOF
name: claude
executable: "$CLAUDE_WRAPPER"
type: executable
timeout: "30s"
description: "Call Anthropic Claude API"
version: "1.0.0"

operations:
  chat:
    cmd_template: "{{executable}} {{value}}"

config:
  output:
    method: stdout
EOF

# === Remove existing extensions if present ===
echo "✅ Removing existing extensions (if any)..."
../fabric-fix --rmextension=openai 2>/dev/null || true
../fabric-fix --rmextension=claude 2>/dev/null || true

# === Register extensions ===
echo "✅ Registering OpenAI extension with Fabric ..."
../fabric-fix --addextension "$YAML"

echo "✅ Registering Claude extension with Fabric ..."
../fabric-fix --addextension "$CLAUDE_YAML"

# === Create ai_echo pattern ===
echo "✅ Creating system.md pattern at $SYSTEM_MD ..."
cat > "$SYSTEM_MD" <<'EOF'
Summarize the responses from both AI models:

OpenAI Response:
{{ext:openai:chat:{{input}}}}

Claude Response:
{{ext:claude:chat:{{input}}}}
EOF

# === Verify pattern is registered ===
echo "✅ Verifying pattern is registered..."
../fabric-fix --listpatterns | grep -q "ai_echo" && echo "   ✓ Pattern 'ai_echo' found!" || echo "   ⚠️  Pattern might not be visible yet"

# === Run tests ===
echo
echo "🎉 Setup complete!"
echo
echo "🔹 Test 1: Direct extension calls"
echo '   echo "{{ext:openai:chat:What is Artificial Intelligence}}" | ../fabric-fix'
echo '   echo "{{ext:claude:chat:What is Artificial Intelligence}}" | ../fabric-fix'
echo
echo "🔹 Test 2: Pattern call (calls both OpenAI and Claude)"
echo '   echo "What is Artificial Intelligence" | ../fabric-fix -p ai_echo'
echo
echo "Pattern system.md looks like:"
cat "$SYSTEM_MD"