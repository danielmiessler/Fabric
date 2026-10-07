# Bash completion for fabric CLI
#
# Installation:
# 1. Place this file in a standard completion directory, e.g.,
#    - /etc/bash_completion.d/
#    - /usr/local/etc/bash_completion.d/
#    - ~/.local/share/bash-completion/completions/
# 2. Or, source it directly in your ~/.bashrc or ~/.bash_profile:
#    source /path/to/fabric.bash

_fabric() {
  local cur prev words cword
  if declare -F _comp_get_words &>/dev/null; then
      _comp_get_words cur prev words cword
   else
      _get_comp_words_by_ref cur prev words cword
   fi

  # Define all possible options/flags
  local opts="--pattern -p --workflow --variable -v --context -C --session --attachment -a --setup -S --temperature -t --topp -T --stream -s --presencepenalty -P --raw -r --frequencypenalty -F --listpatterns -l --readpattern --listmodels -L --listcontexts -x --listsessions -X --updatepatterns -U --copy -c --model -m --vendor -V --modelContextLength --maxTokens --output -o --output-session --extract --extract-last --latest -n --changeDefaultModel -d --youtube -y --playlist --transcript --transcript-with-timestamps --visual --visual-sensitivity --visual-fps --comments --metadata --yt-dlp-args --spotify --language -g --scrape_url -u --scrape_question -q --serply_search --firecrawl_search --seed -e --thinking --wipecontext -w --wipesession -W --printcontext --printsession --readability --input-has-vars --no-variable-replacement --dry-run --print-prompt --serve --serveOllama --address --api-key --cors-origins --config --search --search-location --image-file --image-size --image-quality --image-compression --image-background --suppress-think --think-start-tag --think-end-tag --disable-responses-api --transcribe-file --transcribe-model --split-media-file --voice --list-gemini-voices --list-transcription-models --notification --notification-command --show-metadata --debug --version --listextensions --addextension --rmextension --strategy --liststrategies --listvendors --shell-complete-list --help -h"

  # Helper function for dynamic completions
  _fabric_get_list() {
    "${COMP_WORDS[0]}" "$1" --shell-complete-list 2>/dev/null
  }

  # Set COMPREPLY to the names from a dynamic list that start with $cur.
  #
  # SECURITY: do not give these names to `compgen -W "<list>"`. compgen expands
  # the word list, with command substitution, so a name that contains `$(...)`
  # runs as a command when the user pushes Tab. Read each name as a plain line
  # and compare it as literal text. Use `while read`, not mapfile, for bash 3.2.
  _fabric_complete_dynamic() {
    local _line
    COMPREPLY=()
    while IFS= read -r _line; do
      [[ -n ${_line} && ${_line} == "${cur}"* ]] && COMPREPLY+=("${_line}")
    done < <(_fabric_get_list "$1")
  }

  # Handle completions based on the previous word
  case "${prev}" in
  -p | --pattern | --readpattern)
    _fabric_complete_dynamic --listpatterns
    return 0
    ;;
  -C | --context)
    _fabric_complete_dynamic --listcontexts
    return 0
    ;;
  --session)
    _fabric_complete_dynamic --listsessions
    return 0
    ;;
  -m | --model)
    _fabric_complete_dynamic --listmodels
    return 0
    ;;
  -V | --vendor)
    _fabric_complete_dynamic --listvendors
    return 0
    ;;
  -w | --wipecontext)
    _fabric_complete_dynamic --listcontexts
    return 0
    ;;
  -W | --wipesession)
    _fabric_complete_dynamic --listsessions
    return 0
    ;;
  --printcontext)
    _fabric_complete_dynamic --listcontexts
    return 0
    ;;
  --printsession)
    _fabric_complete_dynamic --listsessions
    return 0
    ;;
  --thinking)
    COMPREPLY=($(compgen -W "off low medium high" -- "${cur}"))
    return 0
    ;;
  --rmextension)
    _fabric_complete_dynamic --listextensions
    return 0
    ;;
  --strategy)
    _fabric_complete_dynamic --liststrategies
    return 0
    ;;
  --voice)
    _fabric_complete_dynamic --list-gemini-voices
    return 0
    ;;
  --transcribe-model)
    _fabric_complete_dynamic --list-transcription-models
    return 0
    ;;
  --debug)
    COMPREPLY=($(compgen -W "0 1 2 3 4" -- "${cur}"))
    return 0
    ;;
  # Options requiring file/directory paths
  -a | --attachment | -o | --output | --config | --workflow | --addextension | --image-file | --transcribe-file)
    _filedir
    return 0
    ;;
  # Image generation options with specific values
  --image-size)
    COMPREPLY=($(compgen -W "1024x1024 1536x1024 1024x1536 auto" -- "$cur"))
    return 0
    ;;
  --image-quality)
    COMPREPLY=($(compgen -W "low medium high auto" -- "$cur"))
    return 0
    ;;
  --image-background)
    COMPREPLY=($(compgen -W "opaque transparent" -- "$cur"))
    return 0
    ;;
  # Options requiring simple arguments (no specific completion logic here)
  -v | --variable | -t | --temperature | -T | --topp | -P | --presencepenalty | -F | --frequencypenalty | --modelContextLength | --maxTokens | -n | --latest | -y | --youtube | --visual-sensitivity | --visual-fps | --yt-dlp-args | -g | --language | -u | --scrape_url | -q | --scrape_question | --serply_search | --firecrawl_search | -e | --seed | --address | --api-key | --cors-origins | --search-location | --image-compression | --think-start-tag | --think-end-tag | --notification-command)
    # No specific completion suggestions, user types the value
    return 0
    ;;
  --spotify)
    return 0
    ;;
  esac

  # If the current word starts with '-', suggest options
  if [[ "${cur}" == -* ]]; then
    COMPREPLY=($(compgen -W "${opts}" -- "${cur}"))
    return 0
  fi

  # Default: complete files/directories if no other rule matches
  # _filedir
  # Or provide no completions if it's not an option or argument following a known flag
  COMPREPLY=()

}

complete -F _fabric fabric fabric-ai
