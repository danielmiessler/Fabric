# Audit Transparency

Evaluate whether decisions or systems that affect others are explainable in terms those affected can understand.

## Why This Matters

Opacity combined with power is coercion's favorite disguise. When the powerful are opaque to the powerless:
- "Consent" becomes meaningless (you can't consent to what you don't understand)
- Accountability becomes impossible (you can't challenge what you can't see)
- Correction becomes blocked (errors hide behind complexity)

## Origin

Transparency was the #1 gap identified by consensus across 5+ AI models when 19 systems evaluated the Ultimate Law ethical framework (2026). Proposed as the 8th principle: "Every decision affecting others must be explainable in terms the affected party can understand."

## Five Transparency Dimensions

| Dimension | Question |
|-----------|----------|
| Decision | Can affected parties see how decisions are made? |
| Algorithmic | Can system behavior be explained in plain language? |
| Financial | Are costs, fees, and flows visible? |
| Governance | Are rules visible before they take effect? |
| Data | Do people know what's collected and how it's used? |

## Usage

```bash
# Audit an AI system
echo "GPT-4 determines loan eligibility" | fabric -p audit_transparency

# Evaluate a policy
echo "Content moderation decisions are made by automated systems" | fabric -p audit_transparency

# Check a contract
cat employment_contract.txt | fabric -p audit_transparency

# Audit governance
echo "Platform rules can change at any time without notice" | fabric -p audit_transparency
```

## The Reversal Test

> "Would the decision-maker accept this level of opacity if they were the affected party?"

## Source

From the Ultimate Law framework: github.com/ghrom/ultimatelaw
Developed after cross-model AI dialogue series (19 models, 10+ organizations, 2026)
