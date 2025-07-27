# create_mnemonic_phrases

Generate short, memorable sentences that embed Diceware‑style words **unchanged and in order**. This pattern is ideal for turning a raw Diceware word list into phrases that are easier to recall while preserving the exact secret.

## What is Diceware?

Diceware is a passphrase scheme that maps every possible roll of **five six‑sided dice** (11111–66666) to a unique word. Because there are `6^5 = 7776` combinations, the canonical list contains the same number of entries.

### Entropy of the standard 7776‑word list

```text
words = 7776
entropy_per_word = log2(words) ≈ 12.925 bits
```

A passphrase that strings *N* independently chosen words together therefore carries `N × 12.925 bits` of entropy—≈ 77.5 bits for six words, ≈ 129 bits for ten, and so on. Four or more words already outclass most human‑made passwords.

## Pattern overview

The accompanying **`system.md`** file instructs Fabric to:

1. Echo the supplied words back in **bold**, separated by commas.
2. Generate **five** distinct, short sentences that include the words **in the same order and spelling**, enabling rapid rote learning or spaced‑repetition drills.

The output is deliberately minimalist—no extra commentary—so you can pipe it straight into other scripts.

## Quick start

```bash
# 1  Pick five random words from any Diceware‑compatible list
shuf -n 5 diceware_wordlist.txt | \
  # 2  Feed them to Fabric with this pattern
  fabric --pattern create_mnemonic_phrases -s
```

You’ll see the words echoed in bold, followed by five candidate mnemonic sentences ready for memorisation.

