#!/bin/bash

cd "$(dirname "$0")/.." || exit

if command -v npm &>/dev/null; then
  echo "pnpm is installed"
else
  echo "pnpm is not installed. Please install pnpm first."
  exit 1
fi

# Install the GUI and its dependencies
pnpm install
# Install PDF-to-Markdown components in this order
pnpm install -D patch-package
pnpm install -D pdfjs-dist
pnpm install -D github:jzillmann/pdf-to-markdown#modularize

pnpm exec svelte-kit sync
