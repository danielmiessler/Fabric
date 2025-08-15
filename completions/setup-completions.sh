#!/bin/sh

# Fabric Shell Completions Setup Script
# This script automatically installs shell completions for the fabric CLI
# based on your current shell and the installed fabric command name.

set -e

# Global variables
DRY_RUN=false
# Base URL to fetch completion files when not available locally
# Can be overridden via environment variable FABRIC_COMPLETIONS_BASE_URL
FABRIC_COMPLETIONS_BASE_URL="${FABRIC_COMPLETIONS_BASE_URL:-https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions}"
TEMP_DIR=""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    printf "${BLUE}[INFO]${NC} %s\n" "$1"
}

print_success() {
    printf "${GREEN}[SUCCESS]${NC} %s\n" "$1"
}

print_warning() {
    printf "${YELLOW}[WARNING]${NC} %s\n" "$1"
}

print_error() {
    printf "${RED}[ERROR]${NC} %s\n" "$1"
}

print_dry_run() {
    printf "${CYAN}[DRY-RUN]${NC} %s\n" "$1"
}

# Function to execute commands with dry-run support
execute_command() {
    cmd="$1"

    if [ "$DRY_RUN" = true ]; then
        print_dry_run "Would run: $cmd"
        return 0
    else
        eval "$cmd" 2>/dev/null
    fi
}

# Simple downloader that prefers curl, falls back to wget
to_github_raw_url() {
    in_url="$1"
    case "$in_url" in
        https://github.com/*/*/blob/*)
            # Convert blob URL to raw
            # https://github.com/{owner}/{repo}/blob/{ref}/path -> https://raw.githubusercontent.com/{owner}/{repo}/{ref}/path
            echo "$in_url" | sed -E 's#https://github.com/([^/]+)/([^/]+)/blob/([^/]+)/#https://raw.githubusercontent.com/\1/\2/\3/#'
            ;;
        https://github.com/*/*/tree/*)
            # Convert tree URL base + file path to raw
            # https://github.com/{owner}/{repo}/tree/{ref}/path -> https://raw.githubusercontent.com/{owner}/{repo}/{ref}/path
            echo "$in_url" | sed -E 's#https://github.com/([^/]+)/([^/]+)/tree/([^/]+)/#https://raw.githubusercontent.com/\1/\2/\3/#'
            ;;
        *)
            echo "$in_url"
            ;;
    esac
}

# Simple downloader that prefers curl, falls back to wget
download_file() {
    url="$1"
    dest="$2"

    if [ "$DRY_RUN" = true ]; then
        print_dry_run "Would download: $url -> $dest"
        return 0
    fi

    eff_url="$(to_github_raw_url "$url")"

    if command -v curl >/dev/null 2>&1; then
        curl -fsSL "$eff_url" -o "$dest"
        return $?
    elif command -v wget >/dev/null 2>&1; then
        wget -q "$eff_url" -O "$dest"
        return $?
    else
        print_error "Neither 'curl' nor 'wget' is available to download: $url"
        return 1
    fi
}

# Attempt to obtain completion files. If local copies are missing,
# download them into a temporary directory and return that directory path.
obtain_completion_files() {
    obf_script_dir="$1"
    obf_need_download=false

    if [ ! -f "$obf_script_dir/_fabric" ] || [ ! -f "$obf_script_dir/fabric.bash" ] || [ ! -f "$obf_script_dir/fabric.fish" ]; then
        obf_need_download=true
    fi

    if [ "$obf_need_download" = false ]; then
        echo "$obf_script_dir"
        return 0
    fi

    # Note: write only to stderr in this function except for the final echo which returns the path
    printf "%s\n" "[INFO] Local completion files not found; will download from GitHub." 1>&2
    printf "%s\n" "[INFO] Source: $FABRIC_COMPLETIONS_BASE_URL" 1>&2

    if [ "$DRY_RUN" = true ]; then
    printf "%s\n" "[DRY-RUN] Would create temporary directory for downloads" 1>&2
    echo "$obf_script_dir" # Keep using original for dry-run copies
        return 0
    fi

    TEMP_DIR="$(mktemp -d 2>/dev/null || mktemp -d -t fabric-completions)"
    if [ ! -d "$TEMP_DIR" ]; then
    print_error "Failed to create temporary directory for downloads."
        return 1
    fi

    if ! download_file "$FABRIC_COMPLETIONS_BASE_URL/_fabric" "$TEMP_DIR/_fabric"; then
        print_error "Failed to download _fabric"
        return 1
    fi
    if [ ! -s "$TEMP_DIR/_fabric" ] || head -n1 "$TEMP_DIR/_fabric" | grep -qi "^<!DOCTYPE\|^<html"; then
        print_error "Downloaded _fabric appears invalid (empty or HTML). Check FABRIC_COMPLETIONS_BASE_URL."
        return 1
    fi
    if ! download_file "$FABRIC_COMPLETIONS_BASE_URL/fabric.bash" "$TEMP_DIR/fabric.bash"; then
        print_error "Failed to download fabric.bash"
        return 1
    fi
    if [ ! -s "$TEMP_DIR/fabric.bash" ] || head -n1 "$TEMP_DIR/fabric.bash" | grep -qi "^<!DOCTYPE\|^<html"; then
        print_error "Downloaded fabric.bash appears invalid (empty or HTML). Check FABRIC_COMPLETIONS_BASE_URL."
        return 1
    fi
    if ! download_file "$FABRIC_COMPLETIONS_BASE_URL/fabric.fish" "$TEMP_DIR/fabric.fish"; then
        print_error "Failed to download fabric.fish"
        return 1
    fi
    if [ ! -s "$TEMP_DIR/fabric.fish" ] || head -n1 "$TEMP_DIR/fabric.fish" | grep -qi "^<!DOCTYPE\|^<html"; then
        print_error "Downloaded fabric.fish appears invalid (empty or HTML). Check FABRIC_COMPLETIONS_BASE_URL."
        return 1
    fi

    echo "$TEMP_DIR"
}

# Ensure directory exists, try sudo on permission failure
ensure_dir() {
    dir="$1"
    # Expand ~ if present
    case "$dir" in
        ~/*)
            dir="$HOME${dir#~}"
            ;;
    esac

    if [ -d "$dir" ]; then
        return 0
    fi

    if [ "$DRY_RUN" = true ]; then
        print_dry_run "Would run: mkdir -p \"$dir\""
        print_dry_run "If permission denied, would run: sudo mkdir -p \"$dir\""
        return 0
    fi

    if mkdir -p "$dir" 2>/dev/null; then
        return 0
    fi
    if command -v sudo >/dev/null 2>&1 && sudo mkdir -p "$dir" 2>/dev/null; then
        return 0
    fi
    print_error "Failed to create directory: $dir"
    return 1
}

# Copy file with sudo fallback on permission failure
install_file() {
    src="$1"
    dest="$2"

    if [ "$DRY_RUN" = true ]; then
        print_dry_run "Would run: cp \"$src\" \"$dest\""
        print_dry_run "If permission denied, would run: sudo cp \"$src\" \"$dest\""
        return 0
    fi

    if cp "$src" "$dest" 2>/dev/null; then
        return 0
    fi
    if command -v sudo >/dev/null 2>&1 && sudo cp "$src" "$dest" 2>/dev/null; then
        return 0
    fi
    print_error "Failed to install file to: $dest"
    return 1
}

# Function to detect fabric command name
detect_fabric_command() {
    if command -v fabric >/dev/null 2>&1; then
        echo "fabric"
    elif command -v fabric-ai >/dev/null 2>&1; then
        echo "fabric-ai"
    else
        print_error "Neither 'fabric' nor 'fabric-ai' command found in PATH"
        exit 1
    fi
}

# Function to detect shell
detect_shell() {
    if [ -n "$SHELL" ]; then
        basename "$SHELL"
    else
        print_warning "SHELL environment variable not set, defaulting to sh"
        echo "sh"
    fi
}

# Function to get script directory
get_script_dir() {
    # Get the directory where this script is located
    script_path="$(readlink -f "$0" 2>/dev/null || realpath "$0" 2>/dev/null || echo "$0")"
    dirname "$script_path"
}

# Function to setup Zsh completions
setup_zsh_completions() {
    fabric_cmd="$1"
    script_dir="$2"
    completion_file="_${fabric_cmd}"

    print_info "Setting up Zsh completions for '$fabric_cmd'..."

    # Try to use existing $fpath first, then fall back to default directories
    zsh_dirs=""

    # Check if user's shell is zsh and try to get fpath from it
    if [ "$(basename "$SHELL")" = "zsh" ] && command -v zsh >/dev/null 2>&1; then
        # Get fpath from zsh by sourcing user's .zshrc first
        fpath_output=$(zsh -c "source \$HOME/.zshrc 2>/dev/null && print -l \$fpath" 2>/dev/null | head -5 | tr '\n' ' ')
        if [ -n "$fpath_output" ] && [ "$fpath_output" != "" ]; then
            print_info "Using directories from zsh \$fpath"
            zsh_dirs="$fpath_output"
        fi
    fi

    # If we couldn't get fpath or it's empty, use default directories
    if [ -z "$zsh_dirs" ] || [ "$zsh_dirs" = "" ]; then
        print_info "Using default zsh completion directories"
        zsh_dirs="/usr/local/share/zsh/site-functions /opt/homebrew/share/zsh/site-functions /usr/share/zsh/site-functions ~/.local/share/zsh/site-functions"
    fi

    installed=false

    for dir in $zsh_dirs; do
        # Create directory (with sudo fallback if needed)
        if ensure_dir "$dir"; then
            if install_file "$script_dir/_fabric" "$dir/$completion_file"; then
                if [ "$DRY_RUN" = true ]; then
                    print_success "Would install Zsh completion to: $dir/$completion_file"
                else
                    print_success "Installed Zsh completion to: $dir/$completion_file"
                fi
                installed=true
                break
            fi
        fi
    done

    if [ "$installed" = false ]; then
        if [ "$DRY_RUN" = true ]; then
            print_warning "Would attempt to install Zsh completions but no writable directory found."
        else
            print_error "Failed to install Zsh completions. Try running with sudo or check permissions."
            return 1
        fi
    fi

    if [ "$DRY_RUN" = true ]; then
        print_info "Would suggest: Restart your shell or run 'autoload -U compinit && compinit' to enable completions."
    else
        print_info "Restart your shell or run 'autoload -U compinit && compinit' to enable completions."
    fi
}

# Function to setup Bash completions
setup_bash_completions() {
    fabric_cmd="$1"
    script_dir="$2"
    completion_file="${fabric_cmd}.bash"

    print_info "Setting up Bash completions for '$fabric_cmd'..."

    # Try different completion directories
    bash_dirs="/etc/bash_completion.d /usr/local/etc/bash_completion.d /opt/homebrew/etc/bash_completion.d ~/.local/share/bash-completion/completions"
    installed=false

    for dir in $bash_dirs; do
        if ensure_dir "$dir"; then
            if install_file "$script_dir/fabric.bash" "$dir/$completion_file"; then
                if [ "$DRY_RUN" = true ]; then
                    print_success "Would install Bash completion to: $dir/$completion_file"
                else
                    print_success "Installed Bash completion to: $dir/$completion_file"
                fi
                installed=true
                break
            fi
        fi
    done

    if [ "$installed" = false ]; then
        if [ "$DRY_RUN" = true ]; then
            print_warning "Would attempt to install Bash completions but no writable directory found."
        else
            print_error "Failed to install Bash completions. Try running with sudo or check permissions."
            return 1
        fi
    fi

    if [ "$DRY_RUN" = true ]; then
        print_info "Would suggest: Restart your shell or run 'source ~/.bashrc' to enable completions."
    else
        print_info "Restart your shell or run 'source ~/.bashrc' to enable completions."
    fi
}

# Function to setup Fish completions
setup_fish_completions() {
    fabric_cmd="$1"
    script_dir="$2"
    completion_file="${fabric_cmd}.fish"

    print_info "Setting up Fish completions for '$fabric_cmd'..."

    # Fish completion directory
    fish_dir="$HOME/.config/fish/completions"

    if [ "$DRY_RUN" = true ]; then
        print_dry_run "Would run: mkdir -p \"$fish_dir\""
        print_dry_run "Would run: cp \"$script_dir/fabric.fish\" \"$fish_dir/$completion_file\""
        print_success "Would install Fish completion to: $fish_dir/$completion_file"
        print_info "Fish will automatically load the completions (no restart needed)."
    elif mkdir -p "$fish_dir" 2>/dev/null; then
        if cp "$script_dir/fabric.fish" "$fish_dir/$completion_file"; then
            print_success "Installed Fish completion to: $fish_dir/$completion_file"
            print_info "Fish will automatically load the completions (no restart needed)."
        else
            print_error "Failed to copy Fish completion file."
            return 1
        fi
    else
        print_error "Failed to create Fish completions directory: $fish_dir"
        return 1
    fi
}

# Function to setup completions for other shells
setup_other_shell_completions() {
    fabric_cmd="$1"
    shell_name="$2"
    script_dir="$3"

    print_warning "Shell '$shell_name' is not directly supported."
    print_info "You can manually source the completion files:"
    print_info "  Bash-compatible: source $script_dir/fabric.bash"
    print_info "  Zsh-compatible: source $script_dir/_fabric"
}

# Function to show help
show_help() {
    cat << EOF
Fabric Shell Completions Setup Script

USAGE:
    setup-completions.sh [OPTIONS]

OPTIONS:
    --dry-run    Show what commands would be run without executing them
    --help       Show this help message

DESCRIPTION:
    This script automatically installs shell completions for the fabric CLI
    based on your current shell and the installed fabric command name.

        The script will use completion files from the same directory as the script
        when available. If they are not present (e.g., when running via curl), it
        will download them from GitHub:

            $FABRIC_COMPLETIONS_BASE_URL

        You can override the download source by setting
        FABRIC_COMPLETIONS_BASE_URL to your preferred location.

    Supports: zsh, bash, fish

    The script will:
    1. Detect whether 'fabric' or 'fabric-ai' is installed
    2. Detect your current shell from the SHELL environment variable
    3. Install the appropriate completion file with the correct name
    4. Try multiple standard completion directories

EXAMPLES:
        ./setup-completions.sh                  # Install completions
        ./setup-completions.sh --dry-run        # Show what would be done
        FABRIC_COMPLETIONS_BASE_URL="https://raw.githubusercontent.com/<owner>/<repo>/main/completions" \\
            ./setup-completions.sh               # Override download source
        ./setup-completions.sh --help           # Show this help

EOF
}

# Main function
main() {
    # Parse command line arguments
    while [ $# -gt 0 ]; do
        case "$1" in
            --dry-run)
                DRY_RUN=true
                shift
                ;;
            --help|-h)
                show_help
                exit 0
                ;;
            *)
                print_error "Unknown option: $1"
                print_info "Use --help for usage information."
                exit 1
                ;;
        esac
    done

    print_info "Fabric Shell Completions Setup"
    print_info "==============================="

    if [ "$DRY_RUN" = true ]; then
        print_info "DRY RUN MODE - Commands will be shown but not executed"
        print_info ""
    fi

    # Get script directory and obtain completion files (local or downloaded)
    script_dir="$(get_script_dir)"
    script_dir="$(obtain_completion_files "$script_dir" || echo "")"
    if [ -z "$script_dir" ]; then
        print_error "Unable to obtain completion files. Aborting."
        exit 1
    fi

    # If we downloaded into a temp dir, arrange cleanup at process exit
    if [ -n "$TEMP_DIR" ] && [ -d "$TEMP_DIR" ]; then
        trap 'if [ -n "$TEMP_DIR" ] && [ -d "$TEMP_DIR" ]; then rm -rf "$TEMP_DIR"; fi' EXIT INT TERM
    fi

    # Detect fabric command
    fabric_cmd="$(detect_fabric_command)"
    print_info "Detected fabric command: $fabric_cmd"

    # Detect shell
    shell_name="$(detect_shell)"
    print_info "Detected shell: $shell_name"

    # Setup completions based on shell
    case "$shell_name" in
        zsh)
            setup_zsh_completions "$fabric_cmd" "$script_dir"
            ;;
        bash)
            setup_bash_completions "$fabric_cmd" "$script_dir"
            ;;
        fish)
            setup_fish_completions "$fabric_cmd" "$script_dir"
            ;;
        *)
            setup_other_shell_completions "$fabric_cmd" "$shell_name" "$script_dir"
            ;;
    esac

    if [ "$DRY_RUN" = true ]; then
        print_success "Dry-run completed! The above commands would set up shell completions."
        print_info "Run without --dry-run to actually install the completions."
    else
        print_success "Shell completion setup completed!"
        print_info "You can now use tab completion with the '$fabric_cmd' command."
    fi
}

# Run main function
main "$@"
