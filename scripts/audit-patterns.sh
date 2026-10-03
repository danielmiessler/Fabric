#!/usr/bin/env bash
# audit-patterns.sh — Find maintenance issues in the Fabric patterns directory.
#
# Usage:
#   ./scripts/audit-patterns.sh [--strict] [patterns_dir]
#
#   --strict   exit 1 when any issue is found (for CI); default is always exit 0
#   -h,--help  show this help
#
# Checks:
#   1. Thin patterns        — system.md files under 15 lines (likely stubs)
#   2. Bloated patterns     — system.md files over 50 KB (likely embedded examples)
#   3. Stale model refs     — mentions of GPT-4, ChatGPT, or other vendor-specific models
#   4. Missing INPUT marker — pattern has no "# INPUT" section
#   5. Locale key gaps      — i18n keys present in en.json but missing from other locales
#   6. Completion gaps      — flags in flags.go with no entry in completions files
#
# Output: plain text report. Exit 0 even when issues found (use --strict for non-zero exit).

set -uo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
STRICT=0
PATTERNS_DIR=""

while [[ $# -gt 0 ]]; do
    case "$1" in
        --strict)   STRICT=1; shift ;;
        -h|--help)  sed -n '2,16p' "$0" | sed 's/^# \{0,1\}//'; exit 0 ;;
        -*)         echo "unknown option: $1 (try --help)" >&2; exit 2 ;;
        *)          PATTERNS_DIR="$1"; shift ;;
    esac
done
PATTERNS_DIR="${PATTERNS_DIR:-$REPO_ROOT/data/patterns}"

if [[ ! -d "$PATTERNS_DIR" ]]; then
    echo "patterns directory not found: $PATTERNS_DIR" >&2
    exit 2
fi

# Colorize only when writing to a terminal, so redirected output stays clean.
if [[ -t 1 ]]; then
    YLW='\033[0;33m'
    GRN='\033[0;32m'
    DIM='\033[2m'
    RST='\033[0m'
else
    YLW='' GRN='' DIM='' RST=''
fi

# Print paths relative to the repo root instead of scripts/../data/...
rel() { echo "${1#"$REPO_ROOT"/}"; }

issues=0

header() { echo; echo "━━━ $1 ━━━"; }
warn()   { echo -e "  ${YLW}⚠${RST}  $*"; issues=$((issues + 1)); }
info()   { echo -e "  ${DIM}·${RST}  $*"; }
ok()     { echo -e "  ${GRN}✓${RST}  $*"; }

# ── 1. Thin patterns ──────────────────────────────────────────────────────────
header "Thin patterns (< 15 lines)"
found=0
while IFS= read -r f; do
    lines=$(wc -l < "$f")
    name=$(basename "$(dirname "$f")")
    if [[ $lines -lt 15 ]]; then
        warn "$name  ($lines lines)  →  $(rel "$f")"
        found=1
    fi
done < <(find "$PATTERNS_DIR" -name "system.md" | sort)
[[ $found -eq 0 ]] && ok "None found"

# ── 2. Bloated patterns ───────────────────────────────────────────────────────
header "Bloated patterns (> 50 KB — likely embedded examples)"
found=0
while IFS= read -r f; do
    size=$(wc -c < "$f")
    name=$(basename "$(dirname "$f")")
    if [[ $size -gt 51200 ]]; then
        kb=$(( size / 1024 ))
        warn "$name  (${kb} KB)  →  $(rel "$f")"
        found=1
    fi
done < <(find "$PATTERNS_DIR" -name "system.md" | sort)
[[ $found -eq 0 ]] && ok "None found"

# ── 3. Stale model/vendor references ─────────────────────────────────────────
header "Stale model/vendor references"
STALE_PATTERN='GPT-4|GPT-3|gpt-4|gpt-3|ChatGPT|OpenAI'\''s|text-davinci|gpt4|gpt3|claude-[0-9]|gemini-pro'
found=0
while IFS= read -r f; do
    matches=$(grep -oE "$STALE_PATTERN" "$f" 2>/dev/null | sort -u | tr '\n' ' ')
    if [[ -n "$matches" ]]; then
        name=$(basename "$(dirname "$f")")
        warn "$name  →  found: ${matches% }"
        found=1
    fi
done < <(find "$PATTERNS_DIR" -name "system.md" | sort)
[[ $found -eq 0 ]] && ok "None found"

# ── 4. Missing INPUT marker ───────────────────────────────────────────────────
header "Patterns missing # INPUT section"
found=0
while IFS= read -r f; do
    if ! grep -q "^# INPUT" "$f" 2>/dev/null; then
        name=$(basename "$(dirname "$f")")
        warn "$name  →  $(rel "$f")"
        found=1
    fi
done < <(find "$PATTERNS_DIR" -name "system.md" | sort)
[[ $found -eq 0 ]] && ok "None found"

# ── 5. i18n locale key gaps ───────────────────────────────────────────────────
LOCALES_DIR="$REPO_ROOT/internal/i18n/locales"
if [[ -d "$LOCALES_DIR" ]]; then
    header "i18n locale key gaps (keys in en.json missing from other locales)"
    found=0
    for locale_file in "$LOCALES_DIR"/*.json; do
        lang=$(basename "$locale_file" .json)
        [[ "$lang" == "en" ]] && continue
        missing=$(python3 -c "
import json
with open('$LOCALES_DIR/en.json') as f:
    en = set(json.load(f).keys())
with open('$locale_file') as f:
    other = set(json.load(f).keys())
missing = sorted(en - other)
if missing:
    print('\n'.join(missing))
" 2>/dev/null)
        if [[ -n "$missing" ]]; then
            count=$(echo "$missing" | wc -l)
            warn "$lang  →  $count missing key(s):"
            echo "$missing" | while IFS= read -r key; do
                info "    $key"
            done
            found=1
        fi
    done
    [[ $found -eq 0 ]] && ok "All locales in sync"
fi

# ── 6. Shell completion gaps ──────────────────────────────────────────────────
FLAGS_FILE="$REPO_ROOT/internal/cli/flags.go"
BASH_FILE="$REPO_ROOT/completions/fabric.bash"
FISH_FILE="$REPO_ROOT/completions/fabric.fish"
ZSH_FILE="$REPO_ROOT/completions/_fabric"

if [[ -f "$FLAGS_FILE" ]]; then
    header "Shell completion gaps (long flags in flags.go missing from completions)"
    # Extract long flag names from struct tags: `long:"flag-name"`
    go_flags=$(grep -oP 'long:"[^"]+"' "$FLAGS_FILE" | sed 's/long:"//;s/"//' | sort -u)

    found=0
    # bash and zsh use --flag-name; fish uses -l flag-name
    for comp_file in "$BASH_FILE" "$ZSH_FILE"; do
        [[ -f "$comp_file" ]] || continue
        comp_name=$(basename "$comp_file")
        while IFS= read -r flag; do
            if ! grep -qF -- "--$flag" "$comp_file" 2>/dev/null; then
                warn "$comp_name  →  missing --$flag"
                found=1
            fi
        done <<< "$go_flags"
    done
    if [[ -f "$FISH_FILE" ]]; then
        comp_name=$(basename "$FISH_FILE")
        while IFS= read -r flag; do
            # fish registers long options with: complete -c cmd -l flag-name
            if ! grep -qF -- "-l $flag" "$FISH_FILE" 2>/dev/null; then
                warn "$comp_name  →  missing -l $flag"
                found=1
            fi
        done <<< "$go_flags"
    fi
    [[ $found -eq 0 ]] && ok "All completions in sync"
fi

# ── Summary ───────────────────────────────────────────────────────────────────
echo
echo "━━━ Summary ━━━"
if [[ $issues -eq 0 ]]; then
    echo -e "${GRN}No issues found.${RST}"
else
    echo -e "${YLW}${issues} issue(s) found.${RST}"
fi

[[ $STRICT -eq 1 && $issues -gt 0 ]] && exit 1
exit 0
