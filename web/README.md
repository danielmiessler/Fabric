# Fabric Web App

A user-friendly web interface for [Fabric](https://github.com/danielmiessler/Fabric) built with [Svelte](https://svelte.dev/), [Skeleton UI](https://www.skeleton.dev/), and [Mdsvex](https://mdsvex.pngwn.io/).

![Fabric Web App Preview](../docs/images/svelte-preview.png)
*Alt: Screenshot of the Fabric web app dashboard showing pattern inputs and outputs.*

## Table of Contents

- [Fabric Web App](#fabric-web-app)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Running the App](#running-the-app)
    - [Prerequisites](#prerequisites)
    - [Launch the Svelte App](#launch-the-svelte-app)
  - [Streamlit UI](#streamlit-ui)
    - [Key Features](#key-features)
    - [Setup and Run](#setup-and-run)
  - [Obsidian Integration](#obsidian-integration)
    - [Quick Setup](#quick-setup)
  - [Contributing](#contributing)

## Installation

> [!NOTE]
> Requires Node.js ≥18 and Fabric installed globally (`fabric --version` to check).

From the Fabric root directory:

**Using npm:**

```bash
./web/scripts/npm-install.sh
```

**Or using pnpm (recommended for speed):**

```bash
./web/scripts/pnpm-install.sh
```

These scripts install Svelte dependencies and patch PDF-to-Markdown libraries (e.g., pdfjs-dist, pdf-to-markdown). Link to scripts:[npm-install.sh](./scripts/npm-install.sh) and [pnpm-install.sh](./scripts/pnpm-install.sh)

## Running the App

### Prerequisites

Start Fabric's server in a separate terminal:

```bash
fabric --serve
```

(This exposes Fabric's API at <http://localhost:8080>)

### Launch the Svelte App

In the `web/` directory:

**Using npm:**

```bash
npm run dev
```

**Or using pnpm:**

```bash
pnpm run dev
```

Visit [http://localhost:5173](http://localhost:5173) (default port).

> [!TIP]
>
> Sync Svelte types if needed: `npx svelte-kit sync`

## Streamlit UI

For Python enthusiasts, this alternative UI excels at data visualization and chaining complex patterns. It supports clipboard ops across platforms (install pyperclip on Windows, xclip on Linux).

- **macOS**: Uses `pbcopy` and `pbpaste` (built-in)
- **Windows**: Uses `pyperclip` library (install with `pip install pyperclip`)
- **Linux**: Uses `xclip` (install with `sudo apt-get install xclip` or equivalent for your Linux distribution)

### Key Features

<!-- - Running and chaining patterns
- Managing pattern outputs
- Creating and editing patterns
- Analyzing pattern results -->

- Run and edit patterns with real-time previews.
- Analyze outputs with charts (via Matplotlib/Seaborn).
- Export results to Markdown or CSV.

### Setup and Run

From `web/`:

```bash
pip install -r requirements.txt #Or: pip install streamlit pandas matplotlib seaborn numpy python-dotenv pyperclip
streamlit run streamlit.py
```

Access at [http://localhost:8501](http://localhost:8501) (default port).

## Obsidian Integration

Turn `web/src/lib/content/` into an [Obsidian](https://obsidian.md) vault for note-taking synced with Fabric patterns. It includes pre-configured `.obsidian/` and `templates/` folders.

### Quick Setup

1. Open Obsidian: File > Open folder as vault > Select `web/src/lib/content/`
2. To publish posts, move them to the posts directory (`web/src/lib/content/posts`).
3. Use Fabric patterns to generate content directly in Markdown files.

> [!TIP]
>
> When creating new posts, make sure to include a date (YYYY-MM-DD), description, tags (e.g., #ai #patterns), and aliases for SEO. Only a date is needed to display a note. Embed images `(![alt](path))`, link patterns `([[pattern-name]])`, or code blocks for reusable snippets—all in standard Markdown.

## Contributing

Refer to the [Contributing Guide](/docs/CONTRIBUTING.md) for details on how to improve this content.
