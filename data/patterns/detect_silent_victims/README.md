# Detect Silent Victims

Identify parties harmed by an action or system who cannot speak up — because they don't exist yet, lack power, lack awareness, or lack voice.

## Why This Matters

"No victim, no crime" is a powerful principle. But it has a critical blind spot: what about victims who can't report their victimhood?

| Category | Example |
|----------|---------|
| Future victims | Climate damage, national debt, resource depletion |
| Voiceless victims | Children, animals, ecosystems, marginalized communities |
| Unaware victims | Data exploitation, slow poisoning, hidden externalities |
| Diffuse victims | Pollution affecting millions trivially each, market manipulation |
| Structural victims | Systems that consistently produce losers by design |

The absence of a complaint is not evidence of the absence of a victim.

## Origin

This pattern emerged from a cross-model AI evaluation where 19 AI systems identified "silent victims" as the most critical gap in the Ultimate Law framework. DeepSeek-R1 proposed "future generations as victims." The devil's advocate (cogito:70b) scored this weakness at 9/10.

## Usage

```bash
# Audit a policy proposal
echo "Build a coal plant to provide cheap energy" | fabric -p detect_silent_victims

# Evaluate a business model
echo "Offer free service funded by selling user data" | fabric -p detect_silent_victims

# Check an AI system
echo "Train AI on scraped internet data" | fabric -p detect_silent_victims

# Audit legislation
cat proposed_law.txt | fabric -p detect_silent_victims
```

## The Reversed Test

> "If every silent victim could speak with equal power, would they consent to this?"

This single question exposes most hidden harm.

## Source

From the Ultimate Law framework: github.com/ghrom/ultimatelaw
Developed after cross-model AI dialogue series (19 models, 10+ organizations, 2026)
