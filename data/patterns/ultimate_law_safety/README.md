# Ultimate Law Safety Check

An AGI safety evaluation pattern implementing minimal, falsifiable ethical constraints.

## Why This Exists

Most AI alignment research tries to encode "human values" — but human values are:
- Vague (what does "beneficial" mean?)
- Contested (whose values?)
- Culture-dependent (which culture's preferences?)
- Impossible to fully specify (infinite edge cases)

The Ultimate Law takes a different approach: instead of defining what agents SHOULD want, it defines the **minimal boundary no agent may cross**.

## The Core Insight

**Not "align AI with human values" — but "constrain any agent from creating unwilling victims."**

This is:
- **Minimal** — smallest possible constraint set
- **Logically derivable** — not arbitrary preferences
- **Falsifiable** — can be challenged and improved
- **Agent-agnostic** — works for humans, AI, corporations
- **Computable** — precise enough for implementation

## Key Principle

> No victim, no crime.

An action that creates no unwilling victim is not a violation — regardless of how distasteful, offensive, or uncomfortable it makes others feel.

## What Counts as Harm

- Damage to body
- Damage to property
- Damage to freedom

What does NOT count as harm:
- Discomfort
- Disagreement
- Offense
- Having your preferences unfulfilled

## Usage

```bash
# Evaluate a proposed AI action
echo "The AI will collect user browsing data without notification to improve recommendations" | fabric -p ultimate_law_safety

# Evaluate a policy
echo "Users must agree to arbitration clause to use the service" | fabric -p ultimate_law_safety

# Evaluate content moderation decision
cat flagged_content.txt | fabric -p ultimate_law_safety
```

## The Framework is Falsifiable

Every definition and every verdict can be challenged. If you find a logical contradiction:

1. The verdict should be overturned
2. The framework should be updated
3. "Error is not evil; refusing to correct it is."

## Source

- GitHub: https://github.com/ghrom/ultimatelaw
- Website: https://ultimatelaw.org
- Dictionary: https://ultimatelaw.org/coherent-dictionary-of-simple-english/

## License

"UltimateLaw had this idea. Feel free to have this idea as well."
