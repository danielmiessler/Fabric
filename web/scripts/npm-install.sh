#!/bin/bash

cd "$(dirname "$0")/.." || exit

if command -v npm &>/dev/null; then
  echo "npm is installed"
else
  echo "npm is not installed. Please install npm first."
  exit 1
fi

# Install the GUI and its dependencies
npm install
# Install PDF-to-Markdown components in this order
npm install -D patch-package
npm install -D pdfjs-dist
npm install -D github:jzillmann/pdf-to-markdown#modularize

npx svelte-kit sync
