package gemini

import (
	"fmt"
	"sort"
)

// GeminiVoice represents a Gemini TTS voice with its characteristics
type GeminiVoice struct {
	Name            string
	Description     string
	Characteristics []string
}

// GetGeminiVoices returns the current list of supported Gemini TTS voices
// This list is maintained based on official Google Gemini documentation
// https://ai.google.dev/gemini-api/docs/speech-generation
func GetGeminiVoices() []GeminiVoice {
	return []GeminiVoice{
		// Firm voices
		{Name: "Kore", Description: "Firm and confident", Characteristics: []string{"firm", "confident", "default"}},
		{Name: "Orus", Description: "Firm and decisive", Characteristics: []string{"firm", "decisive"}},
		{Name: "Alnilam", Description: "Firm and strong", Characteristics: []string{"firm", "strong"}},

		// Upbeat voices
		{Name: "Puck", Description: "Upbeat and energetic", Characteristics: []string{"upbeat", "energetic"}},
		{Name: "Laomedeia", Description: "Upbeat and lively", Characteristics: []string{"upbeat", "lively"}},

		// Bright voices
		{Name: "Zephyr", Description: "Bright and cheerful", Characteristics: []string{"bright", "cheerful"}},
		{Name: "Autonoe", Description: "Bright and optimistic", Characteristics: []string{"bright", "optimistic"}},

		// Informative voices
		{Name: "Charon", Description: "Informative and clear", Characteristics: []string{"informative", "clear"}},
		{Name: "Rasalgethi", Description: "Informative and professional", Characteristics: []string{"informative", "professional"}},

		// Natural voices
		{Name: "Aoede", Description: "Breezy and natural", Characteristics: []string{"breezy", "natural"}},
		{Name: "Leda", Description: "Youthful and energetic", Characteristics: []string{"youthful", "energetic"}},

		// Gentle voices
		{Name: "Vindemiatrix", Description: "Gentle and kind", Characteristics: []string{"gentle", "kind"}},
		{Name: "Achernar", Description: "Soft and gentle", Characteristics: []string{"soft", "gentle"}},
		{Name: "Enceladus", Description: "Breathy and soft", Characteristics: []string{"breathy", "soft"}},

		// Warm voices
		{Name: "Sulafat", Description: "Warm and welcoming", Characteristics: []string{"warm", "welcoming"}},
		{Name: "Capella", Description: "Warm and approachable", Characteristics: []string{"warm", "approachable"}},

		// Clear voices
		{Name: "Iapetus", Description: "Clear and articulate", Characteristics: []string{"clear", "articulate"}},
		{Name: "Erinome", Description: "Clear and precise", Characteristics: []string{"clear", "precise"}},

		// Pleasant voices
		{Name: "Algieba", Description: "Smooth and pleasant", Characteristics: []string{"smooth", "pleasant"}},
		{Name: "Vega", Description: "Smooth and flowing", Characteristics: []string{"smooth", "flowing"}},

		// Textured voices
		{Name: "Algenib", Description: "Gravelly texture", Characteristics: []string{"gravelly", "textured"}},

		// Relaxed voices
		{Name: "Callirrhoe", Description: "Easy-going and relaxed", Characteristics: []string{"relaxed", "easy-going"}},
		{Name: "Despina", Description: "Calm and serene", Characteristics: []string{"calm", "serene"}},

		// Mature voices
		{Name: "Gacrux", Description: "Mature and experienced", Characteristics: []string{"mature", "experienced"}},

		// Expressive voices
		{Name: "Pulcherrima", Description: "Forward and expressive", Characteristics: []string{"forward", "expressive"}},
		{Name: "Lyra", Description: "Melodic and expressive", Characteristics: []string{"melodic", "expressive"}},

		// Dynamic voices
		{Name: "Fenrir", Description: "Excitable and dynamic", Characteristics: []string{"excitable", "dynamic"}},
		{Name: "Sadachbia", Description: "Lively and animated", Characteristics: []string{"lively", "animated"}},

		// Friendly voices
		{Name: "Achird", Description: "Friendly and approachable", Characteristics: []string{"friendly", "approachable"}},

		// Casual voices
		{Name: "Zubenelgenubi", Description: "Casual and conversational", Characteristics: []string{"casual", "conversational"}},

		// Additional voices from latest API
		{Name: "Sadaltager", Description: "Experimental voice with a calm and neutral tone", Characteristics: []string{"experimental", "calm", "neutral"}},
		{Name: "Schedar", Description: "Experimental voice with a warm and engaging tone", Characteristics: []string{"experimental", "warm", "engaging"}},
		{Name: "Umbriel", Description: "Experimental voice with a deep and resonant tone", Characteristics: []string{"experimental", "deep", "resonant"}},
	}
}

// GetGeminiVoiceNames returns just the voice names in alphabetical order
func GetGeminiVoiceNames() []string {
	voices := GetGeminiVoices()
	names := make([]string, len(voices))
	for i, voice := range voices {
		names[i] = voice.Name
	}
	sort.Strings(names)
	return names
}

// IsValidGeminiVoice checks if a voice name is valid
func IsValidGeminiVoice(voiceName string) bool {
	if voiceName == "" {
		return true // Empty voice is valid (will use default)
	}

	for _, voice := range GetGeminiVoices() {
		if voice.Name == voiceName {
			return true
		}
	}
	return false
}

// GetGeminiVoiceByName returns a specific voice by name
func GetGeminiVoiceByName(name string) (*GeminiVoice, error) {
	for _, voice := range GetGeminiVoices() {
		if voice.Name == name {
			return &voice, nil
		}
	}
	return nil, fmt.Errorf("voice '%s' not found", name)
}

// ListGeminiVoices formats the voice list for display
func ListGeminiVoices(shellCompleteMode bool) string {
	if shellCompleteMode {
		// For shell completion, just return voice names
		names := GetGeminiVoiceNames()
		result := ""
		for _, name := range names {
			result += name + "\n"
		}
		return result
	}

	// For human-readable output
	voices := GetGeminiVoices()
	result := "Available Gemini Text-to-Speech voices:\n\n"

	// Group by characteristics for better readability
	groups := map[string][]GeminiVoice{
		"Firm & Confident":     {},
		"Bright & Cheerful":    {},
		"Warm & Welcoming":     {},
		"Clear & Professional": {},
		"Natural & Expressive": {},
		"Other Voices":         {},
	}

	for _, voice := range voices {
		placed := false
		for _, char := range voice.Characteristics {
			switch char {
			case "firm", "confident", "decisive", "strong":
				if !placed {
					groups["Firm & Confident"] = append(groups["Firm & Confident"], voice)
					placed = true
				}
			case "bright", "cheerful", "upbeat", "energetic", "lively":
				if !placed {
					groups["Bright & Cheerful"] = append(groups["Bright & Cheerful"], voice)
					placed = true
				}
			case "warm", "welcoming", "friendly", "approachable":
				if !placed {
					groups["Warm & Welcoming"] = append(groups["Warm & Welcoming"], voice)
					placed = true
				}
			case "clear", "informative", "professional", "articulate":
				if !placed {
					groups["Clear & Professional"] = append(groups["Clear & Professional"], voice)
					placed = true
				}
			case "natural", "expressive", "melodic", "breezy":
				if !placed {
					groups["Natural & Expressive"] = append(groups["Natural & Expressive"], voice)
					placed = true
				}
			}
		}
		if !placed {
			groups["Other Voices"] = append(groups["Other Voices"], voice)
		}
	}

	// Output grouped voices
	for groupName, groupVoices := range groups {
		if len(groupVoices) > 0 {
			result += fmt.Sprintf("%s:\n", groupName)
			for _, voice := range groupVoices {
				defaultStr := ""
				if voice.Name == "Kore" {
					defaultStr = " (default)"
				}
				result += fmt.Sprintf("  %-15s - %s%s\n", voice.Name, voice.Description, defaultStr)
			}
			result += "\n"
		}
	}

	result += "Use --voice <voice_name> to select a specific voice.\n"
	result += "Example: fabric --voice Charon -m gemini-2.5-flash-preview-tts -o output.wav \"Hello world\"\n"

	return result
}

// NOTE: This implementation maintains a curated list based on official Google documentation.
// In the future, if Google provides a dynamic voice discovery API, this can be updated
// to make API calls for real-time voice discovery.
//
// The current approach ensures:
// 1. Fast response times (no API calls needed)
// 2. Reliable voice information with descriptions
// 3. Easy maintenance when new voices are added
// 4. Offline functionality
//
// To update voices: Monitor Google's Gemini TTS documentation at:
// https://ai.google.dev/gemini-api/docs/speech-generation
