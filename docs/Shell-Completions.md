# Shell Completions for Fabric

Fabric comes with shell completion support for Zsh, Bash, and Fish shells. These completions provide intelligent tab-completion for commands, flags, patterns, models, contexts, and more.

## Quick Setup (Automated)

You can install completions without cloning the repo:

```bash
# No-clone install (Zsh/Bash/Fish supported)
curl -fsSL https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions/setup-completions.sh | sh

# Optional: dry-run first
curl -fsSL https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions/setup-completions.sh | sh -s -- --dry-run

# Optional: override the download source
FABRIC_COMPLETIONS_BASE_URL="https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions" \
   sh -c "$(curl -fsSL https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions/setup-completions.sh)"
```

Or, if you have the repository locally:

```bash
# Run the automated setup script from a cloned repo
./completions/setup-completions.sh

# Or see what it would do first
./completions/setup-completions.sh --dry-run
```

The script will:

- Detect whether you have `fabric` or `fabric-ai` installed
- Detect your current shell (zsh, bash, or fish)
- Use your existing `$fpath` directories (for zsh) or standard completion directories
- Install the completion file with the correct name
- Provide instructions for enabling the completions

If the completion files aren't present locally (e.g., when running via `curl`), the script will automatically download them from GitHub.

For manual installation or troubleshooting, see the detailed instructions below.

## Manual Installation

### Zsh

1. Copy the completion file to a directory in your `$fpath`:

   ```bash
   sudo cp completions/_fabric /usr/local/share/zsh/site-functions/
   ```

2. **Important**: If you installed fabric as `fabric-ai`, create a symlink so completions work:

   ```bash
   sudo ln -s /usr/local/share/zsh/site-functions/_fabric /usr/local/share/zsh/site-functions/_fabric-ai
   ```

3. Restart your shell or reload completions:

   ```bash
   autoload -U compinit && compinit
   ```

### Bash

1. Copy the completion file to a standard completion directory:

   ```bash
   # System-wide installation
   sudo cp completions/fabric.bash /etc/bash_completion.d/

   # Or user-specific installation
   mkdir -p ~/.local/share/bash-completion/completions/
   cp completions/fabric.bash ~/.local/share/bash-completion/completions/fabric
   ```

2. **Important**: If you installed fabric as `fabric-ai`, create a symlink:

   ```bash
   # For system-wide installation
   sudo ln -s /etc/bash_completion.d/fabric.bash /etc/bash_completion.d/fabric-ai.bash

   # Or for user-specific installation
   ln -s ~/.local/share/bash-completion/completions/fabric ~/.local/share/bash-completion/completions/fabric-ai
   ```

3. Restart your shell or source the completion:

   ```bash
   source ~/.bashrc
   ```

### Fish

1. Copy the completion file to Fish's completion directory:

   ```bash
   mkdir -p ~/.config/fish/completions
   cp completions/fabric.fish ~/.config/fish/completions/
   ```

2. **Important**: If you installed fabric as `fabric-ai`, create a symlink:

   ```bash
   ln -s ~/.config/fish/completions/fabric.fish ~/.config/fish/completions/fabric-ai.fish
   ```

3. Fish will automatically load the completions (no restart needed).

## Features

The completions provide intelligent suggestions for:

- **Patterns**: Tab-complete available patterns with `-p` or `--pattern`
- **Models**: Tab-complete available models with `-m` or `--model`
- **Contexts**: Tab-complete contexts for context-related flags
- **Sessions**: Tab-complete sessions for session-related flags
- **Strategies**: Tab-complete available strategies
- **Extensions**: Tab-complete registered extensions
- **Gemini Voices**: Tab-complete TTS voices for `--voice`
- **File paths**: Smart file completion for attachment, output, and config options
- **Flag completion**: All available command-line flags and options

## Alternative Installation Method

You can also source the completion files directly in your shell's configuration file:

- **Zsh**: Add to `~/.zshrc`: `source /path/to/fabric/completions/_fabric`
- **Bash**: Add to `~/.bashrc`: `source /path/to/fabric/completions/fabric.bash`
- **Fish**: The file-based installation method above is preferred for Fish

## Troubleshooting

- If completions don't work, ensure the completion files have proper permissions
- For Zsh, verify that the completion directory is in your `$fpath`
- If you renamed the fabric binary, make sure to create the appropriate symlinks as described above
- Restart your shell after installation to ensure completions are loaded

The completion system dynamically queries the fabric command for current patterns, models, and other resources, so your completions will always be up-to-date with your fabric installation.
