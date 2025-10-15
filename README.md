<div align="center">
    <a href="https://go.warp.dev/fabric" target="_blank">
        <sup>Special thanks to:</sup>
        <br>
        <img alt="Warp sponsorship" width="400" src="https://raw.githubusercontent.com/warpdotdev/brand-assets/refs/heads/main/Github/Sponsor/Warp-Github-LG-02.png">
        <br>
        <h>Warp, built for coding with multiple AI agents</b>
        <br>
        <sup>Available for macOS, Linux and Windows</sup>
    </a>
</div>

<br>

<div align="center">

<img src="./docs/images/fabric-logo-gif.gif" alt="fabriclogo" width="400" height="400"/>

# `fabric`

![Static Badge](https://img.shields.io/badge/mission-human_flourishing_via_AI_augmentation-purple)
<br />
![GitHub top language](https://img.shields.io/github/languages/top/danielmiessler/fabric)
![GitHub last commit](https://img.shields.io/github/last-commit/danielmiessler/fabric)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/danielmiessler/fabric)

<div align="center">
<h4><code>fabric</code> is an open-source framework for augmenting humans using AI.</h4>
</div>

![Screenshot of fabric](./docs/images/fabric-summarize.png)

</div>

[Updates](#updates) •
[What and Why](#what-and-why) •
[Philosophy](#philosophy) •
[Installation](#installation) •
[Usage](#usage) •
[Examples](#examples) •
[Just Use the Patterns](#just-use-the-patterns) •
[Custom Patterns](#custom-patterns) •
[Helper Apps](#helper-apps) •
[Meta](#meta)

</div>

## What and why

Since the start of modern AI in late 2022 we've seen an **_extraordinary_** number of AI applications for accomplishing tasks. There are thousands of websites, chat-bots, mobile apps, and other interfaces for using all the different AI out there.

It's all really exciting and powerful, but _it's not easy to integrate this functionality into our lives._

<div class="align center">
<h4>In other words, AI doesn't have a capabilities problem—it has an <em>integration</em> problem.</h4>
</div>

**Fabric was created to address this by creating and organizing the fundamental units of AI—the prompts themselves!**

Fabric organizes prompts by real-world task, allowing people to create, collect, and organize their most important AI solutions in a single place for use in their favorite tools. And if you're command-line focused, you can use Fabric itself as the interface!

## Updates

<details>
<summary>Click to view recent updates</summary>

Dear Users,

We've been doing so many exciting things here at Fabric, I wanted to give a quick summary here to give you a sense of our development velocity!

Below are the **new features and capabilities** we've added (newest first):

### Recent Major Features

- [v1.4.311](https://github.com/danielmiessler/fabric/releases/tag/v1.4.311) (Sep 13, 2025) — **More internationalization support**: Adds de (German), fa (Persian / Farsi), fr (French), it (Italian),
  ja (Japanese), pt (Portuguese), zh (Chinese)
- [v1.4.309](https://github.com/danielmiessler/fabric/releases/tag/v1.4.309) (Sep 9, 2025) — **Comprehensive internationalization support**: Includes English and Spanish locale files.
- [v1.4.303](https://github.com/danielmiessler/fabric/releases/tag/v1.4.303) (Aug 29, 2025) — **New Binary Releases**: Linux ARM and Windows ARM targets. You can run Fabric on the Raspberry PI and on your Windows Surface!
- [v1.4.294](https://github.com/danielmiessler/fabric/releases/tag/v1.4.294) (Aug 20, 2025) — **Venice AI Support**: Added the Venice AI provider. Venice is a Privacy-First, Open-Source AI provider. See their ["About Venice"](https://docs.venice.ai/overview/about-venice) page for details.
- [v1.4.291](https://github.com/danielmiessler/fabric/releases/tag/v1.4.291) (Aug 18, 2025) — **Speech To Text**: Add OpenAI speech-to-text support with `--transcribe-file`, `--transcribe-model`, and `--split-media-file` flags.
- [v1.4.287](https://github.com/danielmiessler/fabric/releases/tag/v1.4.287) (Aug 16, 2025) — **AI Reasoning**: Add Thinking to Gemini models and introduce `readme_updates` python script
- [v1.4.286](https://github.com/danielmiessler/fabric/releases/tag/v1.4.286) (Aug 14, 2025) — **AI Reasoning**: Introduce Thinking Config Across Anthropic and OpenAI Providers
- [v1.4.285](https://github.com/danielmiessler/fabric/releases/tag/v1.4.285) (Aug 13, 2025) — **Extended Context**: Enable One Million Token Context Beta Feature for Sonnet-4
- [v1.4.284](https://github.com/danielmiessler/fabric/releases/tag/v1.4.284) (Aug 12, 2025) — **Easy Shell Completions Setup**: Introduce One-Liner Curl Install for Completions
- [v1.4.283](https://github.com/danielmiessler/fabric/releases/tag/v1.4.283) (Aug 12, 2025) — **Model Management**: Add Vendor Selection Support for Models
- [v1.4.282](https://github.com/danielmiessler/fabric/releases/tag/v1.4.282) (Aug 11, 2025) — **Enhanced Shell Completions**: Enhanced Shell Completions for Fabric CLI Binaries
- [v1.4.281](https://github.com/danielmiessler/fabric/releases/tag/v1.4.281) (Aug 11, 2025) — **Gemini Search Tool**: Add Web Search Tool Support for Gemini Models
- [v1.4.278](https://github.com/danielmiessler/fabric/releases/tag/v1.4.278) (Aug 9, 2025) — **Enhance YouTube Transcripts**: Enhance YouTube Support with Custom yt-dlp Arguments
- [v1.4.277](https://github.com/danielmiessler/fabric/releases/tag/v1.4.277) (Aug 8, 2025) — **Desktop Notifications**: Add cross-platform desktop notifications to Fabric CLI
- [v1.4.274](https://github.com/danielmiessler/fabric/releases/tag/v1.4.274) (Aug 7, 2025) — **Claude 4.1 Added**: Add Support for Claude Opus 4.1 Model
- [v1.4.271](https://github.com/danielmiessler/fabric/releases/tag/v1.4.271) (Jul 28, 2025) — **AI Summarized Release Notes**: Enable AI summary updates for GitHub releases
- [v1.4.268](https://github.com/danielmiessler/fabric/releases/tag/v1.4.268) (Jul 26, 2025) — **Gemini TTS Voice Selection**: add Gemini TTS voice selection and listing functionality
- [v1.4.267](https://github.com/danielmiessler/fabric/releases/tag/v1.4.267) (Jul 26, 2025) — **Text-to-Speech**: Update Gemini Plugin to New SDK with TTS Support
- [v1.4.258](https://github.com/danielmiessler/fabric/releases/tag/v1.4.258) (Jul 17, 2025) — **Onboarding Improved**: Add startup check to initialize config and .env file automatically
- [v1.4.257](https://github.com/danielmiessler/fabric/releases/tag/v1.4.257) (Jul 17, 2025) — **OpenAI Routing Control**: Introduce CLI Flag to Disable OpenAI Responses API
- [v1.4.252](https://github.com/danielmiessler/fabric/releases/tag/v1.4.252) (Jul 16, 2025) — **Hide Thinking Block**: Optional Hiding of Model Thinking Process with Configurable Tags
- [v1.4.246](https://github.com/danielmiessler/fabric/releases/tag/v1.4.246) (Jul 14, 2025) — **Automatic ChangeLog Updates**: Add AI-powered changelog generation with high-performance Go tool and comprehensive caching
- [v1.4.245](https://github.com/danielmiessler/fabric/releases/tag/v1.4.245) (Jul 11, 2025) — **Together AI**: Together AI Support with OpenAI Fallback Mechanism Added
- [v1.4.232](https://github.com/danielmiessler/fabric/releases/tag/v1.4.232) (Jul 6, 2025) — **Add Custom**: Add Custom Patterns Directory Support
- [v1.4.231](https://github.com/danielmiessler/fabric/releases/tag/v1.4.231) (Jul 5, 2025) — **OAuth Auto-Auth**: OAuth Authentication Support for Anthropic (Use your Max Subscription)
- [v1.4.230](https://github.com/danielmiessler/fabric/releases/tag/v1.4.230) (Jul 5, 2025) — **Model Management**: Add advanced image generation parameters for OpenAI models with four new CLI flags
- [v1.4.227](https://github.com/danielmiessler/fabric/releases/tag/v1.4.227) (Jul 4, 2025) — **Add Image**: Add Image Generation Support to Fabric
- [v1.4.226](https://github.com/danielmiessler/fabric/releases/tag/v1.4.226) (Jul 4, 2025) — **Web Search**: OpenAI Plugin Now Supports Web Search Functionality
- [v1.4.225](https://github.com/danielmiessler/fabric/releases/tag/v1.4.225) (Jul 4, 2025) — **Web Search**: Runtime Web Search Control via Command-Line `--search` Flag
- [v1.4.224](https://github.com/danielmiessler/fabric/releases/tag/v1.4.224) (Jul 1, 2025) — **Add code_review**: Add code_review pattern and updates in Pattern_Descriptions
- [v1.4.222](https://github.com/danielmiessler/fabric/releases/tag/v1.4.222) (Jul 1, 2025) — **OpenAI Plugin**: OpenAI Plugin Migrates to New Responses API
- [v1.4.218](https://github.com/danielmiessler/fabric/releases/tag/v1.4.218) (Jun 27, 2025) — **Model Management**: Add Support for OpenAI Search and Research Model Variants
- [v1.4.217](https://github.com/danielmiessler/fabric/releases/tag/v1.4.217) (Jun 26, 2025) — **New YouTube**: New YouTube Transcript Endpoint Added to REST API
- [v1.4.212](https://github.com/danielmiessler/fabric/releases/tag/v1.4.212) (Jun 23, 2025) — **Add Langdock**: Add Langdock AI and enhance generic OpenAI compatible support
- [v1.4.211](https://github.com/danielmiessler/fabric/releases/tag/v1.4.211) (Jun 19, 2025) — **REST API**: REST API and Web UI Now Support Dynamic Pattern Variables
- [v1.4.210](https://github.com/danielmiessler/fabric/releases/tag/v1.4.210) (Jun 18, 2025) — **Add Citations**: Add Citation Support to Perplexity Response
- [v1.4.208](https://github.com/danielmiessler/fabric/releases/tag/v1.4.208) (Jun 17, 2025) — **Add Perplexity**: Add Perplexity AI Provider with Token Limits Support
- [v1.4.203](https://github.com/danielmiessler/fabric/releases/tag/v1.4.203) (Jun 14, 2025) — **Add Amazon Bedrock**: Add support for Amazon Bedrock

These features represent our commitment to making Fabric the most powerful and flexible AI augmentation framework available!

</details>

## Intro videos

Keep in mind that many of these were recorded when Fabric was Python-based, so remember to use the current [install instructions](#installation) below.

- [Network Chuck](https://www.youtube.com/watch?v=UbDyjIIGaxQ)
- [David Bombal](https://www.youtube.com/watch?v=vF-MQmVxnCs)
- [My Own Intro to the Tool](https://www.youtube.com/watch?v=wPEyyigh10g)
- [More Fabric YouTube Videos](https://www.youtube.com/results?search_query=fabric+ai)

## Navigation

- [`fabric`](#fabric)
  - [What and why](#what-and-why)
  - [Updates](#updates)
    - [Recent Major Features](#recent-major-features)
  - [Intro videos](#intro-videos)
  - [Navigation](#navigation)
  - [Changelog](#changelog)
  - [Philosophy](#philosophy)
    - [Breaking problems into components](#breaking-problems-into-components)
    - [Too many prompts](#too-many-prompts)
  - [Installation](#installation)
    - [One-Line Install (Recommended)](#one-line-install-recommended)
    - [Manual Binary Downloads](#manual-binary-downloads)
    - [Using package managers](#using-package-managers)
      - [macOS (Homebrew)](#macos-homebrew)
      - [Arch Linux (AUR)](#arch-linux-aur)
      - [Windows](#windows)
    - [From Source](#from-source)
    - [Docker](#docker)
    - [Environment Variables](#environment-variables)
    - [Setup](#setup)
    - [Per-Pattern Model Mapping](#per-pattern-model-mapping)
    - [Add aliases for all patterns](#add-aliases-for-all-patterns)
      - [Save your files in markdown using aliases](#save-your-files-in-markdown-using-aliases)
    - [Migration](#migration)
    - [Upgrading](#upgrading)
    - [Shell Completions](#shell-completions)
      - [Quick install (no clone required)](#quick-install-no-clone-required)
      - [Zsh Completion](#zsh-completion)
      - [Bash Completion](#bash-completion)
      - [Fish Completion](#fish-completion)
  - [Usage](#usage)
    - [Debug Levels](#debug-levels)
  - [Our approach to prompting](#our-approach-to-prompting)
  - [Examples](#examples)
  - [Just use the Patterns](#just-use-the-patterns)
    - [Prompt Strategies](#prompt-strategies)
  - [Custom Patterns](#custom-patterns)
    - [Setting Up Custom Patterns](#setting-up-custom-patterns)
    - [Using Custom Patterns](#using-custom-patterns)
    - [How It Works](#how-it-works)
  - [Helper Apps](#helper-apps)
    - [`to_pdf`](#to_pdf)
    - [`to_pdf` Installation](#to_pdf-installation)
    - [`code_helper`](#code_helper)
  - [pbpaste](#pbpaste)
  - [Web Interface (Fabric Web App)](#web-interface-fabric-web-app)
  - [Meta](#meta)
    - [Primary contributors](#primary-contributors)
    - [Contributors](#contributors)

<br />

## Changelog

Fabric is evolving rapidly.

Stay current with the latest features by reviewing the [CHANGELOG](./CHANGELOG.md) for all recent changes.

## Philosophy

> AI isn't a thing; it's a _magnifier_ of a thing. And that thing is **human creativity**.

We believe the purpose of technology is to help humans flourish, so when we talk about AI we start with the **human** problems we want to solve.

### Breaking problems into components

Our approach is to break problems into individual pieces (see below) and then apply AI to them one at a time. See below for some examples.

<img width="2078" alt="augmented_challenges" src="https://github.com/danielmiessler/fabric/assets/50654/31997394-85a9-40c2-879b-b347e4701f06">

### Too many prompts

Prompts are good for this, but the biggest challenge I faced in 2023——which still exists today—is **the sheer number of AI prompts out there**. We all have prompts that are useful, but it's hard to discover new ones, know if they are good or not, _and manage different versions of the ones we like_.

One of `fabric`'s primary features is helping people collect and integrate prompts, which we call _Patterns_, into various parts of their lives.

Fabric has Patterns for all sorts of life and work activities, including:

- Extracting the most interesting parts of YouTube videos and podcasts
- Writing an essay in your own voice with just an idea as an input
- Summarizing opaque academic papers
- Creating perfectly matched AI art prompts for a piece of writing
- Rating the quality of content to see if you want to read/watch the whole thing
- Getting summaries of long, boring content
- Explaining code to you
- Turning bad documentation into usable documentation
- Creating social media posts from any content input
- And a million more…

## Installation

### One-Line Install (Recommended)

**Unix/Linux/macOS:**

```bash
curl -fsSL https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.sh | bash
```

**Windows PowerShell:**

```powershell
iwr -useb https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.ps1 | iex
```

> See [scripts/installer/README.md](./scripts/installer/README.md) for custom installation options and troubleshooting.

### Manual Binary Downloads

The latest release binary archives and their expected SHA256 hashes can be found at <https://github.com/danielmiessler/fabric/releases/latest>

### Using package managers

**NOTE:** using Homebrew or the Arch Linux package managers makes `fabric` available as `fabric-ai`, so add
the following alias to your shell startup files to account for this:

```bash
alias fabric='fabric-ai'
```

#### macOS (Homebrew)

`brew install fabric-ai`

#### Arch Linux (AUR)

`yay -S fabric-ai`

#### Windows

Use the official Microsoft supported `Winget` tool:

`winget install danielmiessler.Fabric`

### From Source

To install Fabric, [make sure Go is installed](https://go.dev/doc/install), and then run the following command.

```bash
# Install Fabric directly from the repo
go install github.com/danielmiessler/fabric/cmd/fabric@latest
```

### Docker

Run Fabric using pre-built Docker images:

```bash
# Use latest image from Docker Hub
docker run --rm -it kayvan/fabric:latest --version

# Use specific version from GHCR
docker run --rm -it ghcr.io/ksylvan/fabric:v1.4.305 --version

# Run setup (first time)
mkdir -p $HOME/.fabric-config
docker run --rm -it -v $HOME/.fabric-config:/root/.config/fabric kayvan/fabric:latest --setup

# Use Fabric with your patterns
docker run --rm -it -v $HOME/.fabric-config:/root/.config/fabric kayvan/fabric:latest -p summarize

# Run the REST API server
docker run --rm -it -p 8080:8080 -v $HOME/.fabric-config:/root/.config/fabric kayvan/fabric:latest --serve
```

**Images available at:**

- Docker Hub: [kayvan/fabric](https://hub.docker.com/repository/docker/kayvan/fabric/general)
- GHCR: [ksylvan/fabric](https://github.com/ksylvan/fabric/pkgs/container/fabric)

See [scripts/docker/README.md](./scripts/docker/README.md) for building custom images and advanced configuration.

### Environment Variables

You may need to set some environment variables in your `~/.bashrc` on linux or `~/.zshrc` file on mac to be able to run the `fabric` command. Here is an example of what you can add:

For Intel based macs or linux

```bash
# Golang environment variables
export GOROOT=/usr/local/go
export GOPATH=$HOME/go

# Update PATH to include GOPATH and GOROOT binaries
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
```

for Apple Silicon based macs

```bash
# Golang environment variables
export GOROOT=$(brew --prefix go)/libexec
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
```

### Setup

Now run the following command

```bash
# Run the setup to set up your directories and keys
fabric --setup
```

If everything works you are good to go.

### Per-Pattern Model Mapping

 You can configure specific models for individual patterns using environment variables
 like `FABRIC_MODEL_PATTERN_NAME=vendor|model`

 This makes it easy to maintain these per-pattern model mappings in your shell startup files.

### Add aliases for all patterns

In order to add aliases for all your patterns and use them directly as commands, for example, `summarize` instead of `fabric --pattern summarize`
You can add the following to your `.zshrc` or `.bashrc` file. You
can also optionally set the `FABRIC_ALIAS_PREFIX` environment variable
before, if you'd prefer all the fabric aliases to start with the same prefix.

```bash
# Loop through all files in the ~/.config/fabric/patterns directory
for pattern_file in $HOME/.config/fabric/patterns/*; do
    # Get the base name of the file (i.e., remove the directory path)
    pattern_name="$(basename "$pattern_file")"
    alias_name="${FABRIC_ALIAS_PREFIX:-}${pattern_name}"

    # Create an alias in the form: alias pattern_name="fabric --pattern pattern_name"
    alias_command="alias $alias_name='fabric --pattern $pattern_name'"

    # Evaluate the alias command to add it to the current shell
    eval "$alias_command"
done

yt() {
    if [ "$#" -eq 0 ] || [ "$#" -gt 2 ]; then
        echo "Usage: yt [-t | --timestamps] youtube-link"
        echo "Use the '-t' flag to get the transcript with timestamps."
        return 1
    fi

    transcript_flag="--transcript"
    if [ "$1" = "-t" ] || [ "$1" = "--timestamps" ]; then
        transcript_flag="--transcript-with-timestamps"
        shift
    fi
    local video_link="$1"
    fabric -y "$video_link" $transcript_flag
}
```

You can add the below code for the equivalent aliases inside PowerShell by running `notepad $PROFILE` inside a PowerShell window:

```powershell
# Path to the patterns directory
$patternsPath = Join-Path $HOME ".config/fabric/patterns"
foreach ($patternDir in Get-ChildItem -Path $patternsPath -Directory) {
    # Prepend FABRIC_ALIAS_PREFIX if set; otherwise use empty string
    $prefix = $env:FABRIC_ALIAS_PREFIX ?? ''
    $patternName = "$($patternDir.Name)"
    $aliasName = "$prefix$patternName"
    # Dynamically define a function for each pattern
    $functionDefinition = @"
function $aliasName {
    [CmdletBinding()]
    param(
        [Parameter(ValueFromPipeline = `$true)]
        [string] `$InputObject,

        [Parameter(ValueFromRemainingArguments = `$true)]
        [String[]] `$patternArgs
    )

    begin {
        # Initialize an array to collect pipeline input
        `$collector = @()
    }

    process {
        # Collect pipeline input objects
        if (`$InputObject) {
            `$collector += `$InputObject
        }
    }

    end {
        # Join all pipeline input into a single string, separated by newlines
        `$pipelineContent = `$collector -join "`n"

        # If there's pipeline input, include it in the call to fabric
        if (`$pipelineContent) {
            `$pipelineContent | fabric --pattern $patternName `$patternArgs
        } else {
            # No pipeline input; just call fabric with the additional args
            fabric --pattern $patternName `$patternArgs
        }
    }
}
"@
    # Add the function to the current session
    Invoke-Expression $functionDefinition
}

# Define the 'yt' function as well
function yt {
    [CmdletBinding()]
    param(
        [Parameter()]
        [Alias("timestamps")]
        [switch]$t,

        [Parameter(Position = 0, ValueFromPipeline = $true)]
        [string]$videoLink
    )

    begin {
        $transcriptFlag = "--transcript"
        if ($t) {
            $transcriptFlag = "--transcript-with-timestamps"
        }
    }

    process {
        if (-not $videoLink) {
            Write-Error "Usage: yt [-t | --timestamps] youtube-link"
            return
        }
    }

    end {
        if ($videoLink) {
            # Execute and allow output to flow through the pipeline
            fabric -y $videoLink $transcriptFlag
        }
    }
}
```

This also creates a `yt` alias that allows you to use `yt https://www.youtube.com/watch?v=4b0iet22VIk` to get transcripts, comments, and metadata.

#### Save your files in markdown using aliases

If in addition to the above aliases you would like to have the option to save the output to your favorite markdown note vault like Obsidian then instead of the above add the following to your `.zshrc` or `.bashrc` file:

```bash
# Define the base directory for Obsidian notes
obsidian_base="/path/to/obsidian"

# Loop through all files in the ~/.config/fabric/patterns directory
for pattern_file in ~/.config/fabric/patterns/*; do
    # Get the base name of the file (i.e., remove the directory path)
    pattern_name=$(basename "$pattern_file")

    # Remove any existing alias with the same name
    unalias "$pattern_name" 2>/dev/null

    # Define a function dynamically for each pattern
    eval "
    $pattern_name() {
        local title=\$1
        local date_stamp=\$(date +'%Y-%m-%d')
        local output_path=\"\$obsidian_base/\${date_stamp}-\${title}.md\"

        # Check if a title was provided
        if [ -n \"\$title\" ]; then
            # If a title is provided, use the output path
            fabric --pattern \"$pattern_name\" -o \"\$output_path\"
        else
            # If no title is provided, use --stream
            fabric --pattern \"$pattern_name\" --stream
        fi
    }
    "
done
```

This will allow you to use the patterns as aliases like in the above for example `summarize` instead of `fabric --pattern summarize --stream`, however if you pass in an extra argument like this `summarize "my_article_title"` your output will be saved in the destination that you set in `obsidian_base="/path/to/obsidian"` in the following format `YYYY-MM-DD-my_article_title.md` where the date gets autogenerated for you.
You can tweak the date format by tweaking the `date_stamp` format.

### Migration

If you have the Legacy (Python) version installed and want to migrate to the Go version, here's how you do it. It's basically two steps: 1) uninstall the Python version, and 2) install the Go version.

```bash
# Uninstall Legacy Fabric
pipx uninstall fabric

# Clear any old Fabric aliases
(check your .bashrc, .zshrc, etc.)
# Install the Go version
go install github.com/danielmiessler/fabric/cmd/fabric@latest
# Run setup for the new version. Important because things have changed
fabric --setup
```

Then [set your environmental variables](#environment-variables) as shown above.

### Upgrading

The great thing about Go is that it's super easy to upgrade. Just run the same command you used to install it in the first place and you'll always get the latest version.

```bash
go install github.com/danielmiessler/fabric/cmd/fabric@latest
```

### Shell Completions

Fabric provides shell completion scripts for Zsh, Bash, and Fish
shells, making it easier to use the CLI by providing tab completion
for commands and options.

#### Quick install (no clone required)

You can install completions directly via a one-liner:

```bash
curl -fsSL https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions/setup-completions.sh | sh
```

Optional variants:

```bash
# Dry-run (see actions without changing your system)
curl -fsSL https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions/setup-completions.sh | sh -s -- --dry-run

# Override the download source (advanced)
FABRIC_COMPLETIONS_BASE_URL="https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions" \
    sh -c "$(curl -fsSL https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions/setup-completions.sh)"
```

#### Zsh Completion

To enable Zsh completion:

```bash
# Copy the completion file to a directory in your $fpath
mkdir -p ~/.zsh/completions
cp completions/_fabric ~/.zsh/completions/

# Add the directory to fpath in your .zshrc before compinit
echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc
echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
```

#### Bash Completion

To enable Bash completion:

```bash
# Source the completion script in your .bashrc
echo 'source /path/to/fabric/completions/fabric.bash' >> ~/.bashrc

# Or copy to the system-wide bash completion directory
sudo cp completions/fabric.bash /etc/bash_completion.d/
```

#### Fish Completion

To enable Fish completion:

```bash
# Copy the completion file to the fish completions directory
mkdir -p ~/.config/fish/completions
cp completions/fabric.fish ~/.config/fish/completions/
```

## Usage

Once you have it all set up, here's how to use it.

```bash
fabric -h
```

```plaintext
Usage:
  fabric [OPTIONS]

Application Options:
  -p, --pattern=                    Choose a pattern from the available patterns
  -v, --variable=                   Values for pattern variables, e.g. -v=#role:expert -v=#points:30
  -C, --context=                    Choose a context from the available contexts
      --session=                    Choose a session from the available sessions
  -a, --attachment=                 Attachment path or URL (e.g. for OpenAI image recognition messages)
  -S, --setup                       Run setup for all reconfigurable parts of fabric
  -t, --temperature=                Set temperature (default: 0.7)
  -T, --topp=                       Set top P (default: 0.9)
  -s, --stream                      Stream
  -P, --presencepenalty=            Set presence penalty (default: 0.0)
  -r, --raw                         Use the defaults of the model without sending chat options (like
                                    temperature etc.) and use the user role instead of the system role for
                                    patterns.
  -F, --frequencypenalty=           Set frequency penalty (default: 0.0)
  -l, --listpatterns                List all patterns
  -L, --listmodels                  List all available models
  -x, --listcontexts                List all contexts
  -X, --listsessions                List all sessions
  -U, --updatepatterns              Update patterns
  -c, --copy                        Copy to clipboard
  -m, --model=                      Choose model
  -V, --vendor=                     Specify vendor for chosen model (e.g., -V "LM Studio" -m openai/gpt-oss-20b)
      --modelContextLength=         Model context length (only affects ollama)
  -o, --output=                     Output to file
      --output-session              Output the entire session (also a temporary one) to the output file
  -n, --latest=                     Number of latest patterns to list (default: 0)
  -d, --changeDefaultModel          Change default model
  -y, --youtube=                    YouTube video or play list "URL" to grab transcript, comments from it
                                    and send to chat or print it put to the console and store it in the
                                    output file
      --playlist                    Prefer playlist over video if both ids are present in the URL
      --transcript                  Grab transcript from YouTube video and send to chat (it is used per
                                    default).
      --transcript-with-timestamps  Grab transcript from YouTube video with timestamps and send to chat
      --comments                    Grab comments from YouTube video and send to chat
      --metadata                    Output video metadata
  -g, --language=                   Specify the Language Code for the chat, e.g. -g=en -g=zh
  -u, --scrape_url=                 Scrape website URL to markdown using Jina AI
  -q, --scrape_question=            Search question using Jina AI
  -e, --seed=                       Seed to be used for LMM generation
  -w, --wipecontext=                Wipe context
  -W, --wipesession=                Wipe session
      --printcontext=               Print context
      --printsession=               Print session
      --readability                 Convert HTML input into a clean, readable view
      --input-has-vars              Apply variables to user input
      --no-variable-replacement     Disable pattern variable replacement
      --dry-run                     Show what would be sent to the model without actually sending it
      --serve                       Serve the Fabric Rest API
      --serveOllama                 Serve the Fabric Rest API with ollama endpoints
      --address=                    The address to bind the REST API (default: :8080)
      --api-key=                    API key used to secure server routes
      --config=                     Path to YAML config file
      --version                     Print current version
      --listextensions              List all registered extensions
      --addextension=               Register a new extension from config file path
      --rmextension=                Remove a registered extension by name
      --strategy=                   Choose a strategy from the available strategies
      --liststrategies              List all strategies
      --listvendors                 List all vendors
      --shell-complete-list         Output raw list without headers/formatting (for shell completion)
      --search                      Enable web search tool for supported models (Anthropic, OpenAI, Gemini)
      --search-location=            Set location for web search results (e.g., 'America/Los_Angeles')
      --image-file=                 Save generated image to specified file path (e.g., 'output.png')
      --image-size=                 Image dimensions: 1024x1024, 1536x1024, 1024x1536, auto (default: auto)
      --image-quality=              Image quality: low, medium, high, auto (default: auto)
      --image-compression=          Compression level 0-100 for JPEG/WebP formats (default: not set)
      --image-background=           Background type: opaque, transparent (default: opaque, only for
                                    PNG/WebP)
      --suppress-think              Suppress text enclosed in thinking tags
      --think-start-tag=            Start tag for thinking sections (default: <think>)
      --think-end-tag=              End tag for thinking sections (default: </think>)
      --disable-responses-api       Disable OpenAI Responses API (default: false)
      --voice=                      TTS voice name for supported models (e.g., Kore, Charon, Puck)
                                    (default: Kore)
      --list-gemini-voices          List all available Gemini TTS voices
      --notification                Send desktop notification when command completes
      --notification-command=       Custom command to run for notifications (overrides built-in
                                    notifications)
      --yt-dlp-args=                Additional arguments to pass to yt-dlp (e.g. '--cookies-from-browser brave')
      --thinking=                   Set reasoning/thinking level (e.g., off, low, medium, high, or
                                    numeric tokens for Anthropic or Google Gemini)
      --debug=                     Set debug level (0: off, 1: basic, 2: detailed, 3: trace)
Help Options:
  -h, --help                        Show this help message
```

### Debug Levels

Use the `--debug` flag to control runtime logging:

- `0`: off (default)
- `1`: basic debug info
- `2`: detailed debugging
- `3`: trace level

## Our approach to prompting

Fabric _Patterns_ are different than most prompts you'll see.

- **First, we use `Markdown` to help ensure maximum readability and editability**. This not only helps the creator make a good one, but also anyone who wants to deeply understand what it does. _Importantly, this also includes the AI you're sending it to!_

Here's an example of a Fabric Pattern.

```bash
https://github.com/danielmiessler/Fabric/blob/main/data/patterns/extract_wisdom/system.md
```

<img width="1461" alt="pattern-example" src="https://github.com/danielmiessler/fabric/assets/50654/b910c551-9263-405f-9735-71ca69bbab6d">

- **Next, we are extremely clear in our instructions**, and we use the Markdown structure to emphasize what we want the AI to do, and in what order.

- **And finally, we tend to use the System section of the prompt almost exclusively**. In over a year of being heads-down with this stuff, we've just seen more efficacy from doing that. If that changes, or we're shown data that says otherwise, we will adjust.

## Examples

> The following examples use the macOS `pbpaste` to paste from the clipboard. See the [pbpaste](#pbpaste) section below for Windows and Linux alternatives.

Now let's look at some things you can do with Fabric.

1. Run the `summarize` Pattern based on input from `stdin`. In this case, the body of an article.

    ```bash
    pbpaste | fabric --pattern summarize
    ```

2. Run the `analyze_claims` Pattern with the `--stream` option to get immediate and streaming results.

    ```bash
    pbpaste | fabric --stream --pattern analyze_claims
    ```

3. Run the `extract_wisdom` Pattern with the `--stream` option to get immediate and streaming results from any      Youtube video (much like in the original introduction video).

    ```bash
    fabric -y "https://youtube.com/watch?v=uXs-zPc63kM" --stream --pattern extract_wisdom
    ```

4. Create patterns- you must create a .md file with the pattern and save it to `~/.config/fabric/patterns/[yourpatternname]`.

5. Run a `analyze_claims` pattern on a website. Fabric uses Jina AI to scrape the URL into markdown format before sending it to the model.

    ```bash
    fabric -u https://github.com/danielmiessler/fabric/ -p analyze_claims
    ```

## Just use the Patterns

<img width="1173" alt="fabric-patterns-screenshot" src="https://github.com/danielmiessler/fabric/assets/50654/9186a044-652b-4673-89f7-71cf066f32d8">

<br />
<br />

If you're not looking to do anything fancy, and you just want a lot of great prompts, you can navigate to the [`/patterns`](https://github.com/danielmiessler/fabric/tree/main/data/patterns) directory and start exploring!

We hope that if you used nothing else from Fabric, the Patterns by themselves will make the project useful.

You can use any of the Patterns you see there in any AI application that you have, whether that's ChatGPT or some other app or website. Our plan and prediction is that people will soon be sharing many more than those we've published, and they will be way better than ours.

The wisdom of crowds for the win.

### Prompt Strategies

Fabric also implements prompt strategies like "Chain of Thought" or "Chain of Draft" which can
be used in addition to the basic patterns.

See the [Thinking Faster by Writing Less](https://arxiv.org/pdf/2502.18600) paper and
the [Thought Generation section of Learn Prompting](https://learnprompting.org/docs/advanced/thought_generation/introduction) for examples of prompt strategies.

Each strategy is available as a small `json` file in the [`/strategies`](https://github.com/danielmiessler/fabric/tree/main/data/strategies) directory.

The prompt modification of the strategy is applied to the system prompt and passed on to the
LLM in the chat session.

Use `fabric -S` and select the option to install the strategies in your `~/.config/fabric` directory.

## Custom Patterns

You may want to use Fabric to create your own custom Patterns—but not share them with others. No problem!

Fabric now supports a dedicated custom patterns directory that keeps your personal patterns separate from the built-in ones. This means your custom patterns won't be overwritten when you update Fabric's built-in patterns.

### Setting Up Custom Patterns

1. Run the Fabric setup:

   ```bash
   fabric --setup
   ```

2. Select the "Custom Patterns" option from the Tools menu and enter your desired directory path (e.g., `~/my-custom-patterns`)

3. Fabric will automatically create the directory if it does not exist.

### Using Custom Patterns

1. Create your custom pattern directory structure:

   ```bash
   mkdir -p ~/my-custom-patterns/my-analyzer
   ```

2. Create your pattern file

   ```bash
   echo "You are an expert analyzer of ..." > ~/my-custom-patterns/my-analyzer/system.md
   ```

3. **Use your custom pattern:**

   ```bash
   fabric --pattern my-analyzer "analyze this text"
   ```

### How It Works

- **Priority System**: Custom patterns take precedence over built-in patterns with the same name
- **Seamless Integration**: Custom patterns appear in `fabric --listpatterns` alongside built-in ones
- **Update Safe**: Your custom patterns are never affected by `fabric --updatepatterns`
- **Private by Default**: Custom patterns remain private unless you explicitly share them

Your custom patterns are completely private and won't be affected by Fabric updates!

## Helper Apps

Fabric also makes use of some core helper apps (tools) to make it easier to integrate with your various workflows. Here are some examples:

### `to_pdf`

`to_pdf` is a helper command that converts LaTeX files to PDF format. You can use it like this:

```bash
to_pdf input.tex
```

This will create a PDF file from the input LaTeX file in the same directory.

You can also use it with stdin which works perfectly with the `write_latex` pattern:

```bash
echo "ai security primer" | fabric --pattern write_latex | to_pdf
```

This will create a PDF file named `output.pdf` in the current directory.

### `to_pdf` Installation

To install `to_pdf`, install it the same way as you install Fabric, just with a different repo name.

```bash
go install github.com/danielmiessler/fabric/cmd/to_pdf@latest
```

Make sure you have a LaTeX distribution (like TeX Live or MiKTeX) installed on your system, as `to_pdf` requires `pdflatex` to be available in your system's PATH.

### `code_helper`

`code_helper` is used in conjunction with the `create_coding_feature` pattern.
It generates a `json` representation of a directory of code that can be fed into an AI model
with instructions to create a new feature or edit the code in a specified way.

See [the Create Coding Feature Pattern README](./data/patterns/create_coding_feature/README.md) for details.

Install it first using:

```bash
go install github.com/danielmiessler/fabric/cmd/code_helper@latest
```

## pbpaste

The [examples](#examples) use the macOS program `pbpaste` to paste content from the clipboard to pipe into `fabric` as the input. `pbpaste` is not available on Windows or Linux, but there are alternatives.

On Windows, you can use the PowerShell command `Get-Clipboard` from a PowerShell command prompt. If you like, you can also alias it to `pbpaste`. If you are using classic PowerShell, edit the file `~\Documents\WindowsPowerShell\.profile.ps1`, or if you are using PowerShell Core, edit `~\Documents\PowerShell\.profile.ps1` and add the alias,

```powershell
Set-Alias pbpaste Get-Clipboard
```

On Linux, you can use `xclip -selection clipboard -o` to paste from the clipboard. You will likely need to install `xclip` with your package manager. For Debian based systems including Ubuntu,

```sh
sudo apt update
sudo apt install xclip -y
```

You can also create an alias by editing `~/.bashrc` or `~/.zshrc` and adding the alias,

```sh
alias pbpaste='xclip -selection clipboard -o'
```

## Web Interface (Fabric Web App)

Fabric now includes a built-in web interface that provides a GUI alternative to the command-line interface. Refer to [Web App README](/web/README.md) for installation instructions and an overview of features.

## Meta

> [!NOTE]
> Special thanks to the following people for their inspiration and contributions!

- _Jonathan Dunn_ for being the absolute MVP dev on the project, including spearheading the new Go version, as well as the GUI! All this while also being a full-time medical doctor!
- _Caleb Sima_ for pushing me over the edge of whether to make this a public project or not.
- _Eugen Eisler_ and _Frederick Ros_ for their invaluable contributions to the Go version
- _David Peters_ for his work on the web interface.
- _Joel Parish_ for super useful input on the project's Github directory structure..
- _Joseph Thacker_ for the idea of a `-c` context flag that adds pre-created context in the `./config/fabric/` directory to all Pattern queries.
- _Jason Haddix_ for the idea of a stitch (chained Pattern) to filter content using a local model before sending on to a cloud model, i.e., cleaning customer data using `llama2` before sending on to `gpt-4` for analysis.
- _Andre Guerra_ for assisting with numerous components to make things simpler and more maintainable.

### Primary contributors

<a href="https://github.com/danielmiessler"><img src="https://avatars.githubusercontent.com/u/50654?v=4" title="Daniel Miessler" width="50" height="50" alt="Daniel Miessler"></a>
<a href="https://github.com/xssdoctor"><img src="https://avatars.githubusercontent.com/u/9218431?v=4" title="Jonathan Dunn" width="50" height="50" alt="Jonathan Dunn"></a>
<a href="https://github.com/sbehrens"><img src="https://avatars.githubusercontent.com/u/688589?v=4" title="Scott Behrens" width="50" height="50" alt="Scott Behrens"></a>
<a href="https://github.com/agu3rra"><img src="https://avatars.githubusercontent.com/u/10410523?v=4" title="Andre Guerra" width="50" height="50" alt="Andre Guerra"></a>

### Contributors

<a href="https://github.com/danielmiessler/fabric/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=danielmiessler/fabric" alt="contrib.rocks" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

`fabric` was created by <a href="https://danielmiessler.com/subscribe" target="_blank">Daniel Miessler</a> in January of 2024.
<br /><br />
<a href="https://twitter.com/intent/user?screen_name=danielmiessler">![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/danielmiessler)</a>
