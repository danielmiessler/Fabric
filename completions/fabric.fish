# Fish shell completion for fabric CLI
#
# Installation:
# Copy this file to ~/.config/fish/completions/fabric.fish
# or run:
# mkdir -p ~/.config/fish/completions
# cp completions/fabric.fish ~/.config/fish/completions/

# Helper functions for dynamic completions
function __fabric_get_patterns
	fabric --listpatterns --shell-complete-list 2>/dev/null
end

function __fabric_get_models
	fabric --listmodels --shell-complete-list 2>/dev/null
end

function __fabric_get_contexts
	fabric --listcontexts --shell-complete-list 2>/dev/null
end

function __fabric_get_sessions
	fabric --listsessions --shell-complete-list 2>/dev/null
end

function __fabric_get_strategies
	fabric --liststrategies --shell-complete-list 2>/dev/null
end

function __fabric_get_extensions
	fabric --listextensions --shell-complete-list 2>/dev/null
end

function __fabric_get_gemini_voices
	fabric --list-gemini-voices --shell-complete-list 2>/dev/null
end

# Main completion function
complete -c fabric -f

# Flag completions with arguments
complete -c fabric -s p -l pattern -d "Choose a pattern from the available patterns" -a "(__fabric_get_patterns)"
complete -c fabric -s v -l variable -d "Values for pattern variables, e.g. -v=#role:expert -v=#points:30"
complete -c fabric -s C -l context -d "Choose a context from the available contexts" -a "(__fabric_get_contexts)"
complete -c fabric -l session -d "Choose a session from the available sessions" -a "(__fabric_get_sessions)"
complete -c fabric -s a -l attachment -d "Attachment path or URL (e.g. for OpenAI image recognition messages)" -r
complete -c fabric -s t -l temperature -d "Set temperature (default: 0.7)"
complete -c fabric -s T -l topp -d "Set top P (default: 0.9)"
complete -c fabric -s P -l presencepenalty -d "Set presence penalty (default: 0.0)"
complete -c fabric -s F -l frequencypenalty -d "Set frequency penalty (default: 0.0)"
complete -c fabric -s m -l model -d "Choose model" -a "(__fabric_get_models)"
complete -c fabric -l modelContextLength -d "Model context length (only affects ollama)"
complete -c fabric -s o -l output -d "Output to file" -r
complete -c fabric -s n -l latest -d "Number of latest patterns to list (default: 0)"
complete -c fabric -s y -l youtube -d "YouTube video or play list URL to grab transcript, comments from it"
complete -c fabric -s g -l language -d "Specify the Language Code for the chat, e.g. -g=en -g=zh"
complete -c fabric -s u -l scrape_url -d "Scrape website URL to markdown using Jina AI"
complete -c fabric -s q -l scrape_question -d "Search question using Jina AI"
complete -c fabric -s e -l seed -d "Seed to be used for LMM generation"
complete -c fabric -s w -l wipecontext -d "Wipe context" -a "(__fabric_get_contexts)"
complete -c fabric -s W -l wipesession -d "Wipe session" -a "(__fabric_get_sessions)"
complete -c fabric -l printcontext -d "Print context" -a "(__fabric_get_contexts)"
complete -c fabric -l printsession -d "Print session" -a "(__fabric_get_sessions)"
complete -c fabric -l address -d "The address to bind the REST API (default: :8080)"
complete -c fabric -l api-key -d "API key used to secure server routes"
complete -c fabric -l config -d "Path to YAML config file" -r -a "*.yaml *.yml"
complete -c fabric -l search-location -d "Set location for web search results (e.g., 'America/Los_Angeles')"
complete -c fabric -l image-file -d "Save generated image to specified file path (e.g., 'output.png')" -r -a "*.png *.webp *.jpeg *.jpg"
complete -c fabric -l image-size -d "Image dimensions: 1024x1024, 1536x1024, 1024x1536, auto (default: auto)" -a "1024x1024 1536x1024 1024x1536 auto"
complete -c fabric -l image-quality -d "Image quality: low, medium, high, auto (default: auto)" -a "low medium high auto"
complete -c fabric -l image-compression -d "Compression level 0-100 for JPEG/WebP formats (default: not set)" -r
complete -c fabric -l image-background -d "Background type: opaque, transparent (default: opaque, only for PNG/WebP)" -a "opaque transparent"
complete -c fabric -l addextension -d "Register a new extension from config file path" -r -a "*.yaml *.yml"
complete -c fabric -l rmextension -d "Remove a registered extension by name" -a "(__fabric_get_extensions)"
complete -c fabric -l strategy -d "Choose a strategy from the available strategies" -a "(__fabric_get_strategies)"
complete -c fabric -l think-start-tag -d "Start tag for thinking sections (default: <think>)"
complete -c fabric -l think-end-tag -d "End tag for thinking sections (default: </think>)"
complete -c fabric -l voice -d "TTS voice name for supported models (e.g., Kore, Charon, Puck)" -a "(__fabric_get_gemini_voices)"

# Boolean flags (no arguments)
complete -c fabric -s S -l setup -d "Run setup for all reconfigurable parts of fabric"
complete -c fabric -s s -l stream -d "Stream"
complete -c fabric -s r -l raw -d "Use the defaults of the model without sending chat options"
complete -c fabric -s l -l listpatterns -d "List all patterns"
complete -c fabric -s L -l listmodels -d "List all available models"
complete -c fabric -s x -l listcontexts -d "List all contexts"
complete -c fabric -s X -l listsessions -d "List all sessions"
complete -c fabric -s U -l updatepatterns -d "Update patterns"
complete -c fabric -s c -l copy -d "Copy to clipboard"
complete -c fabric -l output-session -d "Output the entire session to the output file"
complete -c fabric -s d -l changeDefaultModel -d "Change default model"
complete -c fabric -l playlist -d "Prefer playlist over video if both ids are present in the URL"
complete -c fabric -l transcript -d "Grab transcript from YouTube video and send to chat"
complete -c fabric -l transcript-with-timestamps -d "Grab transcript from YouTube video with timestamps"
complete -c fabric -l comments -d "Grab comments from YouTube video and send to chat"
complete -c fabric -l metadata -d "Output video metadata"
complete -c fabric -l readability -d "Convert HTML input into a clean, readable view"
complete -c fabric -l input-has-vars -d "Apply variables to user input"
complete -c fabric -l dry-run -d "Show what would be sent to the model without actually sending it"
complete -c fabric -l search -d "Enable web search tool for supported models (Anthropic, OpenAI)"
complete -c fabric -l serve -d "Serve the Fabric Rest API"
complete -c fabric -l serveOllama -d "Serve the Fabric Rest API with ollama endpoints"
complete -c fabric -l version -d "Print current version"
complete -c fabric -l listextensions -d "List all registered extensions"
complete -c fabric -l liststrategies -d "List all strategies"
complete -c fabric -l listvendors -d "List all vendors"
complete -c fabric -l list-gemini-voices -d "List all available Gemini TTS voices"
complete -c fabric -l shell-complete-list -d "Output raw list without headers/formatting (for shell completion)"
complete -c fabric -l suppress-think -d "Suppress text enclosed in thinking tags"
complete -c fabric -l disable-responses-api -d "Disable OpenAI Responses API (default: false)"
complete -c fabric -s h -l help -d "Show this help message"
