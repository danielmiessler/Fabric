# Compress input before you run a pattern

Fabric reads its input from `stdin`. Thus you can put any compression tool in the pipe before `fabric`. Fabric does not need a special flag for this:

```bash
cat long_document.txt | compress "main ideas and claims" | fabric --pattern summarize
```

This guide shows how to make a `compress` command with [SuperCompress](https://github.com/arjunkshah/supercompress). You can use a different tool in the same position.

## Before you start

SuperCompress sends your text to a hosted API at `supercompress.vercel.app`. Do not use it for text that must stay on your computer.

Compression removes text. For legal or audit text, or for patterns that need the full text, do not compress the input.

## Make the `compress` command

1. Install the client:

   ```bash
   pip install supercompress
   ```

2. Get an API key from the [SuperCompress dashboard](https://supercompress.dev/dashboard).
3. Set the key in your shell profile:

   ```bash
   export SUPERCOMPRESS_API_KEY=sc_live_YOUR_KEY
   ```

4. Save this script as `~/.local/bin/compress`:

   ```python
   #!/usr/bin/env python3
   """Compress stdin with SuperCompress. The argument is the query to keep."""
   import sys

   from supercompress.client import SuperCompress

   if len(sys.argv) < 2:
       sys.exit('usage: compress "what to keep" < input.txt')
   result = SuperCompress().compress(context=sys.stdin.read(), query=" ".join(sys.argv[1:]))
   print(result.compressed_text)
   print(f"compress: {result.original_tokens} -> {result.kept_tokens} tokens", file=sys.stderr)
   ```

5. Make the script executable:

   ```bash
   chmod +x ~/.local/bin/compress
   ```

## Use it

SuperCompress keeps the text that is related to the query. Write a query that agrees with the goal of the pattern:

```bash
fabric -y "https://youtube.com/watch?v=uXs-zPc63kM" | compress "key ideas, insights, and recommendations" | fabric --pattern extract_wisdom
pbpaste | compress "factual claims and their evidence" | fabric --pattern analyze_claims
```

The script writes the token counts to `stderr`. Thus the counts do not go into the pattern input.
