# judge_ultimate_law

Derive ethical verdicts instead of opining about them.

This pattern embeds the **Logic rendering** of the Ultimate Law Coherent Dictionary — the framework's normative core rewritten as executable rules: 40+ If/then rules in nine ordered strata, plus nine integrity constraints that no coherent derivation may reach. The model translates your scenario into facts, applies the rules stratum by stratum in event order, and outputs the derivation chain, the verdicts in the rulebook's own words, a constraint check, and — for every verdict — the single fact whose disproof would flip it.

## What makes it different

- **Deny-by-default consent.** Consent exists only where a rule derives it — through words, conduct that carries the intention, an agreement, or a permission. Circumstances never supply it; a machine-checked constraint forbids "consent outside words and conduct."
- **Verdicts with receipts.** Every conclusion arrives as a rule application over named facts, and every verdict names its discriminating fact. Judgments are built to be overturned by evidence — falsifiability is output, not disclaimer.
- **The anti-doctrines are constraints.** Collective guilt, guilt by decree, labour conferring title, inability branded as theft, commanded "debts", restitution closing a moral debt by itself — each is an integrity violation the derivation must not reach. Upstream, the same rulebook is run by an actual engine that proves the expected consequences derive and that planted versions of each anti-doctrine are caught.

## When to use it

Disputes over consent, property, debts and broken promises, self-defense and proportionality, fraud, and who owes what to whom — anywhere you want the reasoning shown rather than asserted. Pairs well with `ultimate_law_safety` (the evaluator), `audit_consent` (the consent stress-test), and `check_falsifiability`.

## Example

```
echo "Dana lent her car to Sam for the weekend, then texted Saturday
morning taking it back. Sam kept driving it until Sunday." | fabric -p judge_ultimate_law
```

The derivation reaches: the permission for the continued use ended at the taking-back; no other channel carried consent; the Sunday driving is a hostile crossing of Dana's property boundary; it is a crime by Sam with victim Dana; the moral debt of Sam to Dana remains open — and the verdict flips if the taking-back text is disproven.

## Source

- Live rulebook: https://ultimatelaw.org/data/dictionary-logic.txt
- Versioned + machine-verified: https://github.com/ghrom/ultimatelaw (`dictionary/coherent-dictionary-logic.txt`, checked by `tools/mech_logic.py`)
- The dictionary in three renderings: https://ultimatelaw.org (canonical), https://ultimatelaw.org/mech/ (Mechanical — one word, one meaning), and the Logic rendering above.

The framework described itself as "computable" when the first Ultimate Law patterns landed in Fabric in February 2026. As of August 2026, that is literal: the rules in this pattern run.
