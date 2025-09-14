package cli

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/danielmiessler/fabric/internal/i18n"
	"github.com/jessevdk/go-flags"
)

// flagDescriptionMap maps flag names to their i18n keys
var flagDescriptionMap = map[string]string{
	"pattern":                    "choose_pattern_from_available",
	"variable":                   "pattern_variables_help",
	"context":                    "choose_context_from_available",
	"session":                    "choose_session_from_available",
	"attachment":                 "attachment_path_or_url_help",
	"setup":                      "run_setup_for_reconfigurable_parts",
	"temperature":                "set_temperature",
	"topp":                       "set_top_p",
	"stream":                     "stream_help",
	"presencepenalty":            "set_presence_penalty",
	"raw":                        "use_model_defaults_raw_help",
	"frequencypenalty":           "set_frequency_penalty",
	"listpatterns":               "list_all_patterns",
	"listmodels":                 "list_all_available_models",
	"listcontexts":               "list_all_contexts",
	"listsessions":               "list_all_sessions",
	"updatepatterns":             "update_patterns",
	"copy":                       "copy_to_clipboard",
	"model":                      "choose_model",
	"vendor":                     "specify_vendor_for_model",
	"modelContextLength":         "model_context_length_ollama",
	"output":                     "output_to_file",
	"output-session":             "output_entire_session",
	"latest":                     "number_of_latest_patterns",
	"changeDefaultModel":         "change_default_model",
	"youtube":                    "youtube_url_help",
	"playlist":                   "prefer_playlist_over_video",
	"transcript":                 "grab_transcript_from_youtube",
	"transcript-with-timestamps": "grab_transcript_with_timestamps",
	"comments":                   "grab_comments_from_youtube",
	"metadata":                   "output_video_metadata",
	"yt-dlp-args":                "additional_yt_dlp_args",
	"language":                   "specify_language_code",
	"scrape_url":                 "scrape_website_url",
	"scrape_question":            "search_question_jina",
	"seed":                       "seed_for_lmm_generation",
	"wipecontext":                "wipe_context",
	"wipesession":                "wipe_session",
	"printcontext":               "print_context",
	"printsession":               "print_session",
	"readability":                "convert_html_readability",
	"input-has-vars":             "apply_variables_to_input",
	"no-variable-replacement":    "disable_pattern_variable_replacement",
	"dry-run":                    "show_dry_run",
	"serve":                      "serve_fabric_rest_api",
	"serveOllama":                "serve_fabric_api_ollama_endpoints",
	"address":                    "address_to_bind_rest_api",
	"api-key":                    "api_key_secure_server_routes",
	"config":                     "path_to_yaml_config",
	"version":                    "print_current_version",
	"listextensions":             "list_all_registered_extensions",
	"addextension":               "register_new_extension",
	"rmextension":                "remove_registered_extension",
	"strategy":                   "choose_strategy_from_available",
	"liststrategies":             "list_all_strategies",
	"listvendors":                "list_all_vendors",
	"shell-complete-list":        "output_raw_list_shell_completion",
	"search":                     "enable_web_search_tool",
	"search-location":            "set_location_web_search",
	"image-file":                 "save_generated_image_to_file",
	"image-size":                 "image_dimensions_help",
	"image-quality":              "image_quality_help",
	"image-compression":          "compression_level_jpeg_webp",
	"image-background":           "background_type_help",
	"suppress-think":             "suppress_thinking_tags",
	"think-start-tag":            "start_tag_thinking_sections",
	"think-end-tag":              "end_tag_thinking_sections",
	"disable-responses-api":      "disable_openai_responses_api",
	"transcribe-file":            "audio_video_file_transcribe",
	"transcribe-model":           "model_for_transcription",
	"split-media-file":           "split_media_files_ffmpeg",
	"voice":                      "tts_voice_name",
	"list-gemini-voices":         "list_gemini_tts_voices",
	"list-transcription-models":  "list_transcription_models",
	"notification":               "send_desktop_notification",
	"notification-command":       "custom_notification_command",
	"thinking":                   "set_reasoning_thinking_level",
	"debug":                      "set_debug_level",
}

// TranslatedHelpWriter provides custom help output with translated descriptions
type TranslatedHelpWriter struct {
	parser *flags.Parser
	writer io.Writer
}

// NewTranslatedHelpWriter creates a new help writer with translations
func NewTranslatedHelpWriter(parser *flags.Parser, writer io.Writer) *TranslatedHelpWriter {
	return &TranslatedHelpWriter{
		parser: parser,
		writer: writer,
	}
}

// WriteHelp writes the help output with translated flag descriptions
func (h *TranslatedHelpWriter) WriteHelp() {
	fmt.Fprintf(h.writer, "%s\n", i18n.T("usage_header"))
	fmt.Fprintf(h.writer, "  %s %s\n\n", h.parser.Name, i18n.T("options_placeholder"))

	fmt.Fprintf(h.writer, "%s\n", i18n.T("application_options_header"))
	h.writeAllFlags()

	fmt.Fprintf(h.writer, "\n%s\n", i18n.T("help_options_header"))
	fmt.Fprintf(h.writer, "  -h, --help                        %s\n", i18n.T("help_message"))
}

// getTranslatedDescription gets the translated description for a flag
func (h *TranslatedHelpWriter) getTranslatedDescription(flagName string) string {
	if i18nKey, exists := flagDescriptionMap[flagName]; exists {
		return i18n.T(i18nKey)
	}

	// Fallback 1: Try to get original description from struct tag
	if desc := h.getOriginalDescription(flagName); desc != "" {
		return desc
	}

	// Fallback 2: Provide a user-friendly default message
	return i18n.T("no_description_available")
}

// getOriginalDescription retrieves the original description from struct tags
func (h *TranslatedHelpWriter) getOriginalDescription(flagName string) string {
	flags := &Flags{}
	flagsType := reflect.TypeOf(flags).Elem()

	for i := 0; i < flagsType.NumField(); i++ {
		field := flagsType.Field(i)
		longTag := field.Tag.Get("long")

		if longTag == flagName {
			if description := field.Tag.Get("description"); description != "" {
				return description
			}
			break
		}
	}
	return ""
}

// CustomHelpHandler handles help output with translations
func CustomHelpHandler(parser *flags.Parser, writer io.Writer) {
	// Initialize i18n system with detected language if not already initialized
	ensureI18nInitialized()

	helpWriter := NewTranslatedHelpWriter(parser, writer)
	helpWriter.WriteHelp()
}

// ensureI18nInitialized initializes the i18n system if not already done
func ensureI18nInitialized() {
	// Try to detect language from command line args or environment
	lang := detectLanguageFromArgs()
	if lang == "" {
		// Try to detect from environment variables
		lang = detectLanguageFromEnv()
	}

	// Initialize i18n with detected language (or empty for system default)
	i18n.Init(lang)
}

// detectLanguageFromArgs looks for --language/-g flag in os.Args
func detectLanguageFromArgs() string {
	args := os.Args[1:]
	for i, arg := range args {
		if arg == "--language" || arg == "-g" || (runtime.GOOS == "windows" && arg == "/g") {
			if i+1 < len(args) {
				return args[i+1]
			}
		} else if strings.HasPrefix(arg, "--language=") {
			return strings.TrimPrefix(arg, "--language=")
		} else if strings.HasPrefix(arg, "-g=") {
			return strings.TrimPrefix(arg, "-g=")
		} else if runtime.GOOS == "windows" && strings.HasPrefix(arg, "/g:") {
			return strings.TrimPrefix(arg, "/g:")
		} else if runtime.GOOS == "windows" && strings.HasPrefix(arg, "/g=") {
			return strings.TrimPrefix(arg, "/g=")
		}
	}
	return ""
}

// detectLanguageFromEnv detects language from environment variables
func detectLanguageFromEnv() string {
	// Check standard locale environment variables
	envVars := []string{"LC_ALL", "LC_MESSAGES", "LANG"}
	for _, envVar := range envVars {
		if value := os.Getenv(envVar); value != "" {
			// Extract language code from locale (e.g., "es_ES.UTF-8" -> "es")
			if strings.Contains(value, "_") {
				return strings.Split(value, "_")[0]
			}
			if value != "C" && value != "POSIX" {
				return value
			}
		}
	}
	return ""
}

// writeAllFlags writes all flags with translated descriptions
func (h *TranslatedHelpWriter) writeAllFlags() {
	// Use direct reflection on the Flags struct to get all flag definitions
	flags := &Flags{}
	flagsType := reflect.TypeOf(flags).Elem()

	for i := 0; i < flagsType.NumField(); i++ {
		field := flagsType.Field(i)

		shortTag := field.Tag.Get("short")
		longTag := field.Tag.Get("long")
		defaultTag := field.Tag.Get("default")

		if longTag == "" {
			continue // Skip fields without long tags
		}

		// Get translated description
		description := h.getTranslatedDescription(longTag)

		// Format the flag line
		var flagLine strings.Builder
		flagLine.WriteString("  ")

		if shortTag != "" {
			flagLine.WriteString(fmt.Sprintf("-%s, ", shortTag))
		}

		flagLine.WriteString(fmt.Sprintf("--%s", longTag))

		// Add parameter indicator for non-boolean flags
		isBoolFlag := field.Type.Kind() == reflect.Bool ||
			strings.HasSuffix(longTag, "patterns") ||
			strings.HasSuffix(longTag, "models") ||
			strings.HasSuffix(longTag, "contexts") ||
			strings.HasSuffix(longTag, "sessions") ||
			strings.HasSuffix(longTag, "extensions") ||
			strings.HasSuffix(longTag, "strategies") ||
			strings.HasSuffix(longTag, "vendors") ||
			strings.HasSuffix(longTag, "voices") ||
			longTag == "setup" || longTag == "stream" || longTag == "raw" ||
			longTag == "copy" || longTag == "updatepatterns" ||
			longTag == "output-session" || longTag == "changeDefaultModel" ||
			longTag == "playlist" || longTag == "transcript" ||
			longTag == "transcript-with-timestamps" || longTag == "comments" ||
			longTag == "metadata" || longTag == "readability" ||
			longTag == "input-has-vars" || longTag == "no-variable-replacement" ||
			longTag == "dry-run" || longTag == "serve" || longTag == "serveOllama" ||
			longTag == "version" || longTag == "shell-complete-list" ||
			longTag == "search" || longTag == "suppress-think" ||
			longTag == "disable-responses-api" || longTag == "split-media-file" ||
			longTag == "notification"

		if !isBoolFlag {
			flagLine.WriteString("=")
		}

		// Pad to align descriptions
		flagStr := flagLine.String()
		padding := 34 - len(flagStr)
		if padding < 2 {
			padding = 2
		}

		fmt.Fprintf(h.writer, "%s%s%s", flagStr, strings.Repeat(" ", padding), description)

		// Add default value if present
		if defaultTag != "" && defaultTag != "0" && defaultTag != "false" {
			fmt.Fprintf(h.writer, " (default: %s)", defaultTag)
		}

		fmt.Fprintf(h.writer, "\n")
	}
}
