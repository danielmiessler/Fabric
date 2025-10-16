

# Bug Fix of {{input}} in extensions

## Problem Description

When using `{{input}}` inside extension calls within patterns, the input parameter was not being passed through correctly, resulting in empty values being sent to extensions.

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

## Root Cause

The bug was in `/internal/plugins/db/fsdb/patterns.go` in the `applyVariables()` function at line 104:

```go
if processed, err = template.ApplyTemplate(withSentinel, variables, ""); err != nil {
```

This function was calling `ApplyTemplate()` with an empty string `""` as the third parameter (input), even though the actual user input was available. The sentinel token pattern worked correctly:

1. Replace `{{input}}` in pattern with `__FABRIC_INPUT_SENTINEL_TOKEN__`
2. Process template variables (this step had the bug - passed empty input)
3. Replace sentinel token with actual input

The problem occurred in step 2: when the pattern contained `{{ext:name:op:{{input}}}}`, it would become `{{ext:name:op:__FABRIC_INPUT_SENTINEL_TOKEN__}}` after step 1. Then in step 2, `ApplyTemplate()` would see the sentinel token and replace it with the input parameter (which was empty). This resulted in `{{ext:name:op:}}` with no value.

## Alternative Workarounds Tested

### Could users bypass the bug using `-v=input:"text"`?

**Test:** Can the bug be avoided by using `-v=input:"text"` instead of piping input via stdin?

**Result:** ❌ **NO - This does NOT work**

**Testing Evidence:**

```bash
# Test 1: Unfixed version with -v=input (FAILS)
./fabric-before-fix -p ai_echo -v=input:"What is AI?"
# Result: Extensions receive empty string ""

# Test 2: Unfixed version with stdin (FAILS)  
echo "What is AI?" | ./fabric-before-fix -p ai_echo
# Result: Extensions receive empty string ""

# Test 3: Fixed version with -v=input (STILL FAILS!)
./fabric-fix -p ai_echo -v=input:"What is AI?"
# Result: Extensions receive empty string ""

# Test 4: Fixed version with stdin (WORKS)
echo "What is AI?" | ./fabric-fix -p ai_echo  
# Result: Extensions receive "What is AI?" ✓
```

**Why `-v=input:` doesn't work:**

The `{{input}}` variable is **special** and reserved for stdin content. The sentinel token system processes `{{input}}` replacements **before** variable substitution from `-v` flags occurs. This means:

1. Pattern starts: `{{ext:openai:chat:{{input}}}}`
2. Sentinel replaces `{{input}}`: `{{ext:openai:chat:__SENTINEL__}}`
3. Variable substitution happens (including `-v=input:...`)
4. Sentinel replaced with actual stdin (empty if no stdin)

The `-v=input:` flag attempts to set a variable named "input", but the sentinel token system has already replaced `{{input}}` tokens before this variable substitution occurs. The actual `input` parameter comes from stdin, not from the variables map.

**Conclusion:** 

✅ **No workaround existed** - Users had no way to avoid the bug  
✅ **The fix was essential** - Without it, `{{input}}` in extensions never worked  
✅ **Bug was fundamental** - Affected the core sentinel token system, not just input parsing

This makes the fix even more critical for the Fabric ecosystem, as extensions could not receive user input through the `{{input}}` variable at all.

## The Fix

**File:** `/Users/ourdecisions/devHelp/Fabric-fix/internal/plugins/db/fsdb/patterns.go`  
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

## Testing

### Before Fix
```bash
$ echo "Test input text" | ./fabric-fix -p ai_echo --debug=3 2>&1 | grep "Extension call"
DEBUG: Extension call: name=openai operation=chat value=
```
The `value=` is empty!

### After Fix
```bash
$ echo "Test input text" | ./fabric-fix -p ai_echo --debug=3 2>&1 | grep "Extension call"
DEBUG: Extension call: name=openai operation=chat value=Test input text
```
The `value=Test input text` is correct!

### Full Integration Test
```bash
$ echo "What is Artificial Intelligence" | ./fabric-fix -p ai_echo
Here's a summary:

Artificial Intelligence (AI) refers to the simulation of human intelligence by computer systems...
```
✅ Extension receives input and OpenAI returns proper response!

### Unit Tests
All existing tests continue to pass:
- `TestSentinelTokenReplacement` ✅
- `TestNestedInputInExtension` ✅  
- `TestSentinelInVariableProcessing` ✅
- `TestApplyVariables` ✅
- `TestGetApplyVariables` ✅

## Additional Resources

For general extension debugging techniques, see [debugging-extensions.md](./debugging-extensions.md).

## NOTES

- diff two repos with bc

- run tests and coverage before fix
go test -v -coverprofile=coverage.out ./... > bug-fix-extensions/test_results_with_coverage_before_fix.txt 2>&1;rm coverage.out

- `go mod tidy` will remove this from go.sum
 -github.com/anthropics/anthropic-sdk-go v1.12.0 h1:xPqlGnq7rWrTiHazIvCiumA0u7mGQnwDQtvA1M82h9U=
 -github.com/anthropics/anthropic-sdk-go v1.12.0/go.mod h1:WTz31rIUHUHqai2UslPpw5CwXrQP3geYBioRV4WOLvE=

- ☑️ Compare Fabric Fix and Fabric in BC - see big picture
- Bring Test over that shows it fails as is
- Learn how to debug word generator
- Compare Java to Go in table
- create word counter
- show it fails
- create a test that shows fail in general
- make the fix
- show the test passes
- Make full code execution calling claude and chat via api
- create PR together with your thoughts to Daniel Miessler

