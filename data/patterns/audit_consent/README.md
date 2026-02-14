# Audit Consent

Determine whether "consent" in an interaction is genuine or manufactured through power asymmetries.

## Why This Matters

"I agreed to it" is the most common defense for exploitative arrangements. But consent requires more than a signature or a click:

- **Information**: Do you understand what you're agreeing to?
- **Alternatives**: Can you meaningfully say no?
- **No manipulation**: Is the framing honest?
- **Revocability**: Can you change your mind?
- **Capacity**: Can you assess the consequences?

If any of these are absent, "consent" is theater — not agreement.

## Origin

This pattern emerged from a cross-model AI evaluation where 19 different AI systems stress-tested the Ultimate Law ethical framework. The devil's advocate (cogito:70b) scored "consent theater" at 9/10 — the strongest attack in the series. The framework survived, but identified consent verification as its most critical gap.

## Usage

```bash
# Audit terms of service
cat tos.txt | fabric -p audit_consent

# Evaluate an employment contract
echo "Employee agrees to mandatory arbitration and non-compete" | fabric -p audit_consent

# Check a policy proposal
echo "Citizens consent to taxation through democratic participation" | fabric -p audit_consent

# Audit AI data collection
echo "Users agree to data collection by using the service" | fabric -p audit_consent
```

## The Verdict Scale

| Verdict | Meaning |
|---------|---------|
| GENUINE | All five tests pass, low power asymmetry |
| PRESSURED BUT FUNCTIONAL | Some asymmetry, but refusal is possible |
| MANUFACTURED | Appearance of choice masks predetermined outcome |
| COERCED | Refusal carries disproportionate penalty |
| ILLUSORY | No meaningful alternative exists |

## Source

From the Ultimate Law framework: github.com/ghrom/ultimatelaw
Developed after cross-model AI dialogue series (19 models, 10+ organizations, 2026)
