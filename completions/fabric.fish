# Fish shell completion for fabric CLI
#
# Installation:
# Copy this file to ~/.config/fish/completions/fabric.fish
# or run:
# mkdir -p ~/.config/fish/completions
# cp completions/fabric.fish ~/.config/fish/completions/

# Helper functions for dynamic completions
function __fabric_get_patterns
        set cmd (commandline -opc)[1]
        $cmd --listpatterns --shell-complete-list 2>/dev/null
end

function __fabric_get_models
        set cmd (commandline -opc)[1]
        $cmd --listmodels --shell-complete-list 2>/dev/null
end

function __fabric_get_vendors
        set cmd (commandline -opc)[1]
        $cmd --listvendors --shell-complete-list 2>/dev/null
end

function __fabric_get_contexts
        set cmd (commandline -opc)[1]
        $cmd --listcontexts --shell-complete-list 2>/dev/null
end

function __fabric_get_sessions
        set cmd (commandline -opc)[1]
        $cmd --listsessions --shell-complete-list 2>/dev/null
end

function __fabric_get_strategies
        set cmd (commandline -opc)[1]
        $cmd --liststrategies --shell-complete-list 2>/dev/null
end

function __fabric_get_extensions
        set cmd (commandline -opc)[1]
        $cmd --listextensions --shell-complete-list 2>/dev/null
end

function __fabric_get_gemini_voices
        set cmd (commandline -opc)[1]
        $cmd --list-gemini-voices --shell-complete-list 2>/dev/null
end

function __fabric_get_transcription_models
        set cmd (commandline -opc)[1]
        $cmd --list-transcription-models --shell-complete-list 2>/dev/null
end

# Main completion function
function __fabric_register_completions
        set cmd $argv[1]
        complete -c $cmd -f

        # Flag completions with arguments
        complete -c $cmd -s p -l pattern -d "Choose a pattern from the available patterns" -a "(__fabric_get_patterns)"
        complete -c $cmd -s v -l variable -d "Values for pattern variables, e.g. -v=#role:expert -v=#points:30"
        complete -c $cmd -s C -l context -d "Choose a context from the available contexts" -a "(__fabric_get_contexts)"
        complete -c $cmd -l session -d "Choose a session from the available sessions" -a "(__fabric_get_sessions)"
        complete -c $cmd -s a -l attachment -d "Attachment path or URL (e.g. for OpenAI image recognition messages)" -r
        complete -c $cmd -s t -l temperature -d "Set temperature (default: 0.7)"
        complete -c $cmd -s T -l topp -d "Set top P (default: 0.9)"
        complete -c $cmd -s P -l presencepenalty -d "Set presence penalty (default: 0.0)"
        complete -c $cmd -s F -l frequencypenalty -d "Set frequency penalty (default: 0.0)"
        complete -c $cmd -s m -l model -d "Choose model" -a "(__fabric_get_models)"
        complete -c $cmd -s V -l vendor -d "Specify vendor for chosen model (e.g., -V \"LM Studio\" -m openai/gpt-oss-20b)" -a "(__fabric_get_vendors)"
        complete -c $cmd -l modelContextLength -d "Model context length (only affects ollama)"
        complete -c $cmd -s o -l output -d "Output to file" -r
        complete -c $cmd -s n -l latest -d "Number of latest patterns to list (default: 0)"
        complete -c $cmd -s y -l youtube -d "YouTube video or play list URL to grab transcript, comments from it"
        complete -c $cmd -s g -l language -d "Specify the Language Code for the chat, e.g. -g=en -g=zh"
        complete -c $cmd -s u -l scrape_url -d "Scrape website URL to markdown using Jina AI"
        complete -c $cmd -s q -l scrape_question -d "Search question using Jina AI"
        complete -c $cmd -s e -l seed -d "Seed to be used for LMM generation"
        complete -c $cmd -l thinking -d "Set reasoning/thinking level" -a "off low medium high"
        complete -c $cmd -s w -l wipecontext -d "Wipe context" -a "(__fabric_get_contexts)"
        complete -c $cmd -s W -l wipesession -d "Wipe session" -a "(__fabric_get_sessions)"
        complete -c $cmd -l printcontext -d "Print context" -a "(__fabric_get_contexts)"
        complete -c $cmd -l printsession -d "Print session" -a "(__fabric_get_sessions)"
        complete -c $cmd -l address -d "The address to bind the REST API (default: :8080)"
        complete -c $cmd -l api-key -d "API key used to secure server routes"
        complete -c $cmd -l config -d "Path to YAML config file" -r -a "*.yaml *.yml"
        complete -c $cmd -l search-location -d "Set location for web search results (e.g., 'America/Los_Angeles')"
        complete -c $cmd -l image-file -d "Save generated image to specified file path (e.g., 'output.png')" -r -a "*.png *.webp *.jpeg *.jpg"
        complete -c $cmd -l image-size -d "Image dimensions: 1024x1024, 1536x1024, 1024x1536, auto (default: auto)" -a "1024x1024 1536x1024 1024x1536 auto"
        complete -c $cmd -l image-quality -d "Image quality: low, medium, high, auto (default: auto)" -a "low medium high auto"
        complete -c $cmd -l image-compression -d "Compression level 0-100 for JPEG/WebP formats (default: not set)" -r
        complete -c $cmd -l image-background -d "Background type: opaque, transparent (default: opaque, only for PNG/WebP)" -a "opaque transparent"
        complete -c $cmd -l addextension -d "Register a new extension from config file path" -r -a "*.yaml *.yml"
        complete -c $cmd -l rmextension -d "Remove a registered extension by name" -a "(__fabric_get_extensions)"
        complete -c $cmd -l strategy -d "Choose a strategy from the available strategies" -a "(__fabric_get_strategies)"
        complete -c $cmd -l think-start-tag -d "Start tag for thinking sections (default: <think>)"
        complete -c $cmd -l think-end-tag -d "End tag for thinking sections (default: </think>)"
        complete -c $cmd -l voice -d "TTS voice name for supported models (e.g., Kore, Charon, Puck)" -a "(__fabric_get_gemini_voices)"
        complete -c $cmd -l transcribe-file -d "Audio or video file to transcribe" -r -a "*.mp3 *.mp4 *.mpeg *.mpga *.m4a *.wav *.webm"
        complete -c $cmd -l transcribe-model -d "Model to use for transcription (separate from chat model)" -a "(__fabric_get_transcription_models)"
        complete -c $cmd -l debug -d "Set debug level (0=off, 1=basic, 2=detailed, 3=trace)" -a "0 1 2 3"
        complete -c $cmd -l notification-command -d "Custom command to run for notifications (overrides built-in notifications)"

        # Boolean flags (no arguments)
        complete -c $cmd -s S -l setup -d "Run setup for all reconfigurable parts of fabric"
        complete -c $cmd -s s -l stream -d "Stream"
        complete -c $cmd -s r -l raw -d "Use the defaults of the model without sending chat options"
        complete -c $cmd -s l -l listpatterns -d "List all patterns"
        complete -c $cmd -s L -l listmodels -d "List all available models"
        complete -c $cmd -s x -l listcontexts -d "List all contexts"
        complete -c $cmd -s X -l listsessions -d "List all sessions"
        complete -c $cmd -s U -l updatepatterns -d "Update patterns"
        complete -c $cmd -s c -l copy -d "Copy to clipboard"
        complete -c $cmd -l output-session -d "Output the entire session to the output file"
        complete -c $cmd -s d -l changeDefaultModel -d "Change default model"
        complete -c $cmd -l playlist -d "Prefer playlist over video if both ids are present in the URL"
        complete -c $cmd -l transcript -d "Grab transcript from YouTube video and send to chat"
        complete -c $cmd -l transcript-with-timestamps -d "Grab transcript from YouTube video with timestamps"
        complete -c $cmd -l comments -d "Grab comments from YouTube video and send to chat"
        complete -c $cmd -l metadata -d "Output video metadata"
        complete -c $cmd -l yt-dlp-args -d "Additional arguments to pass to yt-dlp (e.g. '--cookies-from-browser brave')"
        complete -c $cmd -l readability -d "Convert HTML input into a clean, readable view"
       complete -c $cmd -l input-has-vars -d "Apply variables to user input"
       complete -c $cmd -l no-variable-replacement -d "Disable pattern variable replacement"
       complete -c $cmd -l dry-run -d "Show what would be sent to the model without actually sending it"
        complete -c $cmd -l search -d "Enable web search tool for supported models (Anthropic, OpenAI, Gemini)"
        complete -c $cmd -l serve -d "Serve the Fabric Rest API"
        complete -c $cmd -l serveOllama -d "Serve the Fabric Rest API with ollama endpoints"
        complete -c $cmd -l version -d "Print current version"
        complete -c $cmd -l listextensions -d "List all registered extensions"
        complete -c $cmd -l liststrategies -d "List all strategies"
        complete -c $cmd -l listvendors -d "List all vendors"
        complete -c $cmd -l list-gemini-voices -d "List all available Gemini TTS voices"
        complete -c $cmd -l shell-complete-list -d "Output raw list without headers/formatting (for shell completion)"
        complete -c $cmd -l suppress-think -d "Suppress text enclosed in thinking tags"
        complete -c $cmd -l disable-responses-api -d "Disable OpenAI Responses API (default: false)"
        complete -c $cmd -l split-media-file -d "Split audio/video files larger than 25MB using ffmpeg"
        complete -c $cmd -l notification -d "Send desktop notification when command completes"
        complete -c $cmd -s h -l help -d "Show this help message"
end

__fabric_register_completions fabric
__fabric_register_completions fabric-ai
