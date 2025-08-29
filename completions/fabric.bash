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
  _get_comp_words_by_ref -n : cur prev words cword

  # Define all possible options/flags
  local opts="--pattern -p --variable -v --context -C --session --attachment -a --setup -S --temperature -t --topp -T --stream -s --presencepenalty -P --raw -r --frequencypenalty -F --listpatterns -l --listmodels -L --listcontexts -x --listsessions -X --updatepatterns -U --copy -c --model -m --vendor -V --modelContextLength --output -o --output-session --latest -n --changeDefaultModel -d --youtube -y --playlist --transcript --transcript-with-timestamps --comments --metadata --yt-dlp-args --language -g --scrape_url -u --scrape_question -q --seed -e --thinking --wipecontext -w --wipesession -W --printcontext --printsession --readability --input-has-vars --no-variable-replacement --dry-run --serve --serveOllama --address --api-key --config --search --search-location --image-file --image-size --image-quality --image-compression --image-background --suppress-think --think-start-tag --think-end-tag --disable-responses-api --transcribe-file --transcribe-model --split-media-file --voice --list-gemini-voices --notification --notification-command --debug --version --listextensions --addextension --rmextension --strategy --liststrategies --listvendors --shell-complete-list --help -h"

  # Helper function for dynamic completions
  _fabric_get_list() {
    "${COMP_WORDS[0]}" "$1" --shell-complete-list 2>/dev/null
  }

  # Handle completions based on the previous word
  case "${prev}" in
  -p | --pattern)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listpatterns)" -- "${cur}"))
    return 0
    ;;
  -C | --context)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listcontexts)" -- "${cur}"))
    return 0
    ;;
  --session)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listsessions)" -- "${cur}"))
    return 0
    ;;
  -m | --model)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listmodels)" -- "${cur}"))
    return 0
    ;;
  -V | --vendor)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listvendors)" -- "${cur}"))
    return 0
    ;;
  -w | --wipecontext)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listcontexts)" -- "${cur}"))
    return 0
    ;;
  -W | --wipesession)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listsessions)" -- "${cur}"))
    return 0
    ;;
  --printcontext)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listcontexts)" -- "${cur}"))
    return 0
    ;;
  --printsession)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listsessions)" -- "${cur}"))
    return 0
    ;;
  --thinking)
    COMPREPLY=($(compgen -W "off low medium high" -- "${cur}"))
    return 0
    ;;
  --rmextension)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --listextensions)" -- "${cur}"))
    return 0
    ;;
  --strategy)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --liststrategies)" -- "${cur}"))
    return 0
    ;;
  --voice)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --list-gemini-voices)" -- "${cur}"))
    return 0
    ;;
  --transcribe-model)
    COMPREPLY=($(compgen -W "$(_fabric_get_list --list-transcription-models)" -- "${cur}"))
    return 0
    ;;
  --debug)
    COMPREPLY=($(compgen -W "0 1 2 3" -- "${cur}"))
    return 0
    ;;
  # Options requiring file/directory paths
  -a | --attachment | -o | --output | --config | --addextension | --image-file | --transcribe-file)
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
  -v | --variable | -t | --temperature | -T | --topp | -P | --presencepenalty | -F | --frequencypenalty | --modelContextLength | -n | --latest | -y | --youtube | --yt-dlp-args | -g | --language | -u | --scrape_url | -q | --scrape_question | -e | --seed | --address | --api-key | --search-location | --image-compression | --think-start-tag | --think-end-tag | --notification-command)
    # No specific completion suggestions, user types the value
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
