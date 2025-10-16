# Pull Request: Fix input parameter bug in extension calls

## 🎯 Quick Summary

**One-line fix** for a bug where `{{input}}` inside extension calls was receiving empty values instead of actual user input.

**Changed:** 1 line in `patterns.go` (line 104)  
**Impact:** Extensions can now receive user input correctly  
**Tests:** +13 new tests, all 468 tests pass

---

## 📋 PR Title

```
Fix: Pass input parameter to extension calls with {{input}} in patterns
```

---

## 📝 PR Description

```markdown
## Problem

When using `{{input}}` inside extension calls within patterns (e.g., `{{ext:openai:chat:{{input}}}}`), the extension was receiving an empty value instead of the actual user input.

### Example Bug Behavior

**Pattern:** `ai_echo/system.md`
```
Summarize:

{{ext:openai:chat:{{input}}}}
```

**Command:**
```bash
echo "What is Artificial Intelligence" | fabric -p ai_echo
```

**Expected:** Extension receives "What is Artificial Intelligence"  
**Actual (before fix):** Extension receives empty string ""

---

## Root Cause

The bug was in `/internal/plugins/db/fsdb/patterns.go` in the `applyVariables()` function at line 104:

```go
if processed, err = template.ApplyTemplate(withSentinel, variables, ""); err != nil {
```

This function was calling `ApplyTemplate()` with an empty string `""` as the third parameter (input), even though the actual user input was available. 

The sentinel token pattern worked correctly:
1. Replace `{{input}}` in pattern with `__FABRIC_INPUT_SENTINEL_TOKEN__`
2. Process template variables (this step had the bug - passed empty input)
3. Replace sentinel token with actual input

The problem occurred in step 2: when the pattern contained `{{ext:name:op:{{input}}}}`, it would become `{{ext:name:op:__FABRIC_INPUT_SENTINEL_TOKEN__}}` after step 1. Then in step 2, `ApplyTemplate()` would see the sentinel token and replace it with the input parameter (which was empty). This resulted in `{{ext:name:op:}}` with no value.

---

## Solution

**File:** `/internal/plugins/db/fsdb/patterns.go`  
**Line:** 104

Changed:
```go
if processed, err = template.ApplyTemplate(withSentinel, variables, ""); err != nil {
```

To:
```go
if processed, err = template.ApplyTemplate(withSentinel, variables, input); err != nil {
```

This ensures that when `ApplyTemplate()` encounters the sentinel token inside extension value parameters, it has the actual user input available to substitute.

---

## Testing Evidence

### Before Fix
```bash
$ echo "Test input text" | ./fabric-fix -p ai_echo --debug=3 2>&1 | grep "Extension call"
DEBUG: Extension call: name=openai operation=chat value=
```
❌ The `value=` is empty!

### After Fix
```bash
$ echo "Test input text" | ./fabric-fix -p ai_echo --debug=3 2>&1 | grep "Extension call"
DEBUG: Extension call: name=openai operation=chat value=Test input text
```
✅ The `value=Test input text` is correct!

### Full Integration Test
```bash
$ echo "What is Artificial Intelligence" | ./fabric-fix -p ai_echo
Here's a summary:

Artificial Intelligence (AI) refers to the simulation of human intelligence by computer systems...
```
✅ Extension receives input and OpenAI returns proper response!

---

## Test Results

### Test Statistics
- **Before Fix:** 455 tests pass (baseline)
- **After Fix:** 468 tests pass (+13 new tests)
- **Result:** ✅ All tests pass, no regressions

### New Tests Added (13 total)

**Sentinel Token Tests:**
1. `TestSentinelTokenReplacement` - Main test with 4 subtests:
   - `sentinel_token_with_{{input}}_in_extension_value`
   - `direct_input_variable_replacement`
   - `sentinel_with_complex_input`
   - `multiple_words_in_input`

2. `TestSentinelInVariableProcessing` - Variable processing with 3 subtests:
   - `input_variable_works_normally`
   - `multiple_input_references`
   - `input_with_variables`

3. `TestExtensionValueWithSentinel` - Sentinel in extension values
4. `TestNestedInputInExtension` - The core bug test case
5. `TestMultipleExtensionsWithInput` - Multiple extension calls
6. `TestExtensionValueMixedInputAndVariable` - Mixed input and variables

All existing tests continue to pass - **no regressions**.

---

## Files Changed

### Core Fix
- `internal/plugins/db/fsdb/patterns.go` - **1 line fix** (line 104)
- `internal/plugins/template/template.go` - Added debug logging to show input parameter

### Tests
- `internal/plugins/template/template_sentinel_test.go` - Added comprehensive test suite (379 lines, 13 tests)
- `internal/plugins/template/template_extension_multiple_test.go` - Tests multiple extension calls
- `internal/plugins/template/template_extension_mixed_test.go` - Tests mixed input/variables

### Documentation & Testing Tools
- `bug-fix-extensions/input-extension-bug-fix.md` - Detailed bug analysis and fix documentation
- `bug-fix-extensions/debugging-extensions.md` - Debugging guide for extensions
- `bug-fix-extensions/test_results_with_coverage_before_fix.txt` - Baseline test results (455 tests)
- `bug-fix-extensions/test_results_with_coverage_after_fix.txt` - Post-fix test results (468 tests)
- `bug-fix-extensions/setup_openai_extension_test.sh` - OpenAI extension testing script
- `bug-fix-extensions/word-generator-counter.sh` - Extension automation example

---

## Additional Context

### Full Documentation
- **Bug Analysis:** [bug-fix-extensions/input-extension-bug-fix.md](bug-fix-extensions/input-extension-bug-fix.md)
- **Debugging Guide:** [bug-fix-extensions/debugging-extensions.md](bug-fix-extensions/debugging-extensions.md)

### Why This Matters
This bug prevented users from:
- Creating patterns that call external LLM APIs (OpenAI, Anthropic, etc.) via extensions
- Building reusable AI workflows with extensions
- Passing dynamic input to extensions in patterns

Now users can create powerful patterns like:
```markdown
# AI consensus pattern
Compare these responses:
- OpenAI: {{ext:openai:chat:{{input}}}}
- Anthropic: {{ext:anthropic:chat:{{input}}}}
- Local: {{ext:ollama:chat:{{input}}}}
```

### Development Process
1. ✅ Identified bug through user-reported issue
2. ✅ Created comprehensive test suite that initially failed
3. ✅ Implemented minimal fix (1 line change)
4. ✅ Verified all tests pass
5. ✅ Tested with real-world OpenAI integration
6. ✅ Documented thoroughly

---

## Checklist

- [x] Bug fix addresses the root cause
- [x] All existing tests pass (no regressions)
- [x] Added comprehensive test coverage (13 new tests)
- [x] Real-world integration testing completed
- [x] Documentation includes before/after examples
- [x] Code change is minimal and focused
- [x] Debug logging added for troubleshooting
- [x] Test results included (before/after)

---

## Related Issues

This fixes the issue where extensions cannot receive user input when called from patterns with `{{input}}` in the extension value parameter.

---

## Breaking Changes

None. This is a bug fix that makes the system work as originally intended.

```

---

## 🚀 How to Submit

### Option 1: GitHub Web UI (Recommended)

1. Go to: https://github.com/nickarino/Fabric
2. Click "Contribute" → "Open pull request"
3. Target: `danielmiessler/Fabric:main` ← `nickarino/Fabric:main`
4. Copy the title from above
5. Copy the entire PR Description section from above
6. Click "Create pull request"

### Option 2: GitHub CLI

```bash
cd /Users/ourdecisions/devHelp/Fabric-fix

gh pr create \
  --repo danielmiessler/Fabric \
  --base main \
  --head nickarino:main \
  --title "Fix: Pass input parameter to extension calls with {{input}} in patterns" \
  --body-file bug-fix-extensions/PR-SUBMISSION.md
```

### Option 3: Using git (if gh CLI has issues)

1. Push your branch (already done ✓)
2. Visit: https://github.com/danielmiessler/Fabric/compare/main...nickarino:Fabric:main
3. Click "Create pull request"
4. Use the PR title and description from this file

---

## 📊 Commits Included

```
fffa5a73 test: Add test results after fixing input parameter bug
08f18651 docs: Add cross-references and rename documentation files
9eea6830 Fix: Pass input parameter to ApplyTemplate in applyVariables
7acbb519 Move debug_instructions.md to bug-fix-extensions folder
```

---

## ✅ Ready to Submit!

All changes are pushed to your fork and ready for PR submission.
