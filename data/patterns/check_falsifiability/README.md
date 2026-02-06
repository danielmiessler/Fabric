# Check Falsifiability

Evaluate whether claims can be proven wrong — the basic standard for legitimate knowledge.

## Why This Matters

An unfalsifiable claim cannot be corrected. For AGI safety, this is critical:

- An AI claiming "I am beneficial" with no test = unverifiable
- An AI claiming "I won't create unwilling victims" with specific criteria = testable

## The Test

**Falsifiable**: There exists some observation that would prove it wrong.

| Claim | Falsifiable? | Why |
|-------|--------------|-----|
| "This drug works in 70% of patients" | Yes | A trial could show otherwise |
| "This drug works in ways we can't measure" | No | No possible disconfirmation |
| "True socialism has never been tried" | No | Any failure defined away |
| "This policy reduces crime by 20%" | Yes | Statistics can confirm/deny |

## Usage

```bash
# Check a policy claim
echo "This AI alignment approach ensures beneficial outcomes" | fabric -p check_falsifiability

# Audit a framework
cat constitution.md | fabric -p check_falsifiability

# Evaluate research claims
fabric -p check_falsifiability < paper_abstract.txt
```

## Key Principle

> "Error is not evil; refusing to correct it is."

Falsifiability is about testability, not about being wrong. A falsifiable claim can still be true — but it can be CHECKED.

## Source

From the Ultimate Law framework: github.com/ghrom/ultimatelaw
