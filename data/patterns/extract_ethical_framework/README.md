# Extract Ethical Framework

Make implicit ethics explicit. Every prescriptive document contains hidden ethical assumptions — this pattern surfaces them.

## Why This Matters

- Terms of Service contain implicit ethics
- AI system descriptions contain implicit ethics
- Policies and laws contain implicit ethics
- Manifestos contain implicit ethics

Making them explicit allows:
1. Checking for internal consistency
2. Evaluating against minimal ethical standards
3. Identifying hidden coercion
4. Challenging unstated assumptions

## What It Extracts

| Question | What to Find |
|----------|--------------|
| Who counts? | Whose interests matter? |
| What's harm? | What are agents protected from? |
| What's consent? | How is agreement established? |
| Who decides? | Who has authority and why? |
| When is force OK? | What justifies coercion? |
| What wins? | Hierarchy when values conflict |

## Usage

```bash
# Analyze terms of service
cat tos.txt | fabric -p extract_ethical_framework

# Analyze an AI safety proposal
echo "The AI should be beneficial and aligned with human values" | fabric -p extract_ethical_framework

# Audit a policy document
fabric -p extract_ethical_framework < policy.md
```

## The Minimal Standard

Does the framework authorize creating unwilling victims?

If yes → it fails the minimal ethics test, regardless of how coherent it is internally.

## Source

From the Ultimate Law framework: github.com/ghrom/ultimatelaw
