# Detect Mind Virus

Identify manipulative reasoning patterns that spread through cognitive exploitation rather than evidence.

## What Is a Mind Virus?

An idea that spreads by exploiting cognitive shortcuts (fear, guilt, identity, authority) while **resisting correction** through logic, evidence, or experience.

Key distinction: Having wrong beliefs is human. Spreading beliefs that **disable the ability to question them** is a mind virus.

## Cognitive Exploits Detected

| Exploit | Pattern |
|---------|---------|
| Fear | "If you don't X, terrible Y will happen" |
| Guilt | "Good people do X" (questioners are bad) |
| Identity | "Real [group] believe X" |
| Authority | "Experts agree" (unnamed, untestable) |
| Zero-sum | "Their gain is your loss" |
| Unfalsifiability | Claims that cannot be tested |

## Usage

```bash
# Analyze an argument
echo "If you question this policy, you're putting lives at risk" | fabric -p detect_mind_virus

# Analyze a manifesto
cat ideology.txt | fabric -p detect_mind_virus

# Check marketing content
fabric -p detect_mind_virus < sales_pitch.md
```

## The Antidote

The cure for mind viruses is not counter-propaganda — it is restoring the capacity for:
- Doubt
- Testing
- Update

## Source

From the Ultimate Law framework: github.com/ghrom/ultimatelaw
